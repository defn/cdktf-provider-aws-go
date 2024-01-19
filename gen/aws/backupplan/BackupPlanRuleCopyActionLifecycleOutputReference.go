package backupplan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/defn/cdktf-provider-aws-go/gen/aws/jsii"

	"github.com/defn/cdktf-provider-aws-go/gen/aws/backupplan/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BackupPlanRuleCopyActionLifecycleOutputReference interface {
	cdktf.ComplexObject
	ColdStorageAfter() *float64
	SetColdStorageAfter(val *float64)
	ColdStorageAfterInput() *float64
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
	DeleteAfter() *float64
	SetDeleteAfter(val *float64)
	DeleteAfterInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *BackupPlanRuleCopyActionLifecycle
	SetInternalValue(val *BackupPlanRuleCopyActionLifecycle)
	OptInToArchiveForSupportedResources() interface{}
	SetOptInToArchiveForSupportedResources(val interface{})
	OptInToArchiveForSupportedResourcesInput() interface{}
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
	ResetColdStorageAfter()
	ResetDeleteAfter()
	ResetOptInToArchiveForSupportedResources()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BackupPlanRuleCopyActionLifecycleOutputReference
type jsiiProxy_BackupPlanRuleCopyActionLifecycleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BackupPlanRuleCopyActionLifecycleOutputReference) ColdStorageAfter() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"coldStorageAfter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPlanRuleCopyActionLifecycleOutputReference) ColdStorageAfterInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"coldStorageAfterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPlanRuleCopyActionLifecycleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPlanRuleCopyActionLifecycleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPlanRuleCopyActionLifecycleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPlanRuleCopyActionLifecycleOutputReference) DeleteAfter() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deleteAfter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPlanRuleCopyActionLifecycleOutputReference) DeleteAfterInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deleteAfterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPlanRuleCopyActionLifecycleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPlanRuleCopyActionLifecycleOutputReference) InternalValue() *BackupPlanRuleCopyActionLifecycle {
	var returns *BackupPlanRuleCopyActionLifecycle
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPlanRuleCopyActionLifecycleOutputReference) OptInToArchiveForSupportedResources() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"optInToArchiveForSupportedResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPlanRuleCopyActionLifecycleOutputReference) OptInToArchiveForSupportedResourcesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"optInToArchiveForSupportedResourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPlanRuleCopyActionLifecycleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPlanRuleCopyActionLifecycleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBackupPlanRuleCopyActionLifecycleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BackupPlanRuleCopyActionLifecycleOutputReference {
	_init_.Initialize()

	if err := validateNewBackupPlanRuleCopyActionLifecycleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BackupPlanRuleCopyActionLifecycleOutputReference{}

	_jsii_.Create(
		"aws.backupPlan.BackupPlanRuleCopyActionLifecycleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBackupPlanRuleCopyActionLifecycleOutputReference_Override(b BackupPlanRuleCopyActionLifecycleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws.backupPlan.BackupPlanRuleCopyActionLifecycleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BackupPlanRuleCopyActionLifecycleOutputReference)SetColdStorageAfter(val *float64) {
	if err := j.validateSetColdStorageAfterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"coldStorageAfter",
		val,
	)
}

func (j *jsiiProxy_BackupPlanRuleCopyActionLifecycleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BackupPlanRuleCopyActionLifecycleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BackupPlanRuleCopyActionLifecycleOutputReference)SetDeleteAfter(val *float64) {
	if err := j.validateSetDeleteAfterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteAfter",
		val,
	)
}

func (j *jsiiProxy_BackupPlanRuleCopyActionLifecycleOutputReference)SetInternalValue(val *BackupPlanRuleCopyActionLifecycle) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BackupPlanRuleCopyActionLifecycleOutputReference)SetOptInToArchiveForSupportedResources(val interface{}) {
	if err := j.validateSetOptInToArchiveForSupportedResourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"optInToArchiveForSupportedResources",
		val,
	)
}

func (j *jsiiProxy_BackupPlanRuleCopyActionLifecycleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BackupPlanRuleCopyActionLifecycleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BackupPlanRuleCopyActionLifecycleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPlanRuleCopyActionLifecycleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BackupPlanRuleCopyActionLifecycleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BackupPlanRuleCopyActionLifecycleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BackupPlanRuleCopyActionLifecycleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BackupPlanRuleCopyActionLifecycleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BackupPlanRuleCopyActionLifecycleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BackupPlanRuleCopyActionLifecycleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BackupPlanRuleCopyActionLifecycleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BackupPlanRuleCopyActionLifecycleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BackupPlanRuleCopyActionLifecycleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPlanRuleCopyActionLifecycleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BackupPlanRuleCopyActionLifecycleOutputReference) ResetColdStorageAfter() {
	_jsii_.InvokeVoid(
		b,
		"resetColdStorageAfter",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupPlanRuleCopyActionLifecycleOutputReference) ResetDeleteAfter() {
	_jsii_.InvokeVoid(
		b,
		"resetDeleteAfter",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupPlanRuleCopyActionLifecycleOutputReference) ResetOptInToArchiveForSupportedResources() {
	_jsii_.InvokeVoid(
		b,
		"resetOptInToArchiveForSupportedResources",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupPlanRuleCopyActionLifecycleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (b *jsiiProxy_BackupPlanRuleCopyActionLifecycleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

