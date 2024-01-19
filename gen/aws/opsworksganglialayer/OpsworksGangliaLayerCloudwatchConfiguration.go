package opsworksganglialayer


type OpsworksGangliaLayerCloudwatchConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/opsworks_ganglia_layer#enabled OpsworksGangliaLayer#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// log_streams block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/opsworks_ganglia_layer#log_streams OpsworksGangliaLayer#log_streams}
	LogStreams interface{} `field:"optional" json:"logStreams" yaml:"logStreams"`
}

