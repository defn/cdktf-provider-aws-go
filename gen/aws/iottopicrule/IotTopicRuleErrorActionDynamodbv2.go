package iottopicrule


type IotTopicRuleErrorActionDynamodbv2 struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/iot_topic_rule#role_arn IotTopicRule#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// put_item block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/iot_topic_rule#put_item IotTopicRule#put_item}
	PutItem *IotTopicRuleErrorActionDynamodbv2PutItem `field:"optional" json:"putItem" yaml:"putItem"`
}

