package sesv2configurationset


type Sesv2ConfigurationSetSuppressionOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/sesv2_configuration_set#suppressed_reasons Sesv2ConfigurationSet#suppressed_reasons}.
	SuppressedReasons *[]*string `field:"optional" json:"suppressedReasons" yaml:"suppressedReasons"`
}

