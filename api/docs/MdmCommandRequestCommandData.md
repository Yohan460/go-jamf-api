# MdmCommandRequestCommandData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommandType** | Pointer to [**MdmCommandType**](MdmCommandType.md) |  | [optional] 
**LostModeMessage** | Pointer to **string** |  | [optional] 
**LostModePhone** | Pointer to **string** |  | [optional] 
**LostModeFootnote** | Pointer to **string** |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] 
**ForceDeletion** | Pointer to **bool** |  | [optional] 
**DeleteAllUsers** | Pointer to **bool** |  | [optional] 
**RequestType** | Pointer to **string** |  | [optional] 
**RequestRequiresNetworkTether** | Pointer to **bool** |  | [optional] 
**BootstrapTokenAllowed** | Pointer to **bool** |  | [optional] 
**Bluetooth** | Pointer to **bool** |  | [optional] 
**AppAnalytics** | Pointer to [**AppAnalyticsSetting**](AppAnalyticsSetting.md) |  | [optional] 
**DiagnosticSubmission** | Pointer to [**DiagnosticSubmissionSetting**](DiagnosticSubmissionSetting.md) |  | [optional] 
**DataRoaming** | Pointer to [**DataRoamingSetting**](DataRoamingSetting.md) |  | [optional] 
**VoiceRoaming** | Pointer to [**VoiceRoamingSetting**](VoiceRoamingSetting.md) |  | [optional] 
**PersonalHotspot** | Pointer to [**PersonalHotspotSetting**](PersonalHotspotSetting.md) |  | [optional] 
**MaximumResidentUsers** | Pointer to **int32** |  | [optional] 
**DeviceName** | Pointer to **string** |  | [optional] 
**ApplicationAttributes** | Pointer to [**ApplicationAttributes**](ApplicationAttributes.md) |  | [optional] 
**SharedDeviceConfiguration** | Pointer to [**SharedDeviceConfiguration**](SharedDeviceConfiguration.md) |  | [optional] 
**ApplicationConfiguration** | Pointer to [**ApplicationConfiguration**](ApplicationConfiguration.md) |  | [optional] 
**TimeZone** | Pointer to **string** |  | [optional] 
**SoftwareUpdateSettings** | Pointer to [**SoftwareUpdateSettings**](SoftwareUpdateSettings.md) |  | [optional] 
**RebuildKernelCache** | Pointer to **bool** |  | [optional] 
**KextPaths** | Pointer to **[]string** | Only used if RebuildKernelCache is true | [optional] 
**NotifyUser** | Pointer to **bool** |  | [optional] 
**NewPassword** | Pointer to **string** | The new password for Recovery Lock. Set as an empty string to clear the Recovery Lock password. | [optional] 

## Methods

### NewMdmCommandRequestCommandData

`func NewMdmCommandRequestCommandData() *MdmCommandRequestCommandData`

NewMdmCommandRequestCommandData instantiates a new MdmCommandRequestCommandData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMdmCommandRequestCommandDataWithDefaults

`func NewMdmCommandRequestCommandDataWithDefaults() *MdmCommandRequestCommandData`

NewMdmCommandRequestCommandDataWithDefaults instantiates a new MdmCommandRequestCommandData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommandType

`func (o *MdmCommandRequestCommandData) GetCommandType() MdmCommandType`

GetCommandType returns the CommandType field if non-nil, zero value otherwise.

### GetCommandTypeOk

`func (o *MdmCommandRequestCommandData) GetCommandTypeOk() (*MdmCommandType, bool)`

GetCommandTypeOk returns a tuple with the CommandType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandType

`func (o *MdmCommandRequestCommandData) SetCommandType(v MdmCommandType)`

SetCommandType sets CommandType field to given value.

### HasCommandType

`func (o *MdmCommandRequestCommandData) HasCommandType() bool`

HasCommandType returns a boolean if a field has been set.

### GetLostModeMessage

`func (o *MdmCommandRequestCommandData) GetLostModeMessage() string`

GetLostModeMessage returns the LostModeMessage field if non-nil, zero value otherwise.

### GetLostModeMessageOk

`func (o *MdmCommandRequestCommandData) GetLostModeMessageOk() (*string, bool)`

GetLostModeMessageOk returns a tuple with the LostModeMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLostModeMessage

`func (o *MdmCommandRequestCommandData) SetLostModeMessage(v string)`

SetLostModeMessage sets LostModeMessage field to given value.

### HasLostModeMessage

`func (o *MdmCommandRequestCommandData) HasLostModeMessage() bool`

HasLostModeMessage returns a boolean if a field has been set.

### GetLostModePhone

`func (o *MdmCommandRequestCommandData) GetLostModePhone() string`

GetLostModePhone returns the LostModePhone field if non-nil, zero value otherwise.

### GetLostModePhoneOk

`func (o *MdmCommandRequestCommandData) GetLostModePhoneOk() (*string, bool)`

GetLostModePhoneOk returns a tuple with the LostModePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLostModePhone

`func (o *MdmCommandRequestCommandData) SetLostModePhone(v string)`

SetLostModePhone sets LostModePhone field to given value.

### HasLostModePhone

`func (o *MdmCommandRequestCommandData) HasLostModePhone() bool`

HasLostModePhone returns a boolean if a field has been set.

### GetLostModeFootnote

`func (o *MdmCommandRequestCommandData) GetLostModeFootnote() string`

GetLostModeFootnote returns the LostModeFootnote field if non-nil, zero value otherwise.

### GetLostModeFootnoteOk

`func (o *MdmCommandRequestCommandData) GetLostModeFootnoteOk() (*string, bool)`

GetLostModeFootnoteOk returns a tuple with the LostModeFootnote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLostModeFootnote

`func (o *MdmCommandRequestCommandData) SetLostModeFootnote(v string)`

SetLostModeFootnote sets LostModeFootnote field to given value.

### HasLostModeFootnote

`func (o *MdmCommandRequestCommandData) HasLostModeFootnote() bool`

HasLostModeFootnote returns a boolean if a field has been set.

### GetUserName

`func (o *MdmCommandRequestCommandData) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *MdmCommandRequestCommandData) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *MdmCommandRequestCommandData) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *MdmCommandRequestCommandData) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetForceDeletion

`func (o *MdmCommandRequestCommandData) GetForceDeletion() bool`

GetForceDeletion returns the ForceDeletion field if non-nil, zero value otherwise.

### GetForceDeletionOk

`func (o *MdmCommandRequestCommandData) GetForceDeletionOk() (*bool, bool)`

GetForceDeletionOk returns a tuple with the ForceDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceDeletion

`func (o *MdmCommandRequestCommandData) SetForceDeletion(v bool)`

SetForceDeletion sets ForceDeletion field to given value.

### HasForceDeletion

`func (o *MdmCommandRequestCommandData) HasForceDeletion() bool`

HasForceDeletion returns a boolean if a field has been set.

### GetDeleteAllUsers

`func (o *MdmCommandRequestCommandData) GetDeleteAllUsers() bool`

GetDeleteAllUsers returns the DeleteAllUsers field if non-nil, zero value otherwise.

### GetDeleteAllUsersOk

`func (o *MdmCommandRequestCommandData) GetDeleteAllUsersOk() (*bool, bool)`

GetDeleteAllUsersOk returns a tuple with the DeleteAllUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteAllUsers

`func (o *MdmCommandRequestCommandData) SetDeleteAllUsers(v bool)`

SetDeleteAllUsers sets DeleteAllUsers field to given value.

### HasDeleteAllUsers

`func (o *MdmCommandRequestCommandData) HasDeleteAllUsers() bool`

HasDeleteAllUsers returns a boolean if a field has been set.

### GetRequestType

`func (o *MdmCommandRequestCommandData) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *MdmCommandRequestCommandData) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *MdmCommandRequestCommandData) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *MdmCommandRequestCommandData) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetRequestRequiresNetworkTether

`func (o *MdmCommandRequestCommandData) GetRequestRequiresNetworkTether() bool`

GetRequestRequiresNetworkTether returns the RequestRequiresNetworkTether field if non-nil, zero value otherwise.

### GetRequestRequiresNetworkTetherOk

`func (o *MdmCommandRequestCommandData) GetRequestRequiresNetworkTetherOk() (*bool, bool)`

GetRequestRequiresNetworkTetherOk returns a tuple with the RequestRequiresNetworkTether field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestRequiresNetworkTether

`func (o *MdmCommandRequestCommandData) SetRequestRequiresNetworkTether(v bool)`

SetRequestRequiresNetworkTether sets RequestRequiresNetworkTether field to given value.

### HasRequestRequiresNetworkTether

`func (o *MdmCommandRequestCommandData) HasRequestRequiresNetworkTether() bool`

HasRequestRequiresNetworkTether returns a boolean if a field has been set.

### GetBootstrapTokenAllowed

`func (o *MdmCommandRequestCommandData) GetBootstrapTokenAllowed() bool`

GetBootstrapTokenAllowed returns the BootstrapTokenAllowed field if non-nil, zero value otherwise.

### GetBootstrapTokenAllowedOk

`func (o *MdmCommandRequestCommandData) GetBootstrapTokenAllowedOk() (*bool, bool)`

GetBootstrapTokenAllowedOk returns a tuple with the BootstrapTokenAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootstrapTokenAllowed

`func (o *MdmCommandRequestCommandData) SetBootstrapTokenAllowed(v bool)`

SetBootstrapTokenAllowed sets BootstrapTokenAllowed field to given value.

### HasBootstrapTokenAllowed

`func (o *MdmCommandRequestCommandData) HasBootstrapTokenAllowed() bool`

HasBootstrapTokenAllowed returns a boolean if a field has been set.

### GetBluetooth

`func (o *MdmCommandRequestCommandData) GetBluetooth() bool`

GetBluetooth returns the Bluetooth field if non-nil, zero value otherwise.

### GetBluetoothOk

`func (o *MdmCommandRequestCommandData) GetBluetoothOk() (*bool, bool)`

GetBluetoothOk returns a tuple with the Bluetooth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBluetooth

`func (o *MdmCommandRequestCommandData) SetBluetooth(v bool)`

SetBluetooth sets Bluetooth field to given value.

### HasBluetooth

`func (o *MdmCommandRequestCommandData) HasBluetooth() bool`

HasBluetooth returns a boolean if a field has been set.

### GetAppAnalytics

`func (o *MdmCommandRequestCommandData) GetAppAnalytics() AppAnalyticsSetting`

GetAppAnalytics returns the AppAnalytics field if non-nil, zero value otherwise.

### GetAppAnalyticsOk

`func (o *MdmCommandRequestCommandData) GetAppAnalyticsOk() (*AppAnalyticsSetting, bool)`

GetAppAnalyticsOk returns a tuple with the AppAnalytics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppAnalytics

`func (o *MdmCommandRequestCommandData) SetAppAnalytics(v AppAnalyticsSetting)`

SetAppAnalytics sets AppAnalytics field to given value.

### HasAppAnalytics

`func (o *MdmCommandRequestCommandData) HasAppAnalytics() bool`

HasAppAnalytics returns a boolean if a field has been set.

### GetDiagnosticSubmission

`func (o *MdmCommandRequestCommandData) GetDiagnosticSubmission() DiagnosticSubmissionSetting`

GetDiagnosticSubmission returns the DiagnosticSubmission field if non-nil, zero value otherwise.

### GetDiagnosticSubmissionOk

`func (o *MdmCommandRequestCommandData) GetDiagnosticSubmissionOk() (*DiagnosticSubmissionSetting, bool)`

GetDiagnosticSubmissionOk returns a tuple with the DiagnosticSubmission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiagnosticSubmission

`func (o *MdmCommandRequestCommandData) SetDiagnosticSubmission(v DiagnosticSubmissionSetting)`

SetDiagnosticSubmission sets DiagnosticSubmission field to given value.

### HasDiagnosticSubmission

`func (o *MdmCommandRequestCommandData) HasDiagnosticSubmission() bool`

HasDiagnosticSubmission returns a boolean if a field has been set.

### GetDataRoaming

`func (o *MdmCommandRequestCommandData) GetDataRoaming() DataRoamingSetting`

GetDataRoaming returns the DataRoaming field if non-nil, zero value otherwise.

### GetDataRoamingOk

`func (o *MdmCommandRequestCommandData) GetDataRoamingOk() (*DataRoamingSetting, bool)`

GetDataRoamingOk returns a tuple with the DataRoaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRoaming

`func (o *MdmCommandRequestCommandData) SetDataRoaming(v DataRoamingSetting)`

SetDataRoaming sets DataRoaming field to given value.

### HasDataRoaming

`func (o *MdmCommandRequestCommandData) HasDataRoaming() bool`

HasDataRoaming returns a boolean if a field has been set.

### GetVoiceRoaming

`func (o *MdmCommandRequestCommandData) GetVoiceRoaming() VoiceRoamingSetting`

GetVoiceRoaming returns the VoiceRoaming field if non-nil, zero value otherwise.

### GetVoiceRoamingOk

`func (o *MdmCommandRequestCommandData) GetVoiceRoamingOk() (*VoiceRoamingSetting, bool)`

GetVoiceRoamingOk returns a tuple with the VoiceRoaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceRoaming

`func (o *MdmCommandRequestCommandData) SetVoiceRoaming(v VoiceRoamingSetting)`

SetVoiceRoaming sets VoiceRoaming field to given value.

### HasVoiceRoaming

`func (o *MdmCommandRequestCommandData) HasVoiceRoaming() bool`

HasVoiceRoaming returns a boolean if a field has been set.

### GetPersonalHotspot

`func (o *MdmCommandRequestCommandData) GetPersonalHotspot() PersonalHotspotSetting`

GetPersonalHotspot returns the PersonalHotspot field if non-nil, zero value otherwise.

### GetPersonalHotspotOk

`func (o *MdmCommandRequestCommandData) GetPersonalHotspotOk() (*PersonalHotspotSetting, bool)`

GetPersonalHotspotOk returns a tuple with the PersonalHotspot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalHotspot

`func (o *MdmCommandRequestCommandData) SetPersonalHotspot(v PersonalHotspotSetting)`

SetPersonalHotspot sets PersonalHotspot field to given value.

### HasPersonalHotspot

`func (o *MdmCommandRequestCommandData) HasPersonalHotspot() bool`

HasPersonalHotspot returns a boolean if a field has been set.

### GetMaximumResidentUsers

`func (o *MdmCommandRequestCommandData) GetMaximumResidentUsers() int32`

GetMaximumResidentUsers returns the MaximumResidentUsers field if non-nil, zero value otherwise.

### GetMaximumResidentUsersOk

`func (o *MdmCommandRequestCommandData) GetMaximumResidentUsersOk() (*int32, bool)`

GetMaximumResidentUsersOk returns a tuple with the MaximumResidentUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumResidentUsers

`func (o *MdmCommandRequestCommandData) SetMaximumResidentUsers(v int32)`

SetMaximumResidentUsers sets MaximumResidentUsers field to given value.

### HasMaximumResidentUsers

`func (o *MdmCommandRequestCommandData) HasMaximumResidentUsers() bool`

HasMaximumResidentUsers returns a boolean if a field has been set.

### GetDeviceName

`func (o *MdmCommandRequestCommandData) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *MdmCommandRequestCommandData) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *MdmCommandRequestCommandData) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *MdmCommandRequestCommandData) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetApplicationAttributes

`func (o *MdmCommandRequestCommandData) GetApplicationAttributes() ApplicationAttributes`

GetApplicationAttributes returns the ApplicationAttributes field if non-nil, zero value otherwise.

### GetApplicationAttributesOk

`func (o *MdmCommandRequestCommandData) GetApplicationAttributesOk() (*ApplicationAttributes, bool)`

GetApplicationAttributesOk returns a tuple with the ApplicationAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationAttributes

`func (o *MdmCommandRequestCommandData) SetApplicationAttributes(v ApplicationAttributes)`

SetApplicationAttributes sets ApplicationAttributes field to given value.

### HasApplicationAttributes

`func (o *MdmCommandRequestCommandData) HasApplicationAttributes() bool`

HasApplicationAttributes returns a boolean if a field has been set.

### GetSharedDeviceConfiguration

`func (o *MdmCommandRequestCommandData) GetSharedDeviceConfiguration() SharedDeviceConfiguration`

GetSharedDeviceConfiguration returns the SharedDeviceConfiguration field if non-nil, zero value otherwise.

### GetSharedDeviceConfigurationOk

`func (o *MdmCommandRequestCommandData) GetSharedDeviceConfigurationOk() (*SharedDeviceConfiguration, bool)`

GetSharedDeviceConfigurationOk returns a tuple with the SharedDeviceConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedDeviceConfiguration

`func (o *MdmCommandRequestCommandData) SetSharedDeviceConfiguration(v SharedDeviceConfiguration)`

SetSharedDeviceConfiguration sets SharedDeviceConfiguration field to given value.

### HasSharedDeviceConfiguration

`func (o *MdmCommandRequestCommandData) HasSharedDeviceConfiguration() bool`

HasSharedDeviceConfiguration returns a boolean if a field has been set.

### GetApplicationConfiguration

`func (o *MdmCommandRequestCommandData) GetApplicationConfiguration() ApplicationConfiguration`

GetApplicationConfiguration returns the ApplicationConfiguration field if non-nil, zero value otherwise.

### GetApplicationConfigurationOk

`func (o *MdmCommandRequestCommandData) GetApplicationConfigurationOk() (*ApplicationConfiguration, bool)`

GetApplicationConfigurationOk returns a tuple with the ApplicationConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationConfiguration

`func (o *MdmCommandRequestCommandData) SetApplicationConfiguration(v ApplicationConfiguration)`

SetApplicationConfiguration sets ApplicationConfiguration field to given value.

### HasApplicationConfiguration

`func (o *MdmCommandRequestCommandData) HasApplicationConfiguration() bool`

HasApplicationConfiguration returns a boolean if a field has been set.

### GetTimeZone

`func (o *MdmCommandRequestCommandData) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *MdmCommandRequestCommandData) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *MdmCommandRequestCommandData) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *MdmCommandRequestCommandData) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetSoftwareUpdateSettings

`func (o *MdmCommandRequestCommandData) GetSoftwareUpdateSettings() SoftwareUpdateSettings`

GetSoftwareUpdateSettings returns the SoftwareUpdateSettings field if non-nil, zero value otherwise.

### GetSoftwareUpdateSettingsOk

`func (o *MdmCommandRequestCommandData) GetSoftwareUpdateSettingsOk() (*SoftwareUpdateSettings, bool)`

GetSoftwareUpdateSettingsOk returns a tuple with the SoftwareUpdateSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareUpdateSettings

`func (o *MdmCommandRequestCommandData) SetSoftwareUpdateSettings(v SoftwareUpdateSettings)`

SetSoftwareUpdateSettings sets SoftwareUpdateSettings field to given value.

### HasSoftwareUpdateSettings

`func (o *MdmCommandRequestCommandData) HasSoftwareUpdateSettings() bool`

HasSoftwareUpdateSettings returns a boolean if a field has been set.

### GetRebuildKernelCache

`func (o *MdmCommandRequestCommandData) GetRebuildKernelCache() bool`

GetRebuildKernelCache returns the RebuildKernelCache field if non-nil, zero value otherwise.

### GetRebuildKernelCacheOk

`func (o *MdmCommandRequestCommandData) GetRebuildKernelCacheOk() (*bool, bool)`

GetRebuildKernelCacheOk returns a tuple with the RebuildKernelCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebuildKernelCache

`func (o *MdmCommandRequestCommandData) SetRebuildKernelCache(v bool)`

SetRebuildKernelCache sets RebuildKernelCache field to given value.

### HasRebuildKernelCache

`func (o *MdmCommandRequestCommandData) HasRebuildKernelCache() bool`

HasRebuildKernelCache returns a boolean if a field has been set.

### GetKextPaths

`func (o *MdmCommandRequestCommandData) GetKextPaths() []string`

GetKextPaths returns the KextPaths field if non-nil, zero value otherwise.

### GetKextPathsOk

`func (o *MdmCommandRequestCommandData) GetKextPathsOk() (*[]string, bool)`

GetKextPathsOk returns a tuple with the KextPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKextPaths

`func (o *MdmCommandRequestCommandData) SetKextPaths(v []string)`

SetKextPaths sets KextPaths field to given value.

### HasKextPaths

`func (o *MdmCommandRequestCommandData) HasKextPaths() bool`

HasKextPaths returns a boolean if a field has been set.

### GetNotifyUser

`func (o *MdmCommandRequestCommandData) GetNotifyUser() bool`

GetNotifyUser returns the NotifyUser field if non-nil, zero value otherwise.

### GetNotifyUserOk

`func (o *MdmCommandRequestCommandData) GetNotifyUserOk() (*bool, bool)`

GetNotifyUserOk returns a tuple with the NotifyUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyUser

`func (o *MdmCommandRequestCommandData) SetNotifyUser(v bool)`

SetNotifyUser sets NotifyUser field to given value.

### HasNotifyUser

`func (o *MdmCommandRequestCommandData) HasNotifyUser() bool`

HasNotifyUser returns a boolean if a field has been set.

### GetNewPassword

`func (o *MdmCommandRequestCommandData) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *MdmCommandRequestCommandData) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *MdmCommandRequestCommandData) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.

### HasNewPassword

`func (o *MdmCommandRequestCommandData) HasNewPassword() bool`

HasNewPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


