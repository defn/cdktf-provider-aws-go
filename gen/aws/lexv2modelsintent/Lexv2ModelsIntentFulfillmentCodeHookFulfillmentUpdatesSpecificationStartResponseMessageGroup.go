package lexv2modelsintent


type Lexv2ModelsIntentFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroup struct {
	// message block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/lexv2models_intent#message Lexv2ModelsIntent#message}
	Message interface{} `field:"optional" json:"message" yaml:"message"`
	// variation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/lexv2models_intent#variation Lexv2ModelsIntent#variation}
	Variation interface{} `field:"optional" json:"variation" yaml:"variation"`
}

