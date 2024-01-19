package codebuildproject


type CodebuildProjectEnvironmentRegistryCredential struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/codebuild_project#credential CodebuildProject#credential}.
	Credential *string `field:"required" json:"credential" yaml:"credential"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/codebuild_project#credential_provider CodebuildProject#credential_provider}.
	CredentialProvider *string `field:"required" json:"credentialProvider" yaml:"credentialProvider"`
}

