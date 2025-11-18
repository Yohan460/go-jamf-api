# SmartComputerGroupV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Criteria** | Pointer to [**[]SmartSearchCriterion**](SmartSearchCriterion.md) |  | [optional] 
**SiteId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSmartComputerGroupV2

`func NewSmartComputerGroupV2(name string, ) *SmartComputerGroupV2`

NewSmartComputerGroupV2 instantiates a new SmartComputerGroupV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmartComputerGroupV2WithDefaults

`func NewSmartComputerGroupV2WithDefaults() *SmartComputerGroupV2`

NewSmartComputerGroupV2WithDefaults instantiates a new SmartComputerGroupV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SmartComputerGroupV2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SmartComputerGroupV2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SmartComputerGroupV2) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SmartComputerGroupV2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SmartComputerGroupV2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SmartComputerGroupV2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SmartComputerGroupV2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCriteria

`func (o *SmartComputerGroupV2) GetCriteria() []SmartSearchCriterion`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *SmartComputerGroupV2) GetCriteriaOk() (*[]SmartSearchCriterion, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *SmartComputerGroupV2) SetCriteria(v []SmartSearchCriterion)`

SetCriteria sets Criteria field to given value.

### HasCriteria

`func (o *SmartComputerGroupV2) HasCriteria() bool`

HasCriteria returns a boolean if a field has been set.

### GetSiteId

`func (o *SmartComputerGroupV2) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *SmartComputerGroupV2) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *SmartComputerGroupV2) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *SmartComputerGroupV2) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### SetSiteIdNil

`func (o *SmartComputerGroupV2) SetSiteIdNil(b bool)`

 SetSiteIdNil sets the value for SiteId to be an explicit nil

### UnsetSiteId
`func (o *SmartComputerGroupV2) UnsetSiteId()`

UnsetSiteId ensures that no value is present for SiteId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


