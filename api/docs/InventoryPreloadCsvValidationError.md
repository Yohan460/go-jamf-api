# InventoryPreloadCsvValidationError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HttpStatus** | Pointer to **int32** | HTTP status of the response | [optional] 
**Errors** | Pointer to [**[]InventoryPreloadCsvValidationErrorCause**](InventoryPreloadCsvValidationErrorCause.md) |  | [optional] 

## Methods

### NewInventoryPreloadCsvValidationError

`func NewInventoryPreloadCsvValidationError() *InventoryPreloadCsvValidationError`

NewInventoryPreloadCsvValidationError instantiates a new InventoryPreloadCsvValidationError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryPreloadCsvValidationErrorWithDefaults

`func NewInventoryPreloadCsvValidationErrorWithDefaults() *InventoryPreloadCsvValidationError`

NewInventoryPreloadCsvValidationErrorWithDefaults instantiates a new InventoryPreloadCsvValidationError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttpStatus

`func (o *InventoryPreloadCsvValidationError) GetHttpStatus() int32`

GetHttpStatus returns the HttpStatus field if non-nil, zero value otherwise.

### GetHttpStatusOk

`func (o *InventoryPreloadCsvValidationError) GetHttpStatusOk() (*int32, bool)`

GetHttpStatusOk returns a tuple with the HttpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpStatus

`func (o *InventoryPreloadCsvValidationError) SetHttpStatus(v int32)`

SetHttpStatus sets HttpStatus field to given value.

### HasHttpStatus

`func (o *InventoryPreloadCsvValidationError) HasHttpStatus() bool`

HasHttpStatus returns a boolean if a field has been set.

### GetErrors

`func (o *InventoryPreloadCsvValidationError) GetErrors() []InventoryPreloadCsvValidationErrorCause`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *InventoryPreloadCsvValidationError) GetErrorsOk() (*[]InventoryPreloadCsvValidationErrorCause, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *InventoryPreloadCsvValidationError) SetErrors(v []InventoryPreloadCsvValidationErrorCause)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *InventoryPreloadCsvValidationError) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


