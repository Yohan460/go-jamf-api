# ComputerLocalUserAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | Pointer to **string** |  | [optional] 
**UserGuid** | Pointer to **string** |  | [optional] [readonly] 
**Username** | Pointer to **string** |  | [optional] 
**FullName** | Pointer to **string** |  | [optional] 
**Admin** | Pointer to **bool** |  | [optional] 
**HomeDirectory** | Pointer to **string** |  | [optional] 
**HomeDirectorySizeMb** | Pointer to **int64** | Home directory size in MB. | [optional] [readonly] 
**FileVault2Enabled** | Pointer to **bool** |  | [optional] 
**UserAccountType** | Pointer to **string** |  | [optional] 
**PasswordMinLength** | Pointer to **int32** |  | [optional] 
**PasswordMaxAge** | Pointer to **int32** |  | [optional] 
**PasswordMinComplexCharacters** | Pointer to **int32** |  | [optional] 
**PasswordHistoryDepth** | Pointer to **int32** |  | [optional] 
**PasswordRequireAlphanumeric** | Pointer to **bool** |  | [optional] 
**ComputerAzureActiveDirectoryId** | Pointer to **string** |  | [optional] 
**UserAzureActiveDirectoryId** | Pointer to **string** |  | [optional] 
**AzureActiveDirectoryId** | Pointer to **string** |  | [optional] 

## Methods

### NewComputerLocalUserAccount

`func NewComputerLocalUserAccount() *ComputerLocalUserAccount`

NewComputerLocalUserAccount instantiates a new ComputerLocalUserAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputerLocalUserAccountWithDefaults

`func NewComputerLocalUserAccountWithDefaults() *ComputerLocalUserAccount`

NewComputerLocalUserAccountWithDefaults instantiates a new ComputerLocalUserAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *ComputerLocalUserAccount) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ComputerLocalUserAccount) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ComputerLocalUserAccount) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *ComputerLocalUserAccount) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetUserGuid

`func (o *ComputerLocalUserAccount) GetUserGuid() string`

GetUserGuid returns the UserGuid field if non-nil, zero value otherwise.

### GetUserGuidOk

`func (o *ComputerLocalUserAccount) GetUserGuidOk() (*string, bool)`

GetUserGuidOk returns a tuple with the UserGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGuid

`func (o *ComputerLocalUserAccount) SetUserGuid(v string)`

SetUserGuid sets UserGuid field to given value.

### HasUserGuid

`func (o *ComputerLocalUserAccount) HasUserGuid() bool`

HasUserGuid returns a boolean if a field has been set.

### GetUsername

`func (o *ComputerLocalUserAccount) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ComputerLocalUserAccount) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ComputerLocalUserAccount) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ComputerLocalUserAccount) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetFullName

`func (o *ComputerLocalUserAccount) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *ComputerLocalUserAccount) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *ComputerLocalUserAccount) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *ComputerLocalUserAccount) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetAdmin

`func (o *ComputerLocalUserAccount) GetAdmin() bool`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *ComputerLocalUserAccount) GetAdminOk() (*bool, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *ComputerLocalUserAccount) SetAdmin(v bool)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *ComputerLocalUserAccount) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.

### GetHomeDirectory

`func (o *ComputerLocalUserAccount) GetHomeDirectory() string`

GetHomeDirectory returns the HomeDirectory field if non-nil, zero value otherwise.

### GetHomeDirectoryOk

`func (o *ComputerLocalUserAccount) GetHomeDirectoryOk() (*string, bool)`

GetHomeDirectoryOk returns a tuple with the HomeDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeDirectory

`func (o *ComputerLocalUserAccount) SetHomeDirectory(v string)`

SetHomeDirectory sets HomeDirectory field to given value.

### HasHomeDirectory

`func (o *ComputerLocalUserAccount) HasHomeDirectory() bool`

HasHomeDirectory returns a boolean if a field has been set.

### GetHomeDirectorySizeMb

`func (o *ComputerLocalUserAccount) GetHomeDirectorySizeMb() int64`

GetHomeDirectorySizeMb returns the HomeDirectorySizeMb field if non-nil, zero value otherwise.

### GetHomeDirectorySizeMbOk

`func (o *ComputerLocalUserAccount) GetHomeDirectorySizeMbOk() (*int64, bool)`

GetHomeDirectorySizeMbOk returns a tuple with the HomeDirectorySizeMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeDirectorySizeMb

`func (o *ComputerLocalUserAccount) SetHomeDirectorySizeMb(v int64)`

SetHomeDirectorySizeMb sets HomeDirectorySizeMb field to given value.

### HasHomeDirectorySizeMb

`func (o *ComputerLocalUserAccount) HasHomeDirectorySizeMb() bool`

HasHomeDirectorySizeMb returns a boolean if a field has been set.

### GetFileVault2Enabled

`func (o *ComputerLocalUserAccount) GetFileVault2Enabled() bool`

GetFileVault2Enabled returns the FileVault2Enabled field if non-nil, zero value otherwise.

### GetFileVault2EnabledOk

`func (o *ComputerLocalUserAccount) GetFileVault2EnabledOk() (*bool, bool)`

GetFileVault2EnabledOk returns a tuple with the FileVault2Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileVault2Enabled

`func (o *ComputerLocalUserAccount) SetFileVault2Enabled(v bool)`

SetFileVault2Enabled sets FileVault2Enabled field to given value.

### HasFileVault2Enabled

`func (o *ComputerLocalUserAccount) HasFileVault2Enabled() bool`

HasFileVault2Enabled returns a boolean if a field has been set.

### GetUserAccountType

`func (o *ComputerLocalUserAccount) GetUserAccountType() string`

GetUserAccountType returns the UserAccountType field if non-nil, zero value otherwise.

### GetUserAccountTypeOk

`func (o *ComputerLocalUserAccount) GetUserAccountTypeOk() (*string, bool)`

GetUserAccountTypeOk returns a tuple with the UserAccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAccountType

`func (o *ComputerLocalUserAccount) SetUserAccountType(v string)`

SetUserAccountType sets UserAccountType field to given value.

### HasUserAccountType

`func (o *ComputerLocalUserAccount) HasUserAccountType() bool`

HasUserAccountType returns a boolean if a field has been set.

### GetPasswordMinLength

`func (o *ComputerLocalUserAccount) GetPasswordMinLength() int32`

GetPasswordMinLength returns the PasswordMinLength field if non-nil, zero value otherwise.

### GetPasswordMinLengthOk

`func (o *ComputerLocalUserAccount) GetPasswordMinLengthOk() (*int32, bool)`

GetPasswordMinLengthOk returns a tuple with the PasswordMinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordMinLength

`func (o *ComputerLocalUserAccount) SetPasswordMinLength(v int32)`

SetPasswordMinLength sets PasswordMinLength field to given value.

### HasPasswordMinLength

`func (o *ComputerLocalUserAccount) HasPasswordMinLength() bool`

HasPasswordMinLength returns a boolean if a field has been set.

### GetPasswordMaxAge

`func (o *ComputerLocalUserAccount) GetPasswordMaxAge() int32`

GetPasswordMaxAge returns the PasswordMaxAge field if non-nil, zero value otherwise.

### GetPasswordMaxAgeOk

`func (o *ComputerLocalUserAccount) GetPasswordMaxAgeOk() (*int32, bool)`

GetPasswordMaxAgeOk returns a tuple with the PasswordMaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordMaxAge

`func (o *ComputerLocalUserAccount) SetPasswordMaxAge(v int32)`

SetPasswordMaxAge sets PasswordMaxAge field to given value.

### HasPasswordMaxAge

`func (o *ComputerLocalUserAccount) HasPasswordMaxAge() bool`

HasPasswordMaxAge returns a boolean if a field has been set.

### GetPasswordMinComplexCharacters

`func (o *ComputerLocalUserAccount) GetPasswordMinComplexCharacters() int32`

GetPasswordMinComplexCharacters returns the PasswordMinComplexCharacters field if non-nil, zero value otherwise.

### GetPasswordMinComplexCharactersOk

`func (o *ComputerLocalUserAccount) GetPasswordMinComplexCharactersOk() (*int32, bool)`

GetPasswordMinComplexCharactersOk returns a tuple with the PasswordMinComplexCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordMinComplexCharacters

`func (o *ComputerLocalUserAccount) SetPasswordMinComplexCharacters(v int32)`

SetPasswordMinComplexCharacters sets PasswordMinComplexCharacters field to given value.

### HasPasswordMinComplexCharacters

`func (o *ComputerLocalUserAccount) HasPasswordMinComplexCharacters() bool`

HasPasswordMinComplexCharacters returns a boolean if a field has been set.

### GetPasswordHistoryDepth

`func (o *ComputerLocalUserAccount) GetPasswordHistoryDepth() int32`

GetPasswordHistoryDepth returns the PasswordHistoryDepth field if non-nil, zero value otherwise.

### GetPasswordHistoryDepthOk

`func (o *ComputerLocalUserAccount) GetPasswordHistoryDepthOk() (*int32, bool)`

GetPasswordHistoryDepthOk returns a tuple with the PasswordHistoryDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHistoryDepth

`func (o *ComputerLocalUserAccount) SetPasswordHistoryDepth(v int32)`

SetPasswordHistoryDepth sets PasswordHistoryDepth field to given value.

### HasPasswordHistoryDepth

`func (o *ComputerLocalUserAccount) HasPasswordHistoryDepth() bool`

HasPasswordHistoryDepth returns a boolean if a field has been set.

### GetPasswordRequireAlphanumeric

`func (o *ComputerLocalUserAccount) GetPasswordRequireAlphanumeric() bool`

GetPasswordRequireAlphanumeric returns the PasswordRequireAlphanumeric field if non-nil, zero value otherwise.

### GetPasswordRequireAlphanumericOk

`func (o *ComputerLocalUserAccount) GetPasswordRequireAlphanumericOk() (*bool, bool)`

GetPasswordRequireAlphanumericOk returns a tuple with the PasswordRequireAlphanumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordRequireAlphanumeric

`func (o *ComputerLocalUserAccount) SetPasswordRequireAlphanumeric(v bool)`

SetPasswordRequireAlphanumeric sets PasswordRequireAlphanumeric field to given value.

### HasPasswordRequireAlphanumeric

`func (o *ComputerLocalUserAccount) HasPasswordRequireAlphanumeric() bool`

HasPasswordRequireAlphanumeric returns a boolean if a field has been set.

### GetComputerAzureActiveDirectoryId

`func (o *ComputerLocalUserAccount) GetComputerAzureActiveDirectoryId() string`

GetComputerAzureActiveDirectoryId returns the ComputerAzureActiveDirectoryId field if non-nil, zero value otherwise.

### GetComputerAzureActiveDirectoryIdOk

`func (o *ComputerLocalUserAccount) GetComputerAzureActiveDirectoryIdOk() (*string, bool)`

GetComputerAzureActiveDirectoryIdOk returns a tuple with the ComputerAzureActiveDirectoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputerAzureActiveDirectoryId

`func (o *ComputerLocalUserAccount) SetComputerAzureActiveDirectoryId(v string)`

SetComputerAzureActiveDirectoryId sets ComputerAzureActiveDirectoryId field to given value.

### HasComputerAzureActiveDirectoryId

`func (o *ComputerLocalUserAccount) HasComputerAzureActiveDirectoryId() bool`

HasComputerAzureActiveDirectoryId returns a boolean if a field has been set.

### GetUserAzureActiveDirectoryId

`func (o *ComputerLocalUserAccount) GetUserAzureActiveDirectoryId() string`

GetUserAzureActiveDirectoryId returns the UserAzureActiveDirectoryId field if non-nil, zero value otherwise.

### GetUserAzureActiveDirectoryIdOk

`func (o *ComputerLocalUserAccount) GetUserAzureActiveDirectoryIdOk() (*string, bool)`

GetUserAzureActiveDirectoryIdOk returns a tuple with the UserAzureActiveDirectoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAzureActiveDirectoryId

`func (o *ComputerLocalUserAccount) SetUserAzureActiveDirectoryId(v string)`

SetUserAzureActiveDirectoryId sets UserAzureActiveDirectoryId field to given value.

### HasUserAzureActiveDirectoryId

`func (o *ComputerLocalUserAccount) HasUserAzureActiveDirectoryId() bool`

HasUserAzureActiveDirectoryId returns a boolean if a field has been set.

### GetAzureActiveDirectoryId

`func (o *ComputerLocalUserAccount) GetAzureActiveDirectoryId() string`

GetAzureActiveDirectoryId returns the AzureActiveDirectoryId field if non-nil, zero value otherwise.

### GetAzureActiveDirectoryIdOk

`func (o *ComputerLocalUserAccount) GetAzureActiveDirectoryIdOk() (*string, bool)`

GetAzureActiveDirectoryIdOk returns a tuple with the AzureActiveDirectoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureActiveDirectoryId

`func (o *ComputerLocalUserAccount) SetAzureActiveDirectoryId(v string)`

SetAzureActiveDirectoryId sets AzureActiveDirectoryId field to given value.

### HasAzureActiveDirectoryId

`func (o *ComputerLocalUserAccount) HasAzureActiveDirectoryId() bool`

HasAzureActiveDirectoryId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


