# Assignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MobileDeviceId** | Pointer to **string** |  | [optional] 
**Selected** | Pointer to **bool** | If true the device should be added to the group, if false should be removed from the group. | [optional] 

## Methods

### NewAssignment

`func NewAssignment() *Assignment`

NewAssignment instantiates a new Assignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignmentWithDefaults

`func NewAssignmentWithDefaults() *Assignment`

NewAssignmentWithDefaults instantiates a new Assignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMobileDeviceId

`func (o *Assignment) GetMobileDeviceId() string`

GetMobileDeviceId returns the MobileDeviceId field if non-nil, zero value otherwise.

### GetMobileDeviceIdOk

`func (o *Assignment) GetMobileDeviceIdOk() (*string, bool)`

GetMobileDeviceIdOk returns a tuple with the MobileDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileDeviceId

`func (o *Assignment) SetMobileDeviceId(v string)`

SetMobileDeviceId sets MobileDeviceId field to given value.

### HasMobileDeviceId

`func (o *Assignment) HasMobileDeviceId() bool`

HasMobileDeviceId returns a boolean if a field has been set.

### GetSelected

`func (o *Assignment) GetSelected() bool`

GetSelected returns the Selected field if non-nil, zero value otherwise.

### GetSelectedOk

`func (o *Assignment) GetSelectedOk() (*bool, bool)`

GetSelectedOk returns a tuple with the Selected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelected

`func (o *Assignment) SetSelected(v bool)`

SetSelected sets Selected field to given value.

### HasSelected

`func (o *Assignment) HasSelected() bool`

HasSelected returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


