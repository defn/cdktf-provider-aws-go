package eksaccessentry


type EksAccessEntryTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/eks_access_entry#create EksAccessEntry#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/eks_access_entry#delete EksAccessEntry#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

