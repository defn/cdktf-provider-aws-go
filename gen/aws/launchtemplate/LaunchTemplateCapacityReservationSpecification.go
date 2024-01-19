package launchtemplate


type LaunchTemplateCapacityReservationSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/launch_template#capacity_reservation_preference LaunchTemplate#capacity_reservation_preference}.
	CapacityReservationPreference *string `field:"optional" json:"capacityReservationPreference" yaml:"capacityReservationPreference"`
	// capacity_reservation_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/launch_template#capacity_reservation_target LaunchTemplate#capacity_reservation_target}
	CapacityReservationTarget *LaunchTemplateCapacityReservationSpecificationCapacityReservationTarget `field:"optional" json:"capacityReservationTarget" yaml:"capacityReservationTarget"`
}

