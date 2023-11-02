---
page_title: "keyhub_group Resource - terraform-provider-keyhub"
subcategory: ""
description: |-
The group resource allows you to store/retrieve/update/delete information about one KeyHub group.
  
---

# keyhub_group (Resource)

The group resource allows you to store/retrieve/update/delete information about one KeyHub group.
*Note:* KeyHub currently only supports the creation of a keyhub group for a client.
Retreiving group details after creation requires adding the group details permission to the client in the KeyHub interface.
Update or Delete isn't possible 

## Permission Requirements
To set any of the `auditing_auth_groupuuid`, `membership_auth_groupuuid`, `nested_under_groupuuid`, `provisioning_auth_groupuuid` parameters 
the client used MUST have the `Groups - Configure a group for additional authorisation or nesting` permission on the referenced group.


## Example Usage

```terraform
resource "keyhub_group" "example" {
  name = "Example"
 
  member {
    uuid = "00000000-0000-0000-0000-000000000000"
    rights = "MANAGER"
  }
}
```

## Schema

### Required

- **member** (Block) At least one manager or nested_under_groupuuid should be defined
- **name** (String) The Name field of the group

### Optional

- **application_administration** (Bool) Group can be assign as managing group of an application
- **audit_months** (List) List of Months the group must be audited. Possible Values: `JANUARY`,`FEBRUARY`,`MARCH`,`APRIL`,`MAY`,`JUNE`,`JULY`,`AUGUST`,`SEPTEMBER`,`OCTOBER`,`NOVEMBER`,`DECEMBER`
- **auditing_auth_groupuuid** (String) The UUID of the group to set as authorizing group for audits
- **client** (Block) Grant clients permissions on the create group, (client used by terraform provider requires global `GROUPS_GRANT_PERMISSIONS_AFTER_CREATE` permission)
- **description** (String) The description of the group
- **extended_access** (String) Defines extended access. Possible values: `NOT_ALLOWED` (default), `ONE_WEEK`, `TWO_WEEKS`
- **hide_audit_trail** (Bool) Don't show audit trail in KeyHub Dashboard
- **membership_auth_groupuuid** (String) The UUID of the group to set as authorizing group for membership
- **nested_under_groupuuid** (String) The UUID of the group to nest the new group under
- **private_group** (Bool) Set group to invite only
- **provisioning_auth_groupuuid** (String) The UUID of the group to set as authorizing group for provisioning
- **record_trail** (Bool) Require a reason before activating a group
- **rotating_password_required** (Bool) Required rotating password for members
- **vault_recovery** (String) Defines recovery strategy. Possible Values: `FULL` (default), `RECOVERY_KEY_ONLY`, `NONE`

### Read-Only

- **id** (String) The value of the ID field of the group
- **uuid** (String) The UUID of the group

### Blocks

The *member* block supports the following:
- **name** (String) The name of the member
- **rights** (String) The rights of the member. Possible values: `MANAGER` (default), `NORMAL`
- **uuid** (String, Required) The uuid of the keyhub account to add as member

The *client* block supports the following:
- **permissions** (List, Required) List of permissions to grant the client application. Possible values: `GROUP_FULL_VAULT_ACCESS`, `GROUP_READ_CONTENTS`, `GROUP_SET_AUTHORIZATION`
- **uuid** (String, Required) The UUID of the client application


## Import

KeyHub group can be imported using the uuid, e.g.

```
$ terraform import keyhub_group.example "00000000-0000-0000-0000-000000000000"
```
