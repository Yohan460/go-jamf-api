# LapsSettingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoDeployEnabled** | Pointer to **bool** | When enabled, all appropriate computers will have the SetAutoAdminPassword command sent to them automatically. | [optional] 
**PasswordRotationTime** | Pointer to **int32** | The amount of time in seconds that the local admin password will be rotated after viewing. | [optional] 
**AutoExpirationTime** | Pointer to **int32** | The amount of time in seconds that the local admin password will be rotated automatically if it is never viewed. | [optional] 

## Methods

### NewLapsSettingsResponse

`func NewLapsSettingsResponse() *LapsSettingsResponse`

NewLapsSettingsResponse instantiates a new LapsSettingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLapsSettingsResponseWithDefaults

`func NewLapsSettingsResponseWithDefaults() *LapsSettingsResponse`

NewLapsSettingsResponseWithDefaults instantiates a new LapsSettingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoDeployEnabled

`func (o *LapsSettingsResponse) GetAutoDeployEnabled() bool`

GetAutoDeployEnabled returns the AutoDeployEnabled field if non-nil, zero value otherwise.

### GetAutoDeployEnabledOk

`func (o *LapsSettingsResponse) GetAutoDeployEnabledOk() (*bool, bool)`

GetAutoDeployEnabledOk returns a tuple with the AutoDeployEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDeployEnabled

`func (o *LapsSettingsResponse) SetAutoDeployEnabled(v bool)`

SetAutoDeployEnabled sets AutoDeployEnabled field to given value.

### HasAutoDeployEnabled

`func (o *LapsSettingsResponse) HasAutoDeployEnabled() bool`

HasAutoDeployEnabled returns a boolean if a field has been set.

### GetPasswordRotationTime

`func (o *LapsSettingsResponse) GetPasswordRotationTime() int32`

GetPasswordRotationTime returns the PasswordRotationTime field if non-nil, zero value otherwise.

### GetPasswordRotationTimeOk

`func (o *LapsSettingsResponse) GetPasswordRotationTimeOk() (*int32, bool)`

GetPasswordRotationTimeOk returns a tuple with the PasswordRotationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordRotationTime

`func (o *LapsSettingsResponse) SetPasswordRotationTime(v int32)`

SetPasswordRotationTime sets PasswordRotationTime field to given value.

### HasPasswordRotationTime

`func (o *LapsSettingsResponse) HasPasswordRotationTime() bool`

HasPasswordRotationTime returns a boolean if a field has been set.

### GetAutoExpirationTime

`func (o *LapsSettingsResponse) GetAutoExpirationTime() int32`

GetAutoExpirationTime returns the AutoExpirationTime field if non-nil, zero value otherwise.

### GetAutoExpirationTimeOk

`func (o *LapsSettingsResponse) GetAutoExpirationTimeOk() (*int32, bool)`

GetAutoExpirationTimeOk returns a tuple with the AutoExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoExpirationTime

`func (o *LapsSettingsResponse) SetAutoExpirationTime(v int32)`

SetAutoExpirationTime sets AutoExpirationTime field to given value.

### HasAutoExpirationTime

`func (o *LapsSettingsResponse) HasAutoExpirationTime() bool`

HasAutoExpirationTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


