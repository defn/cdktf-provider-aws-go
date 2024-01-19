package sagemakerdataqualityjobdefinition


type SagemakerDataQualityJobDefinitionDataQualityJobOutputConfig struct {
	// monitoring_outputs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/sagemaker_data_quality_job_definition#monitoring_outputs SagemakerDataQualityJobDefinition#monitoring_outputs}
	MonitoringOutputs *SagemakerDataQualityJobDefinitionDataQualityJobOutputConfigMonitoringOutputs `field:"required" json:"monitoringOutputs" yaml:"monitoringOutputs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/sagemaker_data_quality_job_definition#kms_key_id SagemakerDataQualityJobDefinition#kms_key_id}.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

