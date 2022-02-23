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

data "keyhub_accounts" "all" {}

#Returns all accounts
output "all_accounts" {
 value = data.keyhub_accounts.all.accounts
}

data "keyhub_account" "account" {
  uuid = ""
}

#Returns a single keyhub account
output "account" {
 value = data.keyhub_group.account
}