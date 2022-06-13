# ProtectSettingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ApiClientId** | Pointer to **string** |  | [optional] 
**ApiClientName** | Pointer to **string** | display name used when creating the API Client in the Jamf Protect web console | [optional] 
**RegistrationId** | Pointer to **string** | ID used when making requests to identify this particular Protect registration. | [optional] 
**ProtectUrl** | Pointer to **string** |  | [optional] 
**LastSyncTime** | Pointer to **string** |  | [optional] 
**SyncStatus** | Pointer to **string** |  | [optional] 
**AutoInstall** | Pointer to **bool** | determines whether the Jamf Protect agent will be automatically installed on client computers | [optional] 

## Methods

### NewProtectSettingsResponse

`func NewProtectSettingsResponse() *ProtectSettingsResponse`

NewProtectSettingsResponse instantiates a new ProtectSettingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectSettingsResponseWithDefaults

`func NewProtectSettingsResponseWithDefaults() *ProtectSettingsResponse`

NewProtectSettingsResponseWithDefaults instantiates a new ProtectSettingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProtectSettingsResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProtectSettingsResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProtectSettingsResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProtectSettingsResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetApiClientId

`func (o *ProtectSettingsResponse) GetApiClientId() string`

GetApiClientId returns the ApiClientId field if non-nil, zero value otherwise.

### GetApiClientIdOk

`func (o *ProtectSettingsResponse) GetApiClientIdOk() (*string, bool)`

GetApiClientIdOk returns a tuple with the ApiClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiClientId

`func (o *ProtectSettingsResponse) SetApiClientId(v string)`

SetApiClientId sets ApiClientId field to given value.

### HasApiClientId

`func (o *ProtectSettingsResponse) HasApiClientId() bool`

HasApiClientId returns a boolean if a field has been set.

### GetApiClientName

`func (o *ProtectSettingsResponse) GetApiClientName() string`

GetApiClientName returns the ApiClientName field if non-nil, zero value otherwise.

### GetApiClientNameOk

`func (o *ProtectSettingsResponse) GetApiClientNameOk() (*string, bool)`

GetApiClientNameOk returns a tuple with the ApiClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiClientName

`func (o *ProtectSettingsResponse) SetApiClientName(v string)`

SetApiClientName sets ApiClientName field to given value.

### HasApiClientName

`func (o *ProtectSettingsResponse) HasApiClientName() bool`

HasApiClientName returns a boolean if a field has been set.

### GetRegistrationId

`func (o *ProtectSettingsResponse) GetRegistrationId() string`

GetRegistrationId returns the RegistrationId field if non-nil, zero value otherwise.

### GetRegistrationIdOk

`func (o *ProtectSettingsResponse) GetRegistrationIdOk() (*string, bool)`

GetRegistrationIdOk returns a tuple with the RegistrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationId

`func (o *ProtectSettingsResponse) SetRegistrationId(v string)`

SetRegistrationId sets RegistrationId field to given value.

### HasRegistrationId

`func (o *ProtectSettingsResponse) HasRegistrationId() bool`

HasRegistrationId returns a boolean if a field has been set.

### GetProtectUrl

`func (o *ProtectSettingsResponse) GetProtectUrl() string`

GetProtectUrl returns the ProtectUrl field if non-nil, zero value otherwise.

### GetProtectUrlOk

`func (o *ProtectSettingsResponse) GetProtectUrlOk() (*string, bool)`

GetProtectUrlOk returns a tuple with the ProtectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectUrl

`func (o *ProtectSettingsResponse) SetProtectUrl(v string)`

SetProtectUrl sets ProtectUrl field to given value.

### HasProtectUrl

`func (o *ProtectSettingsResponse) HasProtectUrl() bool`

HasProtectUrl returns a boolean if a field has been set.

### GetLastSyncTime

`func (o *ProtectSettingsResponse) GetLastSyncTime() string`

GetLastSyncTime returns the LastSyncTime field if non-nil, zero value otherwise.

### GetLastSyncTimeOk

`func (o *ProtectSettingsResponse) GetLastSyncTimeOk() (*string, bool)`

GetLastSyncTimeOk returns a tuple with the LastSyncTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncTime

`func (o *ProtectSettingsResponse) SetLastSyncTime(v string)`

SetLastSyncTime sets LastSyncTime field to given value.

### HasLastSyncTime

`func (o *ProtectSettingsResponse) HasLastSyncTime() bool`

HasLastSyncTime returns a boolean if a field has been set.

### GetSyncStatus

`func (o *ProtectSettingsResponse) GetSyncStatus() string`

GetSyncStatus returns the SyncStatus field if non-nil, zero value otherwise.

### GetSyncStatusOk

`func (o *ProtectSettingsResponse) GetSyncStatusOk() (*string, bool)`

GetSyncStatusOk returns a tuple with the SyncStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncStatus

`func (o *ProtectSettingsResponse) SetSyncStatus(v string)`

SetSyncStatus sets SyncStatus field to given value.

### HasSyncStatus

`func (o *ProtectSettingsResponse) HasSyncStatus() bool`

HasSyncStatus returns a boolean if a field has been set.

### GetAutoInstall

`func (o *ProtectSettingsResponse) GetAutoInstall() bool`

GetAutoInstall returns the AutoInstall field if non-nil, zero value otherwise.

### GetAutoInstallOk

`func (o *ProtectSettingsResponse) GetAutoInstallOk() (*bool, bool)`

GetAutoInstallOk returns a tuple with the AutoInstall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoInstall

`func (o *ProtectSettingsResponse) SetAutoInstall(v bool)`

SetAutoInstall sets AutoInstall field to given value.

### HasAutoInstall

`func (o *ProtectSettingsResponse) HasAutoInstall() bool`

HasAutoInstall returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


