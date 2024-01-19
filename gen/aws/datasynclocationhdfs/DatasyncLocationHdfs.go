package datasynclocationhdfs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/defn/cdktf-provider-aws-go/gen/aws/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/defn/cdktf-provider-aws-go/gen/aws/datasynclocationhdfs/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/datasync_location_hdfs aws_datasync_location_hdfs}.
type DatasyncLocationHdfs interface {
	cdktf.TerraformResource
	AgentArns() *[]*string
	SetAgentArns(val *[]*string)
	AgentArnsInput() *[]*string
	Arn() *string
	AuthenticationType() *string
	SetAuthenticationType(val *string)
	AuthenticationTypeInput() *string
	BlockSize() *float64
	SetBlockSize(val *float64)
	BlockSizeInput() *float64
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	KerberosKeytab() *string
	SetKerberosKeytab(val *string)
	KerberosKeytabInput() *string
	KerberosKrb5Conf() *string
	SetKerberosKrb5Conf(val *string)
	KerberosKrb5ConfInput() *string
	KerberosPrincipal() *string
	SetKerberosPrincipal(val *string)
	KerberosPrincipalInput() *string
	KmsKeyProviderUri() *string
	SetKmsKeyProviderUri(val *string)
	KmsKeyProviderUriInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	NameNode() DatasyncLocationHdfsNameNodeList
	NameNodeInput() interface{}
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	QopConfiguration() DatasyncLocationHdfsQopConfigurationOutputReference
	QopConfigurationInput() *DatasyncLocationHdfsQopConfiguration
	// Experimental.
	RawOverrides() interface{}
	ReplicationFactor() *float64
	SetReplicationFactor(val *float64)
	ReplicationFactorInput() *float64
	SimpleUser() *string
	SetSimpleUser(val *string)
	SimpleUserInput() *string
	Subdirectory() *string
	SetSubdirectory(val *string)
	SubdirectoryInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Uri() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutNameNode(value interface{})
	PutQopConfiguration(value *DatasyncLocationHdfsQopConfiguration)
	ResetAuthenticationType()
	ResetBlockSize()
	ResetId()
	ResetKerberosKeytab()
	ResetKerberosKrb5Conf()
	ResetKerberosPrincipal()
	ResetKmsKeyProviderUri()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetQopConfiguration()
	ResetReplicationFactor()
	ResetSimpleUser()
	ResetSubdirectory()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DatasyncLocationHdfs
type jsiiProxy_DatasyncLocationHdfs struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DatasyncLocationHdfs) AgentArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"agentArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationHdfs) AgentArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"agentArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationHdfs) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationHdfs) AuthenticationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationHdfs) AuthenticationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationHdfs) BlockSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"blockSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationHdfs) BlockSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"blockSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationHdfs) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationHdfs) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationHdfs) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationHdfs) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationHdfs) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationHdfs) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationHdfs) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationHdfs) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationHdfs) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationHdfs) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationHdfs) KerberosKeytab() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kerberosKeytab",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationHdfs) KerberosKeytabInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kerberosKeytabInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationHdfs) KerberosKrb5Conf() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kerberosKrb5Conf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationHdfs) KerberosKrb5ConfInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kerberosKrb5ConfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationHdfs) KerberosPrincipal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kerberosPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationHdfs) KerberosPrincipalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kerberosPrincipalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationHdfs) KmsKeyProviderUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyProviderUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationHdfs) KmsKeyProviderUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyProviderUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationHdfs) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationHdfs) NameNode() DatasyncLocationHdfsNameNodeList {
	var returns DatasyncLocationHdfsNameNodeList
	_jsii_.Get(
		j,
		"nameNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationHdfs) NameNodeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nameNodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationHdfs) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationHdfs) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationHdfs) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationHdfs) QopConfiguration() DatasyncLocationHdfsQopConfigurationOutputReference {
	var returns DatasyncLocationHdfsQopConfigurationOutputReference
	_jsii_.Get(
		j,
		"qopConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationHdfs) QopConfigurationInput() *DatasyncLocationHdfsQopConfiguration {
	var returns *DatasyncLocationHdfsQopConfiguration
	_jsii_.Get(
		j,
		"qopConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationHdfs) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationHdfs) ReplicationFactor() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicationFactor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationHdfs) ReplicationFactorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicationFactorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationHdfs) SimpleUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"simpleUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationHdfs) SimpleUserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"simpleUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationHdfs) Subdirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subdirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationHdfs) SubdirectoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subdirectoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationHdfs) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationHdfs) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationHdfs) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationHdfs) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationHdfs) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationHdfs) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationHdfs) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationHdfs) Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uri",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/datasync_location_hdfs aws_datasync_location_hdfs} Resource.
func NewDatasyncLocationHdfs(scope constructs.Construct, id *string, config *DatasyncLocationHdfsConfig) DatasyncLocationHdfs {
	_init_.Initialize()

	if err := validateNewDatasyncLocationHdfsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatasyncLocationHdfs{}

	_jsii_.Create(
		"aws.datasyncLocationHdfs.DatasyncLocationHdfs",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/datasync_location_hdfs aws_datasync_location_hdfs} Resource.
func NewDatasyncLocationHdfs_Override(d DatasyncLocationHdfs, scope constructs.Construct, id *string, config *DatasyncLocationHdfsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"aws.datasyncLocationHdfs.DatasyncLocationHdfs",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DatasyncLocationHdfs)SetAgentArns(val *[]*string) {
	if err := j.validateSetAgentArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agentArns",
		val,
	)
}

func (j *jsiiProxy_DatasyncLocationHdfs)SetAuthenticationType(val *string) {
	if err := j.validateSetAuthenticationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationType",
		val,
	)
}

func (j *jsiiProxy_DatasyncLocationHdfs)SetBlockSize(val *float64) {
	if err := j.validateSetBlockSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockSize",
		val,
	)
}

func (j *jsiiProxy_DatasyncLocationHdfs)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DatasyncLocationHdfs)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DatasyncLocationHdfs)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DatasyncLocationHdfs)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DatasyncLocationHdfs)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DatasyncLocationHdfs)SetKerberosKeytab(val *string) {
	if err := j.validateSetKerberosKeytabParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kerberosKeytab",
		val,
	)
}

func (j *jsiiProxy_DatasyncLocationHdfs)SetKerberosKrb5Conf(val *string) {
	if err := j.validateSetKerberosKrb5ConfParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kerberosKrb5Conf",
		val,
	)
}

func (j *jsiiProxy_DatasyncLocationHdfs)SetKerberosPrincipal(val *string) {
	if err := j.validateSetKerberosPrincipalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kerberosPrincipal",
		val,
	)
}

func (j *jsiiProxy_DatasyncLocationHdfs)SetKmsKeyProviderUri(val *string) {
	if err := j.validateSetKmsKeyProviderUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyProviderUri",
		val,
	)
}

func (j *jsiiProxy_DatasyncLocationHdfs)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DatasyncLocationHdfs)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DatasyncLocationHdfs)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DatasyncLocationHdfs)SetReplicationFactor(val *float64) {
	if err := j.validateSetReplicationFactorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicationFactor",
		val,
	)
}

func (j *jsiiProxy_DatasyncLocationHdfs)SetSimpleUser(val *string) {
	if err := j.validateSetSimpleUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"simpleUser",
		val,
	)
}

func (j *jsiiProxy_DatasyncLocationHdfs)SetSubdirectory(val *string) {
	if err := j.validateSetSubdirectoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subdirectory",
		val,
	)
}

func (j *jsiiProxy_DatasyncLocationHdfs)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_DatasyncLocationHdfs)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

// Generates CDKTF code for importing a DatasyncLocationHdfs resource upon running "cdktf plan <stack-name>".
func DatasyncLocationHdfs_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDatasyncLocationHdfs_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"aws.datasyncLocationHdfs.DatasyncLocationHdfs",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func DatasyncLocationHdfs_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatasyncLocationHdfs_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws.datasyncLocationHdfs.DatasyncLocationHdfs",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DatasyncLocationHdfs_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatasyncLocationHdfs_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws.datasyncLocationHdfs.DatasyncLocationHdfs",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DatasyncLocationHdfs_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatasyncLocationHdfs_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws.datasyncLocationHdfs.DatasyncLocationHdfs",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DatasyncLocationHdfs_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws.datasyncLocationHdfs.DatasyncLocationHdfs",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DatasyncLocationHdfs) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DatasyncLocationHdfs) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DatasyncLocationHdfs) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatasyncLocationHdfs) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatasyncLocationHdfs) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatasyncLocationHdfs) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatasyncLocationHdfs) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatasyncLocationHdfs) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatasyncLocationHdfs) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatasyncLocationHdfs) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatasyncLocationHdfs) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatasyncLocationHdfs) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatasyncLocationHdfs) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DatasyncLocationHdfs) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatasyncLocationHdfs) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DatasyncLocationHdfs) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DatasyncLocationHdfs) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DatasyncLocationHdfs) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DatasyncLocationHdfs) PutNameNode(value interface{}) {
	if err := d.validatePutNameNodeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putNameNode",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatasyncLocationHdfs) PutQopConfiguration(value *DatasyncLocationHdfsQopConfiguration) {
	if err := d.validatePutQopConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putQopConfiguration",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatasyncLocationHdfs) ResetAuthenticationType() {
	_jsii_.InvokeVoid(
		d,
		"resetAuthenticationType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatasyncLocationHdfs) ResetBlockSize() {
	_jsii_.InvokeVoid(
		d,
		"resetBlockSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatasyncLocationHdfs) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatasyncLocationHdfs) ResetKerberosKeytab() {
	_jsii_.InvokeVoid(
		d,
		"resetKerberosKeytab",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatasyncLocationHdfs) ResetKerberosKrb5Conf() {
	_jsii_.InvokeVoid(
		d,
		"resetKerberosKrb5Conf",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatasyncLocationHdfs) ResetKerberosPrincipal() {
	_jsii_.InvokeVoid(
		d,
		"resetKerberosPrincipal",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatasyncLocationHdfs) ResetKmsKeyProviderUri() {
	_jsii_.InvokeVoid(
		d,
		"resetKmsKeyProviderUri",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatasyncLocationHdfs) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatasyncLocationHdfs) ResetQopConfiguration() {
	_jsii_.InvokeVoid(
		d,
		"resetQopConfiguration",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatasyncLocationHdfs) ResetReplicationFactor() {
	_jsii_.InvokeVoid(
		d,
		"resetReplicationFactor",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatasyncLocationHdfs) ResetSimpleUser() {
	_jsii_.InvokeVoid(
		d,
		"resetSimpleUser",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatasyncLocationHdfs) ResetSubdirectory() {
	_jsii_.InvokeVoid(
		d,
		"resetSubdirectory",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatasyncLocationHdfs) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatasyncLocationHdfs) ResetTagsAll() {
	_jsii_.InvokeVoid(
		d,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatasyncLocationHdfs) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatasyncLocationHdfs) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatasyncLocationHdfs) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatasyncLocationHdfs) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatasyncLocationHdfs) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatasyncLocationHdfs) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

