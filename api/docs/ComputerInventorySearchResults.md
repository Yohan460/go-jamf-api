# ComputerInventorySearchResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | Pointer to **int32** |  | [optional] 
**Results** | Pointer to [**[]ComputerInventoryResponse**](ComputerInventoryResponse.md) |  | [optional] 

## Methods

### NewComputerInventorySearchResults

`func NewComputerInventorySearchResults() *ComputerInventorySearchResults`

NewComputerInventorySearchResults instantiates a new ComputerInventorySearchResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputerInventorySearchResultsWithDefaults

`func NewComputerInventorySearchResultsWithDefaults() *ComputerInventorySearchResults`

NewComputerInventorySearchResultsWithDefaults instantiates a new ComputerInventorySearchResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *ComputerInventorySearchResults) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ComputerInventorySearchResults) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ComputerInventorySearchResults) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ComputerInventorySearchResults) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetResults

`func (o *ComputerInventorySearchResults) GetResults() []ComputerInventoryResponse`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ComputerInventorySearchResults) GetResultsOk() (*[]ComputerInventoryResponse, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ComputerInventorySearchResults) SetResults(v []ComputerInventoryResponse)`

SetResults sets Results field to given value.

### HasResults

`func (o *ComputerInventorySearchResults) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


