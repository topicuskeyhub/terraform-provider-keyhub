terraform {
  required_providers {
    keyhub = {
      source = "registry.terraform.io/hashicorp/keyhub"
    }
  }
}

variable "keyhub_secret" {
  type        = string
  description = "Client secret on KeyHub"
}

variable "keyhub_secret_local" {
  type        = string
  description = "Client secret on KeyHub"
}

provider "keyhub" {
  #  issuer       = "https://keyhub.topicusonderwijs.nl"
  #  clientid     = "3a5e82ad-3f0d-4a63-846b-4b3e431f1135"
  issuer       = "https://keyhub.localhost:8443"
  clientid     = "ebdf81ac-b02b-4335-9dc4-4a9bc4eb406d"
  clientsecret = var.keyhub_secret_local
}

resource "keyhub_group" "terra" {
  name = "Terraform"
  accounts = [{
    uuid   = "2948741d-f852-4599-be0e-cf187b306b4b"
    rights = "MANAGER"
  },{
  	uuid   = "7ea6622b-f9d2-4e52-a799-217b26f88376"
    rights = "MANAGER"
  }]
  client_permissions = [{
    client_uuid = "ebdf81ac-b02b-4335-9dc4-4a9bc4eb406d"
    value       = "GROUP_FULL_VAULT_ACCESS"
  }]
  application_administration = "true"
  additional                 = ["audit"]
}

resource "keyhub_group_vaultrecord" "terrarecord" {
  name       = "Terraform Record"
  group_uuid = resource.keyhub_group.terra.uuid
  secret = {
    password = "test3"
  }
}

resource "keyhub_grouponsystem" "terragos" {
  provisioned_system_uuid = "47923975-b1af-47c8-bd7a-e52ebb4b9b84"
  owner_uuid              = resource.keyhub_group.terra.uuid
  name_in_system          = "cn=terraform,ou=groups,ou=dev,dc=ad01,dc=keyhub,dc=s25,dc=topicus,dc=education"
  type                    = "GROUP"
  provgroups = [{
    activation_required = "false"
    group_uuid          = "c6c98d08-2cbf-45e9-937a-c5c0427348e2"
  }]
}

resource "keyhub_clientapplication" "oauth2client" {
  name                         = "Created by Terraform"
  technical_administrator_uuid = resource.keyhub_group.terra.uuid
  owner_uuid                   = resource.keyhub_group.terra.uuid
  scopes                       = ["profile", "access_vault"]
  oauth2_client = {
    debug_mode = "true"
  }
}

resource "keyhub_clientapplication" "saml2client" {
  name                         = "SAML2 Created by Terraform"
  client_id                    = "https://keyhub-vm3"
  technical_administrator_uuid = resource.keyhub_group.terra.uuid
  owner_uuid                   = resource.keyhub_group.terra.uuid
  saml2_client = {
    subject_format = "ID"
    metadata_url   = "https://does.not.exist.localhost:8600/metadata"
  }
}

resource "keyhub_serviceaccount" "sa" {
  name                         = "Service account by Terraform"
  technical_administrator_uuid = resource.keyhub_group.terra.uuid
  system_uuid                  = "47923975-b1af-47c8-bd7a-e52ebb4b9b84"
  username                     = "terraform"
}
