package memorydbparametergroup


type MemorydbParameterGroupParameter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/memorydb_parameter_group#name MemorydbParameterGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/memorydb_parameter_group#value MemorydbParameterGroup#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

