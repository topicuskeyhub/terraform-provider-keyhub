---
page_title: "keyhub_vaultrecord Resource - terraform-provider-keyhub"
subcategory: ""
description: |-
  The vaultrecord resource allows you to store/retrieve/update/delete  information about one KeyHub vaultrecord.
  
---

# keyhub_vaultrecord (Resource)

The vaultrecord resource allows you to store/retrieve/update/delete information about one KeyHub vaultrecord.

## Example Usage

```hcl
resource "keyhub_vaultrecord" "example" {
  groupuuid = "example"
  name = "Example"
  password = "Example"
}
```

Write text to keyhub:
```hcl
data "local_file" "text_in" {
  filename = "${path.module}/example.txt"
}

resource "keyhub_vaultrecord" "text_file" {
  groupuuid = "example"
  name      = "example - Text"
  filename  = basename(data.local_file.text_in.filename)
  file      = data.local_file.text_in.content
}
```

Write binary to keyhub:
```hcl
data "local_file" "png_in" {
  filename = "${path.module}/example.png"
}

resource "keyhub_vaultrecord" "png_file" {
  groupuuid      = "example"
  name           = "example - PNG"
  filename       = basename(data.local_file.png_in.filename)
  file           = data.local_file.png_in.content_base64
  base64_encoded = true
}
```


## Schema

### Required

- **groupuuid** (String) The group UUID of the vaultrecord you wish to store/retrieve/update/delete
- **name** (String) The Name field of the vaultrecord

### Optional

- **url** (String)
- **username** (String)
- **filename** (String)
- **file** (string) Content of file 
- **base64_encoded** (boolean) (Bool) If true, the value of `file` must be base64 encoded  

At least one of the following is required:

- **comment** (String, Sensitive) The value of the Comment field of the vaultrecord. This value is sensitive as it might contain secret information.
- **password** (String, Sensitive)  The value of the Password field of the vaultrecord. This value is sensitive as it might contain secret information.
- **totp** (String, Sensitive)  The value of the Totp field of the vaultrecord. This value is sensitive as it might contain secret information.

### Read-Only

- **id** (String) The value of the ID field of the vaultrecord
- **uuid** (String) The UUID of the vaultrecord 


