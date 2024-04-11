/*
Jamf Pro API

## Overview The Jamf Pro API is a RESTful API for Jamf Pro built to enable consistent and efficient programmatic access to Jamf Pro.<br/><br/> The swagger schema can be found [here](/api/schema/). 

API version: production
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the VolumePurchasingLocationPost type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VolumePurchasingLocationPost{}

// VolumePurchasingLocationPost struct for VolumePurchasingLocationPost
type VolumePurchasingLocationPost struct {
	// If no value is provided when creating a VolumePurchasingLocation object, the 'name' will default to the 'locationName' value
	Name *string `json:"name,omitempty"`
	AutomaticallyPopulatePurchasedContent *bool `json:"automaticallyPopulatePurchasedContent,omitempty"`
	SendNotificationWhenNoLongerAssigned *bool `json:"sendNotificationWhenNoLongerAssigned,omitempty"`
	AutoRegisterManagedUsers *bool `json:"autoRegisterManagedUsers,omitempty"`
	SiteId *string `json:"siteId,omitempty"`
	ServiceToken string `json:"serviceToken"`
}

type _VolumePurchasingLocationPost VolumePurchasingLocationPost

// NewVolumePurchasingLocationPost instantiates a new VolumePurchasingLocationPost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolumePurchasingLocationPost(serviceToken string) *VolumePurchasingLocationPost {
	this := VolumePurchasingLocationPost{}
	var automaticallyPopulatePurchasedContent bool = false
	this.AutomaticallyPopulatePurchasedContent = &automaticallyPopulatePurchasedContent
	var sendNotificationWhenNoLongerAssigned bool = false
	this.SendNotificationWhenNoLongerAssigned = &sendNotificationWhenNoLongerAssigned
	var autoRegisterManagedUsers bool = false
	this.AutoRegisterManagedUsers = &autoRegisterManagedUsers
	var siteId string = "-1"
	this.SiteId = &siteId
	this.ServiceToken = serviceToken
	return &this
}

// NewVolumePurchasingLocationPostWithDefaults instantiates a new VolumePurchasingLocationPost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumePurchasingLocationPostWithDefaults() *VolumePurchasingLocationPost {
	this := VolumePurchasingLocationPost{}
	var automaticallyPopulatePurchasedContent bool = false
	this.AutomaticallyPopulatePurchasedContent = &automaticallyPopulatePurchasedContent
	var sendNotificationWhenNoLongerAssigned bool = false
	this.SendNotificationWhenNoLongerAssigned = &sendNotificationWhenNoLongerAssigned
	var autoRegisterManagedUsers bool = false
	this.AutoRegisterManagedUsers = &autoRegisterManagedUsers
	var siteId string = "-1"
	this.SiteId = &siteId
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VolumePurchasingLocationPost) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumePurchasingLocationPost) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VolumePurchasingLocationPost) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VolumePurchasingLocationPost) SetName(v string) {
	o.Name = &v
}

// GetAutomaticallyPopulatePurchasedContent returns the AutomaticallyPopulatePurchasedContent field value if set, zero value otherwise.
func (o *VolumePurchasingLocationPost) GetAutomaticallyPopulatePurchasedContent() bool {
	if o == nil || IsNil(o.AutomaticallyPopulatePurchasedContent) {
		var ret bool
		return ret
	}
	return *o.AutomaticallyPopulatePurchasedContent
}

// GetAutomaticallyPopulatePurchasedContentOk returns a tuple with the AutomaticallyPopulatePurchasedContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumePurchasingLocationPost) GetAutomaticallyPopulatePurchasedContentOk() (*bool, bool) {
	if o == nil || IsNil(o.AutomaticallyPopulatePurchasedContent) {
		return nil, false
	}
	return o.AutomaticallyPopulatePurchasedContent, true
}

// HasAutomaticallyPopulatePurchasedContent returns a boolean if a field has been set.
func (o *VolumePurchasingLocationPost) HasAutomaticallyPopulatePurchasedContent() bool {
	if o != nil && !IsNil(o.AutomaticallyPopulatePurchasedContent) {
		return true
	}

	return false
}

// SetAutomaticallyPopulatePurchasedContent gets a reference to the given bool and assigns it to the AutomaticallyPopulatePurchasedContent field.
func (o *VolumePurchasingLocationPost) SetAutomaticallyPopulatePurchasedContent(v bool) {
	o.AutomaticallyPopulatePurchasedContent = &v
}

// GetSendNotificationWhenNoLongerAssigned returns the SendNotificationWhenNoLongerAssigned field value if set, zero value otherwise.
func (o *VolumePurchasingLocationPost) GetSendNotificationWhenNoLongerAssigned() bool {
	if o == nil || IsNil(o.SendNotificationWhenNoLongerAssigned) {
		var ret bool
		return ret
	}
	return *o.SendNotificationWhenNoLongerAssigned
}

// GetSendNotificationWhenNoLongerAssignedOk returns a tuple with the SendNotificationWhenNoLongerAssigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumePurchasingLocationPost) GetSendNotificationWhenNoLongerAssignedOk() (*bool, bool) {
	if o == nil || IsNil(o.SendNotificationWhenNoLongerAssigned) {
		return nil, false
	}
	return o.SendNotificationWhenNoLongerAssigned, true
}

// HasSendNotificationWhenNoLongerAssigned returns a boolean if a field has been set.
func (o *VolumePurchasingLocationPost) HasSendNotificationWhenNoLongerAssigned() bool {
	if o != nil && !IsNil(o.SendNotificationWhenNoLongerAssigned) {
		return true
	}

	return false
}

// SetSendNotificationWhenNoLongerAssigned gets a reference to the given bool and assigns it to the SendNotificationWhenNoLongerAssigned field.
func (o *VolumePurchasingLocationPost) SetSendNotificationWhenNoLongerAssigned(v bool) {
	o.SendNotificationWhenNoLongerAssigned = &v
}

// GetAutoRegisterManagedUsers returns the AutoRegisterManagedUsers field value if set, zero value otherwise.
func (o *VolumePurchasingLocationPost) GetAutoRegisterManagedUsers() bool {
	if o == nil || IsNil(o.AutoRegisterManagedUsers) {
		var ret bool
		return ret
	}
	return *o.AutoRegisterManagedUsers
}

// GetAutoRegisterManagedUsersOk returns a tuple with the AutoRegisterManagedUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumePurchasingLocationPost) GetAutoRegisterManagedUsersOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoRegisterManagedUsers) {
		return nil, false
	}
	return o.AutoRegisterManagedUsers, true
}

// HasAutoRegisterManagedUsers returns a boolean if a field has been set.
func (o *VolumePurchasingLocationPost) HasAutoRegisterManagedUsers() bool {
	if o != nil && !IsNil(o.AutoRegisterManagedUsers) {
		return true
	}

	return false
}

// SetAutoRegisterManagedUsers gets a reference to the given bool and assigns it to the AutoRegisterManagedUsers field.
func (o *VolumePurchasingLocationPost) SetAutoRegisterManagedUsers(v bool) {
	o.AutoRegisterManagedUsers = &v
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *VolumePurchasingLocationPost) GetSiteId() string {
	if o == nil || IsNil(o.SiteId) {
		var ret string
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumePurchasingLocationPost) GetSiteIdOk() (*string, bool) {
	if o == nil || IsNil(o.SiteId) {
		return nil, false
	}
	return o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *VolumePurchasingLocationPost) HasSiteId() bool {
	if o != nil && !IsNil(o.SiteId) {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given string and assigns it to the SiteId field.
func (o *VolumePurchasingLocationPost) SetSiteId(v string) {
	o.SiteId = &v
}

// GetServiceToken returns the ServiceToken field value
func (o *VolumePurchasingLocationPost) GetServiceToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceToken
}

// GetServiceTokenOk returns a tuple with the ServiceToken field value
// and a boolean to check if the value has been set.
func (o *VolumePurchasingLocationPost) GetServiceTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceToken, true
}

// SetServiceToken sets field value
func (o *VolumePurchasingLocationPost) SetServiceToken(v string) {
	o.ServiceToken = v
}

func (o VolumePurchasingLocationPost) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VolumePurchasingLocationPost) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
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
	toSerialize["serviceToken"] = o.ServiceToken
	return toSerialize, nil
}

func (o *VolumePurchasingLocationPost) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"serviceToken",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varVolumePurchasingLocationPost := _VolumePurchasingLocationPost{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVolumePurchasingLocationPost)

	if err != nil {
		return err
	}

	*o = VolumePurchasingLocationPost(varVolumePurchasingLocationPost)

	return err
}

type NullableVolumePurchasingLocationPost struct {
	value *VolumePurchasingLocationPost
	isSet bool
}

func (v NullableVolumePurchasingLocationPost) Get() *VolumePurchasingLocationPost {
	return v.value
}

func (v *NullableVolumePurchasingLocationPost) Set(val *VolumePurchasingLocationPost) {
	v.value = val
	v.isSet = true
}

func (v NullableVolumePurchasingLocationPost) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumePurchasingLocationPost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumePurchasingLocationPost(val *VolumePurchasingLocationPost) *NullableVolumePurchasingLocationPost {
	return &NullableVolumePurchasingLocationPost{value: val, isSet: true}
}

func (v NullableVolumePurchasingLocationPost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumePurchasingLocationPost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


