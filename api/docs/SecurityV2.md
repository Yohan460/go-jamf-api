# SecurityV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataProtected** | Pointer to **bool** |  | [optional] 
**BlockLevelEncryptionCapable** | Pointer to **bool** |  | [optional] 
**FileLevelEncryptionCapable** | Pointer to **bool** |  | [optional] 
**PasscodePresent** | Pointer to **bool** |  | [optional] 
**PasscodeCompliant** | Pointer to **bool** |  | [optional] 
**PasscodeCompliantWithProfile** | Pointer to **bool** |  | [optional] 
**HardwareEncryption** | Pointer to **int32** |  | [optional] 
**ActivationLockEnabled** | Pointer to **bool** |  | [optional] 
**JailBreakDetected** | Pointer to **bool** |  | [optional] 

## Methods

### NewSecurityV2

`func NewSecurityV2() *SecurityV2`

NewSecurityV2 instantiates a new SecurityV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityV2WithDefaults

`func NewSecurityV2WithDefaults() *SecurityV2`

NewSecurityV2WithDefaults instantiates a new SecurityV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataProtected

`func (o *SecurityV2) GetDataProtected() bool`

GetDataProtected returns the DataProtected field if non-nil, zero value otherwise.

### GetDataProtectedOk

`func (o *SecurityV2) GetDataProtectedOk() (*bool, bool)`

GetDataProtectedOk returns a tuple with the DataProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataProtected

`func (o *SecurityV2) SetDataProtected(v bool)`

SetDataProtected sets DataProtected field to given value.

### HasDataProtected

`func (o *SecurityV2) HasDataProtected() bool`

HasDataProtected returns a boolean if a field has been set.

### GetBlockLevelEncryptionCapable

`func (o *SecurityV2) GetBlockLevelEncryptionCapable() bool`

GetBlockLevelEncryptionCapable returns the BlockLevelEncryptionCapable field if non-nil, zero value otherwise.

### GetBlockLevelEncryptionCapableOk

`func (o *SecurityV2) GetBlockLevelEncryptionCapableOk() (*bool, bool)`

GetBlockLevelEncryptionCapableOk returns a tuple with the BlockLevelEncryptionCapable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockLevelEncryptionCapable

`func (o *SecurityV2) SetBlockLevelEncryptionCapable(v bool)`

SetBlockLevelEncryptionCapable sets BlockLevelEncryptionCapable field to given value.

### HasBlockLevelEncryptionCapable

`func (o *SecurityV2) HasBlockLevelEncryptionCapable() bool`

HasBlockLevelEncryptionCapable returns a boolean if a field has been set.

### GetFileLevelEncryptionCapable

`func (o *SecurityV2) GetFileLevelEncryptionCapable() bool`

GetFileLevelEncryptionCapable returns the FileLevelEncryptionCapable field if non-nil, zero value otherwise.

### GetFileLevelEncryptionCapableOk

`func (o *SecurityV2) GetFileLevelEncryptionCapableOk() (*bool, bool)`

GetFileLevelEncryptionCapableOk returns a tuple with the FileLevelEncryptionCapable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLevelEncryptionCapable

`func (o *SecurityV2) SetFileLevelEncryptionCapable(v bool)`

SetFileLevelEncryptionCapable sets FileLevelEncryptionCapable field to given value.

### HasFileLevelEncryptionCapable

`func (o *SecurityV2) HasFileLevelEncryptionCapable() bool`

HasFileLevelEncryptionCapable returns a boolean if a field has been set.

### GetPasscodePresent

`func (o *SecurityV2) GetPasscodePresent() bool`

GetPasscodePresent returns the PasscodePresent field if non-nil, zero value otherwise.

### GetPasscodePresentOk

`func (o *SecurityV2) GetPasscodePresentOk() (*bool, bool)`

GetPasscodePresentOk returns a tuple with the PasscodePresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasscodePresent

`func (o *SecurityV2) SetPasscodePresent(v bool)`

SetPasscodePresent sets PasscodePresent field to given value.

### HasPasscodePresent

`func (o *SecurityV2) HasPasscodePresent() bool`

HasPasscodePresent returns a boolean if a field has been set.

### GetPasscodeCompliant

`func (o *SecurityV2) GetPasscodeCompliant() bool`

GetPasscodeCompliant returns the PasscodeCompliant field if non-nil, zero value otherwise.

### GetPasscodeCompliantOk

`func (o *SecurityV2) GetPasscodeCompliantOk() (*bool, bool)`

GetPasscodeCompliantOk returns a tuple with the PasscodeCompliant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasscodeCompliant

`func (o *SecurityV2) SetPasscodeCompliant(v bool)`

SetPasscodeCompliant sets PasscodeCompliant field to given value.

### HasPasscodeCompliant

`func (o *SecurityV2) HasPasscodeCompliant() bool`

HasPasscodeCompliant returns a boolean if a field has been set.

### GetPasscodeCompliantWithProfile

`func (o *SecurityV2) GetPasscodeCompliantWithProfile() bool`

GetPasscodeCompliantWithProfile returns the PasscodeCompliantWithProfile field if non-nil, zero value otherwise.

### GetPasscodeCompliantWithProfileOk

`func (o *SecurityV2) GetPasscodeCompliantWithProfileOk() (*bool, bool)`

GetPasscodeCompliantWithProfileOk returns a tuple with the PasscodeCompliantWithProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasscodeCompliantWithProfile

`func (o *SecurityV2) SetPasscodeCompliantWithProfile(v bool)`

SetPasscodeCompliantWithProfile sets PasscodeCompliantWithProfile field to given value.

### HasPasscodeCompliantWithProfile

`func (o *SecurityV2) HasPasscodeCompliantWithProfile() bool`

HasPasscodeCompliantWithProfile returns a boolean if a field has been set.

### GetHardwareEncryption

`func (o *SecurityV2) GetHardwareEncryption() int32`

GetHardwareEncryption returns the HardwareEncryption field if non-nil, zero value otherwise.

### GetHardwareEncryptionOk

`func (o *SecurityV2) GetHardwareEncryptionOk() (*int32, bool)`

GetHardwareEncryptionOk returns a tuple with the HardwareEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareEncryption

`func (o *SecurityV2) SetHardwareEncryption(v int32)`

SetHardwareEncryption sets HardwareEncryption field to given value.

### HasHardwareEncryption

`func (o *SecurityV2) HasHardwareEncryption() bool`

HasHardwareEncryption returns a boolean if a field has been set.

### GetActivationLockEnabled

`func (o *SecurityV2) GetActivationLockEnabled() bool`

GetActivationLockEnabled returns the ActivationLockEnabled field if non-nil, zero value otherwise.

### GetActivationLockEnabledOk

`func (o *SecurityV2) GetActivationLockEnabledOk() (*bool, bool)`

GetActivationLockEnabledOk returns a tuple with the ActivationLockEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationLockEnabled

`func (o *SecurityV2) SetActivationLockEnabled(v bool)`

SetActivationLockEnabled sets ActivationLockEnabled field to given value.

### HasActivationLockEnabled

`func (o *SecurityV2) HasActivationLockEnabled() bool`

HasActivationLockEnabled returns a boolean if a field has been set.

### GetJailBreakDetected

`func (o *SecurityV2) GetJailBreakDetected() bool`

GetJailBreakDetected returns the JailBreakDetected field if non-nil, zero value otherwise.

### GetJailBreakDetectedOk

`func (o *SecurityV2) GetJailBreakDetectedOk() (*bool, bool)`

GetJailBreakDetectedOk returns a tuple with the JailBreakDetected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJailBreakDetected

`func (o *SecurityV2) SetJailBreakDetected(v bool)`

SetJailBreakDetected sets JailBreakDetected field to given value.

### HasJailBreakDetected

`func (o *SecurityV2) HasJailBreakDetected() bool`

HasJailBreakDetected returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


