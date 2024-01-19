package qldbstream


type QldbStreamKinesisConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/qldb_stream#stream_arn QldbStream#stream_arn}.
	StreamArn *string `field:"required" json:"streamArn" yaml:"streamArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/qldb_stream#aggregation_enabled QldbStream#aggregation_enabled}.
	AggregationEnabled interface{} `field:"optional" json:"aggregationEnabled" yaml:"aggregationEnabled"`
}

