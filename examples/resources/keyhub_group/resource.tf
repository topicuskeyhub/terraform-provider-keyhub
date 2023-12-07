resource "keyhub_group" "terra" {
  name                       = "Terraform"
  application_administration = "true"
  accounts = [{
    uuid   = "7ea6622b-f9d2-4e52-a799-217b26f88376"
    rights = "MANAGER"
  }]
  client_permissions = [{
    client_uuid = "ebdf81ac-b02b-4335-9dc4-4a9bc4eb406d"
    value       = "GROUP_FULL_VAULT_ACCESS"
  }]
}
