package cloudwatchmetricstream


type CloudwatchMetricStreamTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/cloudwatch_metric_stream#create CloudwatchMetricStream#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/cloudwatch_metric_stream#delete CloudwatchMetricStream#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/cloudwatch_metric_stream#update CloudwatchMetricStream#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

