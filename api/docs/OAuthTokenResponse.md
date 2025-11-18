# OAuthTokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** |  | 
**TokenType** | **string** |  | 
**Scope** | **string** | Space-delimited list of API Roles, formatted and identified by ID like &#x60;api-role:{apiRoleId}&#x60;. | 
**ExpiresIn** | **int64** | Seconds until this access token is expired. This will reflect the &#x60;accessTokenLifetimeSeconds&#x60; value configured on your API Client. | 

## Methods

### NewOAuthTokenResponse

`func NewOAuthTokenResponse(accessToken string, tokenType string, scope string, expiresIn int64, ) *OAuthTokenResponse`

NewOAuthTokenResponse instantiates a new OAuthTokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthTokenResponseWithDefaults

`func NewOAuthTokenResponseWithDefaults() *OAuthTokenResponse`

NewOAuthTokenResponseWithDefaults instantiates a new OAuthTokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *OAuthTokenResponse) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *OAuthTokenResponse) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *OAuthTokenResponse) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetTokenType

`func (o *OAuthTokenResponse) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *OAuthTokenResponse) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *OAuthTokenResponse) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.


### GetScope

`func (o *OAuthTokenResponse) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *OAuthTokenResponse) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *OAuthTokenResponse) SetScope(v string)`

SetScope sets Scope field to given value.


### GetExpiresIn

`func (o *OAuthTokenResponse) GetExpiresIn() int64`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *OAuthTokenResponse) GetExpiresInOk() (*int64, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *OAuthTokenResponse) SetExpiresIn(v int64)`

SetExpiresIn sets ExpiresIn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


