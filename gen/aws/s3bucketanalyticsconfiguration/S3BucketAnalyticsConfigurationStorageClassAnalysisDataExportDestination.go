package s3bucketanalyticsconfiguration


type S3BucketAnalyticsConfigurationStorageClassAnalysisDataExportDestination struct {
	// s3_bucket_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/s3_bucket_analytics_configuration#s3_bucket_destination S3BucketAnalyticsConfiguration#s3_bucket_destination}
	S3BucketDestination *S3BucketAnalyticsConfigurationStorageClassAnalysisDataExportDestinationS3BucketDestination `field:"required" json:"s3BucketDestination" yaml:"s3BucketDestination"`
}

