package dataawsssoadminprincipalapplicationassignments

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsSsoadminPrincipalApplicationAssignmentsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/ssoadmin_principal_application_assignments#instance_arn DataAwsSsoadminPrincipalApplicationAssignments#instance_arn}.
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/ssoadmin_principal_application_assignments#principal_id DataAwsSsoadminPrincipalApplicationAssignments#principal_id}.
	PrincipalId *string `field:"required" json:"principalId" yaml:"principalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/ssoadmin_principal_application_assignments#principal_type DataAwsSsoadminPrincipalApplicationAssignments#principal_type}.
	PrincipalType *string `field:"required" json:"principalType" yaml:"principalType"`
	// application_assignments block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/ssoadmin_principal_application_assignments#application_assignments DataAwsSsoadminPrincipalApplicationAssignments#application_assignments}
	ApplicationAssignments interface{} `field:"optional" json:"applicationAssignments" yaml:"applicationAssignments"`
}

