package networkfirewallrulegroup


type NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRuleRuleDefinitionMatchAttributesDestinationPort struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/networkfirewall_rule_group#from_port NetworkfirewallRuleGroup#from_port}.
	FromPort *float64 `field:"required" json:"fromPort" yaml:"fromPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/networkfirewall_rule_group#to_port NetworkfirewallRuleGroup#to_port}.
	ToPort *float64 `field:"optional" json:"toPort" yaml:"toPort"`
}

