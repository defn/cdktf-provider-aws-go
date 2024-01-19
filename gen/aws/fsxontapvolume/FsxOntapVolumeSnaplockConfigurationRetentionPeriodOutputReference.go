package fsxontapvolume

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/defn/cdktf-provider-aws-go/gen/aws/jsii"

	"github.com/defn/cdktf-provider-aws-go/gen/aws/fsxontapvolume/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FsxOntapVolumeSnaplockConfigurationRetentionPeriodOutputReference interface {
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
	DefaultRetention() FsxOntapVolumeSnaplockConfigurationRetentionPeriodDefaultRetentionOutputReference
	DefaultRetentionInput() *FsxOntapVolumeSnaplockConfigurationRetentionPeriodDefaultRetention
	// Experimental.
	Fqn() *string
	InternalValue() *FsxOntapVolumeSnaplockConfigurationRetentionPeriod
	SetInternalValue(val *FsxOntapVolumeSnaplockConfigurationRetentionPeriod)
	MaximumRetention() FsxOntapVolumeSnaplockConfigurationRetentionPeriodMaximumRetentionOutputReference
	MaximumRetentionInput() *FsxOntapVolumeSnaplockConfigurationRetentionPeriodMaximumRetention
	MinimumRetention() FsxOntapVolumeSnaplockConfigurationRetentionPeriodMinimumRetentionOutputReference
	MinimumRetentionInput() *FsxOntapVolumeSnaplockConfigurationRetentionPeriodMinimumRetention
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
	PutDefaultRetention(value *FsxOntapVolumeSnaplockConfigurationRetentionPeriodDefaultRetention)
	PutMaximumRetention(value *FsxOntapVolumeSnaplockConfigurationRetentionPeriodMaximumRetention)
	PutMinimumRetention(value *FsxOntapVolumeSnaplockConfigurationRetentionPeriodMinimumRetention)
	ResetDefaultRetention()
	ResetMaximumRetention()
	ResetMinimumRetention()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FsxOntapVolumeSnaplockConfigurationRetentionPeriodOutputReference
type jsiiProxy_FsxOntapVolumeSnaplockConfigurationRetentionPeriodOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FsxOntapVolumeSnaplockConfigurationRetentionPeriodOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeSnaplockConfigurationRetentionPeriodOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeSnaplockConfigurationRetentionPeriodOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeSnaplockConfigurationRetentionPeriodOutputReference) DefaultRetention() FsxOntapVolumeSnaplockConfigurationRetentionPeriodDefaultRetentionOutputReference {
	var returns FsxOntapVolumeSnaplockConfigurationRetentionPeriodDefaultRetentionOutputReference
	_jsii_.Get(
		j,
		"defaultRetention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeSnaplockConfigurationRetentionPeriodOutputReference) DefaultRetentionInput() *FsxOntapVolumeSnaplockConfigurationRetentionPeriodDefaultRetention {
	var returns *FsxOntapVolumeSnaplockConfigurationRetentionPeriodDefaultRetention
	_jsii_.Get(
		j,
		"defaultRetentionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeSnaplockConfigurationRetentionPeriodOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeSnaplockConfigurationRetentionPeriodOutputReference) InternalValue() *FsxOntapVolumeSnaplockConfigurationRetentionPeriod {
	var returns *FsxOntapVolumeSnaplockConfigurationRetentionPeriod
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeSnaplockConfigurationRetentionPeriodOutputReference) MaximumRetention() FsxOntapVolumeSnaplockConfigurationRetentionPeriodMaximumRetentionOutputReference {
	var returns FsxOntapVolumeSnaplockConfigurationRetentionPeriodMaximumRetentionOutputReference
	_jsii_.Get(
		j,
		"maximumRetention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeSnaplockConfigurationRetentionPeriodOutputReference) MaximumRetentionInput() *FsxOntapVolumeSnaplockConfigurationRetentionPeriodMaximumRetention {
	var returns *FsxOntapVolumeSnaplockConfigurationRetentionPeriodMaximumRetention
	_jsii_.Get(
		j,
		"maximumRetentionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeSnaplockConfigurationRetentionPeriodOutputReference) MinimumRetention() FsxOntapVolumeSnaplockConfigurationRetentionPeriodMinimumRetentionOutputReference {
	var returns FsxOntapVolumeSnaplockConfigurationRetentionPeriodMinimumRetentionOutputReference
	_jsii_.Get(
		j,
		"minimumRetention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeSnaplockConfigurationRetentionPeriodOutputReference) MinimumRetentionInput() *FsxOntapVolumeSnaplockConfigurationRetentionPeriodMinimumRetention {
	var returns *FsxOntapVolumeSnaplockConfigurationRetentionPeriodMinimumRetention
	_jsii_.Get(
		j,
		"minimumRetentionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeSnaplockConfigurationRetentionPeriodOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeSnaplockConfigurationRetentionPeriodOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewFsxOntapVolumeSnaplockConfigurationRetentionPeriodOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) FsxOntapVolumeSnaplockConfigurationRetentionPeriodOutputReference {
	_init_.Initialize()

	if err := validateNewFsxOntapVolumeSnaplockConfigurationRetentionPeriodOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_FsxOntapVolumeSnaplockConfigurationRetentionPeriodOutputReference{}

	_jsii_.Create(
		"aws.fsxOntapVolume.FsxOntapVolumeSnaplockConfigurationRetentionPeriodOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewFsxOntapVolumeSnaplockConfigurationRetentionPeriodOutputReference_Override(f FsxOntapVolumeSnaplockConfigurationRetentionPeriodOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws.fsxOntapVolume.FsxOntapVolumeSnaplockConfigurationRetentionPeriodOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		f,
	)
}

func (j *jsiiProxy_FsxOntapVolumeSnaplockConfigurationRetentionPeriodOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolumeSnaplockConfigurationRetentionPeriodOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolumeSnaplockConfigurationRetentionPeriodOutputReference)SetInternalValue(val *FsxOntapVolumeSnaplockConfigurationRetentionPeriod) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolumeSnaplockConfigurationRetentionPeriodOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolumeSnaplockConfigurationRetentionPeriodOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (f *jsiiProxy_FsxOntapVolumeSnaplockConfigurationRetentionPeriodOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOntapVolumeSnaplockConfigurationRetentionPeriodOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := f.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOntapVolumeSnaplockConfigurationRetentionPeriodOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOntapVolumeSnaplockConfigurationRetentionPeriodOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := f.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOntapVolumeSnaplockConfigurationRetentionPeriodOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := f.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOntapVolumeSnaplockConfigurationRetentionPeriodOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := f.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOntapVolumeSnaplockConfigurationRetentionPeriodOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := f.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOntapVolumeSnaplockConfigurationRetentionPeriodOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := f.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOntapVolumeSnaplockConfigurationRetentionPeriodOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := f.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOntapVolumeSnaplockConfigurationRetentionPeriodOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := f.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOntapVolumeSnaplockConfigurationRetentionPeriodOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOntapVolumeSnaplockConfigurationRetentionPeriodOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOntapVolumeSnaplockConfigurationRetentionPeriodOutputReference) PutDefaultRetention(value *FsxOntapVolumeSnaplockConfigurationRetentionPeriodDefaultRetention) {
	if err := f.validatePutDefaultRetentionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putDefaultRetention",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxOntapVolumeSnaplockConfigurationRetentionPeriodOutputReference) PutMaximumRetention(value *FsxOntapVolumeSnaplockConfigurationRetentionPeriodMaximumRetention) {
	if err := f.validatePutMaximumRetentionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putMaximumRetention",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxOntapVolumeSnaplockConfigurationRetentionPeriodOutputReference) PutMinimumRetention(value *FsxOntapVolumeSnaplockConfigurationRetentionPeriodMinimumRetention) {
	if err := f.validatePutMinimumRetentionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putMinimumRetention",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxOntapVolumeSnaplockConfigurationRetentionPeriodOutputReference) ResetDefaultRetention() {
	_jsii_.InvokeVoid(
		f,
		"resetDefaultRetention",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapVolumeSnaplockConfigurationRetentionPeriodOutputReference) ResetMaximumRetention() {
	_jsii_.InvokeVoid(
		f,
		"resetMaximumRetention",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapVolumeSnaplockConfigurationRetentionPeriodOutputReference) ResetMinimumRetention() {
	_jsii_.InvokeVoid(
		f,
		"resetMinimumRetention",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapVolumeSnaplockConfigurationRetentionPeriodOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := f.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOntapVolumeSnaplockConfigurationRetentionPeriodOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

