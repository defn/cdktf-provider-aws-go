package gluecatalogtable


type GlueCatalogTablePartitionIndex struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/glue_catalog_table#index_name GlueCatalogTable#index_name}.
	IndexName *string `field:"required" json:"indexName" yaml:"indexName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/glue_catalog_table#keys GlueCatalogTable#keys}.
	Keys *[]*string `field:"required" json:"keys" yaml:"keys"`
}

