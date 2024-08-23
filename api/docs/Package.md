# Package

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**PackageName** | **string** |  | 
**FileName** | **string** |  | 
**CategoryId** | **string** |  | 
**Info** | Pointer to **NullableString** |  | [optional] 
**Notes** | Pointer to **NullableString** |  | [optional] 
**Priority** | **int64** |  | 
**OsRequirements** | Pointer to **NullableString** |  | [optional] 
**FillUserTemplate** | **bool** |  | 
**Indexed** | Pointer to **bool** |  | [optional] [readonly] 
**FillExistingUsers** | Pointer to **bool** |  | [optional] 
**Swu** | Pointer to **bool** |  | [optional] 
**RebootRequired** | **bool** |  | 
**SelfHealNotify** | Pointer to **bool** |  | [optional] 
**SelfHealingAction** | Pointer to **NullableString** |  | [optional] 
**OsInstall** | **bool** |  | 
**SerialNumber** | Pointer to **NullableString** |  | [optional] 
**ParentPackageId** | Pointer to **NullableString** |  | [optional] 
**BasePath** | Pointer to **NullableString** |  | [optional] 
**SuppressUpdates** | **bool** |  | 
**CloudTransferStatus** | Pointer to **string** |  | [optional] [readonly] 
**IgnoreConflicts** | Pointer to **bool** |  | [optional] 
**SuppressFromDock** | **bool** |  | 
**SuppressEula** | **bool** |  | 
**SuppressRegistration** | **bool** |  | 
**InstallLanguage** | Pointer to **NullableString** |  | [optional] 
**Md5** | Pointer to **NullableString** |  | [optional] 
**Sha256** | Pointer to **NullableString** |  | [optional] 
**HashType** | Pointer to **NullableString** |  | [optional] 
**HashValue** | Pointer to **NullableString** |  | [optional] 
**Size** | Pointer to **NullableString** |  | [optional] [readonly] 
**OsInstallerVersion** | Pointer to **NullableString** |  | [optional] 
**Manifest** | Pointer to **NullableString** |  | [optional] 
**ManifestFileName** | Pointer to **NullableString** |  | [optional] 
**Format** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPackage

`func NewPackage(packageName string, fileName string, categoryId string, priority int64, fillUserTemplate bool, rebootRequired bool, osInstall bool, suppressUpdates bool, suppressFromDock bool, suppressEula bool, suppressRegistration bool, ) *Package`

NewPackage instantiates a new Package object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageWithDefaults

`func NewPackageWithDefaults() *Package`

NewPackageWithDefaults instantiates a new Package object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Package) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Package) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Package) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Package) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPackageName

`func (o *Package) GetPackageName() string`

GetPackageName returns the PackageName field if non-nil, zero value otherwise.

### GetPackageNameOk

`func (o *Package) GetPackageNameOk() (*string, bool)`

GetPackageNameOk returns a tuple with the PackageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageName

`func (o *Package) SetPackageName(v string)`

SetPackageName sets PackageName field to given value.


### GetFileName

`func (o *Package) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *Package) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *Package) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetCategoryId

`func (o *Package) GetCategoryId() string`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *Package) GetCategoryIdOk() (*string, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *Package) SetCategoryId(v string)`

SetCategoryId sets CategoryId field to given value.


### GetInfo

`func (o *Package) GetInfo() string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *Package) GetInfoOk() (*string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *Package) SetInfo(v string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *Package) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### SetInfoNil

`func (o *Package) SetInfoNil(b bool)`

 SetInfoNil sets the value for Info to be an explicit nil

### UnsetInfo
`func (o *Package) UnsetInfo()`

UnsetInfo ensures that no value is present for Info, not even an explicit nil
### GetNotes

`func (o *Package) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Package) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Package) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *Package) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *Package) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *Package) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetPriority

`func (o *Package) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *Package) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *Package) SetPriority(v int64)`

SetPriority sets Priority field to given value.


### GetOsRequirements

`func (o *Package) GetOsRequirements() string`

GetOsRequirements returns the OsRequirements field if non-nil, zero value otherwise.

### GetOsRequirementsOk

`func (o *Package) GetOsRequirementsOk() (*string, bool)`

GetOsRequirementsOk returns a tuple with the OsRequirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsRequirements

`func (o *Package) SetOsRequirements(v string)`

SetOsRequirements sets OsRequirements field to given value.

### HasOsRequirements

`func (o *Package) HasOsRequirements() bool`

HasOsRequirements returns a boolean if a field has been set.

### SetOsRequirementsNil

`func (o *Package) SetOsRequirementsNil(b bool)`

 SetOsRequirementsNil sets the value for OsRequirements to be an explicit nil

### UnsetOsRequirements
`func (o *Package) UnsetOsRequirements()`

UnsetOsRequirements ensures that no value is present for OsRequirements, not even an explicit nil
### GetFillUserTemplate

`func (o *Package) GetFillUserTemplate() bool`

GetFillUserTemplate returns the FillUserTemplate field if non-nil, zero value otherwise.

### GetFillUserTemplateOk

`func (o *Package) GetFillUserTemplateOk() (*bool, bool)`

GetFillUserTemplateOk returns a tuple with the FillUserTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillUserTemplate

`func (o *Package) SetFillUserTemplate(v bool)`

SetFillUserTemplate sets FillUserTemplate field to given value.


### GetIndexed

`func (o *Package) GetIndexed() bool`

GetIndexed returns the Indexed field if non-nil, zero value otherwise.

### GetIndexedOk

`func (o *Package) GetIndexedOk() (*bool, bool)`

GetIndexedOk returns a tuple with the Indexed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexed

`func (o *Package) SetIndexed(v bool)`

SetIndexed sets Indexed field to given value.

### HasIndexed

`func (o *Package) HasIndexed() bool`

HasIndexed returns a boolean if a field has been set.

### GetFillExistingUsers

`func (o *Package) GetFillExistingUsers() bool`

GetFillExistingUsers returns the FillExistingUsers field if non-nil, zero value otherwise.

### GetFillExistingUsersOk

`func (o *Package) GetFillExistingUsersOk() (*bool, bool)`

GetFillExistingUsersOk returns a tuple with the FillExistingUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillExistingUsers

`func (o *Package) SetFillExistingUsers(v bool)`

SetFillExistingUsers sets FillExistingUsers field to given value.

### HasFillExistingUsers

`func (o *Package) HasFillExistingUsers() bool`

HasFillExistingUsers returns a boolean if a field has been set.

### GetSwu

`func (o *Package) GetSwu() bool`

GetSwu returns the Swu field if non-nil, zero value otherwise.

### GetSwuOk

`func (o *Package) GetSwuOk() (*bool, bool)`

GetSwuOk returns a tuple with the Swu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwu

`func (o *Package) SetSwu(v bool)`

SetSwu sets Swu field to given value.

### HasSwu

`func (o *Package) HasSwu() bool`

HasSwu returns a boolean if a field has been set.

### GetRebootRequired

`func (o *Package) GetRebootRequired() bool`

GetRebootRequired returns the RebootRequired field if non-nil, zero value otherwise.

### GetRebootRequiredOk

`func (o *Package) GetRebootRequiredOk() (*bool, bool)`

GetRebootRequiredOk returns a tuple with the RebootRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebootRequired

`func (o *Package) SetRebootRequired(v bool)`

SetRebootRequired sets RebootRequired field to given value.


### GetSelfHealNotify

`func (o *Package) GetSelfHealNotify() bool`

GetSelfHealNotify returns the SelfHealNotify field if non-nil, zero value otherwise.

### GetSelfHealNotifyOk

`func (o *Package) GetSelfHealNotifyOk() (*bool, bool)`

GetSelfHealNotifyOk returns a tuple with the SelfHealNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfHealNotify

`func (o *Package) SetSelfHealNotify(v bool)`

SetSelfHealNotify sets SelfHealNotify field to given value.

### HasSelfHealNotify

`func (o *Package) HasSelfHealNotify() bool`

HasSelfHealNotify returns a boolean if a field has been set.

### GetSelfHealingAction

`func (o *Package) GetSelfHealingAction() string`

GetSelfHealingAction returns the SelfHealingAction field if non-nil, zero value otherwise.

### GetSelfHealingActionOk

`func (o *Package) GetSelfHealingActionOk() (*string, bool)`

GetSelfHealingActionOk returns a tuple with the SelfHealingAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfHealingAction

`func (o *Package) SetSelfHealingAction(v string)`

SetSelfHealingAction sets SelfHealingAction field to given value.

### HasSelfHealingAction

`func (o *Package) HasSelfHealingAction() bool`

HasSelfHealingAction returns a boolean if a field has been set.

### SetSelfHealingActionNil

`func (o *Package) SetSelfHealingActionNil(b bool)`

 SetSelfHealingActionNil sets the value for SelfHealingAction to be an explicit nil

### UnsetSelfHealingAction
`func (o *Package) UnsetSelfHealingAction()`

UnsetSelfHealingAction ensures that no value is present for SelfHealingAction, not even an explicit nil
### GetOsInstall

`func (o *Package) GetOsInstall() bool`

GetOsInstall returns the OsInstall field if non-nil, zero value otherwise.

### GetOsInstallOk

`func (o *Package) GetOsInstallOk() (*bool, bool)`

GetOsInstallOk returns a tuple with the OsInstall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsInstall

`func (o *Package) SetOsInstall(v bool)`

SetOsInstall sets OsInstall field to given value.


### GetSerialNumber

`func (o *Package) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *Package) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *Package) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *Package) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### SetSerialNumberNil

`func (o *Package) SetSerialNumberNil(b bool)`

 SetSerialNumberNil sets the value for SerialNumber to be an explicit nil

### UnsetSerialNumber
`func (o *Package) UnsetSerialNumber()`

UnsetSerialNumber ensures that no value is present for SerialNumber, not even an explicit nil
### GetParentPackageId

`func (o *Package) GetParentPackageId() string`

GetParentPackageId returns the ParentPackageId field if non-nil, zero value otherwise.

### GetParentPackageIdOk

`func (o *Package) GetParentPackageIdOk() (*string, bool)`

GetParentPackageIdOk returns a tuple with the ParentPackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentPackageId

`func (o *Package) SetParentPackageId(v string)`

SetParentPackageId sets ParentPackageId field to given value.

### HasParentPackageId

`func (o *Package) HasParentPackageId() bool`

HasParentPackageId returns a boolean if a field has been set.

### SetParentPackageIdNil

`func (o *Package) SetParentPackageIdNil(b bool)`

 SetParentPackageIdNil sets the value for ParentPackageId to be an explicit nil

### UnsetParentPackageId
`func (o *Package) UnsetParentPackageId()`

UnsetParentPackageId ensures that no value is present for ParentPackageId, not even an explicit nil
### GetBasePath

`func (o *Package) GetBasePath() string`

GetBasePath returns the BasePath field if non-nil, zero value otherwise.

### GetBasePathOk

`func (o *Package) GetBasePathOk() (*string, bool)`

GetBasePathOk returns a tuple with the BasePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasePath

`func (o *Package) SetBasePath(v string)`

SetBasePath sets BasePath field to given value.

### HasBasePath

`func (o *Package) HasBasePath() bool`

HasBasePath returns a boolean if a field has been set.

### SetBasePathNil

`func (o *Package) SetBasePathNil(b bool)`

 SetBasePathNil sets the value for BasePath to be an explicit nil

### UnsetBasePath
`func (o *Package) UnsetBasePath()`

UnsetBasePath ensures that no value is present for BasePath, not even an explicit nil
### GetSuppressUpdates

`func (o *Package) GetSuppressUpdates() bool`

GetSuppressUpdates returns the SuppressUpdates field if non-nil, zero value otherwise.

### GetSuppressUpdatesOk

`func (o *Package) GetSuppressUpdatesOk() (*bool, bool)`

GetSuppressUpdatesOk returns a tuple with the SuppressUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressUpdates

`func (o *Package) SetSuppressUpdates(v bool)`

SetSuppressUpdates sets SuppressUpdates field to given value.


### GetCloudTransferStatus

`func (o *Package) GetCloudTransferStatus() string`

GetCloudTransferStatus returns the CloudTransferStatus field if non-nil, zero value otherwise.

### GetCloudTransferStatusOk

`func (o *Package) GetCloudTransferStatusOk() (*string, bool)`

GetCloudTransferStatusOk returns a tuple with the CloudTransferStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudTransferStatus

`func (o *Package) SetCloudTransferStatus(v string)`

SetCloudTransferStatus sets CloudTransferStatus field to given value.

### HasCloudTransferStatus

`func (o *Package) HasCloudTransferStatus() bool`

HasCloudTransferStatus returns a boolean if a field has been set.

### GetIgnoreConflicts

`func (o *Package) GetIgnoreConflicts() bool`

GetIgnoreConflicts returns the IgnoreConflicts field if non-nil, zero value otherwise.

### GetIgnoreConflictsOk

`func (o *Package) GetIgnoreConflictsOk() (*bool, bool)`

GetIgnoreConflictsOk returns a tuple with the IgnoreConflicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreConflicts

`func (o *Package) SetIgnoreConflicts(v bool)`

SetIgnoreConflicts sets IgnoreConflicts field to given value.

### HasIgnoreConflicts

`func (o *Package) HasIgnoreConflicts() bool`

HasIgnoreConflicts returns a boolean if a field has been set.

### GetSuppressFromDock

`func (o *Package) GetSuppressFromDock() bool`

GetSuppressFromDock returns the SuppressFromDock field if non-nil, zero value otherwise.

### GetSuppressFromDockOk

`func (o *Package) GetSuppressFromDockOk() (*bool, bool)`

GetSuppressFromDockOk returns a tuple with the SuppressFromDock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressFromDock

`func (o *Package) SetSuppressFromDock(v bool)`

SetSuppressFromDock sets SuppressFromDock field to given value.


### GetSuppressEula

`func (o *Package) GetSuppressEula() bool`

GetSuppressEula returns the SuppressEula field if non-nil, zero value otherwise.

### GetSuppressEulaOk

`func (o *Package) GetSuppressEulaOk() (*bool, bool)`

GetSuppressEulaOk returns a tuple with the SuppressEula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressEula

`func (o *Package) SetSuppressEula(v bool)`

SetSuppressEula sets SuppressEula field to given value.


### GetSuppressRegistration

`func (o *Package) GetSuppressRegistration() bool`

GetSuppressRegistration returns the SuppressRegistration field if non-nil, zero value otherwise.

### GetSuppressRegistrationOk

`func (o *Package) GetSuppressRegistrationOk() (*bool, bool)`

GetSuppressRegistrationOk returns a tuple with the SuppressRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressRegistration

`func (o *Package) SetSuppressRegistration(v bool)`

SetSuppressRegistration sets SuppressRegistration field to given value.


### GetInstallLanguage

`func (o *Package) GetInstallLanguage() string`

GetInstallLanguage returns the InstallLanguage field if non-nil, zero value otherwise.

### GetInstallLanguageOk

`func (o *Package) GetInstallLanguageOk() (*string, bool)`

GetInstallLanguageOk returns a tuple with the InstallLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallLanguage

`func (o *Package) SetInstallLanguage(v string)`

SetInstallLanguage sets InstallLanguage field to given value.

### HasInstallLanguage

`func (o *Package) HasInstallLanguage() bool`

HasInstallLanguage returns a boolean if a field has been set.

### SetInstallLanguageNil

`func (o *Package) SetInstallLanguageNil(b bool)`

 SetInstallLanguageNil sets the value for InstallLanguage to be an explicit nil

### UnsetInstallLanguage
`func (o *Package) UnsetInstallLanguage()`

UnsetInstallLanguage ensures that no value is present for InstallLanguage, not even an explicit nil
### GetMd5

`func (o *Package) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *Package) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *Package) SetMd5(v string)`

SetMd5 sets Md5 field to given value.

### HasMd5

`func (o *Package) HasMd5() bool`

HasMd5 returns a boolean if a field has been set.

### SetMd5Nil

`func (o *Package) SetMd5Nil(b bool)`

 SetMd5Nil sets the value for Md5 to be an explicit nil

### UnsetMd5
`func (o *Package) UnsetMd5()`

UnsetMd5 ensures that no value is present for Md5, not even an explicit nil
### GetSha256

`func (o *Package) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *Package) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *Package) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *Package) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.

### SetSha256Nil

`func (o *Package) SetSha256Nil(b bool)`

 SetSha256Nil sets the value for Sha256 to be an explicit nil

### UnsetSha256
`func (o *Package) UnsetSha256()`

UnsetSha256 ensures that no value is present for Sha256, not even an explicit nil
### GetHashType

`func (o *Package) GetHashType() string`

GetHashType returns the HashType field if non-nil, zero value otherwise.

### GetHashTypeOk

`func (o *Package) GetHashTypeOk() (*string, bool)`

GetHashTypeOk returns a tuple with the HashType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashType

`func (o *Package) SetHashType(v string)`

SetHashType sets HashType field to given value.

### HasHashType

`func (o *Package) HasHashType() bool`

HasHashType returns a boolean if a field has been set.

### SetHashTypeNil

`func (o *Package) SetHashTypeNil(b bool)`

 SetHashTypeNil sets the value for HashType to be an explicit nil

### UnsetHashType
`func (o *Package) UnsetHashType()`

UnsetHashType ensures that no value is present for HashType, not even an explicit nil
### GetHashValue

`func (o *Package) GetHashValue() string`

GetHashValue returns the HashValue field if non-nil, zero value otherwise.

### GetHashValueOk

`func (o *Package) GetHashValueOk() (*string, bool)`

GetHashValueOk returns a tuple with the HashValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashValue

`func (o *Package) SetHashValue(v string)`

SetHashValue sets HashValue field to given value.

### HasHashValue

`func (o *Package) HasHashValue() bool`

HasHashValue returns a boolean if a field has been set.

### SetHashValueNil

`func (o *Package) SetHashValueNil(b bool)`

 SetHashValueNil sets the value for HashValue to be an explicit nil

### UnsetHashValue
`func (o *Package) UnsetHashValue()`

UnsetHashValue ensures that no value is present for HashValue, not even an explicit nil
### GetSize

`func (o *Package) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Package) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Package) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *Package) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *Package) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *Package) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetOsInstallerVersion

`func (o *Package) GetOsInstallerVersion() string`

GetOsInstallerVersion returns the OsInstallerVersion field if non-nil, zero value otherwise.

### GetOsInstallerVersionOk

`func (o *Package) GetOsInstallerVersionOk() (*string, bool)`

GetOsInstallerVersionOk returns a tuple with the OsInstallerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsInstallerVersion

`func (o *Package) SetOsInstallerVersion(v string)`

SetOsInstallerVersion sets OsInstallerVersion field to given value.

### HasOsInstallerVersion

`func (o *Package) HasOsInstallerVersion() bool`

HasOsInstallerVersion returns a boolean if a field has been set.

### SetOsInstallerVersionNil

`func (o *Package) SetOsInstallerVersionNil(b bool)`

 SetOsInstallerVersionNil sets the value for OsInstallerVersion to be an explicit nil

### UnsetOsInstallerVersion
`func (o *Package) UnsetOsInstallerVersion()`

UnsetOsInstallerVersion ensures that no value is present for OsInstallerVersion, not even an explicit nil
### GetManifest

`func (o *Package) GetManifest() string`

GetManifest returns the Manifest field if non-nil, zero value otherwise.

### GetManifestOk

`func (o *Package) GetManifestOk() (*string, bool)`

GetManifestOk returns a tuple with the Manifest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifest

`func (o *Package) SetManifest(v string)`

SetManifest sets Manifest field to given value.

### HasManifest

`func (o *Package) HasManifest() bool`

HasManifest returns a boolean if a field has been set.

### SetManifestNil

`func (o *Package) SetManifestNil(b bool)`

 SetManifestNil sets the value for Manifest to be an explicit nil

### UnsetManifest
`func (o *Package) UnsetManifest()`

UnsetManifest ensures that no value is present for Manifest, not even an explicit nil
### GetManifestFileName

`func (o *Package) GetManifestFileName() string`

GetManifestFileName returns the ManifestFileName field if non-nil, zero value otherwise.

### GetManifestFileNameOk

`func (o *Package) GetManifestFileNameOk() (*string, bool)`

GetManifestFileNameOk returns a tuple with the ManifestFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifestFileName

`func (o *Package) SetManifestFileName(v string)`

SetManifestFileName sets ManifestFileName field to given value.

### HasManifestFileName

`func (o *Package) HasManifestFileName() bool`

HasManifestFileName returns a boolean if a field has been set.

### SetManifestFileNameNil

`func (o *Package) SetManifestFileNameNil(b bool)`

 SetManifestFileNameNil sets the value for ManifestFileName to be an explicit nil

### UnsetManifestFileName
`func (o *Package) UnsetManifestFileName()`

UnsetManifestFileName ensures that no value is present for ManifestFileName, not even an explicit nil
### GetFormat

`func (o *Package) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *Package) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *Package) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *Package) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormatNil

`func (o *Package) SetFormatNil(b bool)`

 SetFormatNil sets the value for Format to be an explicit nil

### UnsetFormat
`func (o *Package) UnsetFormat()`

UnsetFormat ensures that no value is present for Format, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


