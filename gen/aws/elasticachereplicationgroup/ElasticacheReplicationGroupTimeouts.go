package elasticachereplicationgroup


type ElasticacheReplicationGroupTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/elasticache_replication_group#create ElasticacheReplicationGroup#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/elasticache_replication_group#delete ElasticacheReplicationGroup#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/elasticache_replication_group#update ElasticacheReplicationGroup#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

