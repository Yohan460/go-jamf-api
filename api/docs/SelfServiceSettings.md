# SelfServiceSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstallSettings** | Pointer to [**SelfServiceInstallSettings**](SelfServiceInstallSettings.md) |  | [optional] 
**LoginSettings** | Pointer to [**SelfServiceLoginSettings**](SelfServiceLoginSettings.md) |  | [optional] 
**ConfigurationSettings** | Pointer to [**SelfServiceInteractionSettings**](SelfServiceInteractionSettings.md) |  | [optional] 

## Methods

### NewSelfServiceSettings

`func NewSelfServiceSettings() *SelfServiceSettings`

NewSelfServiceSettings instantiates a new SelfServiceSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelfServiceSettingsWithDefaults

`func NewSelfServiceSettingsWithDefaults() *SelfServiceSettings`

NewSelfServiceSettingsWithDefaults instantiates a new SelfServiceSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstallSettings

`func (o *SelfServiceSettings) GetInstallSettings() SelfServiceInstallSettings`

GetInstallSettings returns the InstallSettings field if non-nil, zero value otherwise.

### GetInstallSettingsOk

`func (o *SelfServiceSettings) GetInstallSettingsOk() (*SelfServiceInstallSettings, bool)`

GetInstallSettingsOk returns a tuple with the InstallSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallSettings

`func (o *SelfServiceSettings) SetInstallSettings(v SelfServiceInstallSettings)`

SetInstallSettings sets InstallSettings field to given value.

### HasInstallSettings

`func (o *SelfServiceSettings) HasInstallSettings() bool`

HasInstallSettings returns a boolean if a field has been set.

### GetLoginSettings

`func (o *SelfServiceSettings) GetLoginSettings() SelfServiceLoginSettings`

GetLoginSettings returns the LoginSettings field if non-nil, zero value otherwise.

### GetLoginSettingsOk

`func (o *SelfServiceSettings) GetLoginSettingsOk() (*SelfServiceLoginSettings, bool)`

GetLoginSettingsOk returns a tuple with the LoginSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginSettings

`func (o *SelfServiceSettings) SetLoginSettings(v SelfServiceLoginSettings)`

SetLoginSettings sets LoginSettings field to given value.

### HasLoginSettings

`func (o *SelfServiceSettings) HasLoginSettings() bool`

HasLoginSettings returns a boolean if a field has been set.

### GetConfigurationSettings

`func (o *SelfServiceSettings) GetConfigurationSettings() SelfServiceInteractionSettings`

GetConfigurationSettings returns the ConfigurationSettings field if non-nil, zero value otherwise.

### GetConfigurationSettingsOk

`func (o *SelfServiceSettings) GetConfigurationSettingsOk() (*SelfServiceInteractionSettings, bool)`

GetConfigurationSettingsOk returns a tuple with the ConfigurationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationSettings

`func (o *SelfServiceSettings) SetConfigurationSettings(v SelfServiceInteractionSettings)`

SetConfigurationSettings sets ConfigurationSettings field to given value.

### HasConfigurationSettings

`func (o *SelfServiceSettings) HasConfigurationSettings() bool`

HasConfigurationSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


