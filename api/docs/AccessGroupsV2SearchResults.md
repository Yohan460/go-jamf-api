# AccessGroupsV2SearchResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | Pointer to **int64** |  | [optional] 
**Results** | Pointer to [**[]EnrollmentAccessGroupV2**](EnrollmentAccessGroupV2.md) |  | [optional] 

## Methods

### NewAccessGroupsV2SearchResults

`func NewAccessGroupsV2SearchResults() *AccessGroupsV2SearchResults`

NewAccessGroupsV2SearchResults instantiates a new AccessGroupsV2SearchResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessGroupsV2SearchResultsWithDefaults

`func NewAccessGroupsV2SearchResultsWithDefaults() *AccessGroupsV2SearchResults`

NewAccessGroupsV2SearchResultsWithDefaults instantiates a new AccessGroupsV2SearchResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *AccessGroupsV2SearchResults) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *AccessGroupsV2SearchResults) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *AccessGroupsV2SearchResults) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *AccessGroupsV2SearchResults) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetResults

`func (o *AccessGroupsV2SearchResults) GetResults() []EnrollmentAccessGroupV2`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *AccessGroupsV2SearchResults) GetResultsOk() (*[]EnrollmentAccessGroupV2, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *AccessGroupsV2SearchResults) SetResults(v []EnrollmentAccessGroupV2)`

SetResults sets Results field to given value.

### HasResults

`func (o *AccessGroupsV2SearchResults) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


