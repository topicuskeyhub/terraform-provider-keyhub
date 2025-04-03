# Updating and releasing Topicus KeyHub Terraform Provider 

> **Note:** First update and release the Topicus KeyHub Terraform Provider Generator code

## 1. Updating

### 1.1 Dependencies

Use the just-released version of the Topicus KeyHub Terraform Provider Generator.
This should also then update the go-sdk to the latest version transitively.

> **Note:** Make really sure the tag is pushed for the Terraform Provider Generator because a resolution failure is cached globally for an entire day.

```Shell
go get github.com/topicuskeyhub/terraform-provider-keyhub-generator@v1.0.26
```

Then update the other go dependencies

```Shell
go get -u
go mod tidy
```

### 1.2 Commit the results

```Shell
git add .
git commit -m "Upgrade dependencies"
git push
```

### 1.3 Generate provider

> **Note:** Note the three periods!

```Shell
go generate ./...
```

### 1.4 Commit the results

```Shell
git add .
git commit -m "Upgrade to KeyHub 40"
git push
```

## 2. Tag and release

```Shell
git tag v2.40.0
git push origin v2.40.0
```
