package wafwebacl


type WafWebAclLoggingConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/waf_web_acl#log_destination WafWebAcl#log_destination}.
	LogDestination *string `field:"required" json:"logDestination" yaml:"logDestination"`
	// redacted_fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/waf_web_acl#redacted_fields WafWebAcl#redacted_fields}
	RedactedFields *WafWebAclLoggingConfigurationRedactedFields `field:"optional" json:"redactedFields" yaml:"redactedFields"`
}

