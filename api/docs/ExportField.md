# ExportField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldName** | Pointer to **string** | English name of the field to be exported. | [optional] 
**FieldLabelOverride** | Pointer to **NullableString** | Name which should be used for the label in the response - can be in any language. When null the fieldName itself will be used as the label. | [optional] 

## Methods

### NewExportField

`func NewExportField() *ExportField`

NewExportField instantiates a new ExportField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportFieldWithDefaults

`func NewExportFieldWithDefaults() *ExportField`

NewExportFieldWithDefaults instantiates a new ExportField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldName

`func (o *ExportField) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *ExportField) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *ExportField) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.

### HasFieldName

`func (o *ExportField) HasFieldName() bool`

HasFieldName returns a boolean if a field has been set.

### GetFieldLabelOverride

`func (o *ExportField) GetFieldLabelOverride() string`

GetFieldLabelOverride returns the FieldLabelOverride field if non-nil, zero value otherwise.

### GetFieldLabelOverrideOk

`func (o *ExportField) GetFieldLabelOverrideOk() (*string, bool)`

GetFieldLabelOverrideOk returns a tuple with the FieldLabelOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldLabelOverride

`func (o *ExportField) SetFieldLabelOverride(v string)`

SetFieldLabelOverride sets FieldLabelOverride field to given value.

### HasFieldLabelOverride

`func (o *ExportField) HasFieldLabelOverride() bool`

HasFieldLabelOverride returns a boolean if a field has been set.

### SetFieldLabelOverrideNil

`func (o *ExportField) SetFieldLabelOverrideNil(b bool)`

 SetFieldLabelOverrideNil sets the value for FieldLabelOverride to be an explicit nil

### UnsetFieldLabelOverride
`func (o *ExportField) UnsetFieldLabelOverride()`

UnsetFieldLabelOverride ensures that no value is present for FieldLabelOverride, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


