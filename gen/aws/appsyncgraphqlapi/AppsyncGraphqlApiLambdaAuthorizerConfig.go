package appsyncgraphqlapi


type AppsyncGraphqlApiLambdaAuthorizerConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appsync_graphql_api#authorizer_uri AppsyncGraphqlApi#authorizer_uri}.
	AuthorizerUri *string `field:"required" json:"authorizerUri" yaml:"authorizerUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appsync_graphql_api#authorizer_result_ttl_in_seconds AppsyncGraphqlApi#authorizer_result_ttl_in_seconds}.
	AuthorizerResultTtlInSeconds *float64 `field:"optional" json:"authorizerResultTtlInSeconds" yaml:"authorizerResultTtlInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appsync_graphql_api#identity_validation_expression AppsyncGraphqlApi#identity_validation_expression}.
	IdentityValidationExpression *string `field:"optional" json:"identityValidationExpression" yaml:"identityValidationExpression"`
}

