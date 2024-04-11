# CertificateRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubjectX500Principal** | Pointer to **string** |  | [optional] 
**IssuerX500Principal** | Pointer to **string** |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 
**NotAfter** | Pointer to **int64** |  | [optional] 
**NotBefore** | Pointer to **int64** |  | [optional] 
**Signature** | Pointer to [**Signature**](Signature.md) |  | [optional] 
**KeyUsage** | Pointer to **[]string** |  | [optional] 
**KeyUsageExtended** | Pointer to **[]string** |  | [optional] 
**Sha1Fingerprint** | Pointer to **string** |  | [optional] 
**Sha256Fingerprint** | Pointer to **string** |  | [optional] 

## Methods

### NewCertificateRecord

`func NewCertificateRecord() *CertificateRecord`

NewCertificateRecord instantiates a new CertificateRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateRecordWithDefaults

`func NewCertificateRecordWithDefaults() *CertificateRecord`

NewCertificateRecordWithDefaults instantiates a new CertificateRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubjectX500Principal

`func (o *CertificateRecord) GetSubjectX500Principal() string`

GetSubjectX500Principal returns the SubjectX500Principal field if non-nil, zero value otherwise.

### GetSubjectX500PrincipalOk

`func (o *CertificateRecord) GetSubjectX500PrincipalOk() (*string, bool)`

GetSubjectX500PrincipalOk returns a tuple with the SubjectX500Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectX500Principal

`func (o *CertificateRecord) SetSubjectX500Principal(v string)`

SetSubjectX500Principal sets SubjectX500Principal field to given value.

### HasSubjectX500Principal

`func (o *CertificateRecord) HasSubjectX500Principal() bool`

HasSubjectX500Principal returns a boolean if a field has been set.

### GetIssuerX500Principal

`func (o *CertificateRecord) GetIssuerX500Principal() string`

GetIssuerX500Principal returns the IssuerX500Principal field if non-nil, zero value otherwise.

### GetIssuerX500PrincipalOk

`func (o *CertificateRecord) GetIssuerX500PrincipalOk() (*string, bool)`

GetIssuerX500PrincipalOk returns a tuple with the IssuerX500Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerX500Principal

`func (o *CertificateRecord) SetIssuerX500Principal(v string)`

SetIssuerX500Principal sets IssuerX500Principal field to given value.

### HasIssuerX500Principal

`func (o *CertificateRecord) HasIssuerX500Principal() bool`

HasIssuerX500Principal returns a boolean if a field has been set.

### GetSerialNumber

`func (o *CertificateRecord) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *CertificateRecord) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *CertificateRecord) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *CertificateRecord) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetVersion

`func (o *CertificateRecord) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CertificateRecord) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CertificateRecord) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CertificateRecord) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetNotAfter

`func (o *CertificateRecord) GetNotAfter() int64`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *CertificateRecord) GetNotAfterOk() (*int64, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *CertificateRecord) SetNotAfter(v int64)`

SetNotAfter sets NotAfter field to given value.

### HasNotAfter

`func (o *CertificateRecord) HasNotAfter() bool`

HasNotAfter returns a boolean if a field has been set.

### GetNotBefore

`func (o *CertificateRecord) GetNotBefore() int64`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *CertificateRecord) GetNotBeforeOk() (*int64, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *CertificateRecord) SetNotBefore(v int64)`

SetNotBefore sets NotBefore field to given value.

### HasNotBefore

`func (o *CertificateRecord) HasNotBefore() bool`

HasNotBefore returns a boolean if a field has been set.

### GetSignature

`func (o *CertificateRecord) GetSignature() Signature`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *CertificateRecord) GetSignatureOk() (*Signature, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *CertificateRecord) SetSignature(v Signature)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *CertificateRecord) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetKeyUsage

`func (o *CertificateRecord) GetKeyUsage() []string`

GetKeyUsage returns the KeyUsage field if non-nil, zero value otherwise.

### GetKeyUsageOk

`func (o *CertificateRecord) GetKeyUsageOk() (*[]string, bool)`

GetKeyUsageOk returns a tuple with the KeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyUsage

`func (o *CertificateRecord) SetKeyUsage(v []string)`

SetKeyUsage sets KeyUsage field to given value.

### HasKeyUsage

`func (o *CertificateRecord) HasKeyUsage() bool`

HasKeyUsage returns a boolean if a field has been set.

### GetKeyUsageExtended

`func (o *CertificateRecord) GetKeyUsageExtended() []string`

GetKeyUsageExtended returns the KeyUsageExtended field if non-nil, zero value otherwise.

### GetKeyUsageExtendedOk

`func (o *CertificateRecord) GetKeyUsageExtendedOk() (*[]string, bool)`

GetKeyUsageExtendedOk returns a tuple with the KeyUsageExtended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyUsageExtended

`func (o *CertificateRecord) SetKeyUsageExtended(v []string)`

SetKeyUsageExtended sets KeyUsageExtended field to given value.

### HasKeyUsageExtended

`func (o *CertificateRecord) HasKeyUsageExtended() bool`

HasKeyUsageExtended returns a boolean if a field has been set.

### GetSha1Fingerprint

`func (o *CertificateRecord) GetSha1Fingerprint() string`

GetSha1Fingerprint returns the Sha1Fingerprint field if non-nil, zero value otherwise.

### GetSha1FingerprintOk

`func (o *CertificateRecord) GetSha1FingerprintOk() (*string, bool)`

GetSha1FingerprintOk returns a tuple with the Sha1Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha1Fingerprint

`func (o *CertificateRecord) SetSha1Fingerprint(v string)`

SetSha1Fingerprint sets Sha1Fingerprint field to given value.

### HasSha1Fingerprint

`func (o *CertificateRecord) HasSha1Fingerprint() bool`

HasSha1Fingerprint returns a boolean if a field has been set.

### GetSha256Fingerprint

`func (o *CertificateRecord) GetSha256Fingerprint() string`

GetSha256Fingerprint returns the Sha256Fingerprint field if non-nil, zero value otherwise.

### GetSha256FingerprintOk

`func (o *CertificateRecord) GetSha256FingerprintOk() (*string, bool)`

GetSha256FingerprintOk returns a tuple with the Sha256Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256Fingerprint

`func (o *CertificateRecord) SetSha256Fingerprint(v string)`

SetSha256Fingerprint sets Sha256Fingerprint field to given value.

### HasSha256Fingerprint

`func (o *CertificateRecord) HasSha256Fingerprint() bool`

HasSha256Fingerprint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


