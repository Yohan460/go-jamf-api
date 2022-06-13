# AdvancedUserContentSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Name** | **string** |  | 
**Criteria** | Pointer to [**[]SmartSearchCriterion**](SmartSearchCriterion.md) |  | [optional] 
**DisplayFields** | Pointer to **[]string** |  | [optional] 
**SiteId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAdvancedUserContentSearch

`func NewAdvancedUserContentSearch(name string, ) *AdvancedUserContentSearch`

NewAdvancedUserContentSearch instantiates a new AdvancedUserContentSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvancedUserContentSearchWithDefaults

`func NewAdvancedUserContentSearchWithDefaults() *AdvancedUserContentSearch`

NewAdvancedUserContentSearchWithDefaults instantiates a new AdvancedUserContentSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AdvancedUserContentSearch) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdvancedUserContentSearch) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdvancedUserContentSearch) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AdvancedUserContentSearch) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AdvancedUserContentSearch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdvancedUserContentSearch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdvancedUserContentSearch) SetName(v string)`

SetName sets Name field to given value.


### GetCriteria

`func (o *AdvancedUserContentSearch) GetCriteria() []SmartSearchCriterion`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *AdvancedUserContentSearch) GetCriteriaOk() (*[]SmartSearchCriterion, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *AdvancedUserContentSearch) SetCriteria(v []SmartSearchCriterion)`

SetCriteria sets Criteria field to given value.

### HasCriteria

`func (o *AdvancedUserContentSearch) HasCriteria() bool`

HasCriteria returns a boolean if a field has been set.

### GetDisplayFields

`func (o *AdvancedUserContentSearch) GetDisplayFields() []string`

GetDisplayFields returns the DisplayFields field if non-nil, zero value otherwise.

### GetDisplayFieldsOk

`func (o *AdvancedUserContentSearch) GetDisplayFieldsOk() (*[]string, bool)`

GetDisplayFieldsOk returns a tuple with the DisplayFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayFields

`func (o *AdvancedUserContentSearch) SetDisplayFields(v []string)`

SetDisplayFields sets DisplayFields field to given value.

### HasDisplayFields

`func (o *AdvancedUserContentSearch) HasDisplayFields() bool`

HasDisplayFields returns a boolean if a field has been set.

### GetSiteId

`func (o *AdvancedUserContentSearch) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *AdvancedUserContentSearch) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *AdvancedUserContentSearch) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *AdvancedUserContentSearch) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### SetSiteIdNil

`func (o *AdvancedUserContentSearch) SetSiteIdNil(b bool)`

 SetSiteIdNil sets the value for SiteId to be an explicit nil

### UnsetSiteId
`func (o *AdvancedUserContentSearch) UnsetSiteId()`

UnsetSiteId ensures that no value is present for SiteId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


