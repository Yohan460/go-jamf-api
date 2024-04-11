# Recipient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**RealName** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 

## Methods

### NewRecipient

`func NewRecipient() *Recipient`

NewRecipient instantiates a new Recipient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecipientWithDefaults

`func NewRecipientWithDefaults() *Recipient`

NewRecipientWithDefaults instantiates a new Recipient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Recipient) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Recipient) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Recipient) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Recipient) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRealName

`func (o *Recipient) GetRealName() string`

GetRealName returns the RealName field if non-nil, zero value otherwise.

### GetRealNameOk

`func (o *Recipient) GetRealNameOk() (*string, bool)`

GetRealNameOk returns a tuple with the RealName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealName

`func (o *Recipient) SetRealName(v string)`

SetRealName sets RealName field to given value.

### HasRealName

`func (o *Recipient) HasRealName() bool`

HasRealName returns a boolean if a field has been set.

### GetEmail

`func (o *Recipient) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Recipient) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Recipient) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Recipient) HasEmail() bool`

HasEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


