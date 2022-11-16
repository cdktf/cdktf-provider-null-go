//go:build no_runtime_type_checking

package resource

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_Resource) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (r *jsiiProxy_Resource) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Resource) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Resource) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Resource) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Resource) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Resource) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Resource) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Resource) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Resource) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Resource) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Resource) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateResource_IsConstructParameters(x interface{}) error {
	return nil
}

func validateResource_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateResource_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Resource) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Resource) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Resource) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Resource) validateSetTriggersParameters(val *map[string]*string) error {
	return nil
}

func validateNewResourceParameters(scope constructs.Construct, id *string, config *ResourceConfig) error {
	return nil
}

