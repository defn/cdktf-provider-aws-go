package ecsservice


type EcsServiceDeploymentCircuitBreaker struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/ecs_service#enable EcsService#enable}.
	Enable interface{} `field:"required" json:"enable" yaml:"enable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/ecs_service#rollback EcsService#rollback}.
	Rollback interface{} `field:"required" json:"rollback" yaml:"rollback"`
}

