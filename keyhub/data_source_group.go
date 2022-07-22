package keyhub

import (
	"context"
	"fmt"
	keyhubmodel "github.com/topicuskeyhub/go-keyhub/model"
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
		"description": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"extended_access": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"vault_recovery": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"audit_months": {
			Type: schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Computed: true,
		},

		"member": {
			Type:     schema.TypeSet,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"uuid": {
						Type:     schema.TypeString,
						Computed: true,
					},
					"rights": {
						Type:     schema.TypeString,
						Computed: true,
					},
					"name": {
						Type:     schema.TypeString,
						Computed: true,
					},
				},
			},
		},

		"rotating_password_required": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"record_trail": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"private_group": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"hide_audit_trail": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"application_administration": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"auditor": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"single_managed": {
			Type:     schema.TypeBool,
			Computed: true,
		},

		"provisioning_auth_groupuuid": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"membership_auth_groupuuid": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"auditing_auth_groupuuid": {
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
		diags = append(diags, NewDiagnosticSetError("id", err))
	}
	if err := d.Set("uuid", group.UUID); err != nil {
		diags = append(diags, NewDiagnosticSetError("uuid", err))
	}
	if err := d.Set("name", group.Name); err != nil {
		diags = append(diags, NewDiagnosticSetError("name", err))
	}
	if err := d.Set("description", group.Description); err != nil {
		diags = append(diags, NewDiagnosticSetError("description", err))
	}
	if err := d.Set("extended_access", group.ExtendedAccess); err != nil {
		diags = append(diags, NewDiagnosticSetError("extended_access", err))
	}
	if err := d.Set("vault_recovery", group.VaultRecovery); err != nil {
		diags = append(diags, NewDiagnosticSetError("vault_recovery", err))
	}
	if err := d.Set("audit_months", group.AuditConfig.Months.ToList()); err != nil {
		diags = append(diags, NewDiagnosticSetError("audit_months", err))
	}

	// Bools
	if err := d.Set("rotating_password_required", group.RotatingPasswordRequired); err != nil {
		diags = append(diags, NewDiagnosticSetError("rotating_password_required", err))
	}
	if err := d.Set("record_trail", group.HideAuditTrail); err != nil {
		diags = append(diags, NewDiagnosticSetError("record_trail", err))
	}
	if err := d.Set("private_group", group.PrivateGroup); err != nil {
		diags = append(diags, NewDiagnosticSetError("private_group", err))
	}
	if err := d.Set("hide_audit_trail", group.HideAuditTrail); err != nil {
		diags = append(diags, NewDiagnosticSetError("hide_audit_trail", err))
	}
	if err := d.Set("application_administration", group.ApplicationAdministration); err != nil {
		diags = append(diags, NewDiagnosticSetError("application_administration", err))
	}
	if err := d.Set("auditor", group.Auditor); err != nil {
		diags = append(diags, NewDiagnosticSetError("auditor", err))
	}
	if err := d.Set("single_managed", group.SingleManaged); err != nil {
		diags = append(diags, NewDiagnosticSetError("single_managed", err))
	}
	if group.AuthorizingGroupProvisioning != nil {
		if err := d.Set("provisioning_auth_groupuuid", group.AuthorizingGroupProvisioning.UUID); err != nil {
			diags = append(diags, NewDiagnosticSetError("provisioning_auth_groupuuid", err))
		}
	}
	if group.AuthorizingGroupMembership != nil {
		if err := d.Set("membership_auth_groupuuid", group.AuthorizingGroupMembership.UUID); err != nil {
			diags = append(diags, NewDiagnosticSetError("membership_auth_groupuuid", err))
		}
	}
	if group.AuthorizingGroupAuditing != nil {
		if err := d.Set("auditing_auth_groupuuid", group.AuthorizingGroupAuditing.UUID); err != nil {
			diags = append(diags, NewDiagnosticSetError("auditing_auth_groupuuid", err))
		}
	}

	if group.AdditionalObjects != nil {

		if group.AdditionalObjects.Admins != nil {
			if err := d.Set("member", flattenMembers(group.AdditionalObjects.Admins)); err != nil {
				diags = append(diags, NewDiagnosticSetError("member", err))
			}
		}

		if group.AdditionalObjects.ClientPermissions != nil {

			if err := d.Set("client", flattenClientPermissions(group.AdditionalObjects.ClientPermissions)); err != nil {
				diags = append(diags, NewDiagnosticSetError("client", err))
			}
		}

	}

	d.SetId(strconv.FormatInt(group.Self().ID, 10))

	return diags
}

func flattenMembers(members *keyhubmodel.GroupAccountList) []interface{} {

	if members == nil {
		return make([]interface{}, 0)
	}

	list := make([]interface{}, len(members.Items))

	for i, member := range members.Items {

		m := make(map[string]interface{})
		m["uuid"] = member.UUID
		m["rights"] = member.Rights
		m["name"] = member.DisplayName

		list[i] = m

	}

	return list

}

func flattenClientPermissions(clientPermissions *keyhubmodel.ClientPermissionsWithClient) []interface{} {

	if clientPermissions == nil {
		return make([]interface{}, 0)
	}

	groupMap := map[string][]string{}

	for _, perm := range clientPermissions.Items {
		groupMap[perm.Client.UUID] = append(groupMap[perm.Client.UUID], string(perm.Value))
	}

	list := make([]interface{}, len(groupMap))

	for uuid, permissions := range groupMap {

		p := make(map[string]interface{})
		p["uuid"] = uuid
		p["permissions"] = permissions

		list = append(list, p)

	}

	return list

}

func NewDiagnosticSetError(key string, err error) diag.Diagnostic {
	return diag.Diagnostic{
		Severity: diag.Error,
		Summary:  fmt.Sprintf("Could not set value for %s", key),
		Detail:   err.Error(),
	}
}
