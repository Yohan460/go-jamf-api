# DeviceCommunicationSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoRenewMobileDeviceMdmProfileWhenCaRenewed** | Pointer to **bool** |  | [optional] 
**AutoRenewMobileDeviceMdmProfileWhenDeviceIdentityCertExpiring** | Pointer to **bool** |  | [optional] 
**AutoRenewComputerMdmProfileWhenCaRenewed** | Pointer to **bool** |  | [optional] 
**AutoRenewComputerMdmProfileWhenDeviceIdentityCertExpiring** | Pointer to **bool** |  | [optional] 
**MdmProfileMobileDeviceExpirationLimitInDays** | Pointer to **int32** |  | [optional] [default to 180]
**MdmProfileComputerExpirationLimitInDays** | Pointer to **int32** |  | [optional] [default to 180]

## Methods

### NewDeviceCommunicationSettings

`func NewDeviceCommunicationSettings() *DeviceCommunicationSettings`

NewDeviceCommunicationSettings instantiates a new DeviceCommunicationSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceCommunicationSettingsWithDefaults

`func NewDeviceCommunicationSettingsWithDefaults() *DeviceCommunicationSettings`

NewDeviceCommunicationSettingsWithDefaults instantiates a new DeviceCommunicationSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoRenewMobileDeviceMdmProfileWhenCaRenewed

`func (o *DeviceCommunicationSettings) GetAutoRenewMobileDeviceMdmProfileWhenCaRenewed() bool`

GetAutoRenewMobileDeviceMdmProfileWhenCaRenewed returns the AutoRenewMobileDeviceMdmProfileWhenCaRenewed field if non-nil, zero value otherwise.

### GetAutoRenewMobileDeviceMdmProfileWhenCaRenewedOk

`func (o *DeviceCommunicationSettings) GetAutoRenewMobileDeviceMdmProfileWhenCaRenewedOk() (*bool, bool)`

GetAutoRenewMobileDeviceMdmProfileWhenCaRenewedOk returns a tuple with the AutoRenewMobileDeviceMdmProfileWhenCaRenewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRenewMobileDeviceMdmProfileWhenCaRenewed

`func (o *DeviceCommunicationSettings) SetAutoRenewMobileDeviceMdmProfileWhenCaRenewed(v bool)`

SetAutoRenewMobileDeviceMdmProfileWhenCaRenewed sets AutoRenewMobileDeviceMdmProfileWhenCaRenewed field to given value.

### HasAutoRenewMobileDeviceMdmProfileWhenCaRenewed

`func (o *DeviceCommunicationSettings) HasAutoRenewMobileDeviceMdmProfileWhenCaRenewed() bool`

HasAutoRenewMobileDeviceMdmProfileWhenCaRenewed returns a boolean if a field has been set.

### GetAutoRenewMobileDeviceMdmProfileWhenDeviceIdentityCertExpiring

`func (o *DeviceCommunicationSettings) GetAutoRenewMobileDeviceMdmProfileWhenDeviceIdentityCertExpiring() bool`

GetAutoRenewMobileDeviceMdmProfileWhenDeviceIdentityCertExpiring returns the AutoRenewMobileDeviceMdmProfileWhenDeviceIdentityCertExpiring field if non-nil, zero value otherwise.

### GetAutoRenewMobileDeviceMdmProfileWhenDeviceIdentityCertExpiringOk

`func (o *DeviceCommunicationSettings) GetAutoRenewMobileDeviceMdmProfileWhenDeviceIdentityCertExpiringOk() (*bool, bool)`

GetAutoRenewMobileDeviceMdmProfileWhenDeviceIdentityCertExpiringOk returns a tuple with the AutoRenewMobileDeviceMdmProfileWhenDeviceIdentityCertExpiring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRenewMobileDeviceMdmProfileWhenDeviceIdentityCertExpiring

`func (o *DeviceCommunicationSettings) SetAutoRenewMobileDeviceMdmProfileWhenDeviceIdentityCertExpiring(v bool)`

SetAutoRenewMobileDeviceMdmProfileWhenDeviceIdentityCertExpiring sets AutoRenewMobileDeviceMdmProfileWhenDeviceIdentityCertExpiring field to given value.

### HasAutoRenewMobileDeviceMdmProfileWhenDeviceIdentityCertExpiring

`func (o *DeviceCommunicationSettings) HasAutoRenewMobileDeviceMdmProfileWhenDeviceIdentityCertExpiring() bool`

HasAutoRenewMobileDeviceMdmProfileWhenDeviceIdentityCertExpiring returns a boolean if a field has been set.

### GetAutoRenewComputerMdmProfileWhenCaRenewed

`func (o *DeviceCommunicationSettings) GetAutoRenewComputerMdmProfileWhenCaRenewed() bool`

GetAutoRenewComputerMdmProfileWhenCaRenewed returns the AutoRenewComputerMdmProfileWhenCaRenewed field if non-nil, zero value otherwise.

### GetAutoRenewComputerMdmProfileWhenCaRenewedOk

`func (o *DeviceCommunicationSettings) GetAutoRenewComputerMdmProfileWhenCaRenewedOk() (*bool, bool)`

GetAutoRenewComputerMdmProfileWhenCaRenewedOk returns a tuple with the AutoRenewComputerMdmProfileWhenCaRenewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRenewComputerMdmProfileWhenCaRenewed

`func (o *DeviceCommunicationSettings) SetAutoRenewComputerMdmProfileWhenCaRenewed(v bool)`

SetAutoRenewComputerMdmProfileWhenCaRenewed sets AutoRenewComputerMdmProfileWhenCaRenewed field to given value.

### HasAutoRenewComputerMdmProfileWhenCaRenewed

`func (o *DeviceCommunicationSettings) HasAutoRenewComputerMdmProfileWhenCaRenewed() bool`

HasAutoRenewComputerMdmProfileWhenCaRenewed returns a boolean if a field has been set.

### GetAutoRenewComputerMdmProfileWhenDeviceIdentityCertExpiring

`func (o *DeviceCommunicationSettings) GetAutoRenewComputerMdmProfileWhenDeviceIdentityCertExpiring() bool`

GetAutoRenewComputerMdmProfileWhenDeviceIdentityCertExpiring returns the AutoRenewComputerMdmProfileWhenDeviceIdentityCertExpiring field if non-nil, zero value otherwise.

### GetAutoRenewComputerMdmProfileWhenDeviceIdentityCertExpiringOk

`func (o *DeviceCommunicationSettings) GetAutoRenewComputerMdmProfileWhenDeviceIdentityCertExpiringOk() (*bool, bool)`

GetAutoRenewComputerMdmProfileWhenDeviceIdentityCertExpiringOk returns a tuple with the AutoRenewComputerMdmProfileWhenDeviceIdentityCertExpiring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRenewComputerMdmProfileWhenDeviceIdentityCertExpiring

`func (o *DeviceCommunicationSettings) SetAutoRenewComputerMdmProfileWhenDeviceIdentityCertExpiring(v bool)`

SetAutoRenewComputerMdmProfileWhenDeviceIdentityCertExpiring sets AutoRenewComputerMdmProfileWhenDeviceIdentityCertExpiring field to given value.

### HasAutoRenewComputerMdmProfileWhenDeviceIdentityCertExpiring

`func (o *DeviceCommunicationSettings) HasAutoRenewComputerMdmProfileWhenDeviceIdentityCertExpiring() bool`

HasAutoRenewComputerMdmProfileWhenDeviceIdentityCertExpiring returns a boolean if a field has been set.

### GetMdmProfileMobileDeviceExpirationLimitInDays

`func (o *DeviceCommunicationSettings) GetMdmProfileMobileDeviceExpirationLimitInDays() int32`

GetMdmProfileMobileDeviceExpirationLimitInDays returns the MdmProfileMobileDeviceExpirationLimitInDays field if non-nil, zero value otherwise.

### GetMdmProfileMobileDeviceExpirationLimitInDaysOk

`func (o *DeviceCommunicationSettings) GetMdmProfileMobileDeviceExpirationLimitInDaysOk() (*int32, bool)`

GetMdmProfileMobileDeviceExpirationLimitInDaysOk returns a tuple with the MdmProfileMobileDeviceExpirationLimitInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdmProfileMobileDeviceExpirationLimitInDays

`func (o *DeviceCommunicationSettings) SetMdmProfileMobileDeviceExpirationLimitInDays(v int32)`

SetMdmProfileMobileDeviceExpirationLimitInDays sets MdmProfileMobileDeviceExpirationLimitInDays field to given value.

### HasMdmProfileMobileDeviceExpirationLimitInDays

`func (o *DeviceCommunicationSettings) HasMdmProfileMobileDeviceExpirationLimitInDays() bool`

HasMdmProfileMobileDeviceExpirationLimitInDays returns a boolean if a field has been set.

### GetMdmProfileComputerExpirationLimitInDays

`func (o *DeviceCommunicationSettings) GetMdmProfileComputerExpirationLimitInDays() int32`

GetMdmProfileComputerExpirationLimitInDays returns the MdmProfileComputerExpirationLimitInDays field if non-nil, zero value otherwise.

### GetMdmProfileComputerExpirationLimitInDaysOk

`func (o *DeviceCommunicationSettings) GetMdmProfileComputerExpirationLimitInDaysOk() (*int32, bool)`

GetMdmProfileComputerExpirationLimitInDaysOk returns a tuple with the MdmProfileComputerExpirationLimitInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdmProfileComputerExpirationLimitInDays

`func (o *DeviceCommunicationSettings) SetMdmProfileComputerExpirationLimitInDays(v int32)`

SetMdmProfileComputerExpirationLimitInDays sets MdmProfileComputerExpirationLimitInDays field to given value.

### HasMdmProfileComputerExpirationLimitInDays

`func (o *DeviceCommunicationSettings) HasMdmProfileComputerExpirationLimitInDays() bool`

HasMdmProfileComputerExpirationLimitInDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


