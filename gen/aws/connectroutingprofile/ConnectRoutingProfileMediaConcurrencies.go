package connectroutingprofile


type ConnectRoutingProfileMediaConcurrencies struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/connect_routing_profile#channel ConnectRoutingProfile#channel}.
	Channel *string `field:"required" json:"channel" yaml:"channel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/connect_routing_profile#concurrency ConnectRoutingProfile#concurrency}.
	Concurrency *float64 `field:"required" json:"concurrency" yaml:"concurrency"`
}

