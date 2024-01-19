package redshiftscheduledaction

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/defn/cdktf-provider-aws-go/gen/aws/jsii"

	"github.com/defn/cdktf-provider-aws-go/gen/aws/redshiftscheduledaction/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RedshiftScheduledActionTargetActionOutputReference interface {
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
	InternalValue() *RedshiftScheduledActionTargetAction
	SetInternalValue(val *RedshiftScheduledActionTargetAction)
	PauseCluster() RedshiftScheduledActionTargetActionPauseClusterOutputReference
	PauseClusterInput() *RedshiftScheduledActionTargetActionPauseCluster
	ResizeCluster() RedshiftScheduledActionTargetActionResizeClusterOutputReference
	ResizeClusterInput() *RedshiftScheduledActionTargetActionResizeCluster
	ResumeCluster() RedshiftScheduledActionTargetActionResumeClusterOutputReference
	ResumeClusterInput() *RedshiftScheduledActionTargetActionResumeCluster
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
	PutPauseCluster(value *RedshiftScheduledActionTargetActionPauseCluster)
	PutResizeCluster(value *RedshiftScheduledActionTargetActionResizeCluster)
	PutResumeCluster(value *RedshiftScheduledActionTargetActionResumeCluster)
	ResetPauseCluster()
	ResetResizeCluster()
	ResetResumeCluster()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RedshiftScheduledActionTargetActionOutputReference
type jsiiProxy_RedshiftScheduledActionTargetActionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_RedshiftScheduledActionTargetActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftScheduledActionTargetActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftScheduledActionTargetActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftScheduledActionTargetActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftScheduledActionTargetActionOutputReference) InternalValue() *RedshiftScheduledActionTargetAction {
	var returns *RedshiftScheduledActionTargetAction
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftScheduledActionTargetActionOutputReference) PauseCluster() RedshiftScheduledActionTargetActionPauseClusterOutputReference {
	var returns RedshiftScheduledActionTargetActionPauseClusterOutputReference
	_jsii_.Get(
		j,
		"pauseCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftScheduledActionTargetActionOutputReference) PauseClusterInput() *RedshiftScheduledActionTargetActionPauseCluster {
	var returns *RedshiftScheduledActionTargetActionPauseCluster
	_jsii_.Get(
		j,
		"pauseClusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftScheduledActionTargetActionOutputReference) ResizeCluster() RedshiftScheduledActionTargetActionResizeClusterOutputReference {
	var returns RedshiftScheduledActionTargetActionResizeClusterOutputReference
	_jsii_.Get(
		j,
		"resizeCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftScheduledActionTargetActionOutputReference) ResizeClusterInput() *RedshiftScheduledActionTargetActionResizeCluster {
	var returns *RedshiftScheduledActionTargetActionResizeCluster
	_jsii_.Get(
		j,
		"resizeClusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftScheduledActionTargetActionOutputReference) ResumeCluster() RedshiftScheduledActionTargetActionResumeClusterOutputReference {
	var returns RedshiftScheduledActionTargetActionResumeClusterOutputReference
	_jsii_.Get(
		j,
		"resumeCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftScheduledActionTargetActionOutputReference) ResumeClusterInput() *RedshiftScheduledActionTargetActionResumeCluster {
	var returns *RedshiftScheduledActionTargetActionResumeCluster
	_jsii_.Get(
		j,
		"resumeClusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftScheduledActionTargetActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftScheduledActionTargetActionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewRedshiftScheduledActionTargetActionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) RedshiftScheduledActionTargetActionOutputReference {
	_init_.Initialize()

	if err := validateNewRedshiftScheduledActionTargetActionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_RedshiftScheduledActionTargetActionOutputReference{}

	_jsii_.Create(
		"aws.redshiftScheduledAction.RedshiftScheduledActionTargetActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewRedshiftScheduledActionTargetActionOutputReference_Override(r RedshiftScheduledActionTargetActionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws.redshiftScheduledAction.RedshiftScheduledActionTargetActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_RedshiftScheduledActionTargetActionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_RedshiftScheduledActionTargetActionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_RedshiftScheduledActionTargetActionOutputReference)SetInternalValue(val *RedshiftScheduledActionTargetAction) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RedshiftScheduledActionTargetActionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RedshiftScheduledActionTargetActionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_RedshiftScheduledActionTargetActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftScheduledActionTargetActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftScheduledActionTargetActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftScheduledActionTargetActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftScheduledActionTargetActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftScheduledActionTargetActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftScheduledActionTargetActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftScheduledActionTargetActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftScheduledActionTargetActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftScheduledActionTargetActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftScheduledActionTargetActionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftScheduledActionTargetActionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftScheduledActionTargetActionOutputReference) PutPauseCluster(value *RedshiftScheduledActionTargetActionPauseCluster) {
	if err := r.validatePutPauseClusterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putPauseCluster",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RedshiftScheduledActionTargetActionOutputReference) PutResizeCluster(value *RedshiftScheduledActionTargetActionResizeCluster) {
	if err := r.validatePutResizeClusterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putResizeCluster",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RedshiftScheduledActionTargetActionOutputReference) PutResumeCluster(value *RedshiftScheduledActionTargetActionResumeCluster) {
	if err := r.validatePutResumeClusterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putResumeCluster",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RedshiftScheduledActionTargetActionOutputReference) ResetPauseCluster() {
	_jsii_.InvokeVoid(
		r,
		"resetPauseCluster",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftScheduledActionTargetActionOutputReference) ResetResizeCluster() {
	_jsii_.InvokeVoid(
		r,
		"resetResizeCluster",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftScheduledActionTargetActionOutputReference) ResetResumeCluster() {
	_jsii_.InvokeVoid(
		r,
		"resetResumeCluster",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftScheduledActionTargetActionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := r.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftScheduledActionTargetActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

