// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type NullProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/null/3.2.4/docs#alias NullProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
}

