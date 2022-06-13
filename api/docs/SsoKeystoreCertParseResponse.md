# SsoKeystoreCertParseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** |  | [optional] [default to " "]
**Keys** | Pointer to [**[]CertificateKey**](CertificateKey.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**KeystoreSetupType** | Pointer to **string** |  | [optional] 
**KeystoreFile** | Pointer to **[]string** |  | [optional] 
**KeystoreFileName** | Pointer to **string** |  | [optional] 

## Methods

### NewSsoKeystoreCertParseResponse

`func NewSsoKeystoreCertParseResponse() *SsoKeystoreCertParseResponse`

NewSsoKeystoreCertParseResponse instantiates a new SsoKeystoreCertParseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoKeystoreCertParseResponseWithDefaults

`func NewSsoKeystoreCertParseResponseWithDefaults() *SsoKeystoreCertParseResponse`

NewSsoKeystoreCertParseResponseWithDefaults instantiates a new SsoKeystoreCertParseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *SsoKeystoreCertParseResponse) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SsoKeystoreCertParseResponse) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SsoKeystoreCertParseResponse) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *SsoKeystoreCertParseResponse) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetKeys

`func (o *SsoKeystoreCertParseResponse) GetKeys() []CertificateKey`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *SsoKeystoreCertParseResponse) GetKeysOk() (*[]CertificateKey, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *SsoKeystoreCertParseResponse) SetKeys(v []CertificateKey)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *SsoKeystoreCertParseResponse) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### GetType

`func (o *SsoKeystoreCertParseResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SsoKeystoreCertParseResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SsoKeystoreCertParseResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SsoKeystoreCertParseResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetKeystoreSetupType

`func (o *SsoKeystoreCertParseResponse) GetKeystoreSetupType() string`

GetKeystoreSetupType returns the KeystoreSetupType field if non-nil, zero value otherwise.

### GetKeystoreSetupTypeOk

`func (o *SsoKeystoreCertParseResponse) GetKeystoreSetupTypeOk() (*string, bool)`

GetKeystoreSetupTypeOk returns a tuple with the KeystoreSetupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeystoreSetupType

`func (o *SsoKeystoreCertParseResponse) SetKeystoreSetupType(v string)`

SetKeystoreSetupType sets KeystoreSetupType field to given value.

### HasKeystoreSetupType

`func (o *SsoKeystoreCertParseResponse) HasKeystoreSetupType() bool`

HasKeystoreSetupType returns a boolean if a field has been set.

### GetKeystoreFile

`func (o *SsoKeystoreCertParseResponse) GetKeystoreFile() []string`

GetKeystoreFile returns the KeystoreFile field if non-nil, zero value otherwise.

### GetKeystoreFileOk

`func (o *SsoKeystoreCertParseResponse) GetKeystoreFileOk() (*[]string, bool)`

GetKeystoreFileOk returns a tuple with the KeystoreFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeystoreFile

`func (o *SsoKeystoreCertParseResponse) SetKeystoreFile(v []string)`

SetKeystoreFile sets KeystoreFile field to given value.

### HasKeystoreFile

`func (o *SsoKeystoreCertParseResponse) HasKeystoreFile() bool`

HasKeystoreFile returns a boolean if a field has been set.

### GetKeystoreFileName

`func (o *SsoKeystoreCertParseResponse) GetKeystoreFileName() string`

GetKeystoreFileName returns the KeystoreFileName field if non-nil, zero value otherwise.

### GetKeystoreFileNameOk

`func (o *SsoKeystoreCertParseResponse) GetKeystoreFileNameOk() (*string, bool)`

GetKeystoreFileNameOk returns a tuple with the KeystoreFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeystoreFileName

`func (o *SsoKeystoreCertParseResponse) SetKeystoreFileName(v string)`

SetKeystoreFileName sets KeystoreFileName field to given value.

### HasKeystoreFileName

`func (o *SsoKeystoreCertParseResponse) HasKeystoreFileName() bool`

HasKeystoreFileName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


