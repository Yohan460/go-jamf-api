# SsoSettingsV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigurationType** | **string** |  | 
**OidcSettings** | [**OidcSettings**](OidcSettings.md) |  | 
**SamlSettings** | [**SamlSettings**](SamlSettings.md) |  | 
**SsoForEnrollmentEnabled** | **bool** |  | [default to false]
**SsoBypassAllowed** | **bool** |  | [default to false]
**SsoEnabled** | **bool** |  | [default to false]
**SsoForMacOsSelfServiceEnabled** | **bool** |  | [default to false]
**EnrollmentSsoForAccountDrivenEnrollmentEnabled** | **bool** |  | [default to false]
**EnrollmentSsoConfig** | Pointer to [**EnrollmentSsoConfig**](EnrollmentSsoConfig.md) |  | [optional] 
**GroupEnrollmentAccessEnabled** | **bool** |  | [default to false]
**GroupEnrollmentAccessName** | Pointer to **string** |  | [optional] [default to " "]

## Methods

### NewSsoSettingsV3

`func NewSsoSettingsV3(configurationType string, oidcSettings OidcSettings, samlSettings SamlSettings, ssoForEnrollmentEnabled bool, ssoBypassAllowed bool, ssoEnabled bool, ssoForMacOsSelfServiceEnabled bool, enrollmentSsoForAccountDrivenEnrollmentEnabled bool, groupEnrollmentAccessEnabled bool, ) *SsoSettingsV3`

NewSsoSettingsV3 instantiates a new SsoSettingsV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoSettingsV3WithDefaults

`func NewSsoSettingsV3WithDefaults() *SsoSettingsV3`

NewSsoSettingsV3WithDefaults instantiates a new SsoSettingsV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigurationType

`func (o *SsoSettingsV3) GetConfigurationType() string`

GetConfigurationType returns the ConfigurationType field if non-nil, zero value otherwise.

### GetConfigurationTypeOk

`func (o *SsoSettingsV3) GetConfigurationTypeOk() (*string, bool)`

GetConfigurationTypeOk returns a tuple with the ConfigurationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationType

`func (o *SsoSettingsV3) SetConfigurationType(v string)`

SetConfigurationType sets ConfigurationType field to given value.


### GetOidcSettings

`func (o *SsoSettingsV3) GetOidcSettings() OidcSettings`

GetOidcSettings returns the OidcSettings field if non-nil, zero value otherwise.

### GetOidcSettingsOk

`func (o *SsoSettingsV3) GetOidcSettingsOk() (*OidcSettings, bool)`

GetOidcSettingsOk returns a tuple with the OidcSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcSettings

`func (o *SsoSettingsV3) SetOidcSettings(v OidcSettings)`

SetOidcSettings sets OidcSettings field to given value.


### GetSamlSettings

`func (o *SsoSettingsV3) GetSamlSettings() SamlSettings`

GetSamlSettings returns the SamlSettings field if non-nil, zero value otherwise.

### GetSamlSettingsOk

`func (o *SsoSettingsV3) GetSamlSettingsOk() (*SamlSettings, bool)`

GetSamlSettingsOk returns a tuple with the SamlSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlSettings

`func (o *SsoSettingsV3) SetSamlSettings(v SamlSettings)`

SetSamlSettings sets SamlSettings field to given value.


### GetSsoForEnrollmentEnabled

`func (o *SsoSettingsV3) GetSsoForEnrollmentEnabled() bool`

GetSsoForEnrollmentEnabled returns the SsoForEnrollmentEnabled field if non-nil, zero value otherwise.

### GetSsoForEnrollmentEnabledOk

`func (o *SsoSettingsV3) GetSsoForEnrollmentEnabledOk() (*bool, bool)`

GetSsoForEnrollmentEnabledOk returns a tuple with the SsoForEnrollmentEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoForEnrollmentEnabled

`func (o *SsoSettingsV3) SetSsoForEnrollmentEnabled(v bool)`

SetSsoForEnrollmentEnabled sets SsoForEnrollmentEnabled field to given value.


### GetSsoBypassAllowed

`func (o *SsoSettingsV3) GetSsoBypassAllowed() bool`

GetSsoBypassAllowed returns the SsoBypassAllowed field if non-nil, zero value otherwise.

### GetSsoBypassAllowedOk

`func (o *SsoSettingsV3) GetSsoBypassAllowedOk() (*bool, bool)`

GetSsoBypassAllowedOk returns a tuple with the SsoBypassAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoBypassAllowed

`func (o *SsoSettingsV3) SetSsoBypassAllowed(v bool)`

SetSsoBypassAllowed sets SsoBypassAllowed field to given value.


### GetSsoEnabled

`func (o *SsoSettingsV3) GetSsoEnabled() bool`

GetSsoEnabled returns the SsoEnabled field if non-nil, zero value otherwise.

### GetSsoEnabledOk

`func (o *SsoSettingsV3) GetSsoEnabledOk() (*bool, bool)`

GetSsoEnabledOk returns a tuple with the SsoEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoEnabled

`func (o *SsoSettingsV3) SetSsoEnabled(v bool)`

SetSsoEnabled sets SsoEnabled field to given value.


### GetSsoForMacOsSelfServiceEnabled

`func (o *SsoSettingsV3) GetSsoForMacOsSelfServiceEnabled() bool`

GetSsoForMacOsSelfServiceEnabled returns the SsoForMacOsSelfServiceEnabled field if non-nil, zero value otherwise.

### GetSsoForMacOsSelfServiceEnabledOk

`func (o *SsoSettingsV3) GetSsoForMacOsSelfServiceEnabledOk() (*bool, bool)`

GetSsoForMacOsSelfServiceEnabledOk returns a tuple with the SsoForMacOsSelfServiceEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoForMacOsSelfServiceEnabled

`func (o *SsoSettingsV3) SetSsoForMacOsSelfServiceEnabled(v bool)`

SetSsoForMacOsSelfServiceEnabled sets SsoForMacOsSelfServiceEnabled field to given value.


### GetEnrollmentSsoForAccountDrivenEnrollmentEnabled

`func (o *SsoSettingsV3) GetEnrollmentSsoForAccountDrivenEnrollmentEnabled() bool`

GetEnrollmentSsoForAccountDrivenEnrollmentEnabled returns the EnrollmentSsoForAccountDrivenEnrollmentEnabled field if non-nil, zero value otherwise.

### GetEnrollmentSsoForAccountDrivenEnrollmentEnabledOk

`func (o *SsoSettingsV3) GetEnrollmentSsoForAccountDrivenEnrollmentEnabledOk() (*bool, bool)`

GetEnrollmentSsoForAccountDrivenEnrollmentEnabledOk returns a tuple with the EnrollmentSsoForAccountDrivenEnrollmentEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentSsoForAccountDrivenEnrollmentEnabled

`func (o *SsoSettingsV3) SetEnrollmentSsoForAccountDrivenEnrollmentEnabled(v bool)`

SetEnrollmentSsoForAccountDrivenEnrollmentEnabled sets EnrollmentSsoForAccountDrivenEnrollmentEnabled field to given value.


### GetEnrollmentSsoConfig

`func (o *SsoSettingsV3) GetEnrollmentSsoConfig() EnrollmentSsoConfig`

GetEnrollmentSsoConfig returns the EnrollmentSsoConfig field if non-nil, zero value otherwise.

### GetEnrollmentSsoConfigOk

`func (o *SsoSettingsV3) GetEnrollmentSsoConfigOk() (*EnrollmentSsoConfig, bool)`

GetEnrollmentSsoConfigOk returns a tuple with the EnrollmentSsoConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentSsoConfig

`func (o *SsoSettingsV3) SetEnrollmentSsoConfig(v EnrollmentSsoConfig)`

SetEnrollmentSsoConfig sets EnrollmentSsoConfig field to given value.

### HasEnrollmentSsoConfig

`func (o *SsoSettingsV3) HasEnrollmentSsoConfig() bool`

HasEnrollmentSsoConfig returns a boolean if a field has been set.

### GetGroupEnrollmentAccessEnabled

`func (o *SsoSettingsV3) GetGroupEnrollmentAccessEnabled() bool`

GetGroupEnrollmentAccessEnabled returns the GroupEnrollmentAccessEnabled field if non-nil, zero value otherwise.

### GetGroupEnrollmentAccessEnabledOk

`func (o *SsoSettingsV3) GetGroupEnrollmentAccessEnabledOk() (*bool, bool)`

GetGroupEnrollmentAccessEnabledOk returns a tuple with the GroupEnrollmentAccessEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupEnrollmentAccessEnabled

`func (o *SsoSettingsV3) SetGroupEnrollmentAccessEnabled(v bool)`

SetGroupEnrollmentAccessEnabled sets GroupEnrollmentAccessEnabled field to given value.


### GetGroupEnrollmentAccessName

`func (o *SsoSettingsV3) GetGroupEnrollmentAccessName() string`

GetGroupEnrollmentAccessName returns the GroupEnrollmentAccessName field if non-nil, zero value otherwise.

### GetGroupEnrollmentAccessNameOk

`func (o *SsoSettingsV3) GetGroupEnrollmentAccessNameOk() (*string, bool)`

GetGroupEnrollmentAccessNameOk returns a tuple with the GroupEnrollmentAccessName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupEnrollmentAccessName

`func (o *SsoSettingsV3) SetGroupEnrollmentAccessName(v string)`

SetGroupEnrollmentAccessName sets GroupEnrollmentAccessName field to given value.

### HasGroupEnrollmentAccessName

`func (o *SsoSettingsV3) HasGroupEnrollmentAccessName() bool`

HasGroupEnrollmentAccessName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


