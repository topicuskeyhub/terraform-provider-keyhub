package keyhub

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/topicuskeyhub/go-keyhub"
	keyhubmodel "github.com/topicuskeyhub/go-keyhub/model"
)

func resourceVaultRecord() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceVaultRecordCreate,
		ReadContext:   resourceVaultRecordRead,
		UpdateContext: resourceVaultRecordUpdate,
		DeleteContext: resourceVaultRecordDelete,
		Schema:        VaultRecordResourceSchema(),
	}
}

func VaultRecordResourceSchema() map[string]*schema.Schema {
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
			Required: false,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"url": {
			Type:     schema.TypeString,
			Required: false,
		},
		"username": {
			Type:     schema.TypeString,
			Required: false,
		},
		"filename": {
			Type:     schema.TypeString,
			Required: false,
		},
		"password": {
			Type:      schema.TypeString,
			Sensitive: true,
			Required:  false,
		},
		"totp": {
			Type:      schema.TypeString,
			Sensitive: true,
			Required:  false,
		},
		// "file": {
		// 	Type:      schema.TypeString,
		// 	Sensitive: true,
		// 	Required:  false,
		// },
		"comment": {
			Type:      schema.TypeString,
			Sensitive: true,
			Required:  false,
		},
	}
}

func resourceVaultRecordCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*keyhub.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	groupUUID := d.Get("groupuuid").(string)
	name := d.Get("name").(string)

	group, err := client.Groups.Get(groupUUID)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not GET group " + groupUUID + " for new vault record " + name,
			Detail:   err.Error(),
		})
	}

	vaultRecord := new(keyhubmodel.VaultRecord)
	vaultRecord.Name = d.Get("name").(string)
	vaultRecord.URL = d.Get("url").(string)
	vaultRecord.Username = d.Get("username").(string)
	vaultRecord.Color = d.Get("color").(string)
	vaultRecord.Filename = d.Get("filename").(string)

	_, existsPassword := d.GetOk("password")
	_, existsTotp := d.GetOk("totp")
	// _, existsFile := d.GetOk("file")
	_, existsComment := d.GetOk("comment")
	if existsPassword || existsTotp || existsComment {
		vaultRecord.AdditionalObjects = new(keyhubmodel.VaultRecordAdditionalObjects)
		vaultRecord.AdditionalObjects.Secret = new(keyhubmodel.VaultRecordSecretAdditionalObject)
	}

	if existsPassword {
		vaultRecord.AdditionalObjects.Secret.Password = d.Get("password").(string)
	}
	if existsTotp {
		vaultRecord.AdditionalObjects.Secret.Totp = d.Get("totp").(string)
	}
	// if existsFile {
	// vaultRecord.AdditionalObjects.Secret.File = d.Get("file").(string)
	// }
	if existsComment {
		vaultRecord.AdditionalObjects.Secret.Comment = d.Get("comment").(string)
	}

	client.Vaults.

	return diags
}

func resourceVaultRecordRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*keyhub.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}

func resourceVaultRecordUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*keyhub.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return resourceVaultRecordRead(ctx, d, m)
}

func resourceVaultRecordDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*keyhub.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}
