terraform {
  required_providers {
    keyhub = {
      source  = "github.com/topicuskeyhub/keyhub"
      version = "0.1"
    }
    random = {
      source = "hashicorp/random"
      version = "3.1.0"
    }
  }
}

provider "random" {
}

provider "keyhub" {
  issuer = "https://my.keyhub.domain.com"
  clientid = ""
  clientsecret = ""
}

data "keyhub_group" "group" {
  uuid = ""
}

data "keyhub_vaultrecord" "firstvaultrecord" {
  groupuuid = "${data.keyhub_group.group.uuid}"
  uuid = ""
}

#Returns a single vaultrecord
output "firstvaultrecord" {
  value = data.keyhub_vaultrecord.firstvaultrecord
  sensitive = true
}

#Returns the comment of a single vaultrecord
output "firstvaultrecord_comment" {
  value = data.keyhub_vaultrecord.firstvaultrecord.comment
  sensitive = true
}

resource "random_uuid" "randomuuid" {
}

resource "keyhub_vaultrecord" "secondvaultrecord" {
  groupuuid = "${data.keyhub_group.group.uuid}"
  name = "VaultRecord ${random_uuid.randomuuid.result}"
  password = "Random"
}

#Returns the created vaultrecord
output "secondvaultrecord" {
  value = resource.keyhub_vaultrecord.secondvaultrecord
  sensitive = true
}