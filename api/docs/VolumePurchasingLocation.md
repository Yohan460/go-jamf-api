# VolumePurchasingLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**TotalPurchasedLicenses** | Pointer to **int64** |  | [optional] [readonly] 
**TotalUsedLicenses** | Pointer to **int64** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**AppleId** | Pointer to **string** |  | [optional] [readonly] 
**OrganizationName** | Pointer to **string** |  | [optional] [readonly] 
**TokenExpiration** | Pointer to **string** |  | [optional] [readonly] 
**CountryCode** | Pointer to **string** | The two-letter ISO 3166-1 code that designates the country where the Volume Purchasing account is located. | [optional] [readonly] 
**LocationName** | Pointer to **string** |  | [optional] [readonly] 
**ClientContextMismatch** | Pointer to **bool** | If this is \&quot;true\&quot;, the clientContext used by this server does not match the clientContext returned by the Volume Purchasing API. | [optional] [readonly] 
**AutomaticallyPopulatePurchasedContent** | Pointer to **bool** |  | [optional] 
**SendNotificationWhenNoLongerAssigned** | Pointer to **bool** |  | [optional] 
**AutoRegisterManagedUsers** | Pointer to **bool** |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] 
**SiteName** | Pointer to **string** |  | [optional] 
**LastSyncTime** | Pointer to **string** |  | [optional] [readonly] 
**Content** | Pointer to [**[]VolumePurchasingContent**](VolumePurchasingContent.md) |  | [optional] [readonly] 

## Methods

### NewVolumePurchasingLocation

`func NewVolumePurchasingLocation() *VolumePurchasingLocation`

NewVolumePurchasingLocation instantiates a new VolumePurchasingLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumePurchasingLocationWithDefaults

`func NewVolumePurchasingLocationWithDefaults() *VolumePurchasingLocation`

NewVolumePurchasingLocationWithDefaults instantiates a new VolumePurchasingLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VolumePurchasingLocation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VolumePurchasingLocation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VolumePurchasingLocation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VolumePurchasingLocation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTotalPurchasedLicenses

`func (o *VolumePurchasingLocation) GetTotalPurchasedLicenses() int64`

GetTotalPurchasedLicenses returns the TotalPurchasedLicenses field if non-nil, zero value otherwise.

### GetTotalPurchasedLicensesOk

`func (o *VolumePurchasingLocation) GetTotalPurchasedLicensesOk() (*int64, bool)`

GetTotalPurchasedLicensesOk returns a tuple with the TotalPurchasedLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPurchasedLicenses

`func (o *VolumePurchasingLocation) SetTotalPurchasedLicenses(v int64)`

SetTotalPurchasedLicenses sets TotalPurchasedLicenses field to given value.

### HasTotalPurchasedLicenses

`func (o *VolumePurchasingLocation) HasTotalPurchasedLicenses() bool`

HasTotalPurchasedLicenses returns a boolean if a field has been set.

### GetTotalUsedLicenses

`func (o *VolumePurchasingLocation) GetTotalUsedLicenses() int64`

GetTotalUsedLicenses returns the TotalUsedLicenses field if non-nil, zero value otherwise.

### GetTotalUsedLicensesOk

`func (o *VolumePurchasingLocation) GetTotalUsedLicensesOk() (*int64, bool)`

GetTotalUsedLicensesOk returns a tuple with the TotalUsedLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsedLicenses

`func (o *VolumePurchasingLocation) SetTotalUsedLicenses(v int64)`

SetTotalUsedLicenses sets TotalUsedLicenses field to given value.

### HasTotalUsedLicenses

`func (o *VolumePurchasingLocation) HasTotalUsedLicenses() bool`

HasTotalUsedLicenses returns a boolean if a field has been set.

### GetId

`func (o *VolumePurchasingLocation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VolumePurchasingLocation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VolumePurchasingLocation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VolumePurchasingLocation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAppleId

`func (o *VolumePurchasingLocation) GetAppleId() string`

GetAppleId returns the AppleId field if non-nil, zero value otherwise.

### GetAppleIdOk

`func (o *VolumePurchasingLocation) GetAppleIdOk() (*string, bool)`

GetAppleIdOk returns a tuple with the AppleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppleId

`func (o *VolumePurchasingLocation) SetAppleId(v string)`

SetAppleId sets AppleId field to given value.

### HasAppleId

`func (o *VolumePurchasingLocation) HasAppleId() bool`

HasAppleId returns a boolean if a field has been set.

### GetOrganizationName

`func (o *VolumePurchasingLocation) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *VolumePurchasingLocation) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *VolumePurchasingLocation) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *VolumePurchasingLocation) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### GetTokenExpiration

`func (o *VolumePurchasingLocation) GetTokenExpiration() string`

GetTokenExpiration returns the TokenExpiration field if non-nil, zero value otherwise.

### GetTokenExpirationOk

`func (o *VolumePurchasingLocation) GetTokenExpirationOk() (*string, bool)`

GetTokenExpirationOk returns a tuple with the TokenExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExpiration

`func (o *VolumePurchasingLocation) SetTokenExpiration(v string)`

SetTokenExpiration sets TokenExpiration field to given value.

### HasTokenExpiration

`func (o *VolumePurchasingLocation) HasTokenExpiration() bool`

HasTokenExpiration returns a boolean if a field has been set.

### GetCountryCode

`func (o *VolumePurchasingLocation) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *VolumePurchasingLocation) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *VolumePurchasingLocation) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *VolumePurchasingLocation) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetLocationName

`func (o *VolumePurchasingLocation) GetLocationName() string`

GetLocationName returns the LocationName field if non-nil, zero value otherwise.

### GetLocationNameOk

`func (o *VolumePurchasingLocation) GetLocationNameOk() (*string, bool)`

GetLocationNameOk returns a tuple with the LocationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationName

`func (o *VolumePurchasingLocation) SetLocationName(v string)`

SetLocationName sets LocationName field to given value.

### HasLocationName

`func (o *VolumePurchasingLocation) HasLocationName() bool`

HasLocationName returns a boolean if a field has been set.

### GetClientContextMismatch

`func (o *VolumePurchasingLocation) GetClientContextMismatch() bool`

GetClientContextMismatch returns the ClientContextMismatch field if non-nil, zero value otherwise.

### GetClientContextMismatchOk

`func (o *VolumePurchasingLocation) GetClientContextMismatchOk() (*bool, bool)`

GetClientContextMismatchOk returns a tuple with the ClientContextMismatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientContextMismatch

`func (o *VolumePurchasingLocation) SetClientContextMismatch(v bool)`

SetClientContextMismatch sets ClientContextMismatch field to given value.

### HasClientContextMismatch

`func (o *VolumePurchasingLocation) HasClientContextMismatch() bool`

HasClientContextMismatch returns a boolean if a field has been set.

### GetAutomaticallyPopulatePurchasedContent

`func (o *VolumePurchasingLocation) GetAutomaticallyPopulatePurchasedContent() bool`

GetAutomaticallyPopulatePurchasedContent returns the AutomaticallyPopulatePurchasedContent field if non-nil, zero value otherwise.

### GetAutomaticallyPopulatePurchasedContentOk

`func (o *VolumePurchasingLocation) GetAutomaticallyPopulatePurchasedContentOk() (*bool, bool)`

GetAutomaticallyPopulatePurchasedContentOk returns a tuple with the AutomaticallyPopulatePurchasedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticallyPopulatePurchasedContent

`func (o *VolumePurchasingLocation) SetAutomaticallyPopulatePurchasedContent(v bool)`

SetAutomaticallyPopulatePurchasedContent sets AutomaticallyPopulatePurchasedContent field to given value.

### HasAutomaticallyPopulatePurchasedContent

`func (o *VolumePurchasingLocation) HasAutomaticallyPopulatePurchasedContent() bool`

HasAutomaticallyPopulatePurchasedContent returns a boolean if a field has been set.

### GetSendNotificationWhenNoLongerAssigned

`func (o *VolumePurchasingLocation) GetSendNotificationWhenNoLongerAssigned() bool`

GetSendNotificationWhenNoLongerAssigned returns the SendNotificationWhenNoLongerAssigned field if non-nil, zero value otherwise.

### GetSendNotificationWhenNoLongerAssignedOk

`func (o *VolumePurchasingLocation) GetSendNotificationWhenNoLongerAssignedOk() (*bool, bool)`

GetSendNotificationWhenNoLongerAssignedOk returns a tuple with the SendNotificationWhenNoLongerAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendNotificationWhenNoLongerAssigned

`func (o *VolumePurchasingLocation) SetSendNotificationWhenNoLongerAssigned(v bool)`

SetSendNotificationWhenNoLongerAssigned sets SendNotificationWhenNoLongerAssigned field to given value.

### HasSendNotificationWhenNoLongerAssigned

`func (o *VolumePurchasingLocation) HasSendNotificationWhenNoLongerAssigned() bool`

HasSendNotificationWhenNoLongerAssigned returns a boolean if a field has been set.

### GetAutoRegisterManagedUsers

`func (o *VolumePurchasingLocation) GetAutoRegisterManagedUsers() bool`

GetAutoRegisterManagedUsers returns the AutoRegisterManagedUsers field if non-nil, zero value otherwise.

### GetAutoRegisterManagedUsersOk

`func (o *VolumePurchasingLocation) GetAutoRegisterManagedUsersOk() (*bool, bool)`

GetAutoRegisterManagedUsersOk returns a tuple with the AutoRegisterManagedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRegisterManagedUsers

`func (o *VolumePurchasingLocation) SetAutoRegisterManagedUsers(v bool)`

SetAutoRegisterManagedUsers sets AutoRegisterManagedUsers field to given value.

### HasAutoRegisterManagedUsers

`func (o *VolumePurchasingLocation) HasAutoRegisterManagedUsers() bool`

HasAutoRegisterManagedUsers returns a boolean if a field has been set.

### GetSiteId

`func (o *VolumePurchasingLocation) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *VolumePurchasingLocation) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *VolumePurchasingLocation) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *VolumePurchasingLocation) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSiteName

`func (o *VolumePurchasingLocation) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *VolumePurchasingLocation) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *VolumePurchasingLocation) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *VolumePurchasingLocation) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetLastSyncTime

`func (o *VolumePurchasingLocation) GetLastSyncTime() string`

GetLastSyncTime returns the LastSyncTime field if non-nil, zero value otherwise.

### GetLastSyncTimeOk

`func (o *VolumePurchasingLocation) GetLastSyncTimeOk() (*string, bool)`

GetLastSyncTimeOk returns a tuple with the LastSyncTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncTime

`func (o *VolumePurchasingLocation) SetLastSyncTime(v string)`

SetLastSyncTime sets LastSyncTime field to given value.

### HasLastSyncTime

`func (o *VolumePurchasingLocation) HasLastSyncTime() bool`

HasLastSyncTime returns a boolean if a field has been set.

### GetContent

`func (o *VolumePurchasingLocation) GetContent() []VolumePurchasingContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *VolumePurchasingLocation) GetContentOk() (*[]VolumePurchasingContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *VolumePurchasingLocation) SetContent(v []VolumePurchasingContent)`

SetContent sets Content field to given value.

### HasContent

`func (o *VolumePurchasingLocation) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


