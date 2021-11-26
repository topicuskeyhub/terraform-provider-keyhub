package keyhub

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/topicuskeyhub/go-keyhub"
	keyhubmodel "github.com/topicuskeyhub/go-keyhub/model"
)

func dataSourceVaultRecord() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVaultRecordRead,
		Schema:      VaultRecordSchema(),
	}
}

func VaultRecordSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeInt,
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
		"secrets": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
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
		// "file": {
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
		},
	}
}

func dataSourceVaultRecordRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*keyhub.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	groupUUID := d.Get("groupuuid").(string)
	UUID := d.Get("uuid").(string)
	secrets := d.Get("secrets").(bool)

	group, err := client.Groups.Get(groupUUID)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not GET group " + groupUUID + " for vault record " + UUID,
			Detail:   err.Error(),
		})
	}

	vaultRecord, err := client.Vaults.GetRecord(group, UUID, keyhubmodel.RecordOptions{Secret: secrets})
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not GET secrets of vault record " + UUID,
			Detail:   err.Error(),
		})
	}

	if err := d.Set("id", vaultRecord.Self().ID); err != nil {
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
			Summary:  "Could not set secrets",
			Detail:   err.Error(),
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

	d.SetId(vaultRecord.UUID)

	return diags
}
