# InventoryPreloadInvalidCsvResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HttpsStatus** | Pointer to **int32** |  | [optional] 
**Errors** | Pointer to [**[]InventoryPreloadCsvError**](InventoryPreloadCsvError.md) |  | [optional] 

## Methods

### NewInventoryPreloadInvalidCsvResponse

`func NewInventoryPreloadInvalidCsvResponse() *InventoryPreloadInvalidCsvResponse`

NewInventoryPreloadInvalidCsvResponse instantiates a new InventoryPreloadInvalidCsvResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryPreloadInvalidCsvResponseWithDefaults

`func NewInventoryPreloadInvalidCsvResponseWithDefaults() *InventoryPreloadInvalidCsvResponse`

NewInventoryPreloadInvalidCsvResponseWithDefaults instantiates a new InventoryPreloadInvalidCsvResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttpsStatus

`func (o *InventoryPreloadInvalidCsvResponse) GetHttpsStatus() int32`

GetHttpsStatus returns the HttpsStatus field if non-nil, zero value otherwise.

### GetHttpsStatusOk

`func (o *InventoryPreloadInvalidCsvResponse) GetHttpsStatusOk() (*int32, bool)`

GetHttpsStatusOk returns a tuple with the HttpsStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsStatus

`func (o *InventoryPreloadInvalidCsvResponse) SetHttpsStatus(v int32)`

SetHttpsStatus sets HttpsStatus field to given value.

### HasHttpsStatus

`func (o *InventoryPreloadInvalidCsvResponse) HasHttpsStatus() bool`

HasHttpsStatus returns a boolean if a field has been set.

### GetErrors

`func (o *InventoryPreloadInvalidCsvResponse) GetErrors() []InventoryPreloadCsvError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *InventoryPreloadInvalidCsvResponse) GetErrorsOk() (*[]InventoryPreloadCsvError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *InventoryPreloadInvalidCsvResponse) SetErrors(v []InventoryPreloadCsvError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *InventoryPreloadInvalidCsvResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


