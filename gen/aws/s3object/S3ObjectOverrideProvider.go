package s3object


type S3ObjectOverrideProvider struct {
	// default_tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/s3_object#default_tags S3Object#default_tags}
	DefaultTags *S3ObjectOverrideProviderDefaultTags `field:"optional" json:"defaultTags" yaml:"defaultTags"`
}

