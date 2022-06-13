# EbookScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllComputers** | Pointer to **bool** |  | [optional] 
**AllMobileDevices** | Pointer to **bool** |  | [optional] 
**AllUsers** | Pointer to **bool** |  | [optional] 
**ComputerIds** | Pointer to **[]string** |  | [optional] 
**ComputerGroupIds** | Pointer to **[]string** |  | [optional] 
**MobileDeviceIds** | Pointer to **[]string** |  | [optional] 
**MobileDeviceGroupIds** | Pointer to **[]string** |  | [optional] 
**BuildingIds** | Pointer to **[]string** |  | [optional] 
**DepartmentIds** | Pointer to **[]string** |  | [optional] 
**UserIds** | Pointer to **[]string** |  | [optional] 
**UserGroupIds** | Pointer to **[]string** |  | [optional] 
**ClassroomIds** | Pointer to **[]string** |  | [optional] 
**Limitations** | Pointer to [**EbookLimitations**](EbookLimitations.md) |  | [optional] 
**Exclusions** | Pointer to [**EbookExclusions**](EbookExclusions.md) |  | [optional] 

## Methods

### NewEbookScope

`func NewEbookScope() *EbookScope`

NewEbookScope instantiates a new EbookScope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEbookScopeWithDefaults

`func NewEbookScopeWithDefaults() *EbookScope`

NewEbookScopeWithDefaults instantiates a new EbookScope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllComputers

`func (o *EbookScope) GetAllComputers() bool`

GetAllComputers returns the AllComputers field if non-nil, zero value otherwise.

### GetAllComputersOk

`func (o *EbookScope) GetAllComputersOk() (*bool, bool)`

GetAllComputersOk returns a tuple with the AllComputers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllComputers

`func (o *EbookScope) SetAllComputers(v bool)`

SetAllComputers sets AllComputers field to given value.

### HasAllComputers

`func (o *EbookScope) HasAllComputers() bool`

HasAllComputers returns a boolean if a field has been set.

### GetAllMobileDevices

`func (o *EbookScope) GetAllMobileDevices() bool`

GetAllMobileDevices returns the AllMobileDevices field if non-nil, zero value otherwise.

### GetAllMobileDevicesOk

`func (o *EbookScope) GetAllMobileDevicesOk() (*bool, bool)`

GetAllMobileDevicesOk returns a tuple with the AllMobileDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllMobileDevices

`func (o *EbookScope) SetAllMobileDevices(v bool)`

SetAllMobileDevices sets AllMobileDevices field to given value.

### HasAllMobileDevices

`func (o *EbookScope) HasAllMobileDevices() bool`

HasAllMobileDevices returns a boolean if a field has been set.

### GetAllUsers

`func (o *EbookScope) GetAllUsers() bool`

GetAllUsers returns the AllUsers field if non-nil, zero value otherwise.

### GetAllUsersOk

`func (o *EbookScope) GetAllUsersOk() (*bool, bool)`

GetAllUsersOk returns a tuple with the AllUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllUsers

`func (o *EbookScope) SetAllUsers(v bool)`

SetAllUsers sets AllUsers field to given value.

### HasAllUsers

`func (o *EbookScope) HasAllUsers() bool`

HasAllUsers returns a boolean if a field has been set.

### GetComputerIds

`func (o *EbookScope) GetComputerIds() []string`

GetComputerIds returns the ComputerIds field if non-nil, zero value otherwise.

### GetComputerIdsOk

`func (o *EbookScope) GetComputerIdsOk() (*[]string, bool)`

GetComputerIdsOk returns a tuple with the ComputerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputerIds

`func (o *EbookScope) SetComputerIds(v []string)`

SetComputerIds sets ComputerIds field to given value.

### HasComputerIds

`func (o *EbookScope) HasComputerIds() bool`

HasComputerIds returns a boolean if a field has been set.

### GetComputerGroupIds

`func (o *EbookScope) GetComputerGroupIds() []string`

GetComputerGroupIds returns the ComputerGroupIds field if non-nil, zero value otherwise.

### GetComputerGroupIdsOk

`func (o *EbookScope) GetComputerGroupIdsOk() (*[]string, bool)`

GetComputerGroupIdsOk returns a tuple with the ComputerGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputerGroupIds

`func (o *EbookScope) SetComputerGroupIds(v []string)`

SetComputerGroupIds sets ComputerGroupIds field to given value.

### HasComputerGroupIds

`func (o *EbookScope) HasComputerGroupIds() bool`

HasComputerGroupIds returns a boolean if a field has been set.

### GetMobileDeviceIds

`func (o *EbookScope) GetMobileDeviceIds() []string`

GetMobileDeviceIds returns the MobileDeviceIds field if non-nil, zero value otherwise.

### GetMobileDeviceIdsOk

`func (o *EbookScope) GetMobileDeviceIdsOk() (*[]string, bool)`

GetMobileDeviceIdsOk returns a tuple with the MobileDeviceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileDeviceIds

`func (o *EbookScope) SetMobileDeviceIds(v []string)`

SetMobileDeviceIds sets MobileDeviceIds field to given value.

### HasMobileDeviceIds

`func (o *EbookScope) HasMobileDeviceIds() bool`

HasMobileDeviceIds returns a boolean if a field has been set.

### GetMobileDeviceGroupIds

`func (o *EbookScope) GetMobileDeviceGroupIds() []string`

GetMobileDeviceGroupIds returns the MobileDeviceGroupIds field if non-nil, zero value otherwise.

### GetMobileDeviceGroupIdsOk

`func (o *EbookScope) GetMobileDeviceGroupIdsOk() (*[]string, bool)`

GetMobileDeviceGroupIdsOk returns a tuple with the MobileDeviceGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileDeviceGroupIds

`func (o *EbookScope) SetMobileDeviceGroupIds(v []string)`

SetMobileDeviceGroupIds sets MobileDeviceGroupIds field to given value.

### HasMobileDeviceGroupIds

`func (o *EbookScope) HasMobileDeviceGroupIds() bool`

HasMobileDeviceGroupIds returns a boolean if a field has been set.

### GetBuildingIds

`func (o *EbookScope) GetBuildingIds() []string`

GetBuildingIds returns the BuildingIds field if non-nil, zero value otherwise.

### GetBuildingIdsOk

`func (o *EbookScope) GetBuildingIdsOk() (*[]string, bool)`

GetBuildingIdsOk returns a tuple with the BuildingIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildingIds

`func (o *EbookScope) SetBuildingIds(v []string)`

SetBuildingIds sets BuildingIds field to given value.

### HasBuildingIds

`func (o *EbookScope) HasBuildingIds() bool`

HasBuildingIds returns a boolean if a field has been set.

### GetDepartmentIds

`func (o *EbookScope) GetDepartmentIds() []string`

GetDepartmentIds returns the DepartmentIds field if non-nil, zero value otherwise.

### GetDepartmentIdsOk

`func (o *EbookScope) GetDepartmentIdsOk() (*[]string, bool)`

GetDepartmentIdsOk returns a tuple with the DepartmentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartmentIds

`func (o *EbookScope) SetDepartmentIds(v []string)`

SetDepartmentIds sets DepartmentIds field to given value.

### HasDepartmentIds

`func (o *EbookScope) HasDepartmentIds() bool`

HasDepartmentIds returns a boolean if a field has been set.

### GetUserIds

`func (o *EbookScope) GetUserIds() []string`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *EbookScope) GetUserIdsOk() (*[]string, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *EbookScope) SetUserIds(v []string)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *EbookScope) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.

### GetUserGroupIds

`func (o *EbookScope) GetUserGroupIds() []string`

GetUserGroupIds returns the UserGroupIds field if non-nil, zero value otherwise.

### GetUserGroupIdsOk

`func (o *EbookScope) GetUserGroupIdsOk() (*[]string, bool)`

GetUserGroupIdsOk returns a tuple with the UserGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupIds

`func (o *EbookScope) SetUserGroupIds(v []string)`

SetUserGroupIds sets UserGroupIds field to given value.

### HasUserGroupIds

`func (o *EbookScope) HasUserGroupIds() bool`

HasUserGroupIds returns a boolean if a field has been set.

### GetClassroomIds

`func (o *EbookScope) GetClassroomIds() []string`

GetClassroomIds returns the ClassroomIds field if non-nil, zero value otherwise.

### GetClassroomIdsOk

`func (o *EbookScope) GetClassroomIdsOk() (*[]string, bool)`

GetClassroomIdsOk returns a tuple with the ClassroomIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassroomIds

`func (o *EbookScope) SetClassroomIds(v []string)`

SetClassroomIds sets ClassroomIds field to given value.

### HasClassroomIds

`func (o *EbookScope) HasClassroomIds() bool`

HasClassroomIds returns a boolean if a field has been set.

### GetLimitations

`func (o *EbookScope) GetLimitations() EbookLimitations`

GetLimitations returns the Limitations field if non-nil, zero value otherwise.

### GetLimitationsOk

`func (o *EbookScope) GetLimitationsOk() (*EbookLimitations, bool)`

GetLimitationsOk returns a tuple with the Limitations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitations

`func (o *EbookScope) SetLimitations(v EbookLimitations)`

SetLimitations sets Limitations field to given value.

### HasLimitations

`func (o *EbookScope) HasLimitations() bool`

HasLimitations returns a boolean if a field has been set.

### GetExclusions

`func (o *EbookScope) GetExclusions() EbookExclusions`

GetExclusions returns the Exclusions field if non-nil, zero value otherwise.

### GetExclusionsOk

`func (o *EbookScope) GetExclusionsOk() (*EbookExclusions, bool)`

GetExclusionsOk returns a tuple with the Exclusions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusions

`func (o *EbookScope) SetExclusions(v EbookExclusions)`

SetExclusions sets Exclusions field to given value.

### HasExclusions

`func (o *EbookScope) HasExclusions() bool`

HasExclusions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


