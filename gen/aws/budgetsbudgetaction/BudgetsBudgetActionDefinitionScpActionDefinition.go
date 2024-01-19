package budgetsbudgetaction


type BudgetsBudgetActionDefinitionScpActionDefinition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/budgets_budget_action#policy_id BudgetsBudgetAction#policy_id}.
	PolicyId *string `field:"required" json:"policyId" yaml:"policyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/budgets_budget_action#target_ids BudgetsBudgetAction#target_ids}.
	TargetIds *[]*string `field:"required" json:"targetIds" yaml:"targetIds"`
}

