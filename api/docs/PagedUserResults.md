# PagedUserResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to **int64** | Current page number (zero-based) | [optional] 
**PageSize** | Pointer to **int64** | Number of results per page | [optional] 
**TotalCount** | Pointer to **int64** | Total number of users matching the search criteria | [optional] 
**TotalPages** | Pointer to **int64** | Total number of pages available | [optional] 
**HasNext** | Pointer to **bool** | True if there are more pages after the current page | [optional] 
**HasPrevious** | Pointer to **bool** | True if there are pages before the current page | [optional] 
**Results** | Pointer to [**[]User**](User.md) | List of users in the current page | [optional] 

## Methods

### NewPagedUserResults

`func NewPagedUserResults() *PagedUserResults`

NewPagedUserResults instantiates a new PagedUserResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagedUserResultsWithDefaults

`func NewPagedUserResultsWithDefaults() *PagedUserResults`

NewPagedUserResultsWithDefaults instantiates a new PagedUserResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *PagedUserResults) GetPage() int64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PagedUserResults) GetPageOk() (*int64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PagedUserResults) SetPage(v int64)`

SetPage sets Page field to given value.

### HasPage

`func (o *PagedUserResults) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *PagedUserResults) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *PagedUserResults) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *PagedUserResults) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *PagedUserResults) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetTotalCount

`func (o *PagedUserResults) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *PagedUserResults) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *PagedUserResults) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *PagedUserResults) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetTotalPages

`func (o *PagedUserResults) GetTotalPages() int64`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *PagedUserResults) GetTotalPagesOk() (*int64, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *PagedUserResults) SetTotalPages(v int64)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *PagedUserResults) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### GetHasNext

`func (o *PagedUserResults) GetHasNext() bool`

GetHasNext returns the HasNext field if non-nil, zero value otherwise.

### GetHasNextOk

`func (o *PagedUserResults) GetHasNextOk() (*bool, bool)`

GetHasNextOk returns a tuple with the HasNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNext

`func (o *PagedUserResults) SetHasNext(v bool)`

SetHasNext sets HasNext field to given value.

### HasHasNext

`func (o *PagedUserResults) HasHasNext() bool`

HasHasNext returns a boolean if a field has been set.

### GetHasPrevious

`func (o *PagedUserResults) GetHasPrevious() bool`

GetHasPrevious returns the HasPrevious field if non-nil, zero value otherwise.

### GetHasPreviousOk

`func (o *PagedUserResults) GetHasPreviousOk() (*bool, bool)`

GetHasPreviousOk returns a tuple with the HasPrevious field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrevious

`func (o *PagedUserResults) SetHasPrevious(v bool)`

SetHasPrevious sets HasPrevious field to given value.

### HasHasPrevious

`func (o *PagedUserResults) HasHasPrevious() bool`

HasHasPrevious returns a boolean if a field has been set.

### GetResults

`func (o *PagedUserResults) GetResults() []User`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PagedUserResults) GetResultsOk() (*[]User, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PagedUserResults) SetResults(v []User)`

SetResults sets Results field to given value.

### HasResults

`func (o *PagedUserResults) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


