package kinesisfirehosedeliverystream


type KinesisFirehoseDeliveryStreamOpensearchConfigurationS3Configuration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/kinesis_firehose_delivery_stream#bucket_arn KinesisFirehoseDeliveryStream#bucket_arn}.
	BucketArn *string `field:"required" json:"bucketArn" yaml:"bucketArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/kinesis_firehose_delivery_stream#role_arn KinesisFirehoseDeliveryStream#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/kinesis_firehose_delivery_stream#buffering_interval KinesisFirehoseDeliveryStream#buffering_interval}.
	BufferingInterval *float64 `field:"optional" json:"bufferingInterval" yaml:"bufferingInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/kinesis_firehose_delivery_stream#buffering_size KinesisFirehoseDeliveryStream#buffering_size}.
	BufferingSize *float64 `field:"optional" json:"bufferingSize" yaml:"bufferingSize"`
	// cloudwatch_logging_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/kinesis_firehose_delivery_stream#cloudwatch_logging_options KinesisFirehoseDeliveryStream#cloudwatch_logging_options}
	CloudwatchLoggingOptions *KinesisFirehoseDeliveryStreamOpensearchConfigurationS3ConfigurationCloudwatchLoggingOptions `field:"optional" json:"cloudwatchLoggingOptions" yaml:"cloudwatchLoggingOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/kinesis_firehose_delivery_stream#compression_format KinesisFirehoseDeliveryStream#compression_format}.
	CompressionFormat *string `field:"optional" json:"compressionFormat" yaml:"compressionFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/kinesis_firehose_delivery_stream#error_output_prefix KinesisFirehoseDeliveryStream#error_output_prefix}.
	ErrorOutputPrefix *string `field:"optional" json:"errorOutputPrefix" yaml:"errorOutputPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/kinesis_firehose_delivery_stream#kms_key_arn KinesisFirehoseDeliveryStream#kms_key_arn}.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/kinesis_firehose_delivery_stream#prefix KinesisFirehoseDeliveryStream#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

