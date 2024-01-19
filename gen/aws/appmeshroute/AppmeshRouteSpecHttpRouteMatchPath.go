package appmeshroute


type AppmeshRouteSpecHttpRouteMatchPath struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appmesh_route#exact AppmeshRoute#exact}.
	Exact *string `field:"optional" json:"exact" yaml:"exact"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appmesh_route#regex AppmeshRoute#regex}.
	Regex *string `field:"optional" json:"regex" yaml:"regex"`
}

