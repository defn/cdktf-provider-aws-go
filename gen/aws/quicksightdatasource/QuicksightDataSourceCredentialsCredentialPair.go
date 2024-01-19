package quicksightdatasource


type QuicksightDataSourceCredentialsCredentialPair struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/quicksight_data_source#password QuicksightDataSource#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/quicksight_data_source#username QuicksightDataSource#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
}

