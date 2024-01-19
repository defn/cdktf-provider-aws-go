package quicksightdataset


type QuicksightDataSetLogicalTableMapDataTransformsRenameColumnOperation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/quicksight_data_set#column_name QuicksightDataSet#column_name}.
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/quicksight_data_set#new_column_name QuicksightDataSet#new_column_name}.
	NewColumnName *string `field:"required" json:"newColumnName" yaml:"newColumnName"`
}

