terraform {
  required_providers {
    keyhub = {
      source  = "github.com/topicuskeyhub/keyhub"
      version = "0.0.1"
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

#Returns all groups
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

#Create a group
resource "keyhub_group" "new_group" {
  name = "new_keyhub_group"

  # UUID of keyhub user/account
  member {
    uuid = "00000000-0000-0000-0000-000000000000"
    rights = "MANAGER"
  }

  audit_months = ["JANUARY"]

  description = "This group is created by terraform"



  # UUID of the group that should be set as authoring group
  provisioning_auth_groupuuid = "00000000-0000-0000-0000-000000000000"
  membership_auth_groupuuid   = "00000000-0000-0000-0000-000000000000"
  auditing_auth_groupuuid     = "00000000-0000-0000-0000-000000000000"


}