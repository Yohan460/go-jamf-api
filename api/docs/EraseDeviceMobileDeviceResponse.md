# EraseDeviceMobileDeviceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | **string** | Id of the mobile device for which eraseDevice command was queued | 
**CommandUuid** | **string** | Uuid of the queued eraseDevice command | 

## Methods

### NewEraseDeviceMobileDeviceResponse

`func NewEraseDeviceMobileDeviceResponse(deviceId string, commandUuid string, ) *EraseDeviceMobileDeviceResponse`

NewEraseDeviceMobileDeviceResponse instantiates a new EraseDeviceMobileDeviceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEraseDeviceMobileDeviceResponseWithDefaults

`func NewEraseDeviceMobileDeviceResponseWithDefaults() *EraseDeviceMobileDeviceResponse`

NewEraseDeviceMobileDeviceResponseWithDefaults instantiates a new EraseDeviceMobileDeviceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *EraseDeviceMobileDeviceResponse) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *EraseDeviceMobileDeviceResponse) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *EraseDeviceMobileDeviceResponse) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.


### GetCommandUuid

`func (o *EraseDeviceMobileDeviceResponse) GetCommandUuid() string`

GetCommandUuid returns the CommandUuid field if non-nil, zero value otherwise.

### GetCommandUuidOk

`func (o *EraseDeviceMobileDeviceResponse) GetCommandUuidOk() (*string, bool)`

GetCommandUuidOk returns a tuple with the CommandUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandUuid

`func (o *EraseDeviceMobileDeviceResponse) SetCommandUuid(v string)`

SetCommandUuid sets CommandUuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


