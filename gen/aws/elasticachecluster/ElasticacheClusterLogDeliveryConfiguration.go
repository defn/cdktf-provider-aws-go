package elasticachecluster


type ElasticacheClusterLogDeliveryConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/elasticache_cluster#destination ElasticacheCluster#destination}.
	Destination *string `field:"required" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/elasticache_cluster#destination_type ElasticacheCluster#destination_type}.
	DestinationType *string `field:"required" json:"destinationType" yaml:"destinationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/elasticache_cluster#log_format ElasticacheCluster#log_format}.
	LogFormat *string `field:"required" json:"logFormat" yaml:"logFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/elasticache_cluster#log_type ElasticacheCluster#log_type}.
	LogType *string `field:"required" json:"logType" yaml:"logType"`
}

