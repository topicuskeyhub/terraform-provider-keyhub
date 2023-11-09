// Code generated by "terraform-provider-keyhub-generator"; DO NOT EDIT.
// Copyright (c) Topicus Security B.V.
// SPDX-License-Identifier: APSL-2.0

package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	keyhubreq "github.com/topicuskeyhub/sdk-go/client"
	keyhubmodels "github.com/topicuskeyhub/sdk-go/models"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = &clientDataSource{}
	_ datasource.DataSourceWithConfigure = &clientDataSource{}
)

func NewClientDataSource() datasource.DataSource {
	return &clientDataSource{}
}

type clientDataSource struct {
	providerData *KeyHubProviderData
}

func (d *clientDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = ProviderName + "_client"
	log.Printf("Registered data source %s", resp.TypeName)
}

func (d *clientDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: dataSourceSchemaAttrsClientClientApplication(true),
	}
}

func (d *clientDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

	d.providerData = providerData
}

func (d *clientDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data clientClientApplicationDataDS
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading client from Topicus KeyHub by UUID")
	additionalBackup := data.Additional
	additional, _ := tfToSlice(data.Additional, func(val attr.Value, diags *diag.Diagnostics) string {
		return val.(basetypes.StringValue).ValueString()
	})
	uuid := data.UUID.ValueString()

	d.providerData.Mutex.RLock()
	defer d.providerData.Mutex.RUnlock()
	wrapper, err := d.providerData.Client.Client().Get(ctx, &keyhubreq.ClientRequestBuilderGetRequestConfiguration{
		QueryParameters: &keyhubreq.ClientRequestBuilderGetQueryParameters{
			Uuid:       []string{uuid},
			Additional: additional,
		},
	})

	tkh, diags := findFirst[keyhubmodels.ClientClientApplicationable](ctx, wrapper, "client", &uuid, false, err)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tf, diags := tkhToTFObjectDSClientClientApplication(true, tkh)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	fillDataStructFromTFObjectDSClientClientApplication(&data, tf)
	data.Additional = additionalBackup

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
