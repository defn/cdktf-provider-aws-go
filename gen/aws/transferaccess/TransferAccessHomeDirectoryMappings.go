package transferaccess


type TransferAccessHomeDirectoryMappings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/transfer_access#entry TransferAccess#entry}.
	Entry *string `field:"required" json:"entry" yaml:"entry"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/transfer_access#target TransferAccess#target}.
	Target *string `field:"required" json:"target" yaml:"target"`
}

