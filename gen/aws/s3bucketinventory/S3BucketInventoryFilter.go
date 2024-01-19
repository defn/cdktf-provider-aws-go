package s3bucketinventory


type S3BucketInventoryFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/s3_bucket_inventory#prefix S3BucketInventory#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

