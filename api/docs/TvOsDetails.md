# TvOsDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | Pointer to **string** |  | [optional] 
**ModelIdentifier** | Pointer to **string** |  | [optional] 
**ModelNumber** | Pointer to **string** |  | [optional] 
**Supervised** | Pointer to **bool** |  | [optional] 
**AirplayPassword** | Pointer to **string** |  | [optional] 
**DeviceId** | Pointer to **string** |  | [optional] 
**Locales** | Pointer to **string** |  | [optional] 
**Purchasing** | Pointer to [**PurchasingV2**](PurchasingV2.md) |  | [optional] 
**ConfigurationProfiles** | Pointer to [**[]ConfigurationProfile**](ConfigurationProfile.md) |  | [optional] 
**Certificates** | Pointer to [**[]MobileDeviceCertificateV2**](MobileDeviceCertificateV2.md) |  | [optional] 

## Methods

### NewTvOsDetails

`func NewTvOsDetails() *TvOsDetails`

NewTvOsDetails instantiates a new TvOsDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTvOsDetailsWithDefaults

`func NewTvOsDetailsWithDefaults() *TvOsDetails`

NewTvOsDetailsWithDefaults instantiates a new TvOsDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *TvOsDetails) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *TvOsDetails) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *TvOsDetails) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *TvOsDetails) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelIdentifier

`func (o *TvOsDetails) GetModelIdentifier() string`

GetModelIdentifier returns the ModelIdentifier field if non-nil, zero value otherwise.

### GetModelIdentifierOk

`func (o *TvOsDetails) GetModelIdentifierOk() (*string, bool)`

GetModelIdentifierOk returns a tuple with the ModelIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelIdentifier

`func (o *TvOsDetails) SetModelIdentifier(v string)`

SetModelIdentifier sets ModelIdentifier field to given value.

### HasModelIdentifier

`func (o *TvOsDetails) HasModelIdentifier() bool`

HasModelIdentifier returns a boolean if a field has been set.

### GetModelNumber

`func (o *TvOsDetails) GetModelNumber() string`

GetModelNumber returns the ModelNumber field if non-nil, zero value otherwise.

### GetModelNumberOk

`func (o *TvOsDetails) GetModelNumberOk() (*string, bool)`

GetModelNumberOk returns a tuple with the ModelNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelNumber

`func (o *TvOsDetails) SetModelNumber(v string)`

SetModelNumber sets ModelNumber field to given value.

### HasModelNumber

`func (o *TvOsDetails) HasModelNumber() bool`

HasModelNumber returns a boolean if a field has been set.

### GetSupervised

`func (o *TvOsDetails) GetSupervised() bool`

GetSupervised returns the Supervised field if non-nil, zero value otherwise.

### GetSupervisedOk

`func (o *TvOsDetails) GetSupervisedOk() (*bool, bool)`

GetSupervisedOk returns a tuple with the Supervised field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupervised

`func (o *TvOsDetails) SetSupervised(v bool)`

SetSupervised sets Supervised field to given value.

### HasSupervised

`func (o *TvOsDetails) HasSupervised() bool`

HasSupervised returns a boolean if a field has been set.

### GetAirplayPassword

`func (o *TvOsDetails) GetAirplayPassword() string`

GetAirplayPassword returns the AirplayPassword field if non-nil, zero value otherwise.

### GetAirplayPasswordOk

`func (o *TvOsDetails) GetAirplayPasswordOk() (*string, bool)`

GetAirplayPasswordOk returns a tuple with the AirplayPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirplayPassword

`func (o *TvOsDetails) SetAirplayPassword(v string)`

SetAirplayPassword sets AirplayPassword field to given value.

### HasAirplayPassword

`func (o *TvOsDetails) HasAirplayPassword() bool`

HasAirplayPassword returns a boolean if a field has been set.

### GetDeviceId

`func (o *TvOsDetails) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *TvOsDetails) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *TvOsDetails) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *TvOsDetails) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetLocales

`func (o *TvOsDetails) GetLocales() string`

GetLocales returns the Locales field if non-nil, zero value otherwise.

### GetLocalesOk

`func (o *TvOsDetails) GetLocalesOk() (*string, bool)`

GetLocalesOk returns a tuple with the Locales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocales

`func (o *TvOsDetails) SetLocales(v string)`

SetLocales sets Locales field to given value.

### HasLocales

`func (o *TvOsDetails) HasLocales() bool`

HasLocales returns a boolean if a field has been set.

### GetPurchasing

`func (o *TvOsDetails) GetPurchasing() PurchasingV2`

GetPurchasing returns the Purchasing field if non-nil, zero value otherwise.

### GetPurchasingOk

`func (o *TvOsDetails) GetPurchasingOk() (*PurchasingV2, bool)`

GetPurchasingOk returns a tuple with the Purchasing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasing

`func (o *TvOsDetails) SetPurchasing(v PurchasingV2)`

SetPurchasing sets Purchasing field to given value.

### HasPurchasing

`func (o *TvOsDetails) HasPurchasing() bool`

HasPurchasing returns a boolean if a field has been set.

### GetConfigurationProfiles

`func (o *TvOsDetails) GetConfigurationProfiles() []ConfigurationProfile`

GetConfigurationProfiles returns the ConfigurationProfiles field if non-nil, zero value otherwise.

### GetConfigurationProfilesOk

`func (o *TvOsDetails) GetConfigurationProfilesOk() (*[]ConfigurationProfile, bool)`

GetConfigurationProfilesOk returns a tuple with the ConfigurationProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationProfiles

`func (o *TvOsDetails) SetConfigurationProfiles(v []ConfigurationProfile)`

SetConfigurationProfiles sets ConfigurationProfiles field to given value.

### HasConfigurationProfiles

`func (o *TvOsDetails) HasConfigurationProfiles() bool`

HasConfigurationProfiles returns a boolean if a field has been set.

### GetCertificates

`func (o *TvOsDetails) GetCertificates() []MobileDeviceCertificateV2`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *TvOsDetails) GetCertificatesOk() (*[]MobileDeviceCertificateV2, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *TvOsDetails) SetCertificates(v []MobileDeviceCertificateV2)`

SetCertificates sets Certificates field to given value.

### HasCertificates

`func (o *TvOsDetails) HasCertificates() bool`

HasCertificates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


