package guarddutyorganizationconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/defn/cdktf-provider-aws-go/gen/aws/jsii"

	"github.com/defn/cdktf-provider-aws-go/gen/aws/guarddutyorganizationconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GuarddutyOrganizationConfigurationDatasourcesOutputReference interface {
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
	InternalValue() *GuarddutyOrganizationConfigurationDatasources
	SetInternalValue(val *GuarddutyOrganizationConfigurationDatasources)
	Kubernetes() GuarddutyOrganizationConfigurationDatasourcesKubernetesOutputReference
	KubernetesInput() *GuarddutyOrganizationConfigurationDatasourcesKubernetes
	MalwareProtection() GuarddutyOrganizationConfigurationDatasourcesMalwareProtectionOutputReference
	MalwareProtectionInput() *GuarddutyOrganizationConfigurationDatasourcesMalwareProtection
	S3Logs() GuarddutyOrganizationConfigurationDatasourcesS3LogsOutputReference
	S3LogsInput() *GuarddutyOrganizationConfigurationDatasourcesS3Logs
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
	PutKubernetes(value *GuarddutyOrganizationConfigurationDatasourcesKubernetes)
	PutMalwareProtection(value *GuarddutyOrganizationConfigurationDatasourcesMalwareProtection)
	PutS3Logs(value *GuarddutyOrganizationConfigurationDatasourcesS3Logs)
	ResetKubernetes()
	ResetMalwareProtection()
	ResetS3Logs()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GuarddutyOrganizationConfigurationDatasourcesOutputReference
type jsiiProxy_GuarddutyOrganizationConfigurationDatasourcesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GuarddutyOrganizationConfigurationDatasourcesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyOrganizationConfigurationDatasourcesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyOrganizationConfigurationDatasourcesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyOrganizationConfigurationDatasourcesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyOrganizationConfigurationDatasourcesOutputReference) InternalValue() *GuarddutyOrganizationConfigurationDatasources {
	var returns *GuarddutyOrganizationConfigurationDatasources
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyOrganizationConfigurationDatasourcesOutputReference) Kubernetes() GuarddutyOrganizationConfigurationDatasourcesKubernetesOutputReference {
	var returns GuarddutyOrganizationConfigurationDatasourcesKubernetesOutputReference
	_jsii_.Get(
		j,
		"kubernetes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyOrganizationConfigurationDatasourcesOutputReference) KubernetesInput() *GuarddutyOrganizationConfigurationDatasourcesKubernetes {
	var returns *GuarddutyOrganizationConfigurationDatasourcesKubernetes
	_jsii_.Get(
		j,
		"kubernetesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyOrganizationConfigurationDatasourcesOutputReference) MalwareProtection() GuarddutyOrganizationConfigurationDatasourcesMalwareProtectionOutputReference {
	var returns GuarddutyOrganizationConfigurationDatasourcesMalwareProtectionOutputReference
	_jsii_.Get(
		j,
		"malwareProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyOrganizationConfigurationDatasourcesOutputReference) MalwareProtectionInput() *GuarddutyOrganizationConfigurationDatasourcesMalwareProtection {
	var returns *GuarddutyOrganizationConfigurationDatasourcesMalwareProtection
	_jsii_.Get(
		j,
		"malwareProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyOrganizationConfigurationDatasourcesOutputReference) S3Logs() GuarddutyOrganizationConfigurationDatasourcesS3LogsOutputReference {
	var returns GuarddutyOrganizationConfigurationDatasourcesS3LogsOutputReference
	_jsii_.Get(
		j,
		"s3Logs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyOrganizationConfigurationDatasourcesOutputReference) S3LogsInput() *GuarddutyOrganizationConfigurationDatasourcesS3Logs {
	var returns *GuarddutyOrganizationConfigurationDatasourcesS3Logs
	_jsii_.Get(
		j,
		"s3LogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyOrganizationConfigurationDatasourcesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyOrganizationConfigurationDatasourcesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGuarddutyOrganizationConfigurationDatasourcesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GuarddutyOrganizationConfigurationDatasourcesOutputReference {
	_init_.Initialize()

	if err := validateNewGuarddutyOrganizationConfigurationDatasourcesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GuarddutyOrganizationConfigurationDatasourcesOutputReference{}

	_jsii_.Create(
		"aws.guarddutyOrganizationConfiguration.GuarddutyOrganizationConfigurationDatasourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGuarddutyOrganizationConfigurationDatasourcesOutputReference_Override(g GuarddutyOrganizationConfigurationDatasourcesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws.guarddutyOrganizationConfiguration.GuarddutyOrganizationConfigurationDatasourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GuarddutyOrganizationConfigurationDatasourcesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GuarddutyOrganizationConfigurationDatasourcesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GuarddutyOrganizationConfigurationDatasourcesOutputReference)SetInternalValue(val *GuarddutyOrganizationConfigurationDatasources) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GuarddutyOrganizationConfigurationDatasourcesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GuarddutyOrganizationConfigurationDatasourcesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GuarddutyOrganizationConfigurationDatasourcesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuarddutyOrganizationConfigurationDatasourcesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuarddutyOrganizationConfigurationDatasourcesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuarddutyOrganizationConfigurationDatasourcesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuarddutyOrganizationConfigurationDatasourcesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuarddutyOrganizationConfigurationDatasourcesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuarddutyOrganizationConfigurationDatasourcesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuarddutyOrganizationConfigurationDatasourcesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuarddutyOrganizationConfigurationDatasourcesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuarddutyOrganizationConfigurationDatasourcesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuarddutyOrganizationConfigurationDatasourcesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuarddutyOrganizationConfigurationDatasourcesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuarddutyOrganizationConfigurationDatasourcesOutputReference) PutKubernetes(value *GuarddutyOrganizationConfigurationDatasourcesKubernetes) {
	if err := g.validatePutKubernetesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putKubernetes",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GuarddutyOrganizationConfigurationDatasourcesOutputReference) PutMalwareProtection(value *GuarddutyOrganizationConfigurationDatasourcesMalwareProtection) {
	if err := g.validatePutMalwareProtectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMalwareProtection",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GuarddutyOrganizationConfigurationDatasourcesOutputReference) PutS3Logs(value *GuarddutyOrganizationConfigurationDatasourcesS3Logs) {
	if err := g.validatePutS3LogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putS3Logs",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GuarddutyOrganizationConfigurationDatasourcesOutputReference) ResetKubernetes() {
	_jsii_.InvokeVoid(
		g,
		"resetKubernetes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GuarddutyOrganizationConfigurationDatasourcesOutputReference) ResetMalwareProtection() {
	_jsii_.InvokeVoid(
		g,
		"resetMalwareProtection",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GuarddutyOrganizationConfigurationDatasourcesOutputReference) ResetS3Logs() {
	_jsii_.InvokeVoid(
		g,
		"resetS3Logs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GuarddutyOrganizationConfigurationDatasourcesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuarddutyOrganizationConfigurationDatasourcesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

