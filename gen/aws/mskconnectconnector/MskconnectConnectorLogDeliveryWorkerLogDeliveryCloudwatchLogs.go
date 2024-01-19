package mskconnectconnector


type MskconnectConnectorLogDeliveryWorkerLogDeliveryCloudwatchLogs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/mskconnect_connector#enabled MskconnectConnector#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/mskconnect_connector#log_group MskconnectConnector#log_group}.
	LogGroup *string `field:"optional" json:"logGroup" yaml:"logGroup"`
}
