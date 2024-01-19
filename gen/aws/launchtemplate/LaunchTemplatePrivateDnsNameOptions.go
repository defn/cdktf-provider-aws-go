package launchtemplate


type LaunchTemplatePrivateDnsNameOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/launch_template#enable_resource_name_dns_aaaa_record LaunchTemplate#enable_resource_name_dns_aaaa_record}.
	EnableResourceNameDnsAaaaRecord interface{} `field:"optional" json:"enableResourceNameDnsAaaaRecord" yaml:"enableResourceNameDnsAaaaRecord"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/launch_template#enable_resource_name_dns_a_record LaunchTemplate#enable_resource_name_dns_a_record}.
	EnableResourceNameDnsARecord interface{} `field:"optional" json:"enableResourceNameDnsARecord" yaml:"enableResourceNameDnsARecord"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/launch_template#hostname_type LaunchTemplate#hostname_type}.
	HostnameType *string `field:"optional" json:"hostnameType" yaml:"hostnameType"`
}

