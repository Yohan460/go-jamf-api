# AccountSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | id of Account Settings | [optional] 
**PayloadConfigured** | Pointer to **bool** |  | [optional] [default to false]
**LocalAdminAccountEnabled** | Pointer to **bool** |  | [optional] [default to false]
**AdminUsername** | Pointer to **string** |  | [optional] [default to ""]
**AdminPassword** | Pointer to **string** |  | [optional] [default to ""]
**HiddenAdminAccount** | Pointer to **bool** |  | [optional] [default to false]
**LocalUserManaged** | Pointer to **bool** |  | [optional] [default to false]
**UserAccountType** | Pointer to **string** |  | [optional] [default to "STANDARD"]
**VersionLock** | Pointer to **int32** |  | [optional] [default to 0]
**PrefillPrimaryAccountInfoFeatureEnabled** | Pointer to **bool** |  | [optional] [default to false]
**PrefillType** | Pointer to **string** | Values accepted are only CUSTOM and DEVICE_OWNER | [optional] [default to "CUSTOM"]
**PrefillAccountFullName** | Pointer to **string** |  | [optional] [default to ""]
**PrefillAccountUserName** | Pointer to **string** |  | [optional] [default to ""]
**PreventPrefillInfoFromModification** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewAccountSettingsRequest

`func NewAccountSettingsRequest() *AccountSettingsRequest`

NewAccountSettingsRequest instantiates a new AccountSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountSettingsRequestWithDefaults

`func NewAccountSettingsRequestWithDefaults() *AccountSettingsRequest`

NewAccountSettingsRequestWithDefaults instantiates a new AccountSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountSettingsRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountSettingsRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountSettingsRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccountSettingsRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPayloadConfigured

`func (o *AccountSettingsRequest) GetPayloadConfigured() bool`

GetPayloadConfigured returns the PayloadConfigured field if non-nil, zero value otherwise.

### GetPayloadConfiguredOk

`func (o *AccountSettingsRequest) GetPayloadConfiguredOk() (*bool, bool)`

GetPayloadConfiguredOk returns a tuple with the PayloadConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadConfigured

`func (o *AccountSettingsRequest) SetPayloadConfigured(v bool)`

SetPayloadConfigured sets PayloadConfigured field to given value.

### HasPayloadConfigured

`func (o *AccountSettingsRequest) HasPayloadConfigured() bool`

HasPayloadConfigured returns a boolean if a field has been set.

### GetLocalAdminAccountEnabled

`func (o *AccountSettingsRequest) GetLocalAdminAccountEnabled() bool`

GetLocalAdminAccountEnabled returns the LocalAdminAccountEnabled field if non-nil, zero value otherwise.

### GetLocalAdminAccountEnabledOk

`func (o *AccountSettingsRequest) GetLocalAdminAccountEnabledOk() (*bool, bool)`

GetLocalAdminAccountEnabledOk returns a tuple with the LocalAdminAccountEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAdminAccountEnabled

`func (o *AccountSettingsRequest) SetLocalAdminAccountEnabled(v bool)`

SetLocalAdminAccountEnabled sets LocalAdminAccountEnabled field to given value.

### HasLocalAdminAccountEnabled

`func (o *AccountSettingsRequest) HasLocalAdminAccountEnabled() bool`

HasLocalAdminAccountEnabled returns a boolean if a field has been set.

### GetAdminUsername

`func (o *AccountSettingsRequest) GetAdminUsername() string`

GetAdminUsername returns the AdminUsername field if non-nil, zero value otherwise.

### GetAdminUsernameOk

`func (o *AccountSettingsRequest) GetAdminUsernameOk() (*string, bool)`

GetAdminUsernameOk returns a tuple with the AdminUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminUsername

`func (o *AccountSettingsRequest) SetAdminUsername(v string)`

SetAdminUsername sets AdminUsername field to given value.

### HasAdminUsername

`func (o *AccountSettingsRequest) HasAdminUsername() bool`

HasAdminUsername returns a boolean if a field has been set.

### GetAdminPassword

`func (o *AccountSettingsRequest) GetAdminPassword() string`

GetAdminPassword returns the AdminPassword field if non-nil, zero value otherwise.

### GetAdminPasswordOk

`func (o *AccountSettingsRequest) GetAdminPasswordOk() (*string, bool)`

GetAdminPasswordOk returns a tuple with the AdminPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPassword

`func (o *AccountSettingsRequest) SetAdminPassword(v string)`

SetAdminPassword sets AdminPassword field to given value.

### HasAdminPassword

`func (o *AccountSettingsRequest) HasAdminPassword() bool`

HasAdminPassword returns a boolean if a field has been set.

### GetHiddenAdminAccount

`func (o *AccountSettingsRequest) GetHiddenAdminAccount() bool`

GetHiddenAdminAccount returns the HiddenAdminAccount field if non-nil, zero value otherwise.

### GetHiddenAdminAccountOk

`func (o *AccountSettingsRequest) GetHiddenAdminAccountOk() (*bool, bool)`

GetHiddenAdminAccountOk returns a tuple with the HiddenAdminAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenAdminAccount

`func (o *AccountSettingsRequest) SetHiddenAdminAccount(v bool)`

SetHiddenAdminAccount sets HiddenAdminAccount field to given value.

### HasHiddenAdminAccount

`func (o *AccountSettingsRequest) HasHiddenAdminAccount() bool`

HasHiddenAdminAccount returns a boolean if a field has been set.

### GetLocalUserManaged

`func (o *AccountSettingsRequest) GetLocalUserManaged() bool`

GetLocalUserManaged returns the LocalUserManaged field if non-nil, zero value otherwise.

### GetLocalUserManagedOk

`func (o *AccountSettingsRequest) GetLocalUserManagedOk() (*bool, bool)`

GetLocalUserManagedOk returns a tuple with the LocalUserManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalUserManaged

`func (o *AccountSettingsRequest) SetLocalUserManaged(v bool)`

SetLocalUserManaged sets LocalUserManaged field to given value.

### HasLocalUserManaged

`func (o *AccountSettingsRequest) HasLocalUserManaged() bool`

HasLocalUserManaged returns a boolean if a field has been set.

### GetUserAccountType

`func (o *AccountSettingsRequest) GetUserAccountType() string`

GetUserAccountType returns the UserAccountType field if non-nil, zero value otherwise.

### GetUserAccountTypeOk

`func (o *AccountSettingsRequest) GetUserAccountTypeOk() (*string, bool)`

GetUserAccountTypeOk returns a tuple with the UserAccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAccountType

`func (o *AccountSettingsRequest) SetUserAccountType(v string)`

SetUserAccountType sets UserAccountType field to given value.

### HasUserAccountType

`func (o *AccountSettingsRequest) HasUserAccountType() bool`

HasUserAccountType returns a boolean if a field has been set.

### GetVersionLock

`func (o *AccountSettingsRequest) GetVersionLock() int32`

GetVersionLock returns the VersionLock field if non-nil, zero value otherwise.

### GetVersionLockOk

`func (o *AccountSettingsRequest) GetVersionLockOk() (*int32, bool)`

GetVersionLockOk returns a tuple with the VersionLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionLock

`func (o *AccountSettingsRequest) SetVersionLock(v int32)`

SetVersionLock sets VersionLock field to given value.

### HasVersionLock

`func (o *AccountSettingsRequest) HasVersionLock() bool`

HasVersionLock returns a boolean if a field has been set.

### GetPrefillPrimaryAccountInfoFeatureEnabled

`func (o *AccountSettingsRequest) GetPrefillPrimaryAccountInfoFeatureEnabled() bool`

GetPrefillPrimaryAccountInfoFeatureEnabled returns the PrefillPrimaryAccountInfoFeatureEnabled field if non-nil, zero value otherwise.

### GetPrefillPrimaryAccountInfoFeatureEnabledOk

`func (o *AccountSettingsRequest) GetPrefillPrimaryAccountInfoFeatureEnabledOk() (*bool, bool)`

GetPrefillPrimaryAccountInfoFeatureEnabledOk returns a tuple with the PrefillPrimaryAccountInfoFeatureEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefillPrimaryAccountInfoFeatureEnabled

`func (o *AccountSettingsRequest) SetPrefillPrimaryAccountInfoFeatureEnabled(v bool)`

SetPrefillPrimaryAccountInfoFeatureEnabled sets PrefillPrimaryAccountInfoFeatureEnabled field to given value.

### HasPrefillPrimaryAccountInfoFeatureEnabled

`func (o *AccountSettingsRequest) HasPrefillPrimaryAccountInfoFeatureEnabled() bool`

HasPrefillPrimaryAccountInfoFeatureEnabled returns a boolean if a field has been set.

### GetPrefillType

`func (o *AccountSettingsRequest) GetPrefillType() string`

GetPrefillType returns the PrefillType field if non-nil, zero value otherwise.

### GetPrefillTypeOk

`func (o *AccountSettingsRequest) GetPrefillTypeOk() (*string, bool)`

GetPrefillTypeOk returns a tuple with the PrefillType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefillType

`func (o *AccountSettingsRequest) SetPrefillType(v string)`

SetPrefillType sets PrefillType field to given value.

### HasPrefillType

`func (o *AccountSettingsRequest) HasPrefillType() bool`

HasPrefillType returns a boolean if a field has been set.

### GetPrefillAccountFullName

`func (o *AccountSettingsRequest) GetPrefillAccountFullName() string`

GetPrefillAccountFullName returns the PrefillAccountFullName field if non-nil, zero value otherwise.

### GetPrefillAccountFullNameOk

`func (o *AccountSettingsRequest) GetPrefillAccountFullNameOk() (*string, bool)`

GetPrefillAccountFullNameOk returns a tuple with the PrefillAccountFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefillAccountFullName

`func (o *AccountSettingsRequest) SetPrefillAccountFullName(v string)`

SetPrefillAccountFullName sets PrefillAccountFullName field to given value.

### HasPrefillAccountFullName

`func (o *AccountSettingsRequest) HasPrefillAccountFullName() bool`

HasPrefillAccountFullName returns a boolean if a field has been set.

### GetPrefillAccountUserName

`func (o *AccountSettingsRequest) GetPrefillAccountUserName() string`

GetPrefillAccountUserName returns the PrefillAccountUserName field if non-nil, zero value otherwise.

### GetPrefillAccountUserNameOk

`func (o *AccountSettingsRequest) GetPrefillAccountUserNameOk() (*string, bool)`

GetPrefillAccountUserNameOk returns a tuple with the PrefillAccountUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefillAccountUserName

`func (o *AccountSettingsRequest) SetPrefillAccountUserName(v string)`

SetPrefillAccountUserName sets PrefillAccountUserName field to given value.

### HasPrefillAccountUserName

`func (o *AccountSettingsRequest) HasPrefillAccountUserName() bool`

HasPrefillAccountUserName returns a boolean if a field has been set.

### GetPreventPrefillInfoFromModification

`func (o *AccountSettingsRequest) GetPreventPrefillInfoFromModification() bool`

GetPreventPrefillInfoFromModification returns the PreventPrefillInfoFromModification field if non-nil, zero value otherwise.

### GetPreventPrefillInfoFromModificationOk

`func (o *AccountSettingsRequest) GetPreventPrefillInfoFromModificationOk() (*bool, bool)`

GetPreventPrefillInfoFromModificationOk returns a tuple with the PreventPrefillInfoFromModification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventPrefillInfoFromModification

`func (o *AccountSettingsRequest) SetPreventPrefillInfoFromModification(v bool)`

SetPreventPrefillInfoFromModification sets PreventPrefillInfoFromModification field to given value.

### HasPreventPrefillInfoFromModification

`func (o *AccountSettingsRequest) HasPreventPrefillInfoFromModification() bool`

HasPreventPrefillInfoFromModification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


