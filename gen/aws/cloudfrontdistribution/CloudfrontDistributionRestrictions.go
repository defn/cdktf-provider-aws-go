package cloudfrontdistribution


type CloudfrontDistributionRestrictions struct {
	// geo_restriction block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/cloudfront_distribution#geo_restriction CloudfrontDistribution#geo_restriction}
	GeoRestriction *CloudfrontDistributionRestrictionsGeoRestriction `field:"required" json:"geoRestriction" yaml:"geoRestriction"`
}

