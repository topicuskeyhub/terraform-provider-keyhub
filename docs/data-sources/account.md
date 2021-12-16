---
page_title: "keyhub_account Data Source - terraform-provider-keyhub"
subcategory: ""
description: |-
  The account data source allows you to retrieve information about one KeyHub account.
  
---

# keyhub_account (Data Source)

The account data source allows you to retrieve information about one KeyHub account.

## Example Usage

```terraform
data "keyhub_account" "example" {
  uuid = "example"
}
```

## Schema

### Required

- **uuid** (String) The UUID of the account you wish to retreive

### Read-Only

- **id** (String) The value of the ID field of the account
- **name** (String) The value of the Name field of the account
- **username** (String) The value of the Username field of the account
