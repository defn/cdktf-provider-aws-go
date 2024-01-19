package cloudfrontdistribution


type CloudfrontDistributionOriginCustomHeader struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/cloudfront_distribution#name CloudfrontDistribution#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/cloudfront_distribution#value CloudfrontDistribution#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}
