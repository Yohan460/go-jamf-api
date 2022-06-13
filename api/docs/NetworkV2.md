# NetworkV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CellularTechnology** | Pointer to **string** |  | [optional] 
**VoiceRoamingEnabled** | Pointer to **bool** |  | [optional] 
**Imei** | Pointer to **string** |  | [optional] 
**Iccid** | Pointer to **string** |  | [optional] 
**Meid** | Pointer to **string** |  | [optional] 
**Eid** | Pointer to **string** | EID or \&quot;embedded identity document\&quot; is a number associated with the eSIM on a device | [optional] [readonly] 
**CarrierSettingsVersion** | Pointer to **string** |  | [optional] 
**CurrentCarrierNetwork** | Pointer to **string** |  | [optional] 
**CurrentMobileCountryCode** | Pointer to **string** |  | [optional] 
**CurrentMobileNetworkCode** | Pointer to **string** |  | [optional] 
**HomeCarrierNetwork** | Pointer to **string** |  | [optional] 
**HomeMobileCountryCode** | Pointer to **string** |  | [optional] 
**HomeMobileNetworkCode** | Pointer to **string** |  | [optional] 
**DataRoamingEnabled** | Pointer to **bool** |  | [optional] 
**Roaming** | Pointer to **bool** |  | [optional] 
**PersonalHotspotEnabled** | Pointer to **bool** |  | [optional] 
**PhoneNumber** | Pointer to **string** |  | [optional] 

## Methods

### NewNetworkV2

`func NewNetworkV2() *NetworkV2`

NewNetworkV2 instantiates a new NetworkV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkV2WithDefaults

`func NewNetworkV2WithDefaults() *NetworkV2`

NewNetworkV2WithDefaults instantiates a new NetworkV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCellularTechnology

`func (o *NetworkV2) GetCellularTechnology() string`

GetCellularTechnology returns the CellularTechnology field if non-nil, zero value otherwise.

### GetCellularTechnologyOk

`func (o *NetworkV2) GetCellularTechnologyOk() (*string, bool)`

GetCellularTechnologyOk returns a tuple with the CellularTechnology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellularTechnology

`func (o *NetworkV2) SetCellularTechnology(v string)`

SetCellularTechnology sets CellularTechnology field to given value.

### HasCellularTechnology

`func (o *NetworkV2) HasCellularTechnology() bool`

HasCellularTechnology returns a boolean if a field has been set.

### GetVoiceRoamingEnabled

`func (o *NetworkV2) GetVoiceRoamingEnabled() bool`

GetVoiceRoamingEnabled returns the VoiceRoamingEnabled field if non-nil, zero value otherwise.

### GetVoiceRoamingEnabledOk

`func (o *NetworkV2) GetVoiceRoamingEnabledOk() (*bool, bool)`

GetVoiceRoamingEnabledOk returns a tuple with the VoiceRoamingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceRoamingEnabled

`func (o *NetworkV2) SetVoiceRoamingEnabled(v bool)`

SetVoiceRoamingEnabled sets VoiceRoamingEnabled field to given value.

### HasVoiceRoamingEnabled

`func (o *NetworkV2) HasVoiceRoamingEnabled() bool`

HasVoiceRoamingEnabled returns a boolean if a field has been set.

### GetImei

`func (o *NetworkV2) GetImei() string`

GetImei returns the Imei field if non-nil, zero value otherwise.

### GetImeiOk

`func (o *NetworkV2) GetImeiOk() (*string, bool)`

GetImeiOk returns a tuple with the Imei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImei

`func (o *NetworkV2) SetImei(v string)`

SetImei sets Imei field to given value.

### HasImei

`func (o *NetworkV2) HasImei() bool`

HasImei returns a boolean if a field has been set.

### GetIccid

`func (o *NetworkV2) GetIccid() string`

GetIccid returns the Iccid field if non-nil, zero value otherwise.

### GetIccidOk

`func (o *NetworkV2) GetIccidOk() (*string, bool)`

GetIccidOk returns a tuple with the Iccid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIccid

`func (o *NetworkV2) SetIccid(v string)`

SetIccid sets Iccid field to given value.

### HasIccid

`func (o *NetworkV2) HasIccid() bool`

HasIccid returns a boolean if a field has been set.

### GetMeid

`func (o *NetworkV2) GetMeid() string`

GetMeid returns the Meid field if non-nil, zero value otherwise.

### GetMeidOk

`func (o *NetworkV2) GetMeidOk() (*string, bool)`

GetMeidOk returns a tuple with the Meid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeid

`func (o *NetworkV2) SetMeid(v string)`

SetMeid sets Meid field to given value.

### HasMeid

`func (o *NetworkV2) HasMeid() bool`

HasMeid returns a boolean if a field has been set.

### GetEid

`func (o *NetworkV2) GetEid() string`

GetEid returns the Eid field if non-nil, zero value otherwise.

### GetEidOk

`func (o *NetworkV2) GetEidOk() (*string, bool)`

GetEidOk returns a tuple with the Eid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEid

`func (o *NetworkV2) SetEid(v string)`

SetEid sets Eid field to given value.

### HasEid

`func (o *NetworkV2) HasEid() bool`

HasEid returns a boolean if a field has been set.

### GetCarrierSettingsVersion

`func (o *NetworkV2) GetCarrierSettingsVersion() string`

GetCarrierSettingsVersion returns the CarrierSettingsVersion field if non-nil, zero value otherwise.

### GetCarrierSettingsVersionOk

`func (o *NetworkV2) GetCarrierSettingsVersionOk() (*string, bool)`

GetCarrierSettingsVersionOk returns a tuple with the CarrierSettingsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierSettingsVersion

`func (o *NetworkV2) SetCarrierSettingsVersion(v string)`

SetCarrierSettingsVersion sets CarrierSettingsVersion field to given value.

### HasCarrierSettingsVersion

`func (o *NetworkV2) HasCarrierSettingsVersion() bool`

HasCarrierSettingsVersion returns a boolean if a field has been set.

### GetCurrentCarrierNetwork

`func (o *NetworkV2) GetCurrentCarrierNetwork() string`

GetCurrentCarrierNetwork returns the CurrentCarrierNetwork field if non-nil, zero value otherwise.

### GetCurrentCarrierNetworkOk

`func (o *NetworkV2) GetCurrentCarrierNetworkOk() (*string, bool)`

GetCurrentCarrierNetworkOk returns a tuple with the CurrentCarrierNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentCarrierNetwork

`func (o *NetworkV2) SetCurrentCarrierNetwork(v string)`

SetCurrentCarrierNetwork sets CurrentCarrierNetwork field to given value.

### HasCurrentCarrierNetwork

`func (o *NetworkV2) HasCurrentCarrierNetwork() bool`

HasCurrentCarrierNetwork returns a boolean if a field has been set.

### GetCurrentMobileCountryCode

`func (o *NetworkV2) GetCurrentMobileCountryCode() string`

GetCurrentMobileCountryCode returns the CurrentMobileCountryCode field if non-nil, zero value otherwise.

### GetCurrentMobileCountryCodeOk

`func (o *NetworkV2) GetCurrentMobileCountryCodeOk() (*string, bool)`

GetCurrentMobileCountryCodeOk returns a tuple with the CurrentMobileCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentMobileCountryCode

`func (o *NetworkV2) SetCurrentMobileCountryCode(v string)`

SetCurrentMobileCountryCode sets CurrentMobileCountryCode field to given value.

### HasCurrentMobileCountryCode

`func (o *NetworkV2) HasCurrentMobileCountryCode() bool`

HasCurrentMobileCountryCode returns a boolean if a field has been set.

### GetCurrentMobileNetworkCode

`func (o *NetworkV2) GetCurrentMobileNetworkCode() string`

GetCurrentMobileNetworkCode returns the CurrentMobileNetworkCode field if non-nil, zero value otherwise.

### GetCurrentMobileNetworkCodeOk

`func (o *NetworkV2) GetCurrentMobileNetworkCodeOk() (*string, bool)`

GetCurrentMobileNetworkCodeOk returns a tuple with the CurrentMobileNetworkCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentMobileNetworkCode

`func (o *NetworkV2) SetCurrentMobileNetworkCode(v string)`

SetCurrentMobileNetworkCode sets CurrentMobileNetworkCode field to given value.

### HasCurrentMobileNetworkCode

`func (o *NetworkV2) HasCurrentMobileNetworkCode() bool`

HasCurrentMobileNetworkCode returns a boolean if a field has been set.

### GetHomeCarrierNetwork

`func (o *NetworkV2) GetHomeCarrierNetwork() string`

GetHomeCarrierNetwork returns the HomeCarrierNetwork field if non-nil, zero value otherwise.

### GetHomeCarrierNetworkOk

`func (o *NetworkV2) GetHomeCarrierNetworkOk() (*string, bool)`

GetHomeCarrierNetworkOk returns a tuple with the HomeCarrierNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeCarrierNetwork

`func (o *NetworkV2) SetHomeCarrierNetwork(v string)`

SetHomeCarrierNetwork sets HomeCarrierNetwork field to given value.

### HasHomeCarrierNetwork

`func (o *NetworkV2) HasHomeCarrierNetwork() bool`

HasHomeCarrierNetwork returns a boolean if a field has been set.

### GetHomeMobileCountryCode

`func (o *NetworkV2) GetHomeMobileCountryCode() string`

GetHomeMobileCountryCode returns the HomeMobileCountryCode field if non-nil, zero value otherwise.

### GetHomeMobileCountryCodeOk

`func (o *NetworkV2) GetHomeMobileCountryCodeOk() (*string, bool)`

GetHomeMobileCountryCodeOk returns a tuple with the HomeMobileCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeMobileCountryCode

`func (o *NetworkV2) SetHomeMobileCountryCode(v string)`

SetHomeMobileCountryCode sets HomeMobileCountryCode field to given value.

### HasHomeMobileCountryCode

`func (o *NetworkV2) HasHomeMobileCountryCode() bool`

HasHomeMobileCountryCode returns a boolean if a field has been set.

### GetHomeMobileNetworkCode

`func (o *NetworkV2) GetHomeMobileNetworkCode() string`

GetHomeMobileNetworkCode returns the HomeMobileNetworkCode field if non-nil, zero value otherwise.

### GetHomeMobileNetworkCodeOk

`func (o *NetworkV2) GetHomeMobileNetworkCodeOk() (*string, bool)`

GetHomeMobileNetworkCodeOk returns a tuple with the HomeMobileNetworkCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeMobileNetworkCode

`func (o *NetworkV2) SetHomeMobileNetworkCode(v string)`

SetHomeMobileNetworkCode sets HomeMobileNetworkCode field to given value.

### HasHomeMobileNetworkCode

`func (o *NetworkV2) HasHomeMobileNetworkCode() bool`

HasHomeMobileNetworkCode returns a boolean if a field has been set.

### GetDataRoamingEnabled

`func (o *NetworkV2) GetDataRoamingEnabled() bool`

GetDataRoamingEnabled returns the DataRoamingEnabled field if non-nil, zero value otherwise.

### GetDataRoamingEnabledOk

`func (o *NetworkV2) GetDataRoamingEnabledOk() (*bool, bool)`

GetDataRoamingEnabledOk returns a tuple with the DataRoamingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRoamingEnabled

`func (o *NetworkV2) SetDataRoamingEnabled(v bool)`

SetDataRoamingEnabled sets DataRoamingEnabled field to given value.

### HasDataRoamingEnabled

`func (o *NetworkV2) HasDataRoamingEnabled() bool`

HasDataRoamingEnabled returns a boolean if a field has been set.

### GetRoaming

`func (o *NetworkV2) GetRoaming() bool`

GetRoaming returns the Roaming field if non-nil, zero value otherwise.

### GetRoamingOk

`func (o *NetworkV2) GetRoamingOk() (*bool, bool)`

GetRoamingOk returns a tuple with the Roaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoaming

`func (o *NetworkV2) SetRoaming(v bool)`

SetRoaming sets Roaming field to given value.

### HasRoaming

`func (o *NetworkV2) HasRoaming() bool`

HasRoaming returns a boolean if a field has been set.

### GetPersonalHotspotEnabled

`func (o *NetworkV2) GetPersonalHotspotEnabled() bool`

GetPersonalHotspotEnabled returns the PersonalHotspotEnabled field if non-nil, zero value otherwise.

### GetPersonalHotspotEnabledOk

`func (o *NetworkV2) GetPersonalHotspotEnabledOk() (*bool, bool)`

GetPersonalHotspotEnabledOk returns a tuple with the PersonalHotspotEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalHotspotEnabled

`func (o *NetworkV2) SetPersonalHotspotEnabled(v bool)`

SetPersonalHotspotEnabled sets PersonalHotspotEnabled field to given value.

### HasPersonalHotspotEnabled

`func (o *NetworkV2) HasPersonalHotspotEnabled() bool`

HasPersonalHotspotEnabled returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *NetworkV2) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *NetworkV2) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *NetworkV2) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *NetworkV2) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


