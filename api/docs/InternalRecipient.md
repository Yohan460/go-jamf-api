# InternalRecipient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** |  | 
**Frequency** | Pointer to **string** |  | [optional] [default to "DAILY"]

## Methods

### NewInternalRecipient

`func NewInternalRecipient(accountId string, ) *InternalRecipient`

NewInternalRecipient instantiates a new InternalRecipient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalRecipientWithDefaults

`func NewInternalRecipientWithDefaults() *InternalRecipient`

NewInternalRecipientWithDefaults instantiates a new InternalRecipient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *InternalRecipient) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *InternalRecipient) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *InternalRecipient) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetFrequency

`func (o *InternalRecipient) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *InternalRecipient) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *InternalRecipient) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *InternalRecipient) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


