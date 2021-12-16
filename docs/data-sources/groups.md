---
page_title: "keyhub_groups Data Source - terraform-provider-keyhub"
subcategory: ""
description: |-
  The groups data source allows you to retrieve information about all KeyHub groups.
  
---

# keyhub_groups (Data Source)

The groups data source allows you to retrieve information about all KeyHub groups.

## Example Usage

```terraform
data "keyhub_groups" "all" {
}
```

## Schema

### Optional

- **id** (String) The ID of this resource. This is never filled.

### Read-Only

- **groups** (List of Object) (see [below for nested schema](#nestedatt--groups))

<a id="nestedatt--groups"></a>
### Nested Schema for `groups`

Read-Only:

- **id** (String) The value of the ID field of the group
- **uuid** (String) The value of the UUID field of the group
- **name** (String) The value of the Name field of the group
