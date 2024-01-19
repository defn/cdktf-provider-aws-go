package lexv2modelsintent


type Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationTimeoutNextStepIntentSlot struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/lexv2models_intent#map_block_key Lexv2ModelsIntent#map_block_key}.
	MapBlockKey *string `field:"required" json:"mapBlockKey" yaml:"mapBlockKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/lexv2models_intent#shape Lexv2ModelsIntent#shape}.
	Shape *string `field:"optional" json:"shape" yaml:"shape"`
	// value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/lexv2models_intent#value Lexv2ModelsIntent#value}
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}

