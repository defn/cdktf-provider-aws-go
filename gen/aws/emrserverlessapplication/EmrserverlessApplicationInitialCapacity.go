package emrserverlessapplication


type EmrserverlessApplicationInitialCapacity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/emrserverless_application#initial_capacity_type EmrserverlessApplication#initial_capacity_type}.
	InitialCapacityType *string `field:"required" json:"initialCapacityType" yaml:"initialCapacityType"`
	// initial_capacity_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/emrserverless_application#initial_capacity_config EmrserverlessApplication#initial_capacity_config}
	InitialCapacityConfig *EmrserverlessApplicationInitialCapacityInitialCapacityConfig `field:"optional" json:"initialCapacityConfig" yaml:"initialCapacityConfig"`
}

