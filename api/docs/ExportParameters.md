# ExportParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to **NullableInt64** |  | [optional] [default to 0]
**PageSize** | Pointer to **NullableInt64** |  | [optional] [default to 100]
**Sort** | Pointer to **[]string** | Sorting criteria in the format: [&lt;property&gt;[:asc/desc]. Default direction when not stated is ascending. | [optional] [default to ["id:desc"]]
**Filter** | Pointer to **NullableString** |  | [optional] 
**Fields** | Pointer to [**[]ExportField**](ExportField.md) | Used to change default order or ignore some of the fields. When null or empty array, all fields will be exported. | [optional] 

## Methods

### NewExportParameters

`func NewExportParameters() *ExportParameters`

NewExportParameters instantiates a new ExportParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportParametersWithDefaults

`func NewExportParametersWithDefaults() *ExportParameters`

NewExportParametersWithDefaults instantiates a new ExportParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *ExportParameters) GetPage() int64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ExportParameters) GetPageOk() (*int64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ExportParameters) SetPage(v int64)`

SetPage sets Page field to given value.

### HasPage

`func (o *ExportParameters) HasPage() bool`

HasPage returns a boolean if a field has been set.

### SetPageNil

`func (o *ExportParameters) SetPageNil(b bool)`

 SetPageNil sets the value for Page to be an explicit nil

### UnsetPage
`func (o *ExportParameters) UnsetPage()`

UnsetPage ensures that no value is present for Page, not even an explicit nil
### GetPageSize

`func (o *ExportParameters) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ExportParameters) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ExportParameters) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ExportParameters) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### SetPageSizeNil

`func (o *ExportParameters) SetPageSizeNil(b bool)`

 SetPageSizeNil sets the value for PageSize to be an explicit nil

### UnsetPageSize
`func (o *ExportParameters) UnsetPageSize()`

UnsetPageSize ensures that no value is present for PageSize, not even an explicit nil
### GetSort

`func (o *ExportParameters) GetSort() []string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *ExportParameters) GetSortOk() (*[]string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *ExportParameters) SetSort(v []string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *ExportParameters) HasSort() bool`

HasSort returns a boolean if a field has been set.

### SetSortNil

`func (o *ExportParameters) SetSortNil(b bool)`

 SetSortNil sets the value for Sort to be an explicit nil

### UnsetSort
`func (o *ExportParameters) UnsetSort()`

UnsetSort ensures that no value is present for Sort, not even an explicit nil
### GetFilter

`func (o *ExportParameters) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ExportParameters) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ExportParameters) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ExportParameters) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *ExportParameters) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *ExportParameters) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil
### GetFields

`func (o *ExportParameters) GetFields() []ExportField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ExportParameters) GetFieldsOk() (*[]ExportField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ExportParameters) SetFields(v []ExportField)`

SetFields sets Fields field to given value.

### HasFields

`func (o *ExportParameters) HasFields() bool`

HasFields returns a boolean if a field has been set.

### SetFieldsNil

`func (o *ExportParameters) SetFieldsNil(b bool)`

 SetFieldsNil sets the value for Fields to be an explicit nil

### UnsetFields
`func (o *ExportParameters) UnsetFields()`

UnsetFields ensures that no value is present for Fields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


