// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NullProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (n *jsiiProxy_NullProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateNullProvider_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateNullProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNullProvider_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateNullProvider_IsTerraformProviderParameters(x interface{}) error {
	return nil
}

func validateNewNullProviderParameters(scope constructs.Construct, id *string, config *NullProviderConfig) error {
	return nil
}

