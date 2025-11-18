# RemoveComputerMdmProfileResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **string** | Id of the computer whose MDM profile was removed | [optional] 
**CommandUuid** | Pointer to **string** | Uuid of the command queued that removes the MDM profile | [optional] 

## Methods

### NewRemoveComputerMdmProfileResponse

`func NewRemoveComputerMdmProfileResponse() *RemoveComputerMdmProfileResponse`

NewRemoveComputerMdmProfileResponse instantiates a new RemoveComputerMdmProfileResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoveComputerMdmProfileResponseWithDefaults

`func NewRemoveComputerMdmProfileResponseWithDefaults() *RemoveComputerMdmProfileResponse`

NewRemoveComputerMdmProfileResponseWithDefaults instantiates a new RemoveComputerMdmProfileResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *RemoveComputerMdmProfileResponse) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *RemoveComputerMdmProfileResponse) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *RemoveComputerMdmProfileResponse) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *RemoveComputerMdmProfileResponse) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetCommandUuid

`func (o *RemoveComputerMdmProfileResponse) GetCommandUuid() string`

GetCommandUuid returns the CommandUuid field if non-nil, zero value otherwise.

### GetCommandUuidOk

`func (o *RemoveComputerMdmProfileResponse) GetCommandUuidOk() (*string, bool)`

GetCommandUuidOk returns a tuple with the CommandUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandUuid

`func (o *RemoveComputerMdmProfileResponse) SetCommandUuid(v string)`

SetCommandUuid sets CommandUuid field to given value.

### HasCommandUuid

`func (o *RemoveComputerMdmProfileResponse) HasCommandUuid() bool`

HasCommandUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


