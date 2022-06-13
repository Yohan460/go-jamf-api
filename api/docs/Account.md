# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**RealName** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Preferences** | Pointer to [**AccountPreferences**](AccountPreferences.md) |  | [optional] 
**IsMultiSiteAdmin** | Pointer to **bool** |  | [optional] 
**AccessLevel** | Pointer to **string** |  | [optional] 
**PrivilegeSet** | Pointer to **string** |  | [optional] 
**PrivilegesBySite** | Pointer to **map[string][]string** |  | [optional] 
**GroupIds** | Pointer to **[]int32** |  | [optional] 
**CurrentSiteId** | Pointer to **int32** |  | [optional] 

## Methods

### NewAccount

`func NewAccount() *Account`

NewAccount instantiates a new Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountWithDefaults

`func NewAccountWithDefaults() *Account`

NewAccountWithDefaults instantiates a new Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Account) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Account) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Account) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Account) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUsername

`func (o *Account) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Account) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Account) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *Account) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetRealName

`func (o *Account) GetRealName() string`

GetRealName returns the RealName field if non-nil, zero value otherwise.

### GetRealNameOk

`func (o *Account) GetRealNameOk() (*string, bool)`

GetRealNameOk returns a tuple with the RealName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealName

`func (o *Account) SetRealName(v string)`

SetRealName sets RealName field to given value.

### HasRealName

`func (o *Account) HasRealName() bool`

HasRealName returns a boolean if a field has been set.

### GetEmail

`func (o *Account) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Account) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Account) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Account) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPreferences

`func (o *Account) GetPreferences() AccountPreferences`

GetPreferences returns the Preferences field if non-nil, zero value otherwise.

### GetPreferencesOk

`func (o *Account) GetPreferencesOk() (*AccountPreferences, bool)`

GetPreferencesOk returns a tuple with the Preferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferences

`func (o *Account) SetPreferences(v AccountPreferences)`

SetPreferences sets Preferences field to given value.

### HasPreferences

`func (o *Account) HasPreferences() bool`

HasPreferences returns a boolean if a field has been set.

### GetIsMultiSiteAdmin

`func (o *Account) GetIsMultiSiteAdmin() bool`

GetIsMultiSiteAdmin returns the IsMultiSiteAdmin field if non-nil, zero value otherwise.

### GetIsMultiSiteAdminOk

`func (o *Account) GetIsMultiSiteAdminOk() (*bool, bool)`

GetIsMultiSiteAdminOk returns a tuple with the IsMultiSiteAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultiSiteAdmin

`func (o *Account) SetIsMultiSiteAdmin(v bool)`

SetIsMultiSiteAdmin sets IsMultiSiteAdmin field to given value.

### HasIsMultiSiteAdmin

`func (o *Account) HasIsMultiSiteAdmin() bool`

HasIsMultiSiteAdmin returns a boolean if a field has been set.

### GetAccessLevel

`func (o *Account) GetAccessLevel() string`

GetAccessLevel returns the AccessLevel field if non-nil, zero value otherwise.

### GetAccessLevelOk

`func (o *Account) GetAccessLevelOk() (*string, bool)`

GetAccessLevelOk returns a tuple with the AccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevel

`func (o *Account) SetAccessLevel(v string)`

SetAccessLevel sets AccessLevel field to given value.

### HasAccessLevel

`func (o *Account) HasAccessLevel() bool`

HasAccessLevel returns a boolean if a field has been set.

### GetPrivilegeSet

`func (o *Account) GetPrivilegeSet() string`

GetPrivilegeSet returns the PrivilegeSet field if non-nil, zero value otherwise.

### GetPrivilegeSetOk

`func (o *Account) GetPrivilegeSetOk() (*string, bool)`

GetPrivilegeSetOk returns a tuple with the PrivilegeSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilegeSet

`func (o *Account) SetPrivilegeSet(v string)`

SetPrivilegeSet sets PrivilegeSet field to given value.

### HasPrivilegeSet

`func (o *Account) HasPrivilegeSet() bool`

HasPrivilegeSet returns a boolean if a field has been set.

### GetPrivilegesBySite

`func (o *Account) GetPrivilegesBySite() map[string][]string`

GetPrivilegesBySite returns the PrivilegesBySite field if non-nil, zero value otherwise.

### GetPrivilegesBySiteOk

`func (o *Account) GetPrivilegesBySiteOk() (*map[string][]string, bool)`

GetPrivilegesBySiteOk returns a tuple with the PrivilegesBySite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilegesBySite

`func (o *Account) SetPrivilegesBySite(v map[string][]string)`

SetPrivilegesBySite sets PrivilegesBySite field to given value.

### HasPrivilegesBySite

`func (o *Account) HasPrivilegesBySite() bool`

HasPrivilegesBySite returns a boolean if a field has been set.

### GetGroupIds

`func (o *Account) GetGroupIds() []int32`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *Account) GetGroupIdsOk() (*[]int32, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *Account) SetGroupIds(v []int32)`

SetGroupIds sets GroupIds field to given value.

### HasGroupIds

`func (o *Account) HasGroupIds() bool`

HasGroupIds returns a boolean if a field has been set.

### GetCurrentSiteId

`func (o *Account) GetCurrentSiteId() int32`

GetCurrentSiteId returns the CurrentSiteId field if non-nil, zero value otherwise.

### GetCurrentSiteIdOk

`func (o *Account) GetCurrentSiteIdOk() (*int32, bool)`

GetCurrentSiteIdOk returns a tuple with the CurrentSiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSiteId

`func (o *Account) SetCurrentSiteId(v int32)`

SetCurrentSiteId sets CurrentSiteId field to given value.

### HasCurrentSiteId

`func (o *Account) HasCurrentSiteId() bool`

HasCurrentSiteId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


