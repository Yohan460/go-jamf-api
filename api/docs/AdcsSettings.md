# AdcsSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** |  | [optional] 
**CaName** | Pointer to **string** |  | [optional] 
**Fqdn** | Pointer to **string** |  | [optional] 
**AdcsUrl** | Pointer to **string** |  | [optional] 
**ServerCert** | Pointer to [**AdcsCertificate**](AdcsCertificate.md) |  | [optional] 
**ClientCert** | Pointer to [**AdcsCertificate**](AdcsCertificate.md) |  | [optional] 
**RevocationEnabled** | Pointer to **bool** |  | [optional] 
**ApiClientId** | Pointer to **string** |  | [optional] 
**Outbound** | Pointer to **bool** |  | [optional] 

## Methods

### NewAdcsSettings

`func NewAdcsSettings() *AdcsSettings`

NewAdcsSettings instantiates a new AdcsSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdcsSettingsWithDefaults

`func NewAdcsSettingsWithDefaults() *AdcsSettings`

NewAdcsSettingsWithDefaults instantiates a new AdcsSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *AdcsSettings) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AdcsSettings) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AdcsSettings) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AdcsSettings) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetCaName

`func (o *AdcsSettings) GetCaName() string`

GetCaName returns the CaName field if non-nil, zero value otherwise.

### GetCaNameOk

`func (o *AdcsSettings) GetCaNameOk() (*string, bool)`

GetCaNameOk returns a tuple with the CaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaName

`func (o *AdcsSettings) SetCaName(v string)`

SetCaName sets CaName field to given value.

### HasCaName

`func (o *AdcsSettings) HasCaName() bool`

HasCaName returns a boolean if a field has been set.

### GetFqdn

`func (o *AdcsSettings) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *AdcsSettings) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *AdcsSettings) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *AdcsSettings) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetAdcsUrl

`func (o *AdcsSettings) GetAdcsUrl() string`

GetAdcsUrl returns the AdcsUrl field if non-nil, zero value otherwise.

### GetAdcsUrlOk

`func (o *AdcsSettings) GetAdcsUrlOk() (*string, bool)`

GetAdcsUrlOk returns a tuple with the AdcsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdcsUrl

`func (o *AdcsSettings) SetAdcsUrl(v string)`

SetAdcsUrl sets AdcsUrl field to given value.

### HasAdcsUrl

`func (o *AdcsSettings) HasAdcsUrl() bool`

HasAdcsUrl returns a boolean if a field has been set.

### GetServerCert

`func (o *AdcsSettings) GetServerCert() AdcsCertificate`

GetServerCert returns the ServerCert field if non-nil, zero value otherwise.

### GetServerCertOk

`func (o *AdcsSettings) GetServerCertOk() (*AdcsCertificate, bool)`

GetServerCertOk returns a tuple with the ServerCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCert

`func (o *AdcsSettings) SetServerCert(v AdcsCertificate)`

SetServerCert sets ServerCert field to given value.

### HasServerCert

`func (o *AdcsSettings) HasServerCert() bool`

HasServerCert returns a boolean if a field has been set.

### GetClientCert

`func (o *AdcsSettings) GetClientCert() AdcsCertificate`

GetClientCert returns the ClientCert field if non-nil, zero value otherwise.

### GetClientCertOk

`func (o *AdcsSettings) GetClientCertOk() (*AdcsCertificate, bool)`

GetClientCertOk returns a tuple with the ClientCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCert

`func (o *AdcsSettings) SetClientCert(v AdcsCertificate)`

SetClientCert sets ClientCert field to given value.

### HasClientCert

`func (o *AdcsSettings) HasClientCert() bool`

HasClientCert returns a boolean if a field has been set.

### GetRevocationEnabled

`func (o *AdcsSettings) GetRevocationEnabled() bool`

GetRevocationEnabled returns the RevocationEnabled field if non-nil, zero value otherwise.

### GetRevocationEnabledOk

`func (o *AdcsSettings) GetRevocationEnabledOk() (*bool, bool)`

GetRevocationEnabledOk returns a tuple with the RevocationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationEnabled

`func (o *AdcsSettings) SetRevocationEnabled(v bool)`

SetRevocationEnabled sets RevocationEnabled field to given value.

### HasRevocationEnabled

`func (o *AdcsSettings) HasRevocationEnabled() bool`

HasRevocationEnabled returns a boolean if a field has been set.

### GetApiClientId

`func (o *AdcsSettings) GetApiClientId() string`

GetApiClientId returns the ApiClientId field if non-nil, zero value otherwise.

### GetApiClientIdOk

`func (o *AdcsSettings) GetApiClientIdOk() (*string, bool)`

GetApiClientIdOk returns a tuple with the ApiClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiClientId

`func (o *AdcsSettings) SetApiClientId(v string)`

SetApiClientId sets ApiClientId field to given value.

### HasApiClientId

`func (o *AdcsSettings) HasApiClientId() bool`

HasApiClientId returns a boolean if a field has been set.

### GetOutbound

`func (o *AdcsSettings) GetOutbound() bool`

GetOutbound returns the Outbound field if non-nil, zero value otherwise.

### GetOutboundOk

`func (o *AdcsSettings) GetOutboundOk() (*bool, bool)`

GetOutboundOk returns a tuple with the Outbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutbound

`func (o *AdcsSettings) SetOutbound(v bool)`

SetOutbound sets Outbound field to given value.

### HasOutbound

`func (o *AdcsSettings) HasOutbound() bool`

HasOutbound returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


