# HistorySearchResultsV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | Pointer to **int64** |  | [optional] 
**Results** | Pointer to [**[]ObjectHistoryV1**](ObjectHistoryV1.md) |  | [optional] 

## Methods

### NewHistorySearchResultsV1

`func NewHistorySearchResultsV1() *HistorySearchResultsV1`

NewHistorySearchResultsV1 instantiates a new HistorySearchResultsV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistorySearchResultsV1WithDefaults

`func NewHistorySearchResultsV1WithDefaults() *HistorySearchResultsV1`

NewHistorySearchResultsV1WithDefaults instantiates a new HistorySearchResultsV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *HistorySearchResultsV1) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *HistorySearchResultsV1) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *HistorySearchResultsV1) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *HistorySearchResultsV1) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetResults

`func (o *HistorySearchResultsV1) GetResults() []ObjectHistoryV1`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *HistorySearchResultsV1) GetResultsOk() (*[]ObjectHistoryV1, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *HistorySearchResultsV1) SetResults(v []ObjectHistoryV1)`

SetResults sets Results field to given value.

### HasResults

`func (o *HistorySearchResultsV1) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


