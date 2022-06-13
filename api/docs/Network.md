# Network

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CellularTechnology** | Pointer to **string** |  | [optional] 
**IsVoiceRoamingEnabled** | Pointer to **bool** |  | [optional] 
**Imei** | Pointer to **string** |  | [optional] 
**Iccid** | Pointer to **string** |  | [optional] 
**Meid** | Pointer to **string** |  | [optional] 
**CarrierSettingsVersion** | Pointer to **string** |  | [optional] 
**CurrentCarrierNetwork** | Pointer to **string** |  | [optional] 
**CurrentMobileCountryCode** | Pointer to **string** |  | [optional] 
**CurrentMobileNetworkCode** | Pointer to **string** |  | [optional] 
**HomeCarrierNetwork** | Pointer to **string** |  | [optional] 
**HomeMobileCountryCode** | Pointer to **string** |  | [optional] 
**HomeMobileNetworkCode** | Pointer to **string** |  | [optional] 
**IsDataRoamingEnabled** | Pointer to **bool** |  | [optional] 
**IsRoaming** | Pointer to **bool** |  | [optional] 
**IsPersonalHotspotEnabled** | Pointer to **bool** |  | [optional] 
**PhoneNumber** | Pointer to **string** |  | [optional] 

## Methods

### NewNetwork

`func NewNetwork() *Network`

NewNetwork instantiates a new Network object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkWithDefaults

`func NewNetworkWithDefaults() *Network`

NewNetworkWithDefaults instantiates a new Network object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCellularTechnology

`func (o *Network) GetCellularTechnology() string`

GetCellularTechnology returns the CellularTechnology field if non-nil, zero value otherwise.

### GetCellularTechnologyOk

`func (o *Network) GetCellularTechnologyOk() (*string, bool)`

GetCellularTechnologyOk returns a tuple with the CellularTechnology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellularTechnology

`func (o *Network) SetCellularTechnology(v string)`

SetCellularTechnology sets CellularTechnology field to given value.

### HasCellularTechnology

`func (o *Network) HasCellularTechnology() bool`

HasCellularTechnology returns a boolean if a field has been set.

### GetIsVoiceRoamingEnabled

`func (o *Network) GetIsVoiceRoamingEnabled() bool`

GetIsVoiceRoamingEnabled returns the IsVoiceRoamingEnabled field if non-nil, zero value otherwise.

### GetIsVoiceRoamingEnabledOk

`func (o *Network) GetIsVoiceRoamingEnabledOk() (*bool, bool)`

GetIsVoiceRoamingEnabledOk returns a tuple with the IsVoiceRoamingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVoiceRoamingEnabled

`func (o *Network) SetIsVoiceRoamingEnabled(v bool)`

SetIsVoiceRoamingEnabled sets IsVoiceRoamingEnabled field to given value.

### HasIsVoiceRoamingEnabled

`func (o *Network) HasIsVoiceRoamingEnabled() bool`

HasIsVoiceRoamingEnabled returns a boolean if a field has been set.

### GetImei

`func (o *Network) GetImei() string`

GetImei returns the Imei field if non-nil, zero value otherwise.

### GetImeiOk

`func (o *Network) GetImeiOk() (*string, bool)`

GetImeiOk returns a tuple with the Imei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImei

`func (o *Network) SetImei(v string)`

SetImei sets Imei field to given value.

### HasImei

`func (o *Network) HasImei() bool`

HasImei returns a boolean if a field has been set.

### GetIccid

`func (o *Network) GetIccid() string`

GetIccid returns the Iccid field if non-nil, zero value otherwise.

### GetIccidOk

`func (o *Network) GetIccidOk() (*string, bool)`

GetIccidOk returns a tuple with the Iccid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIccid

`func (o *Network) SetIccid(v string)`

SetIccid sets Iccid field to given value.

### HasIccid

`func (o *Network) HasIccid() bool`

HasIccid returns a boolean if a field has been set.

### GetMeid

`func (o *Network) GetMeid() string`

GetMeid returns the Meid field if non-nil, zero value otherwise.

### GetMeidOk

`func (o *Network) GetMeidOk() (*string, bool)`

GetMeidOk returns a tuple with the Meid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeid

`func (o *Network) SetMeid(v string)`

SetMeid sets Meid field to given value.

### HasMeid

`func (o *Network) HasMeid() bool`

HasMeid returns a boolean if a field has been set.

### GetCarrierSettingsVersion

`func (o *Network) GetCarrierSettingsVersion() string`

GetCarrierSettingsVersion returns the CarrierSettingsVersion field if non-nil, zero value otherwise.

### GetCarrierSettingsVersionOk

`func (o *Network) GetCarrierSettingsVersionOk() (*string, bool)`

GetCarrierSettingsVersionOk returns a tuple with the CarrierSettingsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierSettingsVersion

`func (o *Network) SetCarrierSettingsVersion(v string)`

SetCarrierSettingsVersion sets CarrierSettingsVersion field to given value.

### HasCarrierSettingsVersion

`func (o *Network) HasCarrierSettingsVersion() bool`

HasCarrierSettingsVersion returns a boolean if a field has been set.

### GetCurrentCarrierNetwork

`func (o *Network) GetCurrentCarrierNetwork() string`

GetCurrentCarrierNetwork returns the CurrentCarrierNetwork field if non-nil, zero value otherwise.

### GetCurrentCarrierNetworkOk

`func (o *Network) GetCurrentCarrierNetworkOk() (*string, bool)`

GetCurrentCarrierNetworkOk returns a tuple with the CurrentCarrierNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentCarrierNetwork

`func (o *Network) SetCurrentCarrierNetwork(v string)`

SetCurrentCarrierNetwork sets CurrentCarrierNetwork field to given value.

### HasCurrentCarrierNetwork

`func (o *Network) HasCurrentCarrierNetwork() bool`

HasCurrentCarrierNetwork returns a boolean if a field has been set.

### GetCurrentMobileCountryCode

`func (o *Network) GetCurrentMobileCountryCode() string`

GetCurrentMobileCountryCode returns the CurrentMobileCountryCode field if non-nil, zero value otherwise.

### GetCurrentMobileCountryCodeOk

`func (o *Network) GetCurrentMobileCountryCodeOk() (*string, bool)`

GetCurrentMobileCountryCodeOk returns a tuple with the CurrentMobileCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentMobileCountryCode

`func (o *Network) SetCurrentMobileCountryCode(v string)`

SetCurrentMobileCountryCode sets CurrentMobileCountryCode field to given value.

### HasCurrentMobileCountryCode

`func (o *Network) HasCurrentMobileCountryCode() bool`

HasCurrentMobileCountryCode returns a boolean if a field has been set.

### GetCurrentMobileNetworkCode

`func (o *Network) GetCurrentMobileNetworkCode() string`

GetCurrentMobileNetworkCode returns the CurrentMobileNetworkCode field if non-nil, zero value otherwise.

### GetCurrentMobileNetworkCodeOk

`func (o *Network) GetCurrentMobileNetworkCodeOk() (*string, bool)`

GetCurrentMobileNetworkCodeOk returns a tuple with the CurrentMobileNetworkCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentMobileNetworkCode

`func (o *Network) SetCurrentMobileNetworkCode(v string)`

SetCurrentMobileNetworkCode sets CurrentMobileNetworkCode field to given value.

### HasCurrentMobileNetworkCode

`func (o *Network) HasCurrentMobileNetworkCode() bool`

HasCurrentMobileNetworkCode returns a boolean if a field has been set.

### GetHomeCarrierNetwork

`func (o *Network) GetHomeCarrierNetwork() string`

GetHomeCarrierNetwork returns the HomeCarrierNetwork field if non-nil, zero value otherwise.

### GetHomeCarrierNetworkOk

`func (o *Network) GetHomeCarrierNetworkOk() (*string, bool)`

GetHomeCarrierNetworkOk returns a tuple with the HomeCarrierNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeCarrierNetwork

`func (o *Network) SetHomeCarrierNetwork(v string)`

SetHomeCarrierNetwork sets HomeCarrierNetwork field to given value.

### HasHomeCarrierNetwork

`func (o *Network) HasHomeCarrierNetwork() bool`

HasHomeCarrierNetwork returns a boolean if a field has been set.

### GetHomeMobileCountryCode

`func (o *Network) GetHomeMobileCountryCode() string`

GetHomeMobileCountryCode returns the HomeMobileCountryCode field if non-nil, zero value otherwise.

### GetHomeMobileCountryCodeOk

`func (o *Network) GetHomeMobileCountryCodeOk() (*string, bool)`

GetHomeMobileCountryCodeOk returns a tuple with the HomeMobileCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeMobileCountryCode

`func (o *Network) SetHomeMobileCountryCode(v string)`

SetHomeMobileCountryCode sets HomeMobileCountryCode field to given value.

### HasHomeMobileCountryCode

`func (o *Network) HasHomeMobileCountryCode() bool`

HasHomeMobileCountryCode returns a boolean if a field has been set.

### GetHomeMobileNetworkCode

`func (o *Network) GetHomeMobileNetworkCode() string`

GetHomeMobileNetworkCode returns the HomeMobileNetworkCode field if non-nil, zero value otherwise.

### GetHomeMobileNetworkCodeOk

`func (o *Network) GetHomeMobileNetworkCodeOk() (*string, bool)`

GetHomeMobileNetworkCodeOk returns a tuple with the HomeMobileNetworkCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeMobileNetworkCode

`func (o *Network) SetHomeMobileNetworkCode(v string)`

SetHomeMobileNetworkCode sets HomeMobileNetworkCode field to given value.

### HasHomeMobileNetworkCode

`func (o *Network) HasHomeMobileNetworkCode() bool`

HasHomeMobileNetworkCode returns a boolean if a field has been set.

### GetIsDataRoamingEnabled

`func (o *Network) GetIsDataRoamingEnabled() bool`

GetIsDataRoamingEnabled returns the IsDataRoamingEnabled field if non-nil, zero value otherwise.

### GetIsDataRoamingEnabledOk

`func (o *Network) GetIsDataRoamingEnabledOk() (*bool, bool)`

GetIsDataRoamingEnabledOk returns a tuple with the IsDataRoamingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDataRoamingEnabled

`func (o *Network) SetIsDataRoamingEnabled(v bool)`

SetIsDataRoamingEnabled sets IsDataRoamingEnabled field to given value.

### HasIsDataRoamingEnabled

`func (o *Network) HasIsDataRoamingEnabled() bool`

HasIsDataRoamingEnabled returns a boolean if a field has been set.

### GetIsRoaming

`func (o *Network) GetIsRoaming() bool`

GetIsRoaming returns the IsRoaming field if non-nil, zero value otherwise.

### GetIsRoamingOk

`func (o *Network) GetIsRoamingOk() (*bool, bool)`

GetIsRoamingOk returns a tuple with the IsRoaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRoaming

`func (o *Network) SetIsRoaming(v bool)`

SetIsRoaming sets IsRoaming field to given value.

### HasIsRoaming

`func (o *Network) HasIsRoaming() bool`

HasIsRoaming returns a boolean if a field has been set.

### GetIsPersonalHotspotEnabled

`func (o *Network) GetIsPersonalHotspotEnabled() bool`

GetIsPersonalHotspotEnabled returns the IsPersonalHotspotEnabled field if non-nil, zero value otherwise.

### GetIsPersonalHotspotEnabledOk

`func (o *Network) GetIsPersonalHotspotEnabledOk() (*bool, bool)`

GetIsPersonalHotspotEnabledOk returns a tuple with the IsPersonalHotspotEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPersonalHotspotEnabled

`func (o *Network) SetIsPersonalHotspotEnabled(v bool)`

SetIsPersonalHotspotEnabled sets IsPersonalHotspotEnabled field to given value.

### HasIsPersonalHotspotEnabled

`func (o *Network) HasIsPersonalHotspotEnabled() bool`

HasIsPersonalHotspotEnabled returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *Network) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *Network) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *Network) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *Network) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


