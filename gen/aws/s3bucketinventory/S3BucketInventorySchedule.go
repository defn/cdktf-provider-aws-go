package s3bucketinventory


type S3BucketInventorySchedule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/s3_bucket_inventory#frequency S3BucketInventory#frequency}.
	Frequency *string `field:"required" json:"frequency" yaml:"frequency"`
}

