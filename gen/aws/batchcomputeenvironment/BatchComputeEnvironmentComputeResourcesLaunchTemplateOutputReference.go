package batchcomputeenvironment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/defn/cdktf-provider-aws-go/gen/aws/jsii"

	"github.com/defn/cdktf-provider-aws-go/gen/aws/batchcomputeenvironment/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *BatchComputeEnvironmentComputeResourcesLaunchTemplate
	SetInternalValue(val *BatchComputeEnvironmentComputeResourcesLaunchTemplate)
	LaunchTemplateId() *string
	SetLaunchTemplateId(val *string)
	LaunchTemplateIdInput() *string
	LaunchTemplateName() *string
	SetLaunchTemplateName(val *string)
	LaunchTemplateNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
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
	ResetLaunchTemplateId()
	ResetLaunchTemplateName()
	ResetVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference
type jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) InternalValue() *BatchComputeEnvironmentComputeResourcesLaunchTemplate {
	var returns *BatchComputeEnvironmentComputeResourcesLaunchTemplate
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) LaunchTemplateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchTemplateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) LaunchTemplateIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchTemplateIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) LaunchTemplateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchTemplateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) LaunchTemplateNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchTemplateNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


func NewBatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference {
	_init_.Initialize()

	if err := validateNewBatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference{}

	_jsii_.Create(
		"aws.batchComputeEnvironment.BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference_Override(b BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws.batchComputeEnvironment.BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference)SetInternalValue(val *BatchComputeEnvironmentComputeResourcesLaunchTemplate) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference)SetLaunchTemplateId(val *string) {
	if err := j.validateSetLaunchTemplateIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"launchTemplateId",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference)SetLaunchTemplateName(val *string) {
	if err := j.validateSetLaunchTemplateNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"launchTemplateName",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) ResetLaunchTemplateId() {
	_jsii_.InvokeVoid(
		b,
		"resetLaunchTemplateId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) ResetLaunchTemplateName() {
	_jsii_.InvokeVoid(
		b,
		"resetLaunchTemplateName",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) ResetVersion() {
	_jsii_.InvokeVoid(
		b,
		"resetVersion",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := b.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

