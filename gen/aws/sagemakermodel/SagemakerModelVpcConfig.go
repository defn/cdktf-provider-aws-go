package sagemakermodel


type SagemakerModelVpcConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/sagemaker_model#security_group_ids SagemakerModel#security_group_ids}.
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/sagemaker_model#subnets SagemakerModel#subnets}.
	Subnets *[]*string `field:"required" json:"subnets" yaml:"subnets"`
}

