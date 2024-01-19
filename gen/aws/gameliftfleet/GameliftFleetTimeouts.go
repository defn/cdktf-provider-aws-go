package gameliftfleet


type GameliftFleetTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/gamelift_fleet#create GameliftFleet#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/gamelift_fleet#delete GameliftFleet#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

