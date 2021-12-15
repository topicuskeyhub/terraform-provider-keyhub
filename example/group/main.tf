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

provider "keyhub" {
  issuer = "https://my.keyhub.domain.com"
  clientid = ""
  clientsecret = ""
}

provider "random" {

}

data "keyhub_groups" "all" {}

# Returns all groups
output "all_groups" {
 value = data.keyhub_groups.all.groups
}

data "keyhub_group" "group" {
  uuid = "8f4e735c-e865-4913-9671-04f4e5214add"
}

#Returns a single keyhub group
output "group" {
 value = data.keyhub_group.group
}