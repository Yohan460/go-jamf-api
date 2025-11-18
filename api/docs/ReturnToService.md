# ReturnToService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** |  | 
**MdmProfileData** | Pointer to **string** | Base64 encoded mdm profile | [optional] 
**WifiProfileData** | Pointer to **string** | Base64 encoded wifi profile | [optional] 
**BootstrapToken** | Pointer to **string** | Base64 encoded bootstrap token for the device. | [optional] 

## Methods

### NewReturnToService

`func NewReturnToService(enabled bool, ) *ReturnToService`

NewReturnToService instantiates a new ReturnToService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnToServiceWithDefaults

`func NewReturnToServiceWithDefaults() *ReturnToService`

NewReturnToServiceWithDefaults instantiates a new ReturnToService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ReturnToService) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ReturnToService) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ReturnToService) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMdmProfileData

`func (o *ReturnToService) GetMdmProfileData() string`

GetMdmProfileData returns the MdmProfileData field if non-nil, zero value otherwise.

### GetMdmProfileDataOk

`func (o *ReturnToService) GetMdmProfileDataOk() (*string, bool)`

GetMdmProfileDataOk returns a tuple with the MdmProfileData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdmProfileData

`func (o *ReturnToService) SetMdmProfileData(v string)`

SetMdmProfileData sets MdmProfileData field to given value.

### HasMdmProfileData

`func (o *ReturnToService) HasMdmProfileData() bool`

HasMdmProfileData returns a boolean if a field has been set.

### GetWifiProfileData

`func (o *ReturnToService) GetWifiProfileData() string`

GetWifiProfileData returns the WifiProfileData field if non-nil, zero value otherwise.

### GetWifiProfileDataOk

`func (o *ReturnToService) GetWifiProfileDataOk() (*string, bool)`

GetWifiProfileDataOk returns a tuple with the WifiProfileData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiProfileData

`func (o *ReturnToService) SetWifiProfileData(v string)`

SetWifiProfileData sets WifiProfileData field to given value.

### HasWifiProfileData

`func (o *ReturnToService) HasWifiProfileData() bool`

HasWifiProfileData returns a boolean if a field has been set.

### GetBootstrapToken

`func (o *ReturnToService) GetBootstrapToken() string`

GetBootstrapToken returns the BootstrapToken field if non-nil, zero value otherwise.

### GetBootstrapTokenOk

`func (o *ReturnToService) GetBootstrapTokenOk() (*string, bool)`

GetBootstrapTokenOk returns a tuple with the BootstrapToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootstrapToken

`func (o *ReturnToService) SetBootstrapToken(v string)`

SetBootstrapToken sets BootstrapToken field to given value.

### HasBootstrapToken

`func (o *ReturnToService) HasBootstrapToken() bool`

HasBootstrapToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


