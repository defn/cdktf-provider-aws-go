package autoscalinggroup


type AutoscalingGroupTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/autoscaling_group#delete AutoscalingGroup#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/autoscaling_group#update AutoscalingGroup#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

