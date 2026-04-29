# ComputerInventorySearchResultsV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | Pointer to **int64** |  | [optional] 
**Results** | Pointer to [**[]ComputerInventoryV3**](ComputerInventoryV3.md) |  | [optional] 

## Methods

### NewComputerInventorySearchResultsV3

`func NewComputerInventorySearchResultsV3() *ComputerInventorySearchResultsV3`

NewComputerInventorySearchResultsV3 instantiates a new ComputerInventorySearchResultsV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputerInventorySearchResultsV3WithDefaults

`func NewComputerInventorySearchResultsV3WithDefaults() *ComputerInventorySearchResultsV3`

NewComputerInventorySearchResultsV3WithDefaults instantiates a new ComputerInventorySearchResultsV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *ComputerInventorySearchResultsV3) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ComputerInventorySearchResultsV3) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ComputerInventorySearchResultsV3) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ComputerInventorySearchResultsV3) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetResults

`func (o *ComputerInventorySearchResultsV3) GetResults() []ComputerInventoryV3`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ComputerInventorySearchResultsV3) GetResultsOk() (*[]ComputerInventoryV3, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ComputerInventorySearchResultsV3) SetResults(v []ComputerInventoryV3)`

SetResults sets Results field to given value.

### HasResults

`func (o *ComputerInventorySearchResultsV3) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


