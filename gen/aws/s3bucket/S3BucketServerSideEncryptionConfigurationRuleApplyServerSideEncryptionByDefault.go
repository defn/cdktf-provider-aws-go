package s3bucket


type S3BucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefault struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/s3_bucket#sse_algorithm S3Bucket#sse_algorithm}.
	SseAlgorithm *string `field:"required" json:"sseAlgorithm" yaml:"sseAlgorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/s3_bucket#kms_master_key_id S3Bucket#kms_master_key_id}.
	KmsMasterKeyId *string `field:"optional" json:"kmsMasterKeyId" yaml:"kmsMasterKeyId"`
}

