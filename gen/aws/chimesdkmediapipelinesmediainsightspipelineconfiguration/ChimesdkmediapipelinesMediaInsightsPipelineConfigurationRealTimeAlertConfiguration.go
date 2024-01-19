package chimesdkmediapipelinesmediainsightspipelineconfiguration


type ChimesdkmediapipelinesMediaInsightsPipelineConfigurationRealTimeAlertConfiguration struct {
	// rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#rules ChimesdkmediapipelinesMediaInsightsPipelineConfiguration#rules}
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#disabled ChimesdkmediapipelinesMediaInsightsPipelineConfiguration#disabled}.
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
}

