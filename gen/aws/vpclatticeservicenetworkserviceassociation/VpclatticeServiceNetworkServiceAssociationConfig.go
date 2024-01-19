package vpclatticeservicenetworkserviceassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpclatticeServiceNetworkServiceAssociationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/vpclattice_service_network_service_association#service_identifier VpclatticeServiceNetworkServiceAssociation#service_identifier}.
	ServiceIdentifier *string `field:"required" json:"serviceIdentifier" yaml:"serviceIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/vpclattice_service_network_service_association#service_network_identifier VpclatticeServiceNetworkServiceAssociation#service_network_identifier}.
	ServiceNetworkIdentifier *string `field:"required" json:"serviceNetworkIdentifier" yaml:"serviceNetworkIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/vpclattice_service_network_service_association#id VpclatticeServiceNetworkServiceAssociation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/vpclattice_service_network_service_association#tags VpclatticeServiceNetworkServiceAssociation#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/vpclattice_service_network_service_association#tags_all VpclatticeServiceNetworkServiceAssociation#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/vpclattice_service_network_service_association#timeouts VpclatticeServiceNetworkServiceAssociation#timeouts}
	Timeouts *VpclatticeServiceNetworkServiceAssociationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

