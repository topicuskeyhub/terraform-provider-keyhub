// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	keyhub "github.com/topicuskeyhub/sdk-go"
	keyhubgroup "github.com/topicuskeyhub/sdk-go/group"
	keyhubmodel "github.com/topicuskeyhub/sdk-go/models"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = &groupDataSource{}
	_ datasource.DataSourceWithConfigure = &groupDataSource{}
)

func NewGroupDataSource() datasource.DataSource {
	return &groupDataSource{}
}

// groupDataSource defines the data source implementation.
type groupDataSource struct {
	client *keyhub.KeyHubClient
}

// GroupDataSourceModel describes the data source data model.
type GroupDataSourceModel struct {
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

func (d *groupDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_group"
}

func (d *groupDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Group data source",

		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed: true,
			},
			"uuid": schema.StringAttribute{
				Required: true,
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"description": schema.StringAttribute{
				Computed: true,
			},
			"extended_access": schema.StringAttribute{
				Computed: true,
			},
			"vault_recovery": schema.StringAttribute{
				Computed: true,
			},
			"audit_months": schema.ListAttribute{
				ElementType: types.StringType,
				Computed:    true,
			},
			"rotating_password_required": schema.BoolAttribute{
				Computed: true,
			},
			"record_trail": schema.BoolAttribute{
				Computed: true,
			},
			"private_group": schema.BoolAttribute{
				Computed: true,
			},
			"hide_audit_trail": schema.BoolAttribute{
				Computed: true,
			},
			"application_administration": schema.BoolAttribute{
				Computed: true,
			},
			"auditor": schema.BoolAttribute{
				Computed: true,
			},
			"single_managed": schema.BoolAttribute{
				Computed: true,
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

func (d *groupDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*keyhub.KeyHubClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *keyhub.KeyHubClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

func (d *groupDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data GroupDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	groups, err := d.client.Group().Get(ctx, &keyhubgroup.GroupRequestBuilderGetRequestConfiguration{
		QueryParameters: &keyhubgroup.GroupRequestBuilderGetQueryParameters{
			Uuid: []string{data.UUID.ValueString()},
		},
	})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read group, got error: %s", err))
		return
	}
	if len(groups.GetItems()) == 0 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to find group with UUID %s", data.UUID.ValueString()))
		return
	}
	group := groups.GetItems()[0]

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

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func toMonthsString(months []keyhubmodel.Month) []string {
	ret := make([]string, len(months))
	for _, month := range months {
		ret = append(ret, month.String())
	}
	return ret
}
