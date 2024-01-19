package autoscalinggroup


type AutoscalingGroupMixedInstancesPolicyLaunchTemplate struct {
	// launch_template_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/autoscaling_group#launch_template_specification AutoscalingGroup#launch_template_specification}
	LaunchTemplateSpecification *AutoscalingGroupMixedInstancesPolicyLaunchTemplateLaunchTemplateSpecification `field:"required" json:"launchTemplateSpecification" yaml:"launchTemplateSpecification"`
	// override block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/autoscaling_group#override AutoscalingGroup#override}
	Override interface{} `field:"optional" json:"override" yaml:"override"`
}

