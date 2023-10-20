# MobileDeviceCertificateV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommonName** | Pointer to **string** |  | [optional] 
**Identity** | Pointer to **bool** |  | [optional] 
**ExpirationDateEpoch** | Pointer to **time.Time** |  | [optional] 
**SubjectName** | Pointer to **string** |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**Sha1Fingerprint** | Pointer to **string** |  | [optional] 
**IssuedDateEpoch** | Pointer to **string** |  | [optional] 
**CertificateStatus** | Pointer to **string** |  | [optional] 
**LifecycleStatus** | Pointer to **string** |  | [optional] 

## Methods

### NewMobileDeviceCertificateV2

`func NewMobileDeviceCertificateV2() *MobileDeviceCertificateV2`

NewMobileDeviceCertificateV2 instantiates a new MobileDeviceCertificateV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileDeviceCertificateV2WithDefaults

`func NewMobileDeviceCertificateV2WithDefaults() *MobileDeviceCertificateV2`

NewMobileDeviceCertificateV2WithDefaults instantiates a new MobileDeviceCertificateV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommonName

`func (o *MobileDeviceCertificateV2) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *MobileDeviceCertificateV2) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *MobileDeviceCertificateV2) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *MobileDeviceCertificateV2) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### GetIdentity

`func (o *MobileDeviceCertificateV2) GetIdentity() bool`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *MobileDeviceCertificateV2) GetIdentityOk() (*bool, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *MobileDeviceCertificateV2) SetIdentity(v bool)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *MobileDeviceCertificateV2) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetExpirationDateEpoch

`func (o *MobileDeviceCertificateV2) GetExpirationDateEpoch() time.Time`

GetExpirationDateEpoch returns the ExpirationDateEpoch field if non-nil, zero value otherwise.

### GetExpirationDateEpochOk

`func (o *MobileDeviceCertificateV2) GetExpirationDateEpochOk() (*time.Time, bool)`

GetExpirationDateEpochOk returns a tuple with the ExpirationDateEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDateEpoch

`func (o *MobileDeviceCertificateV2) SetExpirationDateEpoch(v time.Time)`

SetExpirationDateEpoch sets ExpirationDateEpoch field to given value.

### HasExpirationDateEpoch

`func (o *MobileDeviceCertificateV2) HasExpirationDateEpoch() bool`

HasExpirationDateEpoch returns a boolean if a field has been set.

### GetSubjectName

`func (o *MobileDeviceCertificateV2) GetSubjectName() string`

GetSubjectName returns the SubjectName field if non-nil, zero value otherwise.

### GetSubjectNameOk

`func (o *MobileDeviceCertificateV2) GetSubjectNameOk() (*string, bool)`

GetSubjectNameOk returns a tuple with the SubjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectName

`func (o *MobileDeviceCertificateV2) SetSubjectName(v string)`

SetSubjectName sets SubjectName field to given value.

### HasSubjectName

`func (o *MobileDeviceCertificateV2) HasSubjectName() bool`

HasSubjectName returns a boolean if a field has been set.

### GetSerialNumber

`func (o *MobileDeviceCertificateV2) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *MobileDeviceCertificateV2) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *MobileDeviceCertificateV2) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *MobileDeviceCertificateV2) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetSha1Fingerprint

`func (o *MobileDeviceCertificateV2) GetSha1Fingerprint() string`

GetSha1Fingerprint returns the Sha1Fingerprint field if non-nil, zero value otherwise.

### GetSha1FingerprintOk

`func (o *MobileDeviceCertificateV2) GetSha1FingerprintOk() (*string, bool)`

GetSha1FingerprintOk returns a tuple with the Sha1Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha1Fingerprint

`func (o *MobileDeviceCertificateV2) SetSha1Fingerprint(v string)`

SetSha1Fingerprint sets Sha1Fingerprint field to given value.

### HasSha1Fingerprint

`func (o *MobileDeviceCertificateV2) HasSha1Fingerprint() bool`

HasSha1Fingerprint returns a boolean if a field has been set.

### GetIssuedDateEpoch

`func (o *MobileDeviceCertificateV2) GetIssuedDateEpoch() string`

GetIssuedDateEpoch returns the IssuedDateEpoch field if non-nil, zero value otherwise.

### GetIssuedDateEpochOk

`func (o *MobileDeviceCertificateV2) GetIssuedDateEpochOk() (*string, bool)`

GetIssuedDateEpochOk returns a tuple with the IssuedDateEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedDateEpoch

`func (o *MobileDeviceCertificateV2) SetIssuedDateEpoch(v string)`

SetIssuedDateEpoch sets IssuedDateEpoch field to given value.

### HasIssuedDateEpoch

`func (o *MobileDeviceCertificateV2) HasIssuedDateEpoch() bool`

HasIssuedDateEpoch returns a boolean if a field has been set.

### GetCertificateStatus

`func (o *MobileDeviceCertificateV2) GetCertificateStatus() string`

GetCertificateStatus returns the CertificateStatus field if non-nil, zero value otherwise.

### GetCertificateStatusOk

`func (o *MobileDeviceCertificateV2) GetCertificateStatusOk() (*string, bool)`

GetCertificateStatusOk returns a tuple with the CertificateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateStatus

`func (o *MobileDeviceCertificateV2) SetCertificateStatus(v string)`

SetCertificateStatus sets CertificateStatus field to given value.

### HasCertificateStatus

`func (o *MobileDeviceCertificateV2) HasCertificateStatus() bool`

HasCertificateStatus returns a boolean if a field has been set.

### GetLifecycleStatus

`func (o *MobileDeviceCertificateV2) GetLifecycleStatus() string`

GetLifecycleStatus returns the LifecycleStatus field if non-nil, zero value otherwise.

### GetLifecycleStatusOk

`func (o *MobileDeviceCertificateV2) GetLifecycleStatusOk() (*string, bool)`

GetLifecycleStatusOk returns a tuple with the LifecycleStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleStatus

`func (o *MobileDeviceCertificateV2) SetLifecycleStatus(v string)`

SetLifecycleStatus sets LifecycleStatus field to given value.

### HasLifecycleStatus

`func (o *MobileDeviceCertificateV2) HasLifecycleStatus() bool`

HasLifecycleStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


