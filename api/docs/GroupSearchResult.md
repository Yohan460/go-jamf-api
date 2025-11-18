# GroupSearchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | Pointer to **int64** |  | [optional] 
**Results** | Pointer to [**[]GroupV1**](GroupV1.md) |  | [optional] 

## Methods

### NewGroupSearchResult

`func NewGroupSearchResult() *GroupSearchResult`

NewGroupSearchResult instantiates a new GroupSearchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupSearchResultWithDefaults

`func NewGroupSearchResultWithDefaults() *GroupSearchResult`

NewGroupSearchResultWithDefaults instantiates a new GroupSearchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *GroupSearchResult) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *GroupSearchResult) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *GroupSearchResult) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *GroupSearchResult) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetResults

`func (o *GroupSearchResult) GetResults() []GroupV1`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GroupSearchResult) GetResultsOk() (*[]GroupV1, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GroupSearchResult) SetResults(v []GroupV1)`

SetResults sets Results field to given value.

### HasResults

`func (o *GroupSearchResult) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


