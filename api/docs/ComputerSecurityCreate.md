# ComputerSecurityCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SipStatus** | Pointer to **NullableString** |  | [optional] 
**GatekeeperStatus** | Pointer to **NullableString** |  | [optional] 
**XprotectVersion** | Pointer to **NullableString** |  | [optional] 
**ActivationLockEnabled** | Pointer to **NullableBool** |  | [optional] 
**RecoveryLockEnabled** | Pointer to **NullableBool** |  | [optional] 
**FirewallEnabled** | Pointer to **NullableBool** |  | [optional] 
**SecureBootLevel** | Pointer to **NullableString** |  | [optional] 
**ExternalBootLevel** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewComputerSecurityCreate

`func NewComputerSecurityCreate() *ComputerSecurityCreate`

NewComputerSecurityCreate instantiates a new ComputerSecurityCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputerSecurityCreateWithDefaults

`func NewComputerSecurityCreateWithDefaults() *ComputerSecurityCreate`

NewComputerSecurityCreateWithDefaults instantiates a new ComputerSecurityCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSipStatus

`func (o *ComputerSecurityCreate) GetSipStatus() string`

GetSipStatus returns the SipStatus field if non-nil, zero value otherwise.

### GetSipStatusOk

`func (o *ComputerSecurityCreate) GetSipStatusOk() (*string, bool)`

GetSipStatusOk returns a tuple with the SipStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipStatus

`func (o *ComputerSecurityCreate) SetSipStatus(v string)`

SetSipStatus sets SipStatus field to given value.

### HasSipStatus

`func (o *ComputerSecurityCreate) HasSipStatus() bool`

HasSipStatus returns a boolean if a field has been set.

### SetSipStatusNil

`func (o *ComputerSecurityCreate) SetSipStatusNil(b bool)`

 SetSipStatusNil sets the value for SipStatus to be an explicit nil

### UnsetSipStatus
`func (o *ComputerSecurityCreate) UnsetSipStatus()`

UnsetSipStatus ensures that no value is present for SipStatus, not even an explicit nil
### GetGatekeeperStatus

`func (o *ComputerSecurityCreate) GetGatekeeperStatus() string`

GetGatekeeperStatus returns the GatekeeperStatus field if non-nil, zero value otherwise.

### GetGatekeeperStatusOk

`func (o *ComputerSecurityCreate) GetGatekeeperStatusOk() (*string, bool)`

GetGatekeeperStatusOk returns a tuple with the GatekeeperStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatekeeperStatus

`func (o *ComputerSecurityCreate) SetGatekeeperStatus(v string)`

SetGatekeeperStatus sets GatekeeperStatus field to given value.

### HasGatekeeperStatus

`func (o *ComputerSecurityCreate) HasGatekeeperStatus() bool`

HasGatekeeperStatus returns a boolean if a field has been set.

### SetGatekeeperStatusNil

`func (o *ComputerSecurityCreate) SetGatekeeperStatusNil(b bool)`

 SetGatekeeperStatusNil sets the value for GatekeeperStatus to be an explicit nil

### UnsetGatekeeperStatus
`func (o *ComputerSecurityCreate) UnsetGatekeeperStatus()`

UnsetGatekeeperStatus ensures that no value is present for GatekeeperStatus, not even an explicit nil
### GetXprotectVersion

`func (o *ComputerSecurityCreate) GetXprotectVersion() string`

GetXprotectVersion returns the XprotectVersion field if non-nil, zero value otherwise.

### GetXprotectVersionOk

`func (o *ComputerSecurityCreate) GetXprotectVersionOk() (*string, bool)`

GetXprotectVersionOk returns a tuple with the XprotectVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXprotectVersion

`func (o *ComputerSecurityCreate) SetXprotectVersion(v string)`

SetXprotectVersion sets XprotectVersion field to given value.

### HasXprotectVersion

`func (o *ComputerSecurityCreate) HasXprotectVersion() bool`

HasXprotectVersion returns a boolean if a field has been set.

### SetXprotectVersionNil

`func (o *ComputerSecurityCreate) SetXprotectVersionNil(b bool)`

 SetXprotectVersionNil sets the value for XprotectVersion to be an explicit nil

### UnsetXprotectVersion
`func (o *ComputerSecurityCreate) UnsetXprotectVersion()`

UnsetXprotectVersion ensures that no value is present for XprotectVersion, not even an explicit nil
### GetActivationLockEnabled

`func (o *ComputerSecurityCreate) GetActivationLockEnabled() bool`

GetActivationLockEnabled returns the ActivationLockEnabled field if non-nil, zero value otherwise.

### GetActivationLockEnabledOk

`func (o *ComputerSecurityCreate) GetActivationLockEnabledOk() (*bool, bool)`

GetActivationLockEnabledOk returns a tuple with the ActivationLockEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationLockEnabled

`func (o *ComputerSecurityCreate) SetActivationLockEnabled(v bool)`

SetActivationLockEnabled sets ActivationLockEnabled field to given value.

### HasActivationLockEnabled

`func (o *ComputerSecurityCreate) HasActivationLockEnabled() bool`

HasActivationLockEnabled returns a boolean if a field has been set.

### SetActivationLockEnabledNil

`func (o *ComputerSecurityCreate) SetActivationLockEnabledNil(b bool)`

 SetActivationLockEnabledNil sets the value for ActivationLockEnabled to be an explicit nil

### UnsetActivationLockEnabled
`func (o *ComputerSecurityCreate) UnsetActivationLockEnabled()`

UnsetActivationLockEnabled ensures that no value is present for ActivationLockEnabled, not even an explicit nil
### GetRecoveryLockEnabled

`func (o *ComputerSecurityCreate) GetRecoveryLockEnabled() bool`

GetRecoveryLockEnabled returns the RecoveryLockEnabled field if non-nil, zero value otherwise.

### GetRecoveryLockEnabledOk

`func (o *ComputerSecurityCreate) GetRecoveryLockEnabledOk() (*bool, bool)`

GetRecoveryLockEnabledOk returns a tuple with the RecoveryLockEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryLockEnabled

`func (o *ComputerSecurityCreate) SetRecoveryLockEnabled(v bool)`

SetRecoveryLockEnabled sets RecoveryLockEnabled field to given value.

### HasRecoveryLockEnabled

`func (o *ComputerSecurityCreate) HasRecoveryLockEnabled() bool`

HasRecoveryLockEnabled returns a boolean if a field has been set.

### SetRecoveryLockEnabledNil

`func (o *ComputerSecurityCreate) SetRecoveryLockEnabledNil(b bool)`

 SetRecoveryLockEnabledNil sets the value for RecoveryLockEnabled to be an explicit nil

### UnsetRecoveryLockEnabled
`func (o *ComputerSecurityCreate) UnsetRecoveryLockEnabled()`

UnsetRecoveryLockEnabled ensures that no value is present for RecoveryLockEnabled, not even an explicit nil
### GetFirewallEnabled

`func (o *ComputerSecurityCreate) GetFirewallEnabled() bool`

GetFirewallEnabled returns the FirewallEnabled field if non-nil, zero value otherwise.

### GetFirewallEnabledOk

`func (o *ComputerSecurityCreate) GetFirewallEnabledOk() (*bool, bool)`

GetFirewallEnabledOk returns a tuple with the FirewallEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallEnabled

`func (o *ComputerSecurityCreate) SetFirewallEnabled(v bool)`

SetFirewallEnabled sets FirewallEnabled field to given value.

### HasFirewallEnabled

`func (o *ComputerSecurityCreate) HasFirewallEnabled() bool`

HasFirewallEnabled returns a boolean if a field has been set.

### SetFirewallEnabledNil

`func (o *ComputerSecurityCreate) SetFirewallEnabledNil(b bool)`

 SetFirewallEnabledNil sets the value for FirewallEnabled to be an explicit nil

### UnsetFirewallEnabled
`func (o *ComputerSecurityCreate) UnsetFirewallEnabled()`

UnsetFirewallEnabled ensures that no value is present for FirewallEnabled, not even an explicit nil
### GetSecureBootLevel

`func (o *ComputerSecurityCreate) GetSecureBootLevel() string`

GetSecureBootLevel returns the SecureBootLevel field if non-nil, zero value otherwise.

### GetSecureBootLevelOk

`func (o *ComputerSecurityCreate) GetSecureBootLevelOk() (*string, bool)`

GetSecureBootLevelOk returns a tuple with the SecureBootLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureBootLevel

`func (o *ComputerSecurityCreate) SetSecureBootLevel(v string)`

SetSecureBootLevel sets SecureBootLevel field to given value.

### HasSecureBootLevel

`func (o *ComputerSecurityCreate) HasSecureBootLevel() bool`

HasSecureBootLevel returns a boolean if a field has been set.

### SetSecureBootLevelNil

`func (o *ComputerSecurityCreate) SetSecureBootLevelNil(b bool)`

 SetSecureBootLevelNil sets the value for SecureBootLevel to be an explicit nil

### UnsetSecureBootLevel
`func (o *ComputerSecurityCreate) UnsetSecureBootLevel()`

UnsetSecureBootLevel ensures that no value is present for SecureBootLevel, not even an explicit nil
### GetExternalBootLevel

`func (o *ComputerSecurityCreate) GetExternalBootLevel() string`

GetExternalBootLevel returns the ExternalBootLevel field if non-nil, zero value otherwise.

### GetExternalBootLevelOk

`func (o *ComputerSecurityCreate) GetExternalBootLevelOk() (*string, bool)`

GetExternalBootLevelOk returns a tuple with the ExternalBootLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalBootLevel

`func (o *ComputerSecurityCreate) SetExternalBootLevel(v string)`

SetExternalBootLevel sets ExternalBootLevel field to given value.

### HasExternalBootLevel

`func (o *ComputerSecurityCreate) HasExternalBootLevel() bool`

HasExternalBootLevel returns a boolean if a field has been set.

### SetExternalBootLevelNil

`func (o *ComputerSecurityCreate) SetExternalBootLevelNil(b bool)`

 SetExternalBootLevelNil sets the value for ExternalBootLevel to be an explicit nil

### UnsetExternalBootLevel
`func (o *ComputerSecurityCreate) UnsetExternalBootLevel()`

UnsetExternalBootLevel ensures that no value is present for ExternalBootLevel, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


