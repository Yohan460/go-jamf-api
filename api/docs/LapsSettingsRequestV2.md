# LapsSettingsRequestV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoDeployEnabled** | **bool** | When enabled, all appropriate computers will have the SetAutoAdminPassword command sent to them automatically. | 
**PasswordRotationTime** | **int64** | The amount of time in seconds that the local admin password will be rotated after viewing. | 
**AutoRotateEnabled** | **bool** | When enabled, all appropriate computers will automatically have their password expired and rotated after the configured autoRotateExpirationTime | 
**AutoRotateExpirationTime** | **int64** | The amount of time in seconds that the local admin password will be rotated automatically if it is never viewed. | 

## Methods

### NewLapsSettingsRequestV2

`func NewLapsSettingsRequestV2(autoDeployEnabled bool, passwordRotationTime int64, autoRotateEnabled bool, autoRotateExpirationTime int64, ) *LapsSettingsRequestV2`

NewLapsSettingsRequestV2 instantiates a new LapsSettingsRequestV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLapsSettingsRequestV2WithDefaults

`func NewLapsSettingsRequestV2WithDefaults() *LapsSettingsRequestV2`

NewLapsSettingsRequestV2WithDefaults instantiates a new LapsSettingsRequestV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoDeployEnabled

`func (o *LapsSettingsRequestV2) GetAutoDeployEnabled() bool`

GetAutoDeployEnabled returns the AutoDeployEnabled field if non-nil, zero value otherwise.

### GetAutoDeployEnabledOk

`func (o *LapsSettingsRequestV2) GetAutoDeployEnabledOk() (*bool, bool)`

GetAutoDeployEnabledOk returns a tuple with the AutoDeployEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDeployEnabled

`func (o *LapsSettingsRequestV2) SetAutoDeployEnabled(v bool)`

SetAutoDeployEnabled sets AutoDeployEnabled field to given value.


### GetPasswordRotationTime

`func (o *LapsSettingsRequestV2) GetPasswordRotationTime() int64`

GetPasswordRotationTime returns the PasswordRotationTime field if non-nil, zero value otherwise.

### GetPasswordRotationTimeOk

`func (o *LapsSettingsRequestV2) GetPasswordRotationTimeOk() (*int64, bool)`

GetPasswordRotationTimeOk returns a tuple with the PasswordRotationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordRotationTime

`func (o *LapsSettingsRequestV2) SetPasswordRotationTime(v int64)`

SetPasswordRotationTime sets PasswordRotationTime field to given value.


### GetAutoRotateEnabled

`func (o *LapsSettingsRequestV2) GetAutoRotateEnabled() bool`

GetAutoRotateEnabled returns the AutoRotateEnabled field if non-nil, zero value otherwise.

### GetAutoRotateEnabledOk

`func (o *LapsSettingsRequestV2) GetAutoRotateEnabledOk() (*bool, bool)`

GetAutoRotateEnabledOk returns a tuple with the AutoRotateEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRotateEnabled

`func (o *LapsSettingsRequestV2) SetAutoRotateEnabled(v bool)`

SetAutoRotateEnabled sets AutoRotateEnabled field to given value.


### GetAutoRotateExpirationTime

`func (o *LapsSettingsRequestV2) GetAutoRotateExpirationTime() int64`

GetAutoRotateExpirationTime returns the AutoRotateExpirationTime field if non-nil, zero value otherwise.

### GetAutoRotateExpirationTimeOk

`func (o *LapsSettingsRequestV2) GetAutoRotateExpirationTimeOk() (*int64, bool)`

GetAutoRotateExpirationTimeOk returns a tuple with the AutoRotateExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRotateExpirationTime

`func (o *LapsSettingsRequestV2) SetAutoRotateExpirationTime(v int64)`

SetAutoRotateExpirationTime sets AutoRotateExpirationTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


