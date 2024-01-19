package lightsailcontainerservice


type LightsailContainerServicePublicDomainNamesCertificate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/lightsail_container_service#certificate_name LightsailContainerService#certificate_name}.
	CertificateName *string `field:"required" json:"certificateName" yaml:"certificateName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/lightsail_container_service#domain_names LightsailContainerService#domain_names}.
	DomainNames *[]*string `field:"required" json:"domainNames" yaml:"domainNames"`
}

