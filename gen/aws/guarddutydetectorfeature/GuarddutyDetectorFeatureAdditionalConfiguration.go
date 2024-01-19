package guarddutydetectorfeature


type GuarddutyDetectorFeatureAdditionalConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/guardduty_detector_feature#name GuarddutyDetectorFeature#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/guardduty_detector_feature#status GuarddutyDetectorFeature#status}.
	Status *string `field:"required" json:"status" yaml:"status"`
}

