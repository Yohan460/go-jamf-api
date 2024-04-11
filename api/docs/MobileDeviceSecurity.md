# MobileDeviceSecurity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataProtected** | Pointer to **bool** |  | [optional] 
**BlockLevelEncryptionCapable** | Pointer to **bool** |  | [optional] 
**FileLevelEncryptionCapable** | Pointer to **bool** |  | [optional] 
**PasscodePresent** | Pointer to **bool** |  | [optional] 
**PasscodeCompliant** | Pointer to **bool** |  | [optional] 
**PasscodeCompliantWithProfile** | Pointer to **bool** |  | [optional] 
**HardwareEncryption** | Pointer to **int64** |  | [optional] 
**ActivationLockEnabled** | Pointer to **bool** |  | [optional] 
**JailBreakDetected** | Pointer to **bool** |  | [optional] 
**PasscodeLockGracePeriodEnforcedSeconds** | Pointer to **int64** |  | [optional] 
**PersonalDeviceProfileCurrent** | Pointer to **bool** |  | [optional] 
**LostModeEnabled** | Pointer to **bool** |  | [optional] 
**LostModePersistent** | Pointer to **bool** |  | [optional] 
**LostModeMessage** | Pointer to **string** |  | [optional] 
**LostModePhoneNumber** | Pointer to **string** |  | [optional] 
**LostModeFootnote** | Pointer to **string** |  | [optional] 
**LostModeLocation** | Pointer to [**MobileDeviceLostModeLocation**](MobileDeviceLostModeLocation.md) |  | [optional] 

## Methods

### NewMobileDeviceSecurity

`func NewMobileDeviceSecurity() *MobileDeviceSecurity`

NewMobileDeviceSecurity instantiates a new MobileDeviceSecurity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileDeviceSecurityWithDefaults

`func NewMobileDeviceSecurityWithDefaults() *MobileDeviceSecurity`

NewMobileDeviceSecurityWithDefaults instantiates a new MobileDeviceSecurity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataProtected

`func (o *MobileDeviceSecurity) GetDataProtected() bool`

GetDataProtected returns the DataProtected field if non-nil, zero value otherwise.

### GetDataProtectedOk

`func (o *MobileDeviceSecurity) GetDataProtectedOk() (*bool, bool)`

GetDataProtectedOk returns a tuple with the DataProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataProtected

`func (o *MobileDeviceSecurity) SetDataProtected(v bool)`

SetDataProtected sets DataProtected field to given value.

### HasDataProtected

`func (o *MobileDeviceSecurity) HasDataProtected() bool`

HasDataProtected returns a boolean if a field has been set.

### GetBlockLevelEncryptionCapable

`func (o *MobileDeviceSecurity) GetBlockLevelEncryptionCapable() bool`

GetBlockLevelEncryptionCapable returns the BlockLevelEncryptionCapable field if non-nil, zero value otherwise.

### GetBlockLevelEncryptionCapableOk

`func (o *MobileDeviceSecurity) GetBlockLevelEncryptionCapableOk() (*bool, bool)`

GetBlockLevelEncryptionCapableOk returns a tuple with the BlockLevelEncryptionCapable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockLevelEncryptionCapable

`func (o *MobileDeviceSecurity) SetBlockLevelEncryptionCapable(v bool)`

SetBlockLevelEncryptionCapable sets BlockLevelEncryptionCapable field to given value.

### HasBlockLevelEncryptionCapable

`func (o *MobileDeviceSecurity) HasBlockLevelEncryptionCapable() bool`

HasBlockLevelEncryptionCapable returns a boolean if a field has been set.

### GetFileLevelEncryptionCapable

`func (o *MobileDeviceSecurity) GetFileLevelEncryptionCapable() bool`

GetFileLevelEncryptionCapable returns the FileLevelEncryptionCapable field if non-nil, zero value otherwise.

### GetFileLevelEncryptionCapableOk

`func (o *MobileDeviceSecurity) GetFileLevelEncryptionCapableOk() (*bool, bool)`

GetFileLevelEncryptionCapableOk returns a tuple with the FileLevelEncryptionCapable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLevelEncryptionCapable

`func (o *MobileDeviceSecurity) SetFileLevelEncryptionCapable(v bool)`

SetFileLevelEncryptionCapable sets FileLevelEncryptionCapable field to given value.

### HasFileLevelEncryptionCapable

`func (o *MobileDeviceSecurity) HasFileLevelEncryptionCapable() bool`

HasFileLevelEncryptionCapable returns a boolean if a field has been set.

### GetPasscodePresent

`func (o *MobileDeviceSecurity) GetPasscodePresent() bool`

GetPasscodePresent returns the PasscodePresent field if non-nil, zero value otherwise.

### GetPasscodePresentOk

`func (o *MobileDeviceSecurity) GetPasscodePresentOk() (*bool, bool)`

GetPasscodePresentOk returns a tuple with the PasscodePresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasscodePresent

`func (o *MobileDeviceSecurity) SetPasscodePresent(v bool)`

SetPasscodePresent sets PasscodePresent field to given value.

### HasPasscodePresent

`func (o *MobileDeviceSecurity) HasPasscodePresent() bool`

HasPasscodePresent returns a boolean if a field has been set.

### GetPasscodeCompliant

`func (o *MobileDeviceSecurity) GetPasscodeCompliant() bool`

GetPasscodeCompliant returns the PasscodeCompliant field if non-nil, zero value otherwise.

### GetPasscodeCompliantOk

`func (o *MobileDeviceSecurity) GetPasscodeCompliantOk() (*bool, bool)`

GetPasscodeCompliantOk returns a tuple with the PasscodeCompliant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasscodeCompliant

`func (o *MobileDeviceSecurity) SetPasscodeCompliant(v bool)`

SetPasscodeCompliant sets PasscodeCompliant field to given value.

### HasPasscodeCompliant

`func (o *MobileDeviceSecurity) HasPasscodeCompliant() bool`

HasPasscodeCompliant returns a boolean if a field has been set.

### GetPasscodeCompliantWithProfile

`func (o *MobileDeviceSecurity) GetPasscodeCompliantWithProfile() bool`

GetPasscodeCompliantWithProfile returns the PasscodeCompliantWithProfile field if non-nil, zero value otherwise.

### GetPasscodeCompliantWithProfileOk

`func (o *MobileDeviceSecurity) GetPasscodeCompliantWithProfileOk() (*bool, bool)`

GetPasscodeCompliantWithProfileOk returns a tuple with the PasscodeCompliantWithProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasscodeCompliantWithProfile

`func (o *MobileDeviceSecurity) SetPasscodeCompliantWithProfile(v bool)`

SetPasscodeCompliantWithProfile sets PasscodeCompliantWithProfile field to given value.

### HasPasscodeCompliantWithProfile

`func (o *MobileDeviceSecurity) HasPasscodeCompliantWithProfile() bool`

HasPasscodeCompliantWithProfile returns a boolean if a field has been set.

### GetHardwareEncryption

`func (o *MobileDeviceSecurity) GetHardwareEncryption() int64`

GetHardwareEncryption returns the HardwareEncryption field if non-nil, zero value otherwise.

### GetHardwareEncryptionOk

`func (o *MobileDeviceSecurity) GetHardwareEncryptionOk() (*int64, bool)`

GetHardwareEncryptionOk returns a tuple with the HardwareEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareEncryption

`func (o *MobileDeviceSecurity) SetHardwareEncryption(v int64)`

SetHardwareEncryption sets HardwareEncryption field to given value.

### HasHardwareEncryption

`func (o *MobileDeviceSecurity) HasHardwareEncryption() bool`

HasHardwareEncryption returns a boolean if a field has been set.

### GetActivationLockEnabled

`func (o *MobileDeviceSecurity) GetActivationLockEnabled() bool`

GetActivationLockEnabled returns the ActivationLockEnabled field if non-nil, zero value otherwise.

### GetActivationLockEnabledOk

`func (o *MobileDeviceSecurity) GetActivationLockEnabledOk() (*bool, bool)`

GetActivationLockEnabledOk returns a tuple with the ActivationLockEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationLockEnabled

`func (o *MobileDeviceSecurity) SetActivationLockEnabled(v bool)`

SetActivationLockEnabled sets ActivationLockEnabled field to given value.

### HasActivationLockEnabled

`func (o *MobileDeviceSecurity) HasActivationLockEnabled() bool`

HasActivationLockEnabled returns a boolean if a field has been set.

### GetJailBreakDetected

`func (o *MobileDeviceSecurity) GetJailBreakDetected() bool`

GetJailBreakDetected returns the JailBreakDetected field if non-nil, zero value otherwise.

### GetJailBreakDetectedOk

`func (o *MobileDeviceSecurity) GetJailBreakDetectedOk() (*bool, bool)`

GetJailBreakDetectedOk returns a tuple with the JailBreakDetected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJailBreakDetected

`func (o *MobileDeviceSecurity) SetJailBreakDetected(v bool)`

SetJailBreakDetected sets JailBreakDetected field to given value.

### HasJailBreakDetected

`func (o *MobileDeviceSecurity) HasJailBreakDetected() bool`

HasJailBreakDetected returns a boolean if a field has been set.

### GetPasscodeLockGracePeriodEnforcedSeconds

`func (o *MobileDeviceSecurity) GetPasscodeLockGracePeriodEnforcedSeconds() int64`

GetPasscodeLockGracePeriodEnforcedSeconds returns the PasscodeLockGracePeriodEnforcedSeconds field if non-nil, zero value otherwise.

### GetPasscodeLockGracePeriodEnforcedSecondsOk

`func (o *MobileDeviceSecurity) GetPasscodeLockGracePeriodEnforcedSecondsOk() (*int64, bool)`

GetPasscodeLockGracePeriodEnforcedSecondsOk returns a tuple with the PasscodeLockGracePeriodEnforcedSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasscodeLockGracePeriodEnforcedSeconds

`func (o *MobileDeviceSecurity) SetPasscodeLockGracePeriodEnforcedSeconds(v int64)`

SetPasscodeLockGracePeriodEnforcedSeconds sets PasscodeLockGracePeriodEnforcedSeconds field to given value.

### HasPasscodeLockGracePeriodEnforcedSeconds

`func (o *MobileDeviceSecurity) HasPasscodeLockGracePeriodEnforcedSeconds() bool`

HasPasscodeLockGracePeriodEnforcedSeconds returns a boolean if a field has been set.

### GetPersonalDeviceProfileCurrent

`func (o *MobileDeviceSecurity) GetPersonalDeviceProfileCurrent() bool`

GetPersonalDeviceProfileCurrent returns the PersonalDeviceProfileCurrent field if non-nil, zero value otherwise.

### GetPersonalDeviceProfileCurrentOk

`func (o *MobileDeviceSecurity) GetPersonalDeviceProfileCurrentOk() (*bool, bool)`

GetPersonalDeviceProfileCurrentOk returns a tuple with the PersonalDeviceProfileCurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalDeviceProfileCurrent

`func (o *MobileDeviceSecurity) SetPersonalDeviceProfileCurrent(v bool)`

SetPersonalDeviceProfileCurrent sets PersonalDeviceProfileCurrent field to given value.

### HasPersonalDeviceProfileCurrent

`func (o *MobileDeviceSecurity) HasPersonalDeviceProfileCurrent() bool`

HasPersonalDeviceProfileCurrent returns a boolean if a field has been set.

### GetLostModeEnabled

`func (o *MobileDeviceSecurity) GetLostModeEnabled() bool`

GetLostModeEnabled returns the LostModeEnabled field if non-nil, zero value otherwise.

### GetLostModeEnabledOk

`func (o *MobileDeviceSecurity) GetLostModeEnabledOk() (*bool, bool)`

GetLostModeEnabledOk returns a tuple with the LostModeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLostModeEnabled

`func (o *MobileDeviceSecurity) SetLostModeEnabled(v bool)`

SetLostModeEnabled sets LostModeEnabled field to given value.

### HasLostModeEnabled

`func (o *MobileDeviceSecurity) HasLostModeEnabled() bool`

HasLostModeEnabled returns a boolean if a field has been set.

### GetLostModePersistent

`func (o *MobileDeviceSecurity) GetLostModePersistent() bool`

GetLostModePersistent returns the LostModePersistent field if non-nil, zero value otherwise.

### GetLostModePersistentOk

`func (o *MobileDeviceSecurity) GetLostModePersistentOk() (*bool, bool)`

GetLostModePersistentOk returns a tuple with the LostModePersistent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLostModePersistent

`func (o *MobileDeviceSecurity) SetLostModePersistent(v bool)`

SetLostModePersistent sets LostModePersistent field to given value.

### HasLostModePersistent

`func (o *MobileDeviceSecurity) HasLostModePersistent() bool`

HasLostModePersistent returns a boolean if a field has been set.

### GetLostModeMessage

`func (o *MobileDeviceSecurity) GetLostModeMessage() string`

GetLostModeMessage returns the LostModeMessage field if non-nil, zero value otherwise.

### GetLostModeMessageOk

`func (o *MobileDeviceSecurity) GetLostModeMessageOk() (*string, bool)`

GetLostModeMessageOk returns a tuple with the LostModeMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLostModeMessage

`func (o *MobileDeviceSecurity) SetLostModeMessage(v string)`

SetLostModeMessage sets LostModeMessage field to given value.

### HasLostModeMessage

`func (o *MobileDeviceSecurity) HasLostModeMessage() bool`

HasLostModeMessage returns a boolean if a field has been set.

### GetLostModePhoneNumber

`func (o *MobileDeviceSecurity) GetLostModePhoneNumber() string`

GetLostModePhoneNumber returns the LostModePhoneNumber field if non-nil, zero value otherwise.

### GetLostModePhoneNumberOk

`func (o *MobileDeviceSecurity) GetLostModePhoneNumberOk() (*string, bool)`

GetLostModePhoneNumberOk returns a tuple with the LostModePhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLostModePhoneNumber

`func (o *MobileDeviceSecurity) SetLostModePhoneNumber(v string)`

SetLostModePhoneNumber sets LostModePhoneNumber field to given value.

### HasLostModePhoneNumber

`func (o *MobileDeviceSecurity) HasLostModePhoneNumber() bool`

HasLostModePhoneNumber returns a boolean if a field has been set.

### GetLostModeFootnote

`func (o *MobileDeviceSecurity) GetLostModeFootnote() string`

GetLostModeFootnote returns the LostModeFootnote field if non-nil, zero value otherwise.

### GetLostModeFootnoteOk

`func (o *MobileDeviceSecurity) GetLostModeFootnoteOk() (*string, bool)`

GetLostModeFootnoteOk returns a tuple with the LostModeFootnote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLostModeFootnote

`func (o *MobileDeviceSecurity) SetLostModeFootnote(v string)`

SetLostModeFootnote sets LostModeFootnote field to given value.

### HasLostModeFootnote

`func (o *MobileDeviceSecurity) HasLostModeFootnote() bool`

HasLostModeFootnote returns a boolean if a field has been set.

### GetLostModeLocation

`func (o *MobileDeviceSecurity) GetLostModeLocation() MobileDeviceLostModeLocation`

GetLostModeLocation returns the LostModeLocation field if non-nil, zero value otherwise.

### GetLostModeLocationOk

`func (o *MobileDeviceSecurity) GetLostModeLocationOk() (*MobileDeviceLostModeLocation, bool)`

GetLostModeLocationOk returns a tuple with the LostModeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLostModeLocation

`func (o *MobileDeviceSecurity) SetLostModeLocation(v MobileDeviceLostModeLocation)`

SetLostModeLocation sets LostModeLocation field to given value.

### HasLostModeLocation

`func (o *MobileDeviceSecurity) HasLostModeLocation() bool`

HasLostModeLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


