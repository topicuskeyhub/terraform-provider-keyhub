---
page_title: "keyhub_grouponsystem Resource - terraform-provider-keyhub"
subcategory: ""
description: |-
The group resource allows you to store/retrieve/<s>update</s>/<s>delete</s> information about one KeyHub group on system.

---

# keyhub_grouponsystem (Resource)

The group resource allows you to create and retrieve information about one KeyHub group on system.
*Note:* KeyHub currently only supports the creation of a group on system for a client, Update or Delete isn't possible

## Example Usage

```terraform

resource "keyhub_grouponsystem" "example" {

  type = "POSIX_GROUP"                          
  owner = "00000000-0000-0000-0000-000000000000"   # UUID of keyhub group
  system = "00000000-0000-0000-0000-000000000000"  # UUID of linked system
  name_in_system = "umbrella"

  provgroup {
    group = "00000000-0000-0000-0000-000000000000"
    activation_required = true
  }

}
```

## Schema

### Required

- **name_in_system** (String) The name in the system, value normally to CN of the DN (of the provisioned system). For example: `cn=umbrella,ou=group,dc=example,dc=com`
- **owner** (String) The UUID of the group that will become owner of the grouponsystem
- **system** (String) The UUID of the provisioned system to create the group on

### Optional

- **display_name** (String) Display name of the group on provisioned system. (Only on systems that support a display name)
- **provgroup** (Block) Define the provisioning group for the grouponsystem, can be set multiple times. If omitted the owner group will be the provisioning group
- **type** (String) Type of the resulting group in the provisioned system, for example: POSIX_GROUP for ldap.

### Read-Only

- **id** (String) Type of the resulting group in the provisioned system, for example: POSIX_GROUP for ldap.
- **short_name_in_system** (String) common name part of the resulting DN. For example: `cn=umbrella`

### Blocks

The *provgroup* block supports the following:
- **group** (String, Required) The UUID of the group that will become a provisioning group for the grouponsystem
- **activation_required** (Bool) Set to false to have the group on system statically provisioned




## Import

KeyHub group can be imported using the system and grouponsystem id's , e.g.

```
$ terraform import keyhub_grouponsystem.example "<system_id>:<grouponsystem_id>"
```

But both id's aren't available/accessible through the Keyhub GUI so advanced knowlegde of the keyhub api is required.