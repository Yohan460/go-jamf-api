# ApiErrorCause

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Error-specific code that can be used to identify localization string, etc. | [optional] 
**Field** | **string** | Name of the field that caused the error. | 
**Description** | Pointer to **string** | A general description of error for troubleshooting/debugging. Generally this text should not be displayed to a user; instead refer to errorCode and it&#39;s localized text | [optional] 
**Id** | Pointer to **NullableString** | id of object with error. Optional. | [optional] 

## Methods

### NewApiErrorCause

`func NewApiErrorCause(field string, ) *ApiErrorCause`

NewApiErrorCause instantiates a new ApiErrorCause object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiErrorCauseWithDefaults

`func NewApiErrorCauseWithDefaults() *ApiErrorCause`

NewApiErrorCauseWithDefaults instantiates a new ApiErrorCause object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ApiErrorCause) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ApiErrorCause) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ApiErrorCause) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ApiErrorCause) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetField

`func (o *ApiErrorCause) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *ApiErrorCause) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *ApiErrorCause) SetField(v string)`

SetField sets Field field to given value.


### GetDescription

`func (o *ApiErrorCause) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiErrorCause) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiErrorCause) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiErrorCause) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *ApiErrorCause) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiErrorCause) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiErrorCause) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiErrorCause) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ApiErrorCause) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ApiErrorCause) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


