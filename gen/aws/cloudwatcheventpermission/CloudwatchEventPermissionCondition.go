package cloudwatcheventpermission


type CloudwatchEventPermissionCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/cloudwatch_event_permission#key CloudwatchEventPermission#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/cloudwatch_event_permission#type CloudwatchEventPermission#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/cloudwatch_event_permission#value CloudwatchEventPermission#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

