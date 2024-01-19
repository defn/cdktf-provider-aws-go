package docdbclusterparametergroup


type DocdbClusterParameterGroupParameter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/docdb_cluster_parameter_group#name DocdbClusterParameterGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/docdb_cluster_parameter_group#value DocdbClusterParameterGroup#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/docdb_cluster_parameter_group#apply_method DocdbClusterParameterGroup#apply_method}.
	ApplyMethod *string `field:"optional" json:"applyMethod" yaml:"applyMethod"`
}

