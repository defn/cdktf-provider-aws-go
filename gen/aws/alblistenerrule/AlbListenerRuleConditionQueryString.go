package alblistenerrule


type AlbListenerRuleConditionQueryString struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/alb_listener_rule#value AlbListenerRule#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/alb_listener_rule#key AlbListenerRule#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
}

