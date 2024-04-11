# CsaToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RefreshExpiration** | Pointer to **int64** |  | [optional] 
**Scopes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCsaToken

`func NewCsaToken() *CsaToken`

NewCsaToken instantiates a new CsaToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCsaTokenWithDefaults

`func NewCsaTokenWithDefaults() *CsaToken`

NewCsaTokenWithDefaults instantiates a new CsaToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefreshExpiration

`func (o *CsaToken) GetRefreshExpiration() int64`

GetRefreshExpiration returns the RefreshExpiration field if non-nil, zero value otherwise.

### GetRefreshExpirationOk

`func (o *CsaToken) GetRefreshExpirationOk() (*int64, bool)`

GetRefreshExpirationOk returns a tuple with the RefreshExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshExpiration

`func (o *CsaToken) SetRefreshExpiration(v int64)`

SetRefreshExpiration sets RefreshExpiration field to given value.

### HasRefreshExpiration

`func (o *CsaToken) HasRefreshExpiration() bool`

HasRefreshExpiration returns a boolean if a field has been set.

### GetScopes

`func (o *CsaToken) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *CsaToken) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *CsaToken) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *CsaToken) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


