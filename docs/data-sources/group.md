---
page_title: "keyhub_group Data Source - terraform-provider-keyhub"
subcategory: ""
description: |-
  The group data source allows you to retrieve information about one KeyHub group.
  
---

# keyhub_group (Data Source)

The group data source allows you to retrieve information about one KeyHub group.

## Example Usage

```terraform
data "keyhub_group" "example" {
  uuid = "example"
}
```

## Schema

### Required

- **uuid** (String) The UUID of the group you wish to retrieve

### Read-Only

- **id** (String) The value of the ID field of the group
- **name** (String) The value of the Name field of the group
