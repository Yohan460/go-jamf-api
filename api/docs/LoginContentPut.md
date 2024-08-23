# LoginContentPut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeCustomDisclaimer** | **bool** |  | 
**DisclaimerHeading** | Pointer to **string** |  | [optional] 
**DisclaimerMainText** | Pointer to **string** |  | [optional] 
**ActionText** | Pointer to **string** |  | [optional] 

## Methods

### NewLoginContentPut

`func NewLoginContentPut(includeCustomDisclaimer bool, ) *LoginContentPut`

NewLoginContentPut instantiates a new LoginContentPut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginContentPutWithDefaults

`func NewLoginContentPutWithDefaults() *LoginContentPut`

NewLoginContentPutWithDefaults instantiates a new LoginContentPut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludeCustomDisclaimer

`func (o *LoginContentPut) GetIncludeCustomDisclaimer() bool`

GetIncludeCustomDisclaimer returns the IncludeCustomDisclaimer field if non-nil, zero value otherwise.

### GetIncludeCustomDisclaimerOk

`func (o *LoginContentPut) GetIncludeCustomDisclaimerOk() (*bool, bool)`

GetIncludeCustomDisclaimerOk returns a tuple with the IncludeCustomDisclaimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeCustomDisclaimer

`func (o *LoginContentPut) SetIncludeCustomDisclaimer(v bool)`

SetIncludeCustomDisclaimer sets IncludeCustomDisclaimer field to given value.


### GetDisclaimerHeading

`func (o *LoginContentPut) GetDisclaimerHeading() string`

GetDisclaimerHeading returns the DisclaimerHeading field if non-nil, zero value otherwise.

### GetDisclaimerHeadingOk

`func (o *LoginContentPut) GetDisclaimerHeadingOk() (*string, bool)`

GetDisclaimerHeadingOk returns a tuple with the DisclaimerHeading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclaimerHeading

`func (o *LoginContentPut) SetDisclaimerHeading(v string)`

SetDisclaimerHeading sets DisclaimerHeading field to given value.

### HasDisclaimerHeading

`func (o *LoginContentPut) HasDisclaimerHeading() bool`

HasDisclaimerHeading returns a boolean if a field has been set.

### GetDisclaimerMainText

`func (o *LoginContentPut) GetDisclaimerMainText() string`

GetDisclaimerMainText returns the DisclaimerMainText field if non-nil, zero value otherwise.

### GetDisclaimerMainTextOk

`func (o *LoginContentPut) GetDisclaimerMainTextOk() (*string, bool)`

GetDisclaimerMainTextOk returns a tuple with the DisclaimerMainText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclaimerMainText

`func (o *LoginContentPut) SetDisclaimerMainText(v string)`

SetDisclaimerMainText sets DisclaimerMainText field to given value.

### HasDisclaimerMainText

`func (o *LoginContentPut) HasDisclaimerMainText() bool`

HasDisclaimerMainText returns a boolean if a field has been set.

### GetActionText

`func (o *LoginContentPut) GetActionText() string`

GetActionText returns the ActionText field if non-nil, zero value otherwise.

### GetActionTextOk

`func (o *LoginContentPut) GetActionTextOk() (*string, bool)`

GetActionTextOk returns a tuple with the ActionText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionText

`func (o *LoginContentPut) SetActionText(v string)`

SetActionText sets ActionText field to given value.

### HasActionText

`func (o *LoginContentPut) HasActionText() bool`

HasActionText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


