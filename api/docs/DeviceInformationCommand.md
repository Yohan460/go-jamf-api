# DeviceInformationCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommandType** | [**MdmCommandType**](MdmCommandType.md) |  | 
**Queries** | Pointer to **[]string** |  | [optional] 

## Methods

### NewDeviceInformationCommand

`func NewDeviceInformationCommand(commandType MdmCommandType, ) *DeviceInformationCommand`

NewDeviceInformationCommand instantiates a new DeviceInformationCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceInformationCommandWithDefaults

`func NewDeviceInformationCommandWithDefaults() *DeviceInformationCommand`

NewDeviceInformationCommandWithDefaults instantiates a new DeviceInformationCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommandType

`func (o *DeviceInformationCommand) GetCommandType() MdmCommandType`

GetCommandType returns the CommandType field if non-nil, zero value otherwise.

### GetCommandTypeOk

`func (o *DeviceInformationCommand) GetCommandTypeOk() (*MdmCommandType, bool)`

GetCommandTypeOk returns a tuple with the CommandType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandType

`func (o *DeviceInformationCommand) SetCommandType(v MdmCommandType)`

SetCommandType sets CommandType field to given value.


### GetQueries

`func (o *DeviceInformationCommand) GetQueries() []string`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *DeviceInformationCommand) GetQueriesOk() (*[]string, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *DeviceInformationCommand) SetQueries(v []string)`

SetQueries sets Queries field to given value.

### HasQueries

`func (o *DeviceInformationCommand) HasQueries() bool`

HasQueries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


