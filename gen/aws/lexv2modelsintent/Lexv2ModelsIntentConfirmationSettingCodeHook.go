package lexv2modelsintent


type Lexv2ModelsIntentConfirmationSettingCodeHook struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/lexv2models_intent#active Lexv2ModelsIntent#active}.
	Active interface{} `field:"required" json:"active" yaml:"active"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/lexv2models_intent#enable_code_hook_invocation Lexv2ModelsIntent#enable_code_hook_invocation}.
	EnableCodeHookInvocation interface{} `field:"required" json:"enableCodeHookInvocation" yaml:"enableCodeHookInvocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/lexv2models_intent#invocation_label Lexv2ModelsIntent#invocation_label}.
	InvocationLabel *string `field:"optional" json:"invocationLabel" yaml:"invocationLabel"`
	// post_code_hook_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/lexv2models_intent#post_code_hook_specification Lexv2ModelsIntent#post_code_hook_specification}
	PostCodeHookSpecification interface{} `field:"optional" json:"postCodeHookSpecification" yaml:"postCodeHookSpecification"`
}

