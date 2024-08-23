/*
Jamf Pro API

## Overview The Jamf Pro API is a RESTful API for Jamf Pro built to enable consistent and efficient programmatic access to Jamf Pro.<br/><br/> The swagger schema can be found [here](/api/schema/). 

API version: production
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the VolumePurchasingLocationListView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VolumePurchasingLocationListView{}

// VolumePurchasingLocationListView struct for VolumePurchasingLocationListView
type VolumePurchasingLocationListView struct {
	Name *string `json:"name,omitempty"`
	TotalPurchasedLicenses *int64 `json:"totalPurchasedLicenses,omitempty"`
	TotalUsedLicenses *int64 `json:"totalUsedLicenses,omitempty"`
	Id *string `json:"id,omitempty"`
	AppleId *string `json:"appleId,omitempty"`
	OrganizationName *string `json:"organizationName,omitempty"`
	TokenExpiration *string `json:"tokenExpiration,omitempty"`
	// The two-letter ISO 3166-1 code that designates the country where the Volume Purchasing account is located.
	CountryCode *string `json:"countryCode,omitempty"`
	LocationName *string `json:"locationName,omitempty"`
	// If this is \"true\", the clientContext used by this server does not match the clientContext returned by the Volume Purchasing API.
	ClientContextMismatch *bool `json:"clientContextMismatch,omitempty"`
	AutomaticallyPopulatePurchasedContent *bool `json:"automaticallyPopulatePurchasedContent,omitempty"`
	SendNotificationWhenNoLongerAssigned *bool `json:"sendNotificationWhenNoLongerAssigned,omitempty"`
	AutoRegisterManagedUsers *bool `json:"autoRegisterManagedUsers,omitempty"`
	SiteId *string `json:"siteId,omitempty"`
	SiteName *string `json:"siteName,omitempty"`
	LastSyncTime *string `json:"lastSyncTime,omitempty"`
}

// NewVolumePurchasingLocationListView instantiates a new VolumePurchasingLocationListView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolumePurchasingLocationListView() *VolumePurchasingLocationListView {
	this := VolumePurchasingLocationListView{}
	return &this
}

// NewVolumePurchasingLocationListViewWithDefaults instantiates a new VolumePurchasingLocationListView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumePurchasingLocationListViewWithDefaults() *VolumePurchasingLocationListView {
	this := VolumePurchasingLocationListView{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VolumePurchasingLocationListView) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumePurchasingLocationListView) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VolumePurchasingLocationListView) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VolumePurchasingLocationListView) SetName(v string) {
	o.Name = &v
}

// GetTotalPurchasedLicenses returns the TotalPurchasedLicenses field value if set, zero value otherwise.
func (o *VolumePurchasingLocationListView) GetTotalPurchasedLicenses() int64 {
	if o == nil || IsNil(o.TotalPurchasedLicenses) {
		var ret int64
		return ret
	}
	return *o.TotalPurchasedLicenses
}

// GetTotalPurchasedLicensesOk returns a tuple with the TotalPurchasedLicenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumePurchasingLocationListView) GetTotalPurchasedLicensesOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalPurchasedLicenses) {
		return nil, false
	}
	return o.TotalPurchasedLicenses, true
}

// HasTotalPurchasedLicenses returns a boolean if a field has been set.
func (o *VolumePurchasingLocationListView) HasTotalPurchasedLicenses() bool {
	if o != nil && !IsNil(o.TotalPurchasedLicenses) {
		return true
	}

	return false
}

// SetTotalPurchasedLicenses gets a reference to the given int64 and assigns it to the TotalPurchasedLicenses field.
func (o *VolumePurchasingLocationListView) SetTotalPurchasedLicenses(v int64) {
	o.TotalPurchasedLicenses = &v
}

// GetTotalUsedLicenses returns the TotalUsedLicenses field value if set, zero value otherwise.
func (o *VolumePurchasingLocationListView) GetTotalUsedLicenses() int64 {
	if o == nil || IsNil(o.TotalUsedLicenses) {
		var ret int64
		return ret
	}
	return *o.TotalUsedLicenses
}

// GetTotalUsedLicensesOk returns a tuple with the TotalUsedLicenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumePurchasingLocationListView) GetTotalUsedLicensesOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalUsedLicenses) {
		return nil, false
	}
	return o.TotalUsedLicenses, true
}

// HasTotalUsedLicenses returns a boolean if a field has been set.
func (o *VolumePurchasingLocationListView) HasTotalUsedLicenses() bool {
	if o != nil && !IsNil(o.TotalUsedLicenses) {
		return true
	}

	return false
}

// SetTotalUsedLicenses gets a reference to the given int64 and assigns it to the TotalUsedLicenses field.
func (o *VolumePurchasingLocationListView) SetTotalUsedLicenses(v int64) {
	o.TotalUsedLicenses = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VolumePurchasingLocationListView) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumePurchasingLocationListView) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VolumePurchasingLocationListView) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *VolumePurchasingLocationListView) SetId(v string) {
	o.Id = &v
}

// GetAppleId returns the AppleId field value if set, zero value otherwise.
func (o *VolumePurchasingLocationListView) GetAppleId() string {
	if o == nil || IsNil(o.AppleId) {
		var ret string
		return ret
	}
	return *o.AppleId
}

// GetAppleIdOk returns a tuple with the AppleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumePurchasingLocationListView) GetAppleIdOk() (*string, bool) {
	if o == nil || IsNil(o.AppleId) {
		return nil, false
	}
	return o.AppleId, true
}

// HasAppleId returns a boolean if a field has been set.
func (o *VolumePurchasingLocationListView) HasAppleId() bool {
	if o != nil && !IsNil(o.AppleId) {
		return true
	}

	return false
}

// SetAppleId gets a reference to the given string and assigns it to the AppleId field.
func (o *VolumePurchasingLocationListView) SetAppleId(v string) {
	o.AppleId = &v
}

// GetOrganizationName returns the OrganizationName field value if set, zero value otherwise.
func (o *VolumePurchasingLocationListView) GetOrganizationName() string {
	if o == nil || IsNil(o.OrganizationName) {
		var ret string
		return ret
	}
	return *o.OrganizationName
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumePurchasingLocationListView) GetOrganizationNameOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationName) {
		return nil, false
	}
	return o.OrganizationName, true
}

// HasOrganizationName returns a boolean if a field has been set.
func (o *VolumePurchasingLocationListView) HasOrganizationName() bool {
	if o != nil && !IsNil(o.OrganizationName) {
		return true
	}

	return false
}

// SetOrganizationName gets a reference to the given string and assigns it to the OrganizationName field.
func (o *VolumePurchasingLocationListView) SetOrganizationName(v string) {
	o.OrganizationName = &v
}

// GetTokenExpiration returns the TokenExpiration field value if set, zero value otherwise.
func (o *VolumePurchasingLocationListView) GetTokenExpiration() string {
	if o == nil || IsNil(o.TokenExpiration) {
		var ret string
		return ret
	}
	return *o.TokenExpiration
}

// GetTokenExpirationOk returns a tuple with the TokenExpiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumePurchasingLocationListView) GetTokenExpirationOk() (*string, bool) {
	if o == nil || IsNil(o.TokenExpiration) {
		return nil, false
	}
	return o.TokenExpiration, true
}

// HasTokenExpiration returns a boolean if a field has been set.
func (o *VolumePurchasingLocationListView) HasTokenExpiration() bool {
	if o != nil && !IsNil(o.TokenExpiration) {
		return true
	}

	return false
}

// SetTokenExpiration gets a reference to the given string and assigns it to the TokenExpiration field.
func (o *VolumePurchasingLocationListView) SetTokenExpiration(v string) {
	o.TokenExpiration = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *VolumePurchasingLocationListView) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumePurchasingLocationListView) GetCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *VolumePurchasingLocationListView) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *VolumePurchasingLocationListView) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetLocationName returns the LocationName field value if set, zero value otherwise.
func (o *VolumePurchasingLocationListView) GetLocationName() string {
	if o == nil || IsNil(o.LocationName) {
		var ret string
		return ret
	}
	return *o.LocationName
}

// GetLocationNameOk returns a tuple with the LocationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumePurchasingLocationListView) GetLocationNameOk() (*string, bool) {
	if o == nil || IsNil(o.LocationName) {
		return nil, false
	}
	return o.LocationName, true
}

// HasLocationName returns a boolean if a field has been set.
func (o *VolumePurchasingLocationListView) HasLocationName() bool {
	if o != nil && !IsNil(o.LocationName) {
		return true
	}

	return false
}

// SetLocationName gets a reference to the given string and assigns it to the LocationName field.
func (o *VolumePurchasingLocationListView) SetLocationName(v string) {
	o.LocationName = &v
}

// GetClientContextMismatch returns the ClientContextMismatch field value if set, zero value otherwise.
func (o *VolumePurchasingLocationListView) GetClientContextMismatch() bool {
	if o == nil || IsNil(o.ClientContextMismatch) {
		var ret bool
		return ret
	}
	return *o.ClientContextMismatch
}

// GetClientContextMismatchOk returns a tuple with the ClientContextMismatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumePurchasingLocationListView) GetClientContextMismatchOk() (*bool, bool) {
	if o == nil || IsNil(o.ClientContextMismatch) {
		return nil, false
	}
	return o.ClientContextMismatch, true
}

// HasClientContextMismatch returns a boolean if a field has been set.
func (o *VolumePurchasingLocationListView) HasClientContextMismatch() bool {
	if o != nil && !IsNil(o.ClientContextMismatch) {
		return true
	}

	return false
}

// SetClientContextMismatch gets a reference to the given bool and assigns it to the ClientContextMismatch field.
func (o *VolumePurchasingLocationListView) SetClientContextMismatch(v bool) {
	o.ClientContextMismatch = &v
}

// GetAutomaticallyPopulatePurchasedContent returns the AutomaticallyPopulatePurchasedContent field value if set, zero value otherwise.
func (o *VolumePurchasingLocationListView) GetAutomaticallyPopulatePurchasedContent() bool {
	if o == nil || IsNil(o.AutomaticallyPopulatePurchasedContent) {
		var ret bool
		return ret
	}
	return *o.AutomaticallyPopulatePurchasedContent
}

// GetAutomaticallyPopulatePurchasedContentOk returns a tuple with the AutomaticallyPopulatePurchasedContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumePurchasingLocationListView) GetAutomaticallyPopulatePurchasedContentOk() (*bool, bool) {
	if o == nil || IsNil(o.AutomaticallyPopulatePurchasedContent) {
		return nil, false
	}
	return o.AutomaticallyPopulatePurchasedContent, true
}

// HasAutomaticallyPopulatePurchasedContent returns a boolean if a field has been set.
func (o *VolumePurchasingLocationListView) HasAutomaticallyPopulatePurchasedContent() bool {
	if o != nil && !IsNil(o.AutomaticallyPopulatePurchasedContent) {
		return true
	}

	return false
}

// SetAutomaticallyPopulatePurchasedContent gets a reference to the given bool and assigns it to the AutomaticallyPopulatePurchasedContent field.
func (o *VolumePurchasingLocationListView) SetAutomaticallyPopulatePurchasedContent(v bool) {
	o.AutomaticallyPopulatePurchasedContent = &v
}

// GetSendNotificationWhenNoLongerAssigned returns the SendNotificationWhenNoLongerAssigned field value if set, zero value otherwise.
func (o *VolumePurchasingLocationListView) GetSendNotificationWhenNoLongerAssigned() bool {
	if o == nil || IsNil(o.SendNotificationWhenNoLongerAssigned) {
		var ret bool
		return ret
	}
	return *o.SendNotificationWhenNoLongerAssigned
}

// GetSendNotificationWhenNoLongerAssignedOk returns a tuple with the SendNotificationWhenNoLongerAssigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumePurchasingLocationListView) GetSendNotificationWhenNoLongerAssignedOk() (*bool, bool) {
	if o == nil || IsNil(o.SendNotificationWhenNoLongerAssigned) {
		return nil, false
	}
	return o.SendNotificationWhenNoLongerAssigned, true
}

// HasSendNotificationWhenNoLongerAssigned returns a boolean if a field has been set.
func (o *VolumePurchasingLocationListView) HasSendNotificationWhenNoLongerAssigned() bool {
	if o != nil && !IsNil(o.SendNotificationWhenNoLongerAssigned) {
		return true
	}

	return false
}

// SetSendNotificationWhenNoLongerAssigned gets a reference to the given bool and assigns it to the SendNotificationWhenNoLongerAssigned field.
func (o *VolumePurchasingLocationListView) SetSendNotificationWhenNoLongerAssigned(v bool) {
	o.SendNotificationWhenNoLongerAssigned = &v
}

// GetAutoRegisterManagedUsers returns the AutoRegisterManagedUsers field value if set, zero value otherwise.
func (o *VolumePurchasingLocationListView) GetAutoRegisterManagedUsers() bool {
	if o == nil || IsNil(o.AutoRegisterManagedUsers) {
		var ret bool
		return ret
	}
	return *o.AutoRegisterManagedUsers
}

// GetAutoRegisterManagedUsersOk returns a tuple with the AutoRegisterManagedUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumePurchasingLocationListView) GetAutoRegisterManagedUsersOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoRegisterManagedUsers) {
		return nil, false
	}
	return o.AutoRegisterManagedUsers, true
}

// HasAutoRegisterManagedUsers returns a boolean if a field has been set.
func (o *VolumePurchasingLocationListView) HasAutoRegisterManagedUsers() bool {
	if o != nil && !IsNil(o.AutoRegisterManagedUsers) {
		return true
	}

	return false
}

// SetAutoRegisterManagedUsers gets a reference to the given bool and assigns it to the AutoRegisterManagedUsers field.
func (o *VolumePurchasingLocationListView) SetAutoRegisterManagedUsers(v bool) {
	o.AutoRegisterManagedUsers = &v
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *VolumePurchasingLocationListView) GetSiteId() string {
	if o == nil || IsNil(o.SiteId) {
		var ret string
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumePurchasingLocationListView) GetSiteIdOk() (*string, bool) {
	if o == nil || IsNil(o.SiteId) {
		return nil, false
	}
	return o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *VolumePurchasingLocationListView) HasSiteId() bool {
	if o != nil && !IsNil(o.SiteId) {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given string and assigns it to the SiteId field.
func (o *VolumePurchasingLocationListView) SetSiteId(v string) {
	o.SiteId = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *VolumePurchasingLocationListView) GetSiteName() string {
	if o == nil || IsNil(o.SiteName) {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumePurchasingLocationListView) GetSiteNameOk() (*string, bool) {
	if o == nil || IsNil(o.SiteName) {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *VolumePurchasingLocationListView) HasSiteName() bool {
	if o != nil && !IsNil(o.SiteName) {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *VolumePurchasingLocationListView) SetSiteName(v string) {
	o.SiteName = &v
}

// GetLastSyncTime returns the LastSyncTime field value if set, zero value otherwise.
func (o *VolumePurchasingLocationListView) GetLastSyncTime() string {
	if o == nil || IsNil(o.LastSyncTime) {
		var ret string
		return ret
	}
	return *o.LastSyncTime
}

// GetLastSyncTimeOk returns a tuple with the LastSyncTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumePurchasingLocationListView) GetLastSyncTimeOk() (*string, bool) {
	if o == nil || IsNil(o.LastSyncTime) {
		return nil, false
	}
	return o.LastSyncTime, true
}

// HasLastSyncTime returns a boolean if a field has been set.
func (o *VolumePurchasingLocationListView) HasLastSyncTime() bool {
	if o != nil && !IsNil(o.LastSyncTime) {
		return true
	}

	return false
}

// SetLastSyncTime gets a reference to the given string and assigns it to the LastSyncTime field.
func (o *VolumePurchasingLocationListView) SetLastSyncTime(v string) {
	o.LastSyncTime = &v
}

func (o VolumePurchasingLocationListView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VolumePurchasingLocationListView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.TotalPurchasedLicenses) {
		toSerialize["totalPurchasedLicenses"] = o.TotalPurchasedLicenses
	}
	if !IsNil(o.TotalUsedLicenses) {
		toSerialize["totalUsedLicenses"] = o.TotalUsedLicenses
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.AppleId) {
		toSerialize["appleId"] = o.AppleId
	}
	if !IsNil(o.OrganizationName) {
		toSerialize["organizationName"] = o.OrganizationName
	}
	if !IsNil(o.TokenExpiration) {
		toSerialize["tokenExpiration"] = o.TokenExpiration
	}
	if !IsNil(o.CountryCode) {
		toSerialize["countryCode"] = o.CountryCode
	}
	if !IsNil(o.LocationName) {
		toSerialize["locationName"] = o.LocationName
	}
	if !IsNil(o.ClientContextMismatch) {
		toSerialize["clientContextMismatch"] = o.ClientContextMismatch
	}
	if !IsNil(o.AutomaticallyPopulatePurchasedContent) {
		toSerialize["automaticallyPopulatePurchasedContent"] = o.AutomaticallyPopulatePurchasedContent
	}
	if !IsNil(o.SendNotificationWhenNoLongerAssigned) {
		toSerialize["sendNotificationWhenNoLongerAssigned"] = o.SendNotificationWhenNoLongerAssigned
	}
	if !IsNil(o.AutoRegisterManagedUsers) {
		toSerialize["autoRegisterManagedUsers"] = o.AutoRegisterManagedUsers
	}
	if !IsNil(o.SiteId) {
		toSerialize["siteId"] = o.SiteId
	}
	if !IsNil(o.SiteName) {
		toSerialize["siteName"] = o.SiteName
	}
	if !IsNil(o.LastSyncTime) {
		toSerialize["lastSyncTime"] = o.LastSyncTime
	}
	return toSerialize, nil
}

type NullableVolumePurchasingLocationListView struct {
	value *VolumePurchasingLocationListView
	isSet bool
}

func (v NullableVolumePurchasingLocationListView) Get() *VolumePurchasingLocationListView {
	return v.value
}

func (v *NullableVolumePurchasingLocationListView) Set(val *VolumePurchasingLocationListView) {
	v.value = val
	v.isSet = true
}

func (v NullableVolumePurchasingLocationListView) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumePurchasingLocationListView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumePurchasingLocationListView(val *VolumePurchasingLocationListView) *NullableVolumePurchasingLocationListView {
	return &NullableVolumePurchasingLocationListView{value: val, isSet: true}
}

func (v NullableVolumePurchasingLocationListView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumePurchasingLocationListView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


