# SsoKeystoreResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** |  | [optional] [default to " "]
**Keys** | Pointer to [**[]CertificateKey**](CertificateKey.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**KeystoreSetupType** | Pointer to **string** |  | [optional] 
**KeystoreFileName** | Pointer to **string** |  | [optional] 

## Methods

### NewSsoKeystoreResponse

`func NewSsoKeystoreResponse() *SsoKeystoreResponse`

NewSsoKeystoreResponse instantiates a new SsoKeystoreResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoKeystoreResponseWithDefaults

`func NewSsoKeystoreResponseWithDefaults() *SsoKeystoreResponse`

NewSsoKeystoreResponseWithDefaults instantiates a new SsoKeystoreResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *SsoKeystoreResponse) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SsoKeystoreResponse) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SsoKeystoreResponse) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *SsoKeystoreResponse) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetKeys

`func (o *SsoKeystoreResponse) GetKeys() []CertificateKey`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *SsoKeystoreResponse) GetKeysOk() (*[]CertificateKey, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *SsoKeystoreResponse) SetKeys(v []CertificateKey)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *SsoKeystoreResponse) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### GetType

`func (o *SsoKeystoreResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SsoKeystoreResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SsoKeystoreResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SsoKeystoreResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetKeystoreSetupType

`func (o *SsoKeystoreResponse) GetKeystoreSetupType() string`

GetKeystoreSetupType returns the KeystoreSetupType field if non-nil, zero value otherwise.

### GetKeystoreSetupTypeOk

`func (o *SsoKeystoreResponse) GetKeystoreSetupTypeOk() (*string, bool)`

GetKeystoreSetupTypeOk returns a tuple with the KeystoreSetupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeystoreSetupType

`func (o *SsoKeystoreResponse) SetKeystoreSetupType(v string)`

SetKeystoreSetupType sets KeystoreSetupType field to given value.

### HasKeystoreSetupType

`func (o *SsoKeystoreResponse) HasKeystoreSetupType() bool`

HasKeystoreSetupType returns a boolean if a field has been set.

### GetKeystoreFileName

`func (o *SsoKeystoreResponse) GetKeystoreFileName() string`

GetKeystoreFileName returns the KeystoreFileName field if non-nil, zero value otherwise.

### GetKeystoreFileNameOk

`func (o *SsoKeystoreResponse) GetKeystoreFileNameOk() (*string, bool)`

GetKeystoreFileNameOk returns a tuple with the KeystoreFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeystoreFileName

`func (o *SsoKeystoreResponse) SetKeystoreFileName(v string)`

SetKeystoreFileName sets KeystoreFileName field to given value.

### HasKeystoreFileName

`func (o *SsoKeystoreResponse) HasKeystoreFileName() bool`

HasKeystoreFileName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


