package s3bucket


type S3BucketReplicationConfigurationRulesSourceSelectionCriteria struct {
	// sse_kms_encrypted_objects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/s3_bucket#sse_kms_encrypted_objects S3Bucket#sse_kms_encrypted_objects}
	SseKmsEncryptedObjects *S3BucketReplicationConfigurationRulesSourceSelectionCriteriaSseKmsEncryptedObjects `field:"optional" json:"sseKmsEncryptedObjects" yaml:"sseKmsEncryptedObjects"`
}

