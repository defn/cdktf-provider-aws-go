package kmsgrant


type KmsGrantConstraints struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/kms_grant#encryption_context_equals KmsGrant#encryption_context_equals}.
	EncryptionContextEquals *map[string]*string `field:"optional" json:"encryptionContextEquals" yaml:"encryptionContextEquals"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/kms_grant#encryption_context_subset KmsGrant#encryption_context_subset}.
	EncryptionContextSubset *map[string]*string `field:"optional" json:"encryptionContextSubset" yaml:"encryptionContextSubset"`
}

