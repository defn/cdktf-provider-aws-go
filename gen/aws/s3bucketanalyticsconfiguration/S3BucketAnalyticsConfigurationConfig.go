package s3bucketanalyticsconfiguration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type S3BucketAnalyticsConfigurationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/s3_bucket_analytics_configuration#bucket S3BucketAnalyticsConfiguration#bucket}.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/s3_bucket_analytics_configuration#name S3BucketAnalyticsConfiguration#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/s3_bucket_analytics_configuration#filter S3BucketAnalyticsConfiguration#filter}
	Filter *S3BucketAnalyticsConfigurationFilter `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/s3_bucket_analytics_configuration#id S3BucketAnalyticsConfiguration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// storage_class_analysis block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/s3_bucket_analytics_configuration#storage_class_analysis S3BucketAnalyticsConfiguration#storage_class_analysis}
	StorageClassAnalysis *S3BucketAnalyticsConfigurationStorageClassAnalysis `field:"optional" json:"storageClassAnalysis" yaml:"storageClassAnalysis"`
}

