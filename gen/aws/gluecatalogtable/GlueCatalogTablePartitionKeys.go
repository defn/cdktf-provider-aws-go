package gluecatalogtable


type GlueCatalogTablePartitionKeys struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/glue_catalog_table#name GlueCatalogTable#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/glue_catalog_table#comment GlueCatalogTable#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/glue_catalog_table#type GlueCatalogTable#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

