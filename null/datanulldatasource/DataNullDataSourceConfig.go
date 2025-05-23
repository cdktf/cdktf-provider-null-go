// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datanulldatasource

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataNullDataSourceConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// If set, its literal value will be stored and returned.
	//
	// If not, its value defaults to `"default"`. This argument exists primarily for testing and has little practical use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/null/3.2.4/docs/data-sources/data_source#has_computed_default DataNullDataSource#has_computed_default}
	HasComputedDefault *string `field:"optional" json:"hasComputedDefault" yaml:"hasComputedDefault"`
	// A map of arbitrary strings that is copied into the `outputs` attribute, and accessible directly for interpolation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/null/3.2.4/docs/data-sources/data_source#inputs DataNullDataSource#inputs}
	Inputs *map[string]*string `field:"optional" json:"inputs" yaml:"inputs"`
}

