package appmeshvirtualnode


type AppmeshVirtualNodeSpecLoggingAccessLogFileFormat struct {
	// json block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appmesh_virtual_node#json AppmeshVirtualNode#json}
	Json interface{} `field:"optional" json:"json" yaml:"json"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appmesh_virtual_node#text AppmeshVirtualNode#text}.
	Text *string `field:"optional" json:"text" yaml:"text"`
}

