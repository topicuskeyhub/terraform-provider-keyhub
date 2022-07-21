package keyhub

import (
	"context"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	keyhubclient "github.com/topicuskeyhub/go-keyhub"
	"strconv"
)

func dataSourceProvisionedSystem() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceProvisionedSystemRead,
		Schema:      ProvisionedSystemSchema(),
	}
}

func ProvisionedSystemSchema() map[string]*schema.Schema {
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
		"type": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"accountcount": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"usernameprefix": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"technicaladministrator": {
			Type:     schema.TypeMap,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"externaluuid": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func dataSourceProvisionedSystemRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

	system, err := client.Systems.GetByUUID(UUID)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not GET system " + uuidString,
			Detail:   err.Error(),
		})
		return diags
	}

	idString := strconv.FormatInt(system.Self().ID, 10)
	if err := d.Set("id", idString); err != nil {
		diags = append(diags, NewDiagnosticSetError("id", err))
	}
	if err := d.Set("uuid", system.UUID); err != nil {
		diags = append(diags, NewDiagnosticSetError("uuid", err))
	}
	if err := d.Set("name", system.Name); err != nil {
		diags = append(diags, NewDiagnosticSetError("name", err))
	}
	if err := d.Set("type", system.Type); err != nil {
		diags = append(diags, NewDiagnosticSetError("type", err))
	}

	if err := d.Set("accountcount", system.AccountCount); err != nil {
		diags = append(diags, NewDiagnosticSetError("accountcount", err))
	}
	if err := d.Set("usernameprefix", system.UsernamePrefix); err != nil {
		diags = append(diags, NewDiagnosticSetError("usernameprefix", err))
	}
	if system.TechnicalAdministrator != nil {
		if err := d.Set(
			"technicaladministrator",
			map[string]string{
				"uuid": system.TechnicalAdministrator.UUID,
				"name": system.TechnicalAdministrator.Name,
			},
		); err != nil {
			diags = append(diags, NewDiagnosticSetError("technicaladministrator", err))
		}
	}
	if err := d.Set("externaluuid", system.ExternalUUID); err != nil {
		diags = append(diags, NewDiagnosticSetError("externaluuid", err))
	}

	d.SetId(strconv.FormatInt(system.Self().ID, 10))

	return diags
}
