package imagebuilderimagerecipe


type ImagebuilderImageRecipeComponentParameter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/imagebuilder_image_recipe#name ImagebuilderImageRecipe#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/imagebuilder_image_recipe#value ImagebuilderImageRecipe#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

