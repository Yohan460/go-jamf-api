# VenafiCaRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] [readonly] 
**Name** | **string** |  | 
**Username** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**PasswordConfigured** | Pointer to **bool** |  | [optional] [readonly] 
**ProxyAddress** | Pointer to **string** |  | [optional] 
**RevocationEnabled** | Pointer to **bool** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**RefreshToken** | Pointer to **string** |  | [optional] 
**RefreshTokenConfigured** | Pointer to **bool** |  | [optional] [readonly] 

## Methods

### NewVenafiCaRecord

`func NewVenafiCaRecord(name string, ) *VenafiCaRecord`

NewVenafiCaRecord instantiates a new VenafiCaRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVenafiCaRecordWithDefaults

`func NewVenafiCaRecordWithDefaults() *VenafiCaRecord`

NewVenafiCaRecordWithDefaults instantiates a new VenafiCaRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VenafiCaRecord) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VenafiCaRecord) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VenafiCaRecord) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *VenafiCaRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *VenafiCaRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VenafiCaRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VenafiCaRecord) SetName(v string)`

SetName sets Name field to given value.


### GetUsername

`func (o *VenafiCaRecord) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *VenafiCaRecord) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *VenafiCaRecord) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *VenafiCaRecord) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *VenafiCaRecord) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *VenafiCaRecord) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *VenafiCaRecord) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *VenafiCaRecord) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPasswordConfigured

`func (o *VenafiCaRecord) GetPasswordConfigured() bool`

GetPasswordConfigured returns the PasswordConfigured field if non-nil, zero value otherwise.

### GetPasswordConfiguredOk

`func (o *VenafiCaRecord) GetPasswordConfiguredOk() (*bool, bool)`

GetPasswordConfiguredOk returns a tuple with the PasswordConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordConfigured

`func (o *VenafiCaRecord) SetPasswordConfigured(v bool)`

SetPasswordConfigured sets PasswordConfigured field to given value.

### HasPasswordConfigured

`func (o *VenafiCaRecord) HasPasswordConfigured() bool`

HasPasswordConfigured returns a boolean if a field has been set.

### GetProxyAddress

`func (o *VenafiCaRecord) GetProxyAddress() string`

GetProxyAddress returns the ProxyAddress field if non-nil, zero value otherwise.

### GetProxyAddressOk

`func (o *VenafiCaRecord) GetProxyAddressOk() (*string, bool)`

GetProxyAddressOk returns a tuple with the ProxyAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyAddress

`func (o *VenafiCaRecord) SetProxyAddress(v string)`

SetProxyAddress sets ProxyAddress field to given value.

### HasProxyAddress

`func (o *VenafiCaRecord) HasProxyAddress() bool`

HasProxyAddress returns a boolean if a field has been set.

### GetRevocationEnabled

`func (o *VenafiCaRecord) GetRevocationEnabled() bool`

GetRevocationEnabled returns the RevocationEnabled field if non-nil, zero value otherwise.

### GetRevocationEnabledOk

`func (o *VenafiCaRecord) GetRevocationEnabledOk() (*bool, bool)`

GetRevocationEnabledOk returns a tuple with the RevocationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationEnabled

`func (o *VenafiCaRecord) SetRevocationEnabled(v bool)`

SetRevocationEnabled sets RevocationEnabled field to given value.

### HasRevocationEnabled

`func (o *VenafiCaRecord) HasRevocationEnabled() bool`

HasRevocationEnabled returns a boolean if a field has been set.

### GetClientId

`func (o *VenafiCaRecord) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *VenafiCaRecord) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *VenafiCaRecord) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *VenafiCaRecord) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetRefreshToken

`func (o *VenafiCaRecord) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *VenafiCaRecord) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *VenafiCaRecord) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *VenafiCaRecord) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetRefreshTokenConfigured

`func (o *VenafiCaRecord) GetRefreshTokenConfigured() bool`

GetRefreshTokenConfigured returns the RefreshTokenConfigured field if non-nil, zero value otherwise.

### GetRefreshTokenConfiguredOk

`func (o *VenafiCaRecord) GetRefreshTokenConfiguredOk() (*bool, bool)`

GetRefreshTokenConfiguredOk returns a tuple with the RefreshTokenConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenConfigured

`func (o *VenafiCaRecord) SetRefreshTokenConfigured(v bool)`

SetRefreshTokenConfigured sets RefreshTokenConfigured field to given value.

### HasRefreshTokenConfigured

`func (o *VenafiCaRecord) HasRefreshTokenConfigured() bool`

HasRefreshTokenConfigured returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


