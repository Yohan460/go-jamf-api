# CacheSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] [default to "0"]
**Name** | Pointer to **string** |  | [optional] [default to "cache configuration"]
**CacheType** | **string** |  | 
**TimeToLiveSeconds** | **int32** |  | 
**TimeToIdleSeconds** | Pointer to **int32** |  | [optional] 
**DirectoryTimeToLiveSeconds** | Pointer to **int32** |  | [optional] 
**EhcacheMaxBytesLocalHeap** | Pointer to **string** |  | [optional] [default to "null"]
**CacheUniqueId** | **string** | The default is for Jamf Pro to generate a UUID, so we can only give an example instead. | 
**Elasticache** | Pointer to **bool** |  | [optional] [default to false]
**MemcachedEndpoints** | [**[]MemcachedEndpoints**](MemcachedEndpoints.md) |  | 

## Methods

### NewCacheSettings

`func NewCacheSettings(cacheType string, timeToLiveSeconds int32, cacheUniqueId string, memcachedEndpoints []MemcachedEndpoints, ) *CacheSettings`

NewCacheSettings instantiates a new CacheSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCacheSettingsWithDefaults

`func NewCacheSettingsWithDefaults() *CacheSettings`

NewCacheSettingsWithDefaults instantiates a new CacheSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CacheSettings) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CacheSettings) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CacheSettings) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CacheSettings) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CacheSettings) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CacheSettings) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CacheSettings) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CacheSettings) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCacheType

`func (o *CacheSettings) GetCacheType() string`

GetCacheType returns the CacheType field if non-nil, zero value otherwise.

### GetCacheTypeOk

`func (o *CacheSettings) GetCacheTypeOk() (*string, bool)`

GetCacheTypeOk returns a tuple with the CacheType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheType

`func (o *CacheSettings) SetCacheType(v string)`

SetCacheType sets CacheType field to given value.


### GetTimeToLiveSeconds

`func (o *CacheSettings) GetTimeToLiveSeconds() int32`

GetTimeToLiveSeconds returns the TimeToLiveSeconds field if non-nil, zero value otherwise.

### GetTimeToLiveSecondsOk

`func (o *CacheSettings) GetTimeToLiveSecondsOk() (*int32, bool)`

GetTimeToLiveSecondsOk returns a tuple with the TimeToLiveSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToLiveSeconds

`func (o *CacheSettings) SetTimeToLiveSeconds(v int32)`

SetTimeToLiveSeconds sets TimeToLiveSeconds field to given value.


### GetTimeToIdleSeconds

`func (o *CacheSettings) GetTimeToIdleSeconds() int32`

GetTimeToIdleSeconds returns the TimeToIdleSeconds field if non-nil, zero value otherwise.

### GetTimeToIdleSecondsOk

`func (o *CacheSettings) GetTimeToIdleSecondsOk() (*int32, bool)`

GetTimeToIdleSecondsOk returns a tuple with the TimeToIdleSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToIdleSeconds

`func (o *CacheSettings) SetTimeToIdleSeconds(v int32)`

SetTimeToIdleSeconds sets TimeToIdleSeconds field to given value.

### HasTimeToIdleSeconds

`func (o *CacheSettings) HasTimeToIdleSeconds() bool`

HasTimeToIdleSeconds returns a boolean if a field has been set.

### GetDirectoryTimeToLiveSeconds

`func (o *CacheSettings) GetDirectoryTimeToLiveSeconds() int32`

GetDirectoryTimeToLiveSeconds returns the DirectoryTimeToLiveSeconds field if non-nil, zero value otherwise.

### GetDirectoryTimeToLiveSecondsOk

`func (o *CacheSettings) GetDirectoryTimeToLiveSecondsOk() (*int32, bool)`

GetDirectoryTimeToLiveSecondsOk returns a tuple with the DirectoryTimeToLiveSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryTimeToLiveSeconds

`func (o *CacheSettings) SetDirectoryTimeToLiveSeconds(v int32)`

SetDirectoryTimeToLiveSeconds sets DirectoryTimeToLiveSeconds field to given value.

### HasDirectoryTimeToLiveSeconds

`func (o *CacheSettings) HasDirectoryTimeToLiveSeconds() bool`

HasDirectoryTimeToLiveSeconds returns a boolean if a field has been set.

### GetEhcacheMaxBytesLocalHeap

`func (o *CacheSettings) GetEhcacheMaxBytesLocalHeap() string`

GetEhcacheMaxBytesLocalHeap returns the EhcacheMaxBytesLocalHeap field if non-nil, zero value otherwise.

### GetEhcacheMaxBytesLocalHeapOk

`func (o *CacheSettings) GetEhcacheMaxBytesLocalHeapOk() (*string, bool)`

GetEhcacheMaxBytesLocalHeapOk returns a tuple with the EhcacheMaxBytesLocalHeap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEhcacheMaxBytesLocalHeap

`func (o *CacheSettings) SetEhcacheMaxBytesLocalHeap(v string)`

SetEhcacheMaxBytesLocalHeap sets EhcacheMaxBytesLocalHeap field to given value.

### HasEhcacheMaxBytesLocalHeap

`func (o *CacheSettings) HasEhcacheMaxBytesLocalHeap() bool`

HasEhcacheMaxBytesLocalHeap returns a boolean if a field has been set.

### GetCacheUniqueId

`func (o *CacheSettings) GetCacheUniqueId() string`

GetCacheUniqueId returns the CacheUniqueId field if non-nil, zero value otherwise.

### GetCacheUniqueIdOk

`func (o *CacheSettings) GetCacheUniqueIdOk() (*string, bool)`

GetCacheUniqueIdOk returns a tuple with the CacheUniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheUniqueId

`func (o *CacheSettings) SetCacheUniqueId(v string)`

SetCacheUniqueId sets CacheUniqueId field to given value.


### GetElasticache

`func (o *CacheSettings) GetElasticache() bool`

GetElasticache returns the Elasticache field if non-nil, zero value otherwise.

### GetElasticacheOk

`func (o *CacheSettings) GetElasticacheOk() (*bool, bool)`

GetElasticacheOk returns a tuple with the Elasticache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElasticache

`func (o *CacheSettings) SetElasticache(v bool)`

SetElasticache sets Elasticache field to given value.

### HasElasticache

`func (o *CacheSettings) HasElasticache() bool`

HasElasticache returns a boolean if a field has been set.

### GetMemcachedEndpoints

`func (o *CacheSettings) GetMemcachedEndpoints() []MemcachedEndpoints`

GetMemcachedEndpoints returns the MemcachedEndpoints field if non-nil, zero value otherwise.

### GetMemcachedEndpointsOk

`func (o *CacheSettings) GetMemcachedEndpointsOk() (*[]MemcachedEndpoints, bool)`

GetMemcachedEndpointsOk returns a tuple with the MemcachedEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemcachedEndpoints

`func (o *CacheSettings) SetMemcachedEndpoints(v []MemcachedEndpoints)`

SetMemcachedEndpoints sets MemcachedEndpoints field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


