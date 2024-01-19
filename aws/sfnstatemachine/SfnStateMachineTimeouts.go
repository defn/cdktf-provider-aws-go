package sfnstatemachine


type SfnStateMachineTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/sfn_state_machine#create SfnStateMachine#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/sfn_state_machine#delete SfnStateMachine#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/sfn_state_machine#update SfnStateMachine#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
