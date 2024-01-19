package evidentlyfeature


type EvidentlyFeatureVariations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/evidently_feature#name EvidentlyFeature#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/evidently_feature#value EvidentlyFeature#value}
	Value *EvidentlyFeatureVariationsValue `field:"required" json:"value" yaml:"value"`
}

