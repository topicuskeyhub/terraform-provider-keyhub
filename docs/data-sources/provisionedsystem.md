---
page_title: "keyhub_provisionedsystem Data Source - terraform-provider-keyhub"
subcategory: ""
description: |-
The provisionedsystem data source allows you to retrieve information about one KeyHub provisioned system.
  
---

# keyhub_group (Data Source)

The provisioned system data source allows you to retrieve information about one KeyHub provisioned system.

## Example Usage

```terraform
data "keyhub_provisionedsystem" "example" {
  uuid = "example"
}
```

## Schema

### Required

- **uuid** (String) The UUID of the of the provisioned system

### Read-Only

- **accountcount** (Int) The amount of accounts on the provisioned system
- **externaluuid** (String) The external uuid of the provisioned system
- **id** (String) The value of the ID field of the provisioned system
- **name** (String) The name of the provisioned system
- **technicaladministrator** (Map) The UUID and Name of the group that is set as the technical administrator
- **type** (String) The type of the provisioned system
- **usernameprefix** (String) The username prefix of the provisioned system