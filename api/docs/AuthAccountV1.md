# AuthAccountV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**RealName** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Preferences** | Pointer to [**AccountPreferencesV1**](AccountPreferencesV1.md) |  | [optional] 
**MultiSiteAdmin** | Pointer to **bool** |  | [optional] 
**AccessLevel** | Pointer to **string** |  | [optional] 
**PrivilegeSet** | Pointer to **string** |  | [optional] 
**PrivilegesBySite** | Pointer to **map[string][]string** |  | [optional] 
**GroupIds** | Pointer to **[]string** |  | [optional] 
**CurrentSiteId** | Pointer to **string** |  | [optional] 

## Methods

### NewAuthAccountV1

`func NewAuthAccountV1() *AuthAccountV1`

NewAuthAccountV1 instantiates a new AuthAccountV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthAccountV1WithDefaults

`func NewAuthAccountV1WithDefaults() *AuthAccountV1`

NewAuthAccountV1WithDefaults instantiates a new AuthAccountV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuthAccountV1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthAccountV1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthAccountV1) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthAccountV1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUsername

`func (o *AuthAccountV1) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AuthAccountV1) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AuthAccountV1) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AuthAccountV1) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetRealName

`func (o *AuthAccountV1) GetRealName() string`

GetRealName returns the RealName field if non-nil, zero value otherwise.

### GetRealNameOk

`func (o *AuthAccountV1) GetRealNameOk() (*string, bool)`

GetRealNameOk returns a tuple with the RealName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealName

`func (o *AuthAccountV1) SetRealName(v string)`

SetRealName sets RealName field to given value.

### HasRealName

`func (o *AuthAccountV1) HasRealName() bool`

HasRealName returns a boolean if a field has been set.

### GetEmail

`func (o *AuthAccountV1) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AuthAccountV1) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AuthAccountV1) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AuthAccountV1) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPreferences

`func (o *AuthAccountV1) GetPreferences() AccountPreferencesV1`

GetPreferences returns the Preferences field if non-nil, zero value otherwise.

### GetPreferencesOk

`func (o *AuthAccountV1) GetPreferencesOk() (*AccountPreferencesV1, bool)`

GetPreferencesOk returns a tuple with the Preferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferences

`func (o *AuthAccountV1) SetPreferences(v AccountPreferencesV1)`

SetPreferences sets Preferences field to given value.

### HasPreferences

`func (o *AuthAccountV1) HasPreferences() bool`

HasPreferences returns a boolean if a field has been set.

### GetMultiSiteAdmin

`func (o *AuthAccountV1) GetMultiSiteAdmin() bool`

GetMultiSiteAdmin returns the MultiSiteAdmin field if non-nil, zero value otherwise.

### GetMultiSiteAdminOk

`func (o *AuthAccountV1) GetMultiSiteAdminOk() (*bool, bool)`

GetMultiSiteAdminOk returns a tuple with the MultiSiteAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiSiteAdmin

`func (o *AuthAccountV1) SetMultiSiteAdmin(v bool)`

SetMultiSiteAdmin sets MultiSiteAdmin field to given value.

### HasMultiSiteAdmin

`func (o *AuthAccountV1) HasMultiSiteAdmin() bool`

HasMultiSiteAdmin returns a boolean if a field has been set.

### GetAccessLevel

`func (o *AuthAccountV1) GetAccessLevel() string`

GetAccessLevel returns the AccessLevel field if non-nil, zero value otherwise.

### GetAccessLevelOk

`func (o *AuthAccountV1) GetAccessLevelOk() (*string, bool)`

GetAccessLevelOk returns a tuple with the AccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevel

`func (o *AuthAccountV1) SetAccessLevel(v string)`

SetAccessLevel sets AccessLevel field to given value.

### HasAccessLevel

`func (o *AuthAccountV1) HasAccessLevel() bool`

HasAccessLevel returns a boolean if a field has been set.

### GetPrivilegeSet

`func (o *AuthAccountV1) GetPrivilegeSet() string`

GetPrivilegeSet returns the PrivilegeSet field if non-nil, zero value otherwise.

### GetPrivilegeSetOk

`func (o *AuthAccountV1) GetPrivilegeSetOk() (*string, bool)`

GetPrivilegeSetOk returns a tuple with the PrivilegeSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilegeSet

`func (o *AuthAccountV1) SetPrivilegeSet(v string)`

SetPrivilegeSet sets PrivilegeSet field to given value.

### HasPrivilegeSet

`func (o *AuthAccountV1) HasPrivilegeSet() bool`

HasPrivilegeSet returns a boolean if a field has been set.

### GetPrivilegesBySite

`func (o *AuthAccountV1) GetPrivilegesBySite() map[string][]string`

GetPrivilegesBySite returns the PrivilegesBySite field if non-nil, zero value otherwise.

### GetPrivilegesBySiteOk

`func (o *AuthAccountV1) GetPrivilegesBySiteOk() (*map[string][]string, bool)`

GetPrivilegesBySiteOk returns a tuple with the PrivilegesBySite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilegesBySite

`func (o *AuthAccountV1) SetPrivilegesBySite(v map[string][]string)`

SetPrivilegesBySite sets PrivilegesBySite field to given value.

### HasPrivilegesBySite

`func (o *AuthAccountV1) HasPrivilegesBySite() bool`

HasPrivilegesBySite returns a boolean if a field has been set.

### GetGroupIds

`func (o *AuthAccountV1) GetGroupIds() []string`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *AuthAccountV1) GetGroupIdsOk() (*[]string, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *AuthAccountV1) SetGroupIds(v []string)`

SetGroupIds sets GroupIds field to given value.

### HasGroupIds

`func (o *AuthAccountV1) HasGroupIds() bool`

HasGroupIds returns a boolean if a field has been set.

### GetCurrentSiteId

`func (o *AuthAccountV1) GetCurrentSiteId() string`

GetCurrentSiteId returns the CurrentSiteId field if non-nil, zero value otherwise.

### GetCurrentSiteIdOk

`func (o *AuthAccountV1) GetCurrentSiteIdOk() (*string, bool)`

GetCurrentSiteIdOk returns a tuple with the CurrentSiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSiteId

`func (o *AuthAccountV1) SetCurrentSiteId(v string)`

SetCurrentSiteId sets CurrentSiteId field to given value.

### HasCurrentSiteId

`func (o *AuthAccountV1) HasCurrentSiteId() bool`

HasCurrentSiteId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


