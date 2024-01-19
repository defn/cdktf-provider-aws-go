package snapshotcreatevolumepermission


type SnapshotCreateVolumePermissionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/snapshot_create_volume_permission#create SnapshotCreateVolumePermission#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/snapshot_create_volume_permission#delete SnapshotCreateVolumePermission#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

