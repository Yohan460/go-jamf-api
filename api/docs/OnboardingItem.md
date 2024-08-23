# OnboardingItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**EntityId** | **string** | The id of the Jamf Pro object that should be added to the onboarding workflow for end users. Use this in conjunction with the selfServiceEntityType. For example, if the policy with id 132 should be added to onboarding, then entityId should be 132 and selfServiceEntityType should be  OS_X_POLICY.  | 
**EntityName** | Pointer to **string** |  | [optional] [readonly] 
**ScopeDescription** | Pointer to **string** |  | [optional] [readonly] 
**SiteDescription** | Pointer to **string** |  | [optional] [readonly] 
**SelfServiceEntityType** | **string** |  | 
**Priority** | **int64** |  | 

## Methods

### NewOnboardingItem

`func NewOnboardingItem(entityId string, selfServiceEntityType string, priority int64, ) *OnboardingItem`

NewOnboardingItem instantiates a new OnboardingItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnboardingItemWithDefaults

`func NewOnboardingItemWithDefaults() *OnboardingItem`

NewOnboardingItemWithDefaults instantiates a new OnboardingItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OnboardingItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OnboardingItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OnboardingItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OnboardingItem) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *OnboardingItem) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *OnboardingItem) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetEntityId

`func (o *OnboardingItem) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *OnboardingItem) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *OnboardingItem) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetEntityName

`func (o *OnboardingItem) GetEntityName() string`

GetEntityName returns the EntityName field if non-nil, zero value otherwise.

### GetEntityNameOk

`func (o *OnboardingItem) GetEntityNameOk() (*string, bool)`

GetEntityNameOk returns a tuple with the EntityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityName

`func (o *OnboardingItem) SetEntityName(v string)`

SetEntityName sets EntityName field to given value.

### HasEntityName

`func (o *OnboardingItem) HasEntityName() bool`

HasEntityName returns a boolean if a field has been set.

### GetScopeDescription

`func (o *OnboardingItem) GetScopeDescription() string`

GetScopeDescription returns the ScopeDescription field if non-nil, zero value otherwise.

### GetScopeDescriptionOk

`func (o *OnboardingItem) GetScopeDescriptionOk() (*string, bool)`

GetScopeDescriptionOk returns a tuple with the ScopeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeDescription

`func (o *OnboardingItem) SetScopeDescription(v string)`

SetScopeDescription sets ScopeDescription field to given value.

### HasScopeDescription

`func (o *OnboardingItem) HasScopeDescription() bool`

HasScopeDescription returns a boolean if a field has been set.

### GetSiteDescription

`func (o *OnboardingItem) GetSiteDescription() string`

GetSiteDescription returns the SiteDescription field if non-nil, zero value otherwise.

### GetSiteDescriptionOk

`func (o *OnboardingItem) GetSiteDescriptionOk() (*string, bool)`

GetSiteDescriptionOk returns a tuple with the SiteDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteDescription

`func (o *OnboardingItem) SetSiteDescription(v string)`

SetSiteDescription sets SiteDescription field to given value.

### HasSiteDescription

`func (o *OnboardingItem) HasSiteDescription() bool`

HasSiteDescription returns a boolean if a field has been set.

### GetSelfServiceEntityType

`func (o *OnboardingItem) GetSelfServiceEntityType() string`

GetSelfServiceEntityType returns the SelfServiceEntityType field if non-nil, zero value otherwise.

### GetSelfServiceEntityTypeOk

`func (o *OnboardingItem) GetSelfServiceEntityTypeOk() (*string, bool)`

GetSelfServiceEntityTypeOk returns a tuple with the SelfServiceEntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfServiceEntityType

`func (o *OnboardingItem) SetSelfServiceEntityType(v string)`

SetSelfServiceEntityType sets SelfServiceEntityType field to given value.


### GetPriority

`func (o *OnboardingItem) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *OnboardingItem) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *OnboardingItem) SetPriority(v int64)`

SetPriority sets Priority field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


