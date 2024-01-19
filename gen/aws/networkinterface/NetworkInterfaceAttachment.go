package networkinterface


type NetworkInterfaceAttachment struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/network_interface#device_index NetworkInterface#device_index}.
	DeviceIndex *float64 `field:"required" json:"deviceIndex" yaml:"deviceIndex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/network_interface#instance NetworkInterface#instance}.
	Instance *string `field:"required" json:"instance" yaml:"instance"`
}

