# DeviceLockCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommandType** | [**MdmCommandType**](MdmCommandType.md) |  | 
**Message** | Pointer to **string** |  | [optional] 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**Pin** | Pointer to **string** |  | [optional] 

## Methods

### NewDeviceLockCommand

`func NewDeviceLockCommand(commandType MdmCommandType, ) *DeviceLockCommand`

NewDeviceLockCommand instantiates a new DeviceLockCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceLockCommandWithDefaults

`func NewDeviceLockCommandWithDefaults() *DeviceLockCommand`

NewDeviceLockCommandWithDefaults instantiates a new DeviceLockCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommandType

`func (o *DeviceLockCommand) GetCommandType() MdmCommandType`

GetCommandType returns the CommandType field if non-nil, zero value otherwise.

### GetCommandTypeOk

`func (o *DeviceLockCommand) GetCommandTypeOk() (*MdmCommandType, bool)`

GetCommandTypeOk returns a tuple with the CommandType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandType

`func (o *DeviceLockCommand) SetCommandType(v MdmCommandType)`

SetCommandType sets CommandType field to given value.


### GetMessage

`func (o *DeviceLockCommand) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DeviceLockCommand) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DeviceLockCommand) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DeviceLockCommand) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *DeviceLockCommand) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *DeviceLockCommand) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *DeviceLockCommand) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *DeviceLockCommand) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetPin

`func (o *DeviceLockCommand) GetPin() string`

GetPin returns the Pin field if non-nil, zero value otherwise.

### GetPinOk

`func (o *DeviceLockCommand) GetPinOk() (*string, bool)`

GetPinOk returns a tuple with the Pin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPin

`func (o *DeviceLockCommand) SetPin(v string)`

SetPin sets Pin field to given value.

### HasPin

`func (o *DeviceLockCommand) HasPin() bool`

HasPin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


