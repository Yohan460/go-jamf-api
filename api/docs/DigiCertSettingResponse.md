# DigiCertSettingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**CaName** | Pointer to **string** |  | [optional] 
**Fqdn** | Pointer to **string** |  | [optional] 
**RevocationEnabled** | Pointer to **bool** |  | [optional] 
**ClientCert** | Pointer to [**CertificateResponse**](CertificateResponse.md) |  | [optional] 

## Methods

### NewDigiCertSettingResponse

`func NewDigiCertSettingResponse() *DigiCertSettingResponse`

NewDigiCertSettingResponse instantiates a new DigiCertSettingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDigiCertSettingResponseWithDefaults

`func NewDigiCertSettingResponseWithDefaults() *DigiCertSettingResponse`

NewDigiCertSettingResponseWithDefaults instantiates a new DigiCertSettingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DigiCertSettingResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DigiCertSettingResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DigiCertSettingResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DigiCertSettingResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCaName

`func (o *DigiCertSettingResponse) GetCaName() string`

GetCaName returns the CaName field if non-nil, zero value otherwise.

### GetCaNameOk

`func (o *DigiCertSettingResponse) GetCaNameOk() (*string, bool)`

GetCaNameOk returns a tuple with the CaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaName

`func (o *DigiCertSettingResponse) SetCaName(v string)`

SetCaName sets CaName field to given value.

### HasCaName

`func (o *DigiCertSettingResponse) HasCaName() bool`

HasCaName returns a boolean if a field has been set.

### GetFqdn

`func (o *DigiCertSettingResponse) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *DigiCertSettingResponse) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *DigiCertSettingResponse) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *DigiCertSettingResponse) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetRevocationEnabled

`func (o *DigiCertSettingResponse) GetRevocationEnabled() bool`

GetRevocationEnabled returns the RevocationEnabled field if non-nil, zero value otherwise.

### GetRevocationEnabledOk

`func (o *DigiCertSettingResponse) GetRevocationEnabledOk() (*bool, bool)`

GetRevocationEnabledOk returns a tuple with the RevocationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationEnabled

`func (o *DigiCertSettingResponse) SetRevocationEnabled(v bool)`

SetRevocationEnabled sets RevocationEnabled field to given value.

### HasRevocationEnabled

`func (o *DigiCertSettingResponse) HasRevocationEnabled() bool`

HasRevocationEnabled returns a boolean if a field has been set.

### GetClientCert

`func (o *DigiCertSettingResponse) GetClientCert() CertificateResponse`

GetClientCert returns the ClientCert field if non-nil, zero value otherwise.

### GetClientCertOk

`func (o *DigiCertSettingResponse) GetClientCertOk() (*CertificateResponse, bool)`

GetClientCertOk returns a tuple with the ClientCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCert

`func (o *DigiCertSettingResponse) SetClientCert(v CertificateResponse)`

SetClientCert sets ClientCert field to given value.

### HasClientCert

`func (o *DigiCertSettingResponse) HasClientCert() bool`

HasClientCert returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


