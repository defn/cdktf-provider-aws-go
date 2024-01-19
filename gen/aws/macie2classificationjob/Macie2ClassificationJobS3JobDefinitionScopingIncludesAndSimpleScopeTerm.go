package macie2classificationjob


type Macie2ClassificationJobS3JobDefinitionScopingIncludesAndSimpleScopeTerm struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/macie2_classification_job#comparator Macie2ClassificationJob#comparator}.
	Comparator *string `field:"optional" json:"comparator" yaml:"comparator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/macie2_classification_job#key Macie2ClassificationJob#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/macie2_classification_job#values Macie2ClassificationJob#values}.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

