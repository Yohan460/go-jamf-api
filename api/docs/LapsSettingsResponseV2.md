# LapsSettingsResponseV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoDeployEnabled** | Pointer to **bool** | When enabled, all appropriate computers will have the SetAutoAdminPassword command sent to them automatically. | [optional] 
**PasswordRotationTime** | Pointer to **int64** | The amount of time in seconds that the local admin password will be rotated after viewing. | [optional] 
**AutoRotateEnabled** | Pointer to **bool** | When enabled, all appropriate computers will automatically have their password expired and rotated after the configured autoRotateExpirationTime | [optional] 
**AutoRotateExpirationTime** | Pointer to **int64** | The amount of time in seconds that the local admin password will be rotated automatically if it is never viewed. | [optional] 

## Methods

### NewLapsSettingsResponseV2

`func NewLapsSettingsResponseV2() *LapsSettingsResponseV2`

NewLapsSettingsResponseV2 instantiates a new LapsSettingsResponseV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLapsSettingsResponseV2WithDefaults

`func NewLapsSettingsResponseV2WithDefaults() *LapsSettingsResponseV2`

NewLapsSettingsResponseV2WithDefaults instantiates a new LapsSettingsResponseV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoDeployEnabled

`func (o *LapsSettingsResponseV2) GetAutoDeployEnabled() bool`

GetAutoDeployEnabled returns the AutoDeployEnabled field if non-nil, zero value otherwise.

### GetAutoDeployEnabledOk

`func (o *LapsSettingsResponseV2) GetAutoDeployEnabledOk() (*bool, bool)`

GetAutoDeployEnabledOk returns a tuple with the AutoDeployEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDeployEnabled

`func (o *LapsSettingsResponseV2) SetAutoDeployEnabled(v bool)`

SetAutoDeployEnabled sets AutoDeployEnabled field to given value.

### HasAutoDeployEnabled

`func (o *LapsSettingsResponseV2) HasAutoDeployEnabled() bool`

HasAutoDeployEnabled returns a boolean if a field has been set.

### GetPasswordRotationTime

`func (o *LapsSettingsResponseV2) GetPasswordRotationTime() int64`

GetPasswordRotationTime returns the PasswordRotationTime field if non-nil, zero value otherwise.

### GetPasswordRotationTimeOk

`func (o *LapsSettingsResponseV2) GetPasswordRotationTimeOk() (*int64, bool)`

GetPasswordRotationTimeOk returns a tuple with the PasswordRotationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordRotationTime

`func (o *LapsSettingsResponseV2) SetPasswordRotationTime(v int64)`

SetPasswordRotationTime sets PasswordRotationTime field to given value.

### HasPasswordRotationTime

`func (o *LapsSettingsResponseV2) HasPasswordRotationTime() bool`

HasPasswordRotationTime returns a boolean if a field has been set.

### GetAutoRotateEnabled

`func (o *LapsSettingsResponseV2) GetAutoRotateEnabled() bool`

GetAutoRotateEnabled returns the AutoRotateEnabled field if non-nil, zero value otherwise.

### GetAutoRotateEnabledOk

`func (o *LapsSettingsResponseV2) GetAutoRotateEnabledOk() (*bool, bool)`

GetAutoRotateEnabledOk returns a tuple with the AutoRotateEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRotateEnabled

`func (o *LapsSettingsResponseV2) SetAutoRotateEnabled(v bool)`

SetAutoRotateEnabled sets AutoRotateEnabled field to given value.

### HasAutoRotateEnabled

`func (o *LapsSettingsResponseV2) HasAutoRotateEnabled() bool`

HasAutoRotateEnabled returns a boolean if a field has been set.

### GetAutoRotateExpirationTime

`func (o *LapsSettingsResponseV2) GetAutoRotateExpirationTime() int64`

GetAutoRotateExpirationTime returns the AutoRotateExpirationTime field if non-nil, zero value otherwise.

### GetAutoRotateExpirationTimeOk

`func (o *LapsSettingsResponseV2) GetAutoRotateExpirationTimeOk() (*int64, bool)`

GetAutoRotateExpirationTimeOk returns a tuple with the AutoRotateExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRotateExpirationTime

`func (o *LapsSettingsResponseV2) SetAutoRotateExpirationTime(v int64)`

SetAutoRotateExpirationTime sets AutoRotateExpirationTime field to given value.

### HasAutoRotateExpirationTime

`func (o *LapsSettingsResponseV2) HasAutoRotateExpirationTime() bool`

HasAutoRotateExpirationTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


