# ComputerSecurity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SipStatus** | Pointer to **string** |  | [optional] 
**GatekeeperStatus** | Pointer to **string** |  | [optional] 
**XprotectVersion** | Pointer to **string** |  | [optional] 
**AutoLoginDisabled** | Pointer to **bool** |  | [optional] 
**RemoteDesktopEnabled** | Pointer to **bool** | Collected for macOS 10.14.4 or later | [optional] 
**ActivationLockEnabled** | Pointer to **bool** | Collected for macOS 10.15.0 or later | [optional] 
**RecoveryLockEnabled** | Pointer to **bool** |  | [optional] 
**FirewallEnabled** | Pointer to **bool** |  | [optional] 
**SecureBootLevel** | Pointer to **string** | Collected for macOS 10.15.0 or later | [optional] 
**ExternalBootLevel** | Pointer to **string** | Collected for macOS 10.15.0 or later | [optional] 
**BootstrapTokenAllowed** | Pointer to **bool** | Collected for macOS 11 or later | [optional] 
**BootstrapTokenEscrowedStatus** | Pointer to **string** | Collected for macOS 11 or later | [optional] 

## Methods

### NewComputerSecurity

`func NewComputerSecurity() *ComputerSecurity`

NewComputerSecurity instantiates a new ComputerSecurity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputerSecurityWithDefaults

`func NewComputerSecurityWithDefaults() *ComputerSecurity`

NewComputerSecurityWithDefaults instantiates a new ComputerSecurity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSipStatus

`func (o *ComputerSecurity) GetSipStatus() string`

GetSipStatus returns the SipStatus field if non-nil, zero value otherwise.

### GetSipStatusOk

`func (o *ComputerSecurity) GetSipStatusOk() (*string, bool)`

GetSipStatusOk returns a tuple with the SipStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipStatus

`func (o *ComputerSecurity) SetSipStatus(v string)`

SetSipStatus sets SipStatus field to given value.

### HasSipStatus

`func (o *ComputerSecurity) HasSipStatus() bool`

HasSipStatus returns a boolean if a field has been set.

### GetGatekeeperStatus

`func (o *ComputerSecurity) GetGatekeeperStatus() string`

GetGatekeeperStatus returns the GatekeeperStatus field if non-nil, zero value otherwise.

### GetGatekeeperStatusOk

`func (o *ComputerSecurity) GetGatekeeperStatusOk() (*string, bool)`

GetGatekeeperStatusOk returns a tuple with the GatekeeperStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatekeeperStatus

`func (o *ComputerSecurity) SetGatekeeperStatus(v string)`

SetGatekeeperStatus sets GatekeeperStatus field to given value.

### HasGatekeeperStatus

`func (o *ComputerSecurity) HasGatekeeperStatus() bool`

HasGatekeeperStatus returns a boolean if a field has been set.

### GetXprotectVersion

`func (o *ComputerSecurity) GetXprotectVersion() string`

GetXprotectVersion returns the XprotectVersion field if non-nil, zero value otherwise.

### GetXprotectVersionOk

`func (o *ComputerSecurity) GetXprotectVersionOk() (*string, bool)`

GetXprotectVersionOk returns a tuple with the XprotectVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXprotectVersion

`func (o *ComputerSecurity) SetXprotectVersion(v string)`

SetXprotectVersion sets XprotectVersion field to given value.

### HasXprotectVersion

`func (o *ComputerSecurity) HasXprotectVersion() bool`

HasXprotectVersion returns a boolean if a field has been set.

### GetAutoLoginDisabled

`func (o *ComputerSecurity) GetAutoLoginDisabled() bool`

GetAutoLoginDisabled returns the AutoLoginDisabled field if non-nil, zero value otherwise.

### GetAutoLoginDisabledOk

`func (o *ComputerSecurity) GetAutoLoginDisabledOk() (*bool, bool)`

GetAutoLoginDisabledOk returns a tuple with the AutoLoginDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoLoginDisabled

`func (o *ComputerSecurity) SetAutoLoginDisabled(v bool)`

SetAutoLoginDisabled sets AutoLoginDisabled field to given value.

### HasAutoLoginDisabled

`func (o *ComputerSecurity) HasAutoLoginDisabled() bool`

HasAutoLoginDisabled returns a boolean if a field has been set.

### GetRemoteDesktopEnabled

`func (o *ComputerSecurity) GetRemoteDesktopEnabled() bool`

GetRemoteDesktopEnabled returns the RemoteDesktopEnabled field if non-nil, zero value otherwise.

### GetRemoteDesktopEnabledOk

`func (o *ComputerSecurity) GetRemoteDesktopEnabledOk() (*bool, bool)`

GetRemoteDesktopEnabledOk returns a tuple with the RemoteDesktopEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteDesktopEnabled

`func (o *ComputerSecurity) SetRemoteDesktopEnabled(v bool)`

SetRemoteDesktopEnabled sets RemoteDesktopEnabled field to given value.

### HasRemoteDesktopEnabled

`func (o *ComputerSecurity) HasRemoteDesktopEnabled() bool`

HasRemoteDesktopEnabled returns a boolean if a field has been set.

### GetActivationLockEnabled

`func (o *ComputerSecurity) GetActivationLockEnabled() bool`

GetActivationLockEnabled returns the ActivationLockEnabled field if non-nil, zero value otherwise.

### GetActivationLockEnabledOk

`func (o *ComputerSecurity) GetActivationLockEnabledOk() (*bool, bool)`

GetActivationLockEnabledOk returns a tuple with the ActivationLockEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationLockEnabled

`func (o *ComputerSecurity) SetActivationLockEnabled(v bool)`

SetActivationLockEnabled sets ActivationLockEnabled field to given value.

### HasActivationLockEnabled

`func (o *ComputerSecurity) HasActivationLockEnabled() bool`

HasActivationLockEnabled returns a boolean if a field has been set.

### GetRecoveryLockEnabled

`func (o *ComputerSecurity) GetRecoveryLockEnabled() bool`

GetRecoveryLockEnabled returns the RecoveryLockEnabled field if non-nil, zero value otherwise.

### GetRecoveryLockEnabledOk

`func (o *ComputerSecurity) GetRecoveryLockEnabledOk() (*bool, bool)`

GetRecoveryLockEnabledOk returns a tuple with the RecoveryLockEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryLockEnabled

`func (o *ComputerSecurity) SetRecoveryLockEnabled(v bool)`

SetRecoveryLockEnabled sets RecoveryLockEnabled field to given value.

### HasRecoveryLockEnabled

`func (o *ComputerSecurity) HasRecoveryLockEnabled() bool`

HasRecoveryLockEnabled returns a boolean if a field has been set.

### GetFirewallEnabled

`func (o *ComputerSecurity) GetFirewallEnabled() bool`

GetFirewallEnabled returns the FirewallEnabled field if non-nil, zero value otherwise.

### GetFirewallEnabledOk

`func (o *ComputerSecurity) GetFirewallEnabledOk() (*bool, bool)`

GetFirewallEnabledOk returns a tuple with the FirewallEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallEnabled

`func (o *ComputerSecurity) SetFirewallEnabled(v bool)`

SetFirewallEnabled sets FirewallEnabled field to given value.

### HasFirewallEnabled

`func (o *ComputerSecurity) HasFirewallEnabled() bool`

HasFirewallEnabled returns a boolean if a field has been set.

### GetSecureBootLevel

`func (o *ComputerSecurity) GetSecureBootLevel() string`

GetSecureBootLevel returns the SecureBootLevel field if non-nil, zero value otherwise.

### GetSecureBootLevelOk

`func (o *ComputerSecurity) GetSecureBootLevelOk() (*string, bool)`

GetSecureBootLevelOk returns a tuple with the SecureBootLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureBootLevel

`func (o *ComputerSecurity) SetSecureBootLevel(v string)`

SetSecureBootLevel sets SecureBootLevel field to given value.

### HasSecureBootLevel

`func (o *ComputerSecurity) HasSecureBootLevel() bool`

HasSecureBootLevel returns a boolean if a field has been set.

### GetExternalBootLevel

`func (o *ComputerSecurity) GetExternalBootLevel() string`

GetExternalBootLevel returns the ExternalBootLevel field if non-nil, zero value otherwise.

### GetExternalBootLevelOk

`func (o *ComputerSecurity) GetExternalBootLevelOk() (*string, bool)`

GetExternalBootLevelOk returns a tuple with the ExternalBootLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalBootLevel

`func (o *ComputerSecurity) SetExternalBootLevel(v string)`

SetExternalBootLevel sets ExternalBootLevel field to given value.

### HasExternalBootLevel

`func (o *ComputerSecurity) HasExternalBootLevel() bool`

HasExternalBootLevel returns a boolean if a field has been set.

### GetBootstrapTokenAllowed

`func (o *ComputerSecurity) GetBootstrapTokenAllowed() bool`

GetBootstrapTokenAllowed returns the BootstrapTokenAllowed field if non-nil, zero value otherwise.

### GetBootstrapTokenAllowedOk

`func (o *ComputerSecurity) GetBootstrapTokenAllowedOk() (*bool, bool)`

GetBootstrapTokenAllowedOk returns a tuple with the BootstrapTokenAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootstrapTokenAllowed

`func (o *ComputerSecurity) SetBootstrapTokenAllowed(v bool)`

SetBootstrapTokenAllowed sets BootstrapTokenAllowed field to given value.

### HasBootstrapTokenAllowed

`func (o *ComputerSecurity) HasBootstrapTokenAllowed() bool`

HasBootstrapTokenAllowed returns a boolean if a field has been set.

### GetBootstrapTokenEscrowedStatus

`func (o *ComputerSecurity) GetBootstrapTokenEscrowedStatus() string`

GetBootstrapTokenEscrowedStatus returns the BootstrapTokenEscrowedStatus field if non-nil, zero value otherwise.

### GetBootstrapTokenEscrowedStatusOk

`func (o *ComputerSecurity) GetBootstrapTokenEscrowedStatusOk() (*string, bool)`

GetBootstrapTokenEscrowedStatusOk returns a tuple with the BootstrapTokenEscrowedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootstrapTokenEscrowedStatus

`func (o *ComputerSecurity) SetBootstrapTokenEscrowedStatus(v string)`

SetBootstrapTokenEscrowedStatus sets BootstrapTokenEscrowedStatus field to given value.

### HasBootstrapTokenEscrowedStatus

`func (o *ComputerSecurity) HasBootstrapTokenEscrowedStatus() bool`

HasBootstrapTokenEscrowedStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


