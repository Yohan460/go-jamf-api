# SsoKeystoreAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keys** | Pointer to [**[]CertificateKey**](CertificateKey.md) |  | [optional] 
**Key** | **string** |  | [default to " "]
**Password** | **string** |  | 
**Type** | **string** |  | 
**KeystoreSetupType** | Pointer to **string** |  | [optional] 

## Methods

### NewSsoKeystoreAllOf

`func NewSsoKeystoreAllOf(key string, password string, type_ string, ) *SsoKeystoreAllOf`

NewSsoKeystoreAllOf instantiates a new SsoKeystoreAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoKeystoreAllOfWithDefaults

`func NewSsoKeystoreAllOfWithDefaults() *SsoKeystoreAllOf`

NewSsoKeystoreAllOfWithDefaults instantiates a new SsoKeystoreAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeys

`func (o *SsoKeystoreAllOf) GetKeys() []CertificateKey`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *SsoKeystoreAllOf) GetKeysOk() (*[]CertificateKey, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *SsoKeystoreAllOf) SetKeys(v []CertificateKey)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *SsoKeystoreAllOf) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### GetKey

`func (o *SsoKeystoreAllOf) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SsoKeystoreAllOf) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SsoKeystoreAllOf) SetKey(v string)`

SetKey sets Key field to given value.


### GetPassword

`func (o *SsoKeystoreAllOf) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SsoKeystoreAllOf) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SsoKeystoreAllOf) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetType

`func (o *SsoKeystoreAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SsoKeystoreAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SsoKeystoreAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetKeystoreSetupType

`func (o *SsoKeystoreAllOf) GetKeystoreSetupType() string`

GetKeystoreSetupType returns the KeystoreSetupType field if non-nil, zero value otherwise.

### GetKeystoreSetupTypeOk

`func (o *SsoKeystoreAllOf) GetKeystoreSetupTypeOk() (*string, bool)`

GetKeystoreSetupTypeOk returns a tuple with the KeystoreSetupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeystoreSetupType

`func (o *SsoKeystoreAllOf) SetKeystoreSetupType(v string)`

SetKeystoreSetupType sets KeystoreSetupType field to given value.

### HasKeystoreSetupType

`func (o *SsoKeystoreAllOf) HasKeystoreSetupType() bool`

HasKeystoreSetupType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


