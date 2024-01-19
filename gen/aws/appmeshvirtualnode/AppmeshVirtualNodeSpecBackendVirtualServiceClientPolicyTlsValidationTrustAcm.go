package appmeshvirtualnode


type AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidationTrustAcm struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appmesh_virtual_node#certificate_authority_arns AppmeshVirtualNode#certificate_authority_arns}.
	CertificateAuthorityArns *[]*string `field:"required" json:"certificateAuthorityArns" yaml:"certificateAuthorityArns"`
}

