package keyhub

import (
	"context"
	"strconv"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

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
		// "file":{
		// 	Type:      schema.TypeString,
		// 	Sensitive: true,
		// 	Computed:  true,
		// 	Required:  false,
		// },
		"comment": {
			Type:      schema.TypeString,
			Sensitive: true,
			Computed:  true,
			Required:  false,
		}}

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
			Required: true,
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

	groupUUIDString := d.Get("groupuuid").(string)
	groupUUID, err := uuid.Parse(groupUUIDString)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Field 'groupuuid' is not a valid UUID",
			Detail:   err.Error(),
		})
	}
	group, err := client.Groups.GetByUUID(groupUUID)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not GET group " + groupUUIDString + " to READ vault record(s)",
			Detail:   err.Error(),
		})
	}

	uuidString, valueExists := d.GetOk("uuid")
	if valueExists {
		UUID, err := uuid.Parse(uuidString.(string))
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Field 'uuid' is not a valid UUID",
				Detail:   err.Error(),
			})
		}

		vaultRecord, err = client.Vaults.GetByUUID(group, UUID, &keyhubmodel.VaultRecordAdditionalQueryParams{Secret: true})
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Could not GET vault record " + uuidString.(string),
				Detail:   err.Error(),
			})
		}
	}

	if d.Id() != "" {
		ID, err := strconv.ParseInt(d.Id(), 10, 64)
		if !valueExists && err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Could not parse ID " + d.Id(),
				Detail:   err.Error(),
			})
		} else {
			valueExists = true
			vaultRecord, err = client.Vaults.GetByID(group, ID, &keyhubmodel.VaultRecordAdditionalQueryParams{Secret: true})
			if err != nil {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Could not GET vault record " + d.Id(),
					Detail:   err.Error(),
				})
			}
		}
	}

	if !valueExists {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not GET vault record either by UUID or ID",
			Detail:   "",
		})
		return diags
	}

	idString := strconv.FormatInt(vaultRecord.Self().ID, 10)
	if err := d.Set("id", idString); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not set value for id",
			Detail:   err.Error(),
		})
	}
	if err := d.Set("groupuuid", group.UUID); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not set value for groupuuid",
			Detail:   err.Error(),
		})
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
	// if err := d.Set("file", vaultRecord.AdditionalObjects.Secret.File); err != nil {
	// diags = append(diags, diag.Diagnostic{
	// 	Severity: diag.Error,
	// 	Summary:  "Could not set value for file",
	// 	Detail:   err.Error(),
	// })
	// }
	if err := d.Set("comment", vaultRecord.AdditionalObjects.Secret.Comment); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not set value for comment",
			Detail:   err.Error(),
		})
	}

	d.SetId(strconv.FormatInt(vaultRecord.Self().ID, 10))

	return diags
}
