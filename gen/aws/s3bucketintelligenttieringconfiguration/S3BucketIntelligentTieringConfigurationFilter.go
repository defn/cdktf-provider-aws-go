package s3bucketintelligenttieringconfiguration


type S3BucketIntelligentTieringConfigurationFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/s3_bucket_intelligent_tiering_configuration#prefix S3BucketIntelligentTieringConfiguration#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/s3_bucket_intelligent_tiering_configuration#tags S3BucketIntelligentTieringConfiguration#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

