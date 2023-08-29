// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	keyhub "github.com/topicuskeyhub/sdk-go"
	keyhubmodel "github.com/topicuskeyhub/sdk-go/models"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &groupResource{}
	_ resource.ResourceWithImportState = &groupResource{}
	_ resource.ResourceWithConfigure   = &groupResource{}
)

func NewGroupResource() resource.Resource {
	return &groupResource{}
}

// groupResource defines the resource implementation.
type groupResource struct {
	client *keyhub.KeyHubClient
}

// groupResourceModel describes the resource data model.
type groupResourceModel struct {
	ID                        types.Int64  `tfsdk:"id"`
	UUID                      types.String `tfsdk:"uuid"`
	Name                      types.String `tfsdk:"name"`
	Description               types.String `tfsdk:"description"`
	ExtendedAccess            types.String `tfsdk:"extended_access"`
	VaultRecovery             types.String `tfsdk:"vault_recovery"`
	AuditMonths               types.List   `tfsdk:"audit_months"`
	RotatingPasswordRequired  types.Bool   `tfsdk:"rotating_password_required"`
	RecordTrail               types.Bool   `tfsdk:"record_trail"`
	PrivateGroup              types.Bool   `tfsdk:"private_group"`
	HideAuditTrail            types.Bool   `tfsdk:"hide_audit_trail"`
	ApplicationAdministration types.Bool   `tfsdk:"application_administration"`
	Auditor                   types.Bool   `tfsdk:"auditor"`
	SingleManaged             types.Bool   `tfsdk:"single_managed"`
	ProvisioningAuthGroupUUID types.String `tfsdk:"provisioning_auth_groupuuid"`
	MembershipAuthGroupUUID   types.String `tfsdk:"membership_auth_groupuuid"`
	AuditingAuthGroupUUID     types.String `tfsdk:"auditing_auth_groupuuid"`
	NestedUnderGroupUUID      types.String `tfsdk:"nested_under_groupuuid"`
}

func (r *groupResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_group"
}

func (r *groupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Example resource",

		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"uuid": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"description": schema.StringAttribute{
				Optional: true,
			},
			"extended_access": schema.StringAttribute{
				Computed: true,
				Default:  stringdefault.StaticString("NOT_ALLOWED"),
			},
			"vault_recovery": schema.StringAttribute{
				Computed: true,
				Default:  stringdefault.StaticString("FULL"),
			},
			"audit_months": schema.ListAttribute{
				ElementType: types.StringType,
				Computed:    true,
			},
			"rotating_password_required": schema.BoolAttribute{
				Computed: true,
				Default:  booldefault.StaticBool(false),
			},
			"record_trail": schema.BoolAttribute{
				Computed: true,
				Default:  booldefault.StaticBool(false),
			},
			"private_group": schema.BoolAttribute{
				Computed: true,
				Default:  booldefault.StaticBool(false),
			},
			"hide_audit_trail": schema.BoolAttribute{
				Computed: true,
				Default:  booldefault.StaticBool(false),
			},
			"application_administration": schema.BoolAttribute{
				Computed: true,
				Default:  booldefault.StaticBool(false),
			},
			"auditor": schema.BoolAttribute{
				Computed: true,
				Default:  booldefault.StaticBool(false),
			},
			"single_managed": schema.BoolAttribute{
				Computed: true,
				Default:  booldefault.StaticBool(false),
			},
			"provisioning_auth_groupuuid": schema.StringAttribute{
				Computed: true,
			},
			"membership_auth_groupuuid": schema.StringAttribute{
				Computed: true,
			},
			"auditing_auth_groupuuid": schema.StringAttribute{
				Computed: true,
			},
			"nested_under_groupuuid": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *groupResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*keyhub.KeyHubClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *keyhub.KeyHubClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *groupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *groupResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newgroup := keyhubmodel.NewGroupGroup()
	newgroup.SetName(data.Name.ValueStringPointer())
	newgroup.SetDescription(data.Description.ValueStringPointer())
	extendedaccessstr := data.ExtendedAccess.ValueString()
	if extendedaccessstr != "" {
		extendedaccess, err := keyhubmodel.ParseGroupGroupExtendedAccess(extendedaccessstr)
		if err != nil {
			resp.Diagnostics.AddError("Conversion error", fmt.Sprintf("Cannot convert %s to GroupGroupExtendedAccess: %s", extendedaccessstr, err))
			return
		}
		newgroup.SetExtendedAccess(extendedaccess.(*keyhubmodel.GroupGroupExtendedAccess))
	}

	vaultrecoverystr := data.VaultRecovery.ValueString()
	if vaultrecoverystr != "" {
		vaultrecovery, err := keyhubmodel.ParseGroupVaultRecoveryAvailability(vaultrecoverystr)
		if err != nil {
			resp.Diagnostics.AddError("Conversion error", fmt.Sprintf("Cannot convert %s to GroupVaultRecoveryAvailability: %s", vaultrecoverystr, err))
			return
		}
		newgroup.SetVaultRecovery(vaultrecovery.(*keyhubmodel.GroupVaultRecoveryAvailability))
	}

	// newgroup.SetAuditMonths(data.AuditMonths.ValueStringPointer())
	newgroup.SetRotatingPasswordRequired(data.RotatingPasswordRequired.ValueBoolPointer())
	newgroup.SetRecordTrail(data.RecordTrail.ValueBoolPointer())
	newgroup.SetPrivateGroup(data.PrivateGroup.ValueBoolPointer())
	newgroup.SetHideAuditTrail(data.HideAuditTrail.ValueBoolPointer())
	newgroup.SetApplicationAdministration(data.ApplicationAdministration.ValueBoolPointer())
	newgroup.SetAuditor(data.Auditor.ValueBoolPointer())
	newgroup.SetSingleManaged(data.SingleManaged.ValueBoolPointer())
	// newgroup.SetAuthorizingGroupProvisioning(data.ProvisioningAuthGroupUUID.ValueStringPointer())
	// newgroup.SetAuthorizingGroupMembership(data.MembershipAuthGroupUUID.ValueStringPointer())
	// newgroup.SetAuthorizingGroupAuditing(data.AuditingAuthGroupUUID.ValueStringPointer())
	// newgroup.SetNestedUnder(data.NestedUnderGroupUUID.ValueStringPointer())

	ctx = tflog.SetField(ctx, "keyhub_group_name", data.Name.ValueString())
	tflog.Debug(ctx, "Creating Topicus KeyHub group")
	wrapper := keyhubmodel.NewGroupGroupLinkableWrapper()
	wrapper.SetItems([]keyhubmodel.GroupGroupable{newgroup})
	createdwrapper, err := r.client.Group().Post(ctx, wrapper, nil)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create group, got error: %s", err))
		return
	}

	group := createdwrapper.GetItems()[0]
	fillGroupModelFromResponse(ctx, data, group)

	tflog.Trace(ctx, "Created a new Topicus KeyHub group")
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *groupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data groupResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	ctx = tflog.SetField(ctx, "keyhub_group_id", data.ID.ValueInt64())
	tflog.Debug(ctx, "Reading Topicus KeyHub group by ID")

	group, err := r.client.Group().ByGroupid(strconv.FormatInt(data.ID.ValueInt64(), 10)).Get(ctx, nil)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read group, got error: %s", err))
		return
	}
	fillGroupModelFromResponse(ctx, &data, group)

	tflog.Trace(ctx, "Read a group from Topicus KeyHub")
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *groupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *groupResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := r.client.Do(httpReq)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update example, got error: %s", err))
	//     return
	// }

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *groupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *groupResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := r.client.Do(httpReq)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete example, got error: %s", err))
	//     return
	// }
}

func (r *groupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func fillGroupModelFromResponse(ctx context.Context, data *groupResourceModel, group keyhubmodel.GroupGroupable) {
	data.ID = types.Int64PointerValue(group.GetLinks()[0].GetId())
	data.UUID = types.StringPointerValue(group.GetUuid())
	data.Name = types.StringPointerValue(group.GetName())
	data.Description = types.StringPointerValue(group.GetDescription())
	data.ExtendedAccess = types.StringValue(group.GetExtendedAccess().String())
	data.AuditMonths, _ = types.ListValueFrom(ctx, types.StringType, toMonthsString(group.GetAuditConfig().GetMonths()))
	data.RotatingPasswordRequired = types.BoolPointerValue(group.GetRotatingPasswordRequired())
	data.RecordTrail = types.BoolPointerValue(group.GetRecordTrail())
	data.PrivateGroup = types.BoolPointerValue(group.GetPrivateGroup())
	data.HideAuditTrail = types.BoolPointerValue(group.GetHideAuditTrail())
	data.ApplicationAdministration = types.BoolPointerValue(group.GetApplicationAdministration())
	data.Auditor = types.BoolPointerValue(group.GetAuditor())
	data.SingleManaged = types.BoolPointerValue(group.GetSingleManaged())
	if group.GetAuthorizingGroupProvisioning() == nil {
		data.ProvisioningAuthGroupUUID = types.StringNull()
	} else {
		data.ProvisioningAuthGroupUUID = types.StringPointerValue(group.GetAuthorizingGroupProvisioning().GetUuid())
	}
	if group.GetAuthorizingGroupMembership() == nil {
		data.MembershipAuthGroupUUID = types.StringNull()
	} else {
		data.MembershipAuthGroupUUID = types.StringPointerValue(group.GetAuthorizingGroupMembership().GetUuid())
	}
	if group.GetAuthorizingGroupAuditing() == nil {
		data.AuditingAuthGroupUUID = types.StringNull()
	} else {
		data.AuditingAuthGroupUUID = types.StringPointerValue(group.GetAuthorizingGroupAuditing().GetUuid())
	}
	if group.GetNestedUnder() == nil {
		data.NestedUnderGroupUUID = types.StringNull()
	} else {
		data.NestedUnderGroupUUID = types.StringPointerValue(group.GetNestedUnder().GetUuid())
	}
}

func toMonthsString[T fmt.Stringer](months []T) []attr.Value {
	ret := make([]attr.Value, len(months))
	for _, month := range months {
		ret = append(ret, types.StringValue(month.String()))
	}
	return ret
}
