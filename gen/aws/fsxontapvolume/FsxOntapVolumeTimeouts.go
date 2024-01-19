package fsxontapvolume


type FsxOntapVolumeTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/fsx_ontap_volume#create FsxOntapVolume#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/fsx_ontap_volume#delete FsxOntapVolume#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/fsx_ontap_volume#update FsxOntapVolume#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

