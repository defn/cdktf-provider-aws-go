//go:build no_runtime_type_checking

package autoscalinggroup

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AutoscalingGroupTrafficSourceList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AutoscalingGroupTrafficSourceList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AutoscalingGroupTrafficSourceList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AutoscalingGroupTrafficSourceList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AutoscalingGroupTrafficSourceList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AutoscalingGroupTrafficSourceList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AutoscalingGroupTrafficSourceList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAutoscalingGroupTrafficSourceListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}
