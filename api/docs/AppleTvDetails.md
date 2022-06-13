# AppleTvDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | Pointer to **string** |  | [optional] 
**ModelIdentifier** | Pointer to **string** |  | [optional] 
**ModelNumber** | Pointer to **string** |  | [optional] 
**IsSupervised** | Pointer to **bool** |  | [optional] 
**AirplayPassword** | Pointer to **string** |  | [optional] 
**DeviceId** | Pointer to **string** |  | [optional] 
**Locales** | Pointer to **string** |  | [optional] 
**Purchasing** | Pointer to [**Purchasing**](Purchasing.md) |  | [optional] 
**ConfigurationProfiles** | Pointer to [**[]ConfigurationProfile**](ConfigurationProfile.md) |  | [optional] 

## Methods

### NewAppleTvDetails

`func NewAppleTvDetails() *AppleTvDetails`

NewAppleTvDetails instantiates a new AppleTvDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppleTvDetailsWithDefaults

`func NewAppleTvDetailsWithDefaults() *AppleTvDetails`

NewAppleTvDetailsWithDefaults instantiates a new AppleTvDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *AppleTvDetails) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AppleTvDetails) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AppleTvDetails) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *AppleTvDetails) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelIdentifier

`func (o *AppleTvDetails) GetModelIdentifier() string`

GetModelIdentifier returns the ModelIdentifier field if non-nil, zero value otherwise.

### GetModelIdentifierOk

`func (o *AppleTvDetails) GetModelIdentifierOk() (*string, bool)`

GetModelIdentifierOk returns a tuple with the ModelIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelIdentifier

`func (o *AppleTvDetails) SetModelIdentifier(v string)`

SetModelIdentifier sets ModelIdentifier field to given value.

### HasModelIdentifier

`func (o *AppleTvDetails) HasModelIdentifier() bool`

HasModelIdentifier returns a boolean if a field has been set.

### GetModelNumber

`func (o *AppleTvDetails) GetModelNumber() string`

GetModelNumber returns the ModelNumber field if non-nil, zero value otherwise.

### GetModelNumberOk

`func (o *AppleTvDetails) GetModelNumberOk() (*string, bool)`

GetModelNumberOk returns a tuple with the ModelNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelNumber

`func (o *AppleTvDetails) SetModelNumber(v string)`

SetModelNumber sets ModelNumber field to given value.

### HasModelNumber

`func (o *AppleTvDetails) HasModelNumber() bool`

HasModelNumber returns a boolean if a field has been set.

### GetIsSupervised

`func (o *AppleTvDetails) GetIsSupervised() bool`

GetIsSupervised returns the IsSupervised field if non-nil, zero value otherwise.

### GetIsSupervisedOk

`func (o *AppleTvDetails) GetIsSupervisedOk() (*bool, bool)`

GetIsSupervisedOk returns a tuple with the IsSupervised field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSupervised

`func (o *AppleTvDetails) SetIsSupervised(v bool)`

SetIsSupervised sets IsSupervised field to given value.

### HasIsSupervised

`func (o *AppleTvDetails) HasIsSupervised() bool`

HasIsSupervised returns a boolean if a field has been set.

### GetAirplayPassword

`func (o *AppleTvDetails) GetAirplayPassword() string`

GetAirplayPassword returns the AirplayPassword field if non-nil, zero value otherwise.

### GetAirplayPasswordOk

`func (o *AppleTvDetails) GetAirplayPasswordOk() (*string, bool)`

GetAirplayPasswordOk returns a tuple with the AirplayPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirplayPassword

`func (o *AppleTvDetails) SetAirplayPassword(v string)`

SetAirplayPassword sets AirplayPassword field to given value.

### HasAirplayPassword

`func (o *AppleTvDetails) HasAirplayPassword() bool`

HasAirplayPassword returns a boolean if a field has been set.

### GetDeviceId

`func (o *AppleTvDetails) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *AppleTvDetails) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *AppleTvDetails) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *AppleTvDetails) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetLocales

`func (o *AppleTvDetails) GetLocales() string`

GetLocales returns the Locales field if non-nil, zero value otherwise.

### GetLocalesOk

`func (o *AppleTvDetails) GetLocalesOk() (*string, bool)`

GetLocalesOk returns a tuple with the Locales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocales

`func (o *AppleTvDetails) SetLocales(v string)`

SetLocales sets Locales field to given value.

### HasLocales

`func (o *AppleTvDetails) HasLocales() bool`

HasLocales returns a boolean if a field has been set.

### GetPurchasing

`func (o *AppleTvDetails) GetPurchasing() Purchasing`

GetPurchasing returns the Purchasing field if non-nil, zero value otherwise.

### GetPurchasingOk

`func (o *AppleTvDetails) GetPurchasingOk() (*Purchasing, bool)`

GetPurchasingOk returns a tuple with the Purchasing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasing

`func (o *AppleTvDetails) SetPurchasing(v Purchasing)`

SetPurchasing sets Purchasing field to given value.

### HasPurchasing

`func (o *AppleTvDetails) HasPurchasing() bool`

HasPurchasing returns a boolean if a field has been set.

### GetConfigurationProfiles

`func (o *AppleTvDetails) GetConfigurationProfiles() []ConfigurationProfile`

GetConfigurationProfiles returns the ConfigurationProfiles field if non-nil, zero value otherwise.

### GetConfigurationProfilesOk

`func (o *AppleTvDetails) GetConfigurationProfilesOk() (*[]ConfigurationProfile, bool)`

GetConfigurationProfilesOk returns a tuple with the ConfigurationProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationProfiles

`func (o *AppleTvDetails) SetConfigurationProfiles(v []ConfigurationProfile)`

SetConfigurationProfiles sets ConfigurationProfiles field to given value.

### HasConfigurationProfiles

`func (o *AppleTvDetails) HasConfigurationProfiles() bool`

HasConfigurationProfiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


