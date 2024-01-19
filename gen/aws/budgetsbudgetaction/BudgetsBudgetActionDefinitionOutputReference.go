package budgetsbudgetaction

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/defn/cdktf-provider-aws-go/gen/aws/jsii"

	"github.com/defn/cdktf-provider-aws-go/gen/aws/budgetsbudgetaction/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BudgetsBudgetActionDefinitionOutputReference interface {
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
	IamActionDefinition() BudgetsBudgetActionDefinitionIamActionDefinitionOutputReference
	IamActionDefinitionInput() *BudgetsBudgetActionDefinitionIamActionDefinition
	InternalValue() *BudgetsBudgetActionDefinition
	SetInternalValue(val *BudgetsBudgetActionDefinition)
	ScpActionDefinition() BudgetsBudgetActionDefinitionScpActionDefinitionOutputReference
	ScpActionDefinitionInput() *BudgetsBudgetActionDefinitionScpActionDefinition
	SsmActionDefinition() BudgetsBudgetActionDefinitionSsmActionDefinitionOutputReference
	SsmActionDefinitionInput() *BudgetsBudgetActionDefinitionSsmActionDefinition
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
	PutIamActionDefinition(value *BudgetsBudgetActionDefinitionIamActionDefinition)
	PutScpActionDefinition(value *BudgetsBudgetActionDefinitionScpActionDefinition)
	PutSsmActionDefinition(value *BudgetsBudgetActionDefinitionSsmActionDefinition)
	ResetIamActionDefinition()
	ResetScpActionDefinition()
	ResetSsmActionDefinition()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BudgetsBudgetActionDefinitionOutputReference
type jsiiProxy_BudgetsBudgetActionDefinitionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) IamActionDefinition() BudgetsBudgetActionDefinitionIamActionDefinitionOutputReference {
	var returns BudgetsBudgetActionDefinitionIamActionDefinitionOutputReference
	_jsii_.Get(
		j,
		"iamActionDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) IamActionDefinitionInput() *BudgetsBudgetActionDefinitionIamActionDefinition {
	var returns *BudgetsBudgetActionDefinitionIamActionDefinition
	_jsii_.Get(
		j,
		"iamActionDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) InternalValue() *BudgetsBudgetActionDefinition {
	var returns *BudgetsBudgetActionDefinition
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) ScpActionDefinition() BudgetsBudgetActionDefinitionScpActionDefinitionOutputReference {
	var returns BudgetsBudgetActionDefinitionScpActionDefinitionOutputReference
	_jsii_.Get(
		j,
		"scpActionDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) ScpActionDefinitionInput() *BudgetsBudgetActionDefinitionScpActionDefinition {
	var returns *BudgetsBudgetActionDefinitionScpActionDefinition
	_jsii_.Get(
		j,
		"scpActionDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) SsmActionDefinition() BudgetsBudgetActionDefinitionSsmActionDefinitionOutputReference {
	var returns BudgetsBudgetActionDefinitionSsmActionDefinitionOutputReference
	_jsii_.Get(
		j,
		"ssmActionDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) SsmActionDefinitionInput() *BudgetsBudgetActionDefinitionSsmActionDefinition {
	var returns *BudgetsBudgetActionDefinitionSsmActionDefinition
	_jsii_.Get(
		j,
		"ssmActionDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBudgetsBudgetActionDefinitionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BudgetsBudgetActionDefinitionOutputReference {
	_init_.Initialize()

	if err := validateNewBudgetsBudgetActionDefinitionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BudgetsBudgetActionDefinitionOutputReference{}

	_jsii_.Create(
		"aws.budgetsBudgetAction.BudgetsBudgetActionDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBudgetsBudgetActionDefinitionOutputReference_Override(b BudgetsBudgetActionDefinitionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws.budgetsBudgetAction.BudgetsBudgetActionDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference)SetInternalValue(val *BudgetsBudgetActionDefinition) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) PutIamActionDefinition(value *BudgetsBudgetActionDefinitionIamActionDefinition) {
	if err := b.validatePutIamActionDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putIamActionDefinition",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) PutScpActionDefinition(value *BudgetsBudgetActionDefinitionScpActionDefinition) {
	if err := b.validatePutScpActionDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putScpActionDefinition",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) PutSsmActionDefinition(value *BudgetsBudgetActionDefinitionSsmActionDefinition) {
	if err := b.validatePutSsmActionDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putSsmActionDefinition",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) ResetIamActionDefinition() {
	_jsii_.InvokeVoid(
		b,
		"resetIamActionDefinition",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) ResetScpActionDefinition() {
	_jsii_.InvokeVoid(
		b,
		"resetScpActionDefinition",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) ResetSsmActionDefinition() {
	_jsii_.InvokeVoid(
		b,
		"resetSsmActionDefinition",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (b *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

