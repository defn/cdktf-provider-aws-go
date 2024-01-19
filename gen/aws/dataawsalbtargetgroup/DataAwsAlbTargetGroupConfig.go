package dataawsalbtargetgroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsAlbTargetGroupConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/alb_target_group#arn DataAwsAlbTargetGroup#arn}.
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/alb_target_group#id DataAwsAlbTargetGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/alb_target_group#load_balancing_anomaly_mitigation DataAwsAlbTargetGroup#load_balancing_anomaly_mitigation}.
	LoadBalancingAnomalyMitigation *string `field:"optional" json:"loadBalancingAnomalyMitigation" yaml:"loadBalancingAnomalyMitigation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/alb_target_group#name DataAwsAlbTargetGroup#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/alb_target_group#tags DataAwsAlbTargetGroup#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/alb_target_group#timeouts DataAwsAlbTargetGroup#timeouts}
	Timeouts *DataAwsAlbTargetGroupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

