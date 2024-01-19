package pinpointapp


type PinpointAppQuietTime struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/pinpoint_app#end PinpointApp#end}.
	End *string `field:"optional" json:"end" yaml:"end"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/pinpoint_app#start PinpointApp#start}.
	Start *string `field:"optional" json:"start" yaml:"start"`
}

