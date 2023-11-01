package keyhub

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	keyhubclient "github.com/topicuskeyhub/go-keyhub"
	keyhubmodel "github.com/topicuskeyhub/go-keyhub/model"
)

func resourceVaultRecord() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceVaultRecordCreate,
		ReadContext:   dataSourceVaultRecordRead, //we use the DataSource READ as it supports both data types
		UpdateContext: resourceVaultRecordUpdate,
		DeleteContext: resourceVaultRecordDelete,
		Schema:        VaultRecordResourceSchema(),
		Importer: &schema.ResourceImporter{
			StateContext: resourceVaultRecordImportContext,
		},
	}
}

func resourceVaultRecordImportContext(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	grpUuid, err := uuid.Parse(d.Id())
	if err != nil {
		return nil, fmt.Errorf("`%s` is not a valid uuid", d.Id())
	}

	err = d.Set("uuid", grpUuid.String())
	if err != nil {
		return nil, fmt.Errorf("coult not set uuid: %s", err.Error())
	}

	return []*schema.ResourceData{d}, nil
}

func VaultRecordResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
			Required: false,
		},
		"groupuuid": {
			Type:     schema.TypeString,
			Required: true,
		},
		"uuid": {
			Type:     schema.TypeString,
			Computed: true,
			Required: false,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"url": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"username": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"filename": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"password": {
			Type:      schema.TypeString,
			Sensitive: true,
			Optional:  true,
		},
		"totp": {
			Type:      schema.TypeString,
			Sensitive: true,
			Optional:  true,
		},
		"file": {
			Type:      schema.TypeString,
			Sensitive: true,
			Optional:  true,
		},
		"base64_encoded": {
			Type:        schema.TypeBool,
			Sensitive:   false,
			Optional:    true,
			Default:     false,
			Description: "If true, the value of `file` must be base64 encoded",
		},
		"comment": {
			Type:      schema.TypeString,
			Sensitive: true,
			Optional:  true,
		},
		"enddate": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "End date of vaultrecord, in YYYY-MM-DD format",
			ValidateDiagFunc: func(v any, p cty.Path) diag.Diagnostics {
				strvalue := v.(string)
				value, err := time.Parse("2006-01-02", strvalue)
				var diags diag.Diagnostics
				if err != nil {
					diag := diag.Diagnostic{
						Severity: diag.Error,
						Summary:  "not a valid date",
						Detail:   fmt.Sprintf("%q could not be parsed: %q", strvalue, err.Error()),
					}
					diags = append(diags, diag)
					return diags
				}
				testvalue := value.Format("2006-01-02")
				if testvalue != strvalue {
					diag := diag.Diagnostic{
						Severity: diag.Error,
						Summary:  "parse diff detected",
						Detail:   fmt.Sprintf("%q parsed wrong as %q", strvalue, testvalue),
					}
					diags = append(diags, diag)
					return diags
				}
				return diags
			},
		},
		"warningperiod": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Warning period for the end date of the vaultrecord",
			ValidateDiagFunc: validation.ToDiagFunc(
				validation.StringInSlice(
					[]string{
						string(keyhubmodel.WARNINGPERIOD_NEVER),
						string(keyhubmodel.WARNINGPERIOD_AT_EXPIRATION),
						string(keyhubmodel.WARNINGPERIOD_TWO_WEEKS),
						string(keyhubmodel.WARNINGPERIOD_ONE_MONTH),
						string(keyhubmodel.WARNINGPERIOD_TWO_MONTHS),
						string(keyhubmodel.WARNINGPERIOD_THREE_MONTHS),
						string(keyhubmodel.WARNINGPERIOD_SIX_MONTHS),
					},
					false,
				),
			),
		},
	}
}

func resourceVaultRecordCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*keyhubclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	name := d.Get("name").(string)

	groupUUIDString := d.Get("groupuuid").(string)
	groupUUID, err := uuid.Parse(groupUUIDString)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Field 'groupuuid' is not a valid UUID",
			Detail:   err.Error(),
		})
		return diags
	}
	group, err := client.Groups.GetByUUID(groupUUID)
	if err != nil {
		tflog.Debug(ctx, err.Error(), apiErrorToLogFields(err))
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not GET group " + groupUUIDString + " for new vault record " + name,
			Detail:   err.Error(),
		})
		return diags
	}

	//query vaultrecords by name to prevent duplicates
	existingVaultRecord, err := client.Vaults.List(group, &keyhubmodel.VaultRecordQueryParams{Name: name}, &keyhubmodel.VaultRecordAdditionalQueryParams{Secret: true})
	if err != nil {
		tflog.Debug(ctx, err.Error(), apiErrorToLogFields(err))
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not CREATE vaultRecord in group " + groupUUIDString,
			Detail:   err.Error(),
		})
		return diags
	}
	if len(existingVaultRecord) > 0 {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not CREATE vaultRecord",
			Detail:   "Record with same name exists in group in group " + groupUUIDString,
		})
		return diags
	}

	secrets := &keyhubmodel.VaultRecordSecretAdditionalObject{}
	vaultRecord := keyhubmodel.NewVaultRecord(name, secrets)

	//copy schema data to model
	//use generic copy method. also used in UPDATE.
	diags = append(diags, vaultRecordSchemaToModel(d, vaultRecord)...)
	if diags.HasError() {
		return diags
	}

	newVaultRecord, err := client.Vaults.Create(group, vaultRecord)
	if err != nil {
		tflog.Debug(ctx, err.Error(), apiErrorToLogFields(err))
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not CREATE vaultRecord in group " + groupUUIDString,
			Detail:   err.Error(),
		})
		return diags
	}

	// d.SetId(group.UUID + "/" + newVaultRecord.UUID)
	d.SetId(strconv.FormatInt(newVaultRecord.Self().ID, 10))

	dataSourceVaultRecordRead(ctx, d, m)

	return diags
}

func resourceVaultRecordUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*keyhubclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	groupUUIDString := d.Get("groupuuid").(string)
	groupUUID, err := uuid.Parse(groupUUIDString)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Field 'groupuuid' is not a valid UUID",
			Detail:   err.Error(),
		})
		return diags
	}

	ID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not parse ID " + d.Id(),
			Detail:   err.Error(),
		})
		return diags
	}

	group, err := client.Groups.GetByUUID(groupUUID)
	if err != nil {
		tflog.Debug(ctx, err.Error(), apiErrorToLogFields(err))
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not GET group " + groupUUIDString + " for vault record " + d.Id(),
			Detail:   err.Error(),
		})
		return diags
	}

	vaultRecord, err := client.Vaults.GetByID(group, ID, &keyhubmodel.VaultRecordAdditionalQueryParams{Secret: true})
	if err != nil {
		tflog.Debug(ctx, err.Error(), apiErrorToLogFields(err))
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not GET vault record " + d.Id(),
			Detail:   err.Error(),
		})
		return diags
	}

	//copy schema data to model
	//use generic copy method. also used in CREATE.
	diags = append(diags, vaultRecordSchemaToModel(d, vaultRecord)...)
	if diags.HasError() {
		return diags
	}

	_, err = client.Vaults.Update(group, vaultRecord)
	if err != nil {
		tflog.Debug(ctx, err.Error(), apiErrorToLogFields(err))
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not UPDATE vault record " + vaultRecord.UUID,
			Detail:   err.Error(),
		})
		return diags
	}

	dataSourceVaultRecordRead(ctx, d, m)

	return diags
}

func resourceVaultRecordDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*keyhubclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	groupUUIDString := d.Get("groupuuid").(string)
	groupUUID, err := uuid.Parse(groupUUIDString)
	if err != nil {
		tflog.Debug(ctx, err.Error(), apiErrorToLogFields(err))
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Field 'groupuuid' is not a valid UUID",
			Detail:   err.Error(),
		})
	}

	ID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not parse ID " + d.Id(),
			Detail:   err.Error(),
		})
		return diags
	}

	group, err := client.Groups.GetByUUID(groupUUID)
	if err != nil {
		tflog.Debug(ctx, err.Error(), apiErrorToLogFields(err))
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not GET group " + groupUUIDString + " for vault record " + d.Id(),
			Detail:   err.Error(),
		})
		return diags
	}

	err = client.Vaults.DeleteByID(group, ID)
	if err != nil {
		tflog.Debug(ctx, err.Error(), apiErrorToLogFields(err))
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not DELETE vaultrecord " + d.Id(),
			Detail:   err.Error(),
		})
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}

func vaultRecordSchemaToModel(d *schema.ResourceData, vaultRecord *keyhubmodel.VaultRecord) diag.Diagnostics {

	diags := diag.Diagnostics{}

	if d.HasChange("name") {
		value := d.Get("name")
		vaultRecord.Name = value.(string)
	}
	if d.HasChange("url") {
		value := d.Get("url")
		vaultRecord.URL = value.(string)
	}
	if d.HasChange("username") {
		value := d.Get("username")
		vaultRecord.Username = value.(string)
	}
	if d.HasChange("color") {
		value := d.Get("color")
		vaultRecord.Color = value.(string)
	}
	if d.HasChange("filename") {
		value := d.Get("filename")
		vaultRecord.Filename = value.(string)
	}

	if d.HasChange("password") {
		value := d.Get("password")
		val := value.(string)
		vaultRecord.AdditionalObjects.Secret.Password = &val
	}
	if d.HasChange("totp") {
		value := d.Get("totp")
		val := value.(string)
		vaultRecord.AdditionalObjects.Secret.Totp = &val
	}

	if d.HasChanges("file", "base64_encoded") {
		value := d.Get("file").(string)
		isBase64 := d.Get("base64_encoded").(bool)

		var rawValue []byte
		var err error

		if isBase64 {
			rawValue, err = base64.StdEncoding.DecodeString(value)
			if err != nil {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "base64 decoding error for file",
					Detail:   err.Error(),
				})
			}
		} else {
			rawValue = []byte(value)
		}

		// Check if file isn't larger than accepted by Keyhub before sending.
		if len(rawValue) > 2*1024*1024 {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "File exceeds limit of 2048 KiB",
				Detail:   fmt.Sprintf("Size of file is %d KiB", len(rawValue)/1024),
			})
		}

		vaultRecord.AdditionalObjects.Secret.File = &rawValue
	}
	if d.HasChange("comment") {
		value := d.Get("comment")
		val := value.(string)
		vaultRecord.AdditionalObjects.Secret.Comment = &val
	}

	if d.HasChange("enddate") {
		value := d.Get("enddate")
		if value != "" {
			val, err := time.Parse("2006-01-02", value.(string))
			if err != nil {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Could not parse enddate",
					Detail:   fmt.Sprintf("Could not parse enddate %q: %s", value.(string), err.Error()),
				})
			} else {
				vaultRecord.EndDate = val
			}
		} else {
			vaultRecord.EndDate = time.Time{}
		}
	}

	if d.HasChange("warningperiod") {
		value := d.Get("warningperiod")
		vaultRecord.WarningPeriod = keyhubmodel.RecordWarningPeriod(value.(string))
	}

	return diags
}
