// Code generated by "terraform-provider-keyhub-generator"; DO NOT EDIT.
// Copyright (c) Topicus Security B.V.
// SPDX-License-Identifier: APSL-2.0

package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/sanity-io/litter"
	keyhubreq "github.com/topicuskeyhub/sdk-go/client"
	keyhubmodels "github.com/topicuskeyhub/sdk-go/models"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &clientVaultrecordResource{}
	_ resource.ResourceWithImportState = &clientVaultrecordResource{}
	_ resource.ResourceWithConfigure   = &clientVaultrecordResource{}
)

func NewClientVaultrecordResource() resource.Resource {
	return &clientVaultrecordResource{}
}

type clientVaultrecordResource struct {
	providerData *KeyHubProviderData
}

func (r *clientVaultrecordResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = ProviderName + "_client_vaultrecord"
	tflog.Info(ctx, "Registered resource "+resp.TypeName)
}

func (r *clientVaultrecordResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: resourceSchemaAttrsClientApplicationVaultVaultRecord(true),
	}
}

func (r *clientVaultrecordResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *clientVaultrecordResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var planData clientApplicationVaultVaultRecordDataRS
	resp.Diagnostics.Append(req.Plan.Get(ctx, &planData)...)
	if resp.Diagnostics.HasError() {
		return
	}

	litter.Config.HidePrivateFields = false

	tflog.Trace(ctx, "planData: "+litter.Sdump(planData))

	var configData clientApplicationVaultVaultRecordDataRS
	resp.Diagnostics.Append(req.Config.Get(ctx, &configData)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Trace(ctx, "configData: "+litter.Sdump(configData))

	ctx = context.WithValue(ctx, keyHubClientKey, r.providerData.Client)
	planValues, diags := types.ObjectValueFrom(ctx, clientApplicationVaultVaultRecordAttrTypesRSRecurse, planData)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	configValues, diags := types.ObjectValueFrom(ctx, clientApplicationVaultVaultRecordAttrTypesRSRecurse, configData)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	newTkh, diags := tfObjectToTKHRSClientApplicationVaultVaultRecord(ctx, true, planValues, configValues)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	additionalBackup := planData.Additional
	r.providerData.Mutex.Lock()
	defer r.providerData.Mutex.Unlock()
	tflog.Info(ctx, "Creating Topicus KeyHub client_vaultrecord")
	newWrapper := keyhubmodels.NewVaultVaultRecordLinkableWrapper()
	newWrapper.SetItems([]keyhubmodels.VaultVaultRecordable{newTkh})
	tkhParent, diags := findClientClientApplicationPrimerByUUID(ctx, planData.ClientApplicationUUID.ValueStringPointer())
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	wrapper, err := r.providerData.Client.Client().ByClientidInt64(*tkhParent.GetLinks()[0].GetId()).Vault().Record().Post(
		ctx, newWrapper, &keyhubreq.ItemVaultRecordRequestBuilderPostRequestConfiguration{
			QueryParameters: &keyhubreq.ItemVaultRecordRequestBuilderPostQueryParameters{
				Additional: collectAdditional(ctx, planData, planData.Additional),
			},
		})
	tkh, diags := findFirst[keyhubmodels.VaultVaultRecordable](ctx, wrapper, "client_vaultrecord", nil, false, err)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	postState, diags := tkhToTFObjectRSClientApplicationVaultVaultRecord(true, tkh)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	postState = setAttributeValue(ctx, postState, "client_application_uuid", types.StringValue(planData.ClientApplicationUUID.ValueString()))
	postState = reorderClientApplicationVaultVaultRecord(postState, planValues, true)
	fillDataStructFromTFObjectRSClientApplicationVaultVaultRecord(&planData, postState)
	planData.Additional = additionalBackup

	resp.Diagnostics.Append(resp.State.Set(ctx, &planData)...)

	tflog.Info(ctx, "Created a new Topicus KeyHub client_vaultrecord")
	resp.Diagnostics.Append(resp.State.Set(ctx, &planData)...)
}

func (r *clientVaultrecordResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var planData clientApplicationVaultVaultRecordDataRS
	resp.Diagnostics.Append(req.State.Get(ctx, &planData)...)
	if resp.Diagnostics.HasError() {
		return
	}
	planValues, diags := types.ObjectValueFrom(ctx, clientApplicationVaultVaultRecordAttrTypesRSRecurse, planData)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	additionalBackup := planData.Additional
	r.providerData.Mutex.RLock()
	defer r.providerData.Mutex.RUnlock()
	ctx = context.WithValue(ctx, keyHubClientKey, r.providerData.Client)
	tflog.Info(ctx, "Reading client_vaultrecord from Topicus KeyHub")
	tkhParent, diags := findClientClientApplicationPrimerByUUIDOrNil(ctx, planData.ClientApplicationUUID.ValueStringPointer())
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if tkhParent == nil {
		tflog.Info(ctx, "Parent client_application not found, marking resource as removed")
		resp.State.RemoveResource(ctx)
		return
	}

	wrapper, err := r.providerData.Client.Client().ByClientidInt64(*tkhParent.GetLinks()[0].GetId()).Vault().Record().Get(
		ctx, &keyhubreq.ItemVaultRecordRequestBuilderGetRequestConfiguration{
			QueryParameters: &keyhubreq.ItemVaultRecordRequestBuilderGetQueryParameters{
				Additional: collectAdditional(ctx, planData, planData.Additional),
				Uuid:       []string{planData.UUID.ValueString()},
			},
		})

	if !isHttpStatusCodeOk(ctx, -1, err, &resp.Diagnostics) {
		return
	}

	tkh, diags := findFirst[keyhubmodels.VaultVaultRecordable](ctx, wrapper, "client_vaultrecord", planData.UUID.ValueStringPointer(), true, err)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if tkh == nil {
		tflog.Info(ctx, "client_vaultrecord not found, marking resource as removed")
		resp.State.RemoveResource(ctx)
		return
	}

	postState, diags := tkhToTFObjectRSClientApplicationVaultVaultRecord(true, tkh)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	postState = setAttributeValue(ctx, postState, "client_application_uuid", types.StringValue(planData.ClientApplicationUUID.ValueString()))
	postState = reorderClientApplicationVaultVaultRecord(postState, planValues, true)
	fillDataStructFromTFObjectRSClientApplicationVaultVaultRecord(&planData, postState)
	planData.Additional = additionalBackup

	resp.Diagnostics.Append(resp.State.Set(ctx, &planData)...)
}

func (r *clientVaultrecordResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var planData clientApplicationVaultVaultRecordDataRS
	resp.Diagnostics.Append(req.Plan.Get(ctx, &planData)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var configData clientApplicationVaultVaultRecordDataRS
	resp.Diagnostics.Append(req.Config.Get(ctx, &configData)...)
	if resp.Diagnostics.HasError() {
		return
	}

	ctx = context.WithValue(ctx, keyHubClientKey, r.providerData.Client)
	planValues, diags := types.ObjectValueFrom(ctx, clientApplicationVaultVaultRecordAttrTypesRSRecurse, planData)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	configValues, diags := types.ObjectValueFrom(ctx, clientApplicationVaultVaultRecordAttrTypesRSRecurse, configData)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	newTkh, diags := tfObjectToTKHRSClientApplicationVaultVaultRecord(ctx, true, planValues, configValues)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	additionalBackup := planData.Additional
	r.providerData.Mutex.Lock()
	defer r.providerData.Mutex.Unlock()
	tflog.Info(ctx, "Updating Topicus KeyHub client_vaultrecord")
	tkhParent, diags := findClientClientApplicationPrimerByUUID(ctx, planData.ClientApplicationUUID.ValueStringPointer())
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tkh, err := r.providerData.Client.Client().ByClientidInt64(*tkhParent.GetLinks()[0].GetId()).Vault().Record().ByRecordidInt64(getSelfLink(planData.Links).ID.ValueInt64()).Put(
		ctx, newTkh, &keyhubreq.ItemVaultRecordWithRecordItemRequestBuilderPutRequestConfiguration{
			QueryParameters: &keyhubreq.ItemVaultRecordWithRecordItemRequestBuilderPutQueryParameters{
				Additional: collectAdditional(ctx, planData, planData.Additional),
			},
		})

	if !isHttpStatusCodeOk(ctx, -1, err, &resp.Diagnostics) {
		return
	}

	postState, diags := tkhToTFObjectRSClientApplicationVaultVaultRecord(true, tkh)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	postState = setAttributeValue(ctx, postState, "client_application_uuid", types.StringValue(planData.ClientApplicationUUID.ValueString()))
	postState = reorderClientApplicationVaultVaultRecord(postState, planValues, true)
	fillDataStructFromTFObjectRSClientApplicationVaultVaultRecord(&planData, postState)
	planData.Additional = additionalBackup

	tflog.Info(ctx, "Updated a Topicus KeyHub client_vaultrecord")
	resp.Diagnostics.Append(resp.State.Set(ctx, &planData)...)
}

func (r *clientVaultrecordResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var planData clientApplicationVaultVaultRecordDataRS
	resp.Diagnostics.Append(req.State.Get(ctx, &planData)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.providerData.Mutex.Lock()
	defer r.providerData.Mutex.Unlock()
	ctx = context.WithValue(ctx, keyHubClientKey, r.providerData.Client)
	tflog.Info(ctx, "Deleting client_vaultrecord from Topicus KeyHub")
	err := r.providerData.Client.Client().ByClientidInt64(-1).Vault().Record().ByRecordidInt64(-1).WithUrl(getSelfLink(planData.Links).Href.ValueString()).Delete(ctx, nil)
	if !isHttpStatusCodeOk(ctx, 404, err, &resp.Diagnostics) {
		return
	}
	tflog.Info(ctx, "Deleted client_vaultrecord from Topicus KeyHub")
}

func (r *clientVaultrecordResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.SplitN(req.ID, ".", 2)

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: client_application_uuid.uuid. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("client_application_uuid"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("uuid"), idParts[1])...)
}
