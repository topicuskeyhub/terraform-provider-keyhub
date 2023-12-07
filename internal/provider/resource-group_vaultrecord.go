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
	keyhubreq "github.com/topicuskeyhub/sdk-go/group"
	keyhubmodels "github.com/topicuskeyhub/sdk-go/models"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &groupVaultrecordResource{}
	_ resource.ResourceWithImportState = &groupVaultrecordResource{}
	_ resource.ResourceWithConfigure   = &groupVaultrecordResource{}
)

func NewGroupVaultrecordResource() resource.Resource {
	return &groupVaultrecordResource{}
}

type groupVaultrecordResource struct {
	providerData *KeyHubProviderData
}

func (r *groupVaultrecordResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = ProviderName + "_group_vaultrecord"
	tflog.Info(ctx, "Registred resource "+resp.TypeName)
}

func (r *groupVaultrecordResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: resourceSchemaAttrsGroupVaultVaultRecord(true),
	}
}

func (r *groupVaultrecordResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *groupVaultrecordResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data groupVaultVaultRecordDataRS
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	ctx = context.WithValue(ctx, keyHubClientKey, r.providerData.Client)
	obj, diags := types.ObjectValueFrom(ctx, groupVaultVaultRecordAttrTypesRSRecurse, data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	newTkh, diags := tfObjectToTKHRSGroupVaultVaultRecord(ctx, true, obj)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	additionalBackup := data.Additional
	r.providerData.Mutex.Lock()
	defer r.providerData.Mutex.Unlock()
	tflog.Info(ctx, "Creating Topicus KeyHub group_vaultrecord")
	newWrapper := keyhubmodels.NewVaultVaultRecordLinkableWrapper()
	newWrapper.SetItems([]keyhubmodels.VaultVaultRecordable{newTkh})
	tkhParent, diags := findGroupGroupPrimerByUUID(ctx, data.GroupUUID.ValueStringPointer())
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	wrapper, err := r.providerData.Client.Group().ByGroupidInt64(*tkhParent.GetLinks()[0].GetId()).Vault().Record().Post(
		ctx, newWrapper, &keyhubreq.ItemVaultRecordRequestBuilderPostRequestConfiguration{
			QueryParameters: &keyhubreq.ItemVaultRecordRequestBuilderPostQueryParameters{
				Additional: collectAdditional(ctx, data, data.Additional),
			},
		})
	tkh, diags := findFirst[keyhubmodels.VaultVaultRecordable](ctx, wrapper, "group_vaultrecord", nil, false, err)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tf, diags := tkhToTFObjectRSGroupVaultVaultRecord(true, tkh)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tf = setAttributeValue(ctx, tf, "group_uuid", types.StringValue(data.GroupUUID.ValueString()))
	fillDataStructFromTFObjectRSGroupVaultVaultRecord(&data, tf)
	data.Additional = additionalBackup

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

	tflog.Info(ctx, "Created a new Topicus KeyHub group_vaultrecord")
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *groupVaultrecordResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data groupVaultVaultRecordDataRS
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	additionalBackup := data.Additional
	r.providerData.Mutex.RLock()
	defer r.providerData.Mutex.RUnlock()
	ctx = context.WithValue(ctx, keyHubClientKey, r.providerData.Client)
	tflog.Info(ctx, "Reading group_vaultrecord from Topicus KeyHub")
	tkhParent, diags := findGroupGroupPrimerByUUIDOrNil(ctx, data.GroupUUID.ValueStringPointer())
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if tkhParent == nil {
		tflog.Info(ctx, "Parent group not found, marking resource as removed")
		resp.State.RemoveResource(ctx)
		return
	}

	wrapper, err := r.providerData.Client.Group().ByGroupidInt64(*tkhParent.GetLinks()[0].GetId()).Vault().Record().Get(
		ctx, &keyhubreq.ItemVaultRecordRequestBuilderGetRequestConfiguration{
			QueryParameters: &keyhubreq.ItemVaultRecordRequestBuilderGetQueryParameters{
				Additional: collectAdditional(ctx, data, data.Additional),
				Uuid:       []string{data.UUID.ValueString()},
			},
		})

	if !isHttpStatusCodeOk(ctx, -1, err, &resp.Diagnostics) {
		return
	}

	tkh, diags := findFirst[keyhubmodels.VaultVaultRecordable](ctx, wrapper, "group_vaultrecord", data.UUID.ValueStringPointer(), true, err)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if tkh == nil {
		tflog.Info(ctx, "group_vaultrecord not found, marking resource as removed")
		resp.State.RemoveResource(ctx)
		return
	}

	tf, diags := tkhToTFObjectRSGroupVaultVaultRecord(true, tkh)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tf = setAttributeValue(ctx, tf, "group_uuid", types.StringValue(data.GroupUUID.ValueString()))
	fillDataStructFromTFObjectRSGroupVaultVaultRecord(&data, tf)
	data.Additional = additionalBackup

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *groupVaultrecordResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data groupVaultVaultRecordDataRS
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	ctx = context.WithValue(ctx, keyHubClientKey, r.providerData.Client)
	obj, diags := types.ObjectValueFrom(ctx, groupVaultVaultRecordAttrTypesRSRecurse, data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	newTkh, diags := tfObjectToTKHRSGroupVaultVaultRecord(ctx, true, obj)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	additionalBackup := data.Additional
	r.providerData.Mutex.Lock()
	defer r.providerData.Mutex.Unlock()
	tflog.Info(ctx, "Updating Topicus KeyHub group_vaultrecord")
	tkhParent, diags := findGroupGroupPrimerByUUID(ctx, data.GroupUUID.ValueStringPointer())
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tkh, err := r.providerData.Client.Group().ByGroupidInt64(*tkhParent.GetLinks()[0].GetId()).Vault().Record().ByRecordidInt64(getSelfLink(data.Links).ID.ValueInt64()).Put(
		ctx, newTkh, &keyhubreq.ItemVaultRecordWithRecordItemRequestBuilderPutRequestConfiguration{
			QueryParameters: &keyhubreq.ItemVaultRecordWithRecordItemRequestBuilderPutQueryParameters{
				Additional: collectAdditional(ctx, data, data.Additional),
			},
		})

	if !isHttpStatusCodeOk(ctx, -1, err, &resp.Diagnostics) {
		return
	}

	tf, diags := tkhToTFObjectRSGroupVaultVaultRecord(true, tkh)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tf = setAttributeValue(ctx, tf, "group_uuid", types.StringValue(data.GroupUUID.ValueString()))
	fillDataStructFromTFObjectRSGroupVaultVaultRecord(&data, tf)
	data.Additional = additionalBackup

	tflog.Info(ctx, "Updated a Topicus KeyHub group_vaultrecord")
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *groupVaultrecordResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data groupVaultVaultRecordDataRS
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.providerData.Mutex.Lock()
	defer r.providerData.Mutex.Unlock()
	ctx = context.WithValue(ctx, keyHubClientKey, r.providerData.Client)
	tflog.Info(ctx, "Deleting group_vaultrecord from Topicus KeyHub")
	err := r.providerData.Client.Group().ByGroupidInt64(-1).Vault().Record().ByRecordidInt64(-1).WithUrl(getSelfLink(data.Links).Href.ValueString()).Delete(ctx, nil)
	if !isHttpStatusCodeOk(ctx, 404, err, &resp.Diagnostics) {
		return
	}
	tflog.Info(ctx, "Deleted group_vaultrecord from Topicus KeyHub")
}

func (r *groupVaultrecordResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.SplitN(req.ID, ".", 2)

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: group_uuid.uuid. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("group_uuid"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("uuid"), idParts[1])...)
}
