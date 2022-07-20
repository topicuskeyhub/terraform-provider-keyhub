terraform {
  required_version = ">= 0.14"
  required_providers {
    random = {
      source = "hashicorp/random"
      version = "3.1.3"
    }
    keyhub = {
      source  = "topicuskeyhub/keyhub"
    }

  }
}

provider "random" {
  # Configuration options
}

locals { 
  uuids = {
    "umbrella" : "0c9afdaa-dcaf-4864-9b36-4a3b11a28f11",
  }
  
}



variable "keyhub_instance" { type = string }
variable "keyhub_client_id" { type = string }
variable "keyhub_client_secret" { type = string }

provider "keyhub" {
  issuer       = var.keyhub_instance
  clientid     = var.keyhub_client_id
  clientsecret = var.keyhub_client_secret
}

resource "keyhub_clientapplication" "server2server" {
  is_server2server    = true
  name                = "Terraform - OIDC Client"
  owner               = local.uuids.umbrella
  type                = "OAUTH2"
}


resource "keyhub_clientapplication" "sso" {
  callbackuri         = "https://{1-42}.example.com/oauth2/callback"
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


output "main" {
  value = {
    "server2server" : keyhub_clientapplication.server2server
    "sso" : keyhub_clientapplication.sso
  }
}

