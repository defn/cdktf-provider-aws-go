package schedulerschedule


type SchedulerScheduleTargetEcsParametersPlacementConstraints struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/scheduler_schedule#type SchedulerSchedule#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/scheduler_schedule#expression SchedulerSchedule#expression}.
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
}

