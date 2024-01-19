package imagebuildercontainerrecipe


type ImagebuilderContainerRecipeInstanceConfiguration struct {
	// block_device_mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/imagebuilder_container_recipe#block_device_mapping ImagebuilderContainerRecipe#block_device_mapping}
	BlockDeviceMapping interface{} `field:"optional" json:"blockDeviceMapping" yaml:"blockDeviceMapping"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/imagebuilder_container_recipe#image ImagebuilderContainerRecipe#image}.
	Image *string `field:"optional" json:"image" yaml:"image"`
}

