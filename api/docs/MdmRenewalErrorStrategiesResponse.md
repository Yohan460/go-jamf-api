# MdmRenewalErrorStrategiesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | [**MdmRenewalError**](MdmRenewalError.md) |  | 
**Strategies** | [**[]MdmRenewalStrategy**](MdmRenewalStrategy.md) | List of renewal strategies associated with this error | 

## Methods

### NewMdmRenewalErrorStrategiesResponse

`func NewMdmRenewalErrorStrategiesResponse(error_ MdmRenewalError, strategies []MdmRenewalStrategy, ) *MdmRenewalErrorStrategiesResponse`

NewMdmRenewalErrorStrategiesResponse instantiates a new MdmRenewalErrorStrategiesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMdmRenewalErrorStrategiesResponseWithDefaults

`func NewMdmRenewalErrorStrategiesResponseWithDefaults() *MdmRenewalErrorStrategiesResponse`

NewMdmRenewalErrorStrategiesResponseWithDefaults instantiates a new MdmRenewalErrorStrategiesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *MdmRenewalErrorStrategiesResponse) GetError() MdmRenewalError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MdmRenewalErrorStrategiesResponse) GetErrorOk() (*MdmRenewalError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *MdmRenewalErrorStrategiesResponse) SetError(v MdmRenewalError)`

SetError sets Error field to given value.


### GetStrategies

`func (o *MdmRenewalErrorStrategiesResponse) GetStrategies() []MdmRenewalStrategy`

GetStrategies returns the Strategies field if non-nil, zero value otherwise.

### GetStrategiesOk

`func (o *MdmRenewalErrorStrategiesResponse) GetStrategiesOk() (*[]MdmRenewalStrategy, bool)`

GetStrategiesOk returns a tuple with the Strategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategies

`func (o *MdmRenewalErrorStrategiesResponse) SetStrategies(v []MdmRenewalStrategy)`

SetStrategies sets Strategies field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


