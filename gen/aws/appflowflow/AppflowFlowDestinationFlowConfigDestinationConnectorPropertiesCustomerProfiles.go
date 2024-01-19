package appflowflow


type AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesCustomerProfiles struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appflow_flow#domain_name AppflowFlow#domain_name}.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appflow_flow#object_type_name AppflowFlow#object_type_name}.
	ObjectTypeName *string `field:"optional" json:"objectTypeName" yaml:"objectTypeName"`
}

