# Terraform Provider for Topicus KeyHub

The Terraform Provider for Topicus KeyHub allows managing resources within a Topicus KeyHub appliance.
It requires at least Terraform 1.0 and the most recent version of Terraform is recommended.

~> **Important** Interacting with Topicus KeyHub from Terraform causes any secrets that you read and write to be persisted in plain text in both Terraform's state file *and* in any generated plan files.
For any Terraform module that reads or writes Topicus KeyHub secrets, these files should be treated as sensitive and protected accordingly.
Access to secrets should be restricted as much as possible.

For more information see:
* [Terraform Website](https://www.terraform.io)
* [Topicus KeyHub Provider Documentation](https://registry.terraform.io/providers/topicuskeyhub/keyhub/latest/docs)

The used provider version must match your Topicus KeyHub release.
For example, use the provider version 2.30.0 for Topicus KeyHub 30.
An older version of the provider may work on a newer version of Topicus KeyHub.
A newer version of the provder will fail on an older version of Topicus KeyHub.

## Usage example

```hcl
# 1. Specify the version of the Topicus KeyHub Provider to use
terraform {
  required_providers {
    keyhub = {
      source = "registry.terraform.io/topicuskeyhub/keyhub"
      version = "=2.30.0"
    }
  }
}

# 2. Configure the Topicus KeyHub provider
variable "keyhub_secret" {
  type        = string
  description = "Client secret on KeyHub"
}

provider "keyhub" {
  issuer       = "https://keyhub.example.com"
  clientid     = "ebdf81ac-b02b-4335-9dc4-4a9bc4eb406d"
  clientsecret = var.keyhub_secret
}

# 3. Create a group in Topicus KeyHub
resource "keyhub_group" "group_in_keyhub" {
  name = "Terraform"
  accounts = [{
    uuid   = "7ea6622b-f9d2-4e52-a799-217b26f88376"
    rights = "MANAGER"
  }]
  client_permissions = [{
    client_uuid = "ebdf81ac-b02b-4335-9dc4-4a9bc4eb406d"
    value       = "GROUP_FULL_VAULT_ACCESS"
  }]
}

# 4. Create a vault record in the newly created group
resource "keyhub_group_vaultrecord" "vaultrecord_in_keyhub" {
  name       = "Terraform Record"
  group_uuid = resource.keyhub_group.group_in_keyhub.uuid
  secret = {
    password = "test3"
  }
}

# 5. Setup provisioning for the group
resource "keyhub_grouponsystem" "provisioning" {
  provisioned_system_uuid = "47923975-b1af-47c8-bd7a-e52ebb4b9b84"
  owner_uuid              = resource.keyhub_group.group_in_keyhub.uuid
  name_in_system          = "cn=terraform,ou=groups,dc=demo,dc=topicus-keyhub,dc=com"
  type                    = "GROUP"
  provgroups = [{
    activation_required = "false"
    group_uuid          = "c6c98d08-2cbf-45e9-937a-c5c0427348e2"
  }]
}
```
