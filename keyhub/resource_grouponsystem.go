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
	"strings"
)

func resourceGroupOnSystem() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGroupOnSystemCreate,
		UpdateContext: resourceGroupOnSystemUpdate,
		DeleteContext: resourceGroupOnSystemDelete,
		ReadContext:   resourceGroupOnSystemRead,
		Schema:        GroupOnSystemResourceSchema(),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}

}

func GroupOnSystemResourceSchema() map[string]*schema.Schema {
	resourceSchema := map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"system": {
			Type:             schema.TypeString,
			Required:         true,
			ValidateDiagFunc: validation.ToDiagFunc(validation.IsUUID),
		},
		"owner": {
			Type:             schema.TypeString,
			Required:         true,
			ValidateDiagFunc: validation.ToDiagFunc(validation.IsUUID),
		},

		"type": {
			Type:     schema.TypeString,
			Optional: true,
			ValidateDiagFunc: validation.ToDiagFunc(
				validation.StringInSlice(
					[]string{
						string(keyhubmodel.GOS_TYPE_POSIX),
						string(keyhubmodel.GOS_TYPE_GROUP_OF_NAMES),
						string(keyhubmodel.GOS_TYPE_GROUP_OF_UNIQUE_NAMES),
						string(keyhubmodel.GOS_TYPE_GROUP),
						string(keyhubmodel.GOS_TYPE_AZURE_ROLE),
						string(keyhubmodel.GOS_TYPE_AZURE_UNIFIED_GROUP),
						string(keyhubmodel.GOS_TYPE_AZURE_SECURITY_GROUP),
					},
					false,
				),
			),
		},

		"name_in_system": {
			Type:     schema.TypeString,
			Required: true,
			ValidateDiagFunc: validation.ToDiagFunc(func(i interface{}, k string) (warnings []string, errors []error) {
				v, ok := i.(string)
				if !ok {
					errors = append(errors, fmt.Errorf("expected type of %s to be string", k))
					return warnings, errors
				}
				min := 1
				max := 255
				if len(v) < min || len(v) > max {
					errors = append(errors, fmt.Errorf("expected length of %s to be in the range (%d - %d), got %s", k, min, max, v))
				}

				if v != strings.ToLower(v) {
					errors = append(errors, fmt.Errorf("expected value of %s to be in lowercase, got %s", k, v))
				}

				return warnings, errors
			}),
			DiffSuppressFunc: func(k, oldValue, newValue string, d *schema.ResourceData) bool {
				return strings.HasPrefix(oldValue, "cn="+newValue+",")
			},
		},
		"short_name_in_system": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"display_name": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"provgroup": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"group": {
						Type:             schema.TypeString,
						Required:         true,
						ValidateDiagFunc: validation.ToDiagFunc(validation.IsUUID),
					},
					"securitylevel": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: validation.ToDiagFunc(
							validation.StringInSlice(
								[]string{
									string(keyhubmodel.PRGRP_SECURITY_LEVEL_HIGH),
									string(keyhubmodel.PRGRP_SECURITY_LEVEL_MEDIUM),
									string(keyhubmodel.PRGRP_SECURITY_LEVEL_LOW),
								},
								false,
							),
						),
						Default: string(keyhubmodel.PRGRP_SECURITY_LEVEL_HIGH),
					},
					"static": {
						Type:     schema.TypeBool,
						Optional: true,
						Default:  false,
					},
				},
			},
		},
	}

	return resourceSchema
}

func resourceGroupOnSystemRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*keyhubclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	var system *keyhubmodel.ProvisionedSystem
	var gosId int64
	var err error

	if id, ok := d.GetOk("id"); ok {

		diagUnparseableId := diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not parse Id",
			Detail:   fmt.Sprintf("Id value %s", id),
		}

		idParts := strings.SplitN(id.(string), ":", 2)
		if len(idParts) != 2 {
			return append(diags, diagUnparseableId)
		}
		var systemId int64
		systemId, err = strconv.ParseInt(idParts[0], 10, 64)
		if err != nil {
			return append(diags, diagUnparseableId)
		}
		gosId, err = strconv.ParseInt(idParts[1], 10, 64)
		if err != nil {
			return append(diags, diagUnparseableId)
		}

		system, err = client.Systems.GetById(systemId)
		if err != nil {
			return append(diags, diagUnparseableId)
		}

	} else {
		systemUuid, err := uuid.Parse(d.Get("systemuuid").(string))
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "systemuuid is not a valid uuid",
				Detail:   err.Error(),
			})
			return diags
		}

		system, err = client.Systems.GetByUUID(systemUuid)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Could not GET system " + systemUuid.String(),
				Detail:   err.Error(),
			})
			return diags
		}

		gosId, err = strconv.ParseInt(d.Get("id").(string), 10, 64)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "id is not a valid int",
				Detail:   err.Error(),
			})
			return diags
		}

	}

	gos, err := client.Systems.GetGroupOnSystem(system, gosId, nil)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not GET groupOnSystem " + d.Get("id").(string),
			Detail:   fmt.Sprintf("This might mean the grouponsystem has been deleted from keyhub, \nfor more info see the raw error: %s", err.Error()),
		})
		return diags
	}

	if err := d.Set("type", string(gos.Type)); err != nil {
		diags = append(diags, NewDiagnosticSetError("type", err))
	}

	if err := d.Set("name_in_system", gos.NameInSystem); err != nil {
		diags = append(diags, NewDiagnosticSetError("name_in_system", err))
	}
	if err := d.Set("short_name_in_system", gos.ShortNameInSystem); err != nil {
		diags = append(diags, NewDiagnosticSetError("short_name_in_system", err))
	}
	if err := d.Set("display_name", gos.DisplayName); err != nil {
		diags = append(diags, NewDiagnosticSetError("display_name", err))
	}
	if gos.System != nil {
		if err := d.Set("system", gos.System.UUID); err != nil {
			diags = append(diags, NewDiagnosticSetError("system", err))
		}
	}
	if gos.Owner != nil {
		if err := d.Set("owner", gos.Owner.UUID); err != nil {
			diags = append(diags, NewDiagnosticSetError("owner", err))
		}
	}

	var provgroups []map[string]interface{}

	if gos.AdditionalObjects != nil && gos.AdditionalObjects.ProvGroups != nil && len(gos.AdditionalObjects.ProvGroups.Items) > 0 {

		for _, pvgrp := range gos.AdditionalObjects.ProvGroups.Items {

			provgroup := make(map[string]interface{})
			provgroup["group"] = pvgrp.Group.UUID
			provgroup["securitylevel"] = string(pvgrp.SecurityLevel)
			provgroup["static"] = pvgrp.StaticProvisioning

			provgroups = append(provgroups, provgroup)
		}

	}
	if err := d.Set("provgroup", provgroups); err != nil {
		diags = append(diags, NewDiagnosticSetError("provgroup", err))
	}

	d.SetId(fmt.Sprintf("%d:%d", gos.System.Self().ID, gos.Self().ID))

	return diags
}

func resourceGroupOnSystemCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*keyhubclient.Client)
	_ = client
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	var err error

	var uuidOwner uuid.UUID
	var owner *keyhubmodel.Group
	var uuidSystem uuid.UUID
	var system *keyhubmodel.ProvisionedSystem

	var gos *keyhubmodel.GroupOnSystem
	gos = keyhubmodel.NewGroupOnSystem()

	if ownerUuid, ok := d.GetOk("owner"); ok {
		uuidOwner, err = uuid.Parse(ownerUuid.(string))
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Invalid uuid",
				Detail:   fmt.Sprintf("Value `%s` is not a valid uuid for `%s`", ownerUuid.(string), "owner"),
			})
			return diags
		}

		owner, err = client.Groups.GetByUUID(uuidOwner)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Owner does not exist",
				Detail:   fmt.Sprintf("Could not find group with uuid: %s", ownerUuid.(string)),
			})
		}

	} else {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Owner uuid is required",
			Detail:   "Owner uuid is required",
		})
		return diags
	}

	if systemUuid, ok := d.GetOk("system"); ok {
		uuidSystem, err = uuid.Parse(systemUuid.(string))
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Invalid uuid",
				Detail:   fmt.Sprintf("Value `%s` is not a valid uuid for `%s`", systemUuid.(string), "system"),
			})
			return diags
		}

		system, err = client.Systems.GetByUUID(uuidSystem)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "System does not exist",
				Detail:   fmt.Sprintf("Could not find group with uuid: %s", systemUuid.(string)),
			})
		}

	} else {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "System uuid is required",
			Detail:   "System uuid is required",
		})
		return diags
	}

	if typeName, ok := d.GetOk("type"); ok {
		err = gos.SetTypeString(typeName.(string))
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Type is not valid",
				Detail:   fmt.Sprintf("Could not set type with value: %s", typeName.(string)),
			})
			return diags
		}
	}

	if value, ok := d.GetOk("name_in_system"); ok {
		gos.NameInSystem = value.(string)
	}

	if value, ok := d.GetOk("display_name"); ok {
		gos.DisplayName = value.(string)
	}

	if _, ok := d.GetOk("provgroup"); ok {
		provgroups := d.Get("provgroup").(*schema.Set)
		for _, provgrpiface := range provgroups.List() {
			provgrp := provgrpiface.(map[string]interface{})

			pggrp, err := client.Groups.GetByUUID(uuid.MustParse(provgrp["group"].(string)))
			if err != nil {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Can not configure provgroup for group",
					Detail:   fmt.Sprintf("Could retrieve group with  uuid %s", provgrp["group"].(string)),
				})
				return diags
			}

			pg := keyhubmodel.NewProvisioningGroup()
			pg.Group = pggrp.AsPrimer()
			err = pg.SetSecurityLevelString(provgrp["securitylevel"].(string))
			if err != nil {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Can not configure provgroup for group",
					Detail:   err.Error(),
				})
				return diags
			}

			pg.StaticProvisioning = provgrp["static"].(bool)

			gos.AddProvGroup(*pg)

		}
	}

	gos.Owner = owner.AsPrimer()
	gos.System = system.AsPrimer()

	newGos, err := client.Systems.CreateGroupOnSystem(gos)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not create GroupOnSystem",
			Detail:   fmt.Sprintf("Error: %s", err.Error()),
		})
		return diags
	}

	d.SetId(fmt.Sprintf("%d:%d", newGos.System.Self().ID, newGos.Self().ID))
	diags = append(diags, resourceGroupOnSystemRead(ctx, d, m)...)

	return diags
}

func resourceGroupOnSystemImportContext(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

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

func resourceGroupOnSystemUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*keyhubclient.Client)
	_ = client
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Cannot update a group on system",
		Detail:   "Currently Keyhub doesn't allow a client to update a group on system after it's created, so any changes aren't stored",
	})

	resourceGroupOnSystemRead(ctx, d, m)

	return diags
}

func resourceGroupOnSystemDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*keyhubclient.Client)
	_ = client
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Cannot delete a group on system",
		Detail:   "Currently Keyhub doesn't allow a client to delete a group on system. We will only delete the group on system from the terraform state",
	})

	d.SetId("")

	return diags
}
