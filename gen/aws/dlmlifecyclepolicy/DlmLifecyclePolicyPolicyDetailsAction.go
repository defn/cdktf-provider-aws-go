package dlmlifecyclepolicy


type DlmLifecyclePolicyPolicyDetailsAction struct {
	// cross_region_copy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/dlm_lifecycle_policy#cross_region_copy DlmLifecyclePolicy#cross_region_copy}
	CrossRegionCopy interface{} `field:"required" json:"crossRegionCopy" yaml:"crossRegionCopy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/dlm_lifecycle_policy#name DlmLifecyclePolicy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

