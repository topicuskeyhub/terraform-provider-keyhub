package keyhub

import (
	"context"
	"strconv"

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
			Computed: true,
		},
		"uuid": {
			Type:     schema.TypeString,
			Computed: true,
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

func dataSourceVaultRecordRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*keyhub.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	groupUUID := d.Get("groupuuid").(string)
	UUID := d.Get("uuid").(string)

	group, err := client.Groups.Get(groupUUID)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not GET group " + groupUUID + " for vault record " + UUID,
			Detail:   err.Error(),
		})
	}

	vaultrecord, err := client.Vaults.GetRecord(group, UUID, keyhubmodel.RecordOptions{Secret: true})
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not GET vault record " + UUID,
			Detail:   err.Error(),
		})
	}

	if err := d.Set("id", vaultrecord.Self().ID); err != nil {
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
	if err := d.Set("uuid", vaultrecord.UUID); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not set value for uuid",
			Detail:   err.Error(),
		})
	}
	if err := d.Set("name", vaultrecord.Name); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not set value for name",
			Detail:   err.Error(),
		})
	}
	if err := d.Set("url", vaultrecord.URL); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not set value for url",
			Detail:   err.Error(),
		})
	}
	if err := d.Set("username", vaultrecord.Username); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not set value for username",
			Detail:   err.Error(),
		})
	}
	if err := d.Set("filename", vaultrecord.Filename); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not set value for filename",
			Detail:   err.Error(),
		})
	}

	if vaultrecord.AdditionalObjects != nil &&
		vaultrecord.AdditionalObjects.Secret != nil {

		if err := d.Set("password", vaultrecord.AdditionalObjects.Secret.Password); err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Could not set value for password",
				Detail:   err.Error(),
			})
		}
		if err := d.Set("totp", vaultrecord.AdditionalObjects.Secret.Totp); err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Could not set value for totp",
				Detail:   err.Error(),
			})
		}
		// if err := d.Set("file", vaultrecord.AdditionalObjects.Secret.File); err != nil {
		// diags = append(diags, diag.Diagnostic{
		// 	Severity: diag.Error,
		// 	Summary:  "Could not set value for file",
		// 	Detail:   err.Error(),
		// })
		// }
		if err := d.Set("comment", vaultrecord.AdditionalObjects.Secret.Comment); err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Could not set value for comment",
				Detail:   err.Error(),
			})
		}
	}

	d.SetId(strconv.FormatInt(int64(vaultrecord.Self().ID), 10))

	return diags
}
