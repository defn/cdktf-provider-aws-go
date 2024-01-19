package backupvault


type BackupVaultTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/backup_vault#delete BackupVault#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

