package vpcendpointservice


type VpcEndpointServiceTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/vpc_endpoint_service#create VpcEndpointService#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/vpc_endpoint_service#delete VpcEndpointService#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/vpc_endpoint_service#update VpcEndpointService#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

