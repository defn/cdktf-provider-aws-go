package fisexperimenttemplate


type FisExperimentTemplateTargetFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/fis_experiment_template#path FisExperimentTemplate#path}.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/fis_experiment_template#values FisExperimentTemplate#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

