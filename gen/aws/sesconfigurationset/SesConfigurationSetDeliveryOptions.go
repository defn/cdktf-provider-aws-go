package sesconfigurationset


type SesConfigurationSetDeliveryOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/ses_configuration_set#tls_policy SesConfigurationSet#tls_policy}.
	TlsPolicy *string `field:"optional" json:"tlsPolicy" yaml:"tlsPolicy"`
}

