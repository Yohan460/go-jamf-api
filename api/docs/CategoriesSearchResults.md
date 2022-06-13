# CategoriesSearchResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | Pointer to **int32** |  | [optional] 
**Results** | Pointer to [**[]Category**](Category.md) |  | [optional] 

## Methods

### NewCategoriesSearchResults

`func NewCategoriesSearchResults() *CategoriesSearchResults`

NewCategoriesSearchResults instantiates a new CategoriesSearchResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoriesSearchResultsWithDefaults

`func NewCategoriesSearchResultsWithDefaults() *CategoriesSearchResults`

NewCategoriesSearchResultsWithDefaults instantiates a new CategoriesSearchResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *CategoriesSearchResults) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *CategoriesSearchResults) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *CategoriesSearchResults) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *CategoriesSearchResults) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetResults

`func (o *CategoriesSearchResults) GetResults() []Category`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *CategoriesSearchResults) GetResultsOk() (*[]Category, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *CategoriesSearchResults) SetResults(v []Category)`

SetResults sets Results field to given value.

### HasResults

`func (o *CategoriesSearchResults) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


