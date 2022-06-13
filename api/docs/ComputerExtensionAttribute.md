# ComputerExtensionAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefinitionId** | Pointer to **string** | An identifier of extension attribute definition. | [optional] 
**Name** | Pointer to **string** | A human-readable name by which attribute can be referred to. | [optional] [readonly] 
**Description** | Pointer to **NullableString** | An additional explanation of exact attribute meaning, possible values, etc. | [optional] [readonly] 
**Enabled** | Pointer to **bool** |  | [optional] [readonly] 
**MultiValue** | Pointer to **bool** |  | [optional] [readonly] 
**Values** | Pointer to **[]string** | A value of extension attribute, in some rare cases there may be multiple values present, hence the array.  | [optional] 
**DataType** | Pointer to **NullableString** | A data type of extension attribute. | [optional] [readonly] 
**Options** | Pointer to **[]string** | A closed list of possible values (applies to &#x60;popup&#x60; input type).  | [optional] [readonly] 
**InputType** | Pointer to **NullableString** | The input method. &#x60;text&#x60; is most common and means simply free text, &#x60;popup&#x60; i a closed list of values from which one or many can be selected and &#x60;script&#x60; value is calculated and can never be set directly.  | [optional] [readonly] 

## Methods

### NewComputerExtensionAttribute

`func NewComputerExtensionAttribute() *ComputerExtensionAttribute`

NewComputerExtensionAttribute instantiates a new ComputerExtensionAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputerExtensionAttributeWithDefaults

`func NewComputerExtensionAttributeWithDefaults() *ComputerExtensionAttribute`

NewComputerExtensionAttributeWithDefaults instantiates a new ComputerExtensionAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefinitionId

`func (o *ComputerExtensionAttribute) GetDefinitionId() string`

GetDefinitionId returns the DefinitionId field if non-nil, zero value otherwise.

### GetDefinitionIdOk

`func (o *ComputerExtensionAttribute) GetDefinitionIdOk() (*string, bool)`

GetDefinitionIdOk returns a tuple with the DefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitionId

`func (o *ComputerExtensionAttribute) SetDefinitionId(v string)`

SetDefinitionId sets DefinitionId field to given value.

### HasDefinitionId

`func (o *ComputerExtensionAttribute) HasDefinitionId() bool`

HasDefinitionId returns a boolean if a field has been set.

### GetName

`func (o *ComputerExtensionAttribute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComputerExtensionAttribute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComputerExtensionAttribute) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ComputerExtensionAttribute) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ComputerExtensionAttribute) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ComputerExtensionAttribute) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ComputerExtensionAttribute) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ComputerExtensionAttribute) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ComputerExtensionAttribute) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ComputerExtensionAttribute) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnabled

`func (o *ComputerExtensionAttribute) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ComputerExtensionAttribute) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ComputerExtensionAttribute) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ComputerExtensionAttribute) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMultiValue

`func (o *ComputerExtensionAttribute) GetMultiValue() bool`

GetMultiValue returns the MultiValue field if non-nil, zero value otherwise.

### GetMultiValueOk

`func (o *ComputerExtensionAttribute) GetMultiValueOk() (*bool, bool)`

GetMultiValueOk returns a tuple with the MultiValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiValue

`func (o *ComputerExtensionAttribute) SetMultiValue(v bool)`

SetMultiValue sets MultiValue field to given value.

### HasMultiValue

`func (o *ComputerExtensionAttribute) HasMultiValue() bool`

HasMultiValue returns a boolean if a field has been set.

### GetValues

`func (o *ComputerExtensionAttribute) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ComputerExtensionAttribute) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ComputerExtensionAttribute) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *ComputerExtensionAttribute) HasValues() bool`

HasValues returns a boolean if a field has been set.

### SetValuesNil

`func (o *ComputerExtensionAttribute) SetValuesNil(b bool)`

 SetValuesNil sets the value for Values to be an explicit nil

### UnsetValues
`func (o *ComputerExtensionAttribute) UnsetValues()`

UnsetValues ensures that no value is present for Values, not even an explicit nil
### GetDataType

`func (o *ComputerExtensionAttribute) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *ComputerExtensionAttribute) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *ComputerExtensionAttribute) SetDataType(v string)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *ComputerExtensionAttribute) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### SetDataTypeNil

`func (o *ComputerExtensionAttribute) SetDataTypeNil(b bool)`

 SetDataTypeNil sets the value for DataType to be an explicit nil

### UnsetDataType
`func (o *ComputerExtensionAttribute) UnsetDataType()`

UnsetDataType ensures that no value is present for DataType, not even an explicit nil
### GetOptions

`func (o *ComputerExtensionAttribute) GetOptions() []string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ComputerExtensionAttribute) GetOptionsOk() (*[]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ComputerExtensionAttribute) SetOptions(v []string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ComputerExtensionAttribute) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *ComputerExtensionAttribute) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *ComputerExtensionAttribute) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetInputType

`func (o *ComputerExtensionAttribute) GetInputType() string`

GetInputType returns the InputType field if non-nil, zero value otherwise.

### GetInputTypeOk

`func (o *ComputerExtensionAttribute) GetInputTypeOk() (*string, bool)`

GetInputTypeOk returns a tuple with the InputType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputType

`func (o *ComputerExtensionAttribute) SetInputType(v string)`

SetInputType sets InputType field to given value.

### HasInputType

`func (o *ComputerExtensionAttribute) HasInputType() bool`

HasInputType returns a boolean if a field has been set.

### SetInputTypeNil

`func (o *ComputerExtensionAttribute) SetInputTypeNil(b bool)`

 SetInputTypeNil sets the value for InputType to be an explicit nil

### UnsetInputType
`func (o *ComputerExtensionAttribute) UnsetInputType()`

UnsetInputType ensures that no value is present for InputType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


