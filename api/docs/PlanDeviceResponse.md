# PlanDeviceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | Pointer to [**PlanDevice**](PlanDevice.md) |  | [optional] 
**PlanId** | Pointer to **string** |  | [optional] [readonly] 
**Href** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewPlanDeviceResponse

`func NewPlanDeviceResponse() *PlanDeviceResponse`

NewPlanDeviceResponse instantiates a new PlanDeviceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanDeviceResponseWithDefaults

`func NewPlanDeviceResponseWithDefaults() *PlanDeviceResponse`

NewPlanDeviceResponseWithDefaults instantiates a new PlanDeviceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *PlanDeviceResponse) GetDevice() PlanDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *PlanDeviceResponse) GetDeviceOk() (*PlanDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *PlanDeviceResponse) SetDevice(v PlanDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *PlanDeviceResponse) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetPlanId

`func (o *PlanDeviceResponse) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *PlanDeviceResponse) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *PlanDeviceResponse) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *PlanDeviceResponse) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetHref

`func (o *PlanDeviceResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *PlanDeviceResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *PlanDeviceResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *PlanDeviceResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


