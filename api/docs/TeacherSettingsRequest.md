# TeacherSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | Pointer to **bool** |  | [optional] 
**TimezoneId** | Pointer to **string** |  | [optional] 
**AutoClear** | Pointer to **NullableString** |  | [optional] 
**MaxRestrictionLengthSeconds** | Pointer to **int32** |  | [optional] 
**SafelistedApps** | Pointer to [**[]SafelistedAppsInner**](SafelistedAppsInner.md) |  | [optional] 

## Methods

### NewTeacherSettingsRequest

`func NewTeacherSettingsRequest() *TeacherSettingsRequest`

NewTeacherSettingsRequest instantiates a new TeacherSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeacherSettingsRequestWithDefaults

`func NewTeacherSettingsRequestWithDefaults() *TeacherSettingsRequest`

NewTeacherSettingsRequestWithDefaults instantiates a new TeacherSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *TeacherSettingsRequest) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *TeacherSettingsRequest) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *TeacherSettingsRequest) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *TeacherSettingsRequest) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetTimezoneId

`func (o *TeacherSettingsRequest) GetTimezoneId() string`

GetTimezoneId returns the TimezoneId field if non-nil, zero value otherwise.

### GetTimezoneIdOk

`func (o *TeacherSettingsRequest) GetTimezoneIdOk() (*string, bool)`

GetTimezoneIdOk returns a tuple with the TimezoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezoneId

`func (o *TeacherSettingsRequest) SetTimezoneId(v string)`

SetTimezoneId sets TimezoneId field to given value.

### HasTimezoneId

`func (o *TeacherSettingsRequest) HasTimezoneId() bool`

HasTimezoneId returns a boolean if a field has been set.

### GetAutoClear

`func (o *TeacherSettingsRequest) GetAutoClear() string`

GetAutoClear returns the AutoClear field if non-nil, zero value otherwise.

### GetAutoClearOk

`func (o *TeacherSettingsRequest) GetAutoClearOk() (*string, bool)`

GetAutoClearOk returns a tuple with the AutoClear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoClear

`func (o *TeacherSettingsRequest) SetAutoClear(v string)`

SetAutoClear sets AutoClear field to given value.

### HasAutoClear

`func (o *TeacherSettingsRequest) HasAutoClear() bool`

HasAutoClear returns a boolean if a field has been set.

### SetAutoClearNil

`func (o *TeacherSettingsRequest) SetAutoClearNil(b bool)`

 SetAutoClearNil sets the value for AutoClear to be an explicit nil

### UnsetAutoClear
`func (o *TeacherSettingsRequest) UnsetAutoClear()`

UnsetAutoClear ensures that no value is present for AutoClear, not even an explicit nil
### GetMaxRestrictionLengthSeconds

`func (o *TeacherSettingsRequest) GetMaxRestrictionLengthSeconds() int32`

GetMaxRestrictionLengthSeconds returns the MaxRestrictionLengthSeconds field if non-nil, zero value otherwise.

### GetMaxRestrictionLengthSecondsOk

`func (o *TeacherSettingsRequest) GetMaxRestrictionLengthSecondsOk() (*int32, bool)`

GetMaxRestrictionLengthSecondsOk returns a tuple with the MaxRestrictionLengthSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRestrictionLengthSeconds

`func (o *TeacherSettingsRequest) SetMaxRestrictionLengthSeconds(v int32)`

SetMaxRestrictionLengthSeconds sets MaxRestrictionLengthSeconds field to given value.

### HasMaxRestrictionLengthSeconds

`func (o *TeacherSettingsRequest) HasMaxRestrictionLengthSeconds() bool`

HasMaxRestrictionLengthSeconds returns a boolean if a field has been set.

### GetSafelistedApps

`func (o *TeacherSettingsRequest) GetSafelistedApps() []SafelistedAppsInner`

GetSafelistedApps returns the SafelistedApps field if non-nil, zero value otherwise.

### GetSafelistedAppsOk

`func (o *TeacherSettingsRequest) GetSafelistedAppsOk() (*[]SafelistedAppsInner, bool)`

GetSafelistedAppsOk returns a tuple with the SafelistedApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafelistedApps

`func (o *TeacherSettingsRequest) SetSafelistedApps(v []SafelistedAppsInner)`

SetSafelistedApps sets SafelistedApps field to given value.

### HasSafelistedApps

`func (o *TeacherSettingsRequest) HasSafelistedApps() bool`

HasSafelistedApps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


