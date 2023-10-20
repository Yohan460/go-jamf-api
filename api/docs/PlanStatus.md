# PlanStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** |  | [optional] [readonly] 
**ErrorReasons** | Pointer to **[]string** |  | [optional] [readonly] 

## Methods

### NewPlanStatus

`func NewPlanStatus() *PlanStatus`

NewPlanStatus instantiates a new PlanStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanStatusWithDefaults

`func NewPlanStatusWithDefaults() *PlanStatus`

NewPlanStatusWithDefaults instantiates a new PlanStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *PlanStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PlanStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PlanStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *PlanStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetErrorReasons

`func (o *PlanStatus) GetErrorReasons() []string`

GetErrorReasons returns the ErrorReasons field if non-nil, zero value otherwise.

### GetErrorReasonsOk

`func (o *PlanStatus) GetErrorReasonsOk() (*[]string, bool)`

GetErrorReasonsOk returns a tuple with the ErrorReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorReasons

`func (o *PlanStatus) SetErrorReasons(v []string)`

SetErrorReasons sets ErrorReasons field to given value.

### HasErrorReasons

`func (o *PlanStatus) HasErrorReasons() bool`

HasErrorReasons returns a boolean if a field has been set.

### SetErrorReasonsNil

`func (o *PlanStatus) SetErrorReasonsNil(b bool)`

 SetErrorReasonsNil sets the value for ErrorReasons to be an explicit nil

### UnsetErrorReasons
`func (o *PlanStatus) UnsetErrorReasons()`

UnsetErrorReasons ensures that no value is present for ErrorReasons, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


