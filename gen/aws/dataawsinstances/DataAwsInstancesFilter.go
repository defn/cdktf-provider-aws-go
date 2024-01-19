package dataawsinstances


type DataAwsInstancesFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/instances#name DataAwsInstances#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/instances#values DataAwsInstances#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

