package macie2classificationjob


type Macie2ClassificationJobS3JobDefinitionScopingIncludesAnd struct {
	// simple_scope_term block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/macie2_classification_job#simple_scope_term Macie2ClassificationJob#simple_scope_term}
	SimpleScopeTerm *Macie2ClassificationJobS3JobDefinitionScopingIncludesAndSimpleScopeTerm `field:"optional" json:"simpleScopeTerm" yaml:"simpleScopeTerm"`
	// tag_scope_term block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/macie2_classification_job#tag_scope_term Macie2ClassificationJob#tag_scope_term}
	TagScopeTerm *Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTerm `field:"optional" json:"tagScopeTerm" yaml:"tagScopeTerm"`
}

