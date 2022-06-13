# DeviceEnrollmentDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DeviceEnrollmentProgramInstanceId** | Pointer to **string** |  | [optional] 
**PrestageId** | Pointer to **string** |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Color** | Pointer to **string** |  | [optional] 
**AssetTag** | Pointer to **string** |  | [optional] 
**ProfileStatus** | Pointer to **string** |  | [optional] 
**SyncState** | Pointer to [**AssignRemoveProfileResponseSyncState**](AssignRemoveProfileResponseSyncState.md) |  | [optional] 
**ProfileAssignTime** | Pointer to **string** |  | [optional] 
**ProfilePushTime** | Pointer to **string** |  | [optional] 
**DeviceAssignedDate** | Pointer to **string** |  | [optional] 

## Methods

### NewDeviceEnrollmentDevice

`func NewDeviceEnrollmentDevice() *DeviceEnrollmentDevice`

NewDeviceEnrollmentDevice instantiates a new DeviceEnrollmentDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceEnrollmentDeviceWithDefaults

`func NewDeviceEnrollmentDeviceWithDefaults() *DeviceEnrollmentDevice`

NewDeviceEnrollmentDeviceWithDefaults instantiates a new DeviceEnrollmentDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeviceEnrollmentDevice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceEnrollmentDevice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceEnrollmentDevice) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeviceEnrollmentDevice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeviceEnrollmentProgramInstanceId

`func (o *DeviceEnrollmentDevice) GetDeviceEnrollmentProgramInstanceId() string`

GetDeviceEnrollmentProgramInstanceId returns the DeviceEnrollmentProgramInstanceId field if non-nil, zero value otherwise.

### GetDeviceEnrollmentProgramInstanceIdOk

`func (o *DeviceEnrollmentDevice) GetDeviceEnrollmentProgramInstanceIdOk() (*string, bool)`

GetDeviceEnrollmentProgramInstanceIdOk returns a tuple with the DeviceEnrollmentProgramInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceEnrollmentProgramInstanceId

`func (o *DeviceEnrollmentDevice) SetDeviceEnrollmentProgramInstanceId(v string)`

SetDeviceEnrollmentProgramInstanceId sets DeviceEnrollmentProgramInstanceId field to given value.

### HasDeviceEnrollmentProgramInstanceId

`func (o *DeviceEnrollmentDevice) HasDeviceEnrollmentProgramInstanceId() bool`

HasDeviceEnrollmentProgramInstanceId returns a boolean if a field has been set.

### GetPrestageId

`func (o *DeviceEnrollmentDevice) GetPrestageId() string`

GetPrestageId returns the PrestageId field if non-nil, zero value otherwise.

### GetPrestageIdOk

`func (o *DeviceEnrollmentDevice) GetPrestageIdOk() (*string, bool)`

GetPrestageIdOk returns a tuple with the PrestageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrestageId

`func (o *DeviceEnrollmentDevice) SetPrestageId(v string)`

SetPrestageId sets PrestageId field to given value.

### HasPrestageId

`func (o *DeviceEnrollmentDevice) HasPrestageId() bool`

HasPrestageId returns a boolean if a field has been set.

### GetSerialNumber

`func (o *DeviceEnrollmentDevice) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *DeviceEnrollmentDevice) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *DeviceEnrollmentDevice) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *DeviceEnrollmentDevice) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetDescription

`func (o *DeviceEnrollmentDevice) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DeviceEnrollmentDevice) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DeviceEnrollmentDevice) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DeviceEnrollmentDevice) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetModel

`func (o *DeviceEnrollmentDevice) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *DeviceEnrollmentDevice) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *DeviceEnrollmentDevice) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *DeviceEnrollmentDevice) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetColor

`func (o *DeviceEnrollmentDevice) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *DeviceEnrollmentDevice) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *DeviceEnrollmentDevice) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *DeviceEnrollmentDevice) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetAssetTag

`func (o *DeviceEnrollmentDevice) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *DeviceEnrollmentDevice) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *DeviceEnrollmentDevice) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *DeviceEnrollmentDevice) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### GetProfileStatus

`func (o *DeviceEnrollmentDevice) GetProfileStatus() string`

GetProfileStatus returns the ProfileStatus field if non-nil, zero value otherwise.

### GetProfileStatusOk

`func (o *DeviceEnrollmentDevice) GetProfileStatusOk() (*string, bool)`

GetProfileStatusOk returns a tuple with the ProfileStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileStatus

`func (o *DeviceEnrollmentDevice) SetProfileStatus(v string)`

SetProfileStatus sets ProfileStatus field to given value.

### HasProfileStatus

`func (o *DeviceEnrollmentDevice) HasProfileStatus() bool`

HasProfileStatus returns a boolean if a field has been set.

### GetSyncState

`func (o *DeviceEnrollmentDevice) GetSyncState() AssignRemoveProfileResponseSyncState`

GetSyncState returns the SyncState field if non-nil, zero value otherwise.

### GetSyncStateOk

`func (o *DeviceEnrollmentDevice) GetSyncStateOk() (*AssignRemoveProfileResponseSyncState, bool)`

GetSyncStateOk returns a tuple with the SyncState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncState

`func (o *DeviceEnrollmentDevice) SetSyncState(v AssignRemoveProfileResponseSyncState)`

SetSyncState sets SyncState field to given value.

### HasSyncState

`func (o *DeviceEnrollmentDevice) HasSyncState() bool`

HasSyncState returns a boolean if a field has been set.

### GetProfileAssignTime

`func (o *DeviceEnrollmentDevice) GetProfileAssignTime() string`

GetProfileAssignTime returns the ProfileAssignTime field if non-nil, zero value otherwise.

### GetProfileAssignTimeOk

`func (o *DeviceEnrollmentDevice) GetProfileAssignTimeOk() (*string, bool)`

GetProfileAssignTimeOk returns a tuple with the ProfileAssignTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileAssignTime

`func (o *DeviceEnrollmentDevice) SetProfileAssignTime(v string)`

SetProfileAssignTime sets ProfileAssignTime field to given value.

### HasProfileAssignTime

`func (o *DeviceEnrollmentDevice) HasProfileAssignTime() bool`

HasProfileAssignTime returns a boolean if a field has been set.

### GetProfilePushTime

`func (o *DeviceEnrollmentDevice) GetProfilePushTime() string`

GetProfilePushTime returns the ProfilePushTime field if non-nil, zero value otherwise.

### GetProfilePushTimeOk

`func (o *DeviceEnrollmentDevice) GetProfilePushTimeOk() (*string, bool)`

GetProfilePushTimeOk returns a tuple with the ProfilePushTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilePushTime

`func (o *DeviceEnrollmentDevice) SetProfilePushTime(v string)`

SetProfilePushTime sets ProfilePushTime field to given value.

### HasProfilePushTime

`func (o *DeviceEnrollmentDevice) HasProfilePushTime() bool`

HasProfilePushTime returns a boolean if a field has been set.

### GetDeviceAssignedDate

`func (o *DeviceEnrollmentDevice) GetDeviceAssignedDate() string`

GetDeviceAssignedDate returns the DeviceAssignedDate field if non-nil, zero value otherwise.

### GetDeviceAssignedDateOk

`func (o *DeviceEnrollmentDevice) GetDeviceAssignedDateOk() (*string, bool)`

GetDeviceAssignedDateOk returns a tuple with the DeviceAssignedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceAssignedDate

`func (o *DeviceEnrollmentDevice) SetDeviceAssignedDate(v string)`

SetDeviceAssignedDate sets DeviceAssignedDate field to given value.

### HasDeviceAssignedDate

`func (o *DeviceEnrollmentDevice) HasDeviceAssignedDate() bool`

HasDeviceAssignedDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


