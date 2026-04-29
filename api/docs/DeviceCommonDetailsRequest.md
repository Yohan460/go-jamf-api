# DeviceCommonDetailsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientManagementId** | **string** | The client management ID associated with this device (required) | 
**RenewMdmProfileStartDate** | Pointer to **NullableTime** | Timestamp when MDM profile renewal started (ISO 8601 format) | [optional] 
**MdmProfileNeedsRenewalDueToCaRenewed** | Pointer to **NullableBool** | Whether the MDM profile needs renewal due to CA renewal | [optional] 
**MdmProfileNeedsRenewalDueToDeviceIdentityCertExpiring** | Pointer to **NullableBool** | Whether the MDM profile needs renewal due to expiring device identity certificate | [optional] 
**MdmCheckinUrl** | Pointer to **NullableString** | URL for MDM check-in | [optional] 
**MdmServerUrl** | Pointer to **NullableString** | URL for MDM server | [optional] 

## Methods

### NewDeviceCommonDetailsRequest

`func NewDeviceCommonDetailsRequest(clientManagementId string, ) *DeviceCommonDetailsRequest`

NewDeviceCommonDetailsRequest instantiates a new DeviceCommonDetailsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceCommonDetailsRequestWithDefaults

`func NewDeviceCommonDetailsRequestWithDefaults() *DeviceCommonDetailsRequest`

NewDeviceCommonDetailsRequestWithDefaults instantiates a new DeviceCommonDetailsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientManagementId

`func (o *DeviceCommonDetailsRequest) GetClientManagementId() string`

GetClientManagementId returns the ClientManagementId field if non-nil, zero value otherwise.

### GetClientManagementIdOk

`func (o *DeviceCommonDetailsRequest) GetClientManagementIdOk() (*string, bool)`

GetClientManagementIdOk returns a tuple with the ClientManagementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientManagementId

`func (o *DeviceCommonDetailsRequest) SetClientManagementId(v string)`

SetClientManagementId sets ClientManagementId field to given value.


### GetRenewMdmProfileStartDate

`func (o *DeviceCommonDetailsRequest) GetRenewMdmProfileStartDate() time.Time`

GetRenewMdmProfileStartDate returns the RenewMdmProfileStartDate field if non-nil, zero value otherwise.

### GetRenewMdmProfileStartDateOk

`func (o *DeviceCommonDetailsRequest) GetRenewMdmProfileStartDateOk() (*time.Time, bool)`

GetRenewMdmProfileStartDateOk returns a tuple with the RenewMdmProfileStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewMdmProfileStartDate

`func (o *DeviceCommonDetailsRequest) SetRenewMdmProfileStartDate(v time.Time)`

SetRenewMdmProfileStartDate sets RenewMdmProfileStartDate field to given value.

### HasRenewMdmProfileStartDate

`func (o *DeviceCommonDetailsRequest) HasRenewMdmProfileStartDate() bool`

HasRenewMdmProfileStartDate returns a boolean if a field has been set.

### SetRenewMdmProfileStartDateNil

`func (o *DeviceCommonDetailsRequest) SetRenewMdmProfileStartDateNil(b bool)`

 SetRenewMdmProfileStartDateNil sets the value for RenewMdmProfileStartDate to be an explicit nil

### UnsetRenewMdmProfileStartDate
`func (o *DeviceCommonDetailsRequest) UnsetRenewMdmProfileStartDate()`

UnsetRenewMdmProfileStartDate ensures that no value is present for RenewMdmProfileStartDate, not even an explicit nil
### GetMdmProfileNeedsRenewalDueToCaRenewed

`func (o *DeviceCommonDetailsRequest) GetMdmProfileNeedsRenewalDueToCaRenewed() bool`

GetMdmProfileNeedsRenewalDueToCaRenewed returns the MdmProfileNeedsRenewalDueToCaRenewed field if non-nil, zero value otherwise.

### GetMdmProfileNeedsRenewalDueToCaRenewedOk

`func (o *DeviceCommonDetailsRequest) GetMdmProfileNeedsRenewalDueToCaRenewedOk() (*bool, bool)`

GetMdmProfileNeedsRenewalDueToCaRenewedOk returns a tuple with the MdmProfileNeedsRenewalDueToCaRenewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdmProfileNeedsRenewalDueToCaRenewed

`func (o *DeviceCommonDetailsRequest) SetMdmProfileNeedsRenewalDueToCaRenewed(v bool)`

SetMdmProfileNeedsRenewalDueToCaRenewed sets MdmProfileNeedsRenewalDueToCaRenewed field to given value.

### HasMdmProfileNeedsRenewalDueToCaRenewed

`func (o *DeviceCommonDetailsRequest) HasMdmProfileNeedsRenewalDueToCaRenewed() bool`

HasMdmProfileNeedsRenewalDueToCaRenewed returns a boolean if a field has been set.

### SetMdmProfileNeedsRenewalDueToCaRenewedNil

`func (o *DeviceCommonDetailsRequest) SetMdmProfileNeedsRenewalDueToCaRenewedNil(b bool)`

 SetMdmProfileNeedsRenewalDueToCaRenewedNil sets the value for MdmProfileNeedsRenewalDueToCaRenewed to be an explicit nil

### UnsetMdmProfileNeedsRenewalDueToCaRenewed
`func (o *DeviceCommonDetailsRequest) UnsetMdmProfileNeedsRenewalDueToCaRenewed()`

UnsetMdmProfileNeedsRenewalDueToCaRenewed ensures that no value is present for MdmProfileNeedsRenewalDueToCaRenewed, not even an explicit nil
### GetMdmProfileNeedsRenewalDueToDeviceIdentityCertExpiring

`func (o *DeviceCommonDetailsRequest) GetMdmProfileNeedsRenewalDueToDeviceIdentityCertExpiring() bool`

GetMdmProfileNeedsRenewalDueToDeviceIdentityCertExpiring returns the MdmProfileNeedsRenewalDueToDeviceIdentityCertExpiring field if non-nil, zero value otherwise.

### GetMdmProfileNeedsRenewalDueToDeviceIdentityCertExpiringOk

`func (o *DeviceCommonDetailsRequest) GetMdmProfileNeedsRenewalDueToDeviceIdentityCertExpiringOk() (*bool, bool)`

GetMdmProfileNeedsRenewalDueToDeviceIdentityCertExpiringOk returns a tuple with the MdmProfileNeedsRenewalDueToDeviceIdentityCertExpiring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdmProfileNeedsRenewalDueToDeviceIdentityCertExpiring

`func (o *DeviceCommonDetailsRequest) SetMdmProfileNeedsRenewalDueToDeviceIdentityCertExpiring(v bool)`

SetMdmProfileNeedsRenewalDueToDeviceIdentityCertExpiring sets MdmProfileNeedsRenewalDueToDeviceIdentityCertExpiring field to given value.

### HasMdmProfileNeedsRenewalDueToDeviceIdentityCertExpiring

`func (o *DeviceCommonDetailsRequest) HasMdmProfileNeedsRenewalDueToDeviceIdentityCertExpiring() bool`

HasMdmProfileNeedsRenewalDueToDeviceIdentityCertExpiring returns a boolean if a field has been set.

### SetMdmProfileNeedsRenewalDueToDeviceIdentityCertExpiringNil

`func (o *DeviceCommonDetailsRequest) SetMdmProfileNeedsRenewalDueToDeviceIdentityCertExpiringNil(b bool)`

 SetMdmProfileNeedsRenewalDueToDeviceIdentityCertExpiringNil sets the value for MdmProfileNeedsRenewalDueToDeviceIdentityCertExpiring to be an explicit nil

### UnsetMdmProfileNeedsRenewalDueToDeviceIdentityCertExpiring
`func (o *DeviceCommonDetailsRequest) UnsetMdmProfileNeedsRenewalDueToDeviceIdentityCertExpiring()`

UnsetMdmProfileNeedsRenewalDueToDeviceIdentityCertExpiring ensures that no value is present for MdmProfileNeedsRenewalDueToDeviceIdentityCertExpiring, not even an explicit nil
### GetMdmCheckinUrl

`func (o *DeviceCommonDetailsRequest) GetMdmCheckinUrl() string`

GetMdmCheckinUrl returns the MdmCheckinUrl field if non-nil, zero value otherwise.

### GetMdmCheckinUrlOk

`func (o *DeviceCommonDetailsRequest) GetMdmCheckinUrlOk() (*string, bool)`

GetMdmCheckinUrlOk returns a tuple with the MdmCheckinUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdmCheckinUrl

`func (o *DeviceCommonDetailsRequest) SetMdmCheckinUrl(v string)`

SetMdmCheckinUrl sets MdmCheckinUrl field to given value.

### HasMdmCheckinUrl

`func (o *DeviceCommonDetailsRequest) HasMdmCheckinUrl() bool`

HasMdmCheckinUrl returns a boolean if a field has been set.

### SetMdmCheckinUrlNil

`func (o *DeviceCommonDetailsRequest) SetMdmCheckinUrlNil(b bool)`

 SetMdmCheckinUrlNil sets the value for MdmCheckinUrl to be an explicit nil

### UnsetMdmCheckinUrl
`func (o *DeviceCommonDetailsRequest) UnsetMdmCheckinUrl()`

UnsetMdmCheckinUrl ensures that no value is present for MdmCheckinUrl, not even an explicit nil
### GetMdmServerUrl

`func (o *DeviceCommonDetailsRequest) GetMdmServerUrl() string`

GetMdmServerUrl returns the MdmServerUrl field if non-nil, zero value otherwise.

### GetMdmServerUrlOk

`func (o *DeviceCommonDetailsRequest) GetMdmServerUrlOk() (*string, bool)`

GetMdmServerUrlOk returns a tuple with the MdmServerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdmServerUrl

`func (o *DeviceCommonDetailsRequest) SetMdmServerUrl(v string)`

SetMdmServerUrl sets MdmServerUrl field to given value.

### HasMdmServerUrl

`func (o *DeviceCommonDetailsRequest) HasMdmServerUrl() bool`

HasMdmServerUrl returns a boolean if a field has been set.

### SetMdmServerUrlNil

`func (o *DeviceCommonDetailsRequest) SetMdmServerUrlNil(b bool)`

 SetMdmServerUrlNil sets the value for MdmServerUrl to be an explicit nil

### UnsetMdmServerUrl
`func (o *DeviceCommonDetailsRequest) UnsetMdmServerUrl()`

UnsetMdmServerUrl ensures that no value is present for MdmServerUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


