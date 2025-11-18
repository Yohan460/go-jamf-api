# EraseDeviceComputerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pin** | Pointer to **NullableString** | The six-character PIN for Find My. | [optional] 

## Methods

### NewEraseDeviceComputerRequest

`func NewEraseDeviceComputerRequest() *EraseDeviceComputerRequest`

NewEraseDeviceComputerRequest instantiates a new EraseDeviceComputerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEraseDeviceComputerRequestWithDefaults

`func NewEraseDeviceComputerRequestWithDefaults() *EraseDeviceComputerRequest`

NewEraseDeviceComputerRequestWithDefaults instantiates a new EraseDeviceComputerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPin

`func (o *EraseDeviceComputerRequest) GetPin() string`

GetPin returns the Pin field if non-nil, zero value otherwise.

### GetPinOk

`func (o *EraseDeviceComputerRequest) GetPinOk() (*string, bool)`

GetPinOk returns a tuple with the Pin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPin

`func (o *EraseDeviceComputerRequest) SetPin(v string)`

SetPin sets Pin field to given value.

### HasPin

`func (o *EraseDeviceComputerRequest) HasPin() bool`

HasPin returns a boolean if a field has been set.

### SetPinNil

`func (o *EraseDeviceComputerRequest) SetPinNil(b bool)`

 SetPinNil sets the value for Pin to be an explicit nil

### UnsetPin
`func (o *EraseDeviceComputerRequest) UnsetPin()`

UnsetPin ensures that no value is present for Pin, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


