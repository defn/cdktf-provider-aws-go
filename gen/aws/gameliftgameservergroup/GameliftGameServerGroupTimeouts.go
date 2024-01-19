package gameliftgameservergroup


type GameliftGameServerGroupTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/gamelift_game_server_group#create GameliftGameServerGroup#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/gamelift_game_server_group#delete GameliftGameServerGroup#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

