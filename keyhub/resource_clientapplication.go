package keyhub

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	keyhubclient "github.com/topicuskeyhub/go-keyhub"
	keyhubmodel "github.com/topicuskeyhub/go-keyhub/model"
	"strconv"
)

func resourceClientApplication() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceClientApplicationCreate,
		UpdateContext: resourceClientApplicationUpdate,
		DeleteContext: resourceClientApplicationDelete,
		ReadContext:   resourceClientApplicationRead,
		Schema:        ClientApplicationResourceSchema(),
		Importer: &schema.ResourceImporter{
			StateContext: resourceClientApplicationImportContext,
		},
	}

}

func ClientApplicationResourceSchema() map[string]*schema.Schema {
	resourceSchema := map[string]*schema.Schema{
		"id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The value of the ID field of the client application",
		},
		"uuid": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The UUID of the created client application",
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name of the created client application",
		},

		"url": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "",
		},

		"owner": {
			Type:             schema.TypeString,
			Required:         true,
			ValidateDiagFunc: validation.ToDiagFunc(validation.IsUUID),
			Description:      "The UUID of the group that is owner of the client",
		},
		"technical_administrator": {
			Type:             schema.TypeString,
			Optional:         true,
			Computed:         true,
			ValidateDiagFunc: validation.ToDiagFunc(validation.IsUUID),
			Description:      "The UUID of the group that is technical administrator of the client, default to Owner",
		},

		"type": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The type of the client application. Possible values: `OAUTH2` (default), `LDAP`, `SAML2`",
			ValidateDiagFunc: validation.ToDiagFunc(
				validation.StringInSlice(
					[]string{
						string(keyhubmodel.CLIENT_TYPE_OAUTH2),
						string(keyhubmodel.CLIENT_TYPE_LDAP),
						string(keyhubmodel.CLIENT_TYPE_SAML2),
					},
					false,
				),
			),
			Default: string(keyhubmodel.CLIENT_TYPE_OAUTH2),
		},

		"scopes": {
			Type:        schema.TypeList,
			Optional:    true,
			Computed:    true,
			Description: "The Allowed scopes for the client application. For SSO applications this defaults to `profile`",
			Elem: &schema.Schema{
				Type: schema.TypeString,
				ValidateDiagFunc: validation.ToDiagFunc(
					validation.StringInSlice(
						[]string{
							keyhubmodel.CLIENT_SCOPE_PROFILE,
							keyhubmodel.CLIENT_SCOPE_MANAGE_ACCOUNT,
							keyhubmodel.CLIENT_SCOPE_PROVISIONING,
							keyhubmodel.CLIENT_SCOPE_ACCESS_VAULT,
							keyhubmodel.CLIENT_SCOPE_GROUP_ADMIN,
							keyhubmodel.CLIENT_SCOPE_GLOBAL_ADMIN,
						},
						false,
					),
				),
			},
		},

		"clientid": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The clientid (UUID) of the client application.",
		},
		"clientsecret": {
			Type:        schema.TypeString,
			Computed:    true,
			Sensitive:   true,
			Description: "The secret (UUID) of the client application.",
		},

		"is_sso": {
			Type:        schema.TypeBool,
			Optional:    true,
			Computed:    true,
			Description: "If set the client application is a Single Sign On application",
		},

		"attribute": {
			Type:        schema.TypeSet,
			Optional:    true,
			Computed:    true,
			Description: "The additional attributes which can be retrieved through userinfo ",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"name": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "The name of the attribute",
					},
					"script": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "The script defines the body of a `function (account, groups) { <SCRIPT> }` function, and is written in JavaScript using the ECMAScript 5 standard",
					},
				},
			},
		},

		"confidential": {
			Type:        schema.TypeBool,
			Optional:    true,
			Computed:    true,
			Description: "A confidential (or trusted) OAUTH2 client is able to keep its credentials a secret",
		},
		"is_server2server": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "The client is a code and implicit client",
		},
		"callback_uri": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The client is a code and implicit client",
		},
		"initiate_login_uri": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The URI within the application where a third party login can be started.",
		},
		"id_token_claims": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "A space-separated list of claims. These claims are added to the 'id_token', even if the client does not explicitly request them to be added",
		},
		"show_landingpage": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Displays a landing page immediately after login and before redirecting to the SSO application",
		},
		"binddn": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The identifier of the application to be used by a simple bind",
		},
		"used_for_provisioning": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "",
		},
		"client_certificate": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "",
		},
		"metadata_url": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The URL to retrieve the SAML metadata from",
		},
		"metadata": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The metadata of the SAML application, use if metadata_url is not available ",
		},
		"subject_format": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Topicus KeyHub can deliver the subject in three formats: Primary identifier, UPN and username.",
			ValidateDiagFunc: validation.ToDiagFunc(
				validation.StringInSlice(
					[]string{
						keyhubmodel.CLIENT_SUBJECT_FORMAT_UPN,
						keyhubmodel.CLIENT_SUBJECT_FORMAT_ID,
						keyhubmodel.CLIENT_SUBJECT_FORMAT_USERNAME,
						keyhubmodel.CLIENT_SUBJECT_FORMAT_EMAIL,
					},
					false,
				),
			),
		},
		"segments": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "",
		},
	}

	return resourceSchema
}

func resourceClientApplicationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*keyhubclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	var clientApp *keyhubmodel.ClientApplication
	var err error

	if id, ok := d.GetOk("id"); ok {

		var intId int64
		intId, err = strconv.ParseInt(id.(string), 10, 64)
		if err != nil {
			return append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Could not parse Id",
				Detail:   fmt.Sprintf("Id value %s", id),
			})
		}

		clientApp, err = client.ClientApplications.GetById(intId)
		if err != nil {
			return append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Could not GET ClientApplication " + d.Get("id").(string),
				Detail:   fmt.Sprintf("This might mean the ClientApplication has been deleted from keyhub, \nfor more info see the raw error: %s", err.Error()),
			})
		}

	} else {
		clientUuid, err := uuid.Parse(d.Get("uuid").(string))
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "uuid is not a valid uuid",
				Detail:   err.Error(),
			})
			return diags
		}

		clientApp, err = client.ClientApplications.GetByUUID(clientUuid)
		if err != nil {
			tflog.Debug(ctx, err.Error(), apiErrorToLogFields(err))
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Could not GET ClientApplication " + d.Get("id").(string),
				Detail:   fmt.Sprintf("This might mean the ClientApplication has been deleted from keyhub, \nfor more info see the raw error: %s", err.Error()),
			})
			return diags
		}
	}

	if err := d.Set("uuid", clientApp.UUID); err != nil {
		diags = append(diags, NewDiagnosticSetError("uuid", err))
	}

	// Set parameters

	if err := d.Set("type", string(clientApp.Type)); err != nil {
		diags = append(diags, NewDiagnosticSetError("type", err))
	}

	if err := d.Set("name", clientApp.Name); err != nil {
		diags = append(diags, NewDiagnosticSetError("name", err))
	}

	if clientApp.Owner != nil {
		if err := d.Set("owner", clientApp.Owner.UUID); err != nil {
			diags = append(diags, NewDiagnosticSetError("owner", err))
		}
	}

	if clientApp.TechnicalAdministrator != nil {
		if err := d.Set("technical_administrator", clientApp.TechnicalAdministrator.UUID); err != nil {
			diags = append(diags, NewDiagnosticSetError("technical_administrator", err))
		}
	}

	if err := d.Set("scopes", clientApp.Scopes); err != nil {
		diags = append(diags, NewDiagnosticSetError("scopes", err))
	}

	if err := d.Set("clientid", clientApp.ClientId); err != nil {
		diags = append(diags, NewDiagnosticSetError("clientid", err))
	}
	if secret, err := clientApp.GetSecret(); err == nil {
		if err := d.Set("clientsecret", secret); err != nil {
			diags = append(diags, NewDiagnosticSetError("clientid", err))
		}
	}
	if err := d.Set("is_sso", clientApp.SSOApplication); err != nil {
		diags = append(diags, NewDiagnosticSetError("is_sso", err))
	}

	if len(clientApp.Attributes) > 0 {
		var attributes []map[string]interface{}

		for name, script := range clientApp.Attributes {

			attribute := make(map[string]interface{})
			attribute["name"] = name
			attribute["script"] = script

			attributes = append(attributes, attribute)

		}
		if err := d.Set("attribute", attributes); err != nil {
			diags = append(diags, NewDiagnosticSetError("attribute", err))
		}
	}

	if err := d.Set("confidential", clientApp.Confidential); err != nil {
		diags = append(diags, NewDiagnosticSetError("confidential", err))
	}

	if err := d.Set("is_server2server", clientApp.IsOAuth2Server2Server()); err != nil {
		diags = append(diags, NewDiagnosticSetError("is_server2server", err))
	}

	if err := d.Set("callback_uri", clientApp.CallbackURI); err != nil {
		diags = append(diags, NewDiagnosticSetError("callback_uri", err))
	}

	if err := d.Set("initiate_login_uri", clientApp.InitiateLoginURI); err != nil {
		diags = append(diags, NewDiagnosticSetError("initiate_login_uri", err))
	}

	if err := d.Set("id_token_claims", clientApp.IdTokenClaims); err != nil {
		diags = append(diags, NewDiagnosticSetError("id_token_claims", err))
	}

	if err := d.Set("show_landingpage", clientApp.ShowLandingPage); err != nil {
		diags = append(diags, NewDiagnosticSetError("show_landingpage", err))
	}

	if err := d.Set("binddn", clientApp.BindDN); err != nil {
		diags = append(diags, NewDiagnosticSetError("binddn", err))
	}
	if err := d.Set("used_for_provisioning", clientApp.UsedForProvisioning); err != nil {
		diags = append(diags, NewDiagnosticSetError("used_for_provisioning", err))
	}
	if err := d.Set("client_certificate", clientApp.ClientCertificate); err != nil {
		diags = append(diags, NewDiagnosticSetError("client_certificate", err))
	}
	if err := d.Set("metadata_url", clientApp.MetadataUrl); err != nil {
		diags = append(diags, NewDiagnosticSetError("metadata_url", err))
	}
	if err := d.Set("metadata", clientApp.Metadata); err != nil {
		diags = append(diags, NewDiagnosticSetError("metadata", err))
	}
	if err := d.Set("subject_format", clientApp.SubjectFormat); err != nil {
		diags = append(diags, NewDiagnosticSetError("subject_format", err))
	}
	if err := d.Set("segments", clientApp.Segments); err != nil {
		diags = append(diags, NewDiagnosticSetError("segments", err))
	}

	d.SetId(strconv.FormatInt(clientApp.Self().ID, 10))

	return diags
}

func resourceClientApplicationCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*keyhubclient.Client)
	_ = client

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	var err error

	var uuidOwner uuid.UUID
	var owner *keyhubmodel.Group

	var clientApp *keyhubmodel.ClientApplication

	name := d.Get("name") // Field is required by scheme

	ownerUuid := d.Get("owner")
	uuidOwner, err = uuid.Parse(ownerUuid.(string))
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Invalid uuid",
			Detail:   fmt.Sprintf("Value `%s` is not a valid uuid for `%s`", ownerUuid.(string), "owner"),
		})
		return diags
	}

	owner, err = client.Groups.GetByUUID(uuidOwner)
	if err != nil {
		tflog.Debug(ctx, err.Error(), apiErrorToLogFields(err))
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Owner does not exist",
			Detail:   fmt.Sprintf("Could not find group with uuid: %s", ownerUuid.(string)),
		})
	}

	typeName := d.Get("type")
	switch typeName.(string) {
	case string(keyhubmodel.CLIENT_TYPE_OAUTH2):
		clientApp = keyhubmodel.NewOAuth2ClientApplication(name.(string), owner)
	case string(keyhubmodel.CLIENT_TYPE_SAML2):
		clientApp = keyhubmodel.NewSaml2ClientApplication(name.(string), owner)
	case string(keyhubmodel.CLIENT_TYPE_LDAP):
		clientApp = keyhubmodel.NewLdapClientApplication(name.(string), owner)
	}

	//technicaladministrator

	if technicalAdministratorUuid, ok := d.GetOk("technical_administrator"); ok {
		var uuidTechnicalAdministrator uuid.UUID
		var technicalAdministrator *keyhubmodel.Group
		uuidTechnicalAdministrator, err = uuid.Parse(technicalAdministratorUuid.(string))
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Invalid uuid",
				Detail:   fmt.Sprintf("Value `%s` is not a valid uuid for `%s`", ownerUuid.(string), "owner"),
			})
			return diags
		}

		technicalAdministrator, err = client.Groups.GetByUUID(uuidTechnicalAdministrator)
		if err != nil {
			tflog.Debug(ctx, err.Error(), apiErrorToLogFields(err))
			return append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "TechnicalAdministrator group does not exist",
				Detail:   fmt.Sprintf("Could not find group with uuid: %s", ownerUuid.(string)),
			})
		}

		clientApp.TechnicalAdministrator = technicalAdministrator.AsPrimer()
	}
	if value, ok := d.GetOk("is_server2server"); ok && value.(bool) {
		err := clientApp.SetOAuth2Server2Server()
		if err != nil {
			return append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Could not set Server2Server",
				Detail:   fmt.Sprintf("Error: %s", err.Error()),
			})
		}
	}

	if value, ok := d.GetOk("is_sso"); ok && value.(bool) {
		err := clientApp.SetOAuth2SSO()
		if err != nil {
			return append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Could not set SSO",
				Detail:   fmt.Sprintf("Error: %s", err.Error()),
			})
		}
	}

	if attributes, ok := d.GetOk("attribute"); ok {
		set := attributes.(*schema.Set)

		if set.Len() > 0 {
			for _, rawAttribute := range set.List() {
				attribute := rawAttribute.(map[string]interface{})
				err := clientApp.AddAttribute(attribute["name"].(string), attribute["script"].(string))
				if err != nil {
					return append(diags, diag.Diagnostic{
						Severity: diag.Error,
						Summary:  "Could not set attribute",
						Detail:   fmt.Sprintf("Setting the attribute '%s' returned an error: %s", attribute["name"].(string), err.Error()),
					})
				}
			}
		}
	}

	if value, ok := d.GetOk("scopes"); ok {
		clientApp.Scopes = []string{}
		for _, scope := range value.([]interface{}) {
			clientApp.Scopes = append(clientApp.Scopes, scope.(string))
		}
	}

	if value, ok := d.GetOk("url"); ok {
		clientApp.URL = value.(string)
	}

	if value, ok := d.GetOk("confidential"); ok {
		clientApp.Confidential = value.(bool)
	}

	if value, ok := d.GetOk("callback_uri"); ok {
		clientApp.CallbackURI = value.(string)
	}

	if value, ok := d.GetOk("initiate_login_uri"); ok {
		clientApp.InitiateLoginURI = value.(string)
	}

	if value, ok := d.GetOk("id_token_claims"); ok {
		clientApp.IdTokenClaims = value.(string)
	}
	if value, ok := d.GetOk("show_landingpage"); ok {
		clientApp.ShowLandingPage = value.(bool)
	}
	if value, ok := d.GetOk("binddn"); ok {
		clientApp.BindDN = value.(string)
	}
	if value, ok := d.GetOk("used_for_provisioning"); ok {
		clientApp.UsedForProvisioning = value.(bool)
	}
	if value, ok := d.GetOk("metadata_url"); ok {
		clientApp.MetadataUrl = value.(string)
	}
	if value, ok := d.GetOk("metadata"); ok {
		clientApp.Metadata = value.(string)
	}
	if value, ok := d.GetOk("subject_format"); ok {
		clientApp.SubjectFormat = value.(string)
	}
	if value, ok := d.GetOk("segments"); ok {
		clientApp.Segments = value.(string)
	}

	newApp, err := client.ClientApplications.Create(clientApp)
	if err != nil {
		tflog.Debug(ctx, err.Error(), apiErrorToLogFields(err))
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not create ClientApplication",
			Detail:   fmt.Sprintf("Error: %s", err.Error()),
		})
		return diags
	}

	if secret, err := newApp.GetSecret(); err == nil {
		err := d.Set("clientsecret", secret)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Could not set secret from ClientApplication",
				Detail:   fmt.Sprintf("Error: %s", err.Error()),
			})
			return diags
		}
	}

	d.SetId(strconv.FormatInt(newApp.Self().ID, 10))
	diags = append(diags, resourceClientApplicationRead(ctx, d, m)...)

	return diags
}

func resourceClientApplicationImportContext(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	client := m.(*keyhubclient.Client)

	clientUuid, err := uuid.Parse(d.Id())
	if err != nil {
		return nil, fmt.Errorf("`%s` is not a valid uuid", d.Id())
	}

	tmpApp, err := client.ClientApplications.GetByUUID(clientUuid)
	if err != nil {
		return nil, fmt.Errorf("Could not find clientApplication with uuid `%s`", d.Id())
	} else {
		d.SetId(strconv.FormatInt(tmpApp.Self().ID, 10))
	}

	err = d.Set("uuid", clientUuid.String())
	if err != nil {
		return nil, fmt.Errorf("coult not set uuid: %s", err.Error())
	}

	return []*schema.ResourceData{d}, nil
}

func resourceClientApplicationUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*keyhubclient.Client)
	_ = client
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Cannot update a ClientApplication",
		Detail:   "Currently Keyhub doesn't allow a client to update a ClientApplication after it's created, so any changes aren't stored",
	})

	resourceClientApplicationRead(ctx, d, m)

	return diags
}

func resourceClientApplicationDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*keyhubclient.Client)
	_ = client
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Cannot delete a ClientApplication",
		Detail:   "Currently Keyhub doesn't allow a client to delete a ClientApplication. We will only delete the ClientApplication from the terraform state",
	})

	d.SetId("")

	return diags
}
