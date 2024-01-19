package networkmanagerlink


type NetworkmanagerLinkTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/networkmanager_link#create NetworkmanagerLink#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/networkmanager_link#delete NetworkmanagerLink#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/networkmanager_link#update NetworkmanagerLink#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
