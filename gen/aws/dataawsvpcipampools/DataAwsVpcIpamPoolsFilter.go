package dataawsvpcipampools


type DataAwsVpcIpamPoolsFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/vpc_ipam_pools#name DataAwsVpcIpamPools#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/vpc_ipam_pools#values DataAwsVpcIpamPools#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

