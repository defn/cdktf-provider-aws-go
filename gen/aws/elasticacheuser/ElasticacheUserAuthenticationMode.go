package elasticacheuser


type ElasticacheUserAuthenticationMode struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/elasticache_user#type ElasticacheUser#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/elasticache_user#passwords ElasticacheUser#passwords}.
	Passwords *[]*string `field:"optional" json:"passwords" yaml:"passwords"`
}
