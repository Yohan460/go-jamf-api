# ComputerInventorySearchResultsV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | Pointer to **int64** |  | [optional] 
**Results** | Pointer to [**[]ComputerInventoryV2**](ComputerInventoryV2.md) |  | [optional] 

## Methods

### NewComputerInventorySearchResultsV2

`func NewComputerInventorySearchResultsV2() *ComputerInventorySearchResultsV2`

NewComputerInventorySearchResultsV2 instantiates a new ComputerInventorySearchResultsV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputerInventorySearchResultsV2WithDefaults

`func NewComputerInventorySearchResultsV2WithDefaults() *ComputerInventorySearchResultsV2`

NewComputerInventorySearchResultsV2WithDefaults instantiates a new ComputerInventorySearchResultsV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *ComputerInventorySearchResultsV2) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ComputerInventorySearchResultsV2) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ComputerInventorySearchResultsV2) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ComputerInventorySearchResultsV2) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetResults

`func (o *ComputerInventorySearchResultsV2) GetResults() []ComputerInventoryV2`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ComputerInventorySearchResultsV2) GetResultsOk() (*[]ComputerInventoryV2, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ComputerInventorySearchResultsV2) SetResults(v []ComputerInventoryV2)`

SetResults sets Results field to given value.

### HasResults

`func (o *ComputerInventorySearchResultsV2) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


