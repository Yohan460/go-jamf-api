# EraseDeviceCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PreserveDataPlan** | Pointer to **bool** | If true, preserve the data plan on an iPhone or iPad with eSIM functionality, if one exists. This value is available in iOS 11 and later. | [optional] [default to false]
**DisallowProximitySetup** | Pointer to **bool** | If true, disable Proximity Setup on the next reboot and skip the pane in Setup Assistant. This value is available in iOS 11 and later. Prior to iOS 14, don’t use this option with any other option. | [optional] [default to false]
**Pin** | Pointer to **string** | The six-character PIN for Find My. This value is available in macOS 10.8 and later. | [optional] 
**ObliterationBehavior** | Pointer to **string** | This key defines the fallback behavior for erasing a device. | [optional] 
**ReturnToService** | Pointer to [**EraseDeviceCommandReturnToService**](EraseDeviceCommandReturnToService.md) |  | [optional] 

## Methods

### NewEraseDeviceCommand

`func NewEraseDeviceCommand() *EraseDeviceCommand`

NewEraseDeviceCommand instantiates a new EraseDeviceCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEraseDeviceCommandWithDefaults

`func NewEraseDeviceCommandWithDefaults() *EraseDeviceCommand`

NewEraseDeviceCommandWithDefaults instantiates a new EraseDeviceCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreserveDataPlan

`func (o *EraseDeviceCommand) GetPreserveDataPlan() bool`

GetPreserveDataPlan returns the PreserveDataPlan field if non-nil, zero value otherwise.

### GetPreserveDataPlanOk

`func (o *EraseDeviceCommand) GetPreserveDataPlanOk() (*bool, bool)`

GetPreserveDataPlanOk returns a tuple with the PreserveDataPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveDataPlan

`func (o *EraseDeviceCommand) SetPreserveDataPlan(v bool)`

SetPreserveDataPlan sets PreserveDataPlan field to given value.

### HasPreserveDataPlan

`func (o *EraseDeviceCommand) HasPreserveDataPlan() bool`

HasPreserveDataPlan returns a boolean if a field has been set.

### GetDisallowProximitySetup

`func (o *EraseDeviceCommand) GetDisallowProximitySetup() bool`

GetDisallowProximitySetup returns the DisallowProximitySetup field if non-nil, zero value otherwise.

### GetDisallowProximitySetupOk

`func (o *EraseDeviceCommand) GetDisallowProximitySetupOk() (*bool, bool)`

GetDisallowProximitySetupOk returns a tuple with the DisallowProximitySetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisallowProximitySetup

`func (o *EraseDeviceCommand) SetDisallowProximitySetup(v bool)`

SetDisallowProximitySetup sets DisallowProximitySetup field to given value.

### HasDisallowProximitySetup

`func (o *EraseDeviceCommand) HasDisallowProximitySetup() bool`

HasDisallowProximitySetup returns a boolean if a field has been set.

### GetPin

`func (o *EraseDeviceCommand) GetPin() string`

GetPin returns the Pin field if non-nil, zero value otherwise.

### GetPinOk

`func (o *EraseDeviceCommand) GetPinOk() (*string, bool)`

GetPinOk returns a tuple with the Pin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPin

`func (o *EraseDeviceCommand) SetPin(v string)`

SetPin sets Pin field to given value.

### HasPin

`func (o *EraseDeviceCommand) HasPin() bool`

HasPin returns a boolean if a field has been set.

### GetObliterationBehavior

`func (o *EraseDeviceCommand) GetObliterationBehavior() string`

GetObliterationBehavior returns the ObliterationBehavior field if non-nil, zero value otherwise.

### GetObliterationBehaviorOk

`func (o *EraseDeviceCommand) GetObliterationBehaviorOk() (*string, bool)`

GetObliterationBehaviorOk returns a tuple with the ObliterationBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObliterationBehavior

`func (o *EraseDeviceCommand) SetObliterationBehavior(v string)`

SetObliterationBehavior sets ObliterationBehavior field to given value.

### HasObliterationBehavior

`func (o *EraseDeviceCommand) HasObliterationBehavior() bool`

HasObliterationBehavior returns a boolean if a field has been set.

### GetReturnToService

`func (o *EraseDeviceCommand) GetReturnToService() EraseDeviceCommandReturnToService`

GetReturnToService returns the ReturnToService field if non-nil, zero value otherwise.

### GetReturnToServiceOk

`func (o *EraseDeviceCommand) GetReturnToServiceOk() (*EraseDeviceCommandReturnToService, bool)`

GetReturnToServiceOk returns a tuple with the ReturnToService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnToService

`func (o *EraseDeviceCommand) SetReturnToService(v EraseDeviceCommandReturnToService)`

SetReturnToService sets ReturnToService field to given value.

### HasReturnToService

`func (o *EraseDeviceCommand) HasReturnToService() bool`

HasReturnToService returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


