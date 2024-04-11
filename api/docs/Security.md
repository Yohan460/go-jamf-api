# Security

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsDataProtected** | Pointer to **bool** |  | [optional] 
**IsBlockLevelEncryptionCapable** | Pointer to **bool** |  | [optional] 
**IsFileLevelEncryptionCapable** | Pointer to **bool** |  | [optional] 
**IsPasscodePresent** | Pointer to **bool** |  | [optional] 
**IsPasscodeCompliant** | Pointer to **bool** |  | [optional] 
**IsPasscodeCompliantWithProfile** | Pointer to **bool** |  | [optional] 
**HardwareEncryption** | Pointer to **int64** |  | [optional] 
**IsActivationLockEnabled** | Pointer to **bool** |  | [optional] 
**IsJailBreakDetected** | Pointer to **bool** |  | [optional] 

## Methods

### NewSecurity

`func NewSecurity() *Security`

NewSecurity instantiates a new Security object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityWithDefaults

`func NewSecurityWithDefaults() *Security`

NewSecurityWithDefaults instantiates a new Security object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsDataProtected

`func (o *Security) GetIsDataProtected() bool`

GetIsDataProtected returns the IsDataProtected field if non-nil, zero value otherwise.

### GetIsDataProtectedOk

`func (o *Security) GetIsDataProtectedOk() (*bool, bool)`

GetIsDataProtectedOk returns a tuple with the IsDataProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDataProtected

`func (o *Security) SetIsDataProtected(v bool)`

SetIsDataProtected sets IsDataProtected field to given value.

### HasIsDataProtected

`func (o *Security) HasIsDataProtected() bool`

HasIsDataProtected returns a boolean if a field has been set.

### GetIsBlockLevelEncryptionCapable

`func (o *Security) GetIsBlockLevelEncryptionCapable() bool`

GetIsBlockLevelEncryptionCapable returns the IsBlockLevelEncryptionCapable field if non-nil, zero value otherwise.

### GetIsBlockLevelEncryptionCapableOk

`func (o *Security) GetIsBlockLevelEncryptionCapableOk() (*bool, bool)`

GetIsBlockLevelEncryptionCapableOk returns a tuple with the IsBlockLevelEncryptionCapable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBlockLevelEncryptionCapable

`func (o *Security) SetIsBlockLevelEncryptionCapable(v bool)`

SetIsBlockLevelEncryptionCapable sets IsBlockLevelEncryptionCapable field to given value.

### HasIsBlockLevelEncryptionCapable

`func (o *Security) HasIsBlockLevelEncryptionCapable() bool`

HasIsBlockLevelEncryptionCapable returns a boolean if a field has been set.

### GetIsFileLevelEncryptionCapable

`func (o *Security) GetIsFileLevelEncryptionCapable() bool`

GetIsFileLevelEncryptionCapable returns the IsFileLevelEncryptionCapable field if non-nil, zero value otherwise.

### GetIsFileLevelEncryptionCapableOk

`func (o *Security) GetIsFileLevelEncryptionCapableOk() (*bool, bool)`

GetIsFileLevelEncryptionCapableOk returns a tuple with the IsFileLevelEncryptionCapable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFileLevelEncryptionCapable

`func (o *Security) SetIsFileLevelEncryptionCapable(v bool)`

SetIsFileLevelEncryptionCapable sets IsFileLevelEncryptionCapable field to given value.

### HasIsFileLevelEncryptionCapable

`func (o *Security) HasIsFileLevelEncryptionCapable() bool`

HasIsFileLevelEncryptionCapable returns a boolean if a field has been set.

### GetIsPasscodePresent

`func (o *Security) GetIsPasscodePresent() bool`

GetIsPasscodePresent returns the IsPasscodePresent field if non-nil, zero value otherwise.

### GetIsPasscodePresentOk

`func (o *Security) GetIsPasscodePresentOk() (*bool, bool)`

GetIsPasscodePresentOk returns a tuple with the IsPasscodePresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasscodePresent

`func (o *Security) SetIsPasscodePresent(v bool)`

SetIsPasscodePresent sets IsPasscodePresent field to given value.

### HasIsPasscodePresent

`func (o *Security) HasIsPasscodePresent() bool`

HasIsPasscodePresent returns a boolean if a field has been set.

### GetIsPasscodeCompliant

`func (o *Security) GetIsPasscodeCompliant() bool`

GetIsPasscodeCompliant returns the IsPasscodeCompliant field if non-nil, zero value otherwise.

### GetIsPasscodeCompliantOk

`func (o *Security) GetIsPasscodeCompliantOk() (*bool, bool)`

GetIsPasscodeCompliantOk returns a tuple with the IsPasscodeCompliant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasscodeCompliant

`func (o *Security) SetIsPasscodeCompliant(v bool)`

SetIsPasscodeCompliant sets IsPasscodeCompliant field to given value.

### HasIsPasscodeCompliant

`func (o *Security) HasIsPasscodeCompliant() bool`

HasIsPasscodeCompliant returns a boolean if a field has been set.

### GetIsPasscodeCompliantWithProfile

`func (o *Security) GetIsPasscodeCompliantWithProfile() bool`

GetIsPasscodeCompliantWithProfile returns the IsPasscodeCompliantWithProfile field if non-nil, zero value otherwise.

### GetIsPasscodeCompliantWithProfileOk

`func (o *Security) GetIsPasscodeCompliantWithProfileOk() (*bool, bool)`

GetIsPasscodeCompliantWithProfileOk returns a tuple with the IsPasscodeCompliantWithProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasscodeCompliantWithProfile

`func (o *Security) SetIsPasscodeCompliantWithProfile(v bool)`

SetIsPasscodeCompliantWithProfile sets IsPasscodeCompliantWithProfile field to given value.

### HasIsPasscodeCompliantWithProfile

`func (o *Security) HasIsPasscodeCompliantWithProfile() bool`

HasIsPasscodeCompliantWithProfile returns a boolean if a field has been set.

### GetHardwareEncryption

`func (o *Security) GetHardwareEncryption() int64`

GetHardwareEncryption returns the HardwareEncryption field if non-nil, zero value otherwise.

### GetHardwareEncryptionOk

`func (o *Security) GetHardwareEncryptionOk() (*int64, bool)`

GetHardwareEncryptionOk returns a tuple with the HardwareEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareEncryption

`func (o *Security) SetHardwareEncryption(v int64)`

SetHardwareEncryption sets HardwareEncryption field to given value.

### HasHardwareEncryption

`func (o *Security) HasHardwareEncryption() bool`

HasHardwareEncryption returns a boolean if a field has been set.

### GetIsActivationLockEnabled

`func (o *Security) GetIsActivationLockEnabled() bool`

GetIsActivationLockEnabled returns the IsActivationLockEnabled field if non-nil, zero value otherwise.

### GetIsActivationLockEnabledOk

`func (o *Security) GetIsActivationLockEnabledOk() (*bool, bool)`

GetIsActivationLockEnabledOk returns a tuple with the IsActivationLockEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActivationLockEnabled

`func (o *Security) SetIsActivationLockEnabled(v bool)`

SetIsActivationLockEnabled sets IsActivationLockEnabled field to given value.

### HasIsActivationLockEnabled

`func (o *Security) HasIsActivationLockEnabled() bool`

HasIsActivationLockEnabled returns a boolean if a field has been set.

### GetIsJailBreakDetected

`func (o *Security) GetIsJailBreakDetected() bool`

GetIsJailBreakDetected returns the IsJailBreakDetected field if non-nil, zero value otherwise.

### GetIsJailBreakDetectedOk

`func (o *Security) GetIsJailBreakDetectedOk() (*bool, bool)`

GetIsJailBreakDetectedOk returns a tuple with the IsJailBreakDetected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsJailBreakDetected

`func (o *Security) SetIsJailBreakDetected(v bool)`

SetIsJailBreakDetected sets IsJailBreakDetected field to given value.

### HasIsJailBreakDetected

`func (o *Security) HasIsJailBreakDetected() bool`

HasIsJailBreakDetected returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


