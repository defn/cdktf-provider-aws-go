package lambdafunction


type LambdaFunctionSnapStart struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/lambda_function#apply_on LambdaFunction#apply_on}.
	ApplyOn *string `field:"required" json:"applyOn" yaml:"applyOn"`
}

