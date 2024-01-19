package kinesisfirehosedeliverystream


type KinesisFirehoseDeliveryStreamOpensearchConfigurationCloudwatchLoggingOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/kinesis_firehose_delivery_stream#enabled KinesisFirehoseDeliveryStream#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/kinesis_firehose_delivery_stream#log_group_name KinesisFirehoseDeliveryStream#log_group_name}.
	LogGroupName *string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/kinesis_firehose_delivery_stream#log_stream_name KinesisFirehoseDeliveryStream#log_stream_name}.
	LogStreamName *string `field:"optional" json:"logStreamName" yaml:"logStreamName"`
}

