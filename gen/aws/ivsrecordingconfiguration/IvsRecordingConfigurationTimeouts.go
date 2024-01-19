package ivsrecordingconfiguration


type IvsRecordingConfigurationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/ivs_recording_configuration#create IvsRecordingConfiguration#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/ivs_recording_configuration#delete IvsRecordingConfiguration#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}
