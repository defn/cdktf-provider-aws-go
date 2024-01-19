package backupselection


type BackupSelectionConditionStringNotEquals struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/backup_selection#key BackupSelection#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/backup_selection#value BackupSelection#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

