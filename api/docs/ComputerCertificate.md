# ComputerCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommonName** | Pointer to **string** |  | [optional] 
**Identity** | Pointer to **bool** |  | [optional] 
**ExpirationDate** | Pointer to **time.Time** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**LifecycleStatus** | Pointer to **string** |  | [optional] 
**CertificateStatus** | Pointer to **string** |  | [optional] 
**SubjectName** | Pointer to **string** |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**Sha1Fingerprint** | Pointer to **string** |  | [optional] 
**IssuedDate** | Pointer to **string** |  | [optional] 

## Methods

### NewComputerCertificate

`func NewComputerCertificate() *ComputerCertificate`

NewComputerCertificate instantiates a new ComputerCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputerCertificateWithDefaults

`func NewComputerCertificateWithDefaults() *ComputerCertificate`

NewComputerCertificateWithDefaults instantiates a new ComputerCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommonName

`func (o *ComputerCertificate) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *ComputerCertificate) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *ComputerCertificate) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *ComputerCertificate) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### GetIdentity

`func (o *ComputerCertificate) GetIdentity() bool`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *ComputerCertificate) GetIdentityOk() (*bool, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *ComputerCertificate) SetIdentity(v bool)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *ComputerCertificate) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetExpirationDate

`func (o *ComputerCertificate) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *ComputerCertificate) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *ComputerCertificate) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *ComputerCertificate) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetUsername

`func (o *ComputerCertificate) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ComputerCertificate) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ComputerCertificate) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ComputerCertificate) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetLifecycleStatus

`func (o *ComputerCertificate) GetLifecycleStatus() string`

GetLifecycleStatus returns the LifecycleStatus field if non-nil, zero value otherwise.

### GetLifecycleStatusOk

`func (o *ComputerCertificate) GetLifecycleStatusOk() (*string, bool)`

GetLifecycleStatusOk returns a tuple with the LifecycleStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleStatus

`func (o *ComputerCertificate) SetLifecycleStatus(v string)`

SetLifecycleStatus sets LifecycleStatus field to given value.

### HasLifecycleStatus

`func (o *ComputerCertificate) HasLifecycleStatus() bool`

HasLifecycleStatus returns a boolean if a field has been set.

### GetCertificateStatus

`func (o *ComputerCertificate) GetCertificateStatus() string`

GetCertificateStatus returns the CertificateStatus field if non-nil, zero value otherwise.

### GetCertificateStatusOk

`func (o *ComputerCertificate) GetCertificateStatusOk() (*string, bool)`

GetCertificateStatusOk returns a tuple with the CertificateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateStatus

`func (o *ComputerCertificate) SetCertificateStatus(v string)`

SetCertificateStatus sets CertificateStatus field to given value.

### HasCertificateStatus

`func (o *ComputerCertificate) HasCertificateStatus() bool`

HasCertificateStatus returns a boolean if a field has been set.

### GetSubjectName

`func (o *ComputerCertificate) GetSubjectName() string`

GetSubjectName returns the SubjectName field if non-nil, zero value otherwise.

### GetSubjectNameOk

`func (o *ComputerCertificate) GetSubjectNameOk() (*string, bool)`

GetSubjectNameOk returns a tuple with the SubjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectName

`func (o *ComputerCertificate) SetSubjectName(v string)`

SetSubjectName sets SubjectName field to given value.

### HasSubjectName

`func (o *ComputerCertificate) HasSubjectName() bool`

HasSubjectName returns a boolean if a field has been set.

### GetSerialNumber

`func (o *ComputerCertificate) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *ComputerCertificate) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *ComputerCertificate) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *ComputerCertificate) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetSha1Fingerprint

`func (o *ComputerCertificate) GetSha1Fingerprint() string`

GetSha1Fingerprint returns the Sha1Fingerprint field if non-nil, zero value otherwise.

### GetSha1FingerprintOk

`func (o *ComputerCertificate) GetSha1FingerprintOk() (*string, bool)`

GetSha1FingerprintOk returns a tuple with the Sha1Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha1Fingerprint

`func (o *ComputerCertificate) SetSha1Fingerprint(v string)`

SetSha1Fingerprint sets Sha1Fingerprint field to given value.

### HasSha1Fingerprint

`func (o *ComputerCertificate) HasSha1Fingerprint() bool`

HasSha1Fingerprint returns a boolean if a field has been set.

### GetIssuedDate

`func (o *ComputerCertificate) GetIssuedDate() string`

GetIssuedDate returns the IssuedDate field if non-nil, zero value otherwise.

### GetIssuedDateOk

`func (o *ComputerCertificate) GetIssuedDateOk() (*string, bool)`

GetIssuedDateOk returns a tuple with the IssuedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedDate

`func (o *ComputerCertificate) SetIssuedDate(v string)`

SetIssuedDate sets IssuedDate field to given value.

### HasIssuedDate

`func (o *ComputerCertificate) HasIssuedDate() bool`

HasIssuedDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


