package sagemakeruserprofile


type SagemakerUserProfileUserSettingsSpaceStorageSettings struct {
	// default_ebs_storage_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/sagemaker_user_profile#default_ebs_storage_settings SagemakerUserProfile#default_ebs_storage_settings}
	DefaultEbsStorageSettings *SagemakerUserProfileUserSettingsSpaceStorageSettingsDefaultEbsStorageSettings `field:"optional" json:"defaultEbsStorageSettings" yaml:"defaultEbsStorageSettings"`
}

