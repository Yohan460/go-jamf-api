# ReturnToServiceConfigurationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | Display name of the Return to Service Configuration. | [optional] [default to "false"]
**WifiProfileId** | Pointer to **string** | Id of the wifi profile that is associated with the return to service configuration. | [optional] 

## Methods

### NewReturnToServiceConfigurationRequest

`func NewReturnToServiceConfigurationRequest() *ReturnToServiceConfigurationRequest`

NewReturnToServiceConfigurationRequest instantiates a new ReturnToServiceConfigurationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnToServiceConfigurationRequestWithDefaults

`func NewReturnToServiceConfigurationRequestWithDefaults() *ReturnToServiceConfigurationRequest`

NewReturnToServiceConfigurationRequestWithDefaults instantiates a new ReturnToServiceConfigurationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *ReturnToServiceConfigurationRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ReturnToServiceConfigurationRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ReturnToServiceConfigurationRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ReturnToServiceConfigurationRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetWifiProfileId

`func (o *ReturnToServiceConfigurationRequest) GetWifiProfileId() string`

GetWifiProfileId returns the WifiProfileId field if non-nil, zero value otherwise.

### GetWifiProfileIdOk

`func (o *ReturnToServiceConfigurationRequest) GetWifiProfileIdOk() (*string, bool)`

GetWifiProfileIdOk returns a tuple with the WifiProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiProfileId

`func (o *ReturnToServiceConfigurationRequest) SetWifiProfileId(v string)`

SetWifiProfileId sets WifiProfileId field to given value.

### HasWifiProfileId

`func (o *ReturnToServiceConfigurationRequest) HasWifiProfileId() bool`

HasWifiProfileId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


