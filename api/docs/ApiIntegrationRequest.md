# ApiIntegrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizationScopes** | **[]string** | API Role display names.  | 
**DisplayName** | **string** |  | 
**Enabled** | Pointer to **bool** |  | [optional] 
**AccessTokenLifetimeSeconds** | Pointer to **int32** |  | [optional] 

## Methods

### NewApiIntegrationRequest

`func NewApiIntegrationRequest(authorizationScopes []string, displayName string, ) *ApiIntegrationRequest`

NewApiIntegrationRequest instantiates a new ApiIntegrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiIntegrationRequestWithDefaults

`func NewApiIntegrationRequestWithDefaults() *ApiIntegrationRequest`

NewApiIntegrationRequestWithDefaults instantiates a new ApiIntegrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizationScopes

`func (o *ApiIntegrationRequest) GetAuthorizationScopes() []string`

GetAuthorizationScopes returns the AuthorizationScopes field if non-nil, zero value otherwise.

### GetAuthorizationScopesOk

`func (o *ApiIntegrationRequest) GetAuthorizationScopesOk() (*[]string, bool)`

GetAuthorizationScopesOk returns a tuple with the AuthorizationScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationScopes

`func (o *ApiIntegrationRequest) SetAuthorizationScopes(v []string)`

SetAuthorizationScopes sets AuthorizationScopes field to given value.


### GetDisplayName

`func (o *ApiIntegrationRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ApiIntegrationRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ApiIntegrationRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetEnabled

`func (o *ApiIntegrationRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApiIntegrationRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApiIntegrationRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ApiIntegrationRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAccessTokenLifetimeSeconds

`func (o *ApiIntegrationRequest) GetAccessTokenLifetimeSeconds() int32`

GetAccessTokenLifetimeSeconds returns the AccessTokenLifetimeSeconds field if non-nil, zero value otherwise.

### GetAccessTokenLifetimeSecondsOk

`func (o *ApiIntegrationRequest) GetAccessTokenLifetimeSecondsOk() (*int32, bool)`

GetAccessTokenLifetimeSecondsOk returns a tuple with the AccessTokenLifetimeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenLifetimeSeconds

`func (o *ApiIntegrationRequest) SetAccessTokenLifetimeSeconds(v int32)`

SetAccessTokenLifetimeSeconds sets AccessTokenLifetimeSeconds field to given value.

### HasAccessTokenLifetimeSeconds

`func (o *ApiIntegrationRequest) HasAccessTokenLifetimeSeconds() bool`

HasAccessTokenLifetimeSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


