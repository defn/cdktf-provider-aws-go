package provider


type AwsProviderDefaultTags struct {
	// Resource tags to default across all resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs#tags AwsProvider#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

