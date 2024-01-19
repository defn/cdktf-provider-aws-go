package cloudfrontfieldlevelencryptionprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/defn/cdktf-provider-aws-go/gen/aws/jsii"

	"github.com/defn/cdktf-provider-aws-go/gen/aws/cloudfrontfieldlevelencryptionprofile/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItemsOutputReference interface {
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
	FieldPatterns() CloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItemsFieldPatternsOutputReference
	FieldPatternsInput() *CloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItemsFieldPatterns
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ProviderId() *string
	SetProviderId(val *string)
	ProviderIdInput() *string
	PublicKeyId() *string
	SetPublicKeyId(val *string)
	PublicKeyIdInput() *string
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
	PutFieldPatterns(value *CloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItemsFieldPatterns)
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItemsOutputReference
type jsiiProxy_CloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItemsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItemsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItemsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItemsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItemsOutputReference) FieldPatterns() CloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItemsFieldPatternsOutputReference {
	var returns CloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItemsFieldPatternsOutputReference
	_jsii_.Get(
		j,
		"fieldPatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItemsOutputReference) FieldPatternsInput() *CloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItemsFieldPatterns {
	var returns *CloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItemsFieldPatterns
	_jsii_.Get(
		j,
		"fieldPatternsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItemsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItemsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItemsOutputReference) ProviderId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItemsOutputReference) ProviderIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItemsOutputReference) PublicKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItemsOutputReference) PublicKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItemsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItemsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItemsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItemsOutputReference {
	_init_.Initialize()

	if err := validateNewCloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItemsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItemsOutputReference{}

	_jsii_.Create(
		"aws.cloudfrontFieldLevelEncryptionProfile.CloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItemsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItemsOutputReference_Override(c CloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItemsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"aws.cloudfrontFieldLevelEncryptionProfile.CloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItemsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItemsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItemsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItemsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItemsOutputReference)SetProviderId(val *string) {
	if err := j.validateSetProviderIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"providerId",
		val,
	)
}

func (j *jsiiProxy_CloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItemsOutputReference)SetPublicKeyId(val *string) {
	if err := j.validateSetPublicKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicKeyId",
		val,
	)
}

func (j *jsiiProxy_CloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItemsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItemsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItemsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItemsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItemsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItemsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItemsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItemsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItemsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItemsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItemsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItemsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItemsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItemsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItemsOutputReference) PutFieldPatterns(value *CloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItemsFieldPatterns) {
	if err := c.validatePutFieldPatternsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putFieldPatterns",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItemsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItemsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

