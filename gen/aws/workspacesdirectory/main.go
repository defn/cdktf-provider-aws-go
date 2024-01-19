package workspacesdirectory

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws.workspacesDirectory.WorkspacesDirectory",
		reflect.TypeOf((*WorkspacesDirectory)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "alias", GoGetter: "Alias"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "customerUserName", GoGetter: "CustomerUserName"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "directoryId", GoGetter: "DirectoryId"},
			_jsii_.MemberProperty{JsiiProperty: "directoryIdInput", GoGetter: "DirectoryIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "directoryName", GoGetter: "DirectoryName"},
			_jsii_.MemberProperty{JsiiProperty: "directoryType", GoGetter: "DirectoryType"},
			_jsii_.MemberProperty{JsiiProperty: "dnsIpAddresses", GoGetter: "DnsIpAddresses"},
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
			_jsii_.MemberProperty{JsiiProperty: "iamRoleId", GoGetter: "IamRoleId"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ipGroupIds", GoGetter: "IpGroupIds"},
			_jsii_.MemberProperty{JsiiProperty: "ipGroupIdsInput", GoGetter: "IpGroupIdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putSelfServicePermissions", GoMethod: "PutSelfServicePermissions"},
			_jsii_.MemberMethod{JsiiMethod: "putWorkspaceAccessProperties", GoMethod: "PutWorkspaceAccessProperties"},
			_jsii_.MemberMethod{JsiiMethod: "putWorkspaceCreationProperties", GoMethod: "PutWorkspaceCreationProperties"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "registrationCode", GoGetter: "RegistrationCode"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpGroupIds", GoMethod: "ResetIpGroupIds"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSelfServicePermissions", GoMethod: "ResetSelfServicePermissions"},
			_jsii_.MemberMethod{JsiiMethod: "resetSubnetIds", GoMethod: "ResetSubnetIds"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTagsAll", GoMethod: "ResetTagsAll"},
			_jsii_.MemberMethod{JsiiMethod: "resetWorkspaceAccessProperties", GoMethod: "ResetWorkspaceAccessProperties"},
			_jsii_.MemberMethod{JsiiMethod: "resetWorkspaceCreationProperties", GoMethod: "ResetWorkspaceCreationProperties"},
			_jsii_.MemberProperty{JsiiProperty: "selfServicePermissions", GoGetter: "SelfServicePermissions"},
			_jsii_.MemberProperty{JsiiProperty: "selfServicePermissionsInput", GoGetter: "SelfServicePermissionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "subnetIds", GoGetter: "SubnetIds"},
			_jsii_.MemberProperty{JsiiProperty: "subnetIdsInput", GoGetter: "SubnetIdsInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "workspaceAccessProperties", GoGetter: "WorkspaceAccessProperties"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceAccessPropertiesInput", GoGetter: "WorkspaceAccessPropertiesInput"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceCreationProperties", GoGetter: "WorkspaceCreationProperties"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceCreationPropertiesInput", GoGetter: "WorkspaceCreationPropertiesInput"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceSecurityGroupId", GoGetter: "WorkspaceSecurityGroupId"},
		},
		func() interface{} {
			j := jsiiProxy_WorkspacesDirectory{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws.workspacesDirectory.WorkspacesDirectoryConfig",
		reflect.TypeOf((*WorkspacesDirectoryConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws.workspacesDirectory.WorkspacesDirectorySelfServicePermissions",
		reflect.TypeOf((*WorkspacesDirectorySelfServicePermissions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws.workspacesDirectory.WorkspacesDirectorySelfServicePermissionsOutputReference",
		reflect.TypeOf((*WorkspacesDirectorySelfServicePermissionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "changeComputeType", GoGetter: "ChangeComputeType"},
			_jsii_.MemberProperty{JsiiProperty: "changeComputeTypeInput", GoGetter: "ChangeComputeTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "increaseVolumeSize", GoGetter: "IncreaseVolumeSize"},
			_jsii_.MemberProperty{JsiiProperty: "increaseVolumeSizeInput", GoGetter: "IncreaseVolumeSizeInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "rebuildWorkspace", GoGetter: "RebuildWorkspace"},
			_jsii_.MemberProperty{JsiiProperty: "rebuildWorkspaceInput", GoGetter: "RebuildWorkspaceInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetChangeComputeType", GoMethod: "ResetChangeComputeType"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncreaseVolumeSize", GoMethod: "ResetIncreaseVolumeSize"},
			_jsii_.MemberMethod{JsiiMethod: "resetRebuildWorkspace", GoMethod: "ResetRebuildWorkspace"},
			_jsii_.MemberMethod{JsiiMethod: "resetRestartWorkspace", GoMethod: "ResetRestartWorkspace"},
			_jsii_.MemberMethod{JsiiMethod: "resetSwitchRunningMode", GoMethod: "ResetSwitchRunningMode"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "restartWorkspace", GoGetter: "RestartWorkspace"},
			_jsii_.MemberProperty{JsiiProperty: "restartWorkspaceInput", GoGetter: "RestartWorkspaceInput"},
			_jsii_.MemberProperty{JsiiProperty: "switchRunningMode", GoGetter: "SwitchRunningMode"},
			_jsii_.MemberProperty{JsiiProperty: "switchRunningModeInput", GoGetter: "SwitchRunningModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_WorkspacesDirectorySelfServicePermissionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws.workspacesDirectory.WorkspacesDirectoryWorkspaceAccessProperties",
		reflect.TypeOf((*WorkspacesDirectoryWorkspaceAccessProperties)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws.workspacesDirectory.WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference",
		reflect.TypeOf((*WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "deviceTypeAndroid", GoGetter: "DeviceTypeAndroid"},
			_jsii_.MemberProperty{JsiiProperty: "deviceTypeAndroidInput", GoGetter: "DeviceTypeAndroidInput"},
			_jsii_.MemberProperty{JsiiProperty: "deviceTypeChromeos", GoGetter: "DeviceTypeChromeos"},
			_jsii_.MemberProperty{JsiiProperty: "deviceTypeChromeosInput", GoGetter: "DeviceTypeChromeosInput"},
			_jsii_.MemberProperty{JsiiProperty: "deviceTypeIos", GoGetter: "DeviceTypeIos"},
			_jsii_.MemberProperty{JsiiProperty: "deviceTypeIosInput", GoGetter: "DeviceTypeIosInput"},
			_jsii_.MemberProperty{JsiiProperty: "deviceTypeLinux", GoGetter: "DeviceTypeLinux"},
			_jsii_.MemberProperty{JsiiProperty: "deviceTypeLinuxInput", GoGetter: "DeviceTypeLinuxInput"},
			_jsii_.MemberProperty{JsiiProperty: "deviceTypeOsx", GoGetter: "DeviceTypeOsx"},
			_jsii_.MemberProperty{JsiiProperty: "deviceTypeOsxInput", GoGetter: "DeviceTypeOsxInput"},
			_jsii_.MemberProperty{JsiiProperty: "deviceTypeWeb", GoGetter: "DeviceTypeWeb"},
			_jsii_.MemberProperty{JsiiProperty: "deviceTypeWebInput", GoGetter: "DeviceTypeWebInput"},
			_jsii_.MemberProperty{JsiiProperty: "deviceTypeWindows", GoGetter: "DeviceTypeWindows"},
			_jsii_.MemberProperty{JsiiProperty: "deviceTypeWindowsInput", GoGetter: "DeviceTypeWindowsInput"},
			_jsii_.MemberProperty{JsiiProperty: "deviceTypeZeroclient", GoGetter: "DeviceTypeZeroclient"},
			_jsii_.MemberProperty{JsiiProperty: "deviceTypeZeroclientInput", GoGetter: "DeviceTypeZeroclientInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeviceTypeAndroid", GoMethod: "ResetDeviceTypeAndroid"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeviceTypeChromeos", GoMethod: "ResetDeviceTypeChromeos"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeviceTypeIos", GoMethod: "ResetDeviceTypeIos"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeviceTypeLinux", GoMethod: "ResetDeviceTypeLinux"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeviceTypeOsx", GoMethod: "ResetDeviceTypeOsx"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeviceTypeWeb", GoMethod: "ResetDeviceTypeWeb"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeviceTypeWindows", GoMethod: "ResetDeviceTypeWindows"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeviceTypeZeroclient", GoMethod: "ResetDeviceTypeZeroclient"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws.workspacesDirectory.WorkspacesDirectoryWorkspaceCreationProperties",
		reflect.TypeOf((*WorkspacesDirectoryWorkspaceCreationProperties)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws.workspacesDirectory.WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference",
		reflect.TypeOf((*WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "customSecurityGroupId", GoGetter: "CustomSecurityGroupId"},
			_jsii_.MemberProperty{JsiiProperty: "customSecurityGroupIdInput", GoGetter: "CustomSecurityGroupIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultOu", GoGetter: "DefaultOu"},
			_jsii_.MemberProperty{JsiiProperty: "defaultOuInput", GoGetter: "DefaultOuInput"},
			_jsii_.MemberProperty{JsiiProperty: "enableInternetAccess", GoGetter: "EnableInternetAccess"},
			_jsii_.MemberProperty{JsiiProperty: "enableInternetAccessInput", GoGetter: "EnableInternetAccessInput"},
			_jsii_.MemberProperty{JsiiProperty: "enableMaintenanceMode", GoGetter: "EnableMaintenanceMode"},
			_jsii_.MemberProperty{JsiiProperty: "enableMaintenanceModeInput", GoGetter: "EnableMaintenanceModeInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomSecurityGroupId", GoMethod: "ResetCustomSecurityGroupId"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultOu", GoMethod: "ResetDefaultOu"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableInternetAccess", GoMethod: "ResetEnableInternetAccess"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableMaintenanceMode", GoMethod: "ResetEnableMaintenanceMode"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserEnabledAsLocalAdministrator", GoMethod: "ResetUserEnabledAsLocalAdministrator"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "userEnabledAsLocalAdministrator", GoGetter: "UserEnabledAsLocalAdministrator"},
			_jsii_.MemberProperty{JsiiProperty: "userEnabledAsLocalAdministratorInput", GoGetter: "UserEnabledAsLocalAdministratorInput"},
		},
		func() interface{} {
			j := jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}