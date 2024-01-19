package opsworksmemcachedlayer


type OpsworksMemcachedLayerLoadBasedAutoScaling struct {
	// downscaling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/opsworks_memcached_layer#downscaling OpsworksMemcachedLayer#downscaling}
	Downscaling *OpsworksMemcachedLayerLoadBasedAutoScalingDownscaling `field:"optional" json:"downscaling" yaml:"downscaling"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/opsworks_memcached_layer#enable OpsworksMemcachedLayer#enable}.
	Enable interface{} `field:"optional" json:"enable" yaml:"enable"`
	// upscaling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/opsworks_memcached_layer#upscaling OpsworksMemcachedLayer#upscaling}
	Upscaling *OpsworksMemcachedLayerLoadBasedAutoScalingUpscaling `field:"optional" json:"upscaling" yaml:"upscaling"`
}

