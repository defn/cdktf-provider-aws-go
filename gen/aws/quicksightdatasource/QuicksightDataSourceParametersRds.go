package quicksightdatasource


type QuicksightDataSourceParametersRds struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/quicksight_data_source#database QuicksightDataSource#database}.
	Database *string `field:"required" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/quicksight_data_source#instance_id QuicksightDataSource#instance_id}.
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
}

