package cloudformationstacksetinstance


type CloudformationStackSetInstanceTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/cloudformation_stack_set_instance#create CloudformationStackSetInstance#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/cloudformation_stack_set_instance#delete CloudformationStackSetInstance#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/cloudformation_stack_set_instance#update CloudformationStackSetInstance#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

