package quicksightdataset


type QuicksightDataSetRefreshPropertiesRefreshConfigurationIncrementalRefresh struct {
	// lookback_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/quicksight_data_set#lookback_window QuicksightDataSet#lookback_window}
	LookbackWindow *QuicksightDataSetRefreshPropertiesRefreshConfigurationIncrementalRefreshLookbackWindow `field:"required" json:"lookbackWindow" yaml:"lookbackWindow"`
}

