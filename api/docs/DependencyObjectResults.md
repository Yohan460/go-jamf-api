# DependencyObjectResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | Pointer to **int64** |  | [optional] 
**Results** | Pointer to [**[]DependencyObjectResultsResultsInner**](DependencyObjectResultsResultsInner.md) |  | [optional] 

## Methods

### NewDependencyObjectResults

`func NewDependencyObjectResults() *DependencyObjectResults`

NewDependencyObjectResults instantiates a new DependencyObjectResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDependencyObjectResultsWithDefaults

`func NewDependencyObjectResultsWithDefaults() *DependencyObjectResults`

NewDependencyObjectResultsWithDefaults instantiates a new DependencyObjectResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *DependencyObjectResults) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *DependencyObjectResults) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *DependencyObjectResults) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *DependencyObjectResults) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetResults

`func (o *DependencyObjectResults) GetResults() []DependencyObjectResultsResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *DependencyObjectResults) GetResultsOk() (*[]DependencyObjectResultsResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *DependencyObjectResults) SetResults(v []DependencyObjectResultsResultsInner)`

SetResults sets Results field to given value.

### HasResults

`func (o *DependencyObjectResults) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


