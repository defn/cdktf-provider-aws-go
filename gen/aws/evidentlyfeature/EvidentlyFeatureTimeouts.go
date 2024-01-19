package evidentlyfeature


type EvidentlyFeatureTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/evidently_feature#create EvidentlyFeature#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/evidently_feature#delete EvidentlyFeature#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/evidently_feature#update EvidentlyFeature#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

