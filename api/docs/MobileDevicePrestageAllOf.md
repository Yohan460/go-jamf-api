# MobileDevicePrestageAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAllowPairing** | **bool** |  | 
**IsMultiUser** | **bool** |  | 
**IsSupervised** | **bool** |  | 
**MaximumSharedAccounts** | **int32** |  | 
**IsAutoAdvanceSetup** | **bool** |  | 
**IsConfigureDeviceBeforeSetupAssistant** | **bool** |  | 
**Language** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Names** | Pointer to [**MobileDevicePrestageNames**](MobileDevicePrestageNames.md) |  | [optional] 

## Methods

### NewMobileDevicePrestageAllOf

`func NewMobileDevicePrestageAllOf(isAllowPairing bool, isMultiUser bool, isSupervised bool, maximumSharedAccounts int32, isAutoAdvanceSetup bool, isConfigureDeviceBeforeSetupAssistant bool, ) *MobileDevicePrestageAllOf`

NewMobileDevicePrestageAllOf instantiates a new MobileDevicePrestageAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileDevicePrestageAllOfWithDefaults

`func NewMobileDevicePrestageAllOfWithDefaults() *MobileDevicePrestageAllOf`

NewMobileDevicePrestageAllOfWithDefaults instantiates a new MobileDevicePrestageAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAllowPairing

`func (o *MobileDevicePrestageAllOf) GetIsAllowPairing() bool`

GetIsAllowPairing returns the IsAllowPairing field if non-nil, zero value otherwise.

### GetIsAllowPairingOk

`func (o *MobileDevicePrestageAllOf) GetIsAllowPairingOk() (*bool, bool)`

GetIsAllowPairingOk returns a tuple with the IsAllowPairing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAllowPairing

`func (o *MobileDevicePrestageAllOf) SetIsAllowPairing(v bool)`

SetIsAllowPairing sets IsAllowPairing field to given value.


### GetIsMultiUser

`func (o *MobileDevicePrestageAllOf) GetIsMultiUser() bool`

GetIsMultiUser returns the IsMultiUser field if non-nil, zero value otherwise.

### GetIsMultiUserOk

`func (o *MobileDevicePrestageAllOf) GetIsMultiUserOk() (*bool, bool)`

GetIsMultiUserOk returns a tuple with the IsMultiUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultiUser

`func (o *MobileDevicePrestageAllOf) SetIsMultiUser(v bool)`

SetIsMultiUser sets IsMultiUser field to given value.


### GetIsSupervised

`func (o *MobileDevicePrestageAllOf) GetIsSupervised() bool`

GetIsSupervised returns the IsSupervised field if non-nil, zero value otherwise.

### GetIsSupervisedOk

`func (o *MobileDevicePrestageAllOf) GetIsSupervisedOk() (*bool, bool)`

GetIsSupervisedOk returns a tuple with the IsSupervised field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSupervised

`func (o *MobileDevicePrestageAllOf) SetIsSupervised(v bool)`

SetIsSupervised sets IsSupervised field to given value.


### GetMaximumSharedAccounts

`func (o *MobileDevicePrestageAllOf) GetMaximumSharedAccounts() int32`

GetMaximumSharedAccounts returns the MaximumSharedAccounts field if non-nil, zero value otherwise.

### GetMaximumSharedAccountsOk

`func (o *MobileDevicePrestageAllOf) GetMaximumSharedAccountsOk() (*int32, bool)`

GetMaximumSharedAccountsOk returns a tuple with the MaximumSharedAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumSharedAccounts

`func (o *MobileDevicePrestageAllOf) SetMaximumSharedAccounts(v int32)`

SetMaximumSharedAccounts sets MaximumSharedAccounts field to given value.


### GetIsAutoAdvanceSetup

`func (o *MobileDevicePrestageAllOf) GetIsAutoAdvanceSetup() bool`

GetIsAutoAdvanceSetup returns the IsAutoAdvanceSetup field if non-nil, zero value otherwise.

### GetIsAutoAdvanceSetupOk

`func (o *MobileDevicePrestageAllOf) GetIsAutoAdvanceSetupOk() (*bool, bool)`

GetIsAutoAdvanceSetupOk returns a tuple with the IsAutoAdvanceSetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoAdvanceSetup

`func (o *MobileDevicePrestageAllOf) SetIsAutoAdvanceSetup(v bool)`

SetIsAutoAdvanceSetup sets IsAutoAdvanceSetup field to given value.


### GetIsConfigureDeviceBeforeSetupAssistant

`func (o *MobileDevicePrestageAllOf) GetIsConfigureDeviceBeforeSetupAssistant() bool`

GetIsConfigureDeviceBeforeSetupAssistant returns the IsConfigureDeviceBeforeSetupAssistant field if non-nil, zero value otherwise.

### GetIsConfigureDeviceBeforeSetupAssistantOk

`func (o *MobileDevicePrestageAllOf) GetIsConfigureDeviceBeforeSetupAssistantOk() (*bool, bool)`

GetIsConfigureDeviceBeforeSetupAssistantOk returns a tuple with the IsConfigureDeviceBeforeSetupAssistant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConfigureDeviceBeforeSetupAssistant

`func (o *MobileDevicePrestageAllOf) SetIsConfigureDeviceBeforeSetupAssistant(v bool)`

SetIsConfigureDeviceBeforeSetupAssistant sets IsConfigureDeviceBeforeSetupAssistant field to given value.


### GetLanguage

`func (o *MobileDevicePrestageAllOf) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *MobileDevicePrestageAllOf) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *MobileDevicePrestageAllOf) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *MobileDevicePrestageAllOf) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetRegion

`func (o *MobileDevicePrestageAllOf) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *MobileDevicePrestageAllOf) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *MobileDevicePrestageAllOf) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *MobileDevicePrestageAllOf) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetNames

`func (o *MobileDevicePrestageAllOf) GetNames() MobileDevicePrestageNames`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *MobileDevicePrestageAllOf) GetNamesOk() (*MobileDevicePrestageNames, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *MobileDevicePrestageAllOf) SetNames(v MobileDevicePrestageNames)`

SetNames sets Names field to given value.

### HasNames

`func (o *MobileDevicePrestageAllOf) HasNames() bool`

HasNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


