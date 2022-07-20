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

variable "keyhub_instance" { type = string }
variable "keyhub_client_id" { type = string }
variable "keyhub_client_secret" { type = string }

provider "keyhub" {
  issuer       = var.keyhub_instance
  clientid     = var.keyhub_client_id
  clientsecret = var.keyhub_client_secret
}

provider "random" {
  # Configuration options
}

locals { 
  uuids = {
    "user1" : "b268b920-6728-40eb-a28c-668299bcf0a2",
    "user2" : "6cabb44b-19cc-4f92-8d5a-09d26b6aaa53",

    "baseclient" : "626bf2fd-2750-42cd-b2ed-9712a01bf267",

    "umbrella" : "0c9afdaa-dcaf-4864-9b36-4a3b11a28f11",

    "systemMSAD" : "325e713f-98ad-4449-b710-d8196ece2ffb",
    "systemLDAP" : "46a39091-2330-40b6-a198-29d1fb7abcaa",
  }
  
}

data "keyhub_group" "umbrella" {
  uuid = local.uuids.umbrella
}

data "keyhub_provisionedsystem" "ldap" {
  uuid = local.uuids.systemLDAP
}


resource "keyhub_grouponsystem" "umbrella" {

  type = "POSIX_GROUP"
  owner = data.keyhub_group.umbrella.uuid
  system = data.keyhub_provisionedsystem.ldap.uuid
  name = "Umbrella"

}

resource "keyhub_grouponsystem" "umbrella_provgrps" {
  type = "POSIX_GROUP"
  owner = data.keyhub_group.umbrella.uuid
  system = data.keyhub_provisionedsystem.ldap.uuid
  name = "Umbrella Prov"

  provgroup {
    group = "ff3e8fa1-92bd-4149-91f1-d4b85d2387af"
    securitylevel = "HIGH"
    static = false
  }

  provgroup {
    group = "8673c48c-c549-4b94-9752-486473b1030c"
    securitylevel = "HIGH"
    static = false
  }
}





output "main" {
  value = {
    "umbrella" : data.keyhub_group.umbrella
    "ldap" : data.keyhub_provisionedsystem.ldap
    "umbrella_gos" : keyhub_grouponsystem.umbrella
  }
}

