# ComputerContentCachingParent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentCachingParentId** | Pointer to **string** |  | [optional] [readonly] 
**Address** | Pointer to **string** |  | [optional] [readonly] 
**Alerts** | Pointer to [**ComputerContentCachingParentAlert**](ComputerContentCachingParentAlert.md) |  | [optional] 
**Details** | Pointer to [**ComputerContentCachingParentDetails**](ComputerContentCachingParentDetails.md) |  | [optional] 
**Guid** | Pointer to **string** |  | [optional] [readonly] 
**Healthy** | Pointer to **bool** |  | [optional] [readonly] 
**Port** | Pointer to **int64** |  | [optional] [readonly] 
**Version** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewComputerContentCachingParent

`func NewComputerContentCachingParent() *ComputerContentCachingParent`

NewComputerContentCachingParent instantiates a new ComputerContentCachingParent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputerContentCachingParentWithDefaults

`func NewComputerContentCachingParentWithDefaults() *ComputerContentCachingParent`

NewComputerContentCachingParentWithDefaults instantiates a new ComputerContentCachingParent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentCachingParentId

`func (o *ComputerContentCachingParent) GetContentCachingParentId() string`

GetContentCachingParentId returns the ContentCachingParentId field if non-nil, zero value otherwise.

### GetContentCachingParentIdOk

`func (o *ComputerContentCachingParent) GetContentCachingParentIdOk() (*string, bool)`

GetContentCachingParentIdOk returns a tuple with the ContentCachingParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentCachingParentId

`func (o *ComputerContentCachingParent) SetContentCachingParentId(v string)`

SetContentCachingParentId sets ContentCachingParentId field to given value.

### HasContentCachingParentId

`func (o *ComputerContentCachingParent) HasContentCachingParentId() bool`

HasContentCachingParentId returns a boolean if a field has been set.

### GetAddress

`func (o *ComputerContentCachingParent) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ComputerContentCachingParent) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ComputerContentCachingParent) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ComputerContentCachingParent) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAlerts

`func (o *ComputerContentCachingParent) GetAlerts() ComputerContentCachingParentAlert`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *ComputerContentCachingParent) GetAlertsOk() (*ComputerContentCachingParentAlert, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *ComputerContentCachingParent) SetAlerts(v ComputerContentCachingParentAlert)`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *ComputerContentCachingParent) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.

### GetDetails

`func (o *ComputerContentCachingParent) GetDetails() ComputerContentCachingParentDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ComputerContentCachingParent) GetDetailsOk() (*ComputerContentCachingParentDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ComputerContentCachingParent) SetDetails(v ComputerContentCachingParentDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ComputerContentCachingParent) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetGuid

`func (o *ComputerContentCachingParent) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *ComputerContentCachingParent) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *ComputerContentCachingParent) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *ComputerContentCachingParent) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetHealthy

`func (o *ComputerContentCachingParent) GetHealthy() bool`

GetHealthy returns the Healthy field if non-nil, zero value otherwise.

### GetHealthyOk

`func (o *ComputerContentCachingParent) GetHealthyOk() (*bool, bool)`

GetHealthyOk returns a tuple with the Healthy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthy

`func (o *ComputerContentCachingParent) SetHealthy(v bool)`

SetHealthy sets Healthy field to given value.

### HasHealthy

`func (o *ComputerContentCachingParent) HasHealthy() bool`

HasHealthy returns a boolean if a field has been set.

### GetPort

`func (o *ComputerContentCachingParent) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ComputerContentCachingParent) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ComputerContentCachingParent) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *ComputerContentCachingParent) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetVersion

`func (o *ComputerContentCachingParent) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ComputerContentCachingParent) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ComputerContentCachingParent) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ComputerContentCachingParent) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


