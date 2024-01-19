package dataawsvpcs


type DataAwsVpcsFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/vpcs#name DataAwsVpcs#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/vpcs#values DataAwsVpcs#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

