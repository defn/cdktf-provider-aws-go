package medialivechannel


type MedialiveChannelInputAttachmentsInputSettingsAudioSelector struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/medialive_channel#name MedialiveChannel#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// selector_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/medialive_channel#selector_settings MedialiveChannel#selector_settings}
	SelectorSettings *MedialiveChannelInputAttachmentsInputSettingsAudioSelectorSelectorSettings `field:"optional" json:"selectorSettings" yaml:"selectorSettings"`
}

