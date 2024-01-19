package dataawsvpcipampreviewnextcidr

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsVpcIpamPreviewNextCidrConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/vpc_ipam_preview_next_cidr#ipam_pool_id DataAwsVpcIpamPreviewNextCidr#ipam_pool_id}.
	IpamPoolId *string `field:"required" json:"ipamPoolId" yaml:"ipamPoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/vpc_ipam_preview_next_cidr#disallowed_cidrs DataAwsVpcIpamPreviewNextCidr#disallowed_cidrs}.
	DisallowedCidrs *[]*string `field:"optional" json:"disallowedCidrs" yaml:"disallowedCidrs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/vpc_ipam_preview_next_cidr#id DataAwsVpcIpamPreviewNextCidr#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/vpc_ipam_preview_next_cidr#netmask_length DataAwsVpcIpamPreviewNextCidr#netmask_length}.
	NetmaskLength *float64 `field:"optional" json:"netmaskLength" yaml:"netmaskLength"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/vpc_ipam_preview_next_cidr#timeouts DataAwsVpcIpamPreviewNextCidr#timeouts}
	Timeouts *DataAwsVpcIpamPreviewNextCidrTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

