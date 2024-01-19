package imagebuilderimagerecipe


type ImagebuilderImageRecipeComponent struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/imagebuilder_image_recipe#component_arn ImagebuilderImageRecipe#component_arn}.
	ComponentArn *string `field:"required" json:"componentArn" yaml:"componentArn"`
	// parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/imagebuilder_image_recipe#parameter ImagebuilderImageRecipe#parameter}
	Parameter interface{} `field:"optional" json:"parameter" yaml:"parameter"`
}

