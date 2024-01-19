package ivsrecordingconfiguration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IvsRecordingConfigurationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// destination_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/ivs_recording_configuration#destination_configuration IvsRecordingConfiguration#destination_configuration}
	DestinationConfiguration *IvsRecordingConfigurationDestinationConfiguration `field:"required" json:"destinationConfiguration" yaml:"destinationConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/ivs_recording_configuration#id IvsRecordingConfiguration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/ivs_recording_configuration#name IvsRecordingConfiguration#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/ivs_recording_configuration#recording_reconnect_window_seconds IvsRecordingConfiguration#recording_reconnect_window_seconds}.
	RecordingReconnectWindowSeconds *float64 `field:"optional" json:"recordingReconnectWindowSeconds" yaml:"recordingReconnectWindowSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/ivs_recording_configuration#tags IvsRecordingConfiguration#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/ivs_recording_configuration#tags_all IvsRecordingConfiguration#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// thumbnail_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/ivs_recording_configuration#thumbnail_configuration IvsRecordingConfiguration#thumbnail_configuration}
	ThumbnailConfiguration *IvsRecordingConfigurationThumbnailConfiguration `field:"optional" json:"thumbnailConfiguration" yaml:"thumbnailConfiguration"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/ivs_recording_configuration#timeouts IvsRecordingConfiguration#timeouts}
	Timeouts *IvsRecordingConfigurationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

