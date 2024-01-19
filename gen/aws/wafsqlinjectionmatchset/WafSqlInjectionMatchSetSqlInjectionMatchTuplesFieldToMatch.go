package wafsqlinjectionmatchset


type WafSqlInjectionMatchSetSqlInjectionMatchTuplesFieldToMatch struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/waf_sql_injection_match_set#type WafSqlInjectionMatchSet#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/waf_sql_injection_match_set#data WafSqlInjectionMatchSet#data}.
	Data *string `field:"optional" json:"data" yaml:"data"`
}

