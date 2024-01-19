package dataawsidentitystoreuser


type DataAwsIdentitystoreUserAlternateIdentifier struct {
	// external_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/identitystore_user#external_id DataAwsIdentitystoreUser#external_id}
	ExternalId *DataAwsIdentitystoreUserAlternateIdentifierExternalId `field:"optional" json:"externalId" yaml:"externalId"`
	// unique_attribute block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/identitystore_user#unique_attribute DataAwsIdentitystoreUser#unique_attribute}
	UniqueAttribute *DataAwsIdentitystoreUserAlternateIdentifierUniqueAttribute `field:"optional" json:"uniqueAttribute" yaml:"uniqueAttribute"`
}

