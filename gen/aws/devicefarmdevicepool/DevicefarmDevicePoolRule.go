package devicefarmdevicepool


type DevicefarmDevicePoolRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/devicefarm_device_pool#attribute DevicefarmDevicePool#attribute}.
	Attribute *string `field:"optional" json:"attribute" yaml:"attribute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/devicefarm_device_pool#operator DevicefarmDevicePool#operator}.
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/devicefarm_device_pool#value DevicefarmDevicePool#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}
