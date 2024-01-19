package cecostcategory


type CeCostCategoryRuleRuleOrDimension struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/ce_cost_category#key CeCostCategory#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/ce_cost_category#match_options CeCostCategory#match_options}.
	MatchOptions *[]*string `field:"optional" json:"matchOptions" yaml:"matchOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/ce_cost_category#values CeCostCategory#values}.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

