package dataawsinstance


type DataAwsInstanceFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/instance#name DataAwsInstance#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/instance#values DataAwsInstance#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

