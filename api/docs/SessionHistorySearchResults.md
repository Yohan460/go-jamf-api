# SessionHistorySearchResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]SessionHistoryItem**](SessionHistoryItem.md) |  | 
**TotalCount** | **int64** |  | 

## Methods

### NewSessionHistorySearchResults

`func NewSessionHistorySearchResults(results []SessionHistoryItem, totalCount int64, ) *SessionHistorySearchResults`

NewSessionHistorySearchResults instantiates a new SessionHistorySearchResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionHistorySearchResultsWithDefaults

`func NewSessionHistorySearchResultsWithDefaults() *SessionHistorySearchResults`

NewSessionHistorySearchResultsWithDefaults instantiates a new SessionHistorySearchResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *SessionHistorySearchResults) GetResults() []SessionHistoryItem`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SessionHistorySearchResults) GetResultsOk() (*[]SessionHistoryItem, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SessionHistorySearchResults) SetResults(v []SessionHistoryItem)`

SetResults sets Results field to given value.


### GetTotalCount

`func (o *SessionHistorySearchResults) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *SessionHistorySearchResults) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *SessionHistorySearchResults) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


