# ComputerLocalUserAccountCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | Pointer to **NullableString** |  | [optional] 
**UserGuid** | Pointer to **NullableString** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**FullName** | Pointer to **NullableString** |  | [optional] 
**Admin** | Pointer to **NullableBool** |  | [optional] 
**HomeDirectory** | Pointer to **NullableString** |  | [optional] 
**HomeDirectorySizeMb** | Pointer to **NullableInt64** | Home directory size in MB. | [optional] 
**FileVault2Enabled** | Pointer to **NullableBool** |  | [optional] 
**UserAccountType** | Pointer to **NullableString** |  | [optional] 
**PasswordMinLength** | Pointer to **NullableInt64** |  | [optional] 
**PasswordMaxAge** | Pointer to **NullableInt64** |  | [optional] 
**PasswordMinComplexCharacters** | Pointer to **NullableInt64** |  | [optional] 
**PasswordHistoryDepth** | Pointer to **NullableInt64** |  | [optional] 
**PasswordRequireAlphanumeric** | Pointer to **NullableBool** |  | [optional] 
**ComputerAzureActiveDirectoryId** | Pointer to **NullableString** |  | [optional] 
**UserAzureActiveDirectoryId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewComputerLocalUserAccountCreate

`func NewComputerLocalUserAccountCreate() *ComputerLocalUserAccountCreate`

NewComputerLocalUserAccountCreate instantiates a new ComputerLocalUserAccountCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputerLocalUserAccountCreateWithDefaults

`func NewComputerLocalUserAccountCreateWithDefaults() *ComputerLocalUserAccountCreate`

NewComputerLocalUserAccountCreateWithDefaults instantiates a new ComputerLocalUserAccountCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *ComputerLocalUserAccountCreate) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ComputerLocalUserAccountCreate) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ComputerLocalUserAccountCreate) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *ComputerLocalUserAccountCreate) HasUid() bool`

HasUid returns a boolean if a field has been set.

### SetUidNil

`func (o *ComputerLocalUserAccountCreate) SetUidNil(b bool)`

 SetUidNil sets the value for Uid to be an explicit nil

### UnsetUid
`func (o *ComputerLocalUserAccountCreate) UnsetUid()`

UnsetUid ensures that no value is present for Uid, not even an explicit nil
### GetUserGuid

`func (o *ComputerLocalUserAccountCreate) GetUserGuid() string`

GetUserGuid returns the UserGuid field if non-nil, zero value otherwise.

### GetUserGuidOk

`func (o *ComputerLocalUserAccountCreate) GetUserGuidOk() (*string, bool)`

GetUserGuidOk returns a tuple with the UserGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGuid

`func (o *ComputerLocalUserAccountCreate) SetUserGuid(v string)`

SetUserGuid sets UserGuid field to given value.

### HasUserGuid

`func (o *ComputerLocalUserAccountCreate) HasUserGuid() bool`

HasUserGuid returns a boolean if a field has been set.

### SetUserGuidNil

`func (o *ComputerLocalUserAccountCreate) SetUserGuidNil(b bool)`

 SetUserGuidNil sets the value for UserGuid to be an explicit nil

### UnsetUserGuid
`func (o *ComputerLocalUserAccountCreate) UnsetUserGuid()`

UnsetUserGuid ensures that no value is present for UserGuid, not even an explicit nil
### GetUsername

`func (o *ComputerLocalUserAccountCreate) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ComputerLocalUserAccountCreate) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ComputerLocalUserAccountCreate) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ComputerLocalUserAccountCreate) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *ComputerLocalUserAccountCreate) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *ComputerLocalUserAccountCreate) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetFullName

`func (o *ComputerLocalUserAccountCreate) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *ComputerLocalUserAccountCreate) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *ComputerLocalUserAccountCreate) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *ComputerLocalUserAccountCreate) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *ComputerLocalUserAccountCreate) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *ComputerLocalUserAccountCreate) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetAdmin

`func (o *ComputerLocalUserAccountCreate) GetAdmin() bool`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *ComputerLocalUserAccountCreate) GetAdminOk() (*bool, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *ComputerLocalUserAccountCreate) SetAdmin(v bool)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *ComputerLocalUserAccountCreate) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.

### SetAdminNil

`func (o *ComputerLocalUserAccountCreate) SetAdminNil(b bool)`

 SetAdminNil sets the value for Admin to be an explicit nil

### UnsetAdmin
`func (o *ComputerLocalUserAccountCreate) UnsetAdmin()`

UnsetAdmin ensures that no value is present for Admin, not even an explicit nil
### GetHomeDirectory

`func (o *ComputerLocalUserAccountCreate) GetHomeDirectory() string`

GetHomeDirectory returns the HomeDirectory field if non-nil, zero value otherwise.

### GetHomeDirectoryOk

`func (o *ComputerLocalUserAccountCreate) GetHomeDirectoryOk() (*string, bool)`

GetHomeDirectoryOk returns a tuple with the HomeDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeDirectory

`func (o *ComputerLocalUserAccountCreate) SetHomeDirectory(v string)`

SetHomeDirectory sets HomeDirectory field to given value.

### HasHomeDirectory

`func (o *ComputerLocalUserAccountCreate) HasHomeDirectory() bool`

HasHomeDirectory returns a boolean if a field has been set.

### SetHomeDirectoryNil

`func (o *ComputerLocalUserAccountCreate) SetHomeDirectoryNil(b bool)`

 SetHomeDirectoryNil sets the value for HomeDirectory to be an explicit nil

### UnsetHomeDirectory
`func (o *ComputerLocalUserAccountCreate) UnsetHomeDirectory()`

UnsetHomeDirectory ensures that no value is present for HomeDirectory, not even an explicit nil
### GetHomeDirectorySizeMb

`func (o *ComputerLocalUserAccountCreate) GetHomeDirectorySizeMb() int64`

GetHomeDirectorySizeMb returns the HomeDirectorySizeMb field if non-nil, zero value otherwise.

### GetHomeDirectorySizeMbOk

`func (o *ComputerLocalUserAccountCreate) GetHomeDirectorySizeMbOk() (*int64, bool)`

GetHomeDirectorySizeMbOk returns a tuple with the HomeDirectorySizeMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeDirectorySizeMb

`func (o *ComputerLocalUserAccountCreate) SetHomeDirectorySizeMb(v int64)`

SetHomeDirectorySizeMb sets HomeDirectorySizeMb field to given value.

### HasHomeDirectorySizeMb

`func (o *ComputerLocalUserAccountCreate) HasHomeDirectorySizeMb() bool`

HasHomeDirectorySizeMb returns a boolean if a field has been set.

### SetHomeDirectorySizeMbNil

`func (o *ComputerLocalUserAccountCreate) SetHomeDirectorySizeMbNil(b bool)`

 SetHomeDirectorySizeMbNil sets the value for HomeDirectorySizeMb to be an explicit nil

### UnsetHomeDirectorySizeMb
`func (o *ComputerLocalUserAccountCreate) UnsetHomeDirectorySizeMb()`

UnsetHomeDirectorySizeMb ensures that no value is present for HomeDirectorySizeMb, not even an explicit nil
### GetFileVault2Enabled

`func (o *ComputerLocalUserAccountCreate) GetFileVault2Enabled() bool`

GetFileVault2Enabled returns the FileVault2Enabled field if non-nil, zero value otherwise.

### GetFileVault2EnabledOk

`func (o *ComputerLocalUserAccountCreate) GetFileVault2EnabledOk() (*bool, bool)`

GetFileVault2EnabledOk returns a tuple with the FileVault2Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileVault2Enabled

`func (o *ComputerLocalUserAccountCreate) SetFileVault2Enabled(v bool)`

SetFileVault2Enabled sets FileVault2Enabled field to given value.

### HasFileVault2Enabled

`func (o *ComputerLocalUserAccountCreate) HasFileVault2Enabled() bool`

HasFileVault2Enabled returns a boolean if a field has been set.

### SetFileVault2EnabledNil

`func (o *ComputerLocalUserAccountCreate) SetFileVault2EnabledNil(b bool)`

 SetFileVault2EnabledNil sets the value for FileVault2Enabled to be an explicit nil

### UnsetFileVault2Enabled
`func (o *ComputerLocalUserAccountCreate) UnsetFileVault2Enabled()`

UnsetFileVault2Enabled ensures that no value is present for FileVault2Enabled, not even an explicit nil
### GetUserAccountType

`func (o *ComputerLocalUserAccountCreate) GetUserAccountType() string`

GetUserAccountType returns the UserAccountType field if non-nil, zero value otherwise.

### GetUserAccountTypeOk

`func (o *ComputerLocalUserAccountCreate) GetUserAccountTypeOk() (*string, bool)`

GetUserAccountTypeOk returns a tuple with the UserAccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAccountType

`func (o *ComputerLocalUserAccountCreate) SetUserAccountType(v string)`

SetUserAccountType sets UserAccountType field to given value.

### HasUserAccountType

`func (o *ComputerLocalUserAccountCreate) HasUserAccountType() bool`

HasUserAccountType returns a boolean if a field has been set.

### SetUserAccountTypeNil

`func (o *ComputerLocalUserAccountCreate) SetUserAccountTypeNil(b bool)`

 SetUserAccountTypeNil sets the value for UserAccountType to be an explicit nil

### UnsetUserAccountType
`func (o *ComputerLocalUserAccountCreate) UnsetUserAccountType()`

UnsetUserAccountType ensures that no value is present for UserAccountType, not even an explicit nil
### GetPasswordMinLength

`func (o *ComputerLocalUserAccountCreate) GetPasswordMinLength() int64`

GetPasswordMinLength returns the PasswordMinLength field if non-nil, zero value otherwise.

### GetPasswordMinLengthOk

`func (o *ComputerLocalUserAccountCreate) GetPasswordMinLengthOk() (*int64, bool)`

GetPasswordMinLengthOk returns a tuple with the PasswordMinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordMinLength

`func (o *ComputerLocalUserAccountCreate) SetPasswordMinLength(v int64)`

SetPasswordMinLength sets PasswordMinLength field to given value.

### HasPasswordMinLength

`func (o *ComputerLocalUserAccountCreate) HasPasswordMinLength() bool`

HasPasswordMinLength returns a boolean if a field has been set.

### SetPasswordMinLengthNil

`func (o *ComputerLocalUserAccountCreate) SetPasswordMinLengthNil(b bool)`

 SetPasswordMinLengthNil sets the value for PasswordMinLength to be an explicit nil

### UnsetPasswordMinLength
`func (o *ComputerLocalUserAccountCreate) UnsetPasswordMinLength()`

UnsetPasswordMinLength ensures that no value is present for PasswordMinLength, not even an explicit nil
### GetPasswordMaxAge

`func (o *ComputerLocalUserAccountCreate) GetPasswordMaxAge() int64`

GetPasswordMaxAge returns the PasswordMaxAge field if non-nil, zero value otherwise.

### GetPasswordMaxAgeOk

`func (o *ComputerLocalUserAccountCreate) GetPasswordMaxAgeOk() (*int64, bool)`

GetPasswordMaxAgeOk returns a tuple with the PasswordMaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordMaxAge

`func (o *ComputerLocalUserAccountCreate) SetPasswordMaxAge(v int64)`

SetPasswordMaxAge sets PasswordMaxAge field to given value.

### HasPasswordMaxAge

`func (o *ComputerLocalUserAccountCreate) HasPasswordMaxAge() bool`

HasPasswordMaxAge returns a boolean if a field has been set.

### SetPasswordMaxAgeNil

`func (o *ComputerLocalUserAccountCreate) SetPasswordMaxAgeNil(b bool)`

 SetPasswordMaxAgeNil sets the value for PasswordMaxAge to be an explicit nil

### UnsetPasswordMaxAge
`func (o *ComputerLocalUserAccountCreate) UnsetPasswordMaxAge()`

UnsetPasswordMaxAge ensures that no value is present for PasswordMaxAge, not even an explicit nil
### GetPasswordMinComplexCharacters

`func (o *ComputerLocalUserAccountCreate) GetPasswordMinComplexCharacters() int64`

GetPasswordMinComplexCharacters returns the PasswordMinComplexCharacters field if non-nil, zero value otherwise.

### GetPasswordMinComplexCharactersOk

`func (o *ComputerLocalUserAccountCreate) GetPasswordMinComplexCharactersOk() (*int64, bool)`

GetPasswordMinComplexCharactersOk returns a tuple with the PasswordMinComplexCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordMinComplexCharacters

`func (o *ComputerLocalUserAccountCreate) SetPasswordMinComplexCharacters(v int64)`

SetPasswordMinComplexCharacters sets PasswordMinComplexCharacters field to given value.

### HasPasswordMinComplexCharacters

`func (o *ComputerLocalUserAccountCreate) HasPasswordMinComplexCharacters() bool`

HasPasswordMinComplexCharacters returns a boolean if a field has been set.

### SetPasswordMinComplexCharactersNil

`func (o *ComputerLocalUserAccountCreate) SetPasswordMinComplexCharactersNil(b bool)`

 SetPasswordMinComplexCharactersNil sets the value for PasswordMinComplexCharacters to be an explicit nil

### UnsetPasswordMinComplexCharacters
`func (o *ComputerLocalUserAccountCreate) UnsetPasswordMinComplexCharacters()`

UnsetPasswordMinComplexCharacters ensures that no value is present for PasswordMinComplexCharacters, not even an explicit nil
### GetPasswordHistoryDepth

`func (o *ComputerLocalUserAccountCreate) GetPasswordHistoryDepth() int64`

GetPasswordHistoryDepth returns the PasswordHistoryDepth field if non-nil, zero value otherwise.

### GetPasswordHistoryDepthOk

`func (o *ComputerLocalUserAccountCreate) GetPasswordHistoryDepthOk() (*int64, bool)`

GetPasswordHistoryDepthOk returns a tuple with the PasswordHistoryDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHistoryDepth

`func (o *ComputerLocalUserAccountCreate) SetPasswordHistoryDepth(v int64)`

SetPasswordHistoryDepth sets PasswordHistoryDepth field to given value.

### HasPasswordHistoryDepth

`func (o *ComputerLocalUserAccountCreate) HasPasswordHistoryDepth() bool`

HasPasswordHistoryDepth returns a boolean if a field has been set.

### SetPasswordHistoryDepthNil

`func (o *ComputerLocalUserAccountCreate) SetPasswordHistoryDepthNil(b bool)`

 SetPasswordHistoryDepthNil sets the value for PasswordHistoryDepth to be an explicit nil

### UnsetPasswordHistoryDepth
`func (o *ComputerLocalUserAccountCreate) UnsetPasswordHistoryDepth()`

UnsetPasswordHistoryDepth ensures that no value is present for PasswordHistoryDepth, not even an explicit nil
### GetPasswordRequireAlphanumeric

`func (o *ComputerLocalUserAccountCreate) GetPasswordRequireAlphanumeric() bool`

GetPasswordRequireAlphanumeric returns the PasswordRequireAlphanumeric field if non-nil, zero value otherwise.

### GetPasswordRequireAlphanumericOk

`func (o *ComputerLocalUserAccountCreate) GetPasswordRequireAlphanumericOk() (*bool, bool)`

GetPasswordRequireAlphanumericOk returns a tuple with the PasswordRequireAlphanumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordRequireAlphanumeric

`func (o *ComputerLocalUserAccountCreate) SetPasswordRequireAlphanumeric(v bool)`

SetPasswordRequireAlphanumeric sets PasswordRequireAlphanumeric field to given value.

### HasPasswordRequireAlphanumeric

`func (o *ComputerLocalUserAccountCreate) HasPasswordRequireAlphanumeric() bool`

HasPasswordRequireAlphanumeric returns a boolean if a field has been set.

### SetPasswordRequireAlphanumericNil

`func (o *ComputerLocalUserAccountCreate) SetPasswordRequireAlphanumericNil(b bool)`

 SetPasswordRequireAlphanumericNil sets the value for PasswordRequireAlphanumeric to be an explicit nil

### UnsetPasswordRequireAlphanumeric
`func (o *ComputerLocalUserAccountCreate) UnsetPasswordRequireAlphanumeric()`

UnsetPasswordRequireAlphanumeric ensures that no value is present for PasswordRequireAlphanumeric, not even an explicit nil
### GetComputerAzureActiveDirectoryId

`func (o *ComputerLocalUserAccountCreate) GetComputerAzureActiveDirectoryId() string`

GetComputerAzureActiveDirectoryId returns the ComputerAzureActiveDirectoryId field if non-nil, zero value otherwise.

### GetComputerAzureActiveDirectoryIdOk

`func (o *ComputerLocalUserAccountCreate) GetComputerAzureActiveDirectoryIdOk() (*string, bool)`

GetComputerAzureActiveDirectoryIdOk returns a tuple with the ComputerAzureActiveDirectoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputerAzureActiveDirectoryId

`func (o *ComputerLocalUserAccountCreate) SetComputerAzureActiveDirectoryId(v string)`

SetComputerAzureActiveDirectoryId sets ComputerAzureActiveDirectoryId field to given value.

### HasComputerAzureActiveDirectoryId

`func (o *ComputerLocalUserAccountCreate) HasComputerAzureActiveDirectoryId() bool`

HasComputerAzureActiveDirectoryId returns a boolean if a field has been set.

### SetComputerAzureActiveDirectoryIdNil

`func (o *ComputerLocalUserAccountCreate) SetComputerAzureActiveDirectoryIdNil(b bool)`

 SetComputerAzureActiveDirectoryIdNil sets the value for ComputerAzureActiveDirectoryId to be an explicit nil

### UnsetComputerAzureActiveDirectoryId
`func (o *ComputerLocalUserAccountCreate) UnsetComputerAzureActiveDirectoryId()`

UnsetComputerAzureActiveDirectoryId ensures that no value is present for ComputerAzureActiveDirectoryId, not even an explicit nil
### GetUserAzureActiveDirectoryId

`func (o *ComputerLocalUserAccountCreate) GetUserAzureActiveDirectoryId() string`

GetUserAzureActiveDirectoryId returns the UserAzureActiveDirectoryId field if non-nil, zero value otherwise.

### GetUserAzureActiveDirectoryIdOk

`func (o *ComputerLocalUserAccountCreate) GetUserAzureActiveDirectoryIdOk() (*string, bool)`

GetUserAzureActiveDirectoryIdOk returns a tuple with the UserAzureActiveDirectoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAzureActiveDirectoryId

`func (o *ComputerLocalUserAccountCreate) SetUserAzureActiveDirectoryId(v string)`

SetUserAzureActiveDirectoryId sets UserAzureActiveDirectoryId field to given value.

### HasUserAzureActiveDirectoryId

`func (o *ComputerLocalUserAccountCreate) HasUserAzureActiveDirectoryId() bool`

HasUserAzureActiveDirectoryId returns a boolean if a field has been set.

### SetUserAzureActiveDirectoryIdNil

`func (o *ComputerLocalUserAccountCreate) SetUserAzureActiveDirectoryIdNil(b bool)`

 SetUserAzureActiveDirectoryIdNil sets the value for UserAzureActiveDirectoryId to be an explicit nil

### UnsetUserAzureActiveDirectoryId
`func (o *ComputerLocalUserAccountCreate) UnsetUserAzureActiveDirectoryId()`

UnsetUserAzureActiveDirectoryId ensures that no value is present for UserAzureActiveDirectoryId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


