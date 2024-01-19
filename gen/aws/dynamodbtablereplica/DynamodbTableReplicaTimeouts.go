package dynamodbtablereplica


type DynamodbTableReplicaTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/dynamodb_table_replica#create DynamodbTableReplicaA#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/dynamodb_table_replica#delete DynamodbTableReplicaA#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/dynamodb_table_replica#update DynamodbTableReplicaA#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

