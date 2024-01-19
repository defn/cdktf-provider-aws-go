package kinesisfirehosedeliverystream


type KinesisFirehoseDeliveryStreamRedshiftConfigurationProcessingConfigurationProcessorsParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/kinesis_firehose_delivery_stream#parameter_name KinesisFirehoseDeliveryStream#parameter_name}.
	ParameterName *string `field:"required" json:"parameterName" yaml:"parameterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/kinesis_firehose_delivery_stream#parameter_value KinesisFirehoseDeliveryStream#parameter_value}.
	ParameterValue *string `field:"required" json:"parameterValue" yaml:"parameterValue"`
}

