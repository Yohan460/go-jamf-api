# Signature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | Pointer to **string** |  | [optional] 
**AlgorithmOid** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 

## Methods

### NewSignature

`func NewSignature() *Signature`

NewSignature instantiates a new Signature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignatureWithDefaults

`func NewSignatureWithDefaults() *Signature`

NewSignatureWithDefaults instantiates a new Signature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithm

`func (o *Signature) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *Signature) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *Signature) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *Signature) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetAlgorithmOid

`func (o *Signature) GetAlgorithmOid() string`

GetAlgorithmOid returns the AlgorithmOid field if non-nil, zero value otherwise.

### GetAlgorithmOidOk

`func (o *Signature) GetAlgorithmOidOk() (*string, bool)`

GetAlgorithmOidOk returns a tuple with the AlgorithmOid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithmOid

`func (o *Signature) SetAlgorithmOid(v string)`

SetAlgorithmOid sets AlgorithmOid field to given value.

### HasAlgorithmOid

`func (o *Signature) HasAlgorithmOid() bool`

HasAlgorithmOid returns a boolean if a field has been set.

### GetValue

`func (o *Signature) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Signature) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Signature) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *Signature) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


