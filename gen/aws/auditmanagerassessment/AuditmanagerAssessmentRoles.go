package auditmanagerassessment


type AuditmanagerAssessmentRoles struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/auditmanager_assessment#role_arn AuditmanagerAssessment#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/auditmanager_assessment#role_type AuditmanagerAssessment#role_type}.
	RoleType *string `field:"required" json:"roleType" yaml:"roleType"`
}

