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

- **name** (String) The Name field of the group
- **member** (Block) At least one manager should be defined. *Note:* KeyHub currently only supports one member.

### Optional

- **description** (String) Group description

- **extended_access** (String) Defines extended access. Possible values: `NOT_ALLOWED` (default), `ONE_WEEK`, `TWO_WEEKS` 
- **vault_recovery** (String) Defines recovery strategy. Possible Values: `FULL` (default), `RECOVERY_KEY_ONLY`, `NONE`
- **audit_months** (List of Strings) List of Months the group must be audited. Possible Values: `JANUARY`,`FEBRUARY`,`MARCH`,`APRIL`,`MAY`,`JUNE`,`JULY`,`AUGUST`,`SEPTEMBER`,`OCTOBER`,`NOVEMBER`,`DECEMBER` 


- **rotating_password_required** (Bool) Required rotating password for members 
- **record_trail** (Bool) Require a reason before activating a group
- **private_group** (Bool) Set group to invite only
- **hide_audit_trail** (Bool) Don't show audit trail in KeyHub Dashboard
- **application_administration** (Bool) Group can be assign as managing group of an application
- **auditor** (Bool) No idea.


- **provisioning_auth_groupuuid** (String) UUID of the group to set as authorizing group for provisioning
- **membership_auth_groupuuid** (String) UUID of the group to set as authorizing group for membership
- **auditing_auth_groupuuid** (String) UUID of the group to set as authorizing group for audits

 
### Read-Only

- **id** (String) The value of the ID field of the group
- **uuid** (String) The UUID of the group 


## Import

KeyHub group can be imported using the uuid, e.g.

```
$ terraform import keyhub_group.example "00000000-0000-0000-0000-000000000000"
```
