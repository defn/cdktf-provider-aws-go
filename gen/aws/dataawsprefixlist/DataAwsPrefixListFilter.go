package dataawsprefixlist


type DataAwsPrefixListFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/prefix_list#name DataAwsPrefixList#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/prefix_list#values DataAwsPrefixList#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

