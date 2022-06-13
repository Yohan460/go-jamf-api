# SearchPatchPolicyLogParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageNumber** | Pointer to **int32** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**OrderBy** | Pointer to [**[]OrderBy**](OrderBy.md) |  | [optional] 
**Filter** | Pointer to [**[]Filter**](Filter.md) |  | [optional] 
**IsLoadToEnd** | Pointer to **bool** |  | [optional] 
**PatchPolicyId** | Pointer to **int32** |  | [optional] 
**IsLatest** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewSearchPatchPolicyLogParams

`func NewSearchPatchPolicyLogParams() *SearchPatchPolicyLogParams`

NewSearchPatchPolicyLogParams instantiates a new SearchPatchPolicyLogParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchPatchPolicyLogParamsWithDefaults

`func NewSearchPatchPolicyLogParamsWithDefaults() *SearchPatchPolicyLogParams`

NewSearchPatchPolicyLogParamsWithDefaults instantiates a new SearchPatchPolicyLogParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageNumber

`func (o *SearchPatchPolicyLogParams) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *SearchPatchPolicyLogParams) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *SearchPatchPolicyLogParams) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *SearchPatchPolicyLogParams) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *SearchPatchPolicyLogParams) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *SearchPatchPolicyLogParams) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *SearchPatchPolicyLogParams) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *SearchPatchPolicyLogParams) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetOrderBy

`func (o *SearchPatchPolicyLogParams) GetOrderBy() []OrderBy`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *SearchPatchPolicyLogParams) GetOrderByOk() (*[]OrderBy, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *SearchPatchPolicyLogParams) SetOrderBy(v []OrderBy)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *SearchPatchPolicyLogParams) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetFilter

`func (o *SearchPatchPolicyLogParams) GetFilter() []Filter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *SearchPatchPolicyLogParams) GetFilterOk() (*[]Filter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *SearchPatchPolicyLogParams) SetFilter(v []Filter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *SearchPatchPolicyLogParams) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetIsLoadToEnd

`func (o *SearchPatchPolicyLogParams) GetIsLoadToEnd() bool`

GetIsLoadToEnd returns the IsLoadToEnd field if non-nil, zero value otherwise.

### GetIsLoadToEndOk

`func (o *SearchPatchPolicyLogParams) GetIsLoadToEndOk() (*bool, bool)`

GetIsLoadToEndOk returns a tuple with the IsLoadToEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLoadToEnd

`func (o *SearchPatchPolicyLogParams) SetIsLoadToEnd(v bool)`

SetIsLoadToEnd sets IsLoadToEnd field to given value.

### HasIsLoadToEnd

`func (o *SearchPatchPolicyLogParams) HasIsLoadToEnd() bool`

HasIsLoadToEnd returns a boolean if a field has been set.

### GetPatchPolicyId

`func (o *SearchPatchPolicyLogParams) GetPatchPolicyId() int32`

GetPatchPolicyId returns the PatchPolicyId field if non-nil, zero value otherwise.

### GetPatchPolicyIdOk

`func (o *SearchPatchPolicyLogParams) GetPatchPolicyIdOk() (*int32, bool)`

GetPatchPolicyIdOk returns a tuple with the PatchPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchPolicyId

`func (o *SearchPatchPolicyLogParams) SetPatchPolicyId(v int32)`

SetPatchPolicyId sets PatchPolicyId field to given value.

### HasPatchPolicyId

`func (o *SearchPatchPolicyLogParams) HasPatchPolicyId() bool`

HasPatchPolicyId returns a boolean if a field has been set.

### GetIsLatest

`func (o *SearchPatchPolicyLogParams) GetIsLatest() bool`

GetIsLatest returns the IsLatest field if non-nil, zero value otherwise.

### GetIsLatestOk

`func (o *SearchPatchPolicyLogParams) GetIsLatestOk() (*bool, bool)`

GetIsLatestOk returns a tuple with the IsLatest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLatest

`func (o *SearchPatchPolicyLogParams) SetIsLatest(v bool)`

SetIsLatest sets IsLatest field to given value.

### HasIsLatest

`func (o *SearchPatchPolicyLogParams) HasIsLatest() bool`

HasIsLatest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


