package cloudfrontcontinuousdeploymentpolicy


type CloudfrontContinuousDeploymentPolicyStagingDistributionDnsNames struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/cloudfront_continuous_deployment_policy#quantity CloudfrontContinuousDeploymentPolicy#quantity}.
	Quantity *float64 `field:"required" json:"quantity" yaml:"quantity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/cloudfront_continuous_deployment_policy#items CloudfrontContinuousDeploymentPolicy#items}.
	Items *[]*string `field:"optional" json:"items" yaml:"items"`
}

