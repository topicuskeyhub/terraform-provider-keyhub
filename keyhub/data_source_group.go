package keyhub

import (
	"context"
	"strconv"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	keyhubclient "github.com/topicuskeyhub/go-keyhub"
)

func dataSourceGroup() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGroupRead,
		Schema:      GroupSchema(),
	}
}

func GroupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"uuid": {
			Type:     schema.TypeString,
			Required: true,
		},
		"name": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func dataSourceGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*keyhubclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	uuidString := d.Get("uuid").(string)
	UUID, err := uuid.Parse(uuidString)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Field 'uuid' is not a valid UUID",
			Detail:   err.Error(),
		})
		return diags
	}

	group, err := client.Groups.GetByUUID(UUID)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not GET group " + uuidString,
			Detail:   err.Error(),
		})
		return diags
	}

	idString := strconv.FormatInt(group.Self().ID, 10)
	if err := d.Set("id", idString); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not set value for id",
			Detail:   err.Error(),
		})
	}
	if err := d.Set("uuid", group.UUID); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not set value for uuid",
			Detail:   err.Error(),
		})
	}
	if err := d.Set("name", group.Name); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not set value for name",
			Detail:   err.Error(),
		})
	}

	d.SetId(strconv.FormatInt(group.Self().ID, 10))

	return diags
}
