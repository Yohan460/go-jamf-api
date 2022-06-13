# SearchActivePatchHistoryParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageNumber** | Pointer to **int32** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**IsLoadToEnd** | Pointer to **bool** |  | [optional] 
**SoftwareTitleID** | Pointer to **int32** |  | [optional] 
**SoftwareTitleConfigurationID** | Pointer to **int32** |  | [optional] 
**OrderBy** | Pointer to [**[]OrderBy**](OrderBy.md) |  | [optional] 
**Filter** | Pointer to [**[]Filter**](Filter.md) |  | [optional] 

## Methods

### NewSearchActivePatchHistoryParams

`func NewSearchActivePatchHistoryParams() *SearchActivePatchHistoryParams`

NewSearchActivePatchHistoryParams instantiates a new SearchActivePatchHistoryParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchActivePatchHistoryParamsWithDefaults

`func NewSearchActivePatchHistoryParamsWithDefaults() *SearchActivePatchHistoryParams`

NewSearchActivePatchHistoryParamsWithDefaults instantiates a new SearchActivePatchHistoryParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageNumber

`func (o *SearchActivePatchHistoryParams) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *SearchActivePatchHistoryParams) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *SearchActivePatchHistoryParams) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *SearchActivePatchHistoryParams) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *SearchActivePatchHistoryParams) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *SearchActivePatchHistoryParams) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *SearchActivePatchHistoryParams) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *SearchActivePatchHistoryParams) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetIsLoadToEnd

`func (o *SearchActivePatchHistoryParams) GetIsLoadToEnd() bool`

GetIsLoadToEnd returns the IsLoadToEnd field if non-nil, zero value otherwise.

### GetIsLoadToEndOk

`func (o *SearchActivePatchHistoryParams) GetIsLoadToEndOk() (*bool, bool)`

GetIsLoadToEndOk returns a tuple with the IsLoadToEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLoadToEnd

`func (o *SearchActivePatchHistoryParams) SetIsLoadToEnd(v bool)`

SetIsLoadToEnd sets IsLoadToEnd field to given value.

### HasIsLoadToEnd

`func (o *SearchActivePatchHistoryParams) HasIsLoadToEnd() bool`

HasIsLoadToEnd returns a boolean if a field has been set.

### GetSoftwareTitleID

`func (o *SearchActivePatchHistoryParams) GetSoftwareTitleID() int32`

GetSoftwareTitleID returns the SoftwareTitleID field if non-nil, zero value otherwise.

### GetSoftwareTitleIDOk

`func (o *SearchActivePatchHistoryParams) GetSoftwareTitleIDOk() (*int32, bool)`

GetSoftwareTitleIDOk returns a tuple with the SoftwareTitleID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareTitleID

`func (o *SearchActivePatchHistoryParams) SetSoftwareTitleID(v int32)`

SetSoftwareTitleID sets SoftwareTitleID field to given value.

### HasSoftwareTitleID

`func (o *SearchActivePatchHistoryParams) HasSoftwareTitleID() bool`

HasSoftwareTitleID returns a boolean if a field has been set.

### GetSoftwareTitleConfigurationID

`func (o *SearchActivePatchHistoryParams) GetSoftwareTitleConfigurationID() int32`

GetSoftwareTitleConfigurationID returns the SoftwareTitleConfigurationID field if non-nil, zero value otherwise.

### GetSoftwareTitleConfigurationIDOk

`func (o *SearchActivePatchHistoryParams) GetSoftwareTitleConfigurationIDOk() (*int32, bool)`

GetSoftwareTitleConfigurationIDOk returns a tuple with the SoftwareTitleConfigurationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareTitleConfigurationID

`func (o *SearchActivePatchHistoryParams) SetSoftwareTitleConfigurationID(v int32)`

SetSoftwareTitleConfigurationID sets SoftwareTitleConfigurationID field to given value.

### HasSoftwareTitleConfigurationID

`func (o *SearchActivePatchHistoryParams) HasSoftwareTitleConfigurationID() bool`

HasSoftwareTitleConfigurationID returns a boolean if a field has been set.

### GetOrderBy

`func (o *SearchActivePatchHistoryParams) GetOrderBy() []OrderBy`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *SearchActivePatchHistoryParams) GetOrderByOk() (*[]OrderBy, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *SearchActivePatchHistoryParams) SetOrderBy(v []OrderBy)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *SearchActivePatchHistoryParams) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetFilter

`func (o *SearchActivePatchHistoryParams) GetFilter() []Filter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *SearchActivePatchHistoryParams) GetFilterOk() (*[]Filter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *SearchActivePatchHistoryParams) SetFilter(v []Filter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *SearchActivePatchHistoryParams) HasFilter() bool`

HasFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


