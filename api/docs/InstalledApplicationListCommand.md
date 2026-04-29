# InstalledApplicationListCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommandType** | [**MdmCommandType**](MdmCommandType.md) |  | 
**Identifiers** | Pointer to **[]string** | Array of application bundle identifiers to query. If not provided, all installed applications are returned. | [optional] 
**ManagedAppsOnly** | Pointer to **NullableBool** | If true, only managed applications are returned. If false or not provided, all applications are returned. | [optional] 
**Items** | Pointer to **[]string** | Array of keys to include in the response for each application. If not provided, default keys are returned. | [optional] 

## Methods

### NewInstalledApplicationListCommand

`func NewInstalledApplicationListCommand(commandType MdmCommandType, ) *InstalledApplicationListCommand`

NewInstalledApplicationListCommand instantiates a new InstalledApplicationListCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstalledApplicationListCommandWithDefaults

`func NewInstalledApplicationListCommandWithDefaults() *InstalledApplicationListCommand`

NewInstalledApplicationListCommandWithDefaults instantiates a new InstalledApplicationListCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommandType

`func (o *InstalledApplicationListCommand) GetCommandType() MdmCommandType`

GetCommandType returns the CommandType field if non-nil, zero value otherwise.

### GetCommandTypeOk

`func (o *InstalledApplicationListCommand) GetCommandTypeOk() (*MdmCommandType, bool)`

GetCommandTypeOk returns a tuple with the CommandType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandType

`func (o *InstalledApplicationListCommand) SetCommandType(v MdmCommandType)`

SetCommandType sets CommandType field to given value.


### GetIdentifiers

`func (o *InstalledApplicationListCommand) GetIdentifiers() []string`

GetIdentifiers returns the Identifiers field if non-nil, zero value otherwise.

### GetIdentifiersOk

`func (o *InstalledApplicationListCommand) GetIdentifiersOk() (*[]string, bool)`

GetIdentifiersOk returns a tuple with the Identifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiers

`func (o *InstalledApplicationListCommand) SetIdentifiers(v []string)`

SetIdentifiers sets Identifiers field to given value.

### HasIdentifiers

`func (o *InstalledApplicationListCommand) HasIdentifiers() bool`

HasIdentifiers returns a boolean if a field has been set.

### SetIdentifiersNil

`func (o *InstalledApplicationListCommand) SetIdentifiersNil(b bool)`

 SetIdentifiersNil sets the value for Identifiers to be an explicit nil

### UnsetIdentifiers
`func (o *InstalledApplicationListCommand) UnsetIdentifiers()`

UnsetIdentifiers ensures that no value is present for Identifiers, not even an explicit nil
### GetManagedAppsOnly

`func (o *InstalledApplicationListCommand) GetManagedAppsOnly() bool`

GetManagedAppsOnly returns the ManagedAppsOnly field if non-nil, zero value otherwise.

### GetManagedAppsOnlyOk

`func (o *InstalledApplicationListCommand) GetManagedAppsOnlyOk() (*bool, bool)`

GetManagedAppsOnlyOk returns a tuple with the ManagedAppsOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedAppsOnly

`func (o *InstalledApplicationListCommand) SetManagedAppsOnly(v bool)`

SetManagedAppsOnly sets ManagedAppsOnly field to given value.

### HasManagedAppsOnly

`func (o *InstalledApplicationListCommand) HasManagedAppsOnly() bool`

HasManagedAppsOnly returns a boolean if a field has been set.

### SetManagedAppsOnlyNil

`func (o *InstalledApplicationListCommand) SetManagedAppsOnlyNil(b bool)`

 SetManagedAppsOnlyNil sets the value for ManagedAppsOnly to be an explicit nil

### UnsetManagedAppsOnly
`func (o *InstalledApplicationListCommand) UnsetManagedAppsOnly()`

UnsetManagedAppsOnly ensures that no value is present for ManagedAppsOnly, not even an explicit nil
### GetItems

`func (o *InstalledApplicationListCommand) GetItems() []string`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InstalledApplicationListCommand) GetItemsOk() (*[]string, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InstalledApplicationListCommand) SetItems(v []string)`

SetItems sets Items field to given value.

### HasItems

`func (o *InstalledApplicationListCommand) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *InstalledApplicationListCommand) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *InstalledApplicationListCommand) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


