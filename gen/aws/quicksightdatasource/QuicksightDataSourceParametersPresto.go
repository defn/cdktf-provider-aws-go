package quicksightdatasource


type QuicksightDataSourceParametersPresto struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/quicksight_data_source#catalog QuicksightDataSource#catalog}.
	Catalog *string `field:"required" json:"catalog" yaml:"catalog"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/quicksight_data_source#host QuicksightDataSource#host}.
	Host *string `field:"required" json:"host" yaml:"host"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/quicksight_data_source#port QuicksightDataSource#port}.
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

