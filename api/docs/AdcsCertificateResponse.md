# AdcsCertificateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filename** | Pointer to **string** | Server certificate filename should extend .cer or .pem, and client certificate filename should extend .p12 or .pfx. | [optional] [readonly] 
**SerialNumber** | Pointer to **string** |  | [optional] [readonly] 
**Subject** | Pointer to **string** |  | [optional] [readonly] 
**Issuer** | Pointer to **string** |  | [optional] [readonly] 
**ExpirationDate** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewAdcsCertificateResponse

`func NewAdcsCertificateResponse() *AdcsCertificateResponse`

NewAdcsCertificateResponse instantiates a new AdcsCertificateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdcsCertificateResponseWithDefaults

`func NewAdcsCertificateResponseWithDefaults() *AdcsCertificateResponse`

NewAdcsCertificateResponseWithDefaults instantiates a new AdcsCertificateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilename

`func (o *AdcsCertificateResponse) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *AdcsCertificateResponse) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *AdcsCertificateResponse) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *AdcsCertificateResponse) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetSerialNumber

`func (o *AdcsCertificateResponse) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *AdcsCertificateResponse) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *AdcsCertificateResponse) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *AdcsCertificateResponse) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetSubject

`func (o *AdcsCertificateResponse) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *AdcsCertificateResponse) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *AdcsCertificateResponse) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *AdcsCertificateResponse) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetIssuer

`func (o *AdcsCertificateResponse) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *AdcsCertificateResponse) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *AdcsCertificateResponse) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *AdcsCertificateResponse) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetExpirationDate

`func (o *AdcsCertificateResponse) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *AdcsCertificateResponse) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *AdcsCertificateResponse) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *AdcsCertificateResponse) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


