# terraform-provider-keyhub

## Build provider from source

Run the following command to build the provider

```shell
$ go build -o terraform-provider-keyhub
```

## Test sample configuration

First, build and install the provider so that version 0.0.1 is available on your machine. Edit the makefile and set OS_ARCH when needed.

```shell
$ make
```

Then, navigate to the `examples` directory. 

```shell
$ cd examples
```

Run the following command to initialize the workspace and apply the sample configuration.
You might need to edit the .tf file to set some configuration.

```shell
$ terraform init && terraform apply


## Contributing

1. Install terraform:  https://www.terraform.io/downloads
2. Install golang: https://go.dev/dl/
3. Go wild.