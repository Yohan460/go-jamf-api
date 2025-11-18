# SmartGroupSearchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | Pointer to **int64** |  | [optional] 
**Results** | Pointer to [**[]SmartComputerGroupSearch**](SmartComputerGroupSearch.md) |  | [optional] 

## Methods

### NewSmartGroupSearchResult

`func NewSmartGroupSearchResult() *SmartGroupSearchResult`

NewSmartGroupSearchResult instantiates a new SmartGroupSearchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmartGroupSearchResultWithDefaults

`func NewSmartGroupSearchResultWithDefaults() *SmartGroupSearchResult`

NewSmartGroupSearchResultWithDefaults instantiates a new SmartGroupSearchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *SmartGroupSearchResult) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *SmartGroupSearchResult) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *SmartGroupSearchResult) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *SmartGroupSearchResult) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetResults

`func (o *SmartGroupSearchResult) GetResults() []SmartComputerGroupSearch`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SmartGroupSearchResult) GetResultsOk() (*[]SmartComputerGroupSearch, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SmartGroupSearchResult) SetResults(v []SmartComputerGroupSearch)`

SetResults sets Results field to given value.

### HasResults

`func (o *SmartGroupSearchResult) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


