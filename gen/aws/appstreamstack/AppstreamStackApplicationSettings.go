package appstreamstack


type AppstreamStackApplicationSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appstream_stack#enabled AppstreamStack#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appstream_stack#settings_group AppstreamStack#settings_group}.
	SettingsGroup *string `field:"optional" json:"settingsGroup" yaml:"settingsGroup"`
}

