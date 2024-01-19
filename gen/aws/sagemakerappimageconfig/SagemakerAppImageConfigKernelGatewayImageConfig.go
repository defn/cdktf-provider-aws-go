package sagemakerappimageconfig


type SagemakerAppImageConfigKernelGatewayImageConfig struct {
	// kernel_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/sagemaker_app_image_config#kernel_spec SagemakerAppImageConfig#kernel_spec}
	KernelSpec *SagemakerAppImageConfigKernelGatewayImageConfigKernelSpec `field:"required" json:"kernelSpec" yaml:"kernelSpec"`
	// file_system_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/sagemaker_app_image_config#file_system_config SagemakerAppImageConfig#file_system_config}
	FileSystemConfig *SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfig `field:"optional" json:"fileSystemConfig" yaml:"fileSystemConfig"`
}

