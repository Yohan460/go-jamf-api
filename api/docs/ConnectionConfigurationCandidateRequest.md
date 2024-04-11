# ConnectionConfigurationCandidateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SiteId** | **string** | An identifier of a site which Team Viewer Remote Administration will be configured on | 
**DisplayName** | **string** | Name for Team Viewer Connection Configuration | 
**ScriptToken** | **string** | Token which is used for connecting to Team Viewer | 
**Enabled** | **bool** | Defines the intent to enable or disable Team Viewer connection | 
**SessionTimeout** | **int64** | Number of minutes before the session expires | 

## Methods

### NewConnectionConfigurationCandidateRequest

`func NewConnectionConfigurationCandidateRequest(siteId string, displayName string, scriptToken string, enabled bool, sessionTimeout int64, ) *ConnectionConfigurationCandidateRequest`

NewConnectionConfigurationCandidateRequest instantiates a new ConnectionConfigurationCandidateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionConfigurationCandidateRequestWithDefaults

`func NewConnectionConfigurationCandidateRequestWithDefaults() *ConnectionConfigurationCandidateRequest`

NewConnectionConfigurationCandidateRequestWithDefaults instantiates a new ConnectionConfigurationCandidateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiteId

`func (o *ConnectionConfigurationCandidateRequest) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *ConnectionConfigurationCandidateRequest) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *ConnectionConfigurationCandidateRequest) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.


### GetDisplayName

`func (o *ConnectionConfigurationCandidateRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ConnectionConfigurationCandidateRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ConnectionConfigurationCandidateRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetScriptToken

`func (o *ConnectionConfigurationCandidateRequest) GetScriptToken() string`

GetScriptToken returns the ScriptToken field if non-nil, zero value otherwise.

### GetScriptTokenOk

`func (o *ConnectionConfigurationCandidateRequest) GetScriptTokenOk() (*string, bool)`

GetScriptTokenOk returns a tuple with the ScriptToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptToken

`func (o *ConnectionConfigurationCandidateRequest) SetScriptToken(v string)`

SetScriptToken sets ScriptToken field to given value.


### GetEnabled

`func (o *ConnectionConfigurationCandidateRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ConnectionConfigurationCandidateRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ConnectionConfigurationCandidateRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetSessionTimeout

`func (o *ConnectionConfigurationCandidateRequest) GetSessionTimeout() int64`

GetSessionTimeout returns the SessionTimeout field if non-nil, zero value otherwise.

### GetSessionTimeoutOk

`func (o *ConnectionConfigurationCandidateRequest) GetSessionTimeoutOk() (*int64, bool)`

GetSessionTimeoutOk returns a tuple with the SessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTimeout

`func (o *ConnectionConfigurationCandidateRequest) SetSessionTimeout(v int64)`

SetSessionTimeout sets SessionTimeout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


