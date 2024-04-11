# SelfServiceInteractionSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationsEnabled** | Pointer to **bool** | global Self Service setting for if notifications are on or off  | [optional] [default to false]
**AlertUserApprovedMdm** | Pointer to **bool** | whether users should be notified they need to approve organization&#39;s MDM profile  | [optional] [default to true]
**DefaultLandingPage** | Pointer to **string** | the default landing page in Self Service  | [optional] [default to "HOME"]
**DefaultHomeCategoryId** | Pointer to **int64** | id for the default home category in Self Service  | [optional] [default to -1]
**BookmarksName** | **string** | renamed string for bookmarks if the admin wishes  | 

## Methods

### NewSelfServiceInteractionSettings

`func NewSelfServiceInteractionSettings(bookmarksName string, ) *SelfServiceInteractionSettings`

NewSelfServiceInteractionSettings instantiates a new SelfServiceInteractionSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelfServiceInteractionSettingsWithDefaults

`func NewSelfServiceInteractionSettingsWithDefaults() *SelfServiceInteractionSettings`

NewSelfServiceInteractionSettingsWithDefaults instantiates a new SelfServiceInteractionSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotificationsEnabled

`func (o *SelfServiceInteractionSettings) GetNotificationsEnabled() bool`

GetNotificationsEnabled returns the NotificationsEnabled field if non-nil, zero value otherwise.

### GetNotificationsEnabledOk

`func (o *SelfServiceInteractionSettings) GetNotificationsEnabledOk() (*bool, bool)`

GetNotificationsEnabledOk returns a tuple with the NotificationsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationsEnabled

`func (o *SelfServiceInteractionSettings) SetNotificationsEnabled(v bool)`

SetNotificationsEnabled sets NotificationsEnabled field to given value.

### HasNotificationsEnabled

`func (o *SelfServiceInteractionSettings) HasNotificationsEnabled() bool`

HasNotificationsEnabled returns a boolean if a field has been set.

### GetAlertUserApprovedMdm

`func (o *SelfServiceInteractionSettings) GetAlertUserApprovedMdm() bool`

GetAlertUserApprovedMdm returns the AlertUserApprovedMdm field if non-nil, zero value otherwise.

### GetAlertUserApprovedMdmOk

`func (o *SelfServiceInteractionSettings) GetAlertUserApprovedMdmOk() (*bool, bool)`

GetAlertUserApprovedMdmOk returns a tuple with the AlertUserApprovedMdm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertUserApprovedMdm

`func (o *SelfServiceInteractionSettings) SetAlertUserApprovedMdm(v bool)`

SetAlertUserApprovedMdm sets AlertUserApprovedMdm field to given value.

### HasAlertUserApprovedMdm

`func (o *SelfServiceInteractionSettings) HasAlertUserApprovedMdm() bool`

HasAlertUserApprovedMdm returns a boolean if a field has been set.

### GetDefaultLandingPage

`func (o *SelfServiceInteractionSettings) GetDefaultLandingPage() string`

GetDefaultLandingPage returns the DefaultLandingPage field if non-nil, zero value otherwise.

### GetDefaultLandingPageOk

`func (o *SelfServiceInteractionSettings) GetDefaultLandingPageOk() (*string, bool)`

GetDefaultLandingPageOk returns a tuple with the DefaultLandingPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLandingPage

`func (o *SelfServiceInteractionSettings) SetDefaultLandingPage(v string)`

SetDefaultLandingPage sets DefaultLandingPage field to given value.

### HasDefaultLandingPage

`func (o *SelfServiceInteractionSettings) HasDefaultLandingPage() bool`

HasDefaultLandingPage returns a boolean if a field has been set.

### GetDefaultHomeCategoryId

`func (o *SelfServiceInteractionSettings) GetDefaultHomeCategoryId() int64`

GetDefaultHomeCategoryId returns the DefaultHomeCategoryId field if non-nil, zero value otherwise.

### GetDefaultHomeCategoryIdOk

`func (o *SelfServiceInteractionSettings) GetDefaultHomeCategoryIdOk() (*int64, bool)`

GetDefaultHomeCategoryIdOk returns a tuple with the DefaultHomeCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultHomeCategoryId

`func (o *SelfServiceInteractionSettings) SetDefaultHomeCategoryId(v int64)`

SetDefaultHomeCategoryId sets DefaultHomeCategoryId field to given value.

### HasDefaultHomeCategoryId

`func (o *SelfServiceInteractionSettings) HasDefaultHomeCategoryId() bool`

HasDefaultHomeCategoryId returns a boolean if a field has been set.

### GetBookmarksName

`func (o *SelfServiceInteractionSettings) GetBookmarksName() string`

GetBookmarksName returns the BookmarksName field if non-nil, zero value otherwise.

### GetBookmarksNameOk

`func (o *SelfServiceInteractionSettings) GetBookmarksNameOk() (*string, bool)`

GetBookmarksNameOk returns a tuple with the BookmarksName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookmarksName

`func (o *SelfServiceInteractionSettings) SetBookmarksName(v string)`

SetBookmarksName sets BookmarksName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


