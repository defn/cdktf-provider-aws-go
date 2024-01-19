package mskreplicator


type MskReplicatorKafkaClusterVpcConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/msk_replicator#subnet_ids MskReplicator#subnet_ids}.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/msk_replicator#security_groups_ids MskReplicator#security_groups_ids}.
	SecurityGroupsIds *[]*string `field:"optional" json:"securityGroupsIds" yaml:"securityGroupsIds"`
}

