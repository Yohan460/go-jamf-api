# MobileDeviceServiceSubscriptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CarrierSettingsVersion** | Pointer to **string** |  | [optional] 
**CurrentCarrierNetwork** | Pointer to **string** |  | [optional] 
**CurrentMobileCountryCode** | Pointer to **string** |  | [optional] 
**CurrentMobileNetworkCode** | Pointer to **string** |  | [optional] 
**SubscriberCarrierNetwork** | Pointer to **string** |  | [optional] 
**Eid** | Pointer to **string** |  | [optional] 
**Iccid** | Pointer to **string** |  | [optional] 
**Imei** | Pointer to **string** |  | [optional] 
**DataPreferred** | Pointer to **bool** |  | [optional] 
**Roaming** | Pointer to **bool** |  | [optional] 
**VoicePreferred** | Pointer to **bool** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**LabelId** | Pointer to **string** | The unique identifier for this subscription. | [optional] 
**Meid** | Pointer to **string** |  | [optional] 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**Slot** | Pointer to **string** | The description of the slot that contains the SIM representing this subscription. | [optional] 

## Methods

### NewMobileDeviceServiceSubscriptions

`func NewMobileDeviceServiceSubscriptions() *MobileDeviceServiceSubscriptions`

NewMobileDeviceServiceSubscriptions instantiates a new MobileDeviceServiceSubscriptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileDeviceServiceSubscriptionsWithDefaults

`func NewMobileDeviceServiceSubscriptionsWithDefaults() *MobileDeviceServiceSubscriptions`

NewMobileDeviceServiceSubscriptionsWithDefaults instantiates a new MobileDeviceServiceSubscriptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCarrierSettingsVersion

`func (o *MobileDeviceServiceSubscriptions) GetCarrierSettingsVersion() string`

GetCarrierSettingsVersion returns the CarrierSettingsVersion field if non-nil, zero value otherwise.

### GetCarrierSettingsVersionOk

`func (o *MobileDeviceServiceSubscriptions) GetCarrierSettingsVersionOk() (*string, bool)`

GetCarrierSettingsVersionOk returns a tuple with the CarrierSettingsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierSettingsVersion

`func (o *MobileDeviceServiceSubscriptions) SetCarrierSettingsVersion(v string)`

SetCarrierSettingsVersion sets CarrierSettingsVersion field to given value.

### HasCarrierSettingsVersion

`func (o *MobileDeviceServiceSubscriptions) HasCarrierSettingsVersion() bool`

HasCarrierSettingsVersion returns a boolean if a field has been set.

### GetCurrentCarrierNetwork

`func (o *MobileDeviceServiceSubscriptions) GetCurrentCarrierNetwork() string`

GetCurrentCarrierNetwork returns the CurrentCarrierNetwork field if non-nil, zero value otherwise.

### GetCurrentCarrierNetworkOk

`func (o *MobileDeviceServiceSubscriptions) GetCurrentCarrierNetworkOk() (*string, bool)`

GetCurrentCarrierNetworkOk returns a tuple with the CurrentCarrierNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentCarrierNetwork

`func (o *MobileDeviceServiceSubscriptions) SetCurrentCarrierNetwork(v string)`

SetCurrentCarrierNetwork sets CurrentCarrierNetwork field to given value.

### HasCurrentCarrierNetwork

`func (o *MobileDeviceServiceSubscriptions) HasCurrentCarrierNetwork() bool`

HasCurrentCarrierNetwork returns a boolean if a field has been set.

### GetCurrentMobileCountryCode

`func (o *MobileDeviceServiceSubscriptions) GetCurrentMobileCountryCode() string`

GetCurrentMobileCountryCode returns the CurrentMobileCountryCode field if non-nil, zero value otherwise.

### GetCurrentMobileCountryCodeOk

`func (o *MobileDeviceServiceSubscriptions) GetCurrentMobileCountryCodeOk() (*string, bool)`

GetCurrentMobileCountryCodeOk returns a tuple with the CurrentMobileCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentMobileCountryCode

`func (o *MobileDeviceServiceSubscriptions) SetCurrentMobileCountryCode(v string)`

SetCurrentMobileCountryCode sets CurrentMobileCountryCode field to given value.

### HasCurrentMobileCountryCode

`func (o *MobileDeviceServiceSubscriptions) HasCurrentMobileCountryCode() bool`

HasCurrentMobileCountryCode returns a boolean if a field has been set.

### GetCurrentMobileNetworkCode

`func (o *MobileDeviceServiceSubscriptions) GetCurrentMobileNetworkCode() string`

GetCurrentMobileNetworkCode returns the CurrentMobileNetworkCode field if non-nil, zero value otherwise.

### GetCurrentMobileNetworkCodeOk

`func (o *MobileDeviceServiceSubscriptions) GetCurrentMobileNetworkCodeOk() (*string, bool)`

GetCurrentMobileNetworkCodeOk returns a tuple with the CurrentMobileNetworkCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentMobileNetworkCode

`func (o *MobileDeviceServiceSubscriptions) SetCurrentMobileNetworkCode(v string)`

SetCurrentMobileNetworkCode sets CurrentMobileNetworkCode field to given value.

### HasCurrentMobileNetworkCode

`func (o *MobileDeviceServiceSubscriptions) HasCurrentMobileNetworkCode() bool`

HasCurrentMobileNetworkCode returns a boolean if a field has been set.

### GetSubscriberCarrierNetwork

`func (o *MobileDeviceServiceSubscriptions) GetSubscriberCarrierNetwork() string`

GetSubscriberCarrierNetwork returns the SubscriberCarrierNetwork field if non-nil, zero value otherwise.

### GetSubscriberCarrierNetworkOk

`func (o *MobileDeviceServiceSubscriptions) GetSubscriberCarrierNetworkOk() (*string, bool)`

GetSubscriberCarrierNetworkOk returns a tuple with the SubscriberCarrierNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriberCarrierNetwork

`func (o *MobileDeviceServiceSubscriptions) SetSubscriberCarrierNetwork(v string)`

SetSubscriberCarrierNetwork sets SubscriberCarrierNetwork field to given value.

### HasSubscriberCarrierNetwork

`func (o *MobileDeviceServiceSubscriptions) HasSubscriberCarrierNetwork() bool`

HasSubscriberCarrierNetwork returns a boolean if a field has been set.

### GetEid

`func (o *MobileDeviceServiceSubscriptions) GetEid() string`

GetEid returns the Eid field if non-nil, zero value otherwise.

### GetEidOk

`func (o *MobileDeviceServiceSubscriptions) GetEidOk() (*string, bool)`

GetEidOk returns a tuple with the Eid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEid

`func (o *MobileDeviceServiceSubscriptions) SetEid(v string)`

SetEid sets Eid field to given value.

### HasEid

`func (o *MobileDeviceServiceSubscriptions) HasEid() bool`

HasEid returns a boolean if a field has been set.

### GetIccid

`func (o *MobileDeviceServiceSubscriptions) GetIccid() string`

GetIccid returns the Iccid field if non-nil, zero value otherwise.

### GetIccidOk

`func (o *MobileDeviceServiceSubscriptions) GetIccidOk() (*string, bool)`

GetIccidOk returns a tuple with the Iccid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIccid

`func (o *MobileDeviceServiceSubscriptions) SetIccid(v string)`

SetIccid sets Iccid field to given value.

### HasIccid

`func (o *MobileDeviceServiceSubscriptions) HasIccid() bool`

HasIccid returns a boolean if a field has been set.

### GetImei

`func (o *MobileDeviceServiceSubscriptions) GetImei() string`

GetImei returns the Imei field if non-nil, zero value otherwise.

### GetImeiOk

`func (o *MobileDeviceServiceSubscriptions) GetImeiOk() (*string, bool)`

GetImeiOk returns a tuple with the Imei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImei

`func (o *MobileDeviceServiceSubscriptions) SetImei(v string)`

SetImei sets Imei field to given value.

### HasImei

`func (o *MobileDeviceServiceSubscriptions) HasImei() bool`

HasImei returns a boolean if a field has been set.

### GetDataPreferred

`func (o *MobileDeviceServiceSubscriptions) GetDataPreferred() bool`

GetDataPreferred returns the DataPreferred field if non-nil, zero value otherwise.

### GetDataPreferredOk

`func (o *MobileDeviceServiceSubscriptions) GetDataPreferredOk() (*bool, bool)`

GetDataPreferredOk returns a tuple with the DataPreferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPreferred

`func (o *MobileDeviceServiceSubscriptions) SetDataPreferred(v bool)`

SetDataPreferred sets DataPreferred field to given value.

### HasDataPreferred

`func (o *MobileDeviceServiceSubscriptions) HasDataPreferred() bool`

HasDataPreferred returns a boolean if a field has been set.

### GetRoaming

`func (o *MobileDeviceServiceSubscriptions) GetRoaming() bool`

GetRoaming returns the Roaming field if non-nil, zero value otherwise.

### GetRoamingOk

`func (o *MobileDeviceServiceSubscriptions) GetRoamingOk() (*bool, bool)`

GetRoamingOk returns a tuple with the Roaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoaming

`func (o *MobileDeviceServiceSubscriptions) SetRoaming(v bool)`

SetRoaming sets Roaming field to given value.

### HasRoaming

`func (o *MobileDeviceServiceSubscriptions) HasRoaming() bool`

HasRoaming returns a boolean if a field has been set.

### GetVoicePreferred

`func (o *MobileDeviceServiceSubscriptions) GetVoicePreferred() bool`

GetVoicePreferred returns the VoicePreferred field if non-nil, zero value otherwise.

### GetVoicePreferredOk

`func (o *MobileDeviceServiceSubscriptions) GetVoicePreferredOk() (*bool, bool)`

GetVoicePreferredOk returns a tuple with the VoicePreferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoicePreferred

`func (o *MobileDeviceServiceSubscriptions) SetVoicePreferred(v bool)`

SetVoicePreferred sets VoicePreferred field to given value.

### HasVoicePreferred

`func (o *MobileDeviceServiceSubscriptions) HasVoicePreferred() bool`

HasVoicePreferred returns a boolean if a field has been set.

### GetLabel

`func (o *MobileDeviceServiceSubscriptions) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *MobileDeviceServiceSubscriptions) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *MobileDeviceServiceSubscriptions) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *MobileDeviceServiceSubscriptions) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLabelId

`func (o *MobileDeviceServiceSubscriptions) GetLabelId() string`

GetLabelId returns the LabelId field if non-nil, zero value otherwise.

### GetLabelIdOk

`func (o *MobileDeviceServiceSubscriptions) GetLabelIdOk() (*string, bool)`

GetLabelIdOk returns a tuple with the LabelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelId

`func (o *MobileDeviceServiceSubscriptions) SetLabelId(v string)`

SetLabelId sets LabelId field to given value.

### HasLabelId

`func (o *MobileDeviceServiceSubscriptions) HasLabelId() bool`

HasLabelId returns a boolean if a field has been set.

### GetMeid

`func (o *MobileDeviceServiceSubscriptions) GetMeid() string`

GetMeid returns the Meid field if non-nil, zero value otherwise.

### GetMeidOk

`func (o *MobileDeviceServiceSubscriptions) GetMeidOk() (*string, bool)`

GetMeidOk returns a tuple with the Meid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeid

`func (o *MobileDeviceServiceSubscriptions) SetMeid(v string)`

SetMeid sets Meid field to given value.

### HasMeid

`func (o *MobileDeviceServiceSubscriptions) HasMeid() bool`

HasMeid returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *MobileDeviceServiceSubscriptions) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *MobileDeviceServiceSubscriptions) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *MobileDeviceServiceSubscriptions) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *MobileDeviceServiceSubscriptions) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetSlot

`func (o *MobileDeviceServiceSubscriptions) GetSlot() string`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *MobileDeviceServiceSubscriptions) GetSlotOk() (*string, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *MobileDeviceServiceSubscriptions) SetSlot(v string)`

SetSlot sets Slot field to given value.

### HasSlot

`func (o *MobileDeviceServiceSubscriptions) HasSlot() bool`

HasSlot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


