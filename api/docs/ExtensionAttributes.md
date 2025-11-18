# ExtensionAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique Id for Mobile Device Extension Attribute. | [optional] [readonly] 
**Name** | **string** | Display name for the extension attribute. | 
**Description** | Pointer to **string** | Description for the extension attribute. | [optional] 
**DataType** | **string** | Type of data being collected. | [default to "STRING"]

## Methods

### NewExtensionAttributes

`func NewExtensionAttributes(name string, dataType string, ) *ExtensionAttributes`

NewExtensionAttributes instantiates a new ExtensionAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionAttributesWithDefaults

`func NewExtensionAttributesWithDefaults() *ExtensionAttributes`

NewExtensionAttributesWithDefaults instantiates a new ExtensionAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExtensionAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExtensionAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExtensionAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ExtensionAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ExtensionAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExtensionAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExtensionAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ExtensionAttributes) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExtensionAttributes) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExtensionAttributes) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExtensionAttributes) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDataType

`func (o *ExtensionAttributes) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *ExtensionAttributes) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *ExtensionAttributes) SetDataType(v string)`

SetDataType sets DataType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


