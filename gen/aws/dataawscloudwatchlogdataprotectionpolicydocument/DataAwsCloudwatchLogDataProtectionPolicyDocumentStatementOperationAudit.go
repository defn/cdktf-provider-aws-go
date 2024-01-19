package dataawscloudwatchlogdataprotectionpolicydocument


type DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationAudit struct {
	// findings_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/cloudwatch_log_data_protection_policy_document#findings_destination DataAwsCloudwatchLogDataProtectionPolicyDocument#findings_destination}
	FindingsDestination *DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestination `field:"required" json:"findingsDestination" yaml:"findingsDestination"`
}

