---
page_title: "keyhub_accounts Data Source - terraform-provider-keyhub"
subcategory: ""
description: |-
  The accounts data source allows you to retrieve information about all KeyHub accounts.
  
---

# keyhub_accounts (Data Source)

The accounts data source allows you to retrieve information about all KeyHub accounts.

## Example Usage

```terraform
data "keyhub_accounts" "all" {
}
```

## Schema

### Optional

- **id** (String) The ID of this resource. This is never filled.

### Read-Only

- **accounts** (List of Object) (see [below for nested schema](#nestedatt--accounts))

<a id="nestedatt--accounts"></a>
### Nested Schema for `accounts`

Read-Only:

- **id** (String) The value of the ID field of the account
- **uuid** (String) The value of the UUID field of the account
- **name** (String) The value of the Name field of the account
- **username** (String) The value of the Username field of the account
