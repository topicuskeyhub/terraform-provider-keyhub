terraform {
  required_providers {
    keyhub = {
      source = "registry.terraform.io/hashicorp/keyhub-preview"
    }
  }
}

provider "keyhub" {
  issuer       = "https://keyhub.localhost:8443"
  clientid     = "ebdf81ac-b02b-4335-9dc4-4a9bc4eb406d"
  clientsecret = "o9n8b8j3TRk7A4eQfKEIDIoN-IUvRAlA3gGLUNW8"
}

data "keyhub_group" "testgroup" {
  uuid = "c20a6ed4-1ae5-4a9f-91b5-2f56f5a1522f"
}

