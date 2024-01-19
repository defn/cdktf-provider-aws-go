package appsyncgraphqlapi

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppsyncGraphqlApiConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appsync_graphql_api#authentication_type AppsyncGraphqlApi#authentication_type}.
	AuthenticationType *string `field:"required" json:"authenticationType" yaml:"authenticationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appsync_graphql_api#name AppsyncGraphqlApi#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// additional_authentication_provider block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appsync_graphql_api#additional_authentication_provider AppsyncGraphqlApi#additional_authentication_provider}
	AdditionalAuthenticationProvider interface{} `field:"optional" json:"additionalAuthenticationProvider" yaml:"additionalAuthenticationProvider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appsync_graphql_api#id AppsyncGraphqlApi#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// lambda_authorizer_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appsync_graphql_api#lambda_authorizer_config AppsyncGraphqlApi#lambda_authorizer_config}
	LambdaAuthorizerConfig *AppsyncGraphqlApiLambdaAuthorizerConfig `field:"optional" json:"lambdaAuthorizerConfig" yaml:"lambdaAuthorizerConfig"`
	// log_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appsync_graphql_api#log_config AppsyncGraphqlApi#log_config}
	LogConfig *AppsyncGraphqlApiLogConfig `field:"optional" json:"logConfig" yaml:"logConfig"`
	// openid_connect_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appsync_graphql_api#openid_connect_config AppsyncGraphqlApi#openid_connect_config}
	OpenidConnectConfig *AppsyncGraphqlApiOpenidConnectConfig `field:"optional" json:"openidConnectConfig" yaml:"openidConnectConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appsync_graphql_api#schema AppsyncGraphqlApi#schema}.
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appsync_graphql_api#tags AppsyncGraphqlApi#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appsync_graphql_api#tags_all AppsyncGraphqlApi#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// user_pool_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appsync_graphql_api#user_pool_config AppsyncGraphqlApi#user_pool_config}
	UserPoolConfig *AppsyncGraphqlApiUserPoolConfig `field:"optional" json:"userPoolConfig" yaml:"userPoolConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appsync_graphql_api#visibility AppsyncGraphqlApi#visibility}.
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appsync_graphql_api#xray_enabled AppsyncGraphqlApi#xray_enabled}.
	XrayEnabled interface{} `field:"optional" json:"xrayEnabled" yaml:"xrayEnabled"`
}

