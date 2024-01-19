package quicksightdataset


type QuicksightDataSetColumnGroups struct {
	// geo_spatial_column_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/quicksight_data_set#geo_spatial_column_group QuicksightDataSet#geo_spatial_column_group}
	GeoSpatialColumnGroup *QuicksightDataSetColumnGroupsGeoSpatialColumnGroup `field:"optional" json:"geoSpatialColumnGroup" yaml:"geoSpatialColumnGroup"`
}

