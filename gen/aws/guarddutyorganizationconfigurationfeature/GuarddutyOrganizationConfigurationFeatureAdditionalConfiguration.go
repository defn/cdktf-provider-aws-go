package guarddutyorganizationconfigurationfeature


type GuarddutyOrganizationConfigurationFeatureAdditionalConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/guardduty_organization_configuration_feature#auto_enable GuarddutyOrganizationConfigurationFeature#auto_enable}.
	AutoEnable *string `field:"required" json:"autoEnable" yaml:"autoEnable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/guardduty_organization_configuration_feature#name GuarddutyOrganizationConfigurationFeature#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}
