# CurrentAccount

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
**Privileges** | Pointer to **[]string** |  | [optional] 
**GroupIds** | Pointer to **[]int32** |  | [optional] 
**CurrentSiteId** | Pointer to **int32** |  | [optional] 

## Methods

### NewCurrentAccount

`func NewCurrentAccount() *CurrentAccount`

NewCurrentAccount instantiates a new CurrentAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrentAccountWithDefaults

`func NewCurrentAccountWithDefaults() *CurrentAccount`

NewCurrentAccountWithDefaults instantiates a new CurrentAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CurrentAccount) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CurrentAccount) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CurrentAccount) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CurrentAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUsername

`func (o *CurrentAccount) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CurrentAccount) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CurrentAccount) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CurrentAccount) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetRealName

`func (o *CurrentAccount) GetRealName() string`

GetRealName returns the RealName field if non-nil, zero value otherwise.

### GetRealNameOk

`func (o *CurrentAccount) GetRealNameOk() (*string, bool)`

GetRealNameOk returns a tuple with the RealName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealName

`func (o *CurrentAccount) SetRealName(v string)`

SetRealName sets RealName field to given value.

### HasRealName

`func (o *CurrentAccount) HasRealName() bool`

HasRealName returns a boolean if a field has been set.

### GetEmail

`func (o *CurrentAccount) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CurrentAccount) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CurrentAccount) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CurrentAccount) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPreferences

`func (o *CurrentAccount) GetPreferences() AccountPreferences`

GetPreferences returns the Preferences field if non-nil, zero value otherwise.

### GetPreferencesOk

`func (o *CurrentAccount) GetPreferencesOk() (*AccountPreferences, bool)`

GetPreferencesOk returns a tuple with the Preferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferences

`func (o *CurrentAccount) SetPreferences(v AccountPreferences)`

SetPreferences sets Preferences field to given value.

### HasPreferences

`func (o *CurrentAccount) HasPreferences() bool`

HasPreferences returns a boolean if a field has been set.

### GetIsMultiSiteAdmin

`func (o *CurrentAccount) GetIsMultiSiteAdmin() bool`

GetIsMultiSiteAdmin returns the IsMultiSiteAdmin field if non-nil, zero value otherwise.

### GetIsMultiSiteAdminOk

`func (o *CurrentAccount) GetIsMultiSiteAdminOk() (*bool, bool)`

GetIsMultiSiteAdminOk returns a tuple with the IsMultiSiteAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultiSiteAdmin

`func (o *CurrentAccount) SetIsMultiSiteAdmin(v bool)`

SetIsMultiSiteAdmin sets IsMultiSiteAdmin field to given value.

### HasIsMultiSiteAdmin

`func (o *CurrentAccount) HasIsMultiSiteAdmin() bool`

HasIsMultiSiteAdmin returns a boolean if a field has been set.

### GetAccessLevel

`func (o *CurrentAccount) GetAccessLevel() string`

GetAccessLevel returns the AccessLevel field if non-nil, zero value otherwise.

### GetAccessLevelOk

`func (o *CurrentAccount) GetAccessLevelOk() (*string, bool)`

GetAccessLevelOk returns a tuple with the AccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevel

`func (o *CurrentAccount) SetAccessLevel(v string)`

SetAccessLevel sets AccessLevel field to given value.

### HasAccessLevel

`func (o *CurrentAccount) HasAccessLevel() bool`

HasAccessLevel returns a boolean if a field has been set.

### GetPrivilegeSet

`func (o *CurrentAccount) GetPrivilegeSet() string`

GetPrivilegeSet returns the PrivilegeSet field if non-nil, zero value otherwise.

### GetPrivilegeSetOk

`func (o *CurrentAccount) GetPrivilegeSetOk() (*string, bool)`

GetPrivilegeSetOk returns a tuple with the PrivilegeSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilegeSet

`func (o *CurrentAccount) SetPrivilegeSet(v string)`

SetPrivilegeSet sets PrivilegeSet field to given value.

### HasPrivilegeSet

`func (o *CurrentAccount) HasPrivilegeSet() bool`

HasPrivilegeSet returns a boolean if a field has been set.

### GetPrivileges

`func (o *CurrentAccount) GetPrivileges() []string`

GetPrivileges returns the Privileges field if non-nil, zero value otherwise.

### GetPrivilegesOk

`func (o *CurrentAccount) GetPrivilegesOk() (*[]string, bool)`

GetPrivilegesOk returns a tuple with the Privileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileges

`func (o *CurrentAccount) SetPrivileges(v []string)`

SetPrivileges sets Privileges field to given value.

### HasPrivileges

`func (o *CurrentAccount) HasPrivileges() bool`

HasPrivileges returns a boolean if a field has been set.

### GetGroupIds

`func (o *CurrentAccount) GetGroupIds() []int32`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *CurrentAccount) GetGroupIdsOk() (*[]int32, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *CurrentAccount) SetGroupIds(v []int32)`

SetGroupIds sets GroupIds field to given value.

### HasGroupIds

`func (o *CurrentAccount) HasGroupIds() bool`

HasGroupIds returns a boolean if a field has been set.

### GetCurrentSiteId

`func (o *CurrentAccount) GetCurrentSiteId() int32`

GetCurrentSiteId returns the CurrentSiteId field if non-nil, zero value otherwise.

### GetCurrentSiteIdOk

`func (o *CurrentAccount) GetCurrentSiteIdOk() (*int32, bool)`

GetCurrentSiteIdOk returns a tuple with the CurrentSiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSiteId

`func (o *CurrentAccount) SetCurrentSiteId(v int32)`

SetCurrentSiteId sets CurrentSiteId field to given value.

### HasCurrentSiteId

`func (o *CurrentAccount) HasCurrentSiteId() bool`

HasCurrentSiteId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


