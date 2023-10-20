# TeacherSettingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | Pointer to **bool** |  | [optional] 
**TimezoneId** | Pointer to **string** |  | [optional] 
**AutoClear** | Pointer to **string** |  | [optional] 
**MaxRestrictionLengthSeconds** | Pointer to **int32** |  | [optional] 
**DisplayNameType** | Pointer to **string** |  | [optional] 
**Features** | Pointer to [**TeacherFeatures**](TeacherFeatures.md) |  | [optional] 
**SafelistedApps** | Pointer to [**[]SafelistedApp**](SafelistedApp.md) |  | [optional] 

## Methods

### NewTeacherSettingsResponse

`func NewTeacherSettingsResponse() *TeacherSettingsResponse`

NewTeacherSettingsResponse instantiates a new TeacherSettingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeacherSettingsResponseWithDefaults

`func NewTeacherSettingsResponseWithDefaults() *TeacherSettingsResponse`

NewTeacherSettingsResponseWithDefaults instantiates a new TeacherSettingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *TeacherSettingsResponse) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *TeacherSettingsResponse) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *TeacherSettingsResponse) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *TeacherSettingsResponse) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetTimezoneId

`func (o *TeacherSettingsResponse) GetTimezoneId() string`

GetTimezoneId returns the TimezoneId field if non-nil, zero value otherwise.

### GetTimezoneIdOk

`func (o *TeacherSettingsResponse) GetTimezoneIdOk() (*string, bool)`

GetTimezoneIdOk returns a tuple with the TimezoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezoneId

`func (o *TeacherSettingsResponse) SetTimezoneId(v string)`

SetTimezoneId sets TimezoneId field to given value.

### HasTimezoneId

`func (o *TeacherSettingsResponse) HasTimezoneId() bool`

HasTimezoneId returns a boolean if a field has been set.

### GetAutoClear

`func (o *TeacherSettingsResponse) GetAutoClear() string`

GetAutoClear returns the AutoClear field if non-nil, zero value otherwise.

### GetAutoClearOk

`func (o *TeacherSettingsResponse) GetAutoClearOk() (*string, bool)`

GetAutoClearOk returns a tuple with the AutoClear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoClear

`func (o *TeacherSettingsResponse) SetAutoClear(v string)`

SetAutoClear sets AutoClear field to given value.

### HasAutoClear

`func (o *TeacherSettingsResponse) HasAutoClear() bool`

HasAutoClear returns a boolean if a field has been set.

### GetMaxRestrictionLengthSeconds

`func (o *TeacherSettingsResponse) GetMaxRestrictionLengthSeconds() int32`

GetMaxRestrictionLengthSeconds returns the MaxRestrictionLengthSeconds field if non-nil, zero value otherwise.

### GetMaxRestrictionLengthSecondsOk

`func (o *TeacherSettingsResponse) GetMaxRestrictionLengthSecondsOk() (*int32, bool)`

GetMaxRestrictionLengthSecondsOk returns a tuple with the MaxRestrictionLengthSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRestrictionLengthSeconds

`func (o *TeacherSettingsResponse) SetMaxRestrictionLengthSeconds(v int32)`

SetMaxRestrictionLengthSeconds sets MaxRestrictionLengthSeconds field to given value.

### HasMaxRestrictionLengthSeconds

`func (o *TeacherSettingsResponse) HasMaxRestrictionLengthSeconds() bool`

HasMaxRestrictionLengthSeconds returns a boolean if a field has been set.

### GetDisplayNameType

`func (o *TeacherSettingsResponse) GetDisplayNameType() string`

GetDisplayNameType returns the DisplayNameType field if non-nil, zero value otherwise.

### GetDisplayNameTypeOk

`func (o *TeacherSettingsResponse) GetDisplayNameTypeOk() (*string, bool)`

GetDisplayNameTypeOk returns a tuple with the DisplayNameType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNameType

`func (o *TeacherSettingsResponse) SetDisplayNameType(v string)`

SetDisplayNameType sets DisplayNameType field to given value.

### HasDisplayNameType

`func (o *TeacherSettingsResponse) HasDisplayNameType() bool`

HasDisplayNameType returns a boolean if a field has been set.

### GetFeatures

`func (o *TeacherSettingsResponse) GetFeatures() TeacherFeatures`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *TeacherSettingsResponse) GetFeaturesOk() (*TeacherFeatures, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *TeacherSettingsResponse) SetFeatures(v TeacherFeatures)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *TeacherSettingsResponse) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetSafelistedApps

`func (o *TeacherSettingsResponse) GetSafelistedApps() []SafelistedApp`

GetSafelistedApps returns the SafelistedApps field if non-nil, zero value otherwise.

### GetSafelistedAppsOk

`func (o *TeacherSettingsResponse) GetSafelistedAppsOk() (*[]SafelistedApp, bool)`

GetSafelistedAppsOk returns a tuple with the SafelistedApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafelistedApps

`func (o *TeacherSettingsResponse) SetSafelistedApps(v []SafelistedApp)`

SetSafelistedApps sets SafelistedApps field to given value.

### HasSafelistedApps

`func (o *TeacherSettingsResponse) HasSafelistedApps() bool`

HasSafelistedApps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


