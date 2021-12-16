---
page_title: "keyhub_vaultrecord Resource - terraform-provider-keyhub"
subcategory: ""
description: |-
  The vaultrecord resource allows you to store/retrieve/update/delete  information about one KeyHub vaultrecord.
  
---

# keyhub_vaultrecord (Reource)

The vaultrecord resource allows you to store/retrieve/update/delete information about one KeyHub vaultrecord.

## Example Usage

```terraform
data "keyhub_vaultrecord" "example" {
  groupuuid = "example"
  name = "Example"
  password = "Example"
}
```

## Schema

### Required

- **groupuuid** (String) The group UUID of the vaultrecord you wish to store/retrieve/update/delete
- **name** (String) The Name field of the vaultrecord

### Optional

- **url** (String)
- **username** (String)
- **filename** (String)

At least one of the following is required:

- **comment** (String, Sensitive) The value of the Comment field of the vaultrecord. This value is sensitive as it might contain secret information.
- **password** (String, Sensitive)  The value of the Password field of the vaultrecord. This value is sensitive as it might contain secret information.
- **totp** (String, Sensitive)  The value of the Totp field of the vaultrecord. This value is sensitive as it might contain secret information.

### Read-Only

- **id** (String) The value of the ID field of the vaultrecord
- **uuid** (String) The UUID of the vaultrecord you wish to retreive


