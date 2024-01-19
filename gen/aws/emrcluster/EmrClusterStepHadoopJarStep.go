package emrcluster


type EmrClusterStepHadoopJarStep struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/emr_cluster#args EmrCluster#args}.
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/emr_cluster#jar EmrCluster#jar}.
	Jar *string `field:"optional" json:"jar" yaml:"jar"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/emr_cluster#main_class EmrCluster#main_class}.
	MainClass *string `field:"optional" json:"mainClass" yaml:"mainClass"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/emr_cluster#properties EmrCluster#properties}.
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
}

