# DeclarativeManagementCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommandType** | [**MdmCommandType**](MdmCommandType.md) |  | 
**Data** | Pointer to **string** | Base64 encoded data to be sent with the command | [optional] 

## Methods

### NewDeclarativeManagementCommand

`func NewDeclarativeManagementCommand(commandType MdmCommandType, ) *DeclarativeManagementCommand`

NewDeclarativeManagementCommand instantiates a new DeclarativeManagementCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeclarativeManagementCommandWithDefaults

`func NewDeclarativeManagementCommandWithDefaults() *DeclarativeManagementCommand`

NewDeclarativeManagementCommandWithDefaults instantiates a new DeclarativeManagementCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommandType

`func (o *DeclarativeManagementCommand) GetCommandType() MdmCommandType`

GetCommandType returns the CommandType field if non-nil, zero value otherwise.

### GetCommandTypeOk

`func (o *DeclarativeManagementCommand) GetCommandTypeOk() (*MdmCommandType, bool)`

GetCommandTypeOk returns a tuple with the CommandType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandType

`func (o *DeclarativeManagementCommand) SetCommandType(v MdmCommandType)`

SetCommandType sets CommandType field to given value.


### GetData

`func (o *DeclarativeManagementCommand) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DeclarativeManagementCommand) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DeclarativeManagementCommand) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *DeclarativeManagementCommand) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


