package vpclatticelistener


type VpclatticeListenerTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/vpclattice_listener#create VpclatticeListener#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/vpclattice_listener#delete VpclatticeListener#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/vpclattice_listener#update VpclatticeListener#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

