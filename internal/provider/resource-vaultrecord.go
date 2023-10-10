// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	keyhub "github.com/topicuskeyhub/sdk-go"
	keyhubgroup "github.com/topicuskeyhub/sdk-go/group"
	keyhubmodels "github.com/topicuskeyhub/sdk-go/models"
	keyhubvaultrecord "github.com/topicuskeyhub/sdk-go/vaultrecord"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &vaultRecordResource{}
	_ resource.ResourceWithImportState = &vaultRecordResource{}
	_ resource.ResourceWithConfigure   = &vaultRecordResource{}
)

func NewVaultRecordResource() resource.Resource {
	return &vaultRecordResource{}
}

// groupResource defines the resource implementation.
type vaultRecordResource struct {
	client *keyhub.KeyHubClient
}

func (r *vaultRecordResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = ProviderName + "_vaultrecord"
	tflog.Info(ctx, "Registred resource "+resp.TypeName)
}

func (r *vaultRecordResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: resourceSchemaAttrsGroupVaultVaultRecord(true),
	}
}

func (r *vaultRecordResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	r.client = client
}

func (r *vaultRecordResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data groupVaultVaultRecordDataRS
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	ctx = context.WithValue(ctx, keyHubClientKey, r.client)
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

	tkhGroup, diags := findGroupGroupPrimerByUUID(ctx, data.GroupUUID.ValueStringPointer())
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Creating Topicus KeyHub vaultrecord")
	newWrapper := keyhubmodels.NewVaultVaultRecordLinkableWrapper()
	newWrapper.SetItems([]keyhubmodels.VaultVaultRecordable{newTkh})
	wrapper, err := r.client.Group().ByGroupidInt64(*tkhGroup.GetLinks()[0].GetId()).Vault().Record().Post(
		ctx, newWrapper, &keyhubgroup.ItemVaultRecordRequestBuilderPostRequestConfiguration{
			QueryParameters: &keyhubgroup.ItemVaultRecordRequestBuilderPostQueryParameters{
				Additional: collectAdditional(data.AdditionalObjects),
			},
		})
	tkh, diags := findFirst[keyhubmodels.VaultVaultRecordable](ctx, wrapper, "vaultrecord", nil, err)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tf, diags := tkhToTFObjectRSGroupVaultVaultRecord(true, tkh)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tf = fillGroupUuid(ctx, tf, data)
	fillDataStructFromTFObjectRSGroupVaultVaultRecord(&data, tf)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

	tflog.Info(ctx, "Created a new Topicus KeyHub vaultrecord")
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *vaultRecordResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data groupVaultVaultRecordDataRS
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	uuid := data.UUID.ValueString()
	ctx = context.WithValue(ctx, keyHubClientKey, r.client)
	tflog.Info(ctx, "Reading vaultrecord from Topicus KeyHub by UUID")
	wrapper, err := r.client.Vaultrecord().Get(ctx, &keyhubvaultrecord.VaultrecordRequestBuilderGetRequestConfiguration{
		QueryParameters: &keyhubvaultrecord.VaultrecordRequestBuilderGetQueryParameters{
			Uuid:       []string{uuid},
			Additional: collectAdditional(data.AdditionalObjects),
		},
	})
	tkh, diags := findFirst[keyhubmodels.VaultVaultRecordable](ctx, wrapper, "vaultrecord", &uuid, err)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tf, diags := tkhToTFObjectRSGroupVaultVaultRecord(true, tkh)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tf = fillGroupUuid(ctx, tf, data)
	fillDataStructFromTFObjectRSGroupVaultVaultRecord(&data, tf)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *vaultRecordResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data groupVaultVaultRecordDataRS
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	ctx = context.WithValue(ctx, keyHubClientKey, r.client)
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

	tkhGroup, diags := findGroupGroupPrimerByUUID(ctx, data.GroupUUID.ValueStringPointer())
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Updating Topicus KeyHub vaultrecord")
	tkh, err := r.client.Group().ByGroupidInt64(*tkhGroup.GetLinks()[0].GetId()).Vault().Record().ByIdInt64(getSelfLink(data.Links).ID.ValueInt64()).Put(
		ctx, newTkh, &keyhubgroup.ItemVaultRecordRecordItemRequestBuilderPutRequestConfiguration{
			QueryParameters: &keyhubgroup.ItemVaultRecordRecordItemRequestBuilderPutQueryParameters{
				Additional: collectAdditional(data.AdditionalObjects),
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

	tf = fillGroupUuid(ctx, tf, data)
	fillDataStructFromTFObjectRSGroupVaultVaultRecord(&data, tf)

	tflog.Info(ctx, "Updated a Topicus KeyHub vaultrecord")
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *vaultRecordResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data groupGroupDataRS
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	ctx = context.WithValue(ctx, keyHubClientKey, r.client)
	tflog.Info(ctx, "Deleting vaultrecord from Topicus KeyHub")
	err := r.client.Group().ByGroupidInt64(-1).Vault().Record().ByIdInt64(-1).WithUrl(getSelfLink(data.Links).Href.ValueString()).Delete(ctx, nil)
	if !isHttpStatusCodeOk(ctx, 404, err, &resp.Diagnostics) {
		return
	}
	tflog.Info(ctx, "Deleted vaultrecord from Topicus KeyHub")
}

func (r *vaultRecordResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("uuid"), req, resp)
}

func getSelfLink(linksAttr basetypes.ListValue) restLinkDataRS {
	var links restLinkDataRS
	fillDataStructFromTFObjectRSRestLink(&links, linksAttr.Elements()[0].(basetypes.ObjectValue))
	return links
}

func isHttpStatusCodeOk(ctx context.Context, status int32, err error, diags *diag.Diagnostics) bool {
	if err != nil {
		report, ok := err.(keyhubmodels.ErrorReportable)
		if !ok || *report.GetCode() != status {
			diags.AddError("Client Error", fmt.Sprintf("Unexpected status code: %s", errorReportToString(ctx, err)))
			return false
		}
	}
	return true
}

func fillGroupUuid(ctx context.Context, tf basetypes.ObjectValue, data groupVaultVaultRecordDataRS) basetypes.ObjectValue {
	obj := tf.Attributes()
	obj["group_uuid"] = types.StringValue(data.GroupUUID.ValueString())
	return types.ObjectValueMust(tf.AttributeTypes(ctx), obj)
}

func collectAdditional(additionalObjects basetypes.ObjectValue) []string {
	ret := make([]string, 0)
	for name, attr := range additionalObjects.Attributes() {
		if !attr.IsNull() && !attr.IsUnknown() {
			ret = append(ret, name)
		}
	}
	return ret
}
