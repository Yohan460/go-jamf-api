# NotificationV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**NotificationType**](NotificationType.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Params** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewNotificationV1

`func NewNotificationV1() *NotificationV1`

NewNotificationV1 instantiates a new NotificationV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationV1WithDefaults

`func NewNotificationV1WithDefaults() *NotificationV1`

NewNotificationV1WithDefaults instantiates a new NotificationV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NotificationV1) GetType() NotificationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotificationV1) GetTypeOk() (*NotificationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotificationV1) SetType(v NotificationType)`

SetType sets Type field to given value.

### HasType

`func (o *NotificationV1) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *NotificationV1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationV1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationV1) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NotificationV1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetParams

`func (o *NotificationV1) GetParams() map[string]map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *NotificationV1) GetParamsOk() (*map[string]map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *NotificationV1) SetParams(v map[string]map[string]interface{})`

SetParams sets Params field to given value.

### HasParams

`func (o *NotificationV1) HasParams() bool`

HasParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


