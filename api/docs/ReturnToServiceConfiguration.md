# ReturnToServiceConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Id of the Return to Service Configuration. | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] [default to "false"]
**WifiProfileId** | Pointer to **string** | Id of the wifi profile that is associated with the return to service configuration. | [optional] 

## Methods

### NewReturnToServiceConfiguration

`func NewReturnToServiceConfiguration() *ReturnToServiceConfiguration`

NewReturnToServiceConfiguration instantiates a new ReturnToServiceConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnToServiceConfigurationWithDefaults

`func NewReturnToServiceConfigurationWithDefaults() *ReturnToServiceConfiguration`

NewReturnToServiceConfigurationWithDefaults instantiates a new ReturnToServiceConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReturnToServiceConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReturnToServiceConfiguration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReturnToServiceConfiguration) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ReturnToServiceConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *ReturnToServiceConfiguration) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ReturnToServiceConfiguration) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ReturnToServiceConfiguration) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ReturnToServiceConfiguration) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetWifiProfileId

`func (o *ReturnToServiceConfiguration) GetWifiProfileId() string`

GetWifiProfileId returns the WifiProfileId field if non-nil, zero value otherwise.

### GetWifiProfileIdOk

`func (o *ReturnToServiceConfiguration) GetWifiProfileIdOk() (*string, bool)`

GetWifiProfileIdOk returns a tuple with the WifiProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiProfileId

`func (o *ReturnToServiceConfiguration) SetWifiProfileId(v string)`

SetWifiProfileId sets WifiProfileId field to given value.

### HasWifiProfileId

`func (o *ReturnToServiceConfiguration) HasWifiProfileId() bool`

HasWifiProfileId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


