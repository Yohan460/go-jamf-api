# InventoryPreloadCsvValidationErrorCause

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Error-specific code that can be used to identify localization string, etc. | [optional] 
**Field** | **string** | Name of the field that caused the error. | 
**Description** | Pointer to **string** | A general description of error for troubleshooting/debugging. Generally this text should not be displayed to a user; instead refer to errorCode and it&#39;s localized text | [optional] 
**Id** | Pointer to **string** | id of object with error. Optional. | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**Line** | Pointer to **int32** |  | [optional] 
**FieldSize** | Pointer to **int32** |  | [optional] 
**DeviceType** | Pointer to **string** |  | [optional] 

## Methods

### NewInventoryPreloadCsvValidationErrorCause

`func NewInventoryPreloadCsvValidationErrorCause(field string, ) *InventoryPreloadCsvValidationErrorCause`

NewInventoryPreloadCsvValidationErrorCause instantiates a new InventoryPreloadCsvValidationErrorCause object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryPreloadCsvValidationErrorCauseWithDefaults

`func NewInventoryPreloadCsvValidationErrorCauseWithDefaults() *InventoryPreloadCsvValidationErrorCause`

NewInventoryPreloadCsvValidationErrorCauseWithDefaults instantiates a new InventoryPreloadCsvValidationErrorCause object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *InventoryPreloadCsvValidationErrorCause) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *InventoryPreloadCsvValidationErrorCause) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *InventoryPreloadCsvValidationErrorCause) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *InventoryPreloadCsvValidationErrorCause) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetField

`func (o *InventoryPreloadCsvValidationErrorCause) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *InventoryPreloadCsvValidationErrorCause) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *InventoryPreloadCsvValidationErrorCause) SetField(v string)`

SetField sets Field field to given value.


### GetDescription

`func (o *InventoryPreloadCsvValidationErrorCause) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InventoryPreloadCsvValidationErrorCause) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InventoryPreloadCsvValidationErrorCause) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InventoryPreloadCsvValidationErrorCause) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *InventoryPreloadCsvValidationErrorCause) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InventoryPreloadCsvValidationErrorCause) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InventoryPreloadCsvValidationErrorCause) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InventoryPreloadCsvValidationErrorCause) HasId() bool`

HasId returns a boolean if a field has been set.

### GetValue

`func (o *InventoryPreloadCsvValidationErrorCause) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *InventoryPreloadCsvValidationErrorCause) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *InventoryPreloadCsvValidationErrorCause) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *InventoryPreloadCsvValidationErrorCause) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetSerialNumber

`func (o *InventoryPreloadCsvValidationErrorCause) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *InventoryPreloadCsvValidationErrorCause) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *InventoryPreloadCsvValidationErrorCause) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *InventoryPreloadCsvValidationErrorCause) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetLine

`func (o *InventoryPreloadCsvValidationErrorCause) GetLine() int32`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *InventoryPreloadCsvValidationErrorCause) GetLineOk() (*int32, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *InventoryPreloadCsvValidationErrorCause) SetLine(v int32)`

SetLine sets Line field to given value.

### HasLine

`func (o *InventoryPreloadCsvValidationErrorCause) HasLine() bool`

HasLine returns a boolean if a field has been set.

### GetFieldSize

`func (o *InventoryPreloadCsvValidationErrorCause) GetFieldSize() int32`

GetFieldSize returns the FieldSize field if non-nil, zero value otherwise.

### GetFieldSizeOk

`func (o *InventoryPreloadCsvValidationErrorCause) GetFieldSizeOk() (*int32, bool)`

GetFieldSizeOk returns a tuple with the FieldSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldSize

`func (o *InventoryPreloadCsvValidationErrorCause) SetFieldSize(v int32)`

SetFieldSize sets FieldSize field to given value.

### HasFieldSize

`func (o *InventoryPreloadCsvValidationErrorCause) HasFieldSize() bool`

HasFieldSize returns a boolean if a field has been set.

### GetDeviceType

`func (o *InventoryPreloadCsvValidationErrorCause) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *InventoryPreloadCsvValidationErrorCause) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *InventoryPreloadCsvValidationErrorCause) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *InventoryPreloadCsvValidationErrorCause) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


