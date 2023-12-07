// Copyright (c) Topicus Security B.V.
// SPDX-License-Identifier: APSL-2.0

package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/topicuskeyhub/terraform-provider-keyhub/internal/provider"
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

//go:generate go run github.com/topicuskeyhub/terraform-provider-keyhub-generator --mode model
//go:generate go run github.com/topicuskeyhub/terraform-provider-keyhub-generator --mode data --resource account --linkable authAccount
//go:generate go run github.com/topicuskeyhub/terraform-provider-keyhub-generator --mode data --resource certificate --linkable certificateCertificate
//go:generate go run github.com/topicuskeyhub/terraform-provider-keyhub-generator --mode data --resource clientapplication --linkable clientClientApplication
//go:generate go run github.com/topicuskeyhub/terraform-provider-keyhub-generator --mode data --resource directory --linkable directoryAccountDirectory
//go:generate go run github.com/topicuskeyhub/terraform-provider-keyhub-generator --mode data --resource group --linkable groupGroup
//go:generate go run github.com/topicuskeyhub/terraform-provider-keyhub-generator --mode data --resource groupclassification --linkable groupGroupClassification
//go:generate go run github.com/topicuskeyhub/terraform-provider-keyhub-generator --mode data --resource organizationalunit --linkable organizationOrganizationalUnit
//go:generate go run github.com/topicuskeyhub/terraform-provider-keyhub-generator --mode data --resource serviceaccount --linkable serviceaccountServiceAccount
//go:generate go run github.com/topicuskeyhub/terraform-provider-keyhub-generator --mode data --resource system --linkable provisioningProvisionedSystem
//go:generate go run github.com/topicuskeyhub/terraform-provider-keyhub-generator --mode data --resource vaultrecord --linkable vaultVaultRecord
//go:generate go run github.com/topicuskeyhub/terraform-provider-keyhub-generator --mode data --resource webhook --linkable webhookWebhook
//go:generate go run github.com/topicuskeyhub/terraform-provider-keyhub-generator --mode resource --resource clientapplication
//go:generate go run github.com/topicuskeyhub/terraform-provider-keyhub-generator --mode resource --resource client_vaultrecord
//go:generate go run github.com/topicuskeyhub/terraform-provider-keyhub-generator --mode resource --resource group_vaultrecord
//go:generate go run github.com/topicuskeyhub/terraform-provider-keyhub-generator --mode resource --resource group
//go:generate go run github.com/topicuskeyhub/terraform-provider-keyhub-generator --mode resource --resource grouponsystem
//go:generate go run github.com/topicuskeyhub/terraform-provider-keyhub-generator --mode resource --resource serviceaccount
//go:generate terraform fmt -recursive ./examples/
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

var (
	// these will be set by the goreleaser configuration
	// to appropriate values for the compiled binary.
	version string = "dev"
)

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := providerserver.ServeOpts{
		Address: "registry.terraform.io/hashicorp/keyhub",
		Debug:   debug,
	}

	err := providerserver.Serve(context.Background(), provider.New(version), opts)

	if err != nil {
		log.Fatal(err.Error())
	}
}
