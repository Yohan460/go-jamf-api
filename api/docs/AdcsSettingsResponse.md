# AdcsSettingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**DisplayName** | Pointer to **string** |  | [optional] [readonly] 
**CaName** | Pointer to **string** |  | [optional] [readonly] 
**Fqdn** | Pointer to **string** |  | [optional] [readonly] 
**AdcsUrl** | Pointer to **string** |  | [optional] [readonly] 
**ServerCert** | Pointer to [**AdcsCertificateResponse**](AdcsCertificateResponse.md) |  | [optional] 
**ClientCert** | Pointer to [**AdcsCertificateResponse**](AdcsCertificateResponse.md) |  | [optional] 
**RevocationEnabled** | Pointer to **bool** |  | [optional] [readonly] 
**ApiClientId** | Pointer to **string** |  | [optional] [readonly] 
**Outbound** | Pointer to **bool** |  | [optional] [readonly] 
**ConnectorLastCheckInTimestamp** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewAdcsSettingsResponse

`func NewAdcsSettingsResponse() *AdcsSettingsResponse`

NewAdcsSettingsResponse instantiates a new AdcsSettingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdcsSettingsResponseWithDefaults

`func NewAdcsSettingsResponseWithDefaults() *AdcsSettingsResponse`

NewAdcsSettingsResponseWithDefaults instantiates a new AdcsSettingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AdcsSettingsResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdcsSettingsResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdcsSettingsResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AdcsSettingsResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *AdcsSettingsResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AdcsSettingsResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AdcsSettingsResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AdcsSettingsResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetCaName

`func (o *AdcsSettingsResponse) GetCaName() string`

GetCaName returns the CaName field if non-nil, zero value otherwise.

### GetCaNameOk

`func (o *AdcsSettingsResponse) GetCaNameOk() (*string, bool)`

GetCaNameOk returns a tuple with the CaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaName

`func (o *AdcsSettingsResponse) SetCaName(v string)`

SetCaName sets CaName field to given value.

### HasCaName

`func (o *AdcsSettingsResponse) HasCaName() bool`

HasCaName returns a boolean if a field has been set.

### GetFqdn

`func (o *AdcsSettingsResponse) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *AdcsSettingsResponse) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *AdcsSettingsResponse) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *AdcsSettingsResponse) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetAdcsUrl

`func (o *AdcsSettingsResponse) GetAdcsUrl() string`

GetAdcsUrl returns the AdcsUrl field if non-nil, zero value otherwise.

### GetAdcsUrlOk

`func (o *AdcsSettingsResponse) GetAdcsUrlOk() (*string, bool)`

GetAdcsUrlOk returns a tuple with the AdcsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdcsUrl

`func (o *AdcsSettingsResponse) SetAdcsUrl(v string)`

SetAdcsUrl sets AdcsUrl field to given value.

### HasAdcsUrl

`func (o *AdcsSettingsResponse) HasAdcsUrl() bool`

HasAdcsUrl returns a boolean if a field has been set.

### GetServerCert

`func (o *AdcsSettingsResponse) GetServerCert() AdcsCertificateResponse`

GetServerCert returns the ServerCert field if non-nil, zero value otherwise.

### GetServerCertOk

`func (o *AdcsSettingsResponse) GetServerCertOk() (*AdcsCertificateResponse, bool)`

GetServerCertOk returns a tuple with the ServerCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCert

`func (o *AdcsSettingsResponse) SetServerCert(v AdcsCertificateResponse)`

SetServerCert sets ServerCert field to given value.

### HasServerCert

`func (o *AdcsSettingsResponse) HasServerCert() bool`

HasServerCert returns a boolean if a field has been set.

### GetClientCert

`func (o *AdcsSettingsResponse) GetClientCert() AdcsCertificateResponse`

GetClientCert returns the ClientCert field if non-nil, zero value otherwise.

### GetClientCertOk

`func (o *AdcsSettingsResponse) GetClientCertOk() (*AdcsCertificateResponse, bool)`

GetClientCertOk returns a tuple with the ClientCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCert

`func (o *AdcsSettingsResponse) SetClientCert(v AdcsCertificateResponse)`

SetClientCert sets ClientCert field to given value.

### HasClientCert

`func (o *AdcsSettingsResponse) HasClientCert() bool`

HasClientCert returns a boolean if a field has been set.

### GetRevocationEnabled

`func (o *AdcsSettingsResponse) GetRevocationEnabled() bool`

GetRevocationEnabled returns the RevocationEnabled field if non-nil, zero value otherwise.

### GetRevocationEnabledOk

`func (o *AdcsSettingsResponse) GetRevocationEnabledOk() (*bool, bool)`

GetRevocationEnabledOk returns a tuple with the RevocationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationEnabled

`func (o *AdcsSettingsResponse) SetRevocationEnabled(v bool)`

SetRevocationEnabled sets RevocationEnabled field to given value.

### HasRevocationEnabled

`func (o *AdcsSettingsResponse) HasRevocationEnabled() bool`

HasRevocationEnabled returns a boolean if a field has been set.

### GetApiClientId

`func (o *AdcsSettingsResponse) GetApiClientId() string`

GetApiClientId returns the ApiClientId field if non-nil, zero value otherwise.

### GetApiClientIdOk

`func (o *AdcsSettingsResponse) GetApiClientIdOk() (*string, bool)`

GetApiClientIdOk returns a tuple with the ApiClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiClientId

`func (o *AdcsSettingsResponse) SetApiClientId(v string)`

SetApiClientId sets ApiClientId field to given value.

### HasApiClientId

`func (o *AdcsSettingsResponse) HasApiClientId() bool`

HasApiClientId returns a boolean if a field has been set.

### GetOutbound

`func (o *AdcsSettingsResponse) GetOutbound() bool`

GetOutbound returns the Outbound field if non-nil, zero value otherwise.

### GetOutboundOk

`func (o *AdcsSettingsResponse) GetOutboundOk() (*bool, bool)`

GetOutboundOk returns a tuple with the Outbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutbound

`func (o *AdcsSettingsResponse) SetOutbound(v bool)`

SetOutbound sets Outbound field to given value.

### HasOutbound

`func (o *AdcsSettingsResponse) HasOutbound() bool`

HasOutbound returns a boolean if a field has been set.

### GetConnectorLastCheckInTimestamp

`func (o *AdcsSettingsResponse) GetConnectorLastCheckInTimestamp() time.Time`

GetConnectorLastCheckInTimestamp returns the ConnectorLastCheckInTimestamp field if non-nil, zero value otherwise.

### GetConnectorLastCheckInTimestampOk

`func (o *AdcsSettingsResponse) GetConnectorLastCheckInTimestampOk() (*time.Time, bool)`

GetConnectorLastCheckInTimestampOk returns a tuple with the ConnectorLastCheckInTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorLastCheckInTimestamp

`func (o *AdcsSettingsResponse) SetConnectorLastCheckInTimestamp(v time.Time)`

SetConnectorLastCheckInTimestamp sets ConnectorLastCheckInTimestamp field to given value.

### HasConnectorLastCheckInTimestamp

`func (o *AdcsSettingsResponse) HasConnectorLastCheckInTimestamp() bool`

HasConnectorLastCheckInTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


