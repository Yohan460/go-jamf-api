# AccessGroupsPreviewSearchResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | Pointer to **int32** |  | [optional] 
**Results** | Pointer to [**[]EnrollmentAccessGroupPreview**](EnrollmentAccessGroupPreview.md) |  | [optional] 

## Methods

### NewAccessGroupsPreviewSearchResults

`func NewAccessGroupsPreviewSearchResults() *AccessGroupsPreviewSearchResults`

NewAccessGroupsPreviewSearchResults instantiates a new AccessGroupsPreviewSearchResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessGroupsPreviewSearchResultsWithDefaults

`func NewAccessGroupsPreviewSearchResultsWithDefaults() *AccessGroupsPreviewSearchResults`

NewAccessGroupsPreviewSearchResultsWithDefaults instantiates a new AccessGroupsPreviewSearchResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *AccessGroupsPreviewSearchResults) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *AccessGroupsPreviewSearchResults) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *AccessGroupsPreviewSearchResults) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *AccessGroupsPreviewSearchResults) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetResults

`func (o *AccessGroupsPreviewSearchResults) GetResults() []EnrollmentAccessGroupPreview`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *AccessGroupsPreviewSearchResults) GetResultsOk() (*[]EnrollmentAccessGroupPreview, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *AccessGroupsPreviewSearchResults) SetResults(v []EnrollmentAccessGroupPreview)`

SetResults sets Results field to given value.

### HasResults

`func (o *AccessGroupsPreviewSearchResults) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


