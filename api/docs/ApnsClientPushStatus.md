# ApnsClientPushStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceType** | Pointer to **string** | The type of MDM client device | [optional] 
**ClientId** | Pointer to **string** | Id of the Computer or Device record in Jamf Pro | [optional] 
**DisabledAt** | Pointer to **time.Time** | Timestamp when push notifications were disabled for this client (ISO-8601 format) | [optional] 
**ManagementId** | Pointer to **string** | Unique identifier for the device management record | [optional] 

## Methods

### NewApnsClientPushStatus

`func NewApnsClientPushStatus() *ApnsClientPushStatus`

NewApnsClientPushStatus instantiates a new ApnsClientPushStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApnsClientPushStatusWithDefaults

`func NewApnsClientPushStatusWithDefaults() *ApnsClientPushStatus`

NewApnsClientPushStatusWithDefaults instantiates a new ApnsClientPushStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceType

`func (o *ApnsClientPushStatus) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *ApnsClientPushStatus) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *ApnsClientPushStatus) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *ApnsClientPushStatus) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetClientId

`func (o *ApnsClientPushStatus) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ApnsClientPushStatus) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ApnsClientPushStatus) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ApnsClientPushStatus) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetDisabledAt

`func (o *ApnsClientPushStatus) GetDisabledAt() time.Time`

GetDisabledAt returns the DisabledAt field if non-nil, zero value otherwise.

### GetDisabledAtOk

`func (o *ApnsClientPushStatus) GetDisabledAtOk() (*time.Time, bool)`

GetDisabledAtOk returns a tuple with the DisabledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledAt

`func (o *ApnsClientPushStatus) SetDisabledAt(v time.Time)`

SetDisabledAt sets DisabledAt field to given value.

### HasDisabledAt

`func (o *ApnsClientPushStatus) HasDisabledAt() bool`

HasDisabledAt returns a boolean if a field has been set.

### GetManagementId

`func (o *ApnsClientPushStatus) GetManagementId() string`

GetManagementId returns the ManagementId field if non-nil, zero value otherwise.

### GetManagementIdOk

`func (o *ApnsClientPushStatus) GetManagementIdOk() (*string, bool)`

GetManagementIdOk returns a tuple with the ManagementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementId

`func (o *ApnsClientPushStatus) SetManagementId(v string)`

SetManagementId sets ManagementId field to given value.

### HasManagementId

`func (o *ApnsClientPushStatus) HasManagementId() bool`

HasManagementId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


