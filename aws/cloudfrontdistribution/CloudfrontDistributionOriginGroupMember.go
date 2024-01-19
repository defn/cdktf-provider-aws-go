package cloudfrontdistribution


type CloudfrontDistributionOriginGroupMember struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/cloudfront_distribution#origin_id CloudfrontDistribution#origin_id}.
	OriginId *string `field:"required" json:"originId" yaml:"originId"`
}
