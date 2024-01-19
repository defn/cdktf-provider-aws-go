package dataawssecuritygroups


type DataAwsSecurityGroupsFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/security_groups#name DataAwsSecurityGroups#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/security_groups#values DataAwsSecurityGroups#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}
