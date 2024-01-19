package cloudwatcheventtarget


type CloudwatchEventTargetSagemakerPipelineTarget struct {
	// pipeline_parameter_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/cloudwatch_event_target#pipeline_parameter_list CloudwatchEventTarget#pipeline_parameter_list}
	PipelineParameterList interface{} `field:"optional" json:"pipelineParameterList" yaml:"pipelineParameterList"`
}

