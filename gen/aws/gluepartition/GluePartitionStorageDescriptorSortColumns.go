package gluepartition


type GluePartitionStorageDescriptorSortColumns struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/glue_partition#column GluePartition#column}.
	Column *string `field:"required" json:"column" yaml:"column"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/glue_partition#sort_order GluePartition#sort_order}.
	SortOrder *float64 `field:"required" json:"sortOrder" yaml:"sortOrder"`
}

