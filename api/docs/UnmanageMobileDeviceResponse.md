# UnmanageMobileDeviceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | **string** | Id of the mobile device whose MDM profile was removed | 
**CommandUuid** | **string** | Uuid of the command queued that removes the MDM profile | 

## Methods

### NewUnmanageMobileDeviceResponse

`func NewUnmanageMobileDeviceResponse(deviceId string, commandUuid string, ) *UnmanageMobileDeviceResponse`

NewUnmanageMobileDeviceResponse instantiates a new UnmanageMobileDeviceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnmanageMobileDeviceResponseWithDefaults

`func NewUnmanageMobileDeviceResponseWithDefaults() *UnmanageMobileDeviceResponse`

NewUnmanageMobileDeviceResponseWithDefaults instantiates a new UnmanageMobileDeviceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *UnmanageMobileDeviceResponse) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *UnmanageMobileDeviceResponse) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *UnmanageMobileDeviceResponse) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.


### GetCommandUuid

`func (o *UnmanageMobileDeviceResponse) GetCommandUuid() string`

GetCommandUuid returns the CommandUuid field if non-nil, zero value otherwise.

### GetCommandUuidOk

`func (o *UnmanageMobileDeviceResponse) GetCommandUuidOk() (*string, bool)`

GetCommandUuidOk returns a tuple with the CommandUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandUuid

`func (o *UnmanageMobileDeviceResponse) SetCommandUuid(v string)`

SetCommandUuid sets CommandUuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


