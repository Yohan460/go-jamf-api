# DashboardSetupTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** |  | [optional] [default to false]
**Error** | Pointer to [**DashboardApiError**](DashboardApiError.md) |  | [optional] 

## Methods

### NewDashboardSetupTask

`func NewDashboardSetupTask() *DashboardSetupTask`

NewDashboardSetupTask instantiates a new DashboardSetupTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardSetupTaskWithDefaults

`func NewDashboardSetupTaskWithDefaults() *DashboardSetupTask`

NewDashboardSetupTaskWithDefaults instantiates a new DashboardSetupTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *DashboardSetupTask) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *DashboardSetupTask) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *DashboardSetupTask) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *DashboardSetupTask) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetError

`func (o *DashboardSetupTask) GetError() DashboardApiError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DashboardSetupTask) GetErrorOk() (*DashboardApiError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DashboardSetupTask) SetError(v DashboardApiError)`

SetError sets Error field to given value.

### HasError

`func (o *DashboardSetupTask) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


