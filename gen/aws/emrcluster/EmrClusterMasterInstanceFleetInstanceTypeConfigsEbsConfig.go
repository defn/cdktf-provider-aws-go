package emrcluster


type EmrClusterMasterInstanceFleetInstanceTypeConfigsEbsConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/emr_cluster#size EmrCluster#size}.
	Size *float64 `field:"required" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/emr_cluster#type EmrCluster#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/emr_cluster#iops EmrCluster#iops}.
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/emr_cluster#volumes_per_instance EmrCluster#volumes_per_instance}.
	VolumesPerInstance *float64 `field:"optional" json:"volumesPerInstance" yaml:"volumesPerInstance"`
}

