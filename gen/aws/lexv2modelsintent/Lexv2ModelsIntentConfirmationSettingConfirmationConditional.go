package lexv2modelsintent


type Lexv2ModelsIntentConfirmationSettingConfirmationConditional struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/lexv2models_intent#active Lexv2ModelsIntent#active}.
	Active interface{} `field:"required" json:"active" yaml:"active"`
	// conditional_branch block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/lexv2models_intent#conditional_branch Lexv2ModelsIntent#conditional_branch}
	ConditionalBranch interface{} `field:"optional" json:"conditionalBranch" yaml:"conditionalBranch"`
	// default_branch block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/lexv2models_intent#default_branch Lexv2ModelsIntent#default_branch}
	DefaultBranch interface{} `field:"optional" json:"defaultBranch" yaml:"defaultBranch"`
}

