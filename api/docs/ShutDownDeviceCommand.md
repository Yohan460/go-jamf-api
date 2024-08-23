# ShutDownDeviceCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommandType** | [**MdmCommandType**](MdmCommandType.md) |  | 

## Methods

### NewShutDownDeviceCommand

`func NewShutDownDeviceCommand(commandType MdmCommandType, ) *ShutDownDeviceCommand`

NewShutDownDeviceCommand instantiates a new ShutDownDeviceCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShutDownDeviceCommandWithDefaults

`func NewShutDownDeviceCommandWithDefaults() *ShutDownDeviceCommand`

NewShutDownDeviceCommandWithDefaults instantiates a new ShutDownDeviceCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommandType

`func (o *ShutDownDeviceCommand) GetCommandType() MdmCommandType`

GetCommandType returns the CommandType field if non-nil, zero value otherwise.

### GetCommandTypeOk

`func (o *ShutDownDeviceCommand) GetCommandTypeOk() (*MdmCommandType, bool)`

GetCommandTypeOk returns a tuple with the CommandType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandType

`func (o *ShutDownDeviceCommand) SetCommandType(v MdmCommandType)`

SetCommandType sets CommandType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


