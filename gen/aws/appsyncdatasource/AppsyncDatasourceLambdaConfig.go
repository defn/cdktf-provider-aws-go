package appsyncdatasource


type AppsyncDatasourceLambdaConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appsync_datasource#function_arn AppsyncDatasource#function_arn}.
	FunctionArn *string `field:"required" json:"functionArn" yaml:"functionArn"`
}

