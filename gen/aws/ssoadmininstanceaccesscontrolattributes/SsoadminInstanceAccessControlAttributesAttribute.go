package ssoadmininstanceaccesscontrolattributes


type SsoadminInstanceAccessControlAttributesAttribute struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/ssoadmin_instance_access_control_attributes#key SsoadminInstanceAccessControlAttributes#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/ssoadmin_instance_access_control_attributes#value SsoadminInstanceAccessControlAttributes#value}
	Value interface{} `field:"required" json:"value" yaml:"value"`
}

