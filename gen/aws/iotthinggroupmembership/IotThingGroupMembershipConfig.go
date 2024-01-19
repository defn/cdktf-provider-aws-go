package iotthinggroupmembership

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IotThingGroupMembershipConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/iot_thing_group_membership#thing_group_name IotThingGroupMembership#thing_group_name}.
	ThingGroupName *string `field:"required" json:"thingGroupName" yaml:"thingGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/iot_thing_group_membership#thing_name IotThingGroupMembership#thing_name}.
	ThingName *string `field:"required" json:"thingName" yaml:"thingName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/iot_thing_group_membership#id IotThingGroupMembership#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/iot_thing_group_membership#override_dynamic_group IotThingGroupMembership#override_dynamic_group}.
	OverrideDynamicGroup interface{} `field:"optional" json:"overrideDynamicGroup" yaml:"overrideDynamicGroup"`
}

