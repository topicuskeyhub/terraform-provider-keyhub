package keyhub

import (
	"context"
	"encoding/base64"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"regexp"
	"strconv"

	keyhubclient "github.com/topicuskeyhub/go-keyhub"
	keyhubmodel "github.com/topicuskeyhub/go-keyhub/model"
)

func dataSourceVaultRecord() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVaultRecordRead,
		Schema:      VaultRecordSchema(),
	}
}

func VaultRecordSchema() map[string]*schema.Schema {
	baseSchema := VaultRecordBaseSchema()
	secretsSchema := map[string]*schema.Schema{
		"password": {
			Type:      schema.TypeString,
			Sensitive: true,
			Computed:  true,
			Required:  false,
		},
		"totp": {
			Type:      schema.TypeString,
			Sensitive: true,
			Computed:  true,
			Required:  false,
		},
		"file": {
			Type:      schema.TypeString,
			Sensitive: true,
			Computed:  true,
			Required:  false,
		},
		"base64_encoded": {
			Type:        schema.TypeBool,
			Sensitive:   false,
			Optional:    true,
			Default:     false,
			Description: "If true, the content of `file` will be base64 encoded",
		},
		"comment": {
			Type:      schema.TypeString,
			Sensitive: true,
			Computed:  true,
			Required:  false,
		},
		"enddate": {
			Type:      schema.TypeString,
			Sensitive: false,
			Computed:  true,
			Required:  false,
		},
		"warningperiod": {
			Type:      schema.TypeString,
			Sensitive: false,
			Computed:  true,
			Required:  false,
		},
	}

	schema := map[string]*schema.Schema{}
	for k, v := range baseSchema {
		schema[k] = v
	}
	for k, v := range secretsSchema {
		schema[k] = v
	}

	return schema
}

func VaultRecordBaseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"groupuuid": {
			Type:     schema.TypeString,
			Computed: true,
			Optional: true,
		},
		"uuid": {
			Type:     schema.TypeString,
			Required: true,
		},
		"name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"url": {
			Type:     schema.TypeString,
			Computed: true,
			Required: false,
		},
		"username": {
			Type:     schema.TypeString,
			Computed: true,
			Required: false,
		},
		"filename": {
			Type:     schema.TypeString,
			Computed: true,
			Required: false,
		},
	}
}

//
//This function will support both DataSource READ and Resource READ by checking for UUID and ID.
func dataSourceVaultRecordRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*keyhubclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	var vaultRecord *keyhubmodel.VaultRecord
	var group *keyhubmodel.Group
	var groupUUID uuid.UUID
	var err error

	groupUUIDString := d.Get("groupuuid").(string)
	if groupUUIDString != "" {
		groupUUID, err = uuid.Parse(groupUUIDString)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Field 'groupuuid' is not a valid UUID",
				Detail:   err.Error(),
			})
			return diags
		}
		group, err = client.Groups.GetByUUID(groupUUID)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Could not GET group " + groupUUIDString + " to READ vault record(s)",
				Detail:   err.Error(),
			})
			return diags
		}
	}

	uuidString, valueExists := d.GetOk("uuid")
	if valueExists {
		var UUID uuid.UUID
		UUID, err = uuid.Parse(uuidString.(string))
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Field 'uuid' is not a valid UUID",
				Detail:   err.Error(),
			})
			return diags
		}

		if group != nil {
			vaultRecord, err = client.Vaults.GetByUUID(group, UUID, &keyhubmodel.VaultRecordAdditionalQueryParams{Secret: true})
		} else {
			vaultRecord, err = client.Vaults.FindByUUIDForClient(UUID, &keyhubmodel.VaultRecordAdditionalQueryParams{Secret: true})
		}
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Could not GET vault record " + uuidString.(string),
				Detail:   err.Error(),
			})
			return diags
		}
	} else {

		if d.Id() != "" {
			ID, err := strconv.ParseInt(d.Id(), 10, 64)
			if err != nil {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Could not parse ID " + d.Id(),
					Detail:   err.Error(),
				})
				return diags
			}

			if group != nil {
				vaultRecord, err = client.Vaults.GetByID(group, ID, &keyhubmodel.VaultRecordAdditionalQueryParams{Secret: true})
			} else {
				vaultRecord, err = client.Vaults.FindByIDForClient(ID, &keyhubmodel.VaultRecordAdditionalQueryParams{Secret: true})
			}
			if err != nil {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Could not GET vault record " + d.Id(),
					Detail:   err.Error(),
				})
				return diags
			}

		}
	}

	idString := strconv.FormatInt(vaultRecord.Self().ID, 10)
	if err := d.Set("id", idString); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not set value for id",
			Detail:   err.Error(),
		})
	}
	if group == nil {
		// If group is nil, retrieve group from vaultrecord url, so we can set the groupuuid parameter
		var groupId int64
		r, _ := regexp.Compile("^((.+)/group/([0-9]+))/vault/record/([0-9]+)")
		matches := r.FindStringSubmatch(vaultRecord.Self().Href)
		// 0 = full url, 1 = group url, 2 = base url, 3 = group id, 4 = record id
		groupId, err = strconv.ParseInt(matches[3], 10, 64)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Could not parse groupId from record url",
				Detail:   err.Error(),
			})
		} else {
			group, err = client.Groups.GetById(groupId)
			if err != nil {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Could not GET group " + matches[3] + " for found vault record(s)",
					Detail:   err.Error(),
				})
			}
		}
	}
	if group != nil {
		if err := d.Set("groupuuid", group.UUID); err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Could not set value for groupuuid",
				Detail:   err.Error(),
			})
		}
	}

	if err := d.Set("uuid", vaultRecord.UUID); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not set value for uuid",
			Detail:   err.Error(),
		})
	}
	if err := d.Set("name", vaultRecord.Name); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not set value for name",
			Detail:   err.Error(),
		})
	}
	if err := d.Set("url", vaultRecord.URL); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not set value for url",
			Detail:   err.Error(),
		})
	}
	if err := d.Set("username", vaultRecord.Username); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not set value for username",
			Detail:   err.Error(),
		})
	}
	if err := d.Set("filename", vaultRecord.Filename); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not set value for filename",
			Detail:   err.Error(),
		})
	}

	if vaultRecord.AdditionalObjects == nil ||
		vaultRecord.AdditionalObjects.Secret == nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not set secrets, AdditionalObject was not initialized",
			Detail:   "",
		})

		return diags
	}

	if err := d.Set("password", vaultRecord.AdditionalObjects.Secret.Password); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not set value for password",
			Detail:   err.Error(),
		})
	}
	if err := d.Set("totp", vaultRecord.AdditionalObjects.Secret.Totp); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not set value for totp",
			Detail:   err.Error(),
		})
	}

	if vaultRecord.AdditionalObjects.Secret.File != nil {
		var value string
		value = string(*vaultRecord.AdditionalObjects.Secret.File)

		if encoded, ok := d.GetOk("base64_encoded"); ok {
			if encoded.(bool) {
				value = base64.StdEncoding.EncodeToString(*vaultRecord.AdditionalObjects.Secret.File)
			}
		}

		if err := d.Set("file", value); err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Could not set value for file",
				Detail:   err.Error(),
			})
		}
	}

	if err := d.Set("comment", vaultRecord.AdditionalObjects.Secret.Comment); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not set value for comment",
			Detail:   err.Error(),
		})
	}

	if err := d.Set("enddate", vaultRecord.EndDate.Format("2006-01-02")); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not set value for enddate",
			Detail:   err.Error(),
		})
	}

	if err := d.Set("warningperiod", vaultRecord.WarningPeriod); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not set value for warning period",
			Detail:   err.Error(),
		})
	}

	d.SetId(strconv.FormatInt(vaultRecord.Self().ID, 10))

	return diags
}
