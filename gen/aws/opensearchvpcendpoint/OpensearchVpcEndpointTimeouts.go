package opensearchvpcendpoint


type OpensearchVpcEndpointTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/opensearch_vpc_endpoint#create OpensearchVpcEndpoint#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/opensearch_vpc_endpoint#delete OpensearchVpcEndpoint#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/opensearch_vpc_endpoint#update OpensearchVpcEndpoint#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

