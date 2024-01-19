package dataawsidentitystoregroup


type DataAwsIdentitystoreGroupAlternateIdentifierUniqueAttribute struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/identitystore_group#attribute_path DataAwsIdentitystoreGroup#attribute_path}.
	AttributePath *string `field:"required" json:"attributePath" yaml:"attributePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/identitystore_group#attribute_value DataAwsIdentitystoreGroup#attribute_value}.
	AttributeValue *string `field:"required" json:"attributeValue" yaml:"attributeValue"`
}

