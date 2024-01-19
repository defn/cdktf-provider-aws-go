package cloudwatcheventtarget


type CloudwatchEventTargetEcsTargetPlacementConstraint struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/cloudwatch_event_target#type CloudwatchEventTarget#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/cloudwatch_event_target#expression CloudwatchEventTarget#expression}.
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
}

