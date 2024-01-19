package medialivechannel


type MedialiveChannelTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/medialive_channel#create MedialiveChannel#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/medialive_channel#delete MedialiveChannel#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/medialive_channel#update MedialiveChannel#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

