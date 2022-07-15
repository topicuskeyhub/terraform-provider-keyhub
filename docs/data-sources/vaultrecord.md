---
page_title: "keyhub_vaultrecord Data Source - terraform-provider-keyhub"
subcategory: ""
description: |-
  The vaultrecord data source allows you to retrieve information about one KeyHub vaultrecord.
  
---

# keyhub_vaultrecord (Data Source)

The vaultrecord data source allows you to retrieve information about one KeyHub vaultrecord.

## Example Usage

```terraform
data "keyhub_vaultrecord" "example" {
  groupuuid = "example"
  uuid = "example"
}
```

## Schema

### Required

- **uuid** (String) The UUID of the vaultrecord you wish to retrieve

### Optional

- **groupuuid** (String) The group UUID of the vaultrecord you wish to retrieve

### Read-Only

- **id** (String) The value of the ID field of the vaultrecord
- **name** (String) The value of the Name field of the vaultrecord
- **url** (String) The value of the URL field of the vaultrecord
- **username** (String) The value of the Username field of the vaultrecord
- **filename** (String)  The value of the Filename field of the vaultrecord

- **comment** (String, Sensitive) The value of the Comment field of the vaultrecord. This value is sensitive as it might contain secret information.
- **password** (String, Sensitive)  The value of the Password field of the vaultrecord. This value is sensitive as it might contain secret information.
- **totp** (String, Sensitive)  The value of the Totp field of the vaultrecord. This value is sensitive as it might contain secret information.

## Import

KeyHub group can be imported using the uuid, e.g.

```
$ terraform import keyhub_vaultrecord.example "00000000-0000-0000-0000-000000000000"
```
