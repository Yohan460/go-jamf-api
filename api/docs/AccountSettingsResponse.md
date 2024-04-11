# AccountSettingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | id of Account Settings | [optional] 
**PayloadConfigured** | Pointer to **bool** |  | [optional] [default to false]
**LocalAdminAccountEnabled** | Pointer to **bool** |  | [optional] [default to false]
**AdminUsername** | Pointer to **string** |  | [optional] [default to ""]
**HiddenAdminAccount** | Pointer to **bool** |  | [optional] [default to false]
**LocalUserManaged** | Pointer to **bool** |  | [optional] [default to false]
**UserAccountType** | Pointer to **string** |  | [optional] [default to "STANDARD"]
**VersionLock** | Pointer to **int64** |  | [optional] [default to 0]
**PrefillPrimaryAccountInfoFeatureEnabled** | Pointer to **bool** |  | [optional] [default to false]
**PrefillType** | Pointer to **string** | Values accepted are only CUSTOM and DEVICE_OWNER | [optional] [default to "CUSTOM"]
**PrefillAccountFullName** | Pointer to **string** |  | [optional] [default to ""]
**PrefillAccountUserName** | Pointer to **string** |  | [optional] [default to ""]
**PreventPrefillInfoFromModification** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewAccountSettingsResponse

`func NewAccountSettingsResponse() *AccountSettingsResponse`

NewAccountSettingsResponse instantiates a new AccountSettingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountSettingsResponseWithDefaults

`func NewAccountSettingsResponseWithDefaults() *AccountSettingsResponse`

NewAccountSettingsResponseWithDefaults instantiates a new AccountSettingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountSettingsResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountSettingsResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountSettingsResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccountSettingsResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPayloadConfigured

`func (o *AccountSettingsResponse) GetPayloadConfigured() bool`

GetPayloadConfigured returns the PayloadConfigured field if non-nil, zero value otherwise.

### GetPayloadConfiguredOk

`func (o *AccountSettingsResponse) GetPayloadConfiguredOk() (*bool, bool)`

GetPayloadConfiguredOk returns a tuple with the PayloadConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadConfigured

`func (o *AccountSettingsResponse) SetPayloadConfigured(v bool)`

SetPayloadConfigured sets PayloadConfigured field to given value.

### HasPayloadConfigured

`func (o *AccountSettingsResponse) HasPayloadConfigured() bool`

HasPayloadConfigured returns a boolean if a field has been set.

### GetLocalAdminAccountEnabled

`func (o *AccountSettingsResponse) GetLocalAdminAccountEnabled() bool`

GetLocalAdminAccountEnabled returns the LocalAdminAccountEnabled field if non-nil, zero value otherwise.

### GetLocalAdminAccountEnabledOk

`func (o *AccountSettingsResponse) GetLocalAdminAccountEnabledOk() (*bool, bool)`

GetLocalAdminAccountEnabledOk returns a tuple with the LocalAdminAccountEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAdminAccountEnabled

`func (o *AccountSettingsResponse) SetLocalAdminAccountEnabled(v bool)`

SetLocalAdminAccountEnabled sets LocalAdminAccountEnabled field to given value.

### HasLocalAdminAccountEnabled

`func (o *AccountSettingsResponse) HasLocalAdminAccountEnabled() bool`

HasLocalAdminAccountEnabled returns a boolean if a field has been set.

### GetAdminUsername

`func (o *AccountSettingsResponse) GetAdminUsername() string`

GetAdminUsername returns the AdminUsername field if non-nil, zero value otherwise.

### GetAdminUsernameOk

`func (o *AccountSettingsResponse) GetAdminUsernameOk() (*string, bool)`

GetAdminUsernameOk returns a tuple with the AdminUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminUsername

`func (o *AccountSettingsResponse) SetAdminUsername(v string)`

SetAdminUsername sets AdminUsername field to given value.

### HasAdminUsername

`func (o *AccountSettingsResponse) HasAdminUsername() bool`

HasAdminUsername returns a boolean if a field has been set.

### GetHiddenAdminAccount

`func (o *AccountSettingsResponse) GetHiddenAdminAccount() bool`

GetHiddenAdminAccount returns the HiddenAdminAccount field if non-nil, zero value otherwise.

### GetHiddenAdminAccountOk

`func (o *AccountSettingsResponse) GetHiddenAdminAccountOk() (*bool, bool)`

GetHiddenAdminAccountOk returns a tuple with the HiddenAdminAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenAdminAccount

`func (o *AccountSettingsResponse) SetHiddenAdminAccount(v bool)`

SetHiddenAdminAccount sets HiddenAdminAccount field to given value.

### HasHiddenAdminAccount

`func (o *AccountSettingsResponse) HasHiddenAdminAccount() bool`

HasHiddenAdminAccount returns a boolean if a field has been set.

### GetLocalUserManaged

`func (o *AccountSettingsResponse) GetLocalUserManaged() bool`

GetLocalUserManaged returns the LocalUserManaged field if non-nil, zero value otherwise.

### GetLocalUserManagedOk

`func (o *AccountSettingsResponse) GetLocalUserManagedOk() (*bool, bool)`

GetLocalUserManagedOk returns a tuple with the LocalUserManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalUserManaged

`func (o *AccountSettingsResponse) SetLocalUserManaged(v bool)`

SetLocalUserManaged sets LocalUserManaged field to given value.

### HasLocalUserManaged

`func (o *AccountSettingsResponse) HasLocalUserManaged() bool`

HasLocalUserManaged returns a boolean if a field has been set.

### GetUserAccountType

`func (o *AccountSettingsResponse) GetUserAccountType() string`

GetUserAccountType returns the UserAccountType field if non-nil, zero value otherwise.

### GetUserAccountTypeOk

`func (o *AccountSettingsResponse) GetUserAccountTypeOk() (*string, bool)`

GetUserAccountTypeOk returns a tuple with the UserAccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAccountType

`func (o *AccountSettingsResponse) SetUserAccountType(v string)`

SetUserAccountType sets UserAccountType field to given value.

### HasUserAccountType

`func (o *AccountSettingsResponse) HasUserAccountType() bool`

HasUserAccountType returns a boolean if a field has been set.

### GetVersionLock

`func (o *AccountSettingsResponse) GetVersionLock() int64`

GetVersionLock returns the VersionLock field if non-nil, zero value otherwise.

### GetVersionLockOk

`func (o *AccountSettingsResponse) GetVersionLockOk() (*int64, bool)`

GetVersionLockOk returns a tuple with the VersionLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionLock

`func (o *AccountSettingsResponse) SetVersionLock(v int64)`

SetVersionLock sets VersionLock field to given value.

### HasVersionLock

`func (o *AccountSettingsResponse) HasVersionLock() bool`

HasVersionLock returns a boolean if a field has been set.

### GetPrefillPrimaryAccountInfoFeatureEnabled

`func (o *AccountSettingsResponse) GetPrefillPrimaryAccountInfoFeatureEnabled() bool`

GetPrefillPrimaryAccountInfoFeatureEnabled returns the PrefillPrimaryAccountInfoFeatureEnabled field if non-nil, zero value otherwise.

### GetPrefillPrimaryAccountInfoFeatureEnabledOk

`func (o *AccountSettingsResponse) GetPrefillPrimaryAccountInfoFeatureEnabledOk() (*bool, bool)`

GetPrefillPrimaryAccountInfoFeatureEnabledOk returns a tuple with the PrefillPrimaryAccountInfoFeatureEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefillPrimaryAccountInfoFeatureEnabled

`func (o *AccountSettingsResponse) SetPrefillPrimaryAccountInfoFeatureEnabled(v bool)`

SetPrefillPrimaryAccountInfoFeatureEnabled sets PrefillPrimaryAccountInfoFeatureEnabled field to given value.

### HasPrefillPrimaryAccountInfoFeatureEnabled

`func (o *AccountSettingsResponse) HasPrefillPrimaryAccountInfoFeatureEnabled() bool`

HasPrefillPrimaryAccountInfoFeatureEnabled returns a boolean if a field has been set.

### GetPrefillType

`func (o *AccountSettingsResponse) GetPrefillType() string`

GetPrefillType returns the PrefillType field if non-nil, zero value otherwise.

### GetPrefillTypeOk

`func (o *AccountSettingsResponse) GetPrefillTypeOk() (*string, bool)`

GetPrefillTypeOk returns a tuple with the PrefillType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefillType

`func (o *AccountSettingsResponse) SetPrefillType(v string)`

SetPrefillType sets PrefillType field to given value.

### HasPrefillType

`func (o *AccountSettingsResponse) HasPrefillType() bool`

HasPrefillType returns a boolean if a field has been set.

### GetPrefillAccountFullName

`func (o *AccountSettingsResponse) GetPrefillAccountFullName() string`

GetPrefillAccountFullName returns the PrefillAccountFullName field if non-nil, zero value otherwise.

### GetPrefillAccountFullNameOk

`func (o *AccountSettingsResponse) GetPrefillAccountFullNameOk() (*string, bool)`

GetPrefillAccountFullNameOk returns a tuple with the PrefillAccountFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefillAccountFullName

`func (o *AccountSettingsResponse) SetPrefillAccountFullName(v string)`

SetPrefillAccountFullName sets PrefillAccountFullName field to given value.

### HasPrefillAccountFullName

`func (o *AccountSettingsResponse) HasPrefillAccountFullName() bool`

HasPrefillAccountFullName returns a boolean if a field has been set.

### GetPrefillAccountUserName

`func (o *AccountSettingsResponse) GetPrefillAccountUserName() string`

GetPrefillAccountUserName returns the PrefillAccountUserName field if non-nil, zero value otherwise.

### GetPrefillAccountUserNameOk

`func (o *AccountSettingsResponse) GetPrefillAccountUserNameOk() (*string, bool)`

GetPrefillAccountUserNameOk returns a tuple with the PrefillAccountUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefillAccountUserName

`func (o *AccountSettingsResponse) SetPrefillAccountUserName(v string)`

SetPrefillAccountUserName sets PrefillAccountUserName field to given value.

### HasPrefillAccountUserName

`func (o *AccountSettingsResponse) HasPrefillAccountUserName() bool`

HasPrefillAccountUserName returns a boolean if a field has been set.

### GetPreventPrefillInfoFromModification

`func (o *AccountSettingsResponse) GetPreventPrefillInfoFromModification() bool`

GetPreventPrefillInfoFromModification returns the PreventPrefillInfoFromModification field if non-nil, zero value otherwise.

### GetPreventPrefillInfoFromModificationOk

`func (o *AccountSettingsResponse) GetPreventPrefillInfoFromModificationOk() (*bool, bool)`

GetPreventPrefillInfoFromModificationOk returns a tuple with the PreventPrefillInfoFromModification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventPrefillInfoFromModification

`func (o *AccountSettingsResponse) SetPreventPrefillInfoFromModification(v bool)`

SetPreventPrefillInfoFromModification sets PreventPrefillInfoFromModification field to given value.

### HasPreventPrefillInfoFromModification

`func (o *AccountSettingsResponse) HasPreventPrefillInfoFromModification() bool`

HasPreventPrefillInfoFromModification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


