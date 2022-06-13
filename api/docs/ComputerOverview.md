# ComputerOverview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Location** | Pointer to [**ComputerLocation**](ComputerLocation.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Udid** | Pointer to **string** |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**LastContactDate** | Pointer to **string** |  | [optional] 
**LastReportDate** | Pointer to **string** |  | [optional] 
**LastEnrolledDate** | Pointer to **string** |  | [optional] 
**OperatingSystemVersion** | Pointer to **string** |  | [optional] 
**OperatingSystemBuild** | Pointer to **string** |  | [optional] 
**IpAddress** | Pointer to **string** |  | [optional] 
**MacAddress** | Pointer to **string** |  | [optional] 
**AssetTag** | Pointer to **string** |  | [optional] 
**ModelIdentifier** | Pointer to **string** |  | [optional] 
**MdmAccessRights** | Pointer to **int32** |  | [optional] 
**IsManaged** | Pointer to **bool** |  | [optional] 
**ManagementId** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewComputerOverview

`func NewComputerOverview() *ComputerOverview`

NewComputerOverview instantiates a new ComputerOverview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputerOverviewWithDefaults

`func NewComputerOverviewWithDefaults() *ComputerOverview`

NewComputerOverviewWithDefaults instantiates a new ComputerOverview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ComputerOverview) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ComputerOverview) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ComputerOverview) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ComputerOverview) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocation

`func (o *ComputerOverview) GetLocation() ComputerLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ComputerOverview) GetLocationOk() (*ComputerLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ComputerOverview) SetLocation(v ComputerLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ComputerOverview) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetName

`func (o *ComputerOverview) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComputerOverview) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComputerOverview) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ComputerOverview) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUdid

`func (o *ComputerOverview) GetUdid() string`

GetUdid returns the Udid field if non-nil, zero value otherwise.

### GetUdidOk

`func (o *ComputerOverview) GetUdidOk() (*string, bool)`

GetUdidOk returns a tuple with the Udid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdid

`func (o *ComputerOverview) SetUdid(v string)`

SetUdid sets Udid field to given value.

### HasUdid

`func (o *ComputerOverview) HasUdid() bool`

HasUdid returns a boolean if a field has been set.

### GetSerialNumber

`func (o *ComputerOverview) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *ComputerOverview) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *ComputerOverview) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *ComputerOverview) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetLastContactDate

`func (o *ComputerOverview) GetLastContactDate() string`

GetLastContactDate returns the LastContactDate field if non-nil, zero value otherwise.

### GetLastContactDateOk

`func (o *ComputerOverview) GetLastContactDateOk() (*string, bool)`

GetLastContactDateOk returns a tuple with the LastContactDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastContactDate

`func (o *ComputerOverview) SetLastContactDate(v string)`

SetLastContactDate sets LastContactDate field to given value.

### HasLastContactDate

`func (o *ComputerOverview) HasLastContactDate() bool`

HasLastContactDate returns a boolean if a field has been set.

### GetLastReportDate

`func (o *ComputerOverview) GetLastReportDate() string`

GetLastReportDate returns the LastReportDate field if non-nil, zero value otherwise.

### GetLastReportDateOk

`func (o *ComputerOverview) GetLastReportDateOk() (*string, bool)`

GetLastReportDateOk returns a tuple with the LastReportDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReportDate

`func (o *ComputerOverview) SetLastReportDate(v string)`

SetLastReportDate sets LastReportDate field to given value.

### HasLastReportDate

`func (o *ComputerOverview) HasLastReportDate() bool`

HasLastReportDate returns a boolean if a field has been set.

### GetLastEnrolledDate

`func (o *ComputerOverview) GetLastEnrolledDate() string`

GetLastEnrolledDate returns the LastEnrolledDate field if non-nil, zero value otherwise.

### GetLastEnrolledDateOk

`func (o *ComputerOverview) GetLastEnrolledDateOk() (*string, bool)`

GetLastEnrolledDateOk returns a tuple with the LastEnrolledDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEnrolledDate

`func (o *ComputerOverview) SetLastEnrolledDate(v string)`

SetLastEnrolledDate sets LastEnrolledDate field to given value.

### HasLastEnrolledDate

`func (o *ComputerOverview) HasLastEnrolledDate() bool`

HasLastEnrolledDate returns a boolean if a field has been set.

### GetOperatingSystemVersion

`func (o *ComputerOverview) GetOperatingSystemVersion() string`

GetOperatingSystemVersion returns the OperatingSystemVersion field if non-nil, zero value otherwise.

### GetOperatingSystemVersionOk

`func (o *ComputerOverview) GetOperatingSystemVersionOk() (*string, bool)`

GetOperatingSystemVersionOk returns a tuple with the OperatingSystemVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystemVersion

`func (o *ComputerOverview) SetOperatingSystemVersion(v string)`

SetOperatingSystemVersion sets OperatingSystemVersion field to given value.

### HasOperatingSystemVersion

`func (o *ComputerOverview) HasOperatingSystemVersion() bool`

HasOperatingSystemVersion returns a boolean if a field has been set.

### GetOperatingSystemBuild

`func (o *ComputerOverview) GetOperatingSystemBuild() string`

GetOperatingSystemBuild returns the OperatingSystemBuild field if non-nil, zero value otherwise.

### GetOperatingSystemBuildOk

`func (o *ComputerOverview) GetOperatingSystemBuildOk() (*string, bool)`

GetOperatingSystemBuildOk returns a tuple with the OperatingSystemBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystemBuild

`func (o *ComputerOverview) SetOperatingSystemBuild(v string)`

SetOperatingSystemBuild sets OperatingSystemBuild field to given value.

### HasOperatingSystemBuild

`func (o *ComputerOverview) HasOperatingSystemBuild() bool`

HasOperatingSystemBuild returns a boolean if a field has been set.

### GetIpAddress

`func (o *ComputerOverview) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *ComputerOverview) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *ComputerOverview) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *ComputerOverview) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetMacAddress

`func (o *ComputerOverview) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *ComputerOverview) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *ComputerOverview) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *ComputerOverview) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetAssetTag

`func (o *ComputerOverview) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *ComputerOverview) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *ComputerOverview) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *ComputerOverview) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### GetModelIdentifier

`func (o *ComputerOverview) GetModelIdentifier() string`

GetModelIdentifier returns the ModelIdentifier field if non-nil, zero value otherwise.

### GetModelIdentifierOk

`func (o *ComputerOverview) GetModelIdentifierOk() (*string, bool)`

GetModelIdentifierOk returns a tuple with the ModelIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelIdentifier

`func (o *ComputerOverview) SetModelIdentifier(v string)`

SetModelIdentifier sets ModelIdentifier field to given value.

### HasModelIdentifier

`func (o *ComputerOverview) HasModelIdentifier() bool`

HasModelIdentifier returns a boolean if a field has been set.

### GetMdmAccessRights

`func (o *ComputerOverview) GetMdmAccessRights() int32`

GetMdmAccessRights returns the MdmAccessRights field if non-nil, zero value otherwise.

### GetMdmAccessRightsOk

`func (o *ComputerOverview) GetMdmAccessRightsOk() (*int32, bool)`

GetMdmAccessRightsOk returns a tuple with the MdmAccessRights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdmAccessRights

`func (o *ComputerOverview) SetMdmAccessRights(v int32)`

SetMdmAccessRights sets MdmAccessRights field to given value.

### HasMdmAccessRights

`func (o *ComputerOverview) HasMdmAccessRights() bool`

HasMdmAccessRights returns a boolean if a field has been set.

### GetIsManaged

`func (o *ComputerOverview) GetIsManaged() bool`

GetIsManaged returns the IsManaged field if non-nil, zero value otherwise.

### GetIsManagedOk

`func (o *ComputerOverview) GetIsManagedOk() (*bool, bool)`

GetIsManagedOk returns a tuple with the IsManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManaged

`func (o *ComputerOverview) SetIsManaged(v bool)`

SetIsManaged sets IsManaged field to given value.

### HasIsManaged

`func (o *ComputerOverview) HasIsManaged() bool`

HasIsManaged returns a boolean if a field has been set.

### GetManagementId

`func (o *ComputerOverview) GetManagementId() string`

GetManagementId returns the ManagementId field if non-nil, zero value otherwise.

### GetManagementIdOk

`func (o *ComputerOverview) GetManagementIdOk() (*string, bool)`

GetManagementIdOk returns a tuple with the ManagementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementId

`func (o *ComputerOverview) SetManagementId(v string)`

SetManagementId sets ManagementId field to given value.

### HasManagementId

`func (o *ComputerOverview) HasManagementId() bool`

HasManagementId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


