package securitylakedatalake


type SecuritylakeDataLakeConfigurationLifecycleConfiguration struct {
	// expiration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/securitylake_data_lake#expiration SecuritylakeDataLake#expiration}
	Expiration interface{} `field:"optional" json:"expiration" yaml:"expiration"`
	// transition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/securitylake_data_lake#transition SecuritylakeDataLake#transition}
	Transition interface{} `field:"optional" json:"transition" yaml:"transition"`
}

