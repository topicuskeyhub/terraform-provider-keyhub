// Code generated by "terraform-provider-keyhub-generator"; DO NOT EDIT.
// Copyright (c) Topicus Security B.V.
// SPDX-License-Identifier: APSL-2.0

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/sanity-io/litter"
	keyhubreq "github.com/topicuskeyhub/sdk-go/group"
	keyhubmodels "github.com/topicuskeyhub/sdk-go/models"
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

type groupResource struct {
	providerData *KeyHubProviderData
}

func (r *groupResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = ProviderName + "_group"
	tflog.Info(ctx, "Registered resource "+resp.TypeName)
}

func (r *groupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: resourceSchemaAttrsGroupGroup(true),
	}
}

func (r *groupResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	providerData, ok := req.ProviderData.(*KeyHubProviderData)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *keyhub.KeyHubClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.providerData = providerData
}

func (r *groupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var planData groupGroupDataRS
	resp.Diagnostics.Append(req.Plan.Get(ctx, &planData)...)
	if resp.Diagnostics.HasError() {
		return
	}

	litter.Config.HidePrivateFields = false

	tflog.Trace(ctx, "planData: "+litter.Sdump(planData))

	var configData groupGroupDataRS
	resp.Diagnostics.Append(req.Config.Get(ctx, &configData)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Trace(ctx, "configData: "+litter.Sdump(configData))

	ctx = context.WithValue(ctx, keyHubClientKey, r.providerData.Client)
	planValues, diags := types.ObjectValueFrom(ctx, groupGroupAttrTypesRSRecurse, planData)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	configValues, diags := types.ObjectValueFrom(ctx, groupGroupAttrTypesRSRecurse, configData)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	newTkh, diags := tfObjectToTKHRSGroupGroup(ctx, true, planValues, configValues)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	additionalBackup := planData.Additional
	r.providerData.Mutex.Lock()
	defer r.providerData.Mutex.Unlock()
	tflog.Info(ctx, "Creating Topicus KeyHub group")
	newWrapper := keyhubmodels.NewGroupGroupLinkableWrapper()
	newWrapper.SetItems([]keyhubmodels.GroupGroupable{newTkh})
	wrapper, err := r.providerData.Client.Group().Post(
		ctx, newWrapper, &keyhubreq.GroupRequestBuilderPostRequestConfiguration{
			QueryParameters: &keyhubreq.GroupRequestBuilderPostQueryParameters{
				Additional: collectAdditional(ctx, planData, planData.Additional),
			},
		})
	tkh, diags := findFirst[keyhubmodels.GroupGroupable](ctx, wrapper, "group", nil, false, err)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	postState, diags := tkhToTFObjectRSGroupGroup(true, tkh)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	postState = reorderGroupGroup(postState, planValues, true)
	fillDataStructFromTFObjectRSGroupGroup(&planData, postState)
	planData.Additional = additionalBackup

	resp.Diagnostics.Append(resp.State.Set(ctx, &planData)...)

	tflog.Info(ctx, "Created a new Topicus KeyHub group")
	resp.Diagnostics.Append(resp.State.Set(ctx, &planData)...)
}

func (r *groupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var planData groupGroupDataRS
	resp.Diagnostics.Append(req.State.Get(ctx, &planData)...)
	if resp.Diagnostics.HasError() {
		return
	}
	planValues, diags := types.ObjectValueFrom(ctx, groupGroupAttrTypesRSRecurse, planData)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	additionalBackup := planData.Additional
	r.providerData.Mutex.RLock()
	defer r.providerData.Mutex.RUnlock()
	ctx = context.WithValue(ctx, keyHubClientKey, r.providerData.Client)
	tflog.Info(ctx, "Reading group from Topicus KeyHub")
	wrapper, err := r.providerData.Client.Group().Get(
		ctx, &keyhubreq.GroupRequestBuilderGetRequestConfiguration{
			QueryParameters: &keyhubreq.GroupRequestBuilderGetQueryParameters{
				Additional: collectAdditional(ctx, planData, planData.Additional),
				Uuid:       []string{planData.UUID.ValueString()},
			},
		})

	if !isHttpStatusCodeOk(ctx, -1, err, &resp.Diagnostics) {
		return
	}

	tkh, diags := findFirst[keyhubmodels.GroupGroupable](ctx, wrapper, "group", planData.UUID.ValueStringPointer(), true, err)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if tkh == nil {
		tflog.Info(ctx, "group not found, marking resource as removed")
		resp.State.RemoveResource(ctx)
		return
	}

	postState, diags := tkhToTFObjectRSGroupGroup(true, tkh)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	postState = reorderGroupGroup(postState, planValues, true)
	fillDataStructFromTFObjectRSGroupGroup(&planData, postState)
	planData.Additional = additionalBackup

	resp.Diagnostics.Append(resp.State.Set(ctx, &planData)...)
}

func (r *groupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError("Cannot update a group", "Topicus KeyHub does not support updating a group via Terraform. The requested changes are not applied.")
}

func (r *groupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.Diagnostics.AddError("Cannot delete a group", "Topicus KeyHub does not support deleting a group via Terraform. The requested changes are not applied.")
}

func (r *groupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("uuid"), req, resp)
}
