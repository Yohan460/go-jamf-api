# LoginContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RampInstance** | Pointer to **bool** |  | [optional] 
**IncludeCustomDisclaimer** | **bool** |  | 
**DisclaimerHeading** | Pointer to **string** |  | [optional] 
**DisclaimerMainText** | Pointer to **string** |  | [optional] 
**ActionText** | Pointer to **string** |  | [optional] 

## Methods

### NewLoginContent

`func NewLoginContent(includeCustomDisclaimer bool, ) *LoginContent`

NewLoginContent instantiates a new LoginContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginContentWithDefaults

`func NewLoginContentWithDefaults() *LoginContent`

NewLoginContentWithDefaults instantiates a new LoginContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRampInstance

`func (o *LoginContent) GetRampInstance() bool`

GetRampInstance returns the RampInstance field if non-nil, zero value otherwise.

### GetRampInstanceOk

`func (o *LoginContent) GetRampInstanceOk() (*bool, bool)`

GetRampInstanceOk returns a tuple with the RampInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRampInstance

`func (o *LoginContent) SetRampInstance(v bool)`

SetRampInstance sets RampInstance field to given value.

### HasRampInstance

`func (o *LoginContent) HasRampInstance() bool`

HasRampInstance returns a boolean if a field has been set.

### GetIncludeCustomDisclaimer

`func (o *LoginContent) GetIncludeCustomDisclaimer() bool`

GetIncludeCustomDisclaimer returns the IncludeCustomDisclaimer field if non-nil, zero value otherwise.

### GetIncludeCustomDisclaimerOk

`func (o *LoginContent) GetIncludeCustomDisclaimerOk() (*bool, bool)`

GetIncludeCustomDisclaimerOk returns a tuple with the IncludeCustomDisclaimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeCustomDisclaimer

`func (o *LoginContent) SetIncludeCustomDisclaimer(v bool)`

SetIncludeCustomDisclaimer sets IncludeCustomDisclaimer field to given value.


### GetDisclaimerHeading

`func (o *LoginContent) GetDisclaimerHeading() string`

GetDisclaimerHeading returns the DisclaimerHeading field if non-nil, zero value otherwise.

### GetDisclaimerHeadingOk

`func (o *LoginContent) GetDisclaimerHeadingOk() (*string, bool)`

GetDisclaimerHeadingOk returns a tuple with the DisclaimerHeading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclaimerHeading

`func (o *LoginContent) SetDisclaimerHeading(v string)`

SetDisclaimerHeading sets DisclaimerHeading field to given value.

### HasDisclaimerHeading

`func (o *LoginContent) HasDisclaimerHeading() bool`

HasDisclaimerHeading returns a boolean if a field has been set.

### GetDisclaimerMainText

`func (o *LoginContent) GetDisclaimerMainText() string`

GetDisclaimerMainText returns the DisclaimerMainText field if non-nil, zero value otherwise.

### GetDisclaimerMainTextOk

`func (o *LoginContent) GetDisclaimerMainTextOk() (*string, bool)`

GetDisclaimerMainTextOk returns a tuple with the DisclaimerMainText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclaimerMainText

`func (o *LoginContent) SetDisclaimerMainText(v string)`

SetDisclaimerMainText sets DisclaimerMainText field to given value.

### HasDisclaimerMainText

`func (o *LoginContent) HasDisclaimerMainText() bool`

HasDisclaimerMainText returns a boolean if a field has been set.

### GetActionText

`func (o *LoginContent) GetActionText() string`

GetActionText returns the ActionText field if non-nil, zero value otherwise.

### GetActionTextOk

`func (o *LoginContent) GetActionTextOk() (*string, bool)`

GetActionTextOk returns a tuple with the ActionText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionText

`func (o *LoginContent) SetActionText(v string)`

SetActionText sets ActionText field to given value.

### HasActionText

`func (o *LoginContent) HasActionText() bool`

HasActionText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


