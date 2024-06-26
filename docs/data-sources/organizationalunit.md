---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "keyhub_organizationalunit Data Source - terraform-provider-keyhub"
subcategory: ""
description: |-
  
---

# keyhub_organizationalunit (Data Source)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `uuid` (String)

### Optional

- `additional` (List of String)

### Read-Only

- `audit` (Attributes) (see [below for nested schema](#nestedatt--audit))
- `auditor_group` (Attributes) (see [below for nested schema](#nestedatt--auditor_group))
- `create_group_approve_group` (Attributes) (see [below for nested schema](#nestedatt--create_group_approve_group))
- `create_group_placeholder` (String)
- `depth` (Number)
- `description` (String)
- `enable_tech_admin_approve_group` (Attributes) (see [below for nested schema](#nestedatt--enable_tech_admin_approve_group))
- `links` (Attributes List) (see [below for nested schema](#nestedatt--links))
- `name` (String)
- `owner` (Attributes) (see [below for nested schema](#nestedatt--owner))
- `parent` (Attributes) (see [below for nested schema](#nestedatt--parent))
- `permissions` (Attributes List) (see [below for nested schema](#nestedatt--permissions))
- `recovery_fallback_group` (Attributes) (see [below for nested schema](#nestedatt--recovery_fallback_group))
- `remove_group_approve_group` (Attributes) (see [below for nested schema](#nestedatt--remove_group_approve_group))
- `settings` (Attributes) (see [below for nested schema](#nestedatt--settings))

<a id="nestedatt--audit"></a>
### Nested Schema for `audit`

Read-Only:

- `created_at` (String)
- `created_by` (String)
- `last_modified_at` (String)
- `last_modified_by` (String)


<a id="nestedatt--auditor_group"></a>
### Nested Schema for `auditor_group`

Read-Only:

- `admin` (Boolean)
- `links` (Attributes List) (see [below for nested schema](#nestedatt--auditor_group--links))
- `name` (String)
- `organizational_unit` (Attributes) (see [below for nested schema](#nestedatt--auditor_group--organizational_unit))
- `permissions` (Attributes List) (see [below for nested schema](#nestedatt--auditor_group--permissions))
- `uuid` (String)

<a id="nestedatt--auditor_group--links"></a>
### Nested Schema for `auditor_group.links`

Read-Only:

- `href` (String)
- `id` (Number)
- `rel` (String)
- `type_escaped` (String)


<a id="nestedatt--auditor_group--organizational_unit"></a>
### Nested Schema for `auditor_group.organizational_unit`

Read-Only:

- `links` (Attributes List) (see [below for nested schema](#nestedatt--auditor_group--organizational_unit--links))
- `name` (String)
- `permissions` (Attributes List) (see [below for nested schema](#nestedatt--auditor_group--organizational_unit--permissions))
- `uuid` (String)

<a id="nestedatt--auditor_group--organizational_unit--links"></a>
### Nested Schema for `auditor_group.organizational_unit.links`

Read-Only:

- `href` (String)
- `id` (Number)
- `rel` (String)
- `type_escaped` (String)


<a id="nestedatt--auditor_group--organizational_unit--permissions"></a>
### Nested Schema for `auditor_group.organizational_unit.permissions`

Read-Only:

- `full` (String)
- `instances` (List of String)
- `operations` (Set of String)
- `type_escaped` (String)



<a id="nestedatt--auditor_group--permissions"></a>
### Nested Schema for `auditor_group.permissions`

Read-Only:

- `full` (String)
- `instances` (List of String)
- `operations` (Set of String)
- `type_escaped` (String)



<a id="nestedatt--create_group_approve_group"></a>
### Nested Schema for `create_group_approve_group`

Read-Only:

- `admin` (Boolean)
- `links` (Attributes List) (see [below for nested schema](#nestedatt--create_group_approve_group--links))
- `name` (String)
- `organizational_unit` (Attributes) (see [below for nested schema](#nestedatt--create_group_approve_group--organizational_unit))
- `permissions` (Attributes List) (see [below for nested schema](#nestedatt--create_group_approve_group--permissions))
- `uuid` (String)

<a id="nestedatt--create_group_approve_group--links"></a>
### Nested Schema for `create_group_approve_group.links`

Read-Only:

- `href` (String)
- `id` (Number)
- `rel` (String)
- `type_escaped` (String)


<a id="nestedatt--create_group_approve_group--organizational_unit"></a>
### Nested Schema for `create_group_approve_group.organizational_unit`

Read-Only:

- `links` (Attributes List) (see [below for nested schema](#nestedatt--create_group_approve_group--organizational_unit--links))
- `name` (String)
- `permissions` (Attributes List) (see [below for nested schema](#nestedatt--create_group_approve_group--organizational_unit--permissions))
- `uuid` (String)

<a id="nestedatt--create_group_approve_group--organizational_unit--links"></a>
### Nested Schema for `create_group_approve_group.organizational_unit.links`

Read-Only:

- `href` (String)
- `id` (Number)
- `rel` (String)
- `type_escaped` (String)


<a id="nestedatt--create_group_approve_group--organizational_unit--permissions"></a>
### Nested Schema for `create_group_approve_group.organizational_unit.permissions`

Read-Only:

- `full` (String)
- `instances` (List of String)
- `operations` (Set of String)
- `type_escaped` (String)



<a id="nestedatt--create_group_approve_group--permissions"></a>
### Nested Schema for `create_group_approve_group.permissions`

Read-Only:

- `full` (String)
- `instances` (List of String)
- `operations` (Set of String)
- `type_escaped` (String)



<a id="nestedatt--enable_tech_admin_approve_group"></a>
### Nested Schema for `enable_tech_admin_approve_group`

Read-Only:

- `admin` (Boolean)
- `links` (Attributes List) (see [below for nested schema](#nestedatt--enable_tech_admin_approve_group--links))
- `name` (String)
- `organizational_unit` (Attributes) (see [below for nested schema](#nestedatt--enable_tech_admin_approve_group--organizational_unit))
- `permissions` (Attributes List) (see [below for nested schema](#nestedatt--enable_tech_admin_approve_group--permissions))
- `uuid` (String)

<a id="nestedatt--enable_tech_admin_approve_group--links"></a>
### Nested Schema for `enable_tech_admin_approve_group.links`

Read-Only:

- `href` (String)
- `id` (Number)
- `rel` (String)
- `type_escaped` (String)


<a id="nestedatt--enable_tech_admin_approve_group--organizational_unit"></a>
### Nested Schema for `enable_tech_admin_approve_group.organizational_unit`

Read-Only:

- `links` (Attributes List) (see [below for nested schema](#nestedatt--enable_tech_admin_approve_group--organizational_unit--links))
- `name` (String)
- `permissions` (Attributes List) (see [below for nested schema](#nestedatt--enable_tech_admin_approve_group--organizational_unit--permissions))
- `uuid` (String)

<a id="nestedatt--enable_tech_admin_approve_group--organizational_unit--links"></a>
### Nested Schema for `enable_tech_admin_approve_group.organizational_unit.links`

Read-Only:

- `href` (String)
- `id` (Number)
- `rel` (String)
- `type_escaped` (String)


<a id="nestedatt--enable_tech_admin_approve_group--organizational_unit--permissions"></a>
### Nested Schema for `enable_tech_admin_approve_group.organizational_unit.permissions`

Read-Only:

- `full` (String)
- `instances` (List of String)
- `operations` (Set of String)
- `type_escaped` (String)



<a id="nestedatt--enable_tech_admin_approve_group--permissions"></a>
### Nested Schema for `enable_tech_admin_approve_group.permissions`

Read-Only:

- `full` (String)
- `instances` (List of String)
- `operations` (Set of String)
- `type_escaped` (String)



<a id="nestedatt--links"></a>
### Nested Schema for `links`

Read-Only:

- `href` (String)
- `id` (Number)
- `rel` (String)
- `type_escaped` (String)


<a id="nestedatt--owner"></a>
### Nested Schema for `owner`

Read-Only:

- `admin` (Boolean)
- `links` (Attributes List) (see [below for nested schema](#nestedatt--owner--links))
- `name` (String)
- `organizational_unit` (Attributes) (see [below for nested schema](#nestedatt--owner--organizational_unit))
- `permissions` (Attributes List) (see [below for nested schema](#nestedatt--owner--permissions))
- `uuid` (String)

<a id="nestedatt--owner--links"></a>
### Nested Schema for `owner.links`

Read-Only:

- `href` (String)
- `id` (Number)
- `rel` (String)
- `type_escaped` (String)


<a id="nestedatt--owner--organizational_unit"></a>
### Nested Schema for `owner.organizational_unit`

Read-Only:

- `links` (Attributes List) (see [below for nested schema](#nestedatt--owner--organizational_unit--links))
- `name` (String)
- `permissions` (Attributes List) (see [below for nested schema](#nestedatt--owner--organizational_unit--permissions))
- `uuid` (String)

<a id="nestedatt--owner--organizational_unit--links"></a>
### Nested Schema for `owner.organizational_unit.links`

Read-Only:

- `href` (String)
- `id` (Number)
- `rel` (String)
- `type_escaped` (String)


<a id="nestedatt--owner--organizational_unit--permissions"></a>
### Nested Schema for `owner.organizational_unit.permissions`

Read-Only:

- `full` (String)
- `instances` (List of String)
- `operations` (Set of String)
- `type_escaped` (String)



<a id="nestedatt--owner--permissions"></a>
### Nested Schema for `owner.permissions`

Read-Only:

- `full` (String)
- `instances` (List of String)
- `operations` (Set of String)
- `type_escaped` (String)



<a id="nestedatt--parent"></a>
### Nested Schema for `parent`

Read-Only:

- `links` (Attributes List) (see [below for nested schema](#nestedatt--parent--links))
- `name` (String)
- `permissions` (Attributes List) (see [below for nested schema](#nestedatt--parent--permissions))
- `uuid` (String)

<a id="nestedatt--parent--links"></a>
### Nested Schema for `parent.links`

Read-Only:

- `href` (String)
- `id` (Number)
- `rel` (String)
- `type_escaped` (String)


<a id="nestedatt--parent--permissions"></a>
### Nested Schema for `parent.permissions`

Read-Only:

- `full` (String)
- `instances` (List of String)
- `operations` (Set of String)
- `type_escaped` (String)



<a id="nestedatt--permissions"></a>
### Nested Schema for `permissions`

Read-Only:

- `full` (String)
- `instances` (List of String)
- `operations` (Set of String)
- `type_escaped` (String)


<a id="nestedatt--recovery_fallback_group"></a>
### Nested Schema for `recovery_fallback_group`

Read-Only:

- `admin` (Boolean)
- `links` (Attributes List) (see [below for nested schema](#nestedatt--recovery_fallback_group--links))
- `name` (String)
- `organizational_unit` (Attributes) (see [below for nested schema](#nestedatt--recovery_fallback_group--organizational_unit))
- `permissions` (Attributes List) (see [below for nested schema](#nestedatt--recovery_fallback_group--permissions))
- `uuid` (String)

<a id="nestedatt--recovery_fallback_group--links"></a>
### Nested Schema for `recovery_fallback_group.links`

Read-Only:

- `href` (String)
- `id` (Number)
- `rel` (String)
- `type_escaped` (String)


<a id="nestedatt--recovery_fallback_group--organizational_unit"></a>
### Nested Schema for `recovery_fallback_group.organizational_unit`

Read-Only:

- `links` (Attributes List) (see [below for nested schema](#nestedatt--recovery_fallback_group--organizational_unit--links))
- `name` (String)
- `permissions` (Attributes List) (see [below for nested schema](#nestedatt--recovery_fallback_group--organizational_unit--permissions))
- `uuid` (String)

<a id="nestedatt--recovery_fallback_group--organizational_unit--links"></a>
### Nested Schema for `recovery_fallback_group.organizational_unit.links`

Read-Only:

- `href` (String)
- `id` (Number)
- `rel` (String)
- `type_escaped` (String)


<a id="nestedatt--recovery_fallback_group--organizational_unit--permissions"></a>
### Nested Schema for `recovery_fallback_group.organizational_unit.permissions`

Read-Only:

- `full` (String)
- `instances` (List of String)
- `operations` (Set of String)
- `type_escaped` (String)



<a id="nestedatt--recovery_fallback_group--permissions"></a>
### Nested Schema for `recovery_fallback_group.permissions`

Read-Only:

- `full` (String)
- `instances` (List of String)
- `operations` (Set of String)
- `type_escaped` (String)



<a id="nestedatt--remove_group_approve_group"></a>
### Nested Schema for `remove_group_approve_group`

Read-Only:

- `admin` (Boolean)
- `links` (Attributes List) (see [below for nested schema](#nestedatt--remove_group_approve_group--links))
- `name` (String)
- `organizational_unit` (Attributes) (see [below for nested schema](#nestedatt--remove_group_approve_group--organizational_unit))
- `permissions` (Attributes List) (see [below for nested schema](#nestedatt--remove_group_approve_group--permissions))
- `uuid` (String)

<a id="nestedatt--remove_group_approve_group--links"></a>
### Nested Schema for `remove_group_approve_group.links`

Read-Only:

- `href` (String)
- `id` (Number)
- `rel` (String)
- `type_escaped` (String)


<a id="nestedatt--remove_group_approve_group--organizational_unit"></a>
### Nested Schema for `remove_group_approve_group.organizational_unit`

Read-Only:

- `links` (Attributes List) (see [below for nested schema](#nestedatt--remove_group_approve_group--organizational_unit--links))
- `name` (String)
- `permissions` (Attributes List) (see [below for nested schema](#nestedatt--remove_group_approve_group--organizational_unit--permissions))
- `uuid` (String)

<a id="nestedatt--remove_group_approve_group--organizational_unit--links"></a>
### Nested Schema for `remove_group_approve_group.organizational_unit.links`

Read-Only:

- `href` (String)
- `id` (Number)
- `rel` (String)
- `type_escaped` (String)


<a id="nestedatt--remove_group_approve_group--organizational_unit--permissions"></a>
### Nested Schema for `remove_group_approve_group.organizational_unit.permissions`

Read-Only:

- `full` (String)
- `instances` (List of String)
- `operations` (Set of String)
- `type_escaped` (String)



<a id="nestedatt--remove_group_approve_group--permissions"></a>
### Nested Schema for `remove_group_approve_group.permissions`

Read-Only:

- `full` (String)
- `instances` (List of String)
- `operations` (Set of String)
- `type_escaped` (String)



<a id="nestedatt--settings"></a>
### Nested Schema for `settings`

Read-Only:

- `create_group_approve_group` (Attributes) (see [below for nested schema](#nestedatt--settings--create_group_approve_group))
- `create_group_placeholder` (String)
- `enable_tech_admin_approve_group` (Attributes) (see [below for nested schema](#nestedatt--settings--enable_tech_admin_approve_group))
- `recovery_fallback_group` (Attributes) (see [below for nested schema](#nestedatt--settings--recovery_fallback_group))
- `remove_group_approve_group` (Attributes) (see [below for nested schema](#nestedatt--settings--remove_group_approve_group))

<a id="nestedatt--settings--create_group_approve_group"></a>
### Nested Schema for `settings.create_group_approve_group`

Read-Only:

- `admin` (Boolean)
- `links` (Attributes List) (see [below for nested schema](#nestedatt--settings--create_group_approve_group--links))
- `name` (String)
- `organizational_unit` (Attributes) (see [below for nested schema](#nestedatt--settings--create_group_approve_group--organizational_unit))
- `permissions` (Attributes List) (see [below for nested schema](#nestedatt--settings--create_group_approve_group--permissions))
- `uuid` (String)

<a id="nestedatt--settings--create_group_approve_group--links"></a>
### Nested Schema for `settings.create_group_approve_group.links`

Read-Only:

- `href` (String)
- `id` (Number)
- `rel` (String)
- `type_escaped` (String)


<a id="nestedatt--settings--create_group_approve_group--organizational_unit"></a>
### Nested Schema for `settings.create_group_approve_group.organizational_unit`

Read-Only:

- `links` (Attributes List) (see [below for nested schema](#nestedatt--settings--create_group_approve_group--organizational_unit--links))
- `name` (String)
- `permissions` (Attributes List) (see [below for nested schema](#nestedatt--settings--create_group_approve_group--organizational_unit--permissions))
- `uuid` (String)

<a id="nestedatt--settings--create_group_approve_group--organizational_unit--links"></a>
### Nested Schema for `settings.create_group_approve_group.organizational_unit.links`

Read-Only:

- `href` (String)
- `id` (Number)
- `rel` (String)
- `type_escaped` (String)


<a id="nestedatt--settings--create_group_approve_group--organizational_unit--permissions"></a>
### Nested Schema for `settings.create_group_approve_group.organizational_unit.permissions`

Read-Only:

- `full` (String)
- `instances` (List of String)
- `operations` (Set of String)
- `type_escaped` (String)



<a id="nestedatt--settings--create_group_approve_group--permissions"></a>
### Nested Schema for `settings.create_group_approve_group.permissions`

Read-Only:

- `full` (String)
- `instances` (List of String)
- `operations` (Set of String)
- `type_escaped` (String)



<a id="nestedatt--settings--enable_tech_admin_approve_group"></a>
### Nested Schema for `settings.enable_tech_admin_approve_group`

Read-Only:

- `admin` (Boolean)
- `links` (Attributes List) (see [below for nested schema](#nestedatt--settings--enable_tech_admin_approve_group--links))
- `name` (String)
- `organizational_unit` (Attributes) (see [below for nested schema](#nestedatt--settings--enable_tech_admin_approve_group--organizational_unit))
- `permissions` (Attributes List) (see [below for nested schema](#nestedatt--settings--enable_tech_admin_approve_group--permissions))
- `uuid` (String)

<a id="nestedatt--settings--enable_tech_admin_approve_group--links"></a>
### Nested Schema for `settings.enable_tech_admin_approve_group.links`

Read-Only:

- `href` (String)
- `id` (Number)
- `rel` (String)
- `type_escaped` (String)


<a id="nestedatt--settings--enable_tech_admin_approve_group--organizational_unit"></a>
### Nested Schema for `settings.enable_tech_admin_approve_group.organizational_unit`

Read-Only:

- `links` (Attributes List) (see [below for nested schema](#nestedatt--settings--enable_tech_admin_approve_group--organizational_unit--links))
- `name` (String)
- `permissions` (Attributes List) (see [below for nested schema](#nestedatt--settings--enable_tech_admin_approve_group--organizational_unit--permissions))
- `uuid` (String)

<a id="nestedatt--settings--enable_tech_admin_approve_group--organizational_unit--links"></a>
### Nested Schema for `settings.enable_tech_admin_approve_group.organizational_unit.links`

Read-Only:

- `href` (String)
- `id` (Number)
- `rel` (String)
- `type_escaped` (String)


<a id="nestedatt--settings--enable_tech_admin_approve_group--organizational_unit--permissions"></a>
### Nested Schema for `settings.enable_tech_admin_approve_group.organizational_unit.permissions`

Read-Only:

- `full` (String)
- `instances` (List of String)
- `operations` (Set of String)
- `type_escaped` (String)



<a id="nestedatt--settings--enable_tech_admin_approve_group--permissions"></a>
### Nested Schema for `settings.enable_tech_admin_approve_group.permissions`

Read-Only:

- `full` (String)
- `instances` (List of String)
- `operations` (Set of String)
- `type_escaped` (String)



<a id="nestedatt--settings--recovery_fallback_group"></a>
### Nested Schema for `settings.recovery_fallback_group`

Read-Only:

- `admin` (Boolean)
- `links` (Attributes List) (see [below for nested schema](#nestedatt--settings--recovery_fallback_group--links))
- `name` (String)
- `organizational_unit` (Attributes) (see [below for nested schema](#nestedatt--settings--recovery_fallback_group--organizational_unit))
- `permissions` (Attributes List) (see [below for nested schema](#nestedatt--settings--recovery_fallback_group--permissions))
- `uuid` (String)

<a id="nestedatt--settings--recovery_fallback_group--links"></a>
### Nested Schema for `settings.recovery_fallback_group.links`

Read-Only:

- `href` (String)
- `id` (Number)
- `rel` (String)
- `type_escaped` (String)


<a id="nestedatt--settings--recovery_fallback_group--organizational_unit"></a>
### Nested Schema for `settings.recovery_fallback_group.organizational_unit`

Read-Only:

- `links` (Attributes List) (see [below for nested schema](#nestedatt--settings--recovery_fallback_group--organizational_unit--links))
- `name` (String)
- `permissions` (Attributes List) (see [below for nested schema](#nestedatt--settings--recovery_fallback_group--organizational_unit--permissions))
- `uuid` (String)

<a id="nestedatt--settings--recovery_fallback_group--organizational_unit--links"></a>
### Nested Schema for `settings.recovery_fallback_group.organizational_unit.links`

Read-Only:

- `href` (String)
- `id` (Number)
- `rel` (String)
- `type_escaped` (String)


<a id="nestedatt--settings--recovery_fallback_group--organizational_unit--permissions"></a>
### Nested Schema for `settings.recovery_fallback_group.organizational_unit.permissions`

Read-Only:

- `full` (String)
- `instances` (List of String)
- `operations` (Set of String)
- `type_escaped` (String)



<a id="nestedatt--settings--recovery_fallback_group--permissions"></a>
### Nested Schema for `settings.recovery_fallback_group.permissions`

Read-Only:

- `full` (String)
- `instances` (List of String)
- `operations` (Set of String)
- `type_escaped` (String)



<a id="nestedatt--settings--remove_group_approve_group"></a>
### Nested Schema for `settings.remove_group_approve_group`

Read-Only:

- `admin` (Boolean)
- `links` (Attributes List) (see [below for nested schema](#nestedatt--settings--remove_group_approve_group--links))
- `name` (String)
- `organizational_unit` (Attributes) (see [below for nested schema](#nestedatt--settings--remove_group_approve_group--organizational_unit))
- `permissions` (Attributes List) (see [below for nested schema](#nestedatt--settings--remove_group_approve_group--permissions))
- `uuid` (String)

<a id="nestedatt--settings--remove_group_approve_group--links"></a>
### Nested Schema for `settings.remove_group_approve_group.links`

Read-Only:

- `href` (String)
- `id` (Number)
- `rel` (String)
- `type_escaped` (String)


<a id="nestedatt--settings--remove_group_approve_group--organizational_unit"></a>
### Nested Schema for `settings.remove_group_approve_group.organizational_unit`

Read-Only:

- `links` (Attributes List) (see [below for nested schema](#nestedatt--settings--remove_group_approve_group--organizational_unit--links))
- `name` (String)
- `permissions` (Attributes List) (see [below for nested schema](#nestedatt--settings--remove_group_approve_group--organizational_unit--permissions))
- `uuid` (String)

<a id="nestedatt--settings--remove_group_approve_group--organizational_unit--links"></a>
### Nested Schema for `settings.remove_group_approve_group.organizational_unit.links`

Read-Only:

- `href` (String)
- `id` (Number)
- `rel` (String)
- `type_escaped` (String)


<a id="nestedatt--settings--remove_group_approve_group--organizational_unit--permissions"></a>
### Nested Schema for `settings.remove_group_approve_group.organizational_unit.permissions`

Read-Only:

- `full` (String)
- `instances` (List of String)
- `operations` (Set of String)
- `type_escaped` (String)



<a id="nestedatt--settings--remove_group_approve_group--permissions"></a>
### Nested Schema for `settings.remove_group_approve_group.permissions`

Read-Only:

- `full` (String)
- `instances` (List of String)
- `operations` (Set of String)
- `type_escaped` (String)
