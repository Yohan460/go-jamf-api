# MdmRenewalStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the renewal strategy | 
**MdmRenewalErrorId** | **string** | The MDM renewal error ID this strategy is associated with | 
**MdmRenewalStrategyType** | [**MdmRenewalStrategyType**](MdmRenewalStrategyType.md) |  | 
**StrategyTimeStamp** | Pointer to **time.Time** | Timestamp when this renewal strategy was created (ISO 8601 format) | [optional] 
**MdmRenewalCheckInUrl** | Pointer to **string** | URL for MDM renewal check-in | [optional] 
**MdmRenewalServerUrl** | Pointer to **string** | URL for MDM renewal server | [optional] 

## Methods

### NewMdmRenewalStrategy

`func NewMdmRenewalStrategy(id string, mdmRenewalErrorId string, mdmRenewalStrategyType MdmRenewalStrategyType, ) *MdmRenewalStrategy`

NewMdmRenewalStrategy instantiates a new MdmRenewalStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMdmRenewalStrategyWithDefaults

`func NewMdmRenewalStrategyWithDefaults() *MdmRenewalStrategy`

NewMdmRenewalStrategyWithDefaults instantiates a new MdmRenewalStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MdmRenewalStrategy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MdmRenewalStrategy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MdmRenewalStrategy) SetId(v string)`

SetId sets Id field to given value.


### GetMdmRenewalErrorId

`func (o *MdmRenewalStrategy) GetMdmRenewalErrorId() string`

GetMdmRenewalErrorId returns the MdmRenewalErrorId field if non-nil, zero value otherwise.

### GetMdmRenewalErrorIdOk

`func (o *MdmRenewalStrategy) GetMdmRenewalErrorIdOk() (*string, bool)`

GetMdmRenewalErrorIdOk returns a tuple with the MdmRenewalErrorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdmRenewalErrorId

`func (o *MdmRenewalStrategy) SetMdmRenewalErrorId(v string)`

SetMdmRenewalErrorId sets MdmRenewalErrorId field to given value.


### GetMdmRenewalStrategyType

`func (o *MdmRenewalStrategy) GetMdmRenewalStrategyType() MdmRenewalStrategyType`

GetMdmRenewalStrategyType returns the MdmRenewalStrategyType field if non-nil, zero value otherwise.

### GetMdmRenewalStrategyTypeOk

`func (o *MdmRenewalStrategy) GetMdmRenewalStrategyTypeOk() (*MdmRenewalStrategyType, bool)`

GetMdmRenewalStrategyTypeOk returns a tuple with the MdmRenewalStrategyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdmRenewalStrategyType

`func (o *MdmRenewalStrategy) SetMdmRenewalStrategyType(v MdmRenewalStrategyType)`

SetMdmRenewalStrategyType sets MdmRenewalStrategyType field to given value.


### GetStrategyTimeStamp

`func (o *MdmRenewalStrategy) GetStrategyTimeStamp() time.Time`

GetStrategyTimeStamp returns the StrategyTimeStamp field if non-nil, zero value otherwise.

### GetStrategyTimeStampOk

`func (o *MdmRenewalStrategy) GetStrategyTimeStampOk() (*time.Time, bool)`

GetStrategyTimeStampOk returns a tuple with the StrategyTimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategyTimeStamp

`func (o *MdmRenewalStrategy) SetStrategyTimeStamp(v time.Time)`

SetStrategyTimeStamp sets StrategyTimeStamp field to given value.

### HasStrategyTimeStamp

`func (o *MdmRenewalStrategy) HasStrategyTimeStamp() bool`

HasStrategyTimeStamp returns a boolean if a field has been set.

### GetMdmRenewalCheckInUrl

`func (o *MdmRenewalStrategy) GetMdmRenewalCheckInUrl() string`

GetMdmRenewalCheckInUrl returns the MdmRenewalCheckInUrl field if non-nil, zero value otherwise.

### GetMdmRenewalCheckInUrlOk

`func (o *MdmRenewalStrategy) GetMdmRenewalCheckInUrlOk() (*string, bool)`

GetMdmRenewalCheckInUrlOk returns a tuple with the MdmRenewalCheckInUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdmRenewalCheckInUrl

`func (o *MdmRenewalStrategy) SetMdmRenewalCheckInUrl(v string)`

SetMdmRenewalCheckInUrl sets MdmRenewalCheckInUrl field to given value.

### HasMdmRenewalCheckInUrl

`func (o *MdmRenewalStrategy) HasMdmRenewalCheckInUrl() bool`

HasMdmRenewalCheckInUrl returns a boolean if a field has been set.

### GetMdmRenewalServerUrl

`func (o *MdmRenewalStrategy) GetMdmRenewalServerUrl() string`

GetMdmRenewalServerUrl returns the MdmRenewalServerUrl field if non-nil, zero value otherwise.

### GetMdmRenewalServerUrlOk

`func (o *MdmRenewalStrategy) GetMdmRenewalServerUrlOk() (*string, bool)`

GetMdmRenewalServerUrlOk returns a tuple with the MdmRenewalServerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdmRenewalServerUrl

`func (o *MdmRenewalStrategy) SetMdmRenewalServerUrl(v string)`

SetMdmRenewalServerUrl sets MdmRenewalServerUrl field to given value.

### HasMdmRenewalServerUrl

`func (o *MdmRenewalStrategy) HasMdmRenewalServerUrl() bool`

HasMdmRenewalServerUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


