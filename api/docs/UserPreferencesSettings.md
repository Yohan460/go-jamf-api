# UserPreferencesSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** |  | 
**Key** | **string** |  | 
**Values** | **[]string** | List of preferences for the specific key and user. | 

## Methods

### NewUserPreferencesSettings

`func NewUserPreferencesSettings(username string, key string, values []string, ) *UserPreferencesSettings`

NewUserPreferencesSettings instantiates a new UserPreferencesSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPreferencesSettingsWithDefaults

`func NewUserPreferencesSettingsWithDefaults() *UserPreferencesSettings`

NewUserPreferencesSettingsWithDefaults instantiates a new UserPreferencesSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *UserPreferencesSettings) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserPreferencesSettings) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserPreferencesSettings) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetKey

`func (o *UserPreferencesSettings) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UserPreferencesSettings) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UserPreferencesSettings) SetKey(v string)`

SetKey sets Key field to given value.


### GetValues

`func (o *UserPreferencesSettings) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *UserPreferencesSettings) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *UserPreferencesSettings) SetValues(v []string)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


