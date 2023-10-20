# ManagedSoftwareUpdatePlanPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Devices** | [**[]PlanDevicePost**](PlanDevicePost.md) |  | 
**Config** | [**PlanConfigurationPost**](PlanConfigurationPost.md) |  | 

## Methods

### NewManagedSoftwareUpdatePlanPost

`func NewManagedSoftwareUpdatePlanPost(devices []PlanDevicePost, config PlanConfigurationPost, ) *ManagedSoftwareUpdatePlanPost`

NewManagedSoftwareUpdatePlanPost instantiates a new ManagedSoftwareUpdatePlanPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedSoftwareUpdatePlanPostWithDefaults

`func NewManagedSoftwareUpdatePlanPostWithDefaults() *ManagedSoftwareUpdatePlanPost`

NewManagedSoftwareUpdatePlanPostWithDefaults instantiates a new ManagedSoftwareUpdatePlanPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevices

`func (o *ManagedSoftwareUpdatePlanPost) GetDevices() []PlanDevicePost`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *ManagedSoftwareUpdatePlanPost) GetDevicesOk() (*[]PlanDevicePost, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *ManagedSoftwareUpdatePlanPost) SetDevices(v []PlanDevicePost)`

SetDevices sets Devices field to given value.


### GetConfig

`func (o *ManagedSoftwareUpdatePlanPost) GetConfig() PlanConfigurationPost`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ManagedSoftwareUpdatePlanPost) GetConfigOk() (*PlanConfigurationPost, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ManagedSoftwareUpdatePlanPost) SetConfig(v PlanConfigurationPost)`

SetConfig sets Config field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


