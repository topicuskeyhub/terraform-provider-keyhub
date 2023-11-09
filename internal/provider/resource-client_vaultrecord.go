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
	tflog.Info(ctx, "Registred resource "+resp.TypeName)
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
	var data clientApplicationVaultVaultRecordDataRS
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	ctx = context.WithValue(ctx, keyHubClientKey, r.providerData.Client)
	obj, diags := types.ObjectValueFrom(ctx, clientApplicationVaultVaultRecordAttrTypesRSRecurse, data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	newTkh, diags := tfObjectToTKHRSClientApplicationVaultVaultRecord(ctx, true, obj)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.providerData.Mutex.Lock()
	defer r.providerData.Mutex.Unlock()
	tflog.Info(ctx, "Creating Topicus KeyHub client_vaultrecord")
	newWrapper := keyhubmodels.NewVaultVaultRecordLinkableWrapper()
	newWrapper.SetItems([]keyhubmodels.VaultVaultRecordable{newTkh})
	tkhParent, diags := findClientClientApplicationPrimerByUUID(ctx, data.ClientApplicationUUID.ValueStringPointer())
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	wrapper, err := r.providerData.Client.Client().ByClientidInt64(*tkhParent.GetLinks()[0].GetId()).Vault().Record().Post(
		ctx, newWrapper, &keyhubreq.ItemVaultRecordRequestBuilderPostRequestConfiguration{
			QueryParameters: &keyhubreq.ItemVaultRecordRequestBuilderPostQueryParameters{
				Additional: collectAdditional(data),
			},
		})
	tkh, diags := findFirst[keyhubmodels.VaultVaultRecordable](ctx, wrapper, "client_vaultrecord", nil, err)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tf, diags := tkhToTFObjectRSClientApplicationVaultVaultRecord(true, tkh)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tf = setAttributeValue(ctx, tf, "client_application_uuid", types.StringValue(data.ClientApplicationUUID.ValueString()))
	fillDataStructFromTFObjectRSClientApplicationVaultVaultRecord(&data, tf)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

	tflog.Info(ctx, "Created a new Topicus KeyHub client_vaultrecord")
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *clientVaultrecordResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data clientApplicationVaultVaultRecordDataRS
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.providerData.Mutex.RLock()
	defer r.providerData.Mutex.RUnlock()
	ctx = context.WithValue(ctx, keyHubClientKey, r.providerData.Client)
	tflog.Info(ctx, "Reading client_vaultrecord from Topicus KeyHub")
	tkhParent, diags := findClientClientApplicationPrimerByUUID(ctx, data.ClientApplicationUUID.ValueStringPointer())
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tkh, err := r.providerData.Client.Client().ByClientidInt64(*tkhParent.GetLinks()[0].GetId()).Vault().Record().ByRecordidInt64(getSelfLink(data.Links).ID.ValueInt64()).Get(
		ctx, &keyhubreq.ItemVaultRecordWithRecordItemRequestBuilderGetRequestConfiguration{
			QueryParameters: &keyhubreq.ItemVaultRecordWithRecordItemRequestBuilderGetQueryParameters{
				Additional: collectAdditional(data),
			},
		})

	if !isHttpStatusCodeOk(ctx, -1, err, &resp.Diagnostics) {
		return
	}

	tf, diags := tkhToTFObjectRSClientApplicationVaultVaultRecord(true, tkh)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tf = setAttributeValue(ctx, tf, "client_application_uuid", types.StringValue(data.ClientApplicationUUID.ValueString()))
	fillDataStructFromTFObjectRSClientApplicationVaultVaultRecord(&data, tf)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *clientVaultrecordResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data clientApplicationVaultVaultRecordDataRS
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	ctx = context.WithValue(ctx, keyHubClientKey, r.providerData.Client)
	obj, diags := types.ObjectValueFrom(ctx, clientApplicationVaultVaultRecordAttrTypesRSRecurse, data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	newTkh, diags := tfObjectToTKHRSClientApplicationVaultVaultRecord(ctx, true, obj)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.providerData.Mutex.Lock()
	defer r.providerData.Mutex.Unlock()
	tflog.Info(ctx, "Updating Topicus KeyHub client_vaultrecord")
	tkhParent, diags := findClientClientApplicationPrimerByUUID(ctx, data.ClientApplicationUUID.ValueStringPointer())
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tkh, err := r.providerData.Client.Client().ByClientidInt64(*tkhParent.GetLinks()[0].GetId()).Vault().Record().ByRecordidInt64(getSelfLink(data.Links).ID.ValueInt64()).Put(
		ctx, newTkh, &keyhubreq.ItemVaultRecordWithRecordItemRequestBuilderPutRequestConfiguration{
			QueryParameters: &keyhubreq.ItemVaultRecordWithRecordItemRequestBuilderPutQueryParameters{
				Additional: collectAdditional(data),
			},
		})

	if !isHttpStatusCodeOk(ctx, -1, err, &resp.Diagnostics) {
		return
	}

	tf, diags := tkhToTFObjectRSClientApplicationVaultVaultRecord(true, tkh)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tf = setAttributeValue(ctx, tf, "client_application_uuid", types.StringValue(data.ClientApplicationUUID.ValueString()))
	fillDataStructFromTFObjectRSClientApplicationVaultVaultRecord(&data, tf)

	tflog.Info(ctx, "Updated a Topicus KeyHub client_vaultrecord")
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *clientVaultrecordResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data clientApplicationVaultVaultRecordDataRS
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.providerData.Mutex.Lock()
	defer r.providerData.Mutex.Unlock()
	ctx = context.WithValue(ctx, keyHubClientKey, r.providerData.Client)
	tflog.Info(ctx, "Deleting client_vaultrecord from Topicus KeyHub")
	err := r.providerData.Client.Client().ByClientidInt64(-1).Vault().Record().ByRecordidInt64(-1).WithUrl(getSelfLink(data.Links).Href.ValueString()).Delete(ctx, nil)
	if !isHttpStatusCodeOk(ctx, 404, err, &resp.Diagnostics) {
		return
	}
	tflog.Info(ctx, "Deleted client_vaultrecord from Topicus KeyHub")
}

func (r *clientVaultrecordResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("uuid"), req, resp)
}
