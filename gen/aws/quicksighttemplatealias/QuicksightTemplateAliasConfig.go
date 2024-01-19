package quicksighttemplatealias

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type QuicksightTemplateAliasConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/quicksight_template_alias#alias_name QuicksightTemplateAlias#alias_name}.
	AliasName *string `field:"required" json:"aliasName" yaml:"aliasName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/quicksight_template_alias#template_id QuicksightTemplateAlias#template_id}.
	TemplateId *string `field:"required" json:"templateId" yaml:"templateId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/quicksight_template_alias#template_version_number QuicksightTemplateAlias#template_version_number}.
	TemplateVersionNumber *float64 `field:"required" json:"templateVersionNumber" yaml:"templateVersionNumber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/quicksight_template_alias#aws_account_id QuicksightTemplateAlias#aws_account_id}.
	AwsAccountId *string `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
}

