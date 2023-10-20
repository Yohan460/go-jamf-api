# VolumePurchasingLocationListView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**TotalPurchasedLicenses** | Pointer to **int32** |  | [optional] [readonly] 
**TotalUsedLicenses** | Pointer to **int32** |  | [optional] [readonly] 
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
**LastSyncTime** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewVolumePurchasingLocationListView

`func NewVolumePurchasingLocationListView() *VolumePurchasingLocationListView`

NewVolumePurchasingLocationListView instantiates a new VolumePurchasingLocationListView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumePurchasingLocationListViewWithDefaults

`func NewVolumePurchasingLocationListViewWithDefaults() *VolumePurchasingLocationListView`

NewVolumePurchasingLocationListViewWithDefaults instantiates a new VolumePurchasingLocationListView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VolumePurchasingLocationListView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VolumePurchasingLocationListView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VolumePurchasingLocationListView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VolumePurchasingLocationListView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTotalPurchasedLicenses

`func (o *VolumePurchasingLocationListView) GetTotalPurchasedLicenses() int32`

GetTotalPurchasedLicenses returns the TotalPurchasedLicenses field if non-nil, zero value otherwise.

### GetTotalPurchasedLicensesOk

`func (o *VolumePurchasingLocationListView) GetTotalPurchasedLicensesOk() (*int32, bool)`

GetTotalPurchasedLicensesOk returns a tuple with the TotalPurchasedLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPurchasedLicenses

`func (o *VolumePurchasingLocationListView) SetTotalPurchasedLicenses(v int32)`

SetTotalPurchasedLicenses sets TotalPurchasedLicenses field to given value.

### HasTotalPurchasedLicenses

`func (o *VolumePurchasingLocationListView) HasTotalPurchasedLicenses() bool`

HasTotalPurchasedLicenses returns a boolean if a field has been set.

### GetTotalUsedLicenses

`func (o *VolumePurchasingLocationListView) GetTotalUsedLicenses() int32`

GetTotalUsedLicenses returns the TotalUsedLicenses field if non-nil, zero value otherwise.

### GetTotalUsedLicensesOk

`func (o *VolumePurchasingLocationListView) GetTotalUsedLicensesOk() (*int32, bool)`

GetTotalUsedLicensesOk returns a tuple with the TotalUsedLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsedLicenses

`func (o *VolumePurchasingLocationListView) SetTotalUsedLicenses(v int32)`

SetTotalUsedLicenses sets TotalUsedLicenses field to given value.

### HasTotalUsedLicenses

`func (o *VolumePurchasingLocationListView) HasTotalUsedLicenses() bool`

HasTotalUsedLicenses returns a boolean if a field has been set.

### GetId

`func (o *VolumePurchasingLocationListView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VolumePurchasingLocationListView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VolumePurchasingLocationListView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VolumePurchasingLocationListView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAppleId

`func (o *VolumePurchasingLocationListView) GetAppleId() string`

GetAppleId returns the AppleId field if non-nil, zero value otherwise.

### GetAppleIdOk

`func (o *VolumePurchasingLocationListView) GetAppleIdOk() (*string, bool)`

GetAppleIdOk returns a tuple with the AppleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppleId

`func (o *VolumePurchasingLocationListView) SetAppleId(v string)`

SetAppleId sets AppleId field to given value.

### HasAppleId

`func (o *VolumePurchasingLocationListView) HasAppleId() bool`

HasAppleId returns a boolean if a field has been set.

### GetOrganizationName

`func (o *VolumePurchasingLocationListView) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *VolumePurchasingLocationListView) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *VolumePurchasingLocationListView) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *VolumePurchasingLocationListView) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### GetTokenExpiration

`func (o *VolumePurchasingLocationListView) GetTokenExpiration() string`

GetTokenExpiration returns the TokenExpiration field if non-nil, zero value otherwise.

### GetTokenExpirationOk

`func (o *VolumePurchasingLocationListView) GetTokenExpirationOk() (*string, bool)`

GetTokenExpirationOk returns a tuple with the TokenExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExpiration

`func (o *VolumePurchasingLocationListView) SetTokenExpiration(v string)`

SetTokenExpiration sets TokenExpiration field to given value.

### HasTokenExpiration

`func (o *VolumePurchasingLocationListView) HasTokenExpiration() bool`

HasTokenExpiration returns a boolean if a field has been set.

### GetCountryCode

`func (o *VolumePurchasingLocationListView) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *VolumePurchasingLocationListView) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *VolumePurchasingLocationListView) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *VolumePurchasingLocationListView) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetLocationName

`func (o *VolumePurchasingLocationListView) GetLocationName() string`

GetLocationName returns the LocationName field if non-nil, zero value otherwise.

### GetLocationNameOk

`func (o *VolumePurchasingLocationListView) GetLocationNameOk() (*string, bool)`

GetLocationNameOk returns a tuple with the LocationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationName

`func (o *VolumePurchasingLocationListView) SetLocationName(v string)`

SetLocationName sets LocationName field to given value.

### HasLocationName

`func (o *VolumePurchasingLocationListView) HasLocationName() bool`

HasLocationName returns a boolean if a field has been set.

### GetClientContextMismatch

`func (o *VolumePurchasingLocationListView) GetClientContextMismatch() bool`

GetClientContextMismatch returns the ClientContextMismatch field if non-nil, zero value otherwise.

### GetClientContextMismatchOk

`func (o *VolumePurchasingLocationListView) GetClientContextMismatchOk() (*bool, bool)`

GetClientContextMismatchOk returns a tuple with the ClientContextMismatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientContextMismatch

`func (o *VolumePurchasingLocationListView) SetClientContextMismatch(v bool)`

SetClientContextMismatch sets ClientContextMismatch field to given value.

### HasClientContextMismatch

`func (o *VolumePurchasingLocationListView) HasClientContextMismatch() bool`

HasClientContextMismatch returns a boolean if a field has been set.

### GetAutomaticallyPopulatePurchasedContent

`func (o *VolumePurchasingLocationListView) GetAutomaticallyPopulatePurchasedContent() bool`

GetAutomaticallyPopulatePurchasedContent returns the AutomaticallyPopulatePurchasedContent field if non-nil, zero value otherwise.

### GetAutomaticallyPopulatePurchasedContentOk

`func (o *VolumePurchasingLocationListView) GetAutomaticallyPopulatePurchasedContentOk() (*bool, bool)`

GetAutomaticallyPopulatePurchasedContentOk returns a tuple with the AutomaticallyPopulatePurchasedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticallyPopulatePurchasedContent

`func (o *VolumePurchasingLocationListView) SetAutomaticallyPopulatePurchasedContent(v bool)`

SetAutomaticallyPopulatePurchasedContent sets AutomaticallyPopulatePurchasedContent field to given value.

### HasAutomaticallyPopulatePurchasedContent

`func (o *VolumePurchasingLocationListView) HasAutomaticallyPopulatePurchasedContent() bool`

HasAutomaticallyPopulatePurchasedContent returns a boolean if a field has been set.

### GetSendNotificationWhenNoLongerAssigned

`func (o *VolumePurchasingLocationListView) GetSendNotificationWhenNoLongerAssigned() bool`

GetSendNotificationWhenNoLongerAssigned returns the SendNotificationWhenNoLongerAssigned field if non-nil, zero value otherwise.

### GetSendNotificationWhenNoLongerAssignedOk

`func (o *VolumePurchasingLocationListView) GetSendNotificationWhenNoLongerAssignedOk() (*bool, bool)`

GetSendNotificationWhenNoLongerAssignedOk returns a tuple with the SendNotificationWhenNoLongerAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendNotificationWhenNoLongerAssigned

`func (o *VolumePurchasingLocationListView) SetSendNotificationWhenNoLongerAssigned(v bool)`

SetSendNotificationWhenNoLongerAssigned sets SendNotificationWhenNoLongerAssigned field to given value.

### HasSendNotificationWhenNoLongerAssigned

`func (o *VolumePurchasingLocationListView) HasSendNotificationWhenNoLongerAssigned() bool`

HasSendNotificationWhenNoLongerAssigned returns a boolean if a field has been set.

### GetAutoRegisterManagedUsers

`func (o *VolumePurchasingLocationListView) GetAutoRegisterManagedUsers() bool`

GetAutoRegisterManagedUsers returns the AutoRegisterManagedUsers field if non-nil, zero value otherwise.

### GetAutoRegisterManagedUsersOk

`func (o *VolumePurchasingLocationListView) GetAutoRegisterManagedUsersOk() (*bool, bool)`

GetAutoRegisterManagedUsersOk returns a tuple with the AutoRegisterManagedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRegisterManagedUsers

`func (o *VolumePurchasingLocationListView) SetAutoRegisterManagedUsers(v bool)`

SetAutoRegisterManagedUsers sets AutoRegisterManagedUsers field to given value.

### HasAutoRegisterManagedUsers

`func (o *VolumePurchasingLocationListView) HasAutoRegisterManagedUsers() bool`

HasAutoRegisterManagedUsers returns a boolean if a field has been set.

### GetSiteId

`func (o *VolumePurchasingLocationListView) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *VolumePurchasingLocationListView) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *VolumePurchasingLocationListView) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *VolumePurchasingLocationListView) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetLastSyncTime

`func (o *VolumePurchasingLocationListView) GetLastSyncTime() string`

GetLastSyncTime returns the LastSyncTime field if non-nil, zero value otherwise.

### GetLastSyncTimeOk

`func (o *VolumePurchasingLocationListView) GetLastSyncTimeOk() (*string, bool)`

GetLastSyncTimeOk returns a tuple with the LastSyncTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncTime

`func (o *VolumePurchasingLocationListView) SetLastSyncTime(v string)`

SetLastSyncTime sets LastSyncTime field to given value.

### HasLastSyncTime

`func (o *VolumePurchasingLocationListView) HasLastSyncTime() bool`

HasLastSyncTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


