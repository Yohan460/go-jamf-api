# DigiCertSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaName** | Pointer to **string** |  | [optional] 
**Fqdn** | Pointer to **string** |  | [optional] 
**RevocationEnabled** | Pointer to **bool** |  | [optional] 
**ClientCert** | Pointer to [**Certificate**](Certificate.md) |  | [optional] 

## Methods

### NewDigiCertSetting

`func NewDigiCertSetting() *DigiCertSetting`

NewDigiCertSetting instantiates a new DigiCertSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDigiCertSettingWithDefaults

`func NewDigiCertSettingWithDefaults() *DigiCertSetting`

NewDigiCertSettingWithDefaults instantiates a new DigiCertSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaName

`func (o *DigiCertSetting) GetCaName() string`

GetCaName returns the CaName field if non-nil, zero value otherwise.

### GetCaNameOk

`func (o *DigiCertSetting) GetCaNameOk() (*string, bool)`

GetCaNameOk returns a tuple with the CaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaName

`func (o *DigiCertSetting) SetCaName(v string)`

SetCaName sets CaName field to given value.

### HasCaName

`func (o *DigiCertSetting) HasCaName() bool`

HasCaName returns a boolean if a field has been set.

### GetFqdn

`func (o *DigiCertSetting) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *DigiCertSetting) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *DigiCertSetting) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *DigiCertSetting) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetRevocationEnabled

`func (o *DigiCertSetting) GetRevocationEnabled() bool`

GetRevocationEnabled returns the RevocationEnabled field if non-nil, zero value otherwise.

### GetRevocationEnabledOk

`func (o *DigiCertSetting) GetRevocationEnabledOk() (*bool, bool)`

GetRevocationEnabledOk returns a tuple with the RevocationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationEnabled

`func (o *DigiCertSetting) SetRevocationEnabled(v bool)`

SetRevocationEnabled sets RevocationEnabled field to given value.

### HasRevocationEnabled

`func (o *DigiCertSetting) HasRevocationEnabled() bool`

HasRevocationEnabled returns a boolean if a field has been set.

### GetClientCert

`func (o *DigiCertSetting) GetClientCert() Certificate`

GetClientCert returns the ClientCert field if non-nil, zero value otherwise.

### GetClientCertOk

`func (o *DigiCertSetting) GetClientCertOk() (*Certificate, bool)`

GetClientCertOk returns a tuple with the ClientCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCert

`func (o *DigiCertSetting) SetClientCert(v Certificate)`

SetClientCert sets ClientCert field to given value.

### HasClientCert

`func (o *DigiCertSetting) HasClientCert() bool`

HasClientCert returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


