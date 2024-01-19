package s3directorybucket


type S3DirectoryBucketLocation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/s3_directory_bucket#name S3DirectoryBucket#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/s3_directory_bucket#type S3DirectoryBucket#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

