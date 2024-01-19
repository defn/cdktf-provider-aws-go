package appmeshroute


type AppmeshRouteSpecGrpcRouteTimeout struct {
	// idle block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appmesh_route#idle AppmeshRoute#idle}
	Idle *AppmeshRouteSpecGrpcRouteTimeoutIdle `field:"optional" json:"idle" yaml:"idle"`
	// per_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appmesh_route#per_request AppmeshRoute#per_request}
	PerRequest *AppmeshRouteSpecGrpcRouteTimeoutPerRequest `field:"optional" json:"perRequest" yaml:"perRequest"`
}

