package dmsreplicationinstance


type DmsReplicationInstanceTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/dms_replication_instance#create DmsReplicationInstance#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/dms_replication_instance#delete DmsReplicationInstance#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/dms_replication_instance#update DmsReplicationInstance#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

