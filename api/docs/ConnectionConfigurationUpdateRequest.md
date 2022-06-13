# ConnectionConfigurationUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | Name for Team Viewer Connection Configuration | [optional] 
**Enabled** | Pointer to **bool** | Defines the intent to enable or disable Team Viewer connection | [optional] 
**SessionTimeout** | Pointer to **int32** | Number of minutes before the session expires | [optional] 
**Token** | Pointer to **string** | Script token for Team Viewer Connection Configuration | [optional] 

## Methods

### NewConnectionConfigurationUpdateRequest

`func NewConnectionConfigurationUpdateRequest() *ConnectionConfigurationUpdateRequest`

NewConnectionConfigurationUpdateRequest instantiates a new ConnectionConfigurationUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionConfigurationUpdateRequestWithDefaults

`func NewConnectionConfigurationUpdateRequestWithDefaults() *ConnectionConfigurationUpdateRequest`

NewConnectionConfigurationUpdateRequestWithDefaults instantiates a new ConnectionConfigurationUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *ConnectionConfigurationUpdateRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ConnectionConfigurationUpdateRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ConnectionConfigurationUpdateRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ConnectionConfigurationUpdateRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEnabled

`func (o *ConnectionConfigurationUpdateRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ConnectionConfigurationUpdateRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ConnectionConfigurationUpdateRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ConnectionConfigurationUpdateRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSessionTimeout

`func (o *ConnectionConfigurationUpdateRequest) GetSessionTimeout() int32`

GetSessionTimeout returns the SessionTimeout field if non-nil, zero value otherwise.

### GetSessionTimeoutOk

`func (o *ConnectionConfigurationUpdateRequest) GetSessionTimeoutOk() (*int32, bool)`

GetSessionTimeoutOk returns a tuple with the SessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTimeout

`func (o *ConnectionConfigurationUpdateRequest) SetSessionTimeout(v int32)`

SetSessionTimeout sets SessionTimeout field to given value.

### HasSessionTimeout

`func (o *ConnectionConfigurationUpdateRequest) HasSessionTimeout() bool`

HasSessionTimeout returns a boolean if a field has been set.

### GetToken

`func (o *ConnectionConfigurationUpdateRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ConnectionConfigurationUpdateRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ConnectionConfigurationUpdateRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ConnectionConfigurationUpdateRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


