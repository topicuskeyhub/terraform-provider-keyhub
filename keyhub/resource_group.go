package keyhub

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	keyhubclient "github.com/topicuskeyhub/go-keyhub"
	keyhubmodel "github.com/topicuskeyhub/go-keyhub/model"
	"strconv"
)

func resourceGroup() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGroupCreate,
		ReadContext:   dataSourceGroupRead,
		UpdateContext: resourceGroupUpdate,
		DeleteContext: resourceGroupDelete,
		Schema:        GroupResourceSchema(),
		Importer: &schema.ResourceImporter{
			StateContext: resourceGroupImportContext,
		},
	}
}

func resourceGroupImportContext(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	grpUuid, err := uuid.Parse(d.Id())
	if err != nil {
		return nil, fmt.Errorf("`%s` is not a valid uuid", d.Id())
	}

	err = d.Set("uuid", grpUuid.String())
	if err != nil {
		return nil, fmt.Errorf("could not set `%s` as uuid", grpUuid.String())
	}
	return []*schema.ResourceData{d}, nil
}

func GroupResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"uuid": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"member": {
			Type:     schema.TypeSet,
			Required: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"uuid": {
						Type:             schema.TypeString,
						Required:         true,
						ValidateDiagFunc: validation.ToDiagFunc(validation.IsUUID),
					},
					"rights": {
						Type:             schema.TypeString,
						Optional:         true,
						Computed:         true,
						ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{keyhubmodel.GROUP_RIGHT_MANAGER, keyhubmodel.GROUP_RIGHT_MEMBER}, false)),
					},
					"name": {
						Type:     schema.TypeString,
						Computed: true,
					},
				},
			},
		},

		"client": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"uuid": {
						Type:             schema.TypeString,
						Required:         true,
						ValidateDiagFunc: validation.ToDiagFunc(validation.IsUUID),
					},
					"permissions": {
						Type:     schema.TypeList,
						Required: true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
							ValidateDiagFunc: validation.ToDiagFunc(
								validation.StringInSlice(
									[]string{
										string(keyhubmodel.CLIENT_PERM_ACCOUNTS_QUERY),
										string(keyhubmodel.CLIENT_PERM_ACCOUNTS_REMOVE),
										string(keyhubmodel.CLIENT_PERM_GROUPONSYSTEM_CREATE),
										string(keyhubmodel.CLIENT_PERM_GROUPS_CREATE),
										string(keyhubmodel.CLIENT_PERM_GROUPS_VAULT_ACCESS_AFTER_CREATE),
										string(keyhubmodel.CLIENT_PERM_GROUPS_GRANT_PERMISSIONS_AFTER_CREATE),
										string(keyhubmodel.CLIENT_PERM_GROUPS_QUERY),
										string(keyhubmodel.CLIENT_PERM_GROUP_FULL_VAULT_ACCESS),
										string(keyhubmodel.CLIENT_PERM_GROUP_READ_CONTENTS),
										string(keyhubmodel.CLIENT_PERM_GROUP_SET_AUTHORIZATION),
										string(keyhubmodel.CLIENT_PERM_CLIENTS_CREATE),
										string(keyhubmodel.CLIENT_PERM_CLIENTS_QUERY),
									},
									false,
								),
							),
						},
					},
				},
			},
		},

		"extended_access": {
			Type:     schema.TypeString,
			Optional: true,
			ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice(
				[]string{
					keyhubmodel.GROUP_EXT_ACCESS_NOT,
					keyhubmodel.GROUP_EXT_ACCESS_1W,
					keyhubmodel.GROUP_EXT_ACCESS_2W,
				},
				false,
			)),
			Default: keyhubmodel.GROUP_EXT_ACCESS_NOT,
		},
		"vault_recovery": {
			Type:     schema.TypeString,
			Optional: true,
			ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice(
				[]string{
					keyhubmodel.VAULT_RECOVERY_FULL,
					keyhubmodel.VAULT_RECOVERY_KEY_ONLY,
					keyhubmodel.VAULT_RECOVERY_NONE,
				},
				false,
			)),
			Default: keyhubmodel.VAULT_RECOVERY_FULL,
		},
		"audit_months": {
			Type: schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeString,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice(
					[]string{
						keyhubmodel.MONTH_NAME_JAN,
						keyhubmodel.MONTH_NAME_FEB,
						keyhubmodel.MONTH_NAME_MAR,
						keyhubmodel.MONTH_NAME_APR,
						keyhubmodel.MONTH_NAME_MAY,
						keyhubmodel.MONTH_NAME_JUN,
						keyhubmodel.MONTH_NAME_JUL,
						keyhubmodel.MONTH_NAME_AUG,
						keyhubmodel.MONTH_NAME_SEP,
						keyhubmodel.MONTH_NAME_OCT,
						keyhubmodel.MONTH_NAME_NOV,
						keyhubmodel.MONTH_NAME_DEC,
					},
					false,
				)),
			},
			Optional: true,
		},

		"rotating_password_required": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"record_trail": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"private_group": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"hide_audit_trail": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"application_administration": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"auditor": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"single_managed": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"provisioning_auth_groupuuid": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"membership_auth_groupuuid": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"auditing_auth_groupuuid": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

func resourceGroupCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*keyhubclient.Client)
	_ = client
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	name := d.Get("name").(string)

	newGroup := keyhubmodel.NewEmptyGroup(name)

	if d.HasChange("description") {
		newGroup.Description = d.Get("description").(string)
	}

	/* Simple bools **/
	if d.HasChange("rotating_password_required") {
		newGroup.RotatingPasswordRequired = d.Get("rotating_password_required").(bool)
	}
	if d.HasChange("record_trail") {
		newGroup.RecordTrail = d.Get("record_trail").(bool)
	}
	if d.HasChange("private_group") {
		newGroup.PrivateGroup = d.Get("private_group").(bool)
	}
	if d.HasChange("hide_audit_trail") {
		newGroup.HideAuditTrail = d.Get("hide_audit_trail").(bool)
	}
	if d.HasChange("application_administration") {
		newGroup.ApplicationAdministration = d.Get("application_administration").(bool)
	}
	if d.HasChange("auditor") {
		newGroup.Auditor = d.Get("auditor").(bool)
	}

	/** String values validated by validation **/
	if d.HasChange("vault_recovery") {
		newGroup.VaultRecovery = d.Get("vault_recovery").(string)
	}
	if d.HasChange("extended_access") {
		newGroup.ExtendedAccess = d.Get("extended_access").(string)
	}
	if d.HasChange("audit_months") {
		months := d.Get("audit_months").([]interface{})
		newGroup.AuditConfig = keyhubmodel.NewGroupAuditConfig()
		for _, month := range months {
			newGroup.AuditConfig.Months.Enable(month.(string))
		}
	}

	/** Convert member blocks to GroupAccounts **/
	if d.HasChange("member") {
		members := d.Get("member").(*schema.Set)
		for _, memiface := range members.List() {
			member := memiface.(map[string]interface{})
			kh_member, err := client.Accounts.GetByUUID(uuid.MustParse(member["uuid"].(string)))
			if err != nil {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Member does not exist",
					Detail:   fmt.Sprintf("Could not find an account for member with uuid: %s", member["uuid"]),
				})
			} else {
				switch member["rights"].(string) {
				case keyhubmodel.GROUP_RIGHT_MEMBER:
					newGroup.AddMember(kh_member)
				default:
					newGroup.AddManager(kh_member)
				}
			}
		}
	}

	if d.HasChange("client") {
		clients := d.Get("client").(*schema.Set)
		for _, clientIface := range clients.List() {
			permclient := clientIface.(map[string]interface{})

			clientUuid := uuid.MustParse(permclient["uuid"].(string))

			xClient, err := client.ClientApplications.GetByUUID(clientUuid)
			if err != nil {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Could not get Client for setting permissions",
					Detail:   fmt.Sprintf("Client uuid %s resulted in: %s", clientUuid.String(), err.Error()),
				})
			}

			newPermissions := []keyhubmodel.Oauth2ClientPermissionValue{}
			for _, permIface := range permclient["permissions"].([]interface{}) {
				newPermissions = append(newPermissions, keyhubmodel.Oauth2ClientPermissionValue(permIface.(string)))
			}

			newGroup.GrantClientPermission(xClient, newPermissions...)
		}
	}

	/** Load groups by Uuid **/
	group_keys := []string{"provisioning_auth_groupuuid", "membership_auth_groupuuid", "auditing_auth_groupuuid"}

	for _, group_key := range group_keys {
		strUuid := d.Get(group_key).(string)
		if strUuid == "" {
			continue
		}
		if !d.HasChange(group_key) {
			continue
		}
		grpUuid, err := uuid.Parse(strUuid)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Invalid uuid",
				Detail:   fmt.Sprintf("Value `%s` is not a valid uuid for `%s`", strUuid, group_key),
			})
			continue
		}
		kh_group, err := client.Groups.GetByUUID(grpUuid)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Group does not exist",
				Detail:   fmt.Sprintf("Could not find a group with uuid `%s` for `%s`", strUuid, group_key),
			})
			continue
		}
		switch group_key {
		case "provisioning_auth_groupuuid":
			newGroup.AuthorizingGroupProvisioning = kh_group
		case "membership_auth_groupuuid":
			newGroup.AuthorizingGroupMembership = kh_group
		case "auditing_auth_groupuuid":
			newGroup.AuthorizingGroupAuditing = kh_group
		}
	}

	createdGroup, err := client.Groups.Create(newGroup)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not create group",
			Detail:   err.Error(),
		})
		return diags
	}

	err = d.Set("uuid", createdGroup.UUID)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not set uuid",
			Detail:   fmt.Sprintf("Value `%s` is not valid for uuid", createdGroup.UUID),
		})
	}
	d.SetId(strconv.FormatInt(createdGroup.Self().ID, 10))

	/*
		groupJson, _ := json.MarshalIndent(newGroup, "", "  ")

		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "New group:",
			Detail:   fmt.Sprintf("%s", groupJson),
		})
	*/

	return diags

}

func resourceGroupUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*keyhubclient.Client)
	_ = client
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Cannot update a group",
		Detail:   "Currently Keyhub doesn't allow a client to update a group after it's created, so any changes aren't stored",
	})

	dataSourceGroupRead(ctx, d, m)

	return diags
}

func resourceGroupDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*keyhubclient.Client)
	_ = client
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Cannot update a group",
		Detail:   "Currently Keyhub doesn't allow a client to delete a group. We will only delete the group from the terraform state",
	})

	d.SetId("")

	return diags
}
