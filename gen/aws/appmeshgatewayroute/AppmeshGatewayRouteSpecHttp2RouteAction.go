package appmeshgatewayroute


type AppmeshGatewayRouteSpecHttp2RouteAction struct {
	// target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appmesh_gateway_route#target AppmeshGatewayRoute#target}
	Target *AppmeshGatewayRouteSpecHttp2RouteActionTarget `field:"required" json:"target" yaml:"target"`
	// rewrite block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appmesh_gateway_route#rewrite AppmeshGatewayRoute#rewrite}
	Rewrite *AppmeshGatewayRouteSpecHttp2RouteActionRewrite `field:"optional" json:"rewrite" yaml:"rewrite"`
}

