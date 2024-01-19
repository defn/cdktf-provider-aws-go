package globalacceleratorlistener


type GlobalacceleratorListenerPortRange struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/globalaccelerator_listener#from_port GlobalacceleratorListener#from_port}.
	FromPort *float64 `field:"optional" json:"fromPort" yaml:"fromPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/globalaccelerator_listener#to_port GlobalacceleratorListener#to_port}.
	ToPort *float64 `field:"optional" json:"toPort" yaml:"toPort"`
}

