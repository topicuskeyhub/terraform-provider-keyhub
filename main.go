// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/topicuskeyhub/terraform-provider-keyhubpreview/internal/provider"
)

// Run "go generate" to generate the source code, format example terraform files and generate the docs for the registry/website

// no UUID
// info
// launchpadtile
// numberseq
// groupclient
// provisioninggroup

// sub resources
// directory/{directoryid}/internalaccount
// account/{accountid}/group
// account/{accountid}/organizationalunit
// client/{clientid}/permission
// group/{groupid}/account
// organizationalunit/{organizationalunitid}/account
// system/{systemid}/group
// serviceaccount/{accountid}/group

//go:generate go run github.com/topicuskeyhub/terraform-provider-keyhub-generator@v0.0.5 --mode model
//go:generate go run github.com/topicuskeyhub/terraform-provider-keyhub-generator@v0.0.5 --mode data --resource account --linkable authAccount
//go:generate go run github.com/topicuskeyhub/terraform-provider-keyhub-generator@v0.0.5 --mode data --resource certificate --linkable certificateCertificate
//go:generate go run github.com/topicuskeyhub/terraform-provider-keyhub-generator@v0.0.5 --mode data --resource client --linkable clientClientApplication
//go:generate go run github.com/topicuskeyhub/terraform-provider-keyhub-generator@v0.0.5 --mode data --resource directory --linkable directoryAccountDirectory
//go:generate go run github.com/topicuskeyhub/terraform-provider-keyhub-generator@v0.0.5 --mode data --resource group --linkable groupGroup
//go:generate go run github.com/topicuskeyhub/terraform-provider-keyhub-generator@v0.0.5 --mode data --resource groupclassification --linkable groupGroupClassification
//go:generate go run github.com/topicuskeyhub/terraform-provider-keyhub-generator@v0.0.5 --mode data --resource organizationalunit --linkable organizationOrganizationalUnit
//go:generate go run github.com/topicuskeyhub/terraform-provider-keyhub-generator@v0.0.5 --mode data --resource serviceaccount --linkable serviceaccountServiceAccount
//go:generate go run github.com/topicuskeyhub/terraform-provider-keyhub-generator@v0.0.5 --mode data --resource system --linkable provisioningProvisionedSystem
//go:generate go run github.com/topicuskeyhub/terraform-provider-keyhub-generator@v0.0.5 --mode data --resource vaultrecord --linkable vaultVaultRecord
//go:generate go run github.com/topicuskeyhub/terraform-provider-keyhub-generator@v0.0.5 --mode data --resource webhook --linkable webhookWebhook
//go:generate go run github.com/topicuskeyhub/terraform-provider-keyhub-generator@v0.0.5 --mode resource --resource clientapplication
//go:generate go run github.com/topicuskeyhub/terraform-provider-keyhub-generator@v0.0.5 --mode resource --resource client_vaultrecord
//go:generate go run github.com/topicuskeyhub/terraform-provider-keyhub-generator@v0.0.5 --mode resource --resource group_vaultrecord
//go:generate go run github.com/topicuskeyhub/terraform-provider-keyhub-generator@v0.0.5 --mode resource --resource group
//go:generate go run github.com/topicuskeyhub/terraform-provider-keyhub-generator@v0.0.5 --mode resource --resource grouponsystem
//go:generate go run github.com/topicuskeyhub/terraform-provider-keyhub-generator@v0.0.5 --mode resource --resource serviceaccount

// If you do not have terraform installed, you can remove the formatting command, but its suggested to
// ensure the documentation is formatted properly.
//go:generate terraform fmt -recursive ./examples/

// Run the docs generation tool, check its repository for more information on how it works and how docs
// can be customized.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

var (
	// these will be set by the goreleaser configuration
	// to appropriate values for the compiled binary.
	version string = "dev"

	// goreleaser can pass other information to the main package, such as the specific commit
	// https://goreleaser.com/cookbooks/using-main.version/
)

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := providerserver.ServeOpts{
		Address: "registry.terraform.io/hashicorp/keyhubpreview",
		Debug:   debug,
	}

	err := providerserver.Serve(context.Background(), provider.New(version), opts)

	if err != nil {
		log.Fatal(err.Error())
	}
}
