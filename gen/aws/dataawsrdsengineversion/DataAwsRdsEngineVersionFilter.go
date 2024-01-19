package dataawsrdsengineversion


type DataAwsRdsEngineVersionFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/rds_engine_version#name DataAwsRdsEngineVersion#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/rds_engine_version#values DataAwsRdsEngineVersion#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

