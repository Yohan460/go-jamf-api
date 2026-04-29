# DeviceCommonDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for the device common details record | [optional] 
**ClientManagementId** | **string** | The client management ID associated with this device | 
**RenewMdmProfileStartDate** | Pointer to **NullableTime** | Timestamp when MDM profile renewal started (ISO 8601 format) | [optional] 
**MdmProfileNeedsRenewalDueToCaRenewed** | Pointer to **bool** | Whether the MDM profile needs renewal due to CA renewal | [optional] 
**MdmProfileNeedsRenewalDueToDeviceIdentityCertExpiring** | Pointer to **bool** | Whether the MDM profile needs renewal due to expiring device identity certificate | [optional] 
**MdmCheckinUrl** | Pointer to **NullableString** | URL for MDM check-in | [optional] 
**MdmServerUrl** | Pointer to **NullableString** | URL for MDM server | [optional] 

## Methods

### NewDeviceCommonDetails

`func NewDeviceCommonDetails(clientManagementId string, ) *DeviceCommonDetails`

NewDeviceCommonDetails instantiates a new DeviceCommonDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceCommonDetailsWithDefaults

`func NewDeviceCommonDetailsWithDefaults() *DeviceCommonDetails`

NewDeviceCommonDetailsWithDefaults instantiates a new DeviceCommonDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeviceCommonDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceCommonDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceCommonDetails) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeviceCommonDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetClientManagementId

`func (o *DeviceCommonDetails) GetClientManagementId() string`

GetClientManagementId returns the ClientManagementId field if non-nil, zero value otherwise.

### GetClientManagementIdOk

`func (o *DeviceCommonDetails) GetClientManagementIdOk() (*string, bool)`

GetClientManagementIdOk returns a tuple with the ClientManagementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientManagementId

`func (o *DeviceCommonDetails) SetClientManagementId(v string)`

SetClientManagementId sets ClientManagementId field to given value.


### GetRenewMdmProfileStartDate

`func (o *DeviceCommonDetails) GetRenewMdmProfileStartDate() time.Time`

GetRenewMdmProfileStartDate returns the RenewMdmProfileStartDate field if non-nil, zero value otherwise.

### GetRenewMdmProfileStartDateOk

`func (o *DeviceCommonDetails) GetRenewMdmProfileStartDateOk() (*time.Time, bool)`

GetRenewMdmProfileStartDateOk returns a tuple with the RenewMdmProfileStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewMdmProfileStartDate

`func (o *DeviceCommonDetails) SetRenewMdmProfileStartDate(v time.Time)`

SetRenewMdmProfileStartDate sets RenewMdmProfileStartDate field to given value.

### HasRenewMdmProfileStartDate

`func (o *DeviceCommonDetails) HasRenewMdmProfileStartDate() bool`

HasRenewMdmProfileStartDate returns a boolean if a field has been set.

### SetRenewMdmProfileStartDateNil

`func (o *DeviceCommonDetails) SetRenewMdmProfileStartDateNil(b bool)`

 SetRenewMdmProfileStartDateNil sets the value for RenewMdmProfileStartDate to be an explicit nil

### UnsetRenewMdmProfileStartDate
`func (o *DeviceCommonDetails) UnsetRenewMdmProfileStartDate()`

UnsetRenewMdmProfileStartDate ensures that no value is present for RenewMdmProfileStartDate, not even an explicit nil
### GetMdmProfileNeedsRenewalDueToCaRenewed

`func (o *DeviceCommonDetails) GetMdmProfileNeedsRenewalDueToCaRenewed() bool`

GetMdmProfileNeedsRenewalDueToCaRenewed returns the MdmProfileNeedsRenewalDueToCaRenewed field if non-nil, zero value otherwise.

### GetMdmProfileNeedsRenewalDueToCaRenewedOk

`func (o *DeviceCommonDetails) GetMdmProfileNeedsRenewalDueToCaRenewedOk() (*bool, bool)`

GetMdmProfileNeedsRenewalDueToCaRenewedOk returns a tuple with the MdmProfileNeedsRenewalDueToCaRenewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdmProfileNeedsRenewalDueToCaRenewed

`func (o *DeviceCommonDetails) SetMdmProfileNeedsRenewalDueToCaRenewed(v bool)`

SetMdmProfileNeedsRenewalDueToCaRenewed sets MdmProfileNeedsRenewalDueToCaRenewed field to given value.

### HasMdmProfileNeedsRenewalDueToCaRenewed

`func (o *DeviceCommonDetails) HasMdmProfileNeedsRenewalDueToCaRenewed() bool`

HasMdmProfileNeedsRenewalDueToCaRenewed returns a boolean if a field has been set.

### GetMdmProfileNeedsRenewalDueToDeviceIdentityCertExpiring

`func (o *DeviceCommonDetails) GetMdmProfileNeedsRenewalDueToDeviceIdentityCertExpiring() bool`

GetMdmProfileNeedsRenewalDueToDeviceIdentityCertExpiring returns the MdmProfileNeedsRenewalDueToDeviceIdentityCertExpiring field if non-nil, zero value otherwise.

### GetMdmProfileNeedsRenewalDueToDeviceIdentityCertExpiringOk

`func (o *DeviceCommonDetails) GetMdmProfileNeedsRenewalDueToDeviceIdentityCertExpiringOk() (*bool, bool)`

GetMdmProfileNeedsRenewalDueToDeviceIdentityCertExpiringOk returns a tuple with the MdmProfileNeedsRenewalDueToDeviceIdentityCertExpiring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdmProfileNeedsRenewalDueToDeviceIdentityCertExpiring

`func (o *DeviceCommonDetails) SetMdmProfileNeedsRenewalDueToDeviceIdentityCertExpiring(v bool)`

SetMdmProfileNeedsRenewalDueToDeviceIdentityCertExpiring sets MdmProfileNeedsRenewalDueToDeviceIdentityCertExpiring field to given value.

### HasMdmProfileNeedsRenewalDueToDeviceIdentityCertExpiring

`func (o *DeviceCommonDetails) HasMdmProfileNeedsRenewalDueToDeviceIdentityCertExpiring() bool`

HasMdmProfileNeedsRenewalDueToDeviceIdentityCertExpiring returns a boolean if a field has been set.

### GetMdmCheckinUrl

`func (o *DeviceCommonDetails) GetMdmCheckinUrl() string`

GetMdmCheckinUrl returns the MdmCheckinUrl field if non-nil, zero value otherwise.

### GetMdmCheckinUrlOk

`func (o *DeviceCommonDetails) GetMdmCheckinUrlOk() (*string, bool)`

GetMdmCheckinUrlOk returns a tuple with the MdmCheckinUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdmCheckinUrl

`func (o *DeviceCommonDetails) SetMdmCheckinUrl(v string)`

SetMdmCheckinUrl sets MdmCheckinUrl field to given value.

### HasMdmCheckinUrl

`func (o *DeviceCommonDetails) HasMdmCheckinUrl() bool`

HasMdmCheckinUrl returns a boolean if a field has been set.

### SetMdmCheckinUrlNil

`func (o *DeviceCommonDetails) SetMdmCheckinUrlNil(b bool)`

 SetMdmCheckinUrlNil sets the value for MdmCheckinUrl to be an explicit nil

### UnsetMdmCheckinUrl
`func (o *DeviceCommonDetails) UnsetMdmCheckinUrl()`

UnsetMdmCheckinUrl ensures that no value is present for MdmCheckinUrl, not even an explicit nil
### GetMdmServerUrl

`func (o *DeviceCommonDetails) GetMdmServerUrl() string`

GetMdmServerUrl returns the MdmServerUrl field if non-nil, zero value otherwise.

### GetMdmServerUrlOk

`func (o *DeviceCommonDetails) GetMdmServerUrlOk() (*string, bool)`

GetMdmServerUrlOk returns a tuple with the MdmServerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdmServerUrl

`func (o *DeviceCommonDetails) SetMdmServerUrl(v string)`

SetMdmServerUrl sets MdmServerUrl field to given value.

### HasMdmServerUrl

`func (o *DeviceCommonDetails) HasMdmServerUrl() bool`

HasMdmServerUrl returns a boolean if a field has been set.

### SetMdmServerUrlNil

`func (o *DeviceCommonDetails) SetMdmServerUrlNil(b bool)`

 SetMdmServerUrlNil sets the value for MdmServerUrl to be an explicit nil

### UnsetMdmServerUrl
`func (o *DeviceCommonDetails) UnsetMdmServerUrl()`

UnsetMdmServerUrl ensures that no value is present for MdmServerUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


