# UserAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**PlainPassword** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Realname** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**LdapServerId** | Pointer to **int64** |  | [optional] [default to -1]
**DistinguishedName** | Pointer to **string** |  | [optional] 
**SiteId** | Pointer to **int64** |  | [optional] [default to -1]
**AccessLevel** | Pointer to **string** | Access level for the account | [optional] [default to "FullAccess"]
**PrivilegeLevel** | Pointer to **string** | Privilege level for the account | [optional] [default to "ADMINISTRATOR"]
**LastPasswordChange** | Pointer to **time.Time** |  | [optional] [readonly] 
**ChangePasswordOnNextLogin** | Pointer to **bool** |  | [optional] 
**FailedLoginAttempts** | Pointer to **int64** |  | [optional] [readonly] 
**AccountStatus** | Pointer to **string** | Status of the account | [optional] [default to "Enabled"]

## Methods

### NewUserAccount

`func NewUserAccount() *UserAccount`

NewUserAccount instantiates a new UserAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAccountWithDefaults

`func NewUserAccountWithDefaults() *UserAccount`

NewUserAccountWithDefaults instantiates a new UserAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserAccount) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPlainPassword

`func (o *UserAccount) GetPlainPassword() string`

GetPlainPassword returns the PlainPassword field if non-nil, zero value otherwise.

### GetPlainPasswordOk

`func (o *UserAccount) GetPlainPasswordOk() (*string, bool)`

GetPlainPasswordOk returns a tuple with the PlainPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlainPassword

`func (o *UserAccount) SetPlainPassword(v string)`

SetPlainPassword sets PlainPassword field to given value.

### HasPlainPassword

`func (o *UserAccount) HasPlainPassword() bool`

HasPlainPassword returns a boolean if a field has been set.

### GetUsername

`func (o *UserAccount) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserAccount) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserAccount) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserAccount) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetRealname

`func (o *UserAccount) GetRealname() string`

GetRealname returns the Realname field if non-nil, zero value otherwise.

### GetRealnameOk

`func (o *UserAccount) GetRealnameOk() (*string, bool)`

GetRealnameOk returns a tuple with the Realname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealname

`func (o *UserAccount) SetRealname(v string)`

SetRealname sets Realname field to given value.

### HasRealname

`func (o *UserAccount) HasRealname() bool`

HasRealname returns a boolean if a field has been set.

### GetEmail

`func (o *UserAccount) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserAccount) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserAccount) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserAccount) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhone

`func (o *UserAccount) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *UserAccount) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *UserAccount) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *UserAccount) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetLdapServerId

`func (o *UserAccount) GetLdapServerId() int64`

GetLdapServerId returns the LdapServerId field if non-nil, zero value otherwise.

### GetLdapServerIdOk

`func (o *UserAccount) GetLdapServerIdOk() (*int64, bool)`

GetLdapServerIdOk returns a tuple with the LdapServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapServerId

`func (o *UserAccount) SetLdapServerId(v int64)`

SetLdapServerId sets LdapServerId field to given value.

### HasLdapServerId

`func (o *UserAccount) HasLdapServerId() bool`

HasLdapServerId returns a boolean if a field has been set.

### GetDistinguishedName

`func (o *UserAccount) GetDistinguishedName() string`

GetDistinguishedName returns the DistinguishedName field if non-nil, zero value otherwise.

### GetDistinguishedNameOk

`func (o *UserAccount) GetDistinguishedNameOk() (*string, bool)`

GetDistinguishedNameOk returns a tuple with the DistinguishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinguishedName

`func (o *UserAccount) SetDistinguishedName(v string)`

SetDistinguishedName sets DistinguishedName field to given value.

### HasDistinguishedName

`func (o *UserAccount) HasDistinguishedName() bool`

HasDistinguishedName returns a boolean if a field has been set.

### GetSiteId

`func (o *UserAccount) GetSiteId() int64`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *UserAccount) GetSiteIdOk() (*int64, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *UserAccount) SetSiteId(v int64)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *UserAccount) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetAccessLevel

`func (o *UserAccount) GetAccessLevel() string`

GetAccessLevel returns the AccessLevel field if non-nil, zero value otherwise.

### GetAccessLevelOk

`func (o *UserAccount) GetAccessLevelOk() (*string, bool)`

GetAccessLevelOk returns a tuple with the AccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevel

`func (o *UserAccount) SetAccessLevel(v string)`

SetAccessLevel sets AccessLevel field to given value.

### HasAccessLevel

`func (o *UserAccount) HasAccessLevel() bool`

HasAccessLevel returns a boolean if a field has been set.

### GetPrivilegeLevel

`func (o *UserAccount) GetPrivilegeLevel() string`

GetPrivilegeLevel returns the PrivilegeLevel field if non-nil, zero value otherwise.

### GetPrivilegeLevelOk

`func (o *UserAccount) GetPrivilegeLevelOk() (*string, bool)`

GetPrivilegeLevelOk returns a tuple with the PrivilegeLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilegeLevel

`func (o *UserAccount) SetPrivilegeLevel(v string)`

SetPrivilegeLevel sets PrivilegeLevel field to given value.

### HasPrivilegeLevel

`func (o *UserAccount) HasPrivilegeLevel() bool`

HasPrivilegeLevel returns a boolean if a field has been set.

### GetLastPasswordChange

`func (o *UserAccount) GetLastPasswordChange() time.Time`

GetLastPasswordChange returns the LastPasswordChange field if non-nil, zero value otherwise.

### GetLastPasswordChangeOk

`func (o *UserAccount) GetLastPasswordChangeOk() (*time.Time, bool)`

GetLastPasswordChangeOk returns a tuple with the LastPasswordChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPasswordChange

`func (o *UserAccount) SetLastPasswordChange(v time.Time)`

SetLastPasswordChange sets LastPasswordChange field to given value.

### HasLastPasswordChange

`func (o *UserAccount) HasLastPasswordChange() bool`

HasLastPasswordChange returns a boolean if a field has been set.

### GetChangePasswordOnNextLogin

`func (o *UserAccount) GetChangePasswordOnNextLogin() bool`

GetChangePasswordOnNextLogin returns the ChangePasswordOnNextLogin field if non-nil, zero value otherwise.

### GetChangePasswordOnNextLoginOk

`func (o *UserAccount) GetChangePasswordOnNextLoginOk() (*bool, bool)`

GetChangePasswordOnNextLoginOk returns a tuple with the ChangePasswordOnNextLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangePasswordOnNextLogin

`func (o *UserAccount) SetChangePasswordOnNextLogin(v bool)`

SetChangePasswordOnNextLogin sets ChangePasswordOnNextLogin field to given value.

### HasChangePasswordOnNextLogin

`func (o *UserAccount) HasChangePasswordOnNextLogin() bool`

HasChangePasswordOnNextLogin returns a boolean if a field has been set.

### GetFailedLoginAttempts

`func (o *UserAccount) GetFailedLoginAttempts() int64`

GetFailedLoginAttempts returns the FailedLoginAttempts field if non-nil, zero value otherwise.

### GetFailedLoginAttemptsOk

`func (o *UserAccount) GetFailedLoginAttemptsOk() (*int64, bool)`

GetFailedLoginAttemptsOk returns a tuple with the FailedLoginAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedLoginAttempts

`func (o *UserAccount) SetFailedLoginAttempts(v int64)`

SetFailedLoginAttempts sets FailedLoginAttempts field to given value.

### HasFailedLoginAttempts

`func (o *UserAccount) HasFailedLoginAttempts() bool`

HasFailedLoginAttempts returns a boolean if a field has been set.

### GetAccountStatus

`func (o *UserAccount) GetAccountStatus() string`

GetAccountStatus returns the AccountStatus field if non-nil, zero value otherwise.

### GetAccountStatusOk

`func (o *UserAccount) GetAccountStatusOk() (*string, bool)`

GetAccountStatusOk returns a tuple with the AccountStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountStatus

`func (o *UserAccount) SetAccountStatus(v string)`

SetAccountStatus sets AccountStatus field to given value.

### HasAccountStatus

`func (o *UserAccount) HasAccountStatus() bool`

HasAccountStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


