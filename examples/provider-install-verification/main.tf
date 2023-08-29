terraform {
  required_providers {
    keyhub = {
      source = "registry.terraform.io/hashicorp/keyhub-preview"
    }
  }
}

provider "keyhub" {
  issuer       = "https://keyhub.topicusonderwijs.nl"
  clientid     = "3a5e82ad-3f0d-4a63-846b-4b3e431f1135"
  clientsecret = "P0OERPwK2hHJGS68Op2epiXvcVy4GDpx98huzagI"
}

data "keyhub_group" "test" {
  uuid = "2fb85263-6406-44f9-9e8a-b1a6d1f43250"
}

output "testgroup" {
  value = data.keyhub_group.test
}
