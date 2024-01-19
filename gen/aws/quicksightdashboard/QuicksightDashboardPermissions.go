package quicksightdashboard


type QuicksightDashboardPermissions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/quicksight_dashboard#actions QuicksightDashboard#actions}.
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/quicksight_dashboard#principal QuicksightDashboard#principal}.
	Principal *string `field:"required" json:"principal" yaml:"principal"`
}

