# Notification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Params** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewNotification

`func NewNotification() *Notification`

NewNotification instantiates a new Notification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationWithDefaults

`func NewNotificationWithDefaults() *Notification`

NewNotificationWithDefaults instantiates a new Notification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Notification) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Notification) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Notification) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Notification) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *Notification) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Notification) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Notification) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Notification) HasId() bool`

HasId returns a boolean if a field has been set.

### GetParams

`func (o *Notification) GetParams() map[string]map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *Notification) GetParamsOk() (*map[string]map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *Notification) SetParams(v map[string]map[string]interface{})`

SetParams sets Params field to given value.

### HasParams

`func (o *Notification) HasParams() bool`

HasParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


