# GroupTestSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | Pointer to **int64** |  | [optional] 
**Results** | Pointer to [**[]GroupTestSearch**](GroupTestSearch.md) |  | [optional] 

## Methods

### NewGroupTestSearchResponse

`func NewGroupTestSearchResponse() *GroupTestSearchResponse`

NewGroupTestSearchResponse instantiates a new GroupTestSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupTestSearchResponseWithDefaults

`func NewGroupTestSearchResponseWithDefaults() *GroupTestSearchResponse`

NewGroupTestSearchResponseWithDefaults instantiates a new GroupTestSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *GroupTestSearchResponse) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *GroupTestSearchResponse) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *GroupTestSearchResponse) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *GroupTestSearchResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetResults

`func (o *GroupTestSearchResponse) GetResults() []GroupTestSearch`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GroupTestSearchResponse) GetResultsOk() (*[]GroupTestSearch, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GroupTestSearchResponse) SetResults(v []GroupTestSearch)`

SetResults sets Results field to given value.

### HasResults

`func (o *GroupTestSearchResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


