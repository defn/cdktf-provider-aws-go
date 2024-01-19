package comprehenddocumentclassifier


type ComprehendDocumentClassifierVpcConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/comprehend_document_classifier#security_group_ids ComprehendDocumentClassifier#security_group_ids}.
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/comprehend_document_classifier#subnets ComprehendDocumentClassifier#subnets}.
	Subnets *[]*string `field:"required" json:"subnets" yaml:"subnets"`
}

