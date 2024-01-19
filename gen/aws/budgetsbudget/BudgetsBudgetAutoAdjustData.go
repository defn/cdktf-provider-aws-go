package budgetsbudget


type BudgetsBudgetAutoAdjustData struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/budgets_budget#auto_adjust_type BudgetsBudget#auto_adjust_type}.
	AutoAdjustType *string `field:"required" json:"autoAdjustType" yaml:"autoAdjustType"`
	// historical_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/budgets_budget#historical_options BudgetsBudget#historical_options}
	HistoricalOptions *BudgetsBudgetAutoAdjustDataHistoricalOptions `field:"optional" json:"historicalOptions" yaml:"historicalOptions"`
}

