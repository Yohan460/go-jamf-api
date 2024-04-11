# ParentApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimezoneId** | **string** |  | 
**RestrictedTimes** | [**ParentAppRestrictedTimes**](ParentAppRestrictedTimes.md) |  | 
**DeviceGroupId** | **int64** |  | 
**IsEnabled** | **bool** |  | 
**AllowTemplates** | Pointer to **bool** |  | [optional] 
**DisassociateOnWipeAndReEnroll** | Pointer to **bool** |  | [optional] 
**AllowClearPasscode** | Pointer to **bool** |  | [optional] 
**SafelistedApps** | Pointer to [**[]SafelistedApp**](SafelistedApp.md) |  | [optional] 

## Methods

### NewParentApp

`func NewParentApp(timezoneId string, restrictedTimes ParentAppRestrictedTimes, deviceGroupId int64, isEnabled bool, ) *ParentApp`

NewParentApp instantiates a new ParentApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParentAppWithDefaults

`func NewParentAppWithDefaults() *ParentApp`

NewParentAppWithDefaults instantiates a new ParentApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimezoneId

`func (o *ParentApp) GetTimezoneId() string`

GetTimezoneId returns the TimezoneId field if non-nil, zero value otherwise.

### GetTimezoneIdOk

`func (o *ParentApp) GetTimezoneIdOk() (*string, bool)`

GetTimezoneIdOk returns a tuple with the TimezoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezoneId

`func (o *ParentApp) SetTimezoneId(v string)`

SetTimezoneId sets TimezoneId field to given value.


### GetRestrictedTimes

`func (o *ParentApp) GetRestrictedTimes() ParentAppRestrictedTimes`

GetRestrictedTimes returns the RestrictedTimes field if non-nil, zero value otherwise.

### GetRestrictedTimesOk

`func (o *ParentApp) GetRestrictedTimesOk() (*ParentAppRestrictedTimes, bool)`

GetRestrictedTimesOk returns a tuple with the RestrictedTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedTimes

`func (o *ParentApp) SetRestrictedTimes(v ParentAppRestrictedTimes)`

SetRestrictedTimes sets RestrictedTimes field to given value.


### GetDeviceGroupId

`func (o *ParentApp) GetDeviceGroupId() int64`

GetDeviceGroupId returns the DeviceGroupId field if non-nil, zero value otherwise.

### GetDeviceGroupIdOk

`func (o *ParentApp) GetDeviceGroupIdOk() (*int64, bool)`

GetDeviceGroupIdOk returns a tuple with the DeviceGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceGroupId

`func (o *ParentApp) SetDeviceGroupId(v int64)`

SetDeviceGroupId sets DeviceGroupId field to given value.


### GetIsEnabled

`func (o *ParentApp) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ParentApp) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ParentApp) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetAllowTemplates

`func (o *ParentApp) GetAllowTemplates() bool`

GetAllowTemplates returns the AllowTemplates field if non-nil, zero value otherwise.

### GetAllowTemplatesOk

`func (o *ParentApp) GetAllowTemplatesOk() (*bool, bool)`

GetAllowTemplatesOk returns a tuple with the AllowTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowTemplates

`func (o *ParentApp) SetAllowTemplates(v bool)`

SetAllowTemplates sets AllowTemplates field to given value.

### HasAllowTemplates

`func (o *ParentApp) HasAllowTemplates() bool`

HasAllowTemplates returns a boolean if a field has been set.

### GetDisassociateOnWipeAndReEnroll

`func (o *ParentApp) GetDisassociateOnWipeAndReEnroll() bool`

GetDisassociateOnWipeAndReEnroll returns the DisassociateOnWipeAndReEnroll field if non-nil, zero value otherwise.

### GetDisassociateOnWipeAndReEnrollOk

`func (o *ParentApp) GetDisassociateOnWipeAndReEnrollOk() (*bool, bool)`

GetDisassociateOnWipeAndReEnrollOk returns a tuple with the DisassociateOnWipeAndReEnroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisassociateOnWipeAndReEnroll

`func (o *ParentApp) SetDisassociateOnWipeAndReEnroll(v bool)`

SetDisassociateOnWipeAndReEnroll sets DisassociateOnWipeAndReEnroll field to given value.

### HasDisassociateOnWipeAndReEnroll

`func (o *ParentApp) HasDisassociateOnWipeAndReEnroll() bool`

HasDisassociateOnWipeAndReEnroll returns a boolean if a field has been set.

### GetAllowClearPasscode

`func (o *ParentApp) GetAllowClearPasscode() bool`

GetAllowClearPasscode returns the AllowClearPasscode field if non-nil, zero value otherwise.

### GetAllowClearPasscodeOk

`func (o *ParentApp) GetAllowClearPasscodeOk() (*bool, bool)`

GetAllowClearPasscodeOk returns a tuple with the AllowClearPasscode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowClearPasscode

`func (o *ParentApp) SetAllowClearPasscode(v bool)`

SetAllowClearPasscode sets AllowClearPasscode field to given value.

### HasAllowClearPasscode

`func (o *ParentApp) HasAllowClearPasscode() bool`

HasAllowClearPasscode returns a boolean if a field has been set.

### GetSafelistedApps

`func (o *ParentApp) GetSafelistedApps() []SafelistedApp`

GetSafelistedApps returns the SafelistedApps field if non-nil, zero value otherwise.

### GetSafelistedAppsOk

`func (o *ParentApp) GetSafelistedAppsOk() (*[]SafelistedApp, bool)`

GetSafelistedAppsOk returns a tuple with the SafelistedApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafelistedApps

`func (o *ParentApp) SetSafelistedApps(v []SafelistedApp)`

SetSafelistedApps sets SafelistedApps field to given value.

### HasSafelistedApps

`func (o *ParentApp) HasSafelistedApps() bool`

HasSafelistedApps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


