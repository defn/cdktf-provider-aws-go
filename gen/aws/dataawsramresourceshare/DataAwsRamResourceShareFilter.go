package dataawsramresourceshare


type DataAwsRamResourceShareFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/ram_resource_share#name DataAwsRamResourceShare#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/ram_resource_share#values DataAwsRamResourceShare#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

