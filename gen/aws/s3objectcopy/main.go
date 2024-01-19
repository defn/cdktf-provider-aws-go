package s3objectcopy

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws.s3ObjectCopy.S3ObjectCopy",
		reflect.TypeOf((*S3ObjectCopy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "acl", GoGetter: "Acl"},
			_jsii_.MemberProperty{JsiiProperty: "aclInput", GoGetter: "AclInput"},
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "bucket", GoGetter: "Bucket"},
			_jsii_.MemberProperty{JsiiProperty: "bucketInput", GoGetter: "BucketInput"},
			_jsii_.MemberProperty{JsiiProperty: "bucketKeyEnabled", GoGetter: "BucketKeyEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "bucketKeyEnabledInput", GoGetter: "BucketKeyEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "cacheControl", GoGetter: "CacheControl"},
			_jsii_.MemberProperty{JsiiProperty: "cacheControlInput", GoGetter: "CacheControlInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "checksumAlgorithm", GoGetter: "ChecksumAlgorithm"},
			_jsii_.MemberProperty{JsiiProperty: "checksumAlgorithmInput", GoGetter: "ChecksumAlgorithmInput"},
			_jsii_.MemberProperty{JsiiProperty: "checksumCrc32", GoGetter: "ChecksumCrc32"},
			_jsii_.MemberProperty{JsiiProperty: "checksumCrc32C", GoGetter: "ChecksumCrc32C"},
			_jsii_.MemberProperty{JsiiProperty: "checksumSha1", GoGetter: "ChecksumSha1"},
			_jsii_.MemberProperty{JsiiProperty: "checksumSha256", GoGetter: "ChecksumSha256"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "contentDisposition", GoGetter: "ContentDisposition"},
			_jsii_.MemberProperty{JsiiProperty: "contentDispositionInput", GoGetter: "ContentDispositionInput"},
			_jsii_.MemberProperty{JsiiProperty: "contentEncoding", GoGetter: "ContentEncoding"},
			_jsii_.MemberProperty{JsiiProperty: "contentEncodingInput", GoGetter: "ContentEncodingInput"},
			_jsii_.MemberProperty{JsiiProperty: "contentLanguage", GoGetter: "ContentLanguage"},
			_jsii_.MemberProperty{JsiiProperty: "contentLanguageInput", GoGetter: "ContentLanguageInput"},
			_jsii_.MemberProperty{JsiiProperty: "contentType", GoGetter: "ContentType"},
			_jsii_.MemberProperty{JsiiProperty: "contentTypeInput", GoGetter: "ContentTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "copyIfMatch", GoGetter: "CopyIfMatch"},
			_jsii_.MemberProperty{JsiiProperty: "copyIfMatchInput", GoGetter: "CopyIfMatchInput"},
			_jsii_.MemberProperty{JsiiProperty: "copyIfModifiedSince", GoGetter: "CopyIfModifiedSince"},
			_jsii_.MemberProperty{JsiiProperty: "copyIfModifiedSinceInput", GoGetter: "CopyIfModifiedSinceInput"},
			_jsii_.MemberProperty{JsiiProperty: "copyIfNoneMatch", GoGetter: "CopyIfNoneMatch"},
			_jsii_.MemberProperty{JsiiProperty: "copyIfNoneMatchInput", GoGetter: "CopyIfNoneMatchInput"},
			_jsii_.MemberProperty{JsiiProperty: "copyIfUnmodifiedSince", GoGetter: "CopyIfUnmodifiedSince"},
			_jsii_.MemberProperty{JsiiProperty: "copyIfUnmodifiedSinceInput", GoGetter: "CopyIfUnmodifiedSinceInput"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "customerAlgorithm", GoGetter: "CustomerAlgorithm"},
			_jsii_.MemberProperty{JsiiProperty: "customerAlgorithmInput", GoGetter: "CustomerAlgorithmInput"},
			_jsii_.MemberProperty{JsiiProperty: "customerKey", GoGetter: "CustomerKey"},
			_jsii_.MemberProperty{JsiiProperty: "customerKeyInput", GoGetter: "CustomerKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "customerKeyMd5", GoGetter: "CustomerKeyMd5"},
			_jsii_.MemberProperty{JsiiProperty: "customerKeyMd5Input", GoGetter: "CustomerKeyMd5Input"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "etag", GoGetter: "Etag"},
			_jsii_.MemberProperty{JsiiProperty: "expectedBucketOwner", GoGetter: "ExpectedBucketOwner"},
			_jsii_.MemberProperty{JsiiProperty: "expectedBucketOwnerInput", GoGetter: "ExpectedBucketOwnerInput"},
			_jsii_.MemberProperty{JsiiProperty: "expectedSourceBucketOwner", GoGetter: "ExpectedSourceBucketOwner"},
			_jsii_.MemberProperty{JsiiProperty: "expectedSourceBucketOwnerInput", GoGetter: "ExpectedSourceBucketOwnerInput"},
			_jsii_.MemberProperty{JsiiProperty: "expiration", GoGetter: "Expiration"},
			_jsii_.MemberProperty{JsiiProperty: "expires", GoGetter: "Expires"},
			_jsii_.MemberProperty{JsiiProperty: "expiresInput", GoGetter: "ExpiresInput"},
			_jsii_.MemberProperty{JsiiProperty: "forceDestroy", GoGetter: "ForceDestroy"},
			_jsii_.MemberProperty{JsiiProperty: "forceDestroyInput", GoGetter: "ForceDestroyInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "grant", GoGetter: "Grant"},
			_jsii_.MemberProperty{JsiiProperty: "grantInput", GoGetter: "GrantInput"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "key", GoGetter: "Key"},
			_jsii_.MemberProperty{JsiiProperty: "keyInput", GoGetter: "KeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "kmsEncryptionContext", GoGetter: "KmsEncryptionContext"},
			_jsii_.MemberProperty{JsiiProperty: "kmsEncryptionContextInput", GoGetter: "KmsEncryptionContextInput"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyId", GoGetter: "KmsKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyIdInput", GoGetter: "KmsKeyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "lastModified", GoGetter: "LastModified"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "metadataDirective", GoGetter: "MetadataDirective"},
			_jsii_.MemberProperty{JsiiProperty: "metadataDirectiveInput", GoGetter: "MetadataDirectiveInput"},
			_jsii_.MemberProperty{JsiiProperty: "metadataInput", GoGetter: "MetadataInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "objectLockLegalHoldStatus", GoGetter: "ObjectLockLegalHoldStatus"},
			_jsii_.MemberProperty{JsiiProperty: "objectLockLegalHoldStatusInput", GoGetter: "ObjectLockLegalHoldStatusInput"},
			_jsii_.MemberProperty{JsiiProperty: "objectLockMode", GoGetter: "ObjectLockMode"},
			_jsii_.MemberProperty{JsiiProperty: "objectLockModeInput", GoGetter: "ObjectLockModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "objectLockRetainUntilDate", GoGetter: "ObjectLockRetainUntilDate"},
			_jsii_.MemberProperty{JsiiProperty: "objectLockRetainUntilDateInput", GoGetter: "ObjectLockRetainUntilDateInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putGrant", GoMethod: "PutGrant"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "requestCharged", GoGetter: "RequestCharged"},
			_jsii_.MemberProperty{JsiiProperty: "requestPayer", GoGetter: "RequestPayer"},
			_jsii_.MemberProperty{JsiiProperty: "requestPayerInput", GoGetter: "RequestPayerInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAcl", GoMethod: "ResetAcl"},
			_jsii_.MemberMethod{JsiiMethod: "resetBucketKeyEnabled", GoMethod: "ResetBucketKeyEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetCacheControl", GoMethod: "ResetCacheControl"},
			_jsii_.MemberMethod{JsiiMethod: "resetChecksumAlgorithm", GoMethod: "ResetChecksumAlgorithm"},
			_jsii_.MemberMethod{JsiiMethod: "resetContentDisposition", GoMethod: "ResetContentDisposition"},
			_jsii_.MemberMethod{JsiiMethod: "resetContentEncoding", GoMethod: "ResetContentEncoding"},
			_jsii_.MemberMethod{JsiiMethod: "resetContentLanguage", GoMethod: "ResetContentLanguage"},
			_jsii_.MemberMethod{JsiiMethod: "resetContentType", GoMethod: "ResetContentType"},
			_jsii_.MemberMethod{JsiiMethod: "resetCopyIfMatch", GoMethod: "ResetCopyIfMatch"},
			_jsii_.MemberMethod{JsiiMethod: "resetCopyIfModifiedSince", GoMethod: "ResetCopyIfModifiedSince"},
			_jsii_.MemberMethod{JsiiMethod: "resetCopyIfNoneMatch", GoMethod: "ResetCopyIfNoneMatch"},
			_jsii_.MemberMethod{JsiiMethod: "resetCopyIfUnmodifiedSince", GoMethod: "ResetCopyIfUnmodifiedSince"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomerAlgorithm", GoMethod: "ResetCustomerAlgorithm"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomerKey", GoMethod: "ResetCustomerKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomerKeyMd5", GoMethod: "ResetCustomerKeyMd5"},
			_jsii_.MemberMethod{JsiiMethod: "resetExpectedBucketOwner", GoMethod: "ResetExpectedBucketOwner"},
			_jsii_.MemberMethod{JsiiMethod: "resetExpectedSourceBucketOwner", GoMethod: "ResetExpectedSourceBucketOwner"},
			_jsii_.MemberMethod{JsiiMethod: "resetExpires", GoMethod: "ResetExpires"},
			_jsii_.MemberMethod{JsiiMethod: "resetForceDestroy", GoMethod: "ResetForceDestroy"},
			_jsii_.MemberMethod{JsiiMethod: "resetGrant", GoMethod: "ResetGrant"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetKmsEncryptionContext", GoMethod: "ResetKmsEncryptionContext"},
			_jsii_.MemberMethod{JsiiMethod: "resetKmsKeyId", GoMethod: "ResetKmsKeyId"},
			_jsii_.MemberMethod{JsiiMethod: "resetMetadata", GoMethod: "ResetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "resetMetadataDirective", GoMethod: "ResetMetadataDirective"},
			_jsii_.MemberMethod{JsiiMethod: "resetObjectLockLegalHoldStatus", GoMethod: "ResetObjectLockLegalHoldStatus"},
			_jsii_.MemberMethod{JsiiMethod: "resetObjectLockMode", GoMethod: "ResetObjectLockMode"},
			_jsii_.MemberMethod{JsiiMethod: "resetObjectLockRetainUntilDate", GoMethod: "ResetObjectLockRetainUntilDate"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequestPayer", GoMethod: "ResetRequestPayer"},
			_jsii_.MemberMethod{JsiiMethod: "resetServerSideEncryption", GoMethod: "ResetServerSideEncryption"},
			_jsii_.MemberMethod{JsiiMethod: "resetSourceCustomerAlgorithm", GoMethod: "ResetSourceCustomerAlgorithm"},
			_jsii_.MemberMethod{JsiiMethod: "resetSourceCustomerKey", GoMethod: "ResetSourceCustomerKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetSourceCustomerKeyMd5", GoMethod: "ResetSourceCustomerKeyMd5"},
			_jsii_.MemberMethod{JsiiMethod: "resetStorageClass", GoMethod: "ResetStorageClass"},
			_jsii_.MemberMethod{JsiiMethod: "resetTaggingDirective", GoMethod: "ResetTaggingDirective"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTagsAll", GoMethod: "ResetTagsAll"},
			_jsii_.MemberMethod{JsiiMethod: "resetWebsiteRedirect", GoMethod: "ResetWebsiteRedirect"},
			_jsii_.MemberProperty{JsiiProperty: "serverSideEncryption", GoGetter: "ServerSideEncryption"},
			_jsii_.MemberProperty{JsiiProperty: "serverSideEncryptionInput", GoGetter: "ServerSideEncryptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "source", GoGetter: "Source"},
			_jsii_.MemberProperty{JsiiProperty: "sourceCustomerAlgorithm", GoGetter: "SourceCustomerAlgorithm"},
			_jsii_.MemberProperty{JsiiProperty: "sourceCustomerAlgorithmInput", GoGetter: "SourceCustomerAlgorithmInput"},
			_jsii_.MemberProperty{JsiiProperty: "sourceCustomerKey", GoGetter: "SourceCustomerKey"},
			_jsii_.MemberProperty{JsiiProperty: "sourceCustomerKeyInput", GoGetter: "SourceCustomerKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "sourceCustomerKeyMd5", GoGetter: "SourceCustomerKeyMd5"},
			_jsii_.MemberProperty{JsiiProperty: "sourceCustomerKeyMd5Input", GoGetter: "SourceCustomerKeyMd5Input"},
			_jsii_.MemberProperty{JsiiProperty: "sourceInput", GoGetter: "SourceInput"},
			_jsii_.MemberProperty{JsiiProperty: "sourceVersionId", GoGetter: "SourceVersionId"},
			_jsii_.MemberProperty{JsiiProperty: "storageClass", GoGetter: "StorageClass"},
			_jsii_.MemberProperty{JsiiProperty: "storageClassInput", GoGetter: "StorageClassInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "taggingDirective", GoGetter: "TaggingDirective"},
			_jsii_.MemberProperty{JsiiProperty: "taggingDirectiveInput", GoGetter: "TaggingDirectiveInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "versionId", GoGetter: "VersionId"},
			_jsii_.MemberProperty{JsiiProperty: "websiteRedirect", GoGetter: "WebsiteRedirect"},
			_jsii_.MemberProperty{JsiiProperty: "websiteRedirectInput", GoGetter: "WebsiteRedirectInput"},
		},
		func() interface{} {
			j := jsiiProxy_S3ObjectCopy{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws.s3ObjectCopy.S3ObjectCopyConfig",
		reflect.TypeOf((*S3ObjectCopyConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws.s3ObjectCopy.S3ObjectCopyGrant",
		reflect.TypeOf((*S3ObjectCopyGrant)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws.s3ObjectCopy.S3ObjectCopyGrantList",
		reflect.TypeOf((*S3ObjectCopyGrantList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_S3ObjectCopyGrantList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws.s3ObjectCopy.S3ObjectCopyGrantOutputReference",
		reflect.TypeOf((*S3ObjectCopyGrantOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "email", GoGetter: "Email"},
			_jsii_.MemberProperty{JsiiProperty: "emailInput", GoGetter: "EmailInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "permissions", GoGetter: "Permissions"},
			_jsii_.MemberProperty{JsiiProperty: "permissionsInput", GoGetter: "PermissionsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetEmail", GoMethod: "ResetEmail"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetUri", GoMethod: "ResetUri"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "uri", GoGetter: "Uri"},
			_jsii_.MemberProperty{JsiiProperty: "uriInput", GoGetter: "UriInput"},
		},
		func() interface{} {
			j := jsiiProxy_S3ObjectCopyGrantOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
