# ApiIntegrationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**AuthorizationScopes** | **[]string** |  | 
**DisplayName** | **string** |  | 
**Enabled** | **bool** |  | 
**AccessTokenLifetimeSeconds** | Pointer to **int32** |  | [optional] 
**AppType** | **string** | Type of API Client:     * &#x60;CLIENT_CREDENTIALS&#x60; - A client ID and secret have been generated for this integration.     * &#x60;NATIVE_APP_OAUTH&#x60; - A native app (i.e., Jamf Reset) has been linked to this integration for auth code grant type via Managed App Config.     * &#x60;NONE&#x60; - No client is currently associated with this integration.  | [readonly] 
**ClientId** | **string** |  | [readonly] 

## Methods

### NewApiIntegrationResponse

`func NewApiIntegrationResponse(id int32, authorizationScopes []string, displayName string, enabled bool, appType string, clientId string, ) *ApiIntegrationResponse`

NewApiIntegrationResponse instantiates a new ApiIntegrationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiIntegrationResponseWithDefaults

`func NewApiIntegrationResponseWithDefaults() *ApiIntegrationResponse`

NewApiIntegrationResponseWithDefaults instantiates a new ApiIntegrationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiIntegrationResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiIntegrationResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiIntegrationResponse) SetId(v int32)`

SetId sets Id field to given value.


### GetAuthorizationScopes

`func (o *ApiIntegrationResponse) GetAuthorizationScopes() []string`

GetAuthorizationScopes returns the AuthorizationScopes field if non-nil, zero value otherwise.

### GetAuthorizationScopesOk

`func (o *ApiIntegrationResponse) GetAuthorizationScopesOk() (*[]string, bool)`

GetAuthorizationScopesOk returns a tuple with the AuthorizationScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationScopes

`func (o *ApiIntegrationResponse) SetAuthorizationScopes(v []string)`

SetAuthorizationScopes sets AuthorizationScopes field to given value.


### GetDisplayName

`func (o *ApiIntegrationResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ApiIntegrationResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ApiIntegrationResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetEnabled

`func (o *ApiIntegrationResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApiIntegrationResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApiIntegrationResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAccessTokenLifetimeSeconds

`func (o *ApiIntegrationResponse) GetAccessTokenLifetimeSeconds() int32`

GetAccessTokenLifetimeSeconds returns the AccessTokenLifetimeSeconds field if non-nil, zero value otherwise.

### GetAccessTokenLifetimeSecondsOk

`func (o *ApiIntegrationResponse) GetAccessTokenLifetimeSecondsOk() (*int32, bool)`

GetAccessTokenLifetimeSecondsOk returns a tuple with the AccessTokenLifetimeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenLifetimeSeconds

`func (o *ApiIntegrationResponse) SetAccessTokenLifetimeSeconds(v int32)`

SetAccessTokenLifetimeSeconds sets AccessTokenLifetimeSeconds field to given value.

### HasAccessTokenLifetimeSeconds

`func (o *ApiIntegrationResponse) HasAccessTokenLifetimeSeconds() bool`

HasAccessTokenLifetimeSeconds returns a boolean if a field has been set.

### GetAppType

`func (o *ApiIntegrationResponse) GetAppType() string`

GetAppType returns the AppType field if non-nil, zero value otherwise.

### GetAppTypeOk

`func (o *ApiIntegrationResponse) GetAppTypeOk() (*string, bool)`

GetAppTypeOk returns a tuple with the AppType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppType

`func (o *ApiIntegrationResponse) SetAppType(v string)`

SetAppType sets AppType field to given value.


### GetClientId

`func (o *ApiIntegrationResponse) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ApiIntegrationResponse) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ApiIntegrationResponse) SetClientId(v string)`

SetClientId sets ClientId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


