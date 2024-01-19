package dxmacseckeyassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DxMacsecKeyAssociationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/dx_macsec_key_association#connection_id DxMacsecKeyAssociation#connection_id}.
	ConnectionId *string `field:"required" json:"connectionId" yaml:"connectionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/dx_macsec_key_association#cak DxMacsecKeyAssociation#cak}.
	Cak *string `field:"optional" json:"cak" yaml:"cak"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/dx_macsec_key_association#ckn DxMacsecKeyAssociation#ckn}.
	Ckn *string `field:"optional" json:"ckn" yaml:"ckn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/dx_macsec_key_association#id DxMacsecKeyAssociation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/dx_macsec_key_association#secret_arn DxMacsecKeyAssociation#secret_arn}.
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
}

