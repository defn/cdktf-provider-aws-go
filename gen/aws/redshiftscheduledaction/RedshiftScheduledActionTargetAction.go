package redshiftscheduledaction


type RedshiftScheduledActionTargetAction struct {
	// pause_cluster block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/redshift_scheduled_action#pause_cluster RedshiftScheduledAction#pause_cluster}
	PauseCluster *RedshiftScheduledActionTargetActionPauseCluster `field:"optional" json:"pauseCluster" yaml:"pauseCluster"`
	// resize_cluster block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/redshift_scheduled_action#resize_cluster RedshiftScheduledAction#resize_cluster}
	ResizeCluster *RedshiftScheduledActionTargetActionResizeCluster `field:"optional" json:"resizeCluster" yaml:"resizeCluster"`
	// resume_cluster block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/redshift_scheduled_action#resume_cluster RedshiftScheduledAction#resume_cluster}
	ResumeCluster *RedshiftScheduledActionTargetActionResumeCluster `field:"optional" json:"resumeCluster" yaml:"resumeCluster"`
}

