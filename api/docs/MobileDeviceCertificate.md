# MobileDeviceCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommonName** | Pointer to **string** |  | [optional] 
**Identity** | Pointer to **bool** |  | [optional] 
**ExpirationDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewMobileDeviceCertificate

`func NewMobileDeviceCertificate() *MobileDeviceCertificate`

NewMobileDeviceCertificate instantiates a new MobileDeviceCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileDeviceCertificateWithDefaults

`func NewMobileDeviceCertificateWithDefaults() *MobileDeviceCertificate`

NewMobileDeviceCertificateWithDefaults instantiates a new MobileDeviceCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommonName

`func (o *MobileDeviceCertificate) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *MobileDeviceCertificate) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *MobileDeviceCertificate) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *MobileDeviceCertificate) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### GetIdentity

`func (o *MobileDeviceCertificate) GetIdentity() bool`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *MobileDeviceCertificate) GetIdentityOk() (*bool, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *MobileDeviceCertificate) SetIdentity(v bool)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *MobileDeviceCertificate) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetExpirationDate

`func (o *MobileDeviceCertificate) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *MobileDeviceCertificate) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *MobileDeviceCertificate) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *MobileDeviceCertificate) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


