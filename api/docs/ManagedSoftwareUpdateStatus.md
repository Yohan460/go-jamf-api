# ManagedSoftwareUpdateStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OsUpdatesStatusId** | Pointer to **string** |  | [optional] 
**Device** | Pointer to [**ManagedSoftwareUpdateStatusDevice**](ManagedSoftwareUpdateStatusDevice.md) |  | [optional] 
**DownloadPercentComplete** | Pointer to **float32** |  | [optional] 
**Downloaded** | Pointer to **bool** |  | [optional] 
**ProductKey** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**DeferralsRemaining** | Pointer to **int64** | not applicable to all managed software update statuses | [optional] 
**MaxDeferrals** | Pointer to **int64** | not applicable to all managed software update statuses | [optional] 
**NextScheduledInstall** | Pointer to **time.Time** | not applicable to all managed software update statuses | [optional] 
**PastNotifications** | Pointer to [**[]time.Time**](time.Time.md) | not applicable to all managed software update statuses | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 
**Updated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewManagedSoftwareUpdateStatus

`func NewManagedSoftwareUpdateStatus() *ManagedSoftwareUpdateStatus`

NewManagedSoftwareUpdateStatus instantiates a new ManagedSoftwareUpdateStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedSoftwareUpdateStatusWithDefaults

`func NewManagedSoftwareUpdateStatusWithDefaults() *ManagedSoftwareUpdateStatus`

NewManagedSoftwareUpdateStatusWithDefaults instantiates a new ManagedSoftwareUpdateStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOsUpdatesStatusId

`func (o *ManagedSoftwareUpdateStatus) GetOsUpdatesStatusId() string`

GetOsUpdatesStatusId returns the OsUpdatesStatusId field if non-nil, zero value otherwise.

### GetOsUpdatesStatusIdOk

`func (o *ManagedSoftwareUpdateStatus) GetOsUpdatesStatusIdOk() (*string, bool)`

GetOsUpdatesStatusIdOk returns a tuple with the OsUpdatesStatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsUpdatesStatusId

`func (o *ManagedSoftwareUpdateStatus) SetOsUpdatesStatusId(v string)`

SetOsUpdatesStatusId sets OsUpdatesStatusId field to given value.

### HasOsUpdatesStatusId

`func (o *ManagedSoftwareUpdateStatus) HasOsUpdatesStatusId() bool`

HasOsUpdatesStatusId returns a boolean if a field has been set.

### GetDevice

`func (o *ManagedSoftwareUpdateStatus) GetDevice() ManagedSoftwareUpdateStatusDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *ManagedSoftwareUpdateStatus) GetDeviceOk() (*ManagedSoftwareUpdateStatusDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *ManagedSoftwareUpdateStatus) SetDevice(v ManagedSoftwareUpdateStatusDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *ManagedSoftwareUpdateStatus) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetDownloadPercentComplete

`func (o *ManagedSoftwareUpdateStatus) GetDownloadPercentComplete() float32`

GetDownloadPercentComplete returns the DownloadPercentComplete field if non-nil, zero value otherwise.

### GetDownloadPercentCompleteOk

`func (o *ManagedSoftwareUpdateStatus) GetDownloadPercentCompleteOk() (*float32, bool)`

GetDownloadPercentCompleteOk returns a tuple with the DownloadPercentComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadPercentComplete

`func (o *ManagedSoftwareUpdateStatus) SetDownloadPercentComplete(v float32)`

SetDownloadPercentComplete sets DownloadPercentComplete field to given value.

### HasDownloadPercentComplete

`func (o *ManagedSoftwareUpdateStatus) HasDownloadPercentComplete() bool`

HasDownloadPercentComplete returns a boolean if a field has been set.

### GetDownloaded

`func (o *ManagedSoftwareUpdateStatus) GetDownloaded() bool`

GetDownloaded returns the Downloaded field if non-nil, zero value otherwise.

### GetDownloadedOk

`func (o *ManagedSoftwareUpdateStatus) GetDownloadedOk() (*bool, bool)`

GetDownloadedOk returns a tuple with the Downloaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloaded

`func (o *ManagedSoftwareUpdateStatus) SetDownloaded(v bool)`

SetDownloaded sets Downloaded field to given value.

### HasDownloaded

`func (o *ManagedSoftwareUpdateStatus) HasDownloaded() bool`

HasDownloaded returns a boolean if a field has been set.

### GetProductKey

`func (o *ManagedSoftwareUpdateStatus) GetProductKey() string`

GetProductKey returns the ProductKey field if non-nil, zero value otherwise.

### GetProductKeyOk

`func (o *ManagedSoftwareUpdateStatus) GetProductKeyOk() (*string, bool)`

GetProductKeyOk returns a tuple with the ProductKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductKey

`func (o *ManagedSoftwareUpdateStatus) SetProductKey(v string)`

SetProductKey sets ProductKey field to given value.

### HasProductKey

`func (o *ManagedSoftwareUpdateStatus) HasProductKey() bool`

HasProductKey returns a boolean if a field has been set.

### GetStatus

`func (o *ManagedSoftwareUpdateStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ManagedSoftwareUpdateStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ManagedSoftwareUpdateStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ManagedSoftwareUpdateStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDeferralsRemaining

`func (o *ManagedSoftwareUpdateStatus) GetDeferralsRemaining() int64`

GetDeferralsRemaining returns the DeferralsRemaining field if non-nil, zero value otherwise.

### GetDeferralsRemainingOk

`func (o *ManagedSoftwareUpdateStatus) GetDeferralsRemainingOk() (*int64, bool)`

GetDeferralsRemainingOk returns a tuple with the DeferralsRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeferralsRemaining

`func (o *ManagedSoftwareUpdateStatus) SetDeferralsRemaining(v int64)`

SetDeferralsRemaining sets DeferralsRemaining field to given value.

### HasDeferralsRemaining

`func (o *ManagedSoftwareUpdateStatus) HasDeferralsRemaining() bool`

HasDeferralsRemaining returns a boolean if a field has been set.

### GetMaxDeferrals

`func (o *ManagedSoftwareUpdateStatus) GetMaxDeferrals() int64`

GetMaxDeferrals returns the MaxDeferrals field if non-nil, zero value otherwise.

### GetMaxDeferralsOk

`func (o *ManagedSoftwareUpdateStatus) GetMaxDeferralsOk() (*int64, bool)`

GetMaxDeferralsOk returns a tuple with the MaxDeferrals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDeferrals

`func (o *ManagedSoftwareUpdateStatus) SetMaxDeferrals(v int64)`

SetMaxDeferrals sets MaxDeferrals field to given value.

### HasMaxDeferrals

`func (o *ManagedSoftwareUpdateStatus) HasMaxDeferrals() bool`

HasMaxDeferrals returns a boolean if a field has been set.

### GetNextScheduledInstall

`func (o *ManagedSoftwareUpdateStatus) GetNextScheduledInstall() time.Time`

GetNextScheduledInstall returns the NextScheduledInstall field if non-nil, zero value otherwise.

### GetNextScheduledInstallOk

`func (o *ManagedSoftwareUpdateStatus) GetNextScheduledInstallOk() (*time.Time, bool)`

GetNextScheduledInstallOk returns a tuple with the NextScheduledInstall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextScheduledInstall

`func (o *ManagedSoftwareUpdateStatus) SetNextScheduledInstall(v time.Time)`

SetNextScheduledInstall sets NextScheduledInstall field to given value.

### HasNextScheduledInstall

`func (o *ManagedSoftwareUpdateStatus) HasNextScheduledInstall() bool`

HasNextScheduledInstall returns a boolean if a field has been set.

### GetPastNotifications

`func (o *ManagedSoftwareUpdateStatus) GetPastNotifications() []time.Time`

GetPastNotifications returns the PastNotifications field if non-nil, zero value otherwise.

### GetPastNotificationsOk

`func (o *ManagedSoftwareUpdateStatus) GetPastNotificationsOk() (*[]time.Time, bool)`

GetPastNotificationsOk returns a tuple with the PastNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPastNotifications

`func (o *ManagedSoftwareUpdateStatus) SetPastNotifications(v []time.Time)`

SetPastNotifications sets PastNotifications field to given value.

### HasPastNotifications

`func (o *ManagedSoftwareUpdateStatus) HasPastNotifications() bool`

HasPastNotifications returns a boolean if a field has been set.

### GetCreated

`func (o *ManagedSoftwareUpdateStatus) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ManagedSoftwareUpdateStatus) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ManagedSoftwareUpdateStatus) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ManagedSoftwareUpdateStatus) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *ManagedSoftwareUpdateStatus) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ManagedSoftwareUpdateStatus) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ManagedSoftwareUpdateStatus) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *ManagedSoftwareUpdateStatus) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


