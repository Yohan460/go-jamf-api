# AccountGroupV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**AccessLevel** | Pointer to **string** | Access level for the account group | [optional] [default to "FullAccess"]
**PrivilegeLevel** | Pointer to **string** | Privilege level for the account group | [optional] [default to "ADMINISTRATOR"]
**SiteId** | Pointer to **string** |  | [optional] 
**LdapServerId** | Pointer to **string** |  | [optional] 
**DirectoryGroupId** | Pointer to **string** |  | [optional] 
**Members** | Pointer to [**[]AccountGroupV1MembersInner**](AccountGroupV1MembersInner.md) | Members of this account group | [optional] 
**Privileges** | Pointer to **[]string** | List of privilege strings assigned to this group | [optional] 

## Methods

### NewAccountGroupV1

`func NewAccountGroupV1() *AccountGroupV1`

NewAccountGroupV1 instantiates a new AccountGroupV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountGroupV1WithDefaults

`func NewAccountGroupV1WithDefaults() *AccountGroupV1`

NewAccountGroupV1WithDefaults instantiates a new AccountGroupV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountGroupV1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountGroupV1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountGroupV1) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccountGroupV1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AccountGroupV1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountGroupV1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountGroupV1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccountGroupV1) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAccessLevel

`func (o *AccountGroupV1) GetAccessLevel() string`

GetAccessLevel returns the AccessLevel field if non-nil, zero value otherwise.

### GetAccessLevelOk

`func (o *AccountGroupV1) GetAccessLevelOk() (*string, bool)`

GetAccessLevelOk returns a tuple with the AccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevel

`func (o *AccountGroupV1) SetAccessLevel(v string)`

SetAccessLevel sets AccessLevel field to given value.

### HasAccessLevel

`func (o *AccountGroupV1) HasAccessLevel() bool`

HasAccessLevel returns a boolean if a field has been set.

### GetPrivilegeLevel

`func (o *AccountGroupV1) GetPrivilegeLevel() string`

GetPrivilegeLevel returns the PrivilegeLevel field if non-nil, zero value otherwise.

### GetPrivilegeLevelOk

`func (o *AccountGroupV1) GetPrivilegeLevelOk() (*string, bool)`

GetPrivilegeLevelOk returns a tuple with the PrivilegeLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilegeLevel

`func (o *AccountGroupV1) SetPrivilegeLevel(v string)`

SetPrivilegeLevel sets PrivilegeLevel field to given value.

### HasPrivilegeLevel

`func (o *AccountGroupV1) HasPrivilegeLevel() bool`

HasPrivilegeLevel returns a boolean if a field has been set.

### GetSiteId

`func (o *AccountGroupV1) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *AccountGroupV1) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *AccountGroupV1) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *AccountGroupV1) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetLdapServerId

`func (o *AccountGroupV1) GetLdapServerId() string`

GetLdapServerId returns the LdapServerId field if non-nil, zero value otherwise.

### GetLdapServerIdOk

`func (o *AccountGroupV1) GetLdapServerIdOk() (*string, bool)`

GetLdapServerIdOk returns a tuple with the LdapServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapServerId

`func (o *AccountGroupV1) SetLdapServerId(v string)`

SetLdapServerId sets LdapServerId field to given value.

### HasLdapServerId

`func (o *AccountGroupV1) HasLdapServerId() bool`

HasLdapServerId returns a boolean if a field has been set.

### GetDirectoryGroupId

`func (o *AccountGroupV1) GetDirectoryGroupId() string`

GetDirectoryGroupId returns the DirectoryGroupId field if non-nil, zero value otherwise.

### GetDirectoryGroupIdOk

`func (o *AccountGroupV1) GetDirectoryGroupIdOk() (*string, bool)`

GetDirectoryGroupIdOk returns a tuple with the DirectoryGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryGroupId

`func (o *AccountGroupV1) SetDirectoryGroupId(v string)`

SetDirectoryGroupId sets DirectoryGroupId field to given value.

### HasDirectoryGroupId

`func (o *AccountGroupV1) HasDirectoryGroupId() bool`

HasDirectoryGroupId returns a boolean if a field has been set.

### GetMembers

`func (o *AccountGroupV1) GetMembers() []AccountGroupV1MembersInner`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *AccountGroupV1) GetMembersOk() (*[]AccountGroupV1MembersInner, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *AccountGroupV1) SetMembers(v []AccountGroupV1MembersInner)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *AccountGroupV1) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetPrivileges

`func (o *AccountGroupV1) GetPrivileges() []string`

GetPrivileges returns the Privileges field if non-nil, zero value otherwise.

### GetPrivilegesOk

`func (o *AccountGroupV1) GetPrivilegesOk() (*[]string, bool)`

GetPrivilegesOk returns a tuple with the Privileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileges

`func (o *AccountGroupV1) SetPrivileges(v []string)`

SetPrivileges sets Privileges field to given value.

### HasPrivileges

`func (o *AccountGroupV1) HasPrivileges() bool`

HasPrivileges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


