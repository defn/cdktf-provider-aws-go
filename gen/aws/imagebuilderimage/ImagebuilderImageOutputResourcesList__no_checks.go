//go:build no_runtime_type_checking

package imagebuilderimage

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ImagebuilderImageOutputResourcesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (i *jsiiProxy_ImagebuilderImageOutputResourcesList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_ImagebuilderImageOutputResourcesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ImagebuilderImageOutputResourcesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ImagebuilderImageOutputResourcesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ImagebuilderImageOutputResourcesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewImagebuilderImageOutputResourcesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

