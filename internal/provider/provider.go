// Copyright (c) Topicus Security B.V.
// SPDX-License-Identifier: APSL-2.0

package provider

import (
	"context"
	"net/http"
	"os"
	"sync"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	keyhub "github.com/topicuskeyhub/sdk-go"
)

// Ensure KeyHubProvider satisfies various provider interfaces.
var _ provider.Provider = &KeyHubProvider{}

// KeyHubProvider defines the provider implementation.
type KeyHubProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

const ProviderName = "keyhub"

type KeyHubProviderModel struct {
	Issuer       types.String `tfsdk:"issuer"`
	ClientID     types.String `tfsdk:"clientid"`
	ClientSecret types.String `tfsdk:"clientsecret"`
}

type KeyHubProviderData struct {
	Client *keyhub.KeyHubClient
	Mutex  sync.RWMutex
}

func (p *KeyHubProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = ProviderName
	resp.Version = p.version
	tflog.Info(ctx, "Provider name set to "+resp.TypeName)
}

func (p *KeyHubProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"issuer": schema.StringAttribute{
				Optional: true,
			},
			"clientid": schema.StringAttribute{
				Optional: true,
			},
			"clientsecret": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
		},
	}
}

func (p *KeyHubProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config KeyHubProviderModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if config.Issuer.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("issuer"),
			"Unknown Topicus KeyHub Issuer URI",
			"The provider cannot create the Topicus KeyHub API client as there is an unknown configuration value for the Topicus KeyHub Issuer URI. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the KEYHUB_ISSUER environment variable.",
		)
	}
	if config.ClientID.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("clientid"),
			"Unknown Topicus KeyHub Client ID",
			"The provider cannot create the Topicus KeyHub API client as there is an unknown configuration value for the Topicus KeyHub Client ID. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the KEYHUB_CLIENTID environment variable.",
		)
	}
	if config.ClientSecret.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("clientsecret"),
			"Unknown Topicus KeyHub Client Secret",
			"The provider cannot create the Topicus KeyHub API client as there is an unknown configuration value for the Topicus KeyHub Client Secret. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the KEYHUB_CLIENTSECRET environment variable.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	issuer := os.Getenv("KEYHUB_ISSUER")
	clientid := os.Getenv("KEYHUB_CLIENTID")
	clientsecret := os.Getenv("KEYHUB_CLIENTSECRET")

	if !config.Issuer.IsNull() {
		issuer = config.Issuer.ValueString()
	}

	if !config.ClientID.IsNull() {
		clientid = config.ClientID.ValueString()
	}

	if !config.ClientSecret.IsNull() {
		clientsecret = config.ClientSecret.ValueString()
	}

	if issuer == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("issuer"),
			"Missing Topicus KeyHub Issuer URI",
			"The provider cannot create the Topicus KeyHub API client as there is a missing or empty value for the Topicus KeyHub Issuer URI. "+
				"Set the issuer value in the configuration or use the KEYHUB_ISSUER environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if clientid == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("clientid"),
			"Missing Topicus KeyHub Client ID",
			"The provider cannot create the Topicus KeyHub API client as there is a missing or empty value for the Topicus KeyHub Client ID. "+
				"Set the clientid value in the configuration or use the KEYHUB_CLIENTID environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if clientsecret == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("clientsecret"),
			"Missing Topicus KeyHub Client ID",
			"The provider cannot create the Topicus KeyHub API client as there is a missing or empty value for the Topicus KeyHub Client ID. "+
				"Set the clientsecret value in the configuration or use the KEYHUB_CLIENTSECRET environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}
	ctx = tflog.SetField(ctx, "keyhub_issuer", issuer)
	ctx = tflog.SetField(ctx, "keyhub_clientid", clientid)
	ctx = tflog.SetField(ctx, "keyhub_clientsecret", clientsecret)
	ctx = tflog.MaskFieldValuesWithFieldKeys(ctx, "keyhub_clientsecret")

	tflog.Info(ctx, "Connecting to Topicus KeyHub")
	adapter, err := keyhub.NewKeyHubRequestAdapter(&http.Client{}, issuer, clientid, clientsecret)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create Topicus KeyHub API client",
			"An unexpected error occurred when creating the Topicus KeyHub API client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"Topicus KeyHub API client Error: "+err.Error(),
		)
		return
	}

	data := &KeyHubProviderData{
		Client: keyhub.NewKeyHubClient(adapter),
	}
	resp.DataSourceData = data
	resp.ResourceData = data

	tflog.Info(ctx, "Connected to Topicus KeyHub", map[string]any{"success": true})
}

func (p *KeyHubProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewClientapplicationResource,
		NewClientVaultrecordResource,
		NewGroupVaultrecordResource,
		NewGroupResource,
		NewGrouponsystemResource,
		NewServiceaccountResource,
	}
}

func (p *KeyHubProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewAccountDataSource,
		NewCertificateDataSource,
		NewClientDataSource,
		NewDirectoryDataSource,
		NewGroupDataSource,
		NewGroupclassificationDataSource,
		NewOrganizationalunitDataSource,
		NewServiceaccountDataSource,
		NewSystemDataSource,
		NewVaultrecordDataSource,
		NewWebhookDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &KeyHubProvider{
			version: version,
		}
	}
}
