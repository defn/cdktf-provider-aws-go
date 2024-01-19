package pinpointapp


type PinpointAppCampaignHook struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/pinpoint_app#lambda_function_name PinpointApp#lambda_function_name}.
	LambdaFunctionName *string `field:"optional" json:"lambdaFunctionName" yaml:"lambdaFunctionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/pinpoint_app#mode PinpointApp#mode}.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/pinpoint_app#web_url PinpointApp#web_url}.
	WebUrl *string `field:"optional" json:"webUrl" yaml:"webUrl"`
}

