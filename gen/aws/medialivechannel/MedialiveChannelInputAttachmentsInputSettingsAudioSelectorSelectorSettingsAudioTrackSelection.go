package medialivechannel


type MedialiveChannelInputAttachmentsInputSettingsAudioSelectorSelectorSettingsAudioTrackSelection struct {
	// tracks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/medialive_channel#tracks MedialiveChannel#tracks}
	Tracks interface{} `field:"required" json:"tracks" yaml:"tracks"`
	// dolby_e_decode block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/medialive_channel#dolby_e_decode MedialiveChannel#dolby_e_decode}
	DolbyEDecode *MedialiveChannelInputAttachmentsInputSettingsAudioSelectorSelectorSettingsAudioTrackSelectionDolbyEDecode `field:"optional" json:"dolbyEDecode" yaml:"dolbyEDecode"`
}

