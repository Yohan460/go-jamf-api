# ActivePatchHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PatchId** | Pointer to **int32** |  | [optional] 
**PatchHistoryId** | Pointer to **int32** |  | [optional] 
**DeviceId** | Pointer to **int32** |  | [optional] 
**DeviceName** | Pointer to **string** |  | [optional] 
**BuildingId** | Pointer to **int32** |  | [optional] 
**BuildingName** | Pointer to **string** |  | [optional] 
**DepartmentId** | Pointer to **int32** |  | [optional] 
**DepartmentName** | Pointer to **string** |  | [optional] 
**SiteId** | Pointer to **int32** |  | [optional] 
**SiteName** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**OsVersion** | Pointer to **string** |  | [optional] 
**LastCheckIn** | Pointer to **time.Time** |  | [optional] 
**InstalledVersion** | Pointer to **string** |  | [optional] 

## Methods

### NewActivePatchHistory

`func NewActivePatchHistory() *ActivePatchHistory`

NewActivePatchHistory instantiates a new ActivePatchHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivePatchHistoryWithDefaults

`func NewActivePatchHistoryWithDefaults() *ActivePatchHistory`

NewActivePatchHistoryWithDefaults instantiates a new ActivePatchHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPatchId

`func (o *ActivePatchHistory) GetPatchId() int32`

GetPatchId returns the PatchId field if non-nil, zero value otherwise.

### GetPatchIdOk

`func (o *ActivePatchHistory) GetPatchIdOk() (*int32, bool)`

GetPatchIdOk returns a tuple with the PatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchId

`func (o *ActivePatchHistory) SetPatchId(v int32)`

SetPatchId sets PatchId field to given value.

### HasPatchId

`func (o *ActivePatchHistory) HasPatchId() bool`

HasPatchId returns a boolean if a field has been set.

### GetPatchHistoryId

`func (o *ActivePatchHistory) GetPatchHistoryId() int32`

GetPatchHistoryId returns the PatchHistoryId field if non-nil, zero value otherwise.

### GetPatchHistoryIdOk

`func (o *ActivePatchHistory) GetPatchHistoryIdOk() (*int32, bool)`

GetPatchHistoryIdOk returns a tuple with the PatchHistoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchHistoryId

`func (o *ActivePatchHistory) SetPatchHistoryId(v int32)`

SetPatchHistoryId sets PatchHistoryId field to given value.

### HasPatchHistoryId

`func (o *ActivePatchHistory) HasPatchHistoryId() bool`

HasPatchHistoryId returns a boolean if a field has been set.

### GetDeviceId

`func (o *ActivePatchHistory) GetDeviceId() int32`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *ActivePatchHistory) GetDeviceIdOk() (*int32, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *ActivePatchHistory) SetDeviceId(v int32)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *ActivePatchHistory) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetDeviceName

`func (o *ActivePatchHistory) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *ActivePatchHistory) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *ActivePatchHistory) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *ActivePatchHistory) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetBuildingId

`func (o *ActivePatchHistory) GetBuildingId() int32`

GetBuildingId returns the BuildingId field if non-nil, zero value otherwise.

### GetBuildingIdOk

`func (o *ActivePatchHistory) GetBuildingIdOk() (*int32, bool)`

GetBuildingIdOk returns a tuple with the BuildingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildingId

`func (o *ActivePatchHistory) SetBuildingId(v int32)`

SetBuildingId sets BuildingId field to given value.

### HasBuildingId

`func (o *ActivePatchHistory) HasBuildingId() bool`

HasBuildingId returns a boolean if a field has been set.

### GetBuildingName

`func (o *ActivePatchHistory) GetBuildingName() string`

GetBuildingName returns the BuildingName field if non-nil, zero value otherwise.

### GetBuildingNameOk

`func (o *ActivePatchHistory) GetBuildingNameOk() (*string, bool)`

GetBuildingNameOk returns a tuple with the BuildingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildingName

`func (o *ActivePatchHistory) SetBuildingName(v string)`

SetBuildingName sets BuildingName field to given value.

### HasBuildingName

`func (o *ActivePatchHistory) HasBuildingName() bool`

HasBuildingName returns a boolean if a field has been set.

### GetDepartmentId

`func (o *ActivePatchHistory) GetDepartmentId() int32`

GetDepartmentId returns the DepartmentId field if non-nil, zero value otherwise.

### GetDepartmentIdOk

`func (o *ActivePatchHistory) GetDepartmentIdOk() (*int32, bool)`

GetDepartmentIdOk returns a tuple with the DepartmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartmentId

`func (o *ActivePatchHistory) SetDepartmentId(v int32)`

SetDepartmentId sets DepartmentId field to given value.

### HasDepartmentId

`func (o *ActivePatchHistory) HasDepartmentId() bool`

HasDepartmentId returns a boolean if a field has been set.

### GetDepartmentName

`func (o *ActivePatchHistory) GetDepartmentName() string`

GetDepartmentName returns the DepartmentName field if non-nil, zero value otherwise.

### GetDepartmentNameOk

`func (o *ActivePatchHistory) GetDepartmentNameOk() (*string, bool)`

GetDepartmentNameOk returns a tuple with the DepartmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartmentName

`func (o *ActivePatchHistory) SetDepartmentName(v string)`

SetDepartmentName sets DepartmentName field to given value.

### HasDepartmentName

`func (o *ActivePatchHistory) HasDepartmentName() bool`

HasDepartmentName returns a boolean if a field has been set.

### GetSiteId

`func (o *ActivePatchHistory) GetSiteId() int32`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *ActivePatchHistory) GetSiteIdOk() (*int32, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *ActivePatchHistory) SetSiteId(v int32)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *ActivePatchHistory) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSiteName

`func (o *ActivePatchHistory) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *ActivePatchHistory) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *ActivePatchHistory) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *ActivePatchHistory) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetUsername

`func (o *ActivePatchHistory) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ActivePatchHistory) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ActivePatchHistory) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ActivePatchHistory) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetOsVersion

`func (o *ActivePatchHistory) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *ActivePatchHistory) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *ActivePatchHistory) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *ActivePatchHistory) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetLastCheckIn

`func (o *ActivePatchHistory) GetLastCheckIn() time.Time`

GetLastCheckIn returns the LastCheckIn field if non-nil, zero value otherwise.

### GetLastCheckInOk

`func (o *ActivePatchHistory) GetLastCheckInOk() (*time.Time, bool)`

GetLastCheckInOk returns a tuple with the LastCheckIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckIn

`func (o *ActivePatchHistory) SetLastCheckIn(v time.Time)`

SetLastCheckIn sets LastCheckIn field to given value.

### HasLastCheckIn

`func (o *ActivePatchHistory) HasLastCheckIn() bool`

HasLastCheckIn returns a boolean if a field has been set.

### GetInstalledVersion

`func (o *ActivePatchHistory) GetInstalledVersion() string`

GetInstalledVersion returns the InstalledVersion field if non-nil, zero value otherwise.

### GetInstalledVersionOk

`func (o *ActivePatchHistory) GetInstalledVersionOk() (*string, bool)`

GetInstalledVersionOk returns a tuple with the InstalledVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledVersion

`func (o *ActivePatchHistory) SetInstalledVersion(v string)`

SetInstalledVersion sets InstalledVersion field to given value.

### HasInstalledVersion

`func (o *ActivePatchHistory) HasInstalledVersion() bool`

HasInstalledVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


