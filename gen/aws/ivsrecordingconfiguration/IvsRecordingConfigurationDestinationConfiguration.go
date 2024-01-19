package ivsrecordingconfiguration


type IvsRecordingConfigurationDestinationConfiguration struct {
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/ivs_recording_configuration#s3 IvsRecordingConfiguration#s3}
	S3 *IvsRecordingConfigurationDestinationConfigurationS3 `field:"required" json:"s3" yaml:"s3"`
}

