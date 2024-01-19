package wafv2rulegroup


type Wafv2RuleGroupVisibilityConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/wafv2_rule_group#cloudwatch_metrics_enabled Wafv2RuleGroup#cloudwatch_metrics_enabled}.
	CloudwatchMetricsEnabled interface{} `field:"required" json:"cloudwatchMetricsEnabled" yaml:"cloudwatchMetricsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/wafv2_rule_group#metric_name Wafv2RuleGroup#metric_name}.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/wafv2_rule_group#sampled_requests_enabled Wafv2RuleGroup#sampled_requests_enabled}.
	SampledRequestsEnabled interface{} `field:"required" json:"sampledRequestsEnabled" yaml:"sampledRequestsEnabled"`
}

