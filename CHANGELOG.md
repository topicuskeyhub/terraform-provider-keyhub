## 2.44.0
* Upgrade API to Topicus KeyHub version 44

## 2.43.5
* Added support to replace a list of objects by a list of UUIDs. As a side-effect, some read-only properties might now return an object instead of just a UUID, possibly necessitating changes in terraform configs.

## 2.43.0
* Upgrade API to Topicus KeyHub version 43

## 2.42.0
* Upgrade API to Topicus KeyHub version 42

## 2.41.0
* Upgrade API to Topicus KeyHub version 41
* Added support for write-only fields, specifically for totp keys in vault records.

## 2.40.1
* Upgraded the terraform-plugin-docs to v0.21.0

## 2.40.0
* Upgrade API to Topicus KeyHub version 40
* Added support for LinkableWrapperWithCount
* Byte-arrays are now handled as base64 encoded string in terraform.

## 2.39.0
* Upgrade API to Topicus KeyHub version 39

## 2.38.0
* Upgrade API to Topicus KeyHub version 38

## 2.37.0
* Upgrade API to Topicus KeyHub version 37

## 2.36.0
* Upgrade API to Topicus KeyHub version 36

## 2.35.0
* Upgrade API to Topicus KeyHub version 35

## 2.34.1
* Fix some issues in the generated documentation

## 2.34.0
* Upgrade API to Topicus KeyHub version 34

## 2.33.0
* Upgrade API to Topicus KeyHub version 33

## 2.32.0
* Upgrade API to Topicus KeyHub version 32

## 2.31.5
* Fix a crash in the reordering code

## 2.31.4
* Switch back to lists and use explicit reordering to fix #35

## 2.31.3
* Fix various crashes caused by #35

## 2.31.3
* Fix various crashes caused by #35

## 2.31.2
* Use sets for all collections to ignore changes in order (#35)
* Force replace when the parent uuid of sub resources changes (#36)

## 2.31.1
* Mark passwords and keys as sensitive

## 2.31.0
* Upgrade API to Topicus KeyHub version 31

## 2.30.7
* Use Set Attribute for unordered sets

## 2.30.6
* Fix OAuth2 client application custom attributes

## 2.30.5
* Minor updates to the documentation

## 2.30.4
* Add support for import command

## 2.30.2
* Renamed client resource to clientapplication
* Fixed naming of some properties

## 2.30.0
* First official release for the new Topicus KeyHub Terraform provider, supporting version 30

## 0.0.2
* Add support for provisioned namespaces
* Updated documentation

## 0.0.1 (Initial release)

FEATURES:
* First implementation of a new Topicus KeyHub Terraform provider
