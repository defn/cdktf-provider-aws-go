package prometheusscraper


type PrometheusScraperDestinationAmp struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/prometheus_scraper#workspace_arn PrometheusScraper#workspace_arn}.
	WorkspaceArn *string `field:"required" json:"workspaceArn" yaml:"workspaceArn"`
}

