package dxhostedprivatevirtualinterface

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DxHostedPrivateVirtualInterfaceConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/dx_hosted_private_virtual_interface#address_family DxHostedPrivateVirtualInterface#address_family}.
	AddressFamily *string `field:"required" json:"addressFamily" yaml:"addressFamily"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/dx_hosted_private_virtual_interface#bgp_asn DxHostedPrivateVirtualInterface#bgp_asn}.
	BgpAsn *float64 `field:"required" json:"bgpAsn" yaml:"bgpAsn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/dx_hosted_private_virtual_interface#connection_id DxHostedPrivateVirtualInterface#connection_id}.
	ConnectionId *string `field:"required" json:"connectionId" yaml:"connectionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/dx_hosted_private_virtual_interface#name DxHostedPrivateVirtualInterface#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/dx_hosted_private_virtual_interface#owner_account_id DxHostedPrivateVirtualInterface#owner_account_id}.
	OwnerAccountId *string `field:"required" json:"ownerAccountId" yaml:"ownerAccountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/dx_hosted_private_virtual_interface#vlan DxHostedPrivateVirtualInterface#vlan}.
	Vlan *float64 `field:"required" json:"vlan" yaml:"vlan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/dx_hosted_private_virtual_interface#amazon_address DxHostedPrivateVirtualInterface#amazon_address}.
	AmazonAddress *string `field:"optional" json:"amazonAddress" yaml:"amazonAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/dx_hosted_private_virtual_interface#bgp_auth_key DxHostedPrivateVirtualInterface#bgp_auth_key}.
	BgpAuthKey *string `field:"optional" json:"bgpAuthKey" yaml:"bgpAuthKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/dx_hosted_private_virtual_interface#customer_address DxHostedPrivateVirtualInterface#customer_address}.
	CustomerAddress *string `field:"optional" json:"customerAddress" yaml:"customerAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/dx_hosted_private_virtual_interface#id DxHostedPrivateVirtualInterface#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/dx_hosted_private_virtual_interface#mtu DxHostedPrivateVirtualInterface#mtu}.
	Mtu *float64 `field:"optional" json:"mtu" yaml:"mtu"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/dx_hosted_private_virtual_interface#timeouts DxHostedPrivateVirtualInterface#timeouts}
	Timeouts *DxHostedPrivateVirtualInterfaceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

