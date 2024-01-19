package emrblockpublicaccessconfiguration


type EmrBlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRange struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/emr_block_public_access_configuration#max_range EmrBlockPublicAccessConfiguration#max_range}.
	MaxRange *float64 `field:"required" json:"maxRange" yaml:"maxRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/emr_block_public_access_configuration#min_range EmrBlockPublicAccessConfiguration#min_range}.
	MinRange *float64 `field:"required" json:"minRange" yaml:"minRange"`
}

