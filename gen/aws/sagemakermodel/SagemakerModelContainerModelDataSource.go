package sagemakermodel


type SagemakerModelContainerModelDataSource struct {
	// s3_data_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/sagemaker_model#s3_data_source SagemakerModel#s3_data_source}
	S3DataSource interface{} `field:"required" json:"s3DataSource" yaml:"s3DataSource"`
}

