package elasticsearchdomain


type ElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDuration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/elasticsearch_domain#unit ElasticsearchDomain#unit}.
	Unit *string `field:"required" json:"unit" yaml:"unit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/elasticsearch_domain#value ElasticsearchDomain#value}.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

