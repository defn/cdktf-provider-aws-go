package wafregionalxssmatchset


type WafregionalXssMatchSetXssMatchTuple struct {
	// field_to_match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/wafregional_xss_match_set#field_to_match WafregionalXssMatchSet#field_to_match}
	FieldToMatch *WafregionalXssMatchSetXssMatchTupleFieldToMatch `field:"required" json:"fieldToMatch" yaml:"fieldToMatch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/wafregional_xss_match_set#text_transformation WafregionalXssMatchSet#text_transformation}.
	TextTransformation *string `field:"required" json:"textTransformation" yaml:"textTransformation"`
}

