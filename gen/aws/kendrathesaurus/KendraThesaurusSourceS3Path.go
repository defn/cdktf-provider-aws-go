package kendrathesaurus


type KendraThesaurusSourceS3Path struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/kendra_thesaurus#bucket KendraThesaurus#bucket}.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/kendra_thesaurus#key KendraThesaurus#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
}

