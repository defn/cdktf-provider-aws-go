package lakeformationpermissions


type LakeformationPermissionsLfTag struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/lakeformation_permissions#key LakeformationPermissions#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/lakeformation_permissions#values LakeformationPermissions#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/lakeformation_permissions#catalog_id LakeformationPermissions#catalog_id}.
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
}

