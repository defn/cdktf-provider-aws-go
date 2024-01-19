package appmeshvirtualnode


type AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsCertificate struct {
	// file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appmesh_virtual_node#file AppmeshVirtualNode#file}
	File *AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsCertificateFile `field:"optional" json:"file" yaml:"file"`
	// sds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appmesh_virtual_node#sds AppmeshVirtualNode#sds}
	Sds *AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsCertificateSds `field:"optional" json:"sds" yaml:"sds"`
}

