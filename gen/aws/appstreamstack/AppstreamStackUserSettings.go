package appstreamstack


type AppstreamStackUserSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appstream_stack#action AppstreamStack#action}.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appstream_stack#permission AppstreamStack#permission}.
	Permission *string `field:"required" json:"permission" yaml:"permission"`
}

