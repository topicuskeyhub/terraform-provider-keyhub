// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build tools

package tools

import (
	// Documentation generation
	_ "github.com/hashicorp/terraform-plugin-docs/schemamd"
	_ "github.com/topicuskeyhub/terraform-provider-keyhub-generator/model"
)
