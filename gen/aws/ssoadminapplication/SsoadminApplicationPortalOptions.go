package ssoadminapplication


type SsoadminApplicationPortalOptions struct {
	// sign_in_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/ssoadmin_application#sign_in_options SsoadminApplication#sign_in_options}
	SignInOptions interface{} `field:"optional" json:"signInOptions" yaml:"signInOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/ssoadmin_application#visibility SsoadminApplication#visibility}.
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

