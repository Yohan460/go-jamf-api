# OnboardingConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Enabled** | **bool** |  | 
**OnboardingItems** | [**[]OnboardingItem**](OnboardingItem.md) |  | 

## Methods

### NewOnboardingConfiguration

`func NewOnboardingConfiguration(enabled bool, onboardingItems []OnboardingItem, ) *OnboardingConfiguration`

NewOnboardingConfiguration instantiates a new OnboardingConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnboardingConfigurationWithDefaults

`func NewOnboardingConfigurationWithDefaults() *OnboardingConfiguration`

NewOnboardingConfigurationWithDefaults instantiates a new OnboardingConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OnboardingConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OnboardingConfiguration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OnboardingConfiguration) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OnboardingConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEnabled

`func (o *OnboardingConfiguration) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OnboardingConfiguration) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OnboardingConfiguration) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetOnboardingItems

`func (o *OnboardingConfiguration) GetOnboardingItems() []OnboardingItem`

GetOnboardingItems returns the OnboardingItems field if non-nil, zero value otherwise.

### GetOnboardingItemsOk

`func (o *OnboardingConfiguration) GetOnboardingItemsOk() (*[]OnboardingItem, bool)`

GetOnboardingItemsOk returns a tuple with the OnboardingItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardingItems

`func (o *OnboardingConfiguration) SetOnboardingItems(v []OnboardingItem)`

SetOnboardingItems sets OnboardingItems field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


