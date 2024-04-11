# ApiError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HttpStatus** | Pointer to **int64** | HTTP status of the response | [optional] 
**Errors** | Pointer to [**[]ApiErrorCause**](ApiErrorCause.md) |  | [optional] 

## Methods

### NewApiError

`func NewApiError() *ApiError`

NewApiError instantiates a new ApiError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiErrorWithDefaults

`func NewApiErrorWithDefaults() *ApiError`

NewApiErrorWithDefaults instantiates a new ApiError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttpStatus

`func (o *ApiError) GetHttpStatus() int64`

GetHttpStatus returns the HttpStatus field if non-nil, zero value otherwise.

### GetHttpStatusOk

`func (o *ApiError) GetHttpStatusOk() (*int64, bool)`

GetHttpStatusOk returns a tuple with the HttpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpStatus

`func (o *ApiError) SetHttpStatus(v int64)`

SetHttpStatus sets HttpStatus field to given value.

### HasHttpStatus

`func (o *ApiError) HasHttpStatus() bool`

HasHttpStatus returns a boolean if a field has been set.

### GetErrors

`func (o *ApiError) GetErrors() []ApiErrorCause`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ApiError) GetErrorsOk() (*[]ApiErrorCause, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ApiError) SetErrors(v []ApiErrorCause)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ApiError) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


