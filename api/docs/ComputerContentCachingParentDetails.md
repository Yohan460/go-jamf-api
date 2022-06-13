# ComputerContentCachingParentDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentCachingParentDetailsId** | Pointer to **string** |  | [optional] [readonly] 
**AcPower** | Pointer to **bool** |  | [optional] [readonly] 
**CacheSizeBytes** | Pointer to **int64** |  | [optional] [readonly] 
**Capabilities** | Pointer to [**ComputerContentCachingParentCapabilities**](ComputerContentCachingParentCapabilities.md) |  | [optional] 
**Portable** | Pointer to **bool** |  | [optional] [readonly] 
**LocalNetwork** | Pointer to [**[]ComputerContentCachingParentLocalNetwork**](ComputerContentCachingParentLocalNetwork.md) |  | [optional] 

## Methods

### NewComputerContentCachingParentDetails

`func NewComputerContentCachingParentDetails() *ComputerContentCachingParentDetails`

NewComputerContentCachingParentDetails instantiates a new ComputerContentCachingParentDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputerContentCachingParentDetailsWithDefaults

`func NewComputerContentCachingParentDetailsWithDefaults() *ComputerContentCachingParentDetails`

NewComputerContentCachingParentDetailsWithDefaults instantiates a new ComputerContentCachingParentDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentCachingParentDetailsId

`func (o *ComputerContentCachingParentDetails) GetContentCachingParentDetailsId() string`

GetContentCachingParentDetailsId returns the ContentCachingParentDetailsId field if non-nil, zero value otherwise.

### GetContentCachingParentDetailsIdOk

`func (o *ComputerContentCachingParentDetails) GetContentCachingParentDetailsIdOk() (*string, bool)`

GetContentCachingParentDetailsIdOk returns a tuple with the ContentCachingParentDetailsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentCachingParentDetailsId

`func (o *ComputerContentCachingParentDetails) SetContentCachingParentDetailsId(v string)`

SetContentCachingParentDetailsId sets ContentCachingParentDetailsId field to given value.

### HasContentCachingParentDetailsId

`func (o *ComputerContentCachingParentDetails) HasContentCachingParentDetailsId() bool`

HasContentCachingParentDetailsId returns a boolean if a field has been set.

### GetAcPower

`func (o *ComputerContentCachingParentDetails) GetAcPower() bool`

GetAcPower returns the AcPower field if non-nil, zero value otherwise.

### GetAcPowerOk

`func (o *ComputerContentCachingParentDetails) GetAcPowerOk() (*bool, bool)`

GetAcPowerOk returns a tuple with the AcPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcPower

`func (o *ComputerContentCachingParentDetails) SetAcPower(v bool)`

SetAcPower sets AcPower field to given value.

### HasAcPower

`func (o *ComputerContentCachingParentDetails) HasAcPower() bool`

HasAcPower returns a boolean if a field has been set.

### GetCacheSizeBytes

`func (o *ComputerContentCachingParentDetails) GetCacheSizeBytes() int64`

GetCacheSizeBytes returns the CacheSizeBytes field if non-nil, zero value otherwise.

### GetCacheSizeBytesOk

`func (o *ComputerContentCachingParentDetails) GetCacheSizeBytesOk() (*int64, bool)`

GetCacheSizeBytesOk returns a tuple with the CacheSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheSizeBytes

`func (o *ComputerContentCachingParentDetails) SetCacheSizeBytes(v int64)`

SetCacheSizeBytes sets CacheSizeBytes field to given value.

### HasCacheSizeBytes

`func (o *ComputerContentCachingParentDetails) HasCacheSizeBytes() bool`

HasCacheSizeBytes returns a boolean if a field has been set.

### GetCapabilities

`func (o *ComputerContentCachingParentDetails) GetCapabilities() ComputerContentCachingParentCapabilities`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *ComputerContentCachingParentDetails) GetCapabilitiesOk() (*ComputerContentCachingParentCapabilities, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *ComputerContentCachingParentDetails) SetCapabilities(v ComputerContentCachingParentCapabilities)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *ComputerContentCachingParentDetails) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetPortable

`func (o *ComputerContentCachingParentDetails) GetPortable() bool`

GetPortable returns the Portable field if non-nil, zero value otherwise.

### GetPortableOk

`func (o *ComputerContentCachingParentDetails) GetPortableOk() (*bool, bool)`

GetPortableOk returns a tuple with the Portable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortable

`func (o *ComputerContentCachingParentDetails) SetPortable(v bool)`

SetPortable sets Portable field to given value.

### HasPortable

`func (o *ComputerContentCachingParentDetails) HasPortable() bool`

HasPortable returns a boolean if a field has been set.

### GetLocalNetwork

`func (o *ComputerContentCachingParentDetails) GetLocalNetwork() []ComputerContentCachingParentLocalNetwork`

GetLocalNetwork returns the LocalNetwork field if non-nil, zero value otherwise.

### GetLocalNetworkOk

`func (o *ComputerContentCachingParentDetails) GetLocalNetworkOk() (*[]ComputerContentCachingParentLocalNetwork, bool)`

GetLocalNetworkOk returns a tuple with the LocalNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalNetwork

`func (o *ComputerContentCachingParentDetails) SetLocalNetwork(v []ComputerContentCachingParentLocalNetwork)`

SetLocalNetwork sets LocalNetwork field to given value.

### HasLocalNetwork

`func (o *ComputerContentCachingParentDetails) HasLocalNetwork() bool`

HasLocalNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


