package batchjobdefinition


type BatchJobDefinitionEksPropertiesPodPropertiesContainersVolumeMounts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/batch_job_definition#mount_path BatchJobDefinition#mount_path}.
	MountPath *string `field:"required" json:"mountPath" yaml:"mountPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/batch_job_definition#name BatchJobDefinition#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/batch_job_definition#read_only BatchJobDefinition#read_only}.
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
}

