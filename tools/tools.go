// Copyright (c) Topicus Security B.V.
// SPDX-License-Identifier: APSL-2.0

//go:build tools

package tools

import (
	// Documentation generation
	_ "github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs"
	_ "github.com/topicuskeyhub/terraform-provider-keyhub-generator"
)
