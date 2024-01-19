package sqsqueue

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws.sqsQueue.SqsQueue",
		reflect.TypeOf((*SqsQueue)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "contentBasedDeduplication", GoGetter: "ContentBasedDeduplication"},
			_jsii_.MemberProperty{JsiiProperty: "contentBasedDeduplicationInput", GoGetter: "ContentBasedDeduplicationInput"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "deduplicationScope", GoGetter: "DeduplicationScope"},
			_jsii_.MemberProperty{JsiiProperty: "deduplicationScopeInput", GoGetter: "DeduplicationScopeInput"},
			_jsii_.MemberProperty{JsiiProperty: "delaySeconds", GoGetter: "DelaySeconds"},
			_jsii_.MemberProperty{JsiiProperty: "delaySecondsInput", GoGetter: "DelaySecondsInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "fifoQueue", GoGetter: "FifoQueue"},
			_jsii_.MemberProperty{JsiiProperty: "fifoQueueInput", GoGetter: "FifoQueueInput"},
			_jsii_.MemberProperty{JsiiProperty: "fifoThroughputLimit", GoGetter: "FifoThroughputLimit"},
			_jsii_.MemberProperty{JsiiProperty: "fifoThroughputLimitInput", GoGetter: "FifoThroughputLimitInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "kmsDataKeyReusePeriodSeconds", GoGetter: "KmsDataKeyReusePeriodSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "kmsDataKeyReusePeriodSecondsInput", GoGetter: "KmsDataKeyReusePeriodSecondsInput"},
			_jsii_.MemberProperty{JsiiProperty: "kmsMasterKeyId", GoGetter: "KmsMasterKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "kmsMasterKeyIdInput", GoGetter: "KmsMasterKeyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "maxMessageSize", GoGetter: "MaxMessageSize"},
			_jsii_.MemberProperty{JsiiProperty: "maxMessageSizeInput", GoGetter: "MaxMessageSizeInput"},
			_jsii_.MemberProperty{JsiiProperty: "messageRetentionSeconds", GoGetter: "MessageRetentionSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "messageRetentionSecondsInput", GoGetter: "MessageRetentionSecondsInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "namePrefix", GoGetter: "NamePrefix"},
			_jsii_.MemberProperty{JsiiProperty: "namePrefixInput", GoGetter: "NamePrefixInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "policy", GoGetter: "Policy"},
			_jsii_.MemberProperty{JsiiProperty: "policyInput", GoGetter: "PolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "receiveWaitTimeSeconds", GoGetter: "ReceiveWaitTimeSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "receiveWaitTimeSecondsInput", GoGetter: "ReceiveWaitTimeSecondsInput"},
			_jsii_.MemberProperty{JsiiProperty: "redriveAllowPolicy", GoGetter: "RedriveAllowPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "redriveAllowPolicyInput", GoGetter: "RedriveAllowPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "redrivePolicy", GoGetter: "RedrivePolicy"},
			_jsii_.MemberProperty{JsiiProperty: "redrivePolicyInput", GoGetter: "RedrivePolicyInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetContentBasedDeduplication", GoMethod: "ResetContentBasedDeduplication"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeduplicationScope", GoMethod: "ResetDeduplicationScope"},
			_jsii_.MemberMethod{JsiiMethod: "resetDelaySeconds", GoMethod: "ResetDelaySeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resetFifoQueue", GoMethod: "ResetFifoQueue"},
			_jsii_.MemberMethod{JsiiMethod: "resetFifoThroughputLimit", GoMethod: "ResetFifoThroughputLimit"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetKmsDataKeyReusePeriodSeconds", GoMethod: "ResetKmsDataKeyReusePeriodSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resetKmsMasterKeyId", GoMethod: "ResetKmsMasterKeyId"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxMessageSize", GoMethod: "ResetMaxMessageSize"},
			_jsii_.MemberMethod{JsiiMethod: "resetMessageRetentionSeconds", GoMethod: "ResetMessageRetentionSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberMethod{JsiiMethod: "resetNamePrefix", GoMethod: "ResetNamePrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPolicy", GoMethod: "ResetPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetReceiveWaitTimeSeconds", GoMethod: "ResetReceiveWaitTimeSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resetRedriveAllowPolicy", GoMethod: "ResetRedriveAllowPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetRedrivePolicy", GoMethod: "ResetRedrivePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetSqsManagedSseEnabled", GoMethod: "ResetSqsManagedSseEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTagsAll", GoMethod: "ResetTagsAll"},
			_jsii_.MemberMethod{JsiiMethod: "resetVisibilityTimeoutSeconds", GoMethod: "ResetVisibilityTimeoutSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "sqsManagedSseEnabled", GoGetter: "SqsManagedSseEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "sqsManagedSseEnabledInput", GoGetter: "SqsManagedSseEnabledInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsAll", GoGetter: "TagsAll"},
			_jsii_.MemberProperty{JsiiProperty: "tagsAllInput", GoGetter: "TagsAllInput"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "url", GoGetter: "Url"},
			_jsii_.MemberProperty{JsiiProperty: "visibilityTimeoutSeconds", GoGetter: "VisibilityTimeoutSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "visibilityTimeoutSecondsInput", GoGetter: "VisibilityTimeoutSecondsInput"},
		},
		func() interface{} {
			j := jsiiProxy_SqsQueue{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws.sqsQueue.SqsQueueConfig",
		reflect.TypeOf((*SqsQueueConfig)(nil)).Elem(),
	)
}
