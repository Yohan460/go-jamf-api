# AccountGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessLevel** | Pointer to **string** |  | [optional] 
**PrivilegeSet** | Pointer to **string** |  | [optional] 
**SiteId** | Pointer to **int64** |  | [optional] 
**Privileges** | Pointer to **[]string** |  | [optional] 
**MemberUserIds** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewAccountGroup

`func NewAccountGroup() *AccountGroup`

NewAccountGroup instantiates a new AccountGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountGroupWithDefaults

`func NewAccountGroupWithDefaults() *AccountGroup`

NewAccountGroupWithDefaults instantiates a new AccountGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessLevel

`func (o *AccountGroup) GetAccessLevel() string`

GetAccessLevel returns the AccessLevel field if non-nil, zero value otherwise.

### GetAccessLevelOk

`func (o *AccountGroup) GetAccessLevelOk() (*string, bool)`

GetAccessLevelOk returns a tuple with the AccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevel

`func (o *AccountGroup) SetAccessLevel(v string)`

SetAccessLevel sets AccessLevel field to given value.

### HasAccessLevel

`func (o *AccountGroup) HasAccessLevel() bool`

HasAccessLevel returns a boolean if a field has been set.

### GetPrivilegeSet

`func (o *AccountGroup) GetPrivilegeSet() string`

GetPrivilegeSet returns the PrivilegeSet field if non-nil, zero value otherwise.

### GetPrivilegeSetOk

`func (o *AccountGroup) GetPrivilegeSetOk() (*string, bool)`

GetPrivilegeSetOk returns a tuple with the PrivilegeSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilegeSet

`func (o *AccountGroup) SetPrivilegeSet(v string)`

SetPrivilegeSet sets PrivilegeSet field to given value.

### HasPrivilegeSet

`func (o *AccountGroup) HasPrivilegeSet() bool`

HasPrivilegeSet returns a boolean if a field has been set.

### GetSiteId

`func (o *AccountGroup) GetSiteId() int64`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *AccountGroup) GetSiteIdOk() (*int64, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *AccountGroup) SetSiteId(v int64)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *AccountGroup) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetPrivileges

`func (o *AccountGroup) GetPrivileges() []string`

GetPrivileges returns the Privileges field if non-nil, zero value otherwise.

### GetPrivilegesOk

`func (o *AccountGroup) GetPrivilegesOk() (*[]string, bool)`

GetPrivilegesOk returns a tuple with the Privileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileges

`func (o *AccountGroup) SetPrivileges(v []string)`

SetPrivileges sets Privileges field to given value.

### HasPrivileges

`func (o *AccountGroup) HasPrivileges() bool`

HasPrivileges returns a boolean if a field has been set.

### GetMemberUserIds

`func (o *AccountGroup) GetMemberUserIds() []int64`

GetMemberUserIds returns the MemberUserIds field if non-nil, zero value otherwise.

### GetMemberUserIdsOk

`func (o *AccountGroup) GetMemberUserIdsOk() (*[]int64, bool)`

GetMemberUserIdsOk returns a tuple with the MemberUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberUserIds

`func (o *AccountGroup) SetMemberUserIds(v []int64)`

SetMemberUserIds sets MemberUserIds field to given value.

### HasMemberUserIds

`func (o *AccountGroup) HasMemberUserIds() bool`

HasMemberUserIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


