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

// Ebook struct for Ebook
type Ebook struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Kind *string `json:"kind,omitempty"`
	Url *string `json:"url,omitempty"`
	Free *bool `json:"free,omitempty"`
	Version *string `json:"version,omitempty"`
	Author *string `json:"author,omitempty"`
	// If true, it will be automatically installed
	DeployAsManaged *bool `json:"deployAsManaged,omitempty"`
	InstallAutomatically *bool `json:"installAutomatically,omitempty"`
	CategoryId *string `json:"categoryId,omitempty"`
	SiteId *string `json:"siteId,omitempty"`
}

// NewEbook instantiates a new Ebook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEbook() *Ebook {
	this := Ebook{}
	return &this
}

// NewEbookWithDefaults instantiates a new Ebook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEbookWithDefaults() *Ebook {
	this := Ebook{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Ebook) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ebook) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Ebook) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Ebook) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Ebook) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ebook) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Ebook) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Ebook) SetName(v string) {
	o.Name = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *Ebook) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ebook) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *Ebook) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *Ebook) SetKind(v string) {
	o.Kind = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *Ebook) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ebook) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *Ebook) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *Ebook) SetUrl(v string) {
	o.Url = &v
}

// GetFree returns the Free field value if set, zero value otherwise.
func (o *Ebook) GetFree() bool {
	if o == nil || o.Free == nil {
		var ret bool
		return ret
	}
	return *o.Free
}

// GetFreeOk returns a tuple with the Free field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ebook) GetFreeOk() (*bool, bool) {
	if o == nil || o.Free == nil {
		return nil, false
	}
	return o.Free, true
}

// HasFree returns a boolean if a field has been set.
func (o *Ebook) HasFree() bool {
	if o != nil && o.Free != nil {
		return true
	}

	return false
}

// SetFree gets a reference to the given bool and assigns it to the Free field.
func (o *Ebook) SetFree(v bool) {
	o.Free = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Ebook) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ebook) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Ebook) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *Ebook) SetVersion(v string) {
	o.Version = &v
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *Ebook) GetAuthor() string {
	if o == nil || o.Author == nil {
		var ret string
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ebook) GetAuthorOk() (*string, bool) {
	if o == nil || o.Author == nil {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *Ebook) HasAuthor() bool {
	if o != nil && o.Author != nil {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given string and assigns it to the Author field.
func (o *Ebook) SetAuthor(v string) {
	o.Author = &v
}

// GetDeployAsManaged returns the DeployAsManaged field value if set, zero value otherwise.
func (o *Ebook) GetDeployAsManaged() bool {
	if o == nil || o.DeployAsManaged == nil {
		var ret bool
		return ret
	}
	return *o.DeployAsManaged
}

// GetDeployAsManagedOk returns a tuple with the DeployAsManaged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ebook) GetDeployAsManagedOk() (*bool, bool) {
	if o == nil || o.DeployAsManaged == nil {
		return nil, false
	}
	return o.DeployAsManaged, true
}

// HasDeployAsManaged returns a boolean if a field has been set.
func (o *Ebook) HasDeployAsManaged() bool {
	if o != nil && o.DeployAsManaged != nil {
		return true
	}

	return false
}

// SetDeployAsManaged gets a reference to the given bool and assigns it to the DeployAsManaged field.
func (o *Ebook) SetDeployAsManaged(v bool) {
	o.DeployAsManaged = &v
}

// GetInstallAutomatically returns the InstallAutomatically field value if set, zero value otherwise.
func (o *Ebook) GetInstallAutomatically() bool {
	if o == nil || o.InstallAutomatically == nil {
		var ret bool
		return ret
	}
	return *o.InstallAutomatically
}

// GetInstallAutomaticallyOk returns a tuple with the InstallAutomatically field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ebook) GetInstallAutomaticallyOk() (*bool, bool) {
	if o == nil || o.InstallAutomatically == nil {
		return nil, false
	}
	return o.InstallAutomatically, true
}

// HasInstallAutomatically returns a boolean if a field has been set.
func (o *Ebook) HasInstallAutomatically() bool {
	if o != nil && o.InstallAutomatically != nil {
		return true
	}

	return false
}

// SetInstallAutomatically gets a reference to the given bool and assigns it to the InstallAutomatically field.
func (o *Ebook) SetInstallAutomatically(v bool) {
	o.InstallAutomatically = &v
}

// GetCategoryId returns the CategoryId field value if set, zero value otherwise.
func (o *Ebook) GetCategoryId() string {
	if o == nil || o.CategoryId == nil {
		var ret string
		return ret
	}
	return *o.CategoryId
}

// GetCategoryIdOk returns a tuple with the CategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ebook) GetCategoryIdOk() (*string, bool) {
	if o == nil || o.CategoryId == nil {
		return nil, false
	}
	return o.CategoryId, true
}

// HasCategoryId returns a boolean if a field has been set.
func (o *Ebook) HasCategoryId() bool {
	if o != nil && o.CategoryId != nil {
		return true
	}

	return false
}

// SetCategoryId gets a reference to the given string and assigns it to the CategoryId field.
func (o *Ebook) SetCategoryId(v string) {
	o.CategoryId = &v
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *Ebook) GetSiteId() string {
	if o == nil || o.SiteId == nil {
		var ret string
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ebook) GetSiteIdOk() (*string, bool) {
	if o == nil || o.SiteId == nil {
		return nil, false
	}
	return o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *Ebook) HasSiteId() bool {
	if o != nil && o.SiteId != nil {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given string and assigns it to the SiteId field.
func (o *Ebook) SetSiteId(v string) {
	o.SiteId = &v
}

func (o Ebook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.Free != nil {
		toSerialize["free"] = o.Free
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Author != nil {
		toSerialize["author"] = o.Author
	}
	if o.DeployAsManaged != nil {
		toSerialize["deployAsManaged"] = o.DeployAsManaged
	}
	if o.InstallAutomatically != nil {
		toSerialize["installAutomatically"] = o.InstallAutomatically
	}
	if o.CategoryId != nil {
		toSerialize["categoryId"] = o.CategoryId
	}
	if o.SiteId != nil {
		toSerialize["siteId"] = o.SiteId
	}
	return json.Marshal(toSerialize)
}

type NullableEbook struct {
	value *Ebook
	isSet bool
}

func (v NullableEbook) Get() *Ebook {
	return v.value
}

func (v *NullableEbook) Set(val *Ebook) {
	v.value = val
	v.isSet = true
}

func (v NullableEbook) IsSet() bool {
	return v.isSet
}

func (v *NullableEbook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEbook(val *Ebook) *NullableEbook {
	return &NullableEbook{value: val, isSet: true}
}

func (v NullableEbook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEbook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


