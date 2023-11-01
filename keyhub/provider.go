package keyhub

import (
	"context"
	"errors"
	"fmt"
	keyhubmodel "github.com/topicuskeyhub/go-keyhub/model"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	keyhubclient "github.com/topicuskeyhub/go-keyhub"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"issuer": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("KEYHUB_ISSUER", nil),
			},
			"clientid": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("KEYHUB_CLIENTID", nil),
			},
			"clientsecret": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("KEYHUB_CLIENTSECRET", nil),
			},
		},
		ConfigureContextFunc: providerConfigure,
		ResourcesMap: map[string]*schema.Resource{
			"keyhub_group":             resourceGroup(),
			"keyhub_vaultrecord":       resourceVaultRecord(),
			"keyhub_grouponsystem":     resourceGroupOnSystem(),
			"keyhub_clientapplication": resourceClientApplication(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"keyhub_group":             dataSourceGroup(),
			"keyhub_groups":            dataSourceGroups(),
			"keyhub_account":           dataSourceAccount(),
			"keyhub_accounts":          dataSourceAccounts(),
			"keyhub_vaultrecord":       dataSourceVaultRecord(),
			"keyhub_vaultrecords":      dataSourceVaultRecords(),
			"keyhub_provisionedsystem": dataSourceProvisionedSystem(),
		},
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	issuer := d.Get("issuer").(string)
	clientid := d.Get("clientid").(string)
	clientsecret := d.Get("clientsecret").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if issuer == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Issuer not set",
			Detail:   "Issuer is required for the KeyHub client to be able to connect to your KeyHub environment",
		})
	}
	if clientid == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "ClientId not set",
			Detail:   "ClientId is required for the KeyHub client to be able to authenticate with your KeyHub environment",
		})
	}
	if clientsecret == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "ClientSecret not set",
			Detail:   "ClientSecret is required for the KeyHub client to be able to authenticate with your KeyHub environment",
		})
	}

	client, err := keyhubclient.NewClient(http.DefaultClient, issuer, clientid, clientsecret)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create Keyhub client",
			Detail:   err.Error(),
		})
		return nil, diags
	}

	return client, diags
}

func apiErrorToLogFields(err error) map[string]interface{} {

	fields := map[string]interface{}{}

	var apiError keyhubmodel.KeyhubApiError
	if errors.As(err, &apiError) {
		fields["code"] = fmt.Sprintf("%d", apiError.Report.Code)
		fields["reason"] = apiError.Report.Reason
		fields["exception"] = apiError.Report.Exception
		fields["message"] = apiError.Report.Message
		fields["applicationError"] = apiError.Report.ApplicationError
		fields["stacktrace"] = strings.Join(apiError.Report.StackTrace, "\n")
	}

	return fields

}
