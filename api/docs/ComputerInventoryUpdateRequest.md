# ComputerInventoryUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Udid** | Pointer to **string** |  | [optional] 
**General** | Pointer to [**ComputerGeneralUpdate**](ComputerGeneralUpdate.md) |  | [optional] 
**Purchasing** | Pointer to [**ComputerPurchase**](ComputerPurchase.md) |  | [optional] 
**UserAndLocation** | Pointer to [**ComputerUserAndLocation**](ComputerUserAndLocation.md) |  | [optional] 
**Hardware** | Pointer to [**ComputerHardwareUpdate**](ComputerHardwareUpdate.md) |  | [optional] 
**OperatingSystem** | Pointer to [**ComputerOperatingSystemUpdate**](ComputerOperatingSystemUpdate.md) |  | [optional] 
**ExtensionAttributes** | Pointer to [**[]ComputerExtensionAttribute**](ComputerExtensionAttribute.md) |  | [optional] 

## Methods

### NewComputerInventoryUpdateRequest

`func NewComputerInventoryUpdateRequest() *ComputerInventoryUpdateRequest`

NewComputerInventoryUpdateRequest instantiates a new ComputerInventoryUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputerInventoryUpdateRequestWithDefaults

`func NewComputerInventoryUpdateRequestWithDefaults() *ComputerInventoryUpdateRequest`

NewComputerInventoryUpdateRequestWithDefaults instantiates a new ComputerInventoryUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUdid

`func (o *ComputerInventoryUpdateRequest) GetUdid() string`

GetUdid returns the Udid field if non-nil, zero value otherwise.

### GetUdidOk

`func (o *ComputerInventoryUpdateRequest) GetUdidOk() (*string, bool)`

GetUdidOk returns a tuple with the Udid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdid

`func (o *ComputerInventoryUpdateRequest) SetUdid(v string)`

SetUdid sets Udid field to given value.

### HasUdid

`func (o *ComputerInventoryUpdateRequest) HasUdid() bool`

HasUdid returns a boolean if a field has been set.

### GetGeneral

`func (o *ComputerInventoryUpdateRequest) GetGeneral() ComputerGeneralUpdate`

GetGeneral returns the General field if non-nil, zero value otherwise.

### GetGeneralOk

`func (o *ComputerInventoryUpdateRequest) GetGeneralOk() (*ComputerGeneralUpdate, bool)`

GetGeneralOk returns a tuple with the General field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneral

`func (o *ComputerInventoryUpdateRequest) SetGeneral(v ComputerGeneralUpdate)`

SetGeneral sets General field to given value.

### HasGeneral

`func (o *ComputerInventoryUpdateRequest) HasGeneral() bool`

HasGeneral returns a boolean if a field has been set.

### GetPurchasing

`func (o *ComputerInventoryUpdateRequest) GetPurchasing() ComputerPurchase`

GetPurchasing returns the Purchasing field if non-nil, zero value otherwise.

### GetPurchasingOk

`func (o *ComputerInventoryUpdateRequest) GetPurchasingOk() (*ComputerPurchase, bool)`

GetPurchasingOk returns a tuple with the Purchasing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasing

`func (o *ComputerInventoryUpdateRequest) SetPurchasing(v ComputerPurchase)`

SetPurchasing sets Purchasing field to given value.

### HasPurchasing

`func (o *ComputerInventoryUpdateRequest) HasPurchasing() bool`

HasPurchasing returns a boolean if a field has been set.

### GetUserAndLocation

`func (o *ComputerInventoryUpdateRequest) GetUserAndLocation() ComputerUserAndLocation`

GetUserAndLocation returns the UserAndLocation field if non-nil, zero value otherwise.

### GetUserAndLocationOk

`func (o *ComputerInventoryUpdateRequest) GetUserAndLocationOk() (*ComputerUserAndLocation, bool)`

GetUserAndLocationOk returns a tuple with the UserAndLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAndLocation

`func (o *ComputerInventoryUpdateRequest) SetUserAndLocation(v ComputerUserAndLocation)`

SetUserAndLocation sets UserAndLocation field to given value.

### HasUserAndLocation

`func (o *ComputerInventoryUpdateRequest) HasUserAndLocation() bool`

HasUserAndLocation returns a boolean if a field has been set.

### GetHardware

`func (o *ComputerInventoryUpdateRequest) GetHardware() ComputerHardwareUpdate`

GetHardware returns the Hardware field if non-nil, zero value otherwise.

### GetHardwareOk

`func (o *ComputerInventoryUpdateRequest) GetHardwareOk() (*ComputerHardwareUpdate, bool)`

GetHardwareOk returns a tuple with the Hardware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardware

`func (o *ComputerInventoryUpdateRequest) SetHardware(v ComputerHardwareUpdate)`

SetHardware sets Hardware field to given value.

### HasHardware

`func (o *ComputerInventoryUpdateRequest) HasHardware() bool`

HasHardware returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *ComputerInventoryUpdateRequest) GetOperatingSystem() ComputerOperatingSystemUpdate`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *ComputerInventoryUpdateRequest) GetOperatingSystemOk() (*ComputerOperatingSystemUpdate, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *ComputerInventoryUpdateRequest) SetOperatingSystem(v ComputerOperatingSystemUpdate)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *ComputerInventoryUpdateRequest) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetExtensionAttributes

`func (o *ComputerInventoryUpdateRequest) GetExtensionAttributes() []ComputerExtensionAttribute`

GetExtensionAttributes returns the ExtensionAttributes field if non-nil, zero value otherwise.

### GetExtensionAttributesOk

`func (o *ComputerInventoryUpdateRequest) GetExtensionAttributesOk() (*[]ComputerExtensionAttribute, bool)`

GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionAttributes

`func (o *ComputerInventoryUpdateRequest) SetExtensionAttributes(v []ComputerExtensionAttribute)`

SetExtensionAttributes sets ExtensionAttributes field to given value.

### HasExtensionAttributes

`func (o *ComputerInventoryUpdateRequest) HasExtensionAttributes() bool`

HasExtensionAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


