package appsyncgraphqlapi


type AppsyncGraphqlApiOpenidConnectConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appsync_graphql_api#issuer AppsyncGraphqlApi#issuer}.
	Issuer *string `field:"required" json:"issuer" yaml:"issuer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appsync_graphql_api#auth_ttl AppsyncGraphqlApi#auth_ttl}.
	AuthTtl *float64 `field:"optional" json:"authTtl" yaml:"authTtl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appsync_graphql_api#client_id AppsyncGraphqlApi#client_id}.
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appsync_graphql_api#iat_ttl AppsyncGraphqlApi#iat_ttl}.
	IatTtl *float64 `field:"optional" json:"iatTtl" yaml:"iatTtl"`
}

