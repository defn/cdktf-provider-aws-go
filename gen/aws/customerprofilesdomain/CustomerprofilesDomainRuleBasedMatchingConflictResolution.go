package customerprofilesdomain


type CustomerprofilesDomainRuleBasedMatchingConflictResolution struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/customerprofiles_domain#conflict_resolving_model CustomerprofilesDomain#conflict_resolving_model}.
	ConflictResolvingModel *string `field:"required" json:"conflictResolvingModel" yaml:"conflictResolvingModel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/customerprofiles_domain#source_name CustomerprofilesDomain#source_name}.
	SourceName *string `field:"optional" json:"sourceName" yaml:"sourceName"`
}

