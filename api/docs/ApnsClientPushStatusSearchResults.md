# ApnsClientPushStatusSearchResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int64** | Total number of records matching the query | 
**Results** | [**[]ApnsClientPushStatus**](ApnsClientPushStatus.md) | Array of APNS client push status records | 

## Methods

### NewApnsClientPushStatusSearchResults

`func NewApnsClientPushStatusSearchResults(totalCount int64, results []ApnsClientPushStatus, ) *ApnsClientPushStatusSearchResults`

NewApnsClientPushStatusSearchResults instantiates a new ApnsClientPushStatusSearchResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApnsClientPushStatusSearchResultsWithDefaults

`func NewApnsClientPushStatusSearchResultsWithDefaults() *ApnsClientPushStatusSearchResults`

NewApnsClientPushStatusSearchResultsWithDefaults instantiates a new ApnsClientPushStatusSearchResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *ApnsClientPushStatusSearchResults) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ApnsClientPushStatusSearchResults) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ApnsClientPushStatusSearchResults) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.


### GetResults

`func (o *ApnsClientPushStatusSearchResults) GetResults() []ApnsClientPushStatus`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ApnsClientPushStatusSearchResults) GetResultsOk() (*[]ApnsClientPushStatus, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ApnsClientPushStatusSearchResults) SetResults(v []ApnsClientPushStatus)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


