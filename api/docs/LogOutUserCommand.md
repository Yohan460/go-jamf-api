# LogOutUserCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestType** | Pointer to **string** |  | [optional] 
**RequestRequiresNetworkTether** | Pointer to **bool** |  | [optional] 

## Methods

### NewLogOutUserCommand

`func NewLogOutUserCommand() *LogOutUserCommand`

NewLogOutUserCommand instantiates a new LogOutUserCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogOutUserCommandWithDefaults

`func NewLogOutUserCommandWithDefaults() *LogOutUserCommand`

NewLogOutUserCommandWithDefaults instantiates a new LogOutUserCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestType

`func (o *LogOutUserCommand) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *LogOutUserCommand) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *LogOutUserCommand) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *LogOutUserCommand) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetRequestRequiresNetworkTether

`func (o *LogOutUserCommand) GetRequestRequiresNetworkTether() bool`

GetRequestRequiresNetworkTether returns the RequestRequiresNetworkTether field if non-nil, zero value otherwise.

### GetRequestRequiresNetworkTetherOk

`func (o *LogOutUserCommand) GetRequestRequiresNetworkTetherOk() (*bool, bool)`

GetRequestRequiresNetworkTetherOk returns a tuple with the RequestRequiresNetworkTether field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestRequiresNetworkTether

`func (o *LogOutUserCommand) SetRequestRequiresNetworkTether(v bool)`

SetRequestRequiresNetworkTether sets RequestRequiresNetworkTether field to given value.

### HasRequestRequiresNetworkTether

`func (o *LogOutUserCommand) HasRequestRequiresNetworkTether() bool`

HasRequestRequiresNetworkTether returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


