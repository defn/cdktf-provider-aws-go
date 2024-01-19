package qldbledger


type QldbLedgerTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/qldb_ledger#create QldbLedger#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/qldb_ledger#delete QldbLedger#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

