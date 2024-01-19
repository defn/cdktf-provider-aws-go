package eksidentityproviderconfig


type EksIdentityProviderConfigTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/eks_identity_provider_config#create EksIdentityProviderConfig#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/eks_identity_provider_config#delete EksIdentityProviderConfig#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

