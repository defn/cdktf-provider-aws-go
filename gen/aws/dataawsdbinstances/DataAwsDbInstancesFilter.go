package dataawsdbinstances


type DataAwsDbInstancesFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/db_instances#name DataAwsDbInstances#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/db_instances#values DataAwsDbInstances#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

