package transferserver


type TransferServerWorkflowDetailsOnUpload struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/transfer_server#execution_role TransferServer#execution_role}.
	ExecutionRole *string `field:"required" json:"executionRole" yaml:"executionRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/transfer_server#workflow_id TransferServer#workflow_id}.
	WorkflowId *string `field:"required" json:"workflowId" yaml:"workflowId"`
}

