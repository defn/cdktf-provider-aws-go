package s3bucketreplicationconfiguration


type S3BucketReplicationConfigurationRuleDestinationReplicationTime struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/s3_bucket_replication_configuration#status S3BucketReplicationConfigurationA#status}.
	Status *string `field:"required" json:"status" yaml:"status"`
	// time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/s3_bucket_replication_configuration#time S3BucketReplicationConfigurationA#time}
	Time *S3BucketReplicationConfigurationRuleDestinationReplicationTimeTime `field:"required" json:"time" yaml:"time"`
}

