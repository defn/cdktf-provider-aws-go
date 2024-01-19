package timestreamwritetable


type TimestreamwriteTableSchema struct {
	// composite_partition_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/timestreamwrite_table#composite_partition_key TimestreamwriteTable#composite_partition_key}
	CompositePartitionKey *TimestreamwriteTableSchemaCompositePartitionKey `field:"optional" json:"compositePartitionKey" yaml:"compositePartitionKey"`
}
