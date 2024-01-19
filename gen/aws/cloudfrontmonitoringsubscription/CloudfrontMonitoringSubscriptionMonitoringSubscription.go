package cloudfrontmonitoringsubscription


type CloudfrontMonitoringSubscriptionMonitoringSubscription struct {
	// realtime_metrics_subscription_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/cloudfront_monitoring_subscription#realtime_metrics_subscription_config CloudfrontMonitoringSubscription#realtime_metrics_subscription_config}
	RealtimeMetricsSubscriptionConfig *CloudfrontMonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfig `field:"required" json:"realtimeMetricsSubscriptionConfig" yaml:"realtimeMetricsSubscriptionConfig"`
}

