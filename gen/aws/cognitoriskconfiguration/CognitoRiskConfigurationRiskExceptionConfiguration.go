package cognitoriskconfiguration


type CognitoRiskConfigurationRiskExceptionConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/cognito_risk_configuration#blocked_ip_range_list CognitoRiskConfiguration#blocked_ip_range_list}.
	BlockedIpRangeList *[]*string `field:"optional" json:"blockedIpRangeList" yaml:"blockedIpRangeList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/cognito_risk_configuration#skipped_ip_range_list CognitoRiskConfiguration#skipped_ip_range_list}.
	SkippedIpRangeList *[]*string `field:"optional" json:"skippedIpRangeList" yaml:"skippedIpRangeList"`
}

