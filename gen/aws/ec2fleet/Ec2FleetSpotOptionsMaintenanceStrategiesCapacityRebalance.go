package ec2fleet


type Ec2FleetSpotOptionsMaintenanceStrategiesCapacityRebalance struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/ec2_fleet#replacement_strategy Ec2Fleet#replacement_strategy}.
	ReplacementStrategy *string `field:"optional" json:"replacementStrategy" yaml:"replacementStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/ec2_fleet#termination_delay Ec2Fleet#termination_delay}.
	TerminationDelay *float64 `field:"optional" json:"terminationDelay" yaml:"terminationDelay"`
}

