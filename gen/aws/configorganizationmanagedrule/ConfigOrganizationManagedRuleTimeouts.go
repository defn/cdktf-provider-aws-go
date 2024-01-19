package configorganizationmanagedrule


type ConfigOrganizationManagedRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/config_organization_managed_rule#create ConfigOrganizationManagedRule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/config_organization_managed_rule#delete ConfigOrganizationManagedRule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/config_organization_managed_rule#update ConfigOrganizationManagedRule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

