package dataawseip


type DataAwsEipFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/eip#name DataAwsEip#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/eip#values DataAwsEip#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

