package ecrreplicationconfiguration


type EcrReplicationConfigurationReplicationConfigurationRule struct {
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/ecr_replication_configuration#destination EcrReplicationConfiguration#destination}
	Destination interface{} `field:"required" json:"destination" yaml:"destination"`
	// repository_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/ecr_replication_configuration#repository_filter EcrReplicationConfiguration#repository_filter}
	RepositoryFilter interface{} `field:"optional" json:"repositoryFilter" yaml:"repositoryFilter"`
}

