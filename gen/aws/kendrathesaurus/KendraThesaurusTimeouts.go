package kendrathesaurus


type KendraThesaurusTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/kendra_thesaurus#create KendraThesaurus#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/kendra_thesaurus#delete KendraThesaurus#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/kendra_thesaurus#update KendraThesaurus#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

