package keyhub

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	keyhubclient "github.com/topicuskeyhub/go-keyhub"
)

func dataSourceAccount() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAccountRead,
		Schema:      AccountSchema(),
	}
}

func AccountSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"uuid": {
			Type:     schema.TypeString,
			Required: true,
		},
		"username": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func dataSourceAccountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*keyhubclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	UUID := d.Get("uuid").(string)
	account, err := client.Accounts.GetByUUID(UUID)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not GET account " + UUID,
			Detail:   err.Error(),
		})
	}

	idString := strconv.FormatInt(account.Self().ID, 10)
	if err := d.Set("id", idString); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not set value for id",
			Detail:   err.Error(),
		})
	}
	if err := d.Set("uuid", account.UUID); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not set value for uuid",
			Detail:   err.Error(),
		})
	}
	if err := d.Set("username", account.Username); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not set value for username",
			Detail:   err.Error(),
		})
	}
	if err := d.Set("name", account.DisplayName); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not set value for name",
			Detail:   err.Error(),
		})
	}

	d.SetId(idString)

	return diags
}
