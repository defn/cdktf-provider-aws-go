package iottopicrule


type IotTopicRuleErrorActionSqs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/iot_topic_rule#queue_url IotTopicRule#queue_url}.
	QueueUrl *string `field:"required" json:"queueUrl" yaml:"queueUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/iot_topic_rule#role_arn IotTopicRule#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/iot_topic_rule#use_base64 IotTopicRule#use_base64}.
	UseBase64 interface{} `field:"required" json:"useBase64" yaml:"useBase64"`
}

