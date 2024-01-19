package kinesisfirehosedeliverystream


type KinesisFirehoseDeliveryStreamSplunkConfigurationProcessingConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/kinesis_firehose_delivery_stream#enabled KinesisFirehoseDeliveryStream#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// processors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/kinesis_firehose_delivery_stream#processors KinesisFirehoseDeliveryStream#processors}
	Processors interface{} `field:"optional" json:"processors" yaml:"processors"`
}

