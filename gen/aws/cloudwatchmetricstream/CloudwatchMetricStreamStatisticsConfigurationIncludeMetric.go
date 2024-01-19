package cloudwatchmetricstream


type CloudwatchMetricStreamStatisticsConfigurationIncludeMetric struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/cloudwatch_metric_stream#metric_name CloudwatchMetricStream#metric_name}.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/cloudwatch_metric_stream#namespace CloudwatchMetricStream#namespace}.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
}

