# CertificateIdentityV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filename** | Pointer to **string** |  | [optional] [default to "null"]
**KeystorePassword** | Pointer to **string** |  | [optional] [default to ""]
**IdentityKeystore** | Pointer to **string** | The base 64 encoded certificate. Will not be returned on a GET. | [optional] 
**Md5Sum** | Pointer to **string** | The md5 checksum of the certificate file. Intended to be used in verifification the cert being used to sign QuickAdd packages. | [optional] [default to ""]

## Methods

### NewCertificateIdentityV1

`func NewCertificateIdentityV1() *CertificateIdentityV1`

NewCertificateIdentityV1 instantiates a new CertificateIdentityV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateIdentityV1WithDefaults

`func NewCertificateIdentityV1WithDefaults() *CertificateIdentityV1`

NewCertificateIdentityV1WithDefaults instantiates a new CertificateIdentityV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilename

`func (o *CertificateIdentityV1) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *CertificateIdentityV1) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *CertificateIdentityV1) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *CertificateIdentityV1) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetKeystorePassword

`func (o *CertificateIdentityV1) GetKeystorePassword() string`

GetKeystorePassword returns the KeystorePassword field if non-nil, zero value otherwise.

### GetKeystorePasswordOk

`func (o *CertificateIdentityV1) GetKeystorePasswordOk() (*string, bool)`

GetKeystorePasswordOk returns a tuple with the KeystorePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeystorePassword

`func (o *CertificateIdentityV1) SetKeystorePassword(v string)`

SetKeystorePassword sets KeystorePassword field to given value.

### HasKeystorePassword

`func (o *CertificateIdentityV1) HasKeystorePassword() bool`

HasKeystorePassword returns a boolean if a field has been set.

### GetIdentityKeystore

`func (o *CertificateIdentityV1) GetIdentityKeystore() string`

GetIdentityKeystore returns the IdentityKeystore field if non-nil, zero value otherwise.

### GetIdentityKeystoreOk

`func (o *CertificateIdentityV1) GetIdentityKeystoreOk() (*string, bool)`

GetIdentityKeystoreOk returns a tuple with the IdentityKeystore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityKeystore

`func (o *CertificateIdentityV1) SetIdentityKeystore(v string)`

SetIdentityKeystore sets IdentityKeystore field to given value.

### HasIdentityKeystore

`func (o *CertificateIdentityV1) HasIdentityKeystore() bool`

HasIdentityKeystore returns a boolean if a field has been set.

### GetMd5Sum

`func (o *CertificateIdentityV1) GetMd5Sum() string`

GetMd5Sum returns the Md5Sum field if non-nil, zero value otherwise.

### GetMd5SumOk

`func (o *CertificateIdentityV1) GetMd5SumOk() (*string, bool)`

GetMd5SumOk returns a tuple with the Md5Sum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5Sum

`func (o *CertificateIdentityV1) SetMd5Sum(v string)`

SetMd5Sum sets Md5Sum field to given value.

### HasMd5Sum

`func (o *CertificateIdentityV1) HasMd5Sum() bool`

HasMd5Sum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


