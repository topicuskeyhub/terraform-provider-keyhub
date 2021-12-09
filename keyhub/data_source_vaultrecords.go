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

func dataSourceVaultRecords() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVaultRecordsRead,
		Schema: map[string]*schema.Schema{
			"groupuuid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vaultrecords": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: VaultRecordSchema(),
				},
			},
		},
	}
}

func dataSourceVaultRecordsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*keyhubclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	groupUUID := d.Get("groupuuid").(string)
	group, err := client.Groups.Get(groupUUID)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not GET group " + groupUUID + " for vault records",
			Detail:   err.Error(),
		})
	}

	vaultrecords, err := client.Vaults.GetRecords(group)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not GET vaultrecords of group " + groupUUID,
			Detail:   err.Error(),
		})
	}

	if err := d.Set("groupuuid", group.UUID); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not set value for groupuuid",
			Detail:   err.Error(),
		})
	}

	result := flattenVaultRecordsData(&vaultrecords)
	if err := d.Set("vaultrecords", result); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not set value for vaultrecords",
			Detail:   err.Error(),
		})
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenVaultRecordsData(vaultrecords *[]keyhubmodel.VaultRecord) []interface{} {
	if vaultrecords != nil {
		datas := make([]interface{}, len(*vaultrecords))

		for i, vaultrecord := range *vaultrecords {
			datas[i] = flattenVaultRecordData(&vaultrecord)
		}

		return datas
	}

	return make([]interface{}, 0)
}

func flattenVaultRecordData(vaultrecord *keyhubmodel.VaultRecord) map[string]interface{} {
	if vaultrecord != nil {
		data := make(map[string]interface{})

		data["id"] = strconv.FormatInt(vaultrecord.Self().ID, 10)
		data["uuid"] = vaultrecord.UUID
		data["name"] = vaultrecord.Name
		data["url"] = vaultrecord.URL
		data["username"] = vaultrecord.Username
		data["filename"] = vaultrecord.Filename

		return data
	}

	return make(map[string]interface{})
}
