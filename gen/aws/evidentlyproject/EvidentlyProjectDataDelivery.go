package evidentlyproject


type EvidentlyProjectDataDelivery struct {
	// cloudwatch_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/evidently_project#cloudwatch_logs EvidentlyProject#cloudwatch_logs}
	CloudwatchLogs *EvidentlyProjectDataDeliveryCloudwatchLogs `field:"optional" json:"cloudwatchLogs" yaml:"cloudwatchLogs"`
	// s3_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/evidently_project#s3_destination EvidentlyProject#s3_destination}
	S3Destination *EvidentlyProjectDataDeliveryS3Destination `field:"optional" json:"s3Destination" yaml:"s3Destination"`
}

