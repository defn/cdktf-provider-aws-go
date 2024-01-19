package quicksightdashboard


type QuicksightDashboardSourceEntity struct {
	// source_template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/quicksight_dashboard#source_template QuicksightDashboard#source_template}
	SourceTemplate *QuicksightDashboardSourceEntitySourceTemplate `field:"optional" json:"sourceTemplate" yaml:"sourceTemplate"`
}

