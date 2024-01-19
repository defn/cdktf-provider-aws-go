package s3bucket


type S3BucketReplicationConfigurationRulesDestinationMetrics struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/s3_bucket#minutes S3Bucket#minutes}.
	Minutes *float64 `field:"optional" json:"minutes" yaml:"minutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/s3_bucket#status S3Bucket#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

