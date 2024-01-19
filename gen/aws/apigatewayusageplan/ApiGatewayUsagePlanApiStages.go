package apigatewayusageplan


type ApiGatewayUsagePlanApiStages struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/api_gateway_usage_plan#api_id ApiGatewayUsagePlan#api_id}.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/api_gateway_usage_plan#stage ApiGatewayUsagePlan#stage}.
	Stage *string `field:"required" json:"stage" yaml:"stage"`
	// throttle block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/api_gateway_usage_plan#throttle ApiGatewayUsagePlan#throttle}
	Throttle interface{} `field:"optional" json:"throttle" yaml:"throttle"`
}

