package cloudformationtype


type CloudformationTypeLoggingConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/cloudformation_type#log_group_name CloudformationType#log_group_name}.
	LogGroupName *string `field:"required" json:"logGroupName" yaml:"logGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/cloudformation_type#log_role_arn CloudformationType#log_role_arn}.
	LogRoleArn *string `field:"required" json:"logRoleArn" yaml:"logRoleArn"`
}

