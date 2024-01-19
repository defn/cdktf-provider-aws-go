package quicksighttemplate


type QuicksightTemplatePermissions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/quicksight_template#actions QuicksightTemplate#actions}.
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/quicksight_template#principal QuicksightTemplate#principal}.
	Principal *string `field:"required" json:"principal" yaml:"principal"`
}
