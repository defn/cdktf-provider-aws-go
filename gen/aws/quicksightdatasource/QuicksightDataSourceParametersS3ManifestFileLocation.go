package quicksightdatasource


type QuicksightDataSourceParametersS3ManifestFileLocation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/quicksight_data_source#bucket QuicksightDataSource#bucket}.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/quicksight_data_source#key QuicksightDataSource#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
}
