package alblistenerrule


type AlbListenerRuleConditionHttpHeader struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/alb_listener_rule#http_header_name AlbListenerRule#http_header_name}.
	HttpHeaderName *string `field:"required" json:"httpHeaderName" yaml:"httpHeaderName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/alb_listener_rule#values AlbListenerRule#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

