package efsmounttarget


type EfsMountTargetTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/efs_mount_target#create EfsMountTarget#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/efs_mount_target#delete EfsMountTarget#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

