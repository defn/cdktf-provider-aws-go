//go:build no_runtime_type_checking

package backupplan

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BackupPlanRuleList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (b *jsiiProxy_BackupPlanRuleList) validateGetParameters(index *float64) error {
	return nil
}

func (b *jsiiProxy_BackupPlanRuleList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_BackupPlanRuleList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BackupPlanRuleList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BackupPlanRuleList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_BackupPlanRuleList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewBackupPlanRuleListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

