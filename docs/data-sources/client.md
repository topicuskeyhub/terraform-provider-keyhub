---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "keyhubpreview_client Data Source - terraform-provider-keyhubpreview"
subcategory: ""
description: |-
  
---

# keyhubpreview_client (Data Source)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `uuid` (String)

### Optional

- `additional` (List of String)

### Read-Only

- `additional_objects` (Attributes) (see [below for nested schema](#nestedatt--additional_objects))
- `client_client_application_primer_type` (String)
- `client_id` (String)
- `last_modified_at` (String)
- `ldap_client` (Attributes) (see [below for nested schema](#nestedatt--ldap_client))
- `links` (Attributes List) (see [below for nested schema](#nestedatt--links))
- `name` (String)
- `o_auth2_client` (Attributes) (see [below for nested schema](#nestedatt--o_auth2_client))
- `owner` (Attributes) (see [below for nested schema](#nestedatt--owner))
- `permissions` (Attributes List) (see [below for nested schema](#nestedatt--permissions))
- `saml2_client` (Attributes) (see [below for nested schema](#nestedatt--saml2_client))
- `scopes` (List of String)
- `sso_application` (Boolean)
- `technical_administrator` (Attributes) (see [below for nested schema](#nestedatt--technical_administrator))

<a id="nestedatt--additional_objects"></a>
### Nested Schema for `additional_objects`

Read-Only:

- `audit` (Attributes) (see [below for nested schema](#nestedatt--additional_objects--audit))
- `groupclients` (Attributes) (see [below for nested schema](#nestedatt--additional_objects--groupclients))
- `groups` (Attributes) (see [below for nested schema](#nestedatt--additional_objects--groups))
- `secret` (Attributes) (see [below for nested schema](#nestedatt--additional_objects--secret))
- `tile` (Attributes) (see [below for nested schema](#nestedatt--additional_objects--tile))
- `vault_record_count` (Number)

<a id="nestedatt--additional_objects--audit"></a>
### Nested Schema for `additional_objects.audit`

Read-Only:

- `created_at` (String)
- `created_by` (String)
- `last_modified_at` (String)
- `last_modified_by` (String)


<a id="nestedatt--additional_objects--groupclients"></a>
### Nested Schema for `additional_objects.groupclients`

Read-Only:

- `items` (Attributes List) (see [below for nested schema](#nestedatt--additional_objects--groupclients--items))

<a id="nestedatt--additional_objects--groupclients--items"></a>
### Nested Schema for `additional_objects.groupclients.items`

Optional:

- `additional` (List of String)

Read-Only:

- `activation_required` (Boolean)
- `client` (Attributes) (see [below for nested schema](#nestedatt--additional_objects--groupclients--items--client))
- `group` (Attributes) (see [below for nested schema](#nestedatt--additional_objects--groupclients--items--group))
- `links` (Attributes List) (see [below for nested schema](#nestedatt--additional_objects--groupclients--items--links))
- `owner` (Attributes) (see [below for nested schema](#nestedatt--additional_objects--groupclients--items--owner))
- `permissions` (Attributes List) (see [below for nested schema](#nestedatt--additional_objects--groupclients--items--permissions))
- `technical_administrator` (Attributes) (see [below for nested schema](#nestedatt--additional_objects--groupclients--items--technical_administrator))

<a id="nestedatt--additional_objects--groupclients--items--client"></a>
### Nested Schema for `additional_objects.groupclients.items.technical_administrator`

Read-Only:

- `client_client_application_primer_type` (String)
- `client_id` (String)
- `links` (Attributes List) (see [below for nested schema](#nestedatt--additional_objects--groupclients--items--technical_administrator--links))
- `name` (String)
- `permissions` (Attributes List) (see [below for nested schema](#nestedatt--additional_objects--groupclients--items--technical_administrator--permissions))
- `scopes` (List of String)
- `sso_application` (Boolean)
- `uuid` (String)

<a id="nestedatt--additional_objects--groupclients--items--technical_administrator--links"></a>
### Nested Schema for `additional_objects.groupclients.items.technical_administrator.links`

Read-Only:

- `href` (String)
- `id` (Number)
- `rel` (String)
- `type_escaped` (String)


<a id="nestedatt--additional_objects--groupclients--items--technical_administrator--permissions"></a>
### Nested Schema for `additional_objects.groupclients.items.technical_administrator.permissions`

Read-Only:

- `full` (String)
- `instances` (List of String)
- `operations` (List of String)
- `type_escaped` (String)



<a id="nestedatt--additional_objects--groupclients--items--group"></a>
### Nested Schema for `additional_objects.groupclients.items.technical_administrator`

Read-Only:

- `admin` (Boolean)
- `links` (Attributes List) (see [below for nested schema](#nestedatt--additional_objects--groupclients--items--technical_administrator--links))
- `name` (String)
- `permissions` (Attributes List) (see [below for nested schema](#nestedatt--additional_objects--groupclients--items--technical_administrator--permissions))
- `uuid` (String)

<a id="nestedatt--additional_objects--groupclients--items--technical_administrator--links"></a>
### Nested Schema for `additional_objects.groupclients.items.technical_administrator.links`

Read-Only:

- `href` (String)
- `id` (Number)
- `rel` (String)
- `type_escaped` (String)


<a id="nestedatt--additional_objects--groupclients--items--technical_administrator--permissions"></a>
### Nested Schema for `additional_objects.groupclients.items.technical_administrator.permissions`

Read-Only:

- `full` (String)
- `instances` (List of String)
- `operations` (List of String)
- `type_escaped` (String)



<a id="nestedatt--additional_objects--groupclients--items--links"></a>
### Nested Schema for `additional_objects.groupclients.items.technical_administrator`

Read-Only:

- `href` (String)
- `id` (Number)
- `rel` (String)
- `type_escaped` (String)


<a id="nestedatt--additional_objects--groupclients--items--owner"></a>
### Nested Schema for `additional_objects.groupclients.items.technical_administrator`

Read-Only:

- `admin` (Boolean)
- `links` (Attributes List) (see [below for nested schema](#nestedatt--additional_objects--groupclients--items--technical_administrator--links))
- `name` (String)
- `permissions` (Attributes List) (see [below for nested schema](#nestedatt--additional_objects--groupclients--items--technical_administrator--permissions))
- `uuid` (String)

<a id="nestedatt--additional_objects--groupclients--items--technical_administrator--links"></a>
### Nested Schema for `additional_objects.groupclients.items.technical_administrator.links`

Read-Only:

- `href` (String)
- `id` (Number)
- `rel` (String)
- `type_escaped` (String)


<a id="nestedatt--additional_objects--groupclients--items--technical_administrator--permissions"></a>
### Nested Schema for `additional_objects.groupclients.items.technical_administrator.permissions`

Read-Only:

- `full` (String)
- `instances` (List of String)
- `operations` (List of String)
- `type_escaped` (String)



<a id="nestedatt--additional_objects--groupclients--items--permissions"></a>
### Nested Schema for `additional_objects.groupclients.items.technical_administrator`

Read-Only:

- `full` (String)
- `instances` (List of String)
- `operations` (List of String)
- `type_escaped` (String)


<a id="nestedatt--additional_objects--groupclients--items--technical_administrator"></a>
### Nested Schema for `additional_objects.groupclients.items.technical_administrator`

Read-Only:

- `admin` (Boolean)
- `links` (Attributes List) (see [below for nested schema](#nestedatt--additional_objects--groupclients--items--technical_administrator--links))
- `name` (String)
- `permissions` (Attributes List) (see [below for nested schema](#nestedatt--additional_objects--groupclients--items--technical_administrator--permissions))
- `uuid` (String)

<a id="nestedatt--additional_objects--groupclients--items--technical_administrator--links"></a>
### Nested Schema for `additional_objects.groupclients.items.technical_administrator.links`

Read-Only:

- `href` (String)
- `id` (Number)
- `rel` (String)
- `type_escaped` (String)


<a id="nestedatt--additional_objects--groupclients--items--technical_administrator--permissions"></a>
### Nested Schema for `additional_objects.groupclients.items.technical_administrator.permissions`

Read-Only:

- `full` (String)
- `instances` (List of String)
- `operations` (List of String)
- `type_escaped` (String)





<a id="nestedatt--additional_objects--groups"></a>
### Nested Schema for `additional_objects.groups`

Read-Only:

- `items` (Attributes List) (see [below for nested schema](#nestedatt--additional_objects--groups--items))

<a id="nestedatt--additional_objects--groups--items"></a>
### Nested Schema for `additional_objects.groups.items`

Optional:

- `additional` (List of String)

Read-Only:

- `admin` (Boolean)
- `application_administration` (Boolean)
- `audit_config` (Attributes) (see [below for nested schema](#nestedatt--additional_objects--groups--items--audit_config))
- `audit_requested` (Boolean)
- `auditor` (Boolean)
- `authorizing_group_auditing` (Attributes) (see [below for nested schema](#nestedatt--additional_objects--groups--items--authorizing_group_auditing))
- `authorizing_group_delegation` (Attributes) (see [below for nested schema](#nestedatt--additional_objects--groups--items--authorizing_group_delegation))
- `authorizing_group_membership` (Attributes) (see [below for nested schema](#nestedatt--additional_objects--groups--items--authorizing_group_membership))
- `authorizing_group_provisioning` (Attributes) (see [below for nested schema](#nestedatt--additional_objects--groups--items--authorizing_group_provisioning))
- `authorizing_group_types` (List of String)
- `classification` (Attributes) (see [below for nested schema](#nestedatt--additional_objects--groups--items--classification))
- `description` (String)
- `extended_access` (String)
- `hide_audit_trail` (Boolean)
- `links` (Attributes List) (see [below for nested schema](#nestedatt--additional_objects--groups--items--links))
- `name` (String)
- `nested_under` (Attributes) (see [below for nested schema](#nestedatt--additional_objects--groups--items--nested_under))
- `organizational_unit` (Attributes) (see [below for nested schema](#nestedatt--additional_objects--groups--items--organizational_unit))
- `permissions` (Attributes List) (see [below for nested schema](#nestedatt--additional_objects--groups--items--permissions))
- `private_group` (Boolean)
- `record_trail` (Boolean)
- `rotating_password_required` (Boolean)
- `single_managed` (Boolean)
- `uuid` (String)
- `vault_recovery` (String)
- `vault_requires_activation` (Boolean)

<a id="nestedatt--additional_objects--groups--items--audit_config"></a>
### Nested Schema for `additional_objects.groups.items.vault_requires_activation`

Read-Only:

- `links` (Attributes List) (see [below for nested schema](#nestedatt--additional_objects--groups--items--vault_requires_activation--links))
- `months` (List of String)
- `permissions` (Attributes List) (see [below for nested schema](#nestedatt--additional_objects--groups--items--vault_requires_activation--permissions))

<a id="nestedatt--additional_objects--groups--items--vault_requires_activation--links"></a>
### Nested Schema for `additional_objects.groups.items.vault_requires_activation.links`

Read-Only:

- `href` (String)
- `id` (Number)
- `rel` (String)
- `type_escaped` (String)


<a id="nestedatt--additional_objects--groups--items--vault_requires_activation--permissions"></a>
### Nested Schema for `additional_objects.groups.items.vault_requires_activation.permissions`

Read-Only:

- `full` (String)
- `instances` (List of String)
- `operations` (List of String)
- `type_escaped` (String)



<a id="nestedatt--additional_objects--groups--items--authorizing_group_auditing"></a>
### Nested Schema for `additional_objects.groups.items.vault_requires_activation`

Read-Only:

- `admin` (Boolean)
- `links` (Attributes List) (see [below for nested schema](#nestedatt--additional_objects--groups--items--vault_requires_activation--links))
- `name` (String)
- `permissions` (Attributes List) (see [below for nested schema](#nestedatt--additional_objects--groups--items--vault_requires_activation--permissions))
- `uuid` (String)

<a id="nestedatt--additional_objects--groups--items--vault_requires_activation--links"></a>
### Nested Schema for `additional_objects.groups.items.vault_requires_activation.links`

Read-Only:

- `href` (String)
- `id` (Number)
- `rel` (String)
- `type_escaped` (String)


<a id="nestedatt--additional_objects--groups--items--vault_requires_activation--permissions"></a>
### Nested Schema for `additional_objects.groups.items.vault_requires_activation.permissions`

Read-Only:

- `full` (String)
- `instances` (List of String)
- `operations` (List of String)
- `type_escaped` (String)



<a id="nestedatt--additional_objects--groups--items--authorizing_group_delegation"></a>
### Nested Schema for `additional_objects.groups.items.vault_requires_activation`

Read-Only:

- `admin` (Boolean)
- `links` (Attributes List) (see [below for nested schema](#nestedatt--additional_objects--groups--items--vault_requires_activation--links))
- `name` (String)
- `permissions` (Attributes List) (see [below for nested schema](#nestedatt--additional_objects--groups--items--vault_requires_activation--permissions))
- `uuid` (String)

<a id="nestedatt--additional_objects--groups--items--vault_requires_activation--links"></a>
### Nested Schema for `additional_objects.groups.items.vault_requires_activation.links`

Read-Only:

- `href` (String)
- `id` (Number)
- `rel` (String)
- `type_escaped` (String)


<a id="nestedatt--additional_objects--groups--items--vault_requires_activation--permissions"></a>
### Nested Schema for `additional_objects.groups.items.vault_requires_activation.permissions`

Read-Only:

- `full` (String)
- `instances` (List of String)
- `operations` (List of String)
- `type_escaped` (String)



<a id="nestedatt--additional_objects--groups--items--authorizing_group_membership"></a>
### Nested Schema for `additional_objects.groups.items.vault_requires_activation`

Read-Only:

- `admin` (Boolean)
- `links` (Attributes List) (see [below for nested schema](#nestedatt--additional_objects--groups--items--vault_requires_activation--links))
- `name` (String)
- `permissions` (Attributes List) (see [below for nested schema](#nestedatt--additional_objects--groups--items--vault_requires_activation--permissions))
- `uuid` (String)

<a id="nestedatt--additional_objects--groups--items--vault_requires_activation--links"></a>
### Nested Schema for `additional_objects.groups.items.vault_requires_activation.links`

Read-Only:

- `href` (String)
- `id` (Number)
- `rel` (String)
- `type_escaped` (String)


<a id="nestedatt--additional_objects--groups--items--vault_requires_activation--permissions"></a>
### Nested Schema for `additional_objects.groups.items.vault_requires_activation.permissions`

Read-Only:

- `full` (String)
- `instances` (List of String)
- `operations` (List of String)
- `type_escaped` (String)



<a id="nestedatt--additional_objects--groups--items--authorizing_group_provisioning"></a>
### Nested Schema for `additional_objects.groups.items.vault_requires_activation`

Read-Only:

- `admin` (Boolean)
- `links` (Attributes List) (see [below for nested schema](#nestedatt--additional_objects--groups--items--vault_requires_activation--links))
- `name` (String)
- `permissions` (Attributes List) (see [below for nested schema](#nestedatt--additional_objects--groups--items--vault_requires_activation--permissions))
- `uuid` (String)

<a id="nestedatt--additional_objects--groups--items--vault_requires_activation--links"></a>
### Nested Schema for `additional_objects.groups.items.vault_requires_activation.links`

Read-Only:

- `href` (String)
- `id` (Number)
- `rel` (String)
- `type_escaped` (String)


<a id="nestedatt--additional_objects--groups--items--vault_requires_activation--permissions"></a>
### Nested Schema for `additional_objects.groups.items.vault_requires_activation.permissions`

Read-Only:

- `full` (String)
- `instances` (List of String)
- `operations` (List of String)
- `type_escaped` (String)



<a id="nestedatt--additional_objects--groups--items--classification"></a>
### Nested Schema for `additional_objects.groups.items.vault_requires_activation`

Read-Only:

- `links` (Attributes List) (see [below for nested schema](#nestedatt--additional_objects--groups--items--vault_requires_activation--links))
- `name` (String)
- `permissions` (Attributes List) (see [below for nested schema](#nestedatt--additional_objects--groups--items--vault_requires_activation--permissions))
- `uuid` (String)

<a id="nestedatt--additional_objects--groups--items--vault_requires_activation--links"></a>
### Nested Schema for `additional_objects.groups.items.vault_requires_activation.links`

Read-Only:

- `href` (String)
- `id` (Number)
- `rel` (String)
- `type_escaped` (String)


<a id="nestedatt--additional_objects--groups--items--vault_requires_activation--permissions"></a>
### Nested Schema for `additional_objects.groups.items.vault_requires_activation.permissions`

Read-Only:

- `full` (String)
- `instances` (List of String)
- `operations` (List of String)
- `type_escaped` (String)



<a id="nestedatt--additional_objects--groups--items--links"></a>
### Nested Schema for `additional_objects.groups.items.vault_requires_activation`

Read-Only:

- `href` (String)
- `id` (Number)
- `rel` (String)
- `type_escaped` (String)


<a id="nestedatt--additional_objects--groups--items--nested_under"></a>
### Nested Schema for `additional_objects.groups.items.vault_requires_activation`

Read-Only:

- `admin` (Boolean)
- `links` (Attributes List) (see [below for nested schema](#nestedatt--additional_objects--groups--items--vault_requires_activation--links))
- `name` (String)
- `permissions` (Attributes List) (see [below for nested schema](#nestedatt--additional_objects--groups--items--vault_requires_activation--permissions))
- `uuid` (String)

<a id="nestedatt--additional_objects--groups--items--vault_requires_activation--links"></a>
### Nested Schema for `additional_objects.groups.items.vault_requires_activation.links`

Read-Only:

- `href` (String)
- `id` (Number)
- `rel` (String)
- `type_escaped` (String)


<a id="nestedatt--additional_objects--groups--items--vault_requires_activation--permissions"></a>
### Nested Schema for `additional_objects.groups.items.vault_requires_activation.permissions`

Read-Only:

- `full` (String)
- `instances` (List of String)
- `operations` (List of String)
- `type_escaped` (String)



<a id="nestedatt--additional_objects--groups--items--organizational_unit"></a>
### Nested Schema for `additional_objects.groups.items.vault_requires_activation`

Read-Only:

- `links` (Attributes List) (see [below for nested schema](#nestedatt--additional_objects--groups--items--vault_requires_activation--links))
- `name` (String)
- `permissions` (Attributes List) (see [below for nested schema](#nestedatt--additional_objects--groups--items--vault_requires_activation--permissions))
- `uuid` (String)

<a id="nestedatt--additional_objects--groups--items--vault_requires_activation--links"></a>
### Nested Schema for `additional_objects.groups.items.vault_requires_activation.links`

Read-Only:

- `href` (String)
- `id` (Number)
- `rel` (String)
- `type_escaped` (String)


<a id="nestedatt--additional_objects--groups--items--vault_requires_activation--permissions"></a>
### Nested Schema for `additional_objects.groups.items.vault_requires_activation.permissions`

Read-Only:

- `full` (String)
- `instances` (List of String)
- `operations` (List of String)
- `type_escaped` (String)



<a id="nestedatt--additional_objects--groups--items--permissions"></a>
### Nested Schema for `additional_objects.groups.items.vault_requires_activation`

Read-Only:

- `full` (String)
- `instances` (List of String)
- `operations` (List of String)
- `type_escaped` (String)




<a id="nestedatt--additional_objects--secret"></a>
### Nested Schema for `additional_objects.secret`

Read-Only:

- `generated_secret` (String)
- `old_secret` (String)
- `regenerate` (Boolean)


<a id="nestedatt--additional_objects--tile"></a>
### Nested Schema for `additional_objects.tile`

Read-Only:

- `uri` (String)



<a id="nestedatt--ldap_client"></a>
### Nested Schema for `ldap_client`

Read-Only:

- `bind_dn` (String)
- `client_certificate` (Attributes) (see [below for nested schema](#nestedatt--ldap_client--client_certificate))
- `share_secret_in_vault` (Boolean)
- `shared_secret` (Attributes) (see [below for nested schema](#nestedatt--ldap_client--shared_secret))
- `used_for_provisioning` (Boolean)

<a id="nestedatt--ldap_client--client_certificate"></a>
### Nested Schema for `ldap_client.client_certificate`

Read-Only:

- `alias` (String)
- `certificate_certificate_primer_type` (String)
- `certificate_data` (List of String)
- `expiration` (String)
- `fingerprint_sha1` (String)
- `fingerprint_sha256` (String)
- `global` (Boolean)
- `links` (Attributes List) (see [below for nested schema](#nestedatt--ldap_client--client_certificate--links))
- `permissions` (Attributes List) (see [below for nested schema](#nestedatt--ldap_client--client_certificate--permissions))
- `subject_dn` (String)
- `uuid` (String)

<a id="nestedatt--ldap_client--client_certificate--links"></a>
### Nested Schema for `ldap_client.client_certificate.links`

Read-Only:

- `href` (String)
- `id` (Number)
- `rel` (String)
- `type_escaped` (String)


<a id="nestedatt--ldap_client--client_certificate--permissions"></a>
### Nested Schema for `ldap_client.client_certificate.permissions`

Read-Only:

- `full` (String)
- `instances` (List of String)
- `operations` (List of String)
- `type_escaped` (String)



<a id="nestedatt--ldap_client--shared_secret"></a>
### Nested Schema for `ldap_client.shared_secret`

Read-Only:

- `color` (String)
- `links` (Attributes List) (see [below for nested schema](#nestedatt--ldap_client--shared_secret--links))
- `name` (String)
- `permissions` (Attributes List) (see [below for nested schema](#nestedatt--ldap_client--shared_secret--permissions))
- `share_end_time` (String)
- `uuid` (String)

<a id="nestedatt--ldap_client--shared_secret--links"></a>
### Nested Schema for `ldap_client.shared_secret.links`

Read-Only:

- `href` (String)
- `id` (Number)
- `rel` (String)
- `type_escaped` (String)


<a id="nestedatt--ldap_client--shared_secret--permissions"></a>
### Nested Schema for `ldap_client.shared_secret.permissions`

Read-Only:

- `full` (String)
- `instances` (List of String)
- `operations` (List of String)
- `type_escaped` (String)




<a id="nestedatt--links"></a>
### Nested Schema for `links`

Read-Only:

- `href` (String)
- `id` (Number)
- `rel` (String)
- `type_escaped` (String)


<a id="nestedatt--o_auth2_client"></a>
### Nested Schema for `o_auth2_client`

Read-Only:

- `account_permissions` (Attributes List) (see [below for nested schema](#nestedatt--o_auth2_client--account_permissions))
- `attributes` (Attributes) (see [below for nested schema](#nestedatt--o_auth2_client--attributes))
- `callback_uri` (String)
- `confidential` (Boolean)
- `debug_mode` (Boolean)
- `id_token_claims` (String)
- `initiate_login_uri` (String)
- `resource_uris` (String)
- `share_secret_in_vault` (Boolean)
- `shared_secret` (Attributes) (see [below for nested schema](#nestedatt--o_auth2_client--shared_secret))
- `show_landing_page` (Boolean)
- `use_client_credentials` (Boolean)

<a id="nestedatt--o_auth2_client--account_permissions"></a>
### Nested Schema for `o_auth2_client.account_permissions`

Read-Only:

- `full` (String)
- `instances` (List of String)
- `operations` (List of String)
- `type_escaped` (String)


<a id="nestedatt--o_auth2_client--attributes"></a>
### Nested Schema for `o_auth2_client.attributes`


<a id="nestedatt--o_auth2_client--shared_secret"></a>
### Nested Schema for `o_auth2_client.shared_secret`

Read-Only:

- `color` (String)
- `links` (Attributes List) (see [below for nested schema](#nestedatt--o_auth2_client--shared_secret--links))
- `name` (String)
- `permissions` (Attributes List) (see [below for nested schema](#nestedatt--o_auth2_client--shared_secret--permissions))
- `share_end_time` (String)
- `uuid` (String)

<a id="nestedatt--o_auth2_client--shared_secret--links"></a>
### Nested Schema for `o_auth2_client.shared_secret.links`

Read-Only:

- `href` (String)
- `id` (Number)
- `rel` (String)
- `type_escaped` (String)


<a id="nestedatt--o_auth2_client--shared_secret--permissions"></a>
### Nested Schema for `o_auth2_client.shared_secret.permissions`

Read-Only:

- `full` (String)
- `instances` (List of String)
- `operations` (List of String)
- `type_escaped` (String)




<a id="nestedatt--owner"></a>
### Nested Schema for `owner`

Read-Only:

- `admin` (Boolean)
- `links` (Attributes List) (see [below for nested schema](#nestedatt--owner--links))
- `name` (String)
- `permissions` (Attributes List) (see [below for nested schema](#nestedatt--owner--permissions))
- `uuid` (String)

<a id="nestedatt--owner--links"></a>
### Nested Schema for `owner.links`

Read-Only:

- `href` (String)
- `id` (Number)
- `rel` (String)
- `type_escaped` (String)


<a id="nestedatt--owner--permissions"></a>
### Nested Schema for `owner.permissions`

Read-Only:

- `full` (String)
- `instances` (List of String)
- `operations` (List of String)
- `type_escaped` (String)



<a id="nestedatt--permissions"></a>
### Nested Schema for `permissions`

Read-Only:

- `full` (String)
- `instances` (List of String)
- `operations` (List of String)
- `type_escaped` (String)


<a id="nestedatt--saml2_client"></a>
### Nested Schema for `saml2_client`

Read-Only:

- `attributes` (Attributes) (see [below for nested schema](#nestedatt--saml2_client--attributes))
- `metadata` (String)
- `metadata_url` (String)
- `subject_format` (String)

<a id="nestedatt--saml2_client--attributes"></a>
### Nested Schema for `saml2_client.attributes`



<a id="nestedatt--technical_administrator"></a>
### Nested Schema for `technical_administrator`

Read-Only:

- `admin` (Boolean)
- `links` (Attributes List) (see [below for nested schema](#nestedatt--technical_administrator--links))
- `name` (String)
- `permissions` (Attributes List) (see [below for nested schema](#nestedatt--technical_administrator--permissions))
- `uuid` (String)

<a id="nestedatt--technical_administrator--links"></a>
### Nested Schema for `technical_administrator.links`

Read-Only:

- `href` (String)
- `id` (Number)
- `rel` (String)
- `type_escaped` (String)


<a id="nestedatt--technical_administrator--permissions"></a>
### Nested Schema for `technical_administrator.permissions`

Read-Only:

- `full` (String)
- `instances` (List of String)
- `operations` (List of String)
- `type_escaped` (String)