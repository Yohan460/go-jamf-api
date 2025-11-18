# ApplyRedemptionCodeCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommandType** | [**MdmCommandType**](MdmCommandType.md) |  | 
**Identifier** | Pointer to **string** | The bundle identifier of the app | [optional] 
**RedemptionCode** | Pointer to **string** | The redemption code that applies to the app pending installation | [optional] 

## Methods

### NewApplyRedemptionCodeCommand

`func NewApplyRedemptionCodeCommand(commandType MdmCommandType, ) *ApplyRedemptionCodeCommand`

NewApplyRedemptionCodeCommand instantiates a new ApplyRedemptionCodeCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplyRedemptionCodeCommandWithDefaults

`func NewApplyRedemptionCodeCommandWithDefaults() *ApplyRedemptionCodeCommand`

NewApplyRedemptionCodeCommandWithDefaults instantiates a new ApplyRedemptionCodeCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommandType

`func (o *ApplyRedemptionCodeCommand) GetCommandType() MdmCommandType`

GetCommandType returns the CommandType field if non-nil, zero value otherwise.

### GetCommandTypeOk

`func (o *ApplyRedemptionCodeCommand) GetCommandTypeOk() (*MdmCommandType, bool)`

GetCommandTypeOk returns a tuple with the CommandType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandType

`func (o *ApplyRedemptionCodeCommand) SetCommandType(v MdmCommandType)`

SetCommandType sets CommandType field to given value.


### GetIdentifier

`func (o *ApplyRedemptionCodeCommand) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *ApplyRedemptionCodeCommand) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *ApplyRedemptionCodeCommand) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *ApplyRedemptionCodeCommand) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetRedemptionCode

`func (o *ApplyRedemptionCodeCommand) GetRedemptionCode() string`

GetRedemptionCode returns the RedemptionCode field if non-nil, zero value otherwise.

### GetRedemptionCodeOk

`func (o *ApplyRedemptionCodeCommand) GetRedemptionCodeOk() (*string, bool)`

GetRedemptionCodeOk returns a tuple with the RedemptionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedemptionCode

`func (o *ApplyRedemptionCodeCommand) SetRedemptionCode(v string)`

SetRedemptionCode sets RedemptionCode field to given value.

### HasRedemptionCode

`func (o *ApplyRedemptionCodeCommand) HasRedemptionCode() bool`

HasRedemptionCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


