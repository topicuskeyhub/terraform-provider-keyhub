package keyhub

import (
	"context"
	"strconv"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	keyhubclient "github.com/topicuskeyhub/go-keyhub"
	keyhubmodel "github.com/topicuskeyhub/go-keyhub/model"
)

func resourceVaultRecord() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceVaultRecordCreate,
		ReadContext:   dataSourceVaultRecordRead, //we use the DataSource READ as it supports both data types
		UpdateContext: resourceVaultRecordUpdate,
		DeleteContext: resourceVaultRecordDelete,
		Schema:        VaultRecordResourceSchema(),
	}
}

func VaultRecordResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
			Required: false,
		},
		"groupuuid": {
			Type:     schema.TypeString,
			Required: true,
		},
		"uuid": {
			Type:     schema.TypeString,
			Computed: true,
			Required: false,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"url": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"username": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"filename": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"password": {
			Type:      schema.TypeString,
			Sensitive: true,
			Optional:  true,
		},
		"totp": {
			Type:      schema.TypeString,
			Sensitive: true,
			Optional:  true,
		},
		// "file": {
		// 	Type:      schema.TypeString,
		// 	Sensitive: true,
		// Optional: true,
		// },
		"comment": {
			Type:      schema.TypeString,
			Sensitive: true,
			Optional:  true,
		},
	}
}

func resourceVaultRecordCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*keyhubclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	name := d.Get("name").(string)

	groupUUIDString := d.Get("groupuuid").(string)
	groupUUID, err := uuid.Parse(groupUUIDString)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Field 'groupuuid' is not a valid UUID",
			Detail:   err.Error(),
		})
	}
	group, err := client.Groups.GetByUUID(groupUUID)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not GET group " + groupUUIDString + " for new vault record " + name,
			Detail:   err.Error(),
		})
		return diags
	}

	//query vaultrecords by name to prevent duplicates
	existingVaultRecord, err := client.Vaults.List(group, &keyhubmodel.VaultRecordQueryParams{Name: name}, &keyhubmodel.VaultRecordAdditionalQueryParams{Secret: true})
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not CREATE vaultRecord in group " + groupUUIDString,
			Detail:   err.Error(),
		})
		return diags
	}
	if len(existingVaultRecord) > 0 {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not CREATE vaultRecord",
			Detail:   "Record with same name exists in group in group " + groupUUIDString,
		})
		return diags
	}

	secrets := &keyhubmodel.VaultRecordSecretAdditionalObject{}
	vaultRecord := keyhubmodel.NewVaultRecord(name, secrets)

	//copy schema data to model
	//use generic copy method. also used in UPDATE.
	vaultRecordSchemaToModel(d, vaultRecord, &diags)
	if diags.HasError() {
		return diags
	}

	newVaultRecord, err := client.Vaults.Create(group, vaultRecord)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not CREATE vaultRecord in group " + groupUUIDString,
			Detail:   err.Error(),
		})
		return diags
	}

	// d.SetId(group.UUID + "/" + newVaultRecord.UUID)
	d.SetId(strconv.FormatInt(newVaultRecord.Self().ID, 10))

	dataSourceVaultRecordRead(ctx, d, m)

	return diags
}

func resourceVaultRecordUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*keyhubclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	groupUUIDString := d.Get("groupuuid").(string)
	groupUUID, err := uuid.Parse(groupUUIDString)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Field 'groupuuid' is not a valid UUID",
			Detail:   err.Error(),
		})
	}

	ID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not parse ID " + d.Id(),
			Detail:   err.Error(),
		})
		return diags
	}

	group, err := client.Groups.GetByUUID(groupUUID)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not GET group " + groupUUIDString + " for vault record " + d.Id(),
			Detail:   err.Error(),
		})
		return diags
	}

	vaultRecord, err := client.Vaults.GetByID(group, ID, &keyhubmodel.VaultRecordAdditionalQueryParams{Secret: true})
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not GET vault record " + d.Id(),
			Detail:   err.Error(),
		})
		return diags
	}

	//copy schema data to model
	//use generic copy method. also used in CREATE.
	vaultRecordSchemaToModel(d, vaultRecord, &diags)
	if diags.HasError() {
		return diags
	}

	_, err = client.Vaults.Update(group, vaultRecord)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not UPDATE vault record " + vaultRecord.UUID,
			Detail:   err.Error(),
		})
		return diags
	}

	dataSourceVaultRecordRead(ctx, d, m)

	return diags
}

func resourceVaultRecordDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*keyhubclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	groupUUIDString := d.Get("groupuuid").(string)
	groupUUID, err := uuid.Parse(groupUUIDString)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Field 'groupuuid' is not a valid UUID",
			Detail:   err.Error(),
		})
	}

	ID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not parse ID " + d.Id(),
			Detail:   err.Error(),
		})
		return diags
	}

	group, err := client.Groups.GetByUUID(groupUUID)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not GET group " + groupUUIDString + " for vault record " + d.Id(),
			Detail:   err.Error(),
		})
		return diags
	}

	err = client.Vaults.DeleteByID(group, ID)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Could not DELETE vaultrecord " + d.Id(),
			Detail:   err.Error(),
		})
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}

func vaultRecordSchemaToModel(d *schema.ResourceData, vaultRecord *keyhubmodel.VaultRecord, diags *diag.Diagnostics) {

	value, valueExists := d.GetOk("name")
	if valueExists {
		vaultRecord.Name = value.(string)
	}
	value, valueExists = d.GetOk("url")
	if valueExists {
		vaultRecord.URL = value.(string)
	}
	value, valueExists = d.GetOk("username")
	if valueExists {
		vaultRecord.Username = value.(string)
	}
	value, valueExists = d.GetOk("color")
	if valueExists {
		vaultRecord.Color = value.(string)
	}
	value, valueExists = d.GetOk("filename")
	if valueExists {
		vaultRecord.Filename = value.(string)
	}

	value, valueExists = d.GetOk("password")
	if valueExists {
		val := value.(string)
		vaultRecord.AdditionalObjects.Secret.Password = &val
	}
	value, valueExists = d.GetOk("totp")
	if valueExists {
		val := value.(string)
		vaultRecord.AdditionalObjects.Secret.Totp = &val
	}
	// value, valueExists = d.GetOk("file")
	// if valueExists {
	// 	val := value.([]byte)
	// 	vaultRecord.AdditionalObjects.Secret.File = &val
	// }
	value, valueExists = d.GetOk("comment")
	if valueExists {
		val := value.(string)
		vaultRecord.AdditionalObjects.Secret.Comment = &val
	}
}
