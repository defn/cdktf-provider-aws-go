package emrserverlessapplication


type EmrserverlessApplicationInitialCapacityInitialCapacityConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/emrserverless_application#worker_count EmrserverlessApplication#worker_count}.
	WorkerCount *float64 `field:"required" json:"workerCount" yaml:"workerCount"`
	// worker_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/emrserverless_application#worker_configuration EmrserverlessApplication#worker_configuration}
	WorkerConfiguration *EmrserverlessApplicationInitialCapacityInitialCapacityConfigWorkerConfiguration `field:"optional" json:"workerConfiguration" yaml:"workerConfiguration"`
}

