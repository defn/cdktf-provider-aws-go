package quicksightdashboard


type QuicksightDashboardSourceEntitySourceTemplate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/quicksight_dashboard#arn QuicksightDashboard#arn}.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// data_set_references block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/quicksight_dashboard#data_set_references QuicksightDashboard#data_set_references}
	DataSetReferences interface{} `field:"required" json:"dataSetReferences" yaml:"dataSetReferences"`
}

