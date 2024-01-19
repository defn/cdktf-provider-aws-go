package launchtemplate


type LaunchTemplateInstanceMarketOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/launch_template#market_type LaunchTemplate#market_type}.
	MarketType *string `field:"optional" json:"marketType" yaml:"marketType"`
	// spot_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/launch_template#spot_options LaunchTemplate#spot_options}
	SpotOptions *LaunchTemplateInstanceMarketOptionsSpotOptions `field:"optional" json:"spotOptions" yaml:"spotOptions"`
}

