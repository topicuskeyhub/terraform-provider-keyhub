--- 
layout: ""
page_title: "Topicus KeyHub Provider"
subcategory: ""
description: |-
  Terraform provider for interacting with your Topicus KeyHub REST API.
  
---

# KeyHub Provider

This is the Topicus KeyHub Terraform Provider. It is used to interact with your Topicus KeyHub REST API.

The provider allows you to safely retrieve accounts, groups and vaultrecords and store groups and vaultrecords.


## Example Usage

We'll assume your Topicus KeyHub installation is externally reachable via "https://keyhub.domain.com". Replace this URL with your actual installation URL.

Create an OAuth2/OIDC Client Application through the console:
https://keyhub.domain.com/console/access

Write down the Client Identifier and Client Secret and use them in your terraform configuration. Do not keep your authentication credentials in HCL for production environments, use Terraform environment variables.

**For local testing:**
```terraform
provider "keyhub" {
  issuer = "https://keyhub.domain.com"
  clientid = "myclientid"
  clientsecret = "myclientsecret"
}
```

**Using environment variables:**
```shell
export KEYHUB_ISSUER=https://keyhub.domain.com
export KEYHUB_CLIENTID=myclientid
export KEYHUB_CLIENTSECRET=myclientsecret
```

```terraform
provider "keyhub" {
}
```

## Schema

### Required

No configuration is required however it must be either defined in your Terraform configuration or using environment variables.

### Optional

- **issuer** (String) The URL of your Topicus KeyHub installation (eg: https://keyhub.domain.com)
- **clientid** (String, Sensitive) The Client Identifier of the OAuth2/OIDC Application defined in your Topicus KeyHub installation
- **clientsecret** (String, Sensitive) The Client Secret of the OAuth2/OIDC Application defined in your Topicus KeyHub installation
