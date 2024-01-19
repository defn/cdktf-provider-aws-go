package dbinstanceautomatedbackupsreplication


type DbInstanceAutomatedBackupsReplicationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/db_instance_automated_backups_replication#create DbInstanceAutomatedBackupsReplication#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/db_instance_automated_backups_replication#delete DbInstanceAutomatedBackupsReplication#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}
