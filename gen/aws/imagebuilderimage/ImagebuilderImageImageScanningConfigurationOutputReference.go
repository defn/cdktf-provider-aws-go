package imagebuilderimage

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/defn/cdktf-provider-aws-go/gen/aws/jsii"

	"github.com/defn/cdktf-provider-aws-go/gen/aws/imagebuilderimage/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ImagebuilderImageImageScanningConfigurationOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EcrConfiguration() ImagebuilderImageImageScanningConfigurationEcrConfigurationOutputReference
	EcrConfigurationInput() *ImagebuilderImageImageScanningConfigurationEcrConfiguration
	// Experimental.
	Fqn() *string
	ImageScanningEnabled() interface{}
	SetImageScanningEnabled(val interface{})
	ImageScanningEnabledInput() interface{}
	InternalValue() *ImagebuilderImageImageScanningConfiguration
	SetInternalValue(val *ImagebuilderImageImageScanningConfiguration)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutEcrConfiguration(value *ImagebuilderImageImageScanningConfigurationEcrConfiguration)
	ResetEcrConfiguration()
	ResetImageScanningEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ImagebuilderImageImageScanningConfigurationOutputReference
type jsiiProxy_ImagebuilderImageImageScanningConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ImagebuilderImageImageScanningConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageImageScanningConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageImageScanningConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageImageScanningConfigurationOutputReference) EcrConfiguration() ImagebuilderImageImageScanningConfigurationEcrConfigurationOutputReference {
	var returns ImagebuilderImageImageScanningConfigurationEcrConfigurationOutputReference
	_jsii_.Get(
		j,
		"ecrConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageImageScanningConfigurationOutputReference) EcrConfigurationInput() *ImagebuilderImageImageScanningConfigurationEcrConfiguration {
	var returns *ImagebuilderImageImageScanningConfigurationEcrConfiguration
	_jsii_.Get(
		j,
		"ecrConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageImageScanningConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageImageScanningConfigurationOutputReference) ImageScanningEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"imageScanningEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageImageScanningConfigurationOutputReference) ImageScanningEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"imageScanningEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageImageScanningConfigurationOutputReference) InternalValue() *ImagebuilderImageImageScanningConfiguration {
	var returns *ImagebuilderImageImageScanningConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageImageScanningConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderImageImageScanningConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewImagebuilderImageImageScanningConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ImagebuilderImageImageScanningConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewImagebuilderImageImageScanningConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ImagebuilderImageImageScanningConfigurationOutputReference{}

	_jsii_.Create(
		"aws.imagebuilderImage.ImagebuilderImageImageScanningConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewImagebuilderImageImageScanningConfigurationOutputReference_Override(i ImagebuilderImageImageScanningConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws.imagebuilderImage.ImagebuilderImageImageScanningConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_ImagebuilderImageImageScanningConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderImageImageScanningConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderImageImageScanningConfigurationOutputReference)SetImageScanningEnabled(val interface{}) {
	if err := j.validateSetImageScanningEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageScanningEnabled",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderImageImageScanningConfigurationOutputReference)SetInternalValue(val *ImagebuilderImageImageScanningConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderImageImageScanningConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderImageImageScanningConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_ImagebuilderImageImageScanningConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderImageImageScanningConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderImageImageScanningConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderImageImageScanningConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderImageImageScanningConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderImageImageScanningConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderImageImageScanningConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderImageImageScanningConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderImageImageScanningConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderImageImageScanningConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderImageImageScanningConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderImageImageScanningConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderImageImageScanningConfigurationOutputReference) PutEcrConfiguration(value *ImagebuilderImageImageScanningConfigurationEcrConfiguration) {
	if err := i.validatePutEcrConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putEcrConfiguration",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_ImagebuilderImageImageScanningConfigurationOutputReference) ResetEcrConfiguration() {
	_jsii_.InvokeVoid(
		i,
		"resetEcrConfiguration",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderImageImageScanningConfigurationOutputReference) ResetImageScanningEnabled() {
	_jsii_.InvokeVoid(
		i,
		"resetImageScanningEnabled",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderImageImageScanningConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := i.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderImageImageScanningConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

