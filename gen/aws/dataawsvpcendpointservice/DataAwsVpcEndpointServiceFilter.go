package dataawsvpcendpointservice


type DataAwsVpcEndpointServiceFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/vpc_endpoint_service#name DataAwsVpcEndpointService#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/vpc_endpoint_service#values DataAwsVpcEndpointService#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

