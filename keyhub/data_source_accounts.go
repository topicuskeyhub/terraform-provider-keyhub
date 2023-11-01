package keyhub

import (
	"context"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	keyhubclient "github.com/topicuskeyhub/go-keyhub"
	keyhubmodel "github.com/topicuskeyhub/go-keyhub/model"
)

func dataSourceAccounts() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAccountsRead,
		Schema: map[string]*schema.Schema{
			"accounts": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: AccountSchema(),
				},
			},
		},
	}
}

func dataSourceAccountsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*keyhubclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	accounts, err := client.Accounts.List()
	if err != nil {
		tflog.Debug(ctx, err.Error(), apiErrorToLogFields(err))
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not GET accounts",
			Detail:   err.Error(),
		})
		return diags
	}

	result := flattenAccountsData(&accounts)
	if err := d.Set("accounts", result); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not set value for accounts",
			Detail:   err.Error(),
		})
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenAccountsData(accounts *[]keyhubmodel.Account) []interface{} {
	if accounts != nil {
		datas := make([]interface{}, len(*accounts))

		for i, account := range *accounts {
			datas[i] = flattenAccountData(&account)
		}

		return datas
	}

	return make([]interface{}, 0)
}

func flattenAccountData(account *keyhubmodel.Account) map[string]interface{} {
	if account != nil {
		data := make(map[string]interface{})

		data["id"] = strconv.FormatInt(account.Self().ID, 10)
		data["uuid"] = account.UUID
		data["username"] = account.Username
		data["name"] = account.DisplayName

		return data
	}

	return make(map[string]interface{})
}
