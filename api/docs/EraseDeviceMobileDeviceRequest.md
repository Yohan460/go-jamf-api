# EraseDeviceMobileDeviceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PreserveDataPlan** | Pointer to **bool** | If &#39;true&#39;, preserve the data plan on an iPhone or iPad with eSIM functionality, if one exists. | [optional] [default to false]
**DisallowProximitySetup** | Pointer to **bool** | If &#39;true&#39;, disable Proximity Setup on the next reboot and skip the pane in Setup Assistant. | [optional] [default to false]
**ClearActivationLock** | Pointer to **bool** | Clear the activation lock on the device. | [optional] [default to false]
**ReturnToService** | Pointer to **bool** | If &#39;true&#39;, the device will be returned to service after the erase is complete. | [optional] [default to false]

## Methods

### NewEraseDeviceMobileDeviceRequest

`func NewEraseDeviceMobileDeviceRequest() *EraseDeviceMobileDeviceRequest`

NewEraseDeviceMobileDeviceRequest instantiates a new EraseDeviceMobileDeviceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEraseDeviceMobileDeviceRequestWithDefaults

`func NewEraseDeviceMobileDeviceRequestWithDefaults() *EraseDeviceMobileDeviceRequest`

NewEraseDeviceMobileDeviceRequestWithDefaults instantiates a new EraseDeviceMobileDeviceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreserveDataPlan

`func (o *EraseDeviceMobileDeviceRequest) GetPreserveDataPlan() bool`

GetPreserveDataPlan returns the PreserveDataPlan field if non-nil, zero value otherwise.

### GetPreserveDataPlanOk

`func (o *EraseDeviceMobileDeviceRequest) GetPreserveDataPlanOk() (*bool, bool)`

GetPreserveDataPlanOk returns a tuple with the PreserveDataPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveDataPlan

`func (o *EraseDeviceMobileDeviceRequest) SetPreserveDataPlan(v bool)`

SetPreserveDataPlan sets PreserveDataPlan field to given value.

### HasPreserveDataPlan

`func (o *EraseDeviceMobileDeviceRequest) HasPreserveDataPlan() bool`

HasPreserveDataPlan returns a boolean if a field has been set.

### GetDisallowProximitySetup

`func (o *EraseDeviceMobileDeviceRequest) GetDisallowProximitySetup() bool`

GetDisallowProximitySetup returns the DisallowProximitySetup field if non-nil, zero value otherwise.

### GetDisallowProximitySetupOk

`func (o *EraseDeviceMobileDeviceRequest) GetDisallowProximitySetupOk() (*bool, bool)`

GetDisallowProximitySetupOk returns a tuple with the DisallowProximitySetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisallowProximitySetup

`func (o *EraseDeviceMobileDeviceRequest) SetDisallowProximitySetup(v bool)`

SetDisallowProximitySetup sets DisallowProximitySetup field to given value.

### HasDisallowProximitySetup

`func (o *EraseDeviceMobileDeviceRequest) HasDisallowProximitySetup() bool`

HasDisallowProximitySetup returns a boolean if a field has been set.

### GetClearActivationLock

`func (o *EraseDeviceMobileDeviceRequest) GetClearActivationLock() bool`

GetClearActivationLock returns the ClearActivationLock field if non-nil, zero value otherwise.

### GetClearActivationLockOk

`func (o *EraseDeviceMobileDeviceRequest) GetClearActivationLockOk() (*bool, bool)`

GetClearActivationLockOk returns a tuple with the ClearActivationLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearActivationLock

`func (o *EraseDeviceMobileDeviceRequest) SetClearActivationLock(v bool)`

SetClearActivationLock sets ClearActivationLock field to given value.

### HasClearActivationLock

`func (o *EraseDeviceMobileDeviceRequest) HasClearActivationLock() bool`

HasClearActivationLock returns a boolean if a field has been set.

### GetReturnToService

`func (o *EraseDeviceMobileDeviceRequest) GetReturnToService() bool`

GetReturnToService returns the ReturnToService field if non-nil, zero value otherwise.

### GetReturnToServiceOk

`func (o *EraseDeviceMobileDeviceRequest) GetReturnToServiceOk() (*bool, bool)`

GetReturnToServiceOk returns a tuple with the ReturnToService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnToService

`func (o *EraseDeviceMobileDeviceRequest) SetReturnToService(v bool)`

SetReturnToService sets ReturnToService field to given value.

### HasReturnToService

`func (o *EraseDeviceMobileDeviceRequest) HasReturnToService() bool`

HasReturnToService returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


