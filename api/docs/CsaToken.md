# CsaToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | Pointer to **NullableString** | The tenant ID | [optional] 
**Subject** | Pointer to **string** | Salesforce CRM account ID | [optional] 
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

### GetTenantId

`func (o *CsaToken) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CsaToken) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CsaToken) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CsaToken) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *CsaToken) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *CsaToken) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetSubject

`func (o *CsaToken) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CsaToken) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CsaToken) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CsaToken) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

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


