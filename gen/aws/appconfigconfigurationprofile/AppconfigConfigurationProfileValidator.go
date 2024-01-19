package appconfigconfigurationprofile


type AppconfigConfigurationProfileValidator struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appconfig_configuration_profile#type AppconfigConfigurationProfile#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appconfig_configuration_profile#content AppconfigConfigurationProfile#content}.
	Content *string `field:"optional" json:"content" yaml:"content"`
}

