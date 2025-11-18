# OauthTokens

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to **string** |  | [optional] 
**IdToken** | Pointer to **string** |  | [optional] 

## Methods

### NewOauthTokens

`func NewOauthTokens() *OauthTokens`

NewOauthTokens instantiates a new OauthTokens object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOauthTokensWithDefaults

`func NewOauthTokensWithDefaults() *OauthTokens`

NewOauthTokensWithDefaults instantiates a new OauthTokens object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *OauthTokens) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *OauthTokens) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *OauthTokens) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *OauthTokens) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetIdToken

`func (o *OauthTokens) GetIdToken() string`

GetIdToken returns the IdToken field if non-nil, zero value otherwise.

### GetIdTokenOk

`func (o *OauthTokens) GetIdTokenOk() (*string, bool)`

GetIdTokenOk returns a tuple with the IdToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdToken

`func (o *OauthTokens) SetIdToken(v string)`

SetIdToken sets IdToken field to given value.

### HasIdToken

`func (o *OauthTokens) HasIdToken() bool`

HasIdToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


