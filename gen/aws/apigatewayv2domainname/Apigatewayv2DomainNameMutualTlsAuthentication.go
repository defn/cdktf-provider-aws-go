package apigatewayv2domainname


type Apigatewayv2DomainNameMutualTlsAuthentication struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/apigatewayv2_domain_name#truststore_uri Apigatewayv2DomainName#truststore_uri}.
	TruststoreUri *string `field:"required" json:"truststoreUri" yaml:"truststoreUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/apigatewayv2_domain_name#truststore_version Apigatewayv2DomainName#truststore_version}.
	TruststoreVersion *string `field:"optional" json:"truststoreVersion" yaml:"truststoreVersion"`
}

