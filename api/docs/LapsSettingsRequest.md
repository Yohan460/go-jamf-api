# LapsSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoDeployEnabled** | **bool** | When enabled, all appropriate computers will have the SetAutoAdminPassword command sent to them automatically. | 
**PasswordRotationTime** | **int64** | The amount of time in seconds that the local admin password will be rotated after viewing. | 
**AutoExpirationTime** | **int64** | The amount of time in seconds that the local admin password will be rotated automatically if it is never viewed. | 

## Methods

### NewLapsSettingsRequest

`func NewLapsSettingsRequest(autoDeployEnabled bool, passwordRotationTime int64, autoExpirationTime int64, ) *LapsSettingsRequest`

NewLapsSettingsRequest instantiates a new LapsSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLapsSettingsRequestWithDefaults

`func NewLapsSettingsRequestWithDefaults() *LapsSettingsRequest`

NewLapsSettingsRequestWithDefaults instantiates a new LapsSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoDeployEnabled

`func (o *LapsSettingsRequest) GetAutoDeployEnabled() bool`

GetAutoDeployEnabled returns the AutoDeployEnabled field if non-nil, zero value otherwise.

### GetAutoDeployEnabledOk

`func (o *LapsSettingsRequest) GetAutoDeployEnabledOk() (*bool, bool)`

GetAutoDeployEnabledOk returns a tuple with the AutoDeployEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDeployEnabled

`func (o *LapsSettingsRequest) SetAutoDeployEnabled(v bool)`

SetAutoDeployEnabled sets AutoDeployEnabled field to given value.


### GetPasswordRotationTime

`func (o *LapsSettingsRequest) GetPasswordRotationTime() int64`

GetPasswordRotationTime returns the PasswordRotationTime field if non-nil, zero value otherwise.

### GetPasswordRotationTimeOk

`func (o *LapsSettingsRequest) GetPasswordRotationTimeOk() (*int64, bool)`

GetPasswordRotationTimeOk returns a tuple with the PasswordRotationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordRotationTime

`func (o *LapsSettingsRequest) SetPasswordRotationTime(v int64)`

SetPasswordRotationTime sets PasswordRotationTime field to given value.


### GetAutoExpirationTime

`func (o *LapsSettingsRequest) GetAutoExpirationTime() int64`

GetAutoExpirationTime returns the AutoExpirationTime field if non-nil, zero value otherwise.

### GetAutoExpirationTimeOk

`func (o *LapsSettingsRequest) GetAutoExpirationTimeOk() (*int64, bool)`

GetAutoExpirationTimeOk returns a tuple with the AutoExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoExpirationTime

`func (o *LapsSettingsRequest) SetAutoExpirationTime(v int64)`

SetAutoExpirationTime sets AutoExpirationTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


