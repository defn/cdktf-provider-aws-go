package rdscluster


type RdsClusterRestoreToPointInTime struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/rds_cluster#source_cluster_identifier RdsCluster#source_cluster_identifier}.
	SourceClusterIdentifier *string `field:"required" json:"sourceClusterIdentifier" yaml:"sourceClusterIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/rds_cluster#restore_to_time RdsCluster#restore_to_time}.
	RestoreToTime *string `field:"optional" json:"restoreToTime" yaml:"restoreToTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/rds_cluster#restore_type RdsCluster#restore_type}.
	RestoreType *string `field:"optional" json:"restoreType" yaml:"restoreType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/rds_cluster#use_latest_restorable_time RdsCluster#use_latest_restorable_time}.
	UseLatestRestorableTime interface{} `field:"optional" json:"useLatestRestorableTime" yaml:"useLatestRestorableTime"`
}

