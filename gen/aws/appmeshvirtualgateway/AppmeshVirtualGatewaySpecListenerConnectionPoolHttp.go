package appmeshvirtualgateway


type AppmeshVirtualGatewaySpecListenerConnectionPoolHttp struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appmesh_virtual_gateway#max_connections AppmeshVirtualGateway#max_connections}.
	MaxConnections *float64 `field:"required" json:"maxConnections" yaml:"maxConnections"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appmesh_virtual_gateway#max_pending_requests AppmeshVirtualGateway#max_pending_requests}.
	MaxPendingRequests *float64 `field:"optional" json:"maxPendingRequests" yaml:"maxPendingRequests"`
}

