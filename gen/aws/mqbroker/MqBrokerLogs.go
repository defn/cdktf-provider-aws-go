package mqbroker


type MqBrokerLogs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/mq_broker#audit MqBroker#audit}.
	Audit *string `field:"optional" json:"audit" yaml:"audit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/mq_broker#general MqBroker#general}.
	General interface{} `field:"optional" json:"general" yaml:"general"`
}

