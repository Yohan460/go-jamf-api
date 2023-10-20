# MobileDeviceNetwork

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

### NewMobileDeviceNetwork

`func NewMobileDeviceNetwork() *MobileDeviceNetwork`

NewMobileDeviceNetwork instantiates a new MobileDeviceNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileDeviceNetworkWithDefaults

`func NewMobileDeviceNetworkWithDefaults() *MobileDeviceNetwork`

NewMobileDeviceNetworkWithDefaults instantiates a new MobileDeviceNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCellularTechnology

`func (o *MobileDeviceNetwork) GetCellularTechnology() string`

GetCellularTechnology returns the CellularTechnology field if non-nil, zero value otherwise.

### GetCellularTechnologyOk

`func (o *MobileDeviceNetwork) GetCellularTechnologyOk() (*string, bool)`

GetCellularTechnologyOk returns a tuple with the CellularTechnology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellularTechnology

`func (o *MobileDeviceNetwork) SetCellularTechnology(v string)`

SetCellularTechnology sets CellularTechnology field to given value.

### HasCellularTechnology

`func (o *MobileDeviceNetwork) HasCellularTechnology() bool`

HasCellularTechnology returns a boolean if a field has been set.

### GetVoiceRoamingEnabled

`func (o *MobileDeviceNetwork) GetVoiceRoamingEnabled() bool`

GetVoiceRoamingEnabled returns the VoiceRoamingEnabled field if non-nil, zero value otherwise.

### GetVoiceRoamingEnabledOk

`func (o *MobileDeviceNetwork) GetVoiceRoamingEnabledOk() (*bool, bool)`

GetVoiceRoamingEnabledOk returns a tuple with the VoiceRoamingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceRoamingEnabled

`func (o *MobileDeviceNetwork) SetVoiceRoamingEnabled(v bool)`

SetVoiceRoamingEnabled sets VoiceRoamingEnabled field to given value.

### HasVoiceRoamingEnabled

`func (o *MobileDeviceNetwork) HasVoiceRoamingEnabled() bool`

HasVoiceRoamingEnabled returns a boolean if a field has been set.

### GetImei

`func (o *MobileDeviceNetwork) GetImei() string`

GetImei returns the Imei field if non-nil, zero value otherwise.

### GetImeiOk

`func (o *MobileDeviceNetwork) GetImeiOk() (*string, bool)`

GetImeiOk returns a tuple with the Imei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImei

`func (o *MobileDeviceNetwork) SetImei(v string)`

SetImei sets Imei field to given value.

### HasImei

`func (o *MobileDeviceNetwork) HasImei() bool`

HasImei returns a boolean if a field has been set.

### GetIccid

`func (o *MobileDeviceNetwork) GetIccid() string`

GetIccid returns the Iccid field if non-nil, zero value otherwise.

### GetIccidOk

`func (o *MobileDeviceNetwork) GetIccidOk() (*string, bool)`

GetIccidOk returns a tuple with the Iccid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIccid

`func (o *MobileDeviceNetwork) SetIccid(v string)`

SetIccid sets Iccid field to given value.

### HasIccid

`func (o *MobileDeviceNetwork) HasIccid() bool`

HasIccid returns a boolean if a field has been set.

### GetMeid

`func (o *MobileDeviceNetwork) GetMeid() string`

GetMeid returns the Meid field if non-nil, zero value otherwise.

### GetMeidOk

`func (o *MobileDeviceNetwork) GetMeidOk() (*string, bool)`

GetMeidOk returns a tuple with the Meid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeid

`func (o *MobileDeviceNetwork) SetMeid(v string)`

SetMeid sets Meid field to given value.

### HasMeid

`func (o *MobileDeviceNetwork) HasMeid() bool`

HasMeid returns a boolean if a field has been set.

### GetEid

`func (o *MobileDeviceNetwork) GetEid() string`

GetEid returns the Eid field if non-nil, zero value otherwise.

### GetEidOk

`func (o *MobileDeviceNetwork) GetEidOk() (*string, bool)`

GetEidOk returns a tuple with the Eid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEid

`func (o *MobileDeviceNetwork) SetEid(v string)`

SetEid sets Eid field to given value.

### HasEid

`func (o *MobileDeviceNetwork) HasEid() bool`

HasEid returns a boolean if a field has been set.

### GetCarrierSettingsVersion

`func (o *MobileDeviceNetwork) GetCarrierSettingsVersion() string`

GetCarrierSettingsVersion returns the CarrierSettingsVersion field if non-nil, zero value otherwise.

### GetCarrierSettingsVersionOk

`func (o *MobileDeviceNetwork) GetCarrierSettingsVersionOk() (*string, bool)`

GetCarrierSettingsVersionOk returns a tuple with the CarrierSettingsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierSettingsVersion

`func (o *MobileDeviceNetwork) SetCarrierSettingsVersion(v string)`

SetCarrierSettingsVersion sets CarrierSettingsVersion field to given value.

### HasCarrierSettingsVersion

`func (o *MobileDeviceNetwork) HasCarrierSettingsVersion() bool`

HasCarrierSettingsVersion returns a boolean if a field has been set.

### GetCurrentCarrierNetwork

`func (o *MobileDeviceNetwork) GetCurrentCarrierNetwork() string`

GetCurrentCarrierNetwork returns the CurrentCarrierNetwork field if non-nil, zero value otherwise.

### GetCurrentCarrierNetworkOk

`func (o *MobileDeviceNetwork) GetCurrentCarrierNetworkOk() (*string, bool)`

GetCurrentCarrierNetworkOk returns a tuple with the CurrentCarrierNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentCarrierNetwork

`func (o *MobileDeviceNetwork) SetCurrentCarrierNetwork(v string)`

SetCurrentCarrierNetwork sets CurrentCarrierNetwork field to given value.

### HasCurrentCarrierNetwork

`func (o *MobileDeviceNetwork) HasCurrentCarrierNetwork() bool`

HasCurrentCarrierNetwork returns a boolean if a field has been set.

### GetCurrentMobileCountryCode

`func (o *MobileDeviceNetwork) GetCurrentMobileCountryCode() string`

GetCurrentMobileCountryCode returns the CurrentMobileCountryCode field if non-nil, zero value otherwise.

### GetCurrentMobileCountryCodeOk

`func (o *MobileDeviceNetwork) GetCurrentMobileCountryCodeOk() (*string, bool)`

GetCurrentMobileCountryCodeOk returns a tuple with the CurrentMobileCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentMobileCountryCode

`func (o *MobileDeviceNetwork) SetCurrentMobileCountryCode(v string)`

SetCurrentMobileCountryCode sets CurrentMobileCountryCode field to given value.

### HasCurrentMobileCountryCode

`func (o *MobileDeviceNetwork) HasCurrentMobileCountryCode() bool`

HasCurrentMobileCountryCode returns a boolean if a field has been set.

### GetCurrentMobileNetworkCode

`func (o *MobileDeviceNetwork) GetCurrentMobileNetworkCode() string`

GetCurrentMobileNetworkCode returns the CurrentMobileNetworkCode field if non-nil, zero value otherwise.

### GetCurrentMobileNetworkCodeOk

`func (o *MobileDeviceNetwork) GetCurrentMobileNetworkCodeOk() (*string, bool)`

GetCurrentMobileNetworkCodeOk returns a tuple with the CurrentMobileNetworkCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentMobileNetworkCode

`func (o *MobileDeviceNetwork) SetCurrentMobileNetworkCode(v string)`

SetCurrentMobileNetworkCode sets CurrentMobileNetworkCode field to given value.

### HasCurrentMobileNetworkCode

`func (o *MobileDeviceNetwork) HasCurrentMobileNetworkCode() bool`

HasCurrentMobileNetworkCode returns a boolean if a field has been set.

### GetHomeCarrierNetwork

`func (o *MobileDeviceNetwork) GetHomeCarrierNetwork() string`

GetHomeCarrierNetwork returns the HomeCarrierNetwork field if non-nil, zero value otherwise.

### GetHomeCarrierNetworkOk

`func (o *MobileDeviceNetwork) GetHomeCarrierNetworkOk() (*string, bool)`

GetHomeCarrierNetworkOk returns a tuple with the HomeCarrierNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeCarrierNetwork

`func (o *MobileDeviceNetwork) SetHomeCarrierNetwork(v string)`

SetHomeCarrierNetwork sets HomeCarrierNetwork field to given value.

### HasHomeCarrierNetwork

`func (o *MobileDeviceNetwork) HasHomeCarrierNetwork() bool`

HasHomeCarrierNetwork returns a boolean if a field has been set.

### GetHomeMobileCountryCode

`func (o *MobileDeviceNetwork) GetHomeMobileCountryCode() string`

GetHomeMobileCountryCode returns the HomeMobileCountryCode field if non-nil, zero value otherwise.

### GetHomeMobileCountryCodeOk

`func (o *MobileDeviceNetwork) GetHomeMobileCountryCodeOk() (*string, bool)`

GetHomeMobileCountryCodeOk returns a tuple with the HomeMobileCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeMobileCountryCode

`func (o *MobileDeviceNetwork) SetHomeMobileCountryCode(v string)`

SetHomeMobileCountryCode sets HomeMobileCountryCode field to given value.

### HasHomeMobileCountryCode

`func (o *MobileDeviceNetwork) HasHomeMobileCountryCode() bool`

HasHomeMobileCountryCode returns a boolean if a field has been set.

### GetHomeMobileNetworkCode

`func (o *MobileDeviceNetwork) GetHomeMobileNetworkCode() string`

GetHomeMobileNetworkCode returns the HomeMobileNetworkCode field if non-nil, zero value otherwise.

### GetHomeMobileNetworkCodeOk

`func (o *MobileDeviceNetwork) GetHomeMobileNetworkCodeOk() (*string, bool)`

GetHomeMobileNetworkCodeOk returns a tuple with the HomeMobileNetworkCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeMobileNetworkCode

`func (o *MobileDeviceNetwork) SetHomeMobileNetworkCode(v string)`

SetHomeMobileNetworkCode sets HomeMobileNetworkCode field to given value.

### HasHomeMobileNetworkCode

`func (o *MobileDeviceNetwork) HasHomeMobileNetworkCode() bool`

HasHomeMobileNetworkCode returns a boolean if a field has been set.

### GetDataRoamingEnabled

`func (o *MobileDeviceNetwork) GetDataRoamingEnabled() bool`

GetDataRoamingEnabled returns the DataRoamingEnabled field if non-nil, zero value otherwise.

### GetDataRoamingEnabledOk

`func (o *MobileDeviceNetwork) GetDataRoamingEnabledOk() (*bool, bool)`

GetDataRoamingEnabledOk returns a tuple with the DataRoamingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRoamingEnabled

`func (o *MobileDeviceNetwork) SetDataRoamingEnabled(v bool)`

SetDataRoamingEnabled sets DataRoamingEnabled field to given value.

### HasDataRoamingEnabled

`func (o *MobileDeviceNetwork) HasDataRoamingEnabled() bool`

HasDataRoamingEnabled returns a boolean if a field has been set.

### GetRoaming

`func (o *MobileDeviceNetwork) GetRoaming() bool`

GetRoaming returns the Roaming field if non-nil, zero value otherwise.

### GetRoamingOk

`func (o *MobileDeviceNetwork) GetRoamingOk() (*bool, bool)`

GetRoamingOk returns a tuple with the Roaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoaming

`func (o *MobileDeviceNetwork) SetRoaming(v bool)`

SetRoaming sets Roaming field to given value.

### HasRoaming

`func (o *MobileDeviceNetwork) HasRoaming() bool`

HasRoaming returns a boolean if a field has been set.

### GetPersonalHotspotEnabled

`func (o *MobileDeviceNetwork) GetPersonalHotspotEnabled() bool`

GetPersonalHotspotEnabled returns the PersonalHotspotEnabled field if non-nil, zero value otherwise.

### GetPersonalHotspotEnabledOk

`func (o *MobileDeviceNetwork) GetPersonalHotspotEnabledOk() (*bool, bool)`

GetPersonalHotspotEnabledOk returns a tuple with the PersonalHotspotEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalHotspotEnabled

`func (o *MobileDeviceNetwork) SetPersonalHotspotEnabled(v bool)`

SetPersonalHotspotEnabled sets PersonalHotspotEnabled field to given value.

### HasPersonalHotspotEnabled

`func (o *MobileDeviceNetwork) HasPersonalHotspotEnabled() bool`

HasPersonalHotspotEnabled returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *MobileDeviceNetwork) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *MobileDeviceNetwork) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *MobileDeviceNetwork) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *MobileDeviceNetwork) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


