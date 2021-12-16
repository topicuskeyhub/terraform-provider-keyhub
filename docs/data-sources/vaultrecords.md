---
page_title: "keyhub_vaultrecords Data Source - terraform-provider-keyhub"
subcategory: ""
description: |-
  The vaultrecords data source allows you to retrieve information about all KeyHub vaultrecords.
  
---

# keyhub_vaultrecords (Data Source)

The vaultrecords data source allows you to retrieve information about all KeyHub vaultrecords.

## Example Usage

```terraform
data "keyhub_vaultrecords" "all" {
  groupuuid = "example"
}
```

## Schema

### Required

- **groupuuid** (String) The group UUID of the vaultrecords you wish to retreive.

### Optional

- **id** (String) The ID of this resource. This is never filled.

### Read-Only

- **vaultrecords** (List of Object) (see [below for nested schema](#nestedatt--vaultrecords))

<a id="nestedatt--vaultrecords"></a>
### Nested Schema for `vaultrecords`

Read-Only

- **id** (String) The value of the ID field of the vaultrecord
- **uuid** (String) The value of the UUID field of the vaultrecord
- **name** (String) The value of the Name field of the vaultrecord
- **url** (String) The value of the URL field of the vaultrecord
- **username** (String) The value of the Username field of the vaultrecord
- **filename** (String)  The value of the Filename field of the vaultrecord

- **comment** (String, Sensitive) The value of the Comment field of the vaultrecord. This value is sensitive as it might contain secret information.
- **password** (String, Sensitive)  The value of the Password field of the vaultrecord. This value is sensitive as it might contain secret information.
- **totp** (String, Sensitive)  The value of the Totp field of the vaultrecord. This value is sensitive as it might contain secret information.
