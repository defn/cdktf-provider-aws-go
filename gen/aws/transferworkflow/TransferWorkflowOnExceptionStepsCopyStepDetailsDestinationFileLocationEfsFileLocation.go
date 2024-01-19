package transferworkflow


type TransferWorkflowOnExceptionStepsCopyStepDetailsDestinationFileLocationEfsFileLocation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/transfer_workflow#file_system_id TransferWorkflow#file_system_id}.
	FileSystemId *string `field:"optional" json:"fileSystemId" yaml:"fileSystemId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/transfer_workflow#path TransferWorkflow#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
}

