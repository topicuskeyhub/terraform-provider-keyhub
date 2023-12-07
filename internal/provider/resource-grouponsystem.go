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
	keyhubmodels "github.com/topicuskeyhub/sdk-go/models"
	keyhubreq "github.com/topicuskeyhub/sdk-go/system"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &grouponsystemResource{}
	_ resource.ResourceWithImportState = &grouponsystemResource{}
	_ resource.ResourceWithConfigure   = &grouponsystemResource{}
)

func NewGrouponsystemResource() resource.Resource {
	return &grouponsystemResource{}
}

type grouponsystemResource struct {
	providerData *KeyHubProviderData
}

func (r *grouponsystemResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = ProviderName + "_grouponsystem"
	tflog.Info(ctx, "Registred resource "+resp.TypeName)
}

func (r *grouponsystemResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: resourceSchemaAttrsNestedProvisioningGroupOnSystem(true),
	}
}

func (r *grouponsystemResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *grouponsystemResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data nestedProvisioningGroupOnSystemDataRS
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	ctx = context.WithValue(ctx, keyHubClientKey, r.providerData.Client)
	obj, diags := types.ObjectValueFrom(ctx, nestedProvisioningGroupOnSystemAttrTypesRSRecurse, data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	newTkh, diags := tfObjectToTKHRSNestedProvisioningGroupOnSystem(ctx, true, obj)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	additionalBackup := data.Additional
	r.providerData.Mutex.Lock()
	defer r.providerData.Mutex.Unlock()
	tflog.Info(ctx, "Creating Topicus KeyHub grouponsystem")
	newWrapper := keyhubmodels.NewProvisioningGroupOnSystemLinkableWrapper()
	newWrapper.SetItems([]keyhubmodels.ProvisioningGroupOnSystemable{newTkh})
	tkhParent, diags := findProvisioningProvisionedSystemPrimerByUUID(ctx, data.ProvisionedSystemUUID.ValueStringPointer())
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	wrapper, err := r.providerData.Client.System().BySystemidInt64(*tkhParent.GetLinks()[0].GetId()).Group().Post(
		ctx, newWrapper, &keyhubreq.ItemGroupRequestBuilderPostRequestConfiguration{
			QueryParameters: &keyhubreq.ItemGroupRequestBuilderPostQueryParameters{
				Additional: collectAdditional(ctx, data, data.Additional),
			},
		})
	tkh, diags := findFirst[keyhubmodels.ProvisioningGroupOnSystemable](ctx, wrapper, "grouponsystem", nil, false, err)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tf, diags := tkhToTFObjectRSNestedProvisioningGroupOnSystem(true, tkh)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tf = setAttributeValue(ctx, tf, "provisioned_system_uuid", types.StringValue(data.ProvisionedSystemUUID.ValueString()))
	fillDataStructFromTFObjectRSNestedProvisioningGroupOnSystem(&data, tf)
	data.Additional = additionalBackup

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

	tflog.Info(ctx, "Created a new Topicus KeyHub grouponsystem")
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *grouponsystemResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data nestedProvisioningGroupOnSystemDataRS
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	additionalBackup := data.Additional
	r.providerData.Mutex.RLock()
	defer r.providerData.Mutex.RUnlock()
	ctx = context.WithValue(ctx, keyHubClientKey, r.providerData.Client)
	tflog.Info(ctx, "Reading grouponsystem from Topicus KeyHub")
	tkhParent, diags := findProvisioningProvisionedSystemPrimerByUUIDOrNil(ctx, data.ProvisionedSystemUUID.ValueStringPointer())
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if tkhParent == nil {
		tflog.Info(ctx, "Parent provisioned_system not found, marking resource as removed")
		resp.State.RemoveResource(ctx)
		return
	}

	wrapper, err := r.providerData.Client.System().BySystemidInt64(*tkhParent.GetLinks()[0].GetId()).Group().Get(
		ctx, &keyhubreq.ItemGroupRequestBuilderGetRequestConfiguration{
			QueryParameters: &keyhubreq.ItemGroupRequestBuilderGetQueryParameters{
				Additional:   collectAdditional(ctx, data, data.Additional),
				NameInSystem: []string{data.NameInSystem.ValueString()},
			},
		})

	if !isHttpStatusCodeOk(ctx, -1, err, &resp.Diagnostics) {
		return
	}

	tkh, diags := findFirst[keyhubmodels.ProvisioningGroupOnSystemable](ctx, wrapper, "grouponsystem", data.NameInSystem.ValueStringPointer(), true, err)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if tkh == nil {
		tflog.Info(ctx, "grouponsystem not found, marking resource as removed")
		resp.State.RemoveResource(ctx)
		return
	}

	tf, diags := tkhToTFObjectRSNestedProvisioningGroupOnSystem(true, tkh)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tf = setAttributeValue(ctx, tf, "provisioned_system_uuid", types.StringValue(data.ProvisionedSystemUUID.ValueString()))
	fillDataStructFromTFObjectRSNestedProvisioningGroupOnSystem(&data, tf)
	data.Additional = additionalBackup

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *grouponsystemResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError("Cannot update a grouponsystem", "Topicus KeyHub does not support updating a grouponsystem via Terraform. The requested changes are not applied.")
}

func (r *grouponsystemResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.Diagnostics.AddError("Cannot delete a grouponsystem", "Topicus KeyHub does not support deleting a grouponsystem via Terraform. The requested changes are not applied.")
}

func (r *grouponsystemResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.SplitN(req.ID, ".", 2)

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: provisioned_system_uuid.name_in_system. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("provisioned_system_uuid"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("name_in_system"), idParts[1])...)
}
