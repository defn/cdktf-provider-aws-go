package datasynclocationfsxontapfilesystem


type DatasyncLocationFsxOntapFileSystemProtocol struct {
	// nfs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/datasync_location_fsx_ontap_file_system#nfs DatasyncLocationFsxOntapFileSystem#nfs}
	Nfs *DatasyncLocationFsxOntapFileSystemProtocolNfs `field:"optional" json:"nfs" yaml:"nfs"`
	// smb block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/datasync_location_fsx_ontap_file_system#smb DatasyncLocationFsxOntapFileSystem#smb}
	Smb *DatasyncLocationFsxOntapFileSystemProtocolSmb `field:"optional" json:"smb" yaml:"smb"`
}
