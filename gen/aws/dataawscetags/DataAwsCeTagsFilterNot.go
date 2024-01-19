package dataawscetags


type DataAwsCeTagsFilterNot struct {
	// cost_category block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/ce_tags#cost_category DataAwsCeTags#cost_category}
	CostCategory *DataAwsCeTagsFilterNotCostCategory `field:"optional" json:"costCategory" yaml:"costCategory"`
	// dimension block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/ce_tags#dimension DataAwsCeTags#dimension}
	Dimension *DataAwsCeTagsFilterNotDimension `field:"optional" json:"dimension" yaml:"dimension"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/ce_tags#tags DataAwsCeTags#tags}
	Tags *DataAwsCeTagsFilterNotTags `field:"optional" json:"tags" yaml:"tags"`
}
