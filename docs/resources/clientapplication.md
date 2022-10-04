---
page_title: "keyhub_clientapplication Resource - terraform-provider-keyhub"
subcategory: ""
description: |-
The group resource allows you to store/retrieve/<s>update</s>/<s>delete</s> information about one KeyHub group on system.

---

# keyhub_clientapplication (Resource)

The group resource allows you to create and retrieve information about a clientApplication.
*Note:* KeyHub currently only supports the creation of a client, Update or Delete isn't possible

## Example Usage

```terraform

resource "keyhub_clientapplication" "server2server" {
  is_server2server    = true
  name                = "Terraform - OIDC Client"
  owner               = local.uuids.umbrella
  type                = "OAUTH2"
}


resource "keyhub_clientapplication" "sso" {
  callback_uri         = "https://{1-42}.example.com/oauth2/callback"
  is_sso              = true
  name                = "Terraform - OAuth2 SSO client"
  owner               = local.uuids.umbrella
  scopes              = [ "profile" ]
  type                = "OAUTH2"

  attribute {
    name   = "email"
    script = "return account.email;"
  }
  attribute {
    name   = "groups"
    script = "return groups.map(function (group) {return group.uuid;});"
  }
}

```

## Schema

### Required

- **name** (String) The name of the created client application
- **owner** (String) The UUID of the group that is owner of the client

### Optional

- **attribute** (Block) The additional attributes which can be retrieved through userinfo
- **callback_uri** (String) The client is a code and implicit client
- **confidential** (Bool) A confidential (or trusted) OAUTH2 client is able to keep its credentials a secret
- **id_token_claims** (String) A space-separated list of claims. These claims are added to the 'id_token', even if the client does not explicitly request them to be added
- **initiate_login_uri** (String) The URI within the application where a third party login can be started.
- **is_server2server** (Bool) The client is a code and implicit client
- **is_sso** (Bool) If set the client application is a Single Sign On application
- **metadata** (String) The metadata of the SAML application, use if metadata_url is not available
- **metadata_url** (String) The URL to retrieve the SAML metadata from
- **scopes** (List) The Allowed scopes for the client application. For SSO applications this defaults to `profile`
- **show_landingpage** (Bool) Displays a landing page immediately after login and before redirecting to the SSO application
- **subject_format** (String) Topicus KeyHub can deliver the subject in three formats: Primary identifier, UPN and username.
- **technical_administrator** (String) The UUID of the group that is technical administrator of the client, default to Owner
- **type** (String) The type of the client application. Possible values: `OAUTH2` (default), `LDAP`, `SAML2`

### Read-Only

- **binddn** (String) The identifier of the application to be used by a simple bind
- **clientid** (String) The clientid (UUID) of the client application.
- **clientsecret** (String) The secret (UUID) of the client application.
- **id** (String) The value of the ID field of the client application
- **uuid** (String) The UUID of the created client application

### Blocks

The *attribute* block supports the following:
- **name** (String, Required) The name of the attribute
- **script** (String, Required) The script defines the body of a `function (account, groups) { <SCRIPT> }` function, and is written in JavaScript using the ECMAScript 5 standard


## Import

KeyHub group can be imported using the uuid, e.g.

```
$ terraform import keyhub_clientapplication.example "00000000-0000-0000-0000-000000000000"
```
