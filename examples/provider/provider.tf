terraform {
  required_providers {
    keyhub = {
      source  = "registry.terraform.io/topicuskeyhub/keyhub"
      version = "=2.30.0"
    }
  }
}

variable "keyhub_secret" {
  type        = string
  description = "Client secret on KeyHub"
}

provider "keyhub" {
  issuer       = "https://keyhub.example.com"
  clientid     = "ebdf81ac-b02b-4335-9dc4-4a9bc4eb406d"
  clientsecret = var.keyhub_secret
}
