package lambdafunction


type LambdaFunctionFileSystemConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/lambda_function#arn LambdaFunction#arn}.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/lambda_function#local_mount_path LambdaFunction#local_mount_path}.
	LocalMountPath *string `field:"required" json:"localMountPath" yaml:"localMountPath"`
}

