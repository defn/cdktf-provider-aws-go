package chimesdkmediapipelinesmediainsightspipelineconfiguration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ChimesdkmediapipelinesMediaInsightsPipelineConfigurationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// elements block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#elements ChimesdkmediapipelinesMediaInsightsPipelineConfiguration#elements}
	Elements interface{} `field:"required" json:"elements" yaml:"elements"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#name ChimesdkmediapipelinesMediaInsightsPipelineConfiguration#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#resource_access_role_arn ChimesdkmediapipelinesMediaInsightsPipelineConfiguration#resource_access_role_arn}.
	ResourceAccessRoleArn *string `field:"required" json:"resourceAccessRoleArn" yaml:"resourceAccessRoleArn"`
	// real_time_alert_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#real_time_alert_configuration ChimesdkmediapipelinesMediaInsightsPipelineConfiguration#real_time_alert_configuration}
	RealTimeAlertConfiguration *ChimesdkmediapipelinesMediaInsightsPipelineConfigurationRealTimeAlertConfiguration `field:"optional" json:"realTimeAlertConfiguration" yaml:"realTimeAlertConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#tags ChimesdkmediapipelinesMediaInsightsPipelineConfiguration#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#tags_all ChimesdkmediapipelinesMediaInsightsPipelineConfiguration#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#timeouts ChimesdkmediapipelinesMediaInsightsPipelineConfiguration#timeouts}
	Timeouts *ChimesdkmediapipelinesMediaInsightsPipelineConfigurationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

