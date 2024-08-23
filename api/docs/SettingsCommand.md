# SettingsCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommandType** | [**MdmCommandType**](MdmCommandType.md) |  | 
**BootstrapTokenAllowed** | Pointer to **bool** |  | [optional] 
**Bluetooth** | Pointer to **bool** |  | [optional] 
**AppAnalytics** | Pointer to [**AppAnalyticsSetting**](AppAnalyticsSetting.md) |  | [optional] 
**DiagnosticSubmission** | Pointer to [**DiagnosticSubmissionSetting**](DiagnosticSubmissionSetting.md) |  | [optional] 
**DataRoaming** | Pointer to [**DataRoamingSetting**](DataRoamingSetting.md) |  | [optional] 
**VoiceRoaming** | Pointer to [**VoiceRoamingSetting**](VoiceRoamingSetting.md) |  | [optional] 
**PersonalHotspot** | Pointer to [**PersonalHotspotSetting**](PersonalHotspotSetting.md) |  | [optional] 
**MaximumResidentUsers** | Pointer to **int64** |  | [optional] 
**DeviceName** | Pointer to **string** |  | [optional] 
**ApplicationAttributes** | Pointer to [**ApplicationAttributes**](ApplicationAttributes.md) |  | [optional] 
**SharedDeviceConfiguration** | Pointer to [**SharedDeviceConfiguration**](SharedDeviceConfiguration.md) |  | [optional] 
**ApplicationConfiguration** | Pointer to [**ApplicationConfiguration**](ApplicationConfiguration.md) |  | [optional] 
**TimeZone** | Pointer to **string** |  | [optional] 
**SoftwareUpdateSettings** | Pointer to [**SoftwareUpdateSettings**](SoftwareUpdateSettings.md) |  | [optional] 
**PasscodeLockGracePeriod** | Pointer to **int64** | The number of seconds before a locked screen requires the user to enter the device passcode to unlock it. (Shared iPad Only) | [optional] 

## Methods

### NewSettingsCommand

`func NewSettingsCommand(commandType MdmCommandType, ) *SettingsCommand`

NewSettingsCommand instantiates a new SettingsCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingsCommandWithDefaults

`func NewSettingsCommandWithDefaults() *SettingsCommand`

NewSettingsCommandWithDefaults instantiates a new SettingsCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommandType

`func (o *SettingsCommand) GetCommandType() MdmCommandType`

GetCommandType returns the CommandType field if non-nil, zero value otherwise.

### GetCommandTypeOk

`func (o *SettingsCommand) GetCommandTypeOk() (*MdmCommandType, bool)`

GetCommandTypeOk returns a tuple with the CommandType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandType

`func (o *SettingsCommand) SetCommandType(v MdmCommandType)`

SetCommandType sets CommandType field to given value.


### GetBootstrapTokenAllowed

`func (o *SettingsCommand) GetBootstrapTokenAllowed() bool`

GetBootstrapTokenAllowed returns the BootstrapTokenAllowed field if non-nil, zero value otherwise.

### GetBootstrapTokenAllowedOk

`func (o *SettingsCommand) GetBootstrapTokenAllowedOk() (*bool, bool)`

GetBootstrapTokenAllowedOk returns a tuple with the BootstrapTokenAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootstrapTokenAllowed

`func (o *SettingsCommand) SetBootstrapTokenAllowed(v bool)`

SetBootstrapTokenAllowed sets BootstrapTokenAllowed field to given value.

### HasBootstrapTokenAllowed

`func (o *SettingsCommand) HasBootstrapTokenAllowed() bool`

HasBootstrapTokenAllowed returns a boolean if a field has been set.

### GetBluetooth

`func (o *SettingsCommand) GetBluetooth() bool`

GetBluetooth returns the Bluetooth field if non-nil, zero value otherwise.

### GetBluetoothOk

`func (o *SettingsCommand) GetBluetoothOk() (*bool, bool)`

GetBluetoothOk returns a tuple with the Bluetooth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBluetooth

`func (o *SettingsCommand) SetBluetooth(v bool)`

SetBluetooth sets Bluetooth field to given value.

### HasBluetooth

`func (o *SettingsCommand) HasBluetooth() bool`

HasBluetooth returns a boolean if a field has been set.

### GetAppAnalytics

`func (o *SettingsCommand) GetAppAnalytics() AppAnalyticsSetting`

GetAppAnalytics returns the AppAnalytics field if non-nil, zero value otherwise.

### GetAppAnalyticsOk

`func (o *SettingsCommand) GetAppAnalyticsOk() (*AppAnalyticsSetting, bool)`

GetAppAnalyticsOk returns a tuple with the AppAnalytics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppAnalytics

`func (o *SettingsCommand) SetAppAnalytics(v AppAnalyticsSetting)`

SetAppAnalytics sets AppAnalytics field to given value.

### HasAppAnalytics

`func (o *SettingsCommand) HasAppAnalytics() bool`

HasAppAnalytics returns a boolean if a field has been set.

### GetDiagnosticSubmission

`func (o *SettingsCommand) GetDiagnosticSubmission() DiagnosticSubmissionSetting`

GetDiagnosticSubmission returns the DiagnosticSubmission field if non-nil, zero value otherwise.

### GetDiagnosticSubmissionOk

`func (o *SettingsCommand) GetDiagnosticSubmissionOk() (*DiagnosticSubmissionSetting, bool)`

GetDiagnosticSubmissionOk returns a tuple with the DiagnosticSubmission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiagnosticSubmission

`func (o *SettingsCommand) SetDiagnosticSubmission(v DiagnosticSubmissionSetting)`

SetDiagnosticSubmission sets DiagnosticSubmission field to given value.

### HasDiagnosticSubmission

`func (o *SettingsCommand) HasDiagnosticSubmission() bool`

HasDiagnosticSubmission returns a boolean if a field has been set.

### GetDataRoaming

`func (o *SettingsCommand) GetDataRoaming() DataRoamingSetting`

GetDataRoaming returns the DataRoaming field if non-nil, zero value otherwise.

### GetDataRoamingOk

`func (o *SettingsCommand) GetDataRoamingOk() (*DataRoamingSetting, bool)`

GetDataRoamingOk returns a tuple with the DataRoaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRoaming

`func (o *SettingsCommand) SetDataRoaming(v DataRoamingSetting)`

SetDataRoaming sets DataRoaming field to given value.

### HasDataRoaming

`func (o *SettingsCommand) HasDataRoaming() bool`

HasDataRoaming returns a boolean if a field has been set.

### GetVoiceRoaming

`func (o *SettingsCommand) GetVoiceRoaming() VoiceRoamingSetting`

GetVoiceRoaming returns the VoiceRoaming field if non-nil, zero value otherwise.

### GetVoiceRoamingOk

`func (o *SettingsCommand) GetVoiceRoamingOk() (*VoiceRoamingSetting, bool)`

GetVoiceRoamingOk returns a tuple with the VoiceRoaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceRoaming

`func (o *SettingsCommand) SetVoiceRoaming(v VoiceRoamingSetting)`

SetVoiceRoaming sets VoiceRoaming field to given value.

### HasVoiceRoaming

`func (o *SettingsCommand) HasVoiceRoaming() bool`

HasVoiceRoaming returns a boolean if a field has been set.

### GetPersonalHotspot

`func (o *SettingsCommand) GetPersonalHotspot() PersonalHotspotSetting`

GetPersonalHotspot returns the PersonalHotspot field if non-nil, zero value otherwise.

### GetPersonalHotspotOk

`func (o *SettingsCommand) GetPersonalHotspotOk() (*PersonalHotspotSetting, bool)`

GetPersonalHotspotOk returns a tuple with the PersonalHotspot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalHotspot

`func (o *SettingsCommand) SetPersonalHotspot(v PersonalHotspotSetting)`

SetPersonalHotspot sets PersonalHotspot field to given value.

### HasPersonalHotspot

`func (o *SettingsCommand) HasPersonalHotspot() bool`

HasPersonalHotspot returns a boolean if a field has been set.

### GetMaximumResidentUsers

`func (o *SettingsCommand) GetMaximumResidentUsers() int64`

GetMaximumResidentUsers returns the MaximumResidentUsers field if non-nil, zero value otherwise.

### GetMaximumResidentUsersOk

`func (o *SettingsCommand) GetMaximumResidentUsersOk() (*int64, bool)`

GetMaximumResidentUsersOk returns a tuple with the MaximumResidentUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumResidentUsers

`func (o *SettingsCommand) SetMaximumResidentUsers(v int64)`

SetMaximumResidentUsers sets MaximumResidentUsers field to given value.

### HasMaximumResidentUsers

`func (o *SettingsCommand) HasMaximumResidentUsers() bool`

HasMaximumResidentUsers returns a boolean if a field has been set.

### GetDeviceName

`func (o *SettingsCommand) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *SettingsCommand) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *SettingsCommand) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *SettingsCommand) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetApplicationAttributes

`func (o *SettingsCommand) GetApplicationAttributes() ApplicationAttributes`

GetApplicationAttributes returns the ApplicationAttributes field if non-nil, zero value otherwise.

### GetApplicationAttributesOk

`func (o *SettingsCommand) GetApplicationAttributesOk() (*ApplicationAttributes, bool)`

GetApplicationAttributesOk returns a tuple with the ApplicationAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationAttributes

`func (o *SettingsCommand) SetApplicationAttributes(v ApplicationAttributes)`

SetApplicationAttributes sets ApplicationAttributes field to given value.

### HasApplicationAttributes

`func (o *SettingsCommand) HasApplicationAttributes() bool`

HasApplicationAttributes returns a boolean if a field has been set.

### GetSharedDeviceConfiguration

`func (o *SettingsCommand) GetSharedDeviceConfiguration() SharedDeviceConfiguration`

GetSharedDeviceConfiguration returns the SharedDeviceConfiguration field if non-nil, zero value otherwise.

### GetSharedDeviceConfigurationOk

`func (o *SettingsCommand) GetSharedDeviceConfigurationOk() (*SharedDeviceConfiguration, bool)`

GetSharedDeviceConfigurationOk returns a tuple with the SharedDeviceConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedDeviceConfiguration

`func (o *SettingsCommand) SetSharedDeviceConfiguration(v SharedDeviceConfiguration)`

SetSharedDeviceConfiguration sets SharedDeviceConfiguration field to given value.

### HasSharedDeviceConfiguration

`func (o *SettingsCommand) HasSharedDeviceConfiguration() bool`

HasSharedDeviceConfiguration returns a boolean if a field has been set.

### GetApplicationConfiguration

`func (o *SettingsCommand) GetApplicationConfiguration() ApplicationConfiguration`

GetApplicationConfiguration returns the ApplicationConfiguration field if non-nil, zero value otherwise.

### GetApplicationConfigurationOk

`func (o *SettingsCommand) GetApplicationConfigurationOk() (*ApplicationConfiguration, bool)`

GetApplicationConfigurationOk returns a tuple with the ApplicationConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationConfiguration

`func (o *SettingsCommand) SetApplicationConfiguration(v ApplicationConfiguration)`

SetApplicationConfiguration sets ApplicationConfiguration field to given value.

### HasApplicationConfiguration

`func (o *SettingsCommand) HasApplicationConfiguration() bool`

HasApplicationConfiguration returns a boolean if a field has been set.

### GetTimeZone

`func (o *SettingsCommand) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *SettingsCommand) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *SettingsCommand) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *SettingsCommand) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetSoftwareUpdateSettings

`func (o *SettingsCommand) GetSoftwareUpdateSettings() SoftwareUpdateSettings`

GetSoftwareUpdateSettings returns the SoftwareUpdateSettings field if non-nil, zero value otherwise.

### GetSoftwareUpdateSettingsOk

`func (o *SettingsCommand) GetSoftwareUpdateSettingsOk() (*SoftwareUpdateSettings, bool)`

GetSoftwareUpdateSettingsOk returns a tuple with the SoftwareUpdateSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareUpdateSettings

`func (o *SettingsCommand) SetSoftwareUpdateSettings(v SoftwareUpdateSettings)`

SetSoftwareUpdateSettings sets SoftwareUpdateSettings field to given value.

### HasSoftwareUpdateSettings

`func (o *SettingsCommand) HasSoftwareUpdateSettings() bool`

HasSoftwareUpdateSettings returns a boolean if a field has been set.

### GetPasscodeLockGracePeriod

`func (o *SettingsCommand) GetPasscodeLockGracePeriod() int64`

GetPasscodeLockGracePeriod returns the PasscodeLockGracePeriod field if non-nil, zero value otherwise.

### GetPasscodeLockGracePeriodOk

`func (o *SettingsCommand) GetPasscodeLockGracePeriodOk() (*int64, bool)`

GetPasscodeLockGracePeriodOk returns a tuple with the PasscodeLockGracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasscodeLockGracePeriod

`func (o *SettingsCommand) SetPasscodeLockGracePeriod(v int64)`

SetPasscodeLockGracePeriod sets PasscodeLockGracePeriod field to given value.

### HasPasscodeLockGracePeriod

`func (o *SettingsCommand) HasPasscodeLockGracePeriod() bool`

HasPasscodeLockGracePeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


