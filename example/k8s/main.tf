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

provider "random" {
  # Configuration options
}

provider "keyhub" {
  issuer       = var.keyhub_instance
  clientid     = var.keyhub_client_id
  clientsecret = var.keyhub_client_secret
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


  cluster      = "demo01"

}

resource "keyhub_clientapplication" "control" {
  name = "Kubernetes Secrets Controller - ${local.cluster}"
  owner = local.uuids.umbrella
  type = "OAUTH2"
  is_server2server = true
}

resource "keyhub_group" "control" {
  name = "K8S ${local.cluster}/Default/keyhub-secrets-cntrl"
  application_administration = true
  member { uuid = local.uuids.user1 }
  member { uuid = local.uuids.user2 }

  client {
    uuid = local.uuids.baseclient
    permissions = [ "GROUP_FULL_VAULT_ACCESS", "GROUP_READ_CONTENTS" ]
  }

  client {
    uuid = keyhub_clientapplication.control.uuid
    permissions = [ "GROUP_FULL_VAULT_ACCESS", "GROUP_READ_CONTENTS" ]
  }

  depends_on = [keyhub_clientapplication.control]

}

resource "keyhub_vaultrecord" "control" {
  name = keyhub_group.control.name
  groupuuid = keyhub_group.control.uuid
  username = keyhub_clientapplication.control.clientid
  password = keyhub_clientapplication.control.clientsecret
  depends_on = [keyhub_group.control, keyhub_clientapplication.control]
}

/** Start PROJECT: Default **/

locals {
  project = "Default"
  project_id = "p-xxxxx"
}
resource "keyhub_clientapplication" "prj_default" {
  name = "Kubernetes Secrets Controller - ${local.cluster}/${local.project}"
  owner = local.uuids.umbrella
  type = "OAUTH2"
  is_server2server = true
}

resource "keyhub_vaultrecord" "prj_default" {
  name = keyhub_clientapplication.prj_default.name
  groupuuid = keyhub_group.control.uuid
  username = keyhub_clientapplication.prj_default.clientid
  password = keyhub_clientapplication.prj_default.clientsecret
  comment = "policies:\n  - type: namespace\n    labelSelector: field.cattle.io/projectId=${local.project_id}"
  depends_on = [keyhub_group.control, keyhub_clientapplication.prj_default]
}

resource "keyhub_group" "prj_default" {
  name = "K8S ${local.cluster}/Default"
  application_administration = true
  member { uuid = local.uuids.user1 }
  member { uuid = local.uuids.user2 }

  client {
    uuid = local.uuids.baseclient
    permissions = [ "GROUP_FULL_VAULT_ACCESS", "GROUP_READ_CONTENTS" ]
  }

  client {
    uuid = keyhub_clientapplication.prj_default.uuid
    permissions = [ "GROUP_FULL_VAULT_ACCESS", "GROUP_READ_CONTENTS" ]
  }
  depends_on = [keyhub_group.control, keyhub_clientapplication.prj_default]
}

/** End PROJECT: Default **/

/** Start records for   Default **/
resource "keyhub_clientapplication" "argocd" {
  name = "Argo CD - ${local.cluster}"
  owner = keyhub_group.control.uuid
  type = "OAUTH2"
  is_sso = true
  callback_uri = "https://argocd.${local.cluster}.example.com/auth/callback"
  attribute {
    name = "groups"
    script = "return groups.map(function (group) {return group.uuid;});"
  }
  depends_on = [keyhub_group.control]
}

resource "keyhub_vaultrecord" "argocd" {
  name = keyhub_clientapplication.argocd.name
  groupuuid = keyhub_group.prj_default.uuid
  username = keyhub_clientapplication.argocd.clientid
  password = keyhub_clientapplication.argocd.clientsecret
  depends_on = [keyhub_group.prj_default, keyhub_clientapplication.argocd]
}

resource "keyhub_clientapplication" "proxy" {
  name = "OAuth2 Proxy - ${local.cluster}"
  owner = keyhub_group.control.uuid
  type = "OAUTH2"
  is_sso = true
  callback_uri = "https://{1-42}.${local.cluster}.example.com/oauth2/callback"
  attribute {
    name = "email"
    script = "return account.email;"
  }
  attribute {
    name = "groups"
    script = "return groups.map(function (group) {return group.uuid;});"
  }
  depends_on = [keyhub_group.control]
}

resource "keyhub_vaultrecord" "proxy" {
  name = keyhub_clientapplication.proxy.name
  groupuuid = keyhub_group.prj_default.uuid
  username = keyhub_clientapplication.proxy.clientid
  password = keyhub_clientapplication.proxy.clientsecret
  depends_on = [keyhub_group.prj_default,keyhub_clientapplication.prj_default]
}