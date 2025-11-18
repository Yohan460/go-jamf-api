# EraseDeviceComputerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | **string** | Id of the computer for which eraseDevice command was queued | 
**CommandUuid** | **string** | Uuid of the queued eraseDevice command | 

## Methods

### NewEraseDeviceComputerResponse

`func NewEraseDeviceComputerResponse(deviceId string, commandUuid string, ) *EraseDeviceComputerResponse`

NewEraseDeviceComputerResponse instantiates a new EraseDeviceComputerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEraseDeviceComputerResponseWithDefaults

`func NewEraseDeviceComputerResponseWithDefaults() *EraseDeviceComputerResponse`

NewEraseDeviceComputerResponseWithDefaults instantiates a new EraseDeviceComputerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *EraseDeviceComputerResponse) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *EraseDeviceComputerResponse) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *EraseDeviceComputerResponse) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.


### GetCommandUuid

`func (o *EraseDeviceComputerResponse) GetCommandUuid() string`

GetCommandUuid returns the CommandUuid field if non-nil, zero value otherwise.

### GetCommandUuidOk

`func (o *EraseDeviceComputerResponse) GetCommandUuidOk() (*string, bool)`

GetCommandUuidOk returns a tuple with the CommandUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandUuid

`func (o *EraseDeviceComputerResponse) SetCommandUuid(v string)`

SetCommandUuid sets CommandUuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


