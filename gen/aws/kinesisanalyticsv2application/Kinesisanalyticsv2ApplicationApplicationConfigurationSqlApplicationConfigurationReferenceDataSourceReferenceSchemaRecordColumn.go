package kinesisanalyticsv2application


type Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordColumn struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/kinesisanalyticsv2_application#name Kinesisanalyticsv2Application#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/kinesisanalyticsv2_application#sql_type Kinesisanalyticsv2Application#sql_type}.
	SqlType *string `field:"required" json:"sqlType" yaml:"sqlType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/kinesisanalyticsv2_application#mapping Kinesisanalyticsv2Application#mapping}.
	Mapping *string `field:"optional" json:"mapping" yaml:"mapping"`
}

