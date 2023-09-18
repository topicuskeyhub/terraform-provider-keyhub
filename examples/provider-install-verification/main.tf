terraform {
  required_providers {
    keyhubpreview = {
      source = "registry.terraform.io/hashicorp/keyhubpreview"
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

provider "keyhubpreview" {
  #  issuer       = "https://keyhub.topicusonderwijs.nl"
  #  clientid     = "3a5e82ad-3f0d-4a63-846b-4b3e431f1135"
  issuer       = "https://keyhub.localhost:8443"
  clientid     = "ebdf81ac-b02b-4335-9dc4-4a9bc4eb406d"
  clientsecret = var.keyhub_secret_local
}

data "keyhubpreview_group" "test" {
  #  uuid = "2fb85263-6406-44f9-9e8a-b1a6d1f43250"
  uuid       = "c6c98d08-2cbf-45e9-937a-c5c0427348e2"
  additional = ["accounts"]
}

output "data_group" {
  value = data.keyhubpreview_group.test
}

#resource "keyhubpreview_group" "terra" {
#  name = "Terraform"
#  additional_objects = {
#    accounts = {
#      items = [{
#        uuid   = "7ea6622b-f9d2-4e52-a799-217b26f88376"
#        rights = "MANAGER"
#      }]
#    }
#  }
#  admin = true
#}
#
#output "resource_group" {
#  value = resource.keyhubpreview_group.terra
#}
