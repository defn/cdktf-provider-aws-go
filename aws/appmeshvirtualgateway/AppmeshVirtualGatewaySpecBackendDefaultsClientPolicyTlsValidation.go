package appmeshvirtualgateway


type AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidation struct {
	// trust block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appmesh_virtual_gateway#trust AppmeshVirtualGateway#trust}
	Trust *AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrust `field:"required" json:"trust" yaml:"trust"`
	// subject_alternative_names block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appmesh_virtual_gateway#subject_alternative_names AppmeshVirtualGateway#subject_alternative_names}
	SubjectAlternativeNames *AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNames `field:"optional" json:"subjectAlternativeNames" yaml:"subjectAlternativeNames"`
}
