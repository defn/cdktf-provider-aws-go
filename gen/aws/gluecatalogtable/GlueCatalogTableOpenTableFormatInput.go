package gluecatalogtable


type GlueCatalogTableOpenTableFormatInput struct {
	// iceberg_input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/glue_catalog_table#iceberg_input GlueCatalogTable#iceberg_input}
	IcebergInput *GlueCatalogTableOpenTableFormatInputIcebergInput `field:"required" json:"icebergInput" yaml:"icebergInput"`
}

