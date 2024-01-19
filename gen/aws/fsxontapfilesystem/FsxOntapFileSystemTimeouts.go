package fsxontapfilesystem


type FsxOntapFileSystemTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/fsx_ontap_file_system#create FsxOntapFileSystem#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/fsx_ontap_file_system#delete FsxOntapFileSystem#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/fsx_ontap_file_system#update FsxOntapFileSystem#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

