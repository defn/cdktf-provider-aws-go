package dataawssecuritygroup


type DataAwsSecurityGroupFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/security_group#name DataAwsSecurityGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/security_group#values DataAwsSecurityGroup#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

