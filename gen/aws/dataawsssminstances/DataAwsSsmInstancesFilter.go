package dataawsssminstances


type DataAwsSsmInstancesFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/ssm_instances#name DataAwsSsmInstances#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/ssm_instances#values DataAwsSsmInstances#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

