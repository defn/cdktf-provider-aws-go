package medialivemultiplexprogram

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MedialiveMultiplexProgramConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/medialive_multiplex_program#multiplex_id MedialiveMultiplexProgram#multiplex_id}.
	MultiplexId *string `field:"required" json:"multiplexId" yaml:"multiplexId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/medialive_multiplex_program#program_name MedialiveMultiplexProgram#program_name}.
	ProgramName *string `field:"required" json:"programName" yaml:"programName"`
	// multiplex_program_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/medialive_multiplex_program#multiplex_program_settings MedialiveMultiplexProgram#multiplex_program_settings}
	MultiplexProgramSettings interface{} `field:"optional" json:"multiplexProgramSettings" yaml:"multiplexProgramSettings"`
}

