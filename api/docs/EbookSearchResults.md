# EbookSearchResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | Pointer to **int32** |  | [optional] 
**Results** | Pointer to [**[]Ebook**](Ebook.md) |  | [optional] 

## Methods

### NewEbookSearchResults

`func NewEbookSearchResults() *EbookSearchResults`

NewEbookSearchResults instantiates a new EbookSearchResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEbookSearchResultsWithDefaults

`func NewEbookSearchResultsWithDefaults() *EbookSearchResults`

NewEbookSearchResultsWithDefaults instantiates a new EbookSearchResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *EbookSearchResults) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *EbookSearchResults) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *EbookSearchResults) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *EbookSearchResults) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetResults

`func (o *EbookSearchResults) GetResults() []Ebook`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *EbookSearchResults) GetResultsOk() (*[]Ebook, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *EbookSearchResults) SetResults(v []Ebook)`

SetResults sets Results field to given value.

### HasResults

`func (o *EbookSearchResults) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


