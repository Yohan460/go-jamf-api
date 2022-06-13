# SsoKeystore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keys** | Pointer to [**[]CertificateKey**](CertificateKey.md) |  | [optional] 
**Key** | **string** |  | [default to " "]
**Password** | **string** |  | 
**Type** | **string** |  | 
**KeystoreSetupType** | Pointer to **string** |  | [optional] 

## Methods

### NewSsoKeystore

`func NewSsoKeystore(key string, password string, type_ string, ) *SsoKeystore`

NewSsoKeystore instantiates a new SsoKeystore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoKeystoreWithDefaults

`func NewSsoKeystoreWithDefaults() *SsoKeystore`

NewSsoKeystoreWithDefaults instantiates a new SsoKeystore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeys

`func (o *SsoKeystore) GetKeys() []CertificateKey`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *SsoKeystore) GetKeysOk() (*[]CertificateKey, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *SsoKeystore) SetKeys(v []CertificateKey)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *SsoKeystore) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### GetKey

`func (o *SsoKeystore) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SsoKeystore) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SsoKeystore) SetKey(v string)`

SetKey sets Key field to given value.


### GetPassword

`func (o *SsoKeystore) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SsoKeystore) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SsoKeystore) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetType

`func (o *SsoKeystore) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SsoKeystore) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SsoKeystore) SetType(v string)`

SetType sets Type field to given value.


### GetKeystoreSetupType

`func (o *SsoKeystore) GetKeystoreSetupType() string`

GetKeystoreSetupType returns the KeystoreSetupType field if non-nil, zero value otherwise.

### GetKeystoreSetupTypeOk

`func (o *SsoKeystore) GetKeystoreSetupTypeOk() (*string, bool)`

GetKeystoreSetupTypeOk returns a tuple with the KeystoreSetupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeystoreSetupType

`func (o *SsoKeystore) SetKeystoreSetupType(v string)`

SetKeystoreSetupType sets KeystoreSetupType field to given value.

### HasKeystoreSetupType

`func (o *SsoKeystore) HasKeystoreSetupType() bool`

HasKeystoreSetupType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


