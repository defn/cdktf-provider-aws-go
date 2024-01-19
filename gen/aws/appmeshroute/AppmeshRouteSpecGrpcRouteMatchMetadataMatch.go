package appmeshroute


type AppmeshRouteSpecGrpcRouteMatchMetadataMatch struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appmesh_route#exact AppmeshRoute#exact}.
	Exact *string `field:"optional" json:"exact" yaml:"exact"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appmesh_route#prefix AppmeshRoute#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appmesh_route#range AppmeshRoute#range}
	Range *AppmeshRouteSpecGrpcRouteMatchMetadataMatchRange `field:"optional" json:"range" yaml:"range"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appmesh_route#regex AppmeshRoute#regex}.
	Regex *string `field:"optional" json:"regex" yaml:"regex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appmesh_route#suffix AppmeshRoute#suffix}.
	Suffix *string `field:"optional" json:"suffix" yaml:"suffix"`
}

