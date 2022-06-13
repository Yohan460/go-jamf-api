# AuthTokenV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **string** |  | [optional] 
**Expires** | Pointer to **string** |  | [optional] 

## Methods

### NewAuthTokenV1

`func NewAuthTokenV1() *AuthTokenV1`

NewAuthTokenV1 instantiates a new AuthTokenV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthTokenV1WithDefaults

`func NewAuthTokenV1WithDefaults() *AuthTokenV1`

NewAuthTokenV1WithDefaults instantiates a new AuthTokenV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *AuthTokenV1) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AuthTokenV1) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AuthTokenV1) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AuthTokenV1) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetExpires

`func (o *AuthTokenV1) GetExpires() string`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *AuthTokenV1) GetExpiresOk() (*string, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *AuthTokenV1) SetExpires(v string)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *AuthTokenV1) HasExpires() bool`

HasExpires returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


