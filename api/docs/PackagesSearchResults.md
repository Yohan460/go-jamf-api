# PackagesSearchResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | Pointer to **int64** |  | [optional] 
**Results** | Pointer to [**[]Package**](Package.md) |  | [optional] 

## Methods

### NewPackagesSearchResults

`func NewPackagesSearchResults() *PackagesSearchResults`

NewPackagesSearchResults instantiates a new PackagesSearchResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackagesSearchResultsWithDefaults

`func NewPackagesSearchResultsWithDefaults() *PackagesSearchResults`

NewPackagesSearchResultsWithDefaults instantiates a new PackagesSearchResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *PackagesSearchResults) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *PackagesSearchResults) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *PackagesSearchResults) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *PackagesSearchResults) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetResults

`func (o *PackagesSearchResults) GetResults() []Package`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PackagesSearchResults) GetResultsOk() (*[]Package, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PackagesSearchResults) SetResults(v []Package)`

SetResults sets Results field to given value.

### HasResults

`func (o *PackagesSearchResults) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


