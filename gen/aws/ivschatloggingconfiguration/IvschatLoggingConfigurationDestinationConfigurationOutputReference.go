package ivschatloggingconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/defn/cdktf-provider-aws-go/gen/aws/jsii"

	"github.com/defn/cdktf-provider-aws-go/gen/aws/ivschatloggingconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IvschatLoggingConfigurationDestinationConfigurationOutputReference interface {
	cdktf.ComplexObject
	CloudwatchLogs() IvschatLoggingConfigurationDestinationConfigurationCloudwatchLogsOutputReference
	CloudwatchLogsInput() *IvschatLoggingConfigurationDestinationConfigurationCloudwatchLogs
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
	Firehose() IvschatLoggingConfigurationDestinationConfigurationFirehoseOutputReference
	FirehoseInput() *IvschatLoggingConfigurationDestinationConfigurationFirehose
	// Experimental.
	Fqn() *string
	InternalValue() *IvschatLoggingConfigurationDestinationConfiguration
	SetInternalValue(val *IvschatLoggingConfigurationDestinationConfiguration)
	S3() IvschatLoggingConfigurationDestinationConfigurationS3OutputReference
	S3Input() *IvschatLoggingConfigurationDestinationConfigurationS3
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
	PutCloudwatchLogs(value *IvschatLoggingConfigurationDestinationConfigurationCloudwatchLogs)
	PutFirehose(value *IvschatLoggingConfigurationDestinationConfigurationFirehose)
	PutS3(value *IvschatLoggingConfigurationDestinationConfigurationS3)
	ResetCloudwatchLogs()
	ResetFirehose()
	ResetS3()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IvschatLoggingConfigurationDestinationConfigurationOutputReference
type jsiiProxy_IvschatLoggingConfigurationDestinationConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IvschatLoggingConfigurationDestinationConfigurationOutputReference) CloudwatchLogs() IvschatLoggingConfigurationDestinationConfigurationCloudwatchLogsOutputReference {
	var returns IvschatLoggingConfigurationDestinationConfigurationCloudwatchLogsOutputReference
	_jsii_.Get(
		j,
		"cloudwatchLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IvschatLoggingConfigurationDestinationConfigurationOutputReference) CloudwatchLogsInput() *IvschatLoggingConfigurationDestinationConfigurationCloudwatchLogs {
	var returns *IvschatLoggingConfigurationDestinationConfigurationCloudwatchLogs
	_jsii_.Get(
		j,
		"cloudwatchLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IvschatLoggingConfigurationDestinationConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IvschatLoggingConfigurationDestinationConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IvschatLoggingConfigurationDestinationConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IvschatLoggingConfigurationDestinationConfigurationOutputReference) Firehose() IvschatLoggingConfigurationDestinationConfigurationFirehoseOutputReference {
	var returns IvschatLoggingConfigurationDestinationConfigurationFirehoseOutputReference
	_jsii_.Get(
		j,
		"firehose",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IvschatLoggingConfigurationDestinationConfigurationOutputReference) FirehoseInput() *IvschatLoggingConfigurationDestinationConfigurationFirehose {
	var returns *IvschatLoggingConfigurationDestinationConfigurationFirehose
	_jsii_.Get(
		j,
		"firehoseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IvschatLoggingConfigurationDestinationConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IvschatLoggingConfigurationDestinationConfigurationOutputReference) InternalValue() *IvschatLoggingConfigurationDestinationConfiguration {
	var returns *IvschatLoggingConfigurationDestinationConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IvschatLoggingConfigurationDestinationConfigurationOutputReference) S3() IvschatLoggingConfigurationDestinationConfigurationS3OutputReference {
	var returns IvschatLoggingConfigurationDestinationConfigurationS3OutputReference
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IvschatLoggingConfigurationDestinationConfigurationOutputReference) S3Input() *IvschatLoggingConfigurationDestinationConfigurationS3 {
	var returns *IvschatLoggingConfigurationDestinationConfigurationS3
	_jsii_.Get(
		j,
		"s3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IvschatLoggingConfigurationDestinationConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IvschatLoggingConfigurationDestinationConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIvschatLoggingConfigurationDestinationConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) IvschatLoggingConfigurationDestinationConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewIvschatLoggingConfigurationDestinationConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IvschatLoggingConfigurationDestinationConfigurationOutputReference{}

	_jsii_.Create(
		"aws.ivschatLoggingConfiguration.IvschatLoggingConfigurationDestinationConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIvschatLoggingConfigurationDestinationConfigurationOutputReference_Override(i IvschatLoggingConfigurationDestinationConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws.ivschatLoggingConfiguration.IvschatLoggingConfigurationDestinationConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IvschatLoggingConfigurationDestinationConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IvschatLoggingConfigurationDestinationConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IvschatLoggingConfigurationDestinationConfigurationOutputReference)SetInternalValue(val *IvschatLoggingConfigurationDestinationConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IvschatLoggingConfigurationDestinationConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IvschatLoggingConfigurationDestinationConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IvschatLoggingConfigurationDestinationConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IvschatLoggingConfigurationDestinationConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IvschatLoggingConfigurationDestinationConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IvschatLoggingConfigurationDestinationConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IvschatLoggingConfigurationDestinationConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IvschatLoggingConfigurationDestinationConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IvschatLoggingConfigurationDestinationConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IvschatLoggingConfigurationDestinationConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IvschatLoggingConfigurationDestinationConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IvschatLoggingConfigurationDestinationConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IvschatLoggingConfigurationDestinationConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IvschatLoggingConfigurationDestinationConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IvschatLoggingConfigurationDestinationConfigurationOutputReference) PutCloudwatchLogs(value *IvschatLoggingConfigurationDestinationConfigurationCloudwatchLogs) {
	if err := i.validatePutCloudwatchLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putCloudwatchLogs",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IvschatLoggingConfigurationDestinationConfigurationOutputReference) PutFirehose(value *IvschatLoggingConfigurationDestinationConfigurationFirehose) {
	if err := i.validatePutFirehoseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putFirehose",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IvschatLoggingConfigurationDestinationConfigurationOutputReference) PutS3(value *IvschatLoggingConfigurationDestinationConfigurationS3) {
	if err := i.validatePutS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putS3",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IvschatLoggingConfigurationDestinationConfigurationOutputReference) ResetCloudwatchLogs() {
	_jsii_.InvokeVoid(
		i,
		"resetCloudwatchLogs",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IvschatLoggingConfigurationDestinationConfigurationOutputReference) ResetFirehose() {
	_jsii_.InvokeVoid(
		i,
		"resetFirehose",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IvschatLoggingConfigurationDestinationConfigurationOutputReference) ResetS3() {
	_jsii_.InvokeVoid(
		i,
		"resetS3",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IvschatLoggingConfigurationDestinationConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (i *jsiiProxy_IvschatLoggingConfigurationDestinationConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

