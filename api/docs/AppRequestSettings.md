# AppRequestSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | Pointer to **bool** |  | [optional] 
**AppStoreLocale** | Pointer to **string** | Can be any of the country codes from /v1/app-store-country-codes or \&quot;deviceLocale\&quot; to use each individual device&#39;s locale | [optional] 
**RequesterUserGroupId** | Pointer to **int32** |  | [optional] 
**ApproverEmails** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAppRequestSettings

`func NewAppRequestSettings() *AppRequestSettings`

NewAppRequestSettings instantiates a new AppRequestSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppRequestSettingsWithDefaults

`func NewAppRequestSettingsWithDefaults() *AppRequestSettings`

NewAppRequestSettingsWithDefaults instantiates a new AppRequestSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *AppRequestSettings) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *AppRequestSettings) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *AppRequestSettings) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *AppRequestSettings) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetAppStoreLocale

`func (o *AppRequestSettings) GetAppStoreLocale() string`

GetAppStoreLocale returns the AppStoreLocale field if non-nil, zero value otherwise.

### GetAppStoreLocaleOk

`func (o *AppRequestSettings) GetAppStoreLocaleOk() (*string, bool)`

GetAppStoreLocaleOk returns a tuple with the AppStoreLocale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppStoreLocale

`func (o *AppRequestSettings) SetAppStoreLocale(v string)`

SetAppStoreLocale sets AppStoreLocale field to given value.

### HasAppStoreLocale

`func (o *AppRequestSettings) HasAppStoreLocale() bool`

HasAppStoreLocale returns a boolean if a field has been set.

### GetRequesterUserGroupId

`func (o *AppRequestSettings) GetRequesterUserGroupId() int32`

GetRequesterUserGroupId returns the RequesterUserGroupId field if non-nil, zero value otherwise.

### GetRequesterUserGroupIdOk

`func (o *AppRequestSettings) GetRequesterUserGroupIdOk() (*int32, bool)`

GetRequesterUserGroupIdOk returns a tuple with the RequesterUserGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterUserGroupId

`func (o *AppRequestSettings) SetRequesterUserGroupId(v int32)`

SetRequesterUserGroupId sets RequesterUserGroupId field to given value.

### HasRequesterUserGroupId

`func (o *AppRequestSettings) HasRequesterUserGroupId() bool`

HasRequesterUserGroupId returns a boolean if a field has been set.

### GetApproverEmails

`func (o *AppRequestSettings) GetApproverEmails() []string`

GetApproverEmails returns the ApproverEmails field if non-nil, zero value otherwise.

### GetApproverEmailsOk

`func (o *AppRequestSettings) GetApproverEmailsOk() (*[]string, bool)`

GetApproverEmailsOk returns a tuple with the ApproverEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverEmails

`func (o *AppRequestSettings) SetApproverEmails(v []string)`

SetApproverEmails sets ApproverEmails field to given value.

### HasApproverEmails

`func (o *AppRequestSettings) HasApproverEmails() bool`

HasApproverEmails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


