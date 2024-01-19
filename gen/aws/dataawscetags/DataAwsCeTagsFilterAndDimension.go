package dataawscetags


type DataAwsCeTagsFilterAndDimension struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/ce_tags#key DataAwsCeTags#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/ce_tags#match_options DataAwsCeTags#match_options}.
	MatchOptions *[]*string `field:"optional" json:"matchOptions" yaml:"matchOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/ce_tags#values DataAwsCeTags#values}.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

