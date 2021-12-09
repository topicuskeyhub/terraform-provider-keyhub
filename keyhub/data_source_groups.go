package keyhub

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	keyhubclient "github.com/topicuskeyhub/go-keyhub"
	keyhubmodel "github.com/topicuskeyhub/go-keyhub/model"
)

func dataSourceGroups() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGroupsRead,
		Schema: map[string]*schema.Schema{
			"groups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: GroupSchema(),
				},
			},
		},
	}
}

func dataSourceGroupsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*keyhubclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	groups, err := client.Groups.List()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not GET groups",
			Detail:   err.Error(),
		})
	}

	result := flattenGroupsData(&groups)
	if err := d.Set("groups", result); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not set value for groups",
			Detail:   err.Error(),
		})
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenGroupsData(groups *[]keyhubmodel.Group) []interface{} {
	if groups != nil {
		datas := make([]interface{}, len(*groups))

		for i, group := range *groups {
			datas[i] = flattenGroupData(&group)
		}

		return datas
	}

	return make([]interface{}, 0)
}

func flattenGroupData(group *keyhubmodel.Group) map[string]interface{} {
	if group != nil {
		data := make(map[string]interface{})

		data["id"] = strconv.FormatInt(group.Self().ID, 10)
		data["uuid"] = group.UUID
		data["name"] = group.Name

		return data
	}

	return make(map[string]interface{})
}
