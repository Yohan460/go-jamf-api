# ConnectionConfigurationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | An identifier of connection configuration for Team Viewer Remote Administration | 
**SiteId** | **string** | An identifier of a site which Team Viewer Remote Administration is configured on | 
**DisplayName** | **string** | Name for Team Viewer Connection Configuration | 
**Enabled** | **bool** | Describes if Team Viewer connection is enabled or disabled | 
**SessionTimeout** | **int32** | Number of minutes before the session expires | 

## Methods

### NewConnectionConfigurationResponse

`func NewConnectionConfigurationResponse(id string, siteId string, displayName string, enabled bool, sessionTimeout int32, ) *ConnectionConfigurationResponse`

NewConnectionConfigurationResponse instantiates a new ConnectionConfigurationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionConfigurationResponseWithDefaults

`func NewConnectionConfigurationResponseWithDefaults() *ConnectionConfigurationResponse`

NewConnectionConfigurationResponseWithDefaults instantiates a new ConnectionConfigurationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConnectionConfigurationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectionConfigurationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectionConfigurationResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSiteId

`func (o *ConnectionConfigurationResponse) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *ConnectionConfigurationResponse) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *ConnectionConfigurationResponse) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.


### GetDisplayName

`func (o *ConnectionConfigurationResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ConnectionConfigurationResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ConnectionConfigurationResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetEnabled

`func (o *ConnectionConfigurationResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ConnectionConfigurationResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ConnectionConfigurationResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetSessionTimeout

`func (o *ConnectionConfigurationResponse) GetSessionTimeout() int32`

GetSessionTimeout returns the SessionTimeout field if non-nil, zero value otherwise.

### GetSessionTimeoutOk

`func (o *ConnectionConfigurationResponse) GetSessionTimeoutOk() (*int32, bool)`

GetSessionTimeoutOk returns a tuple with the SessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTimeout

`func (o *ConnectionConfigurationResponse) SetSessionTimeout(v int32)`

SetSessionTimeout sets SessionTimeout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


