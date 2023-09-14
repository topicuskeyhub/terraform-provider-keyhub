// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	keyhub "github.com/topicuskeyhub/sdk-go"
	keyhubgroup "github.com/topicuskeyhub/sdk-go/group"
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

func (r *groupResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_group"
}

func (r *groupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Group resource",
		Attributes:          resourceSchemaAttrsGroupGroup(true),
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
	var data groupGroupDataRS
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	ctx = context.WithValue(ctx, "keyhub_client", r.client)
	obj, diags := types.ObjectValueFrom(ctx, groupGroupAttrTypesRSRecurse, data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	newgroup, diags := tfObjectToTKHRSGroupGroup(ctx, true, obj)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

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
	tfGroup, diags := tkhToTFObjectRSGroupGroup(true, group)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	fillDataStructFromTFObjectRSGroupGroup(&data, tfGroup)

	tflog.Trace(ctx, "Created a new Topicus KeyHub group")
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *groupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data groupGroupDataRS
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	ctx = tflog.SetField(ctx, "keyhub_group_uuid", data.UUID.ValueString())
	ctx = context.WithValue(ctx, "keyhub_client", r.client)
	tflog.Debug(ctx, "Reading group from Topicus KeyHub by UUID")
	groups, err := r.client.Group().Get(ctx, &keyhubgroup.GroupRequestBuilderGetRequestConfiguration{
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
	tfGroup, diags := tkhToTFObjectRSGroupGroup(true, group)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	fillDataStructFromTFObjectRSGroupGroup(&data, tfGroup)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *groupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data groupGroupDataRS
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
	var data groupGroupDataRS
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
