package networkmanagertransitgatewaypeering


type NetworkmanagerTransitGatewayPeeringTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/networkmanager_transit_gateway_peering#create NetworkmanagerTransitGatewayPeering#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/networkmanager_transit_gateway_peering#delete NetworkmanagerTransitGatewayPeering#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

