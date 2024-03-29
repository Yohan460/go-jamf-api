/*
Jamf Pro API

## Overview The Jamf Pro API is a RESTful API for Jamf Pro built to enable consistent and efficient programmatic access to Jamf Pro.<br/><br/> The swagger schema can be found [here](/api/schema/). 

API version: production
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// checks if the IosDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IosDetails{}

// IosDetails will be populated if the type is ios.
type IosDetails struct {
	Model *string `json:"model,omitempty"`
	ModelIdentifier *string `json:"modelIdentifier,omitempty"`
	ModelNumber *string `json:"modelNumber,omitempty"`
	IsSupervised *bool `json:"isSupervised,omitempty"`
	BatteryLevel *int32 `json:"batteryLevel,omitempty"`
	LastBackupTimestamp *time.Time `json:"lastBackupTimestamp,omitempty"`
	CapacityMb *int32 `json:"capacityMb,omitempty"`
	AvailableMb *int32 `json:"availableMb,omitempty"`
	PercentageUsed *int32 `json:"percentageUsed,omitempty"`
	IsShared *bool `json:"isShared,omitempty"`
	IsDeviceLocatorServiceEnabled *bool `json:"isDeviceLocatorServiceEnabled,omitempty"`
	IsDoNotDisturbEnabled *bool `json:"isDoNotDisturbEnabled,omitempty"`
	IsCloudBackupEnabled *bool `json:"isCloudBackupEnabled,omitempty"`
	LastCloudBackupTimestamp *time.Time `json:"lastCloudBackupTimestamp,omitempty"`
	IsLocationServicesEnabled *bool `json:"isLocationServicesEnabled,omitempty"`
	IsITunesStoreAccountActive *bool `json:"isITunesStoreAccountActive,omitempty"`
	IsBleCapable *bool `json:"isBleCapable,omitempty"`
	Computer *IdAndName `json:"computer,omitempty"`
	Purchasing *Purchasing `json:"purchasing,omitempty"`
	Security *Security `json:"security,omitempty"`
	Network *Network `json:"network,omitempty"`
	Applications []MobileDeviceApplication `json:"applications,omitempty"`
	Certificates []MobileDeviceCertificateV1 `json:"certificates,omitempty"`
	Ebooks []MobileDeviceEbook `json:"ebooks,omitempty"`
	ConfigurationProfiles []ConfigurationProfile `json:"configurationProfiles,omitempty"`
	ProvisioningProfiles []MobileDeviceProvisioningProfiles `json:"provisioningProfiles,omitempty"`
	Attachments []MobileDeviceAttachment `json:"attachments,omitempty"`
}

// NewIosDetails instantiates a new IosDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIosDetails() *IosDetails {
	this := IosDetails{}
	return &this
}

// NewIosDetailsWithDefaults instantiates a new IosDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIosDetailsWithDefaults() *IosDetails {
	this := IosDetails{}
	return &this
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *IosDetails) GetModel() string {
	if o == nil || IsNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IosDetails) GetModelOk() (*string, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *IosDetails) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *IosDetails) SetModel(v string) {
	o.Model = &v
}

// GetModelIdentifier returns the ModelIdentifier field value if set, zero value otherwise.
func (o *IosDetails) GetModelIdentifier() string {
	if o == nil || IsNil(o.ModelIdentifier) {
		var ret string
		return ret
	}
	return *o.ModelIdentifier
}

// GetModelIdentifierOk returns a tuple with the ModelIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IosDetails) GetModelIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.ModelIdentifier) {
		return nil, false
	}
	return o.ModelIdentifier, true
}

// HasModelIdentifier returns a boolean if a field has been set.
func (o *IosDetails) HasModelIdentifier() bool {
	if o != nil && !IsNil(o.ModelIdentifier) {
		return true
	}

	return false
}

// SetModelIdentifier gets a reference to the given string and assigns it to the ModelIdentifier field.
func (o *IosDetails) SetModelIdentifier(v string) {
	o.ModelIdentifier = &v
}

// GetModelNumber returns the ModelNumber field value if set, zero value otherwise.
func (o *IosDetails) GetModelNumber() string {
	if o == nil || IsNil(o.ModelNumber) {
		var ret string
		return ret
	}
	return *o.ModelNumber
}

// GetModelNumberOk returns a tuple with the ModelNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IosDetails) GetModelNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ModelNumber) {
		return nil, false
	}
	return o.ModelNumber, true
}

// HasModelNumber returns a boolean if a field has been set.
func (o *IosDetails) HasModelNumber() bool {
	if o != nil && !IsNil(o.ModelNumber) {
		return true
	}

	return false
}

// SetModelNumber gets a reference to the given string and assigns it to the ModelNumber field.
func (o *IosDetails) SetModelNumber(v string) {
	o.ModelNumber = &v
}

// GetIsSupervised returns the IsSupervised field value if set, zero value otherwise.
func (o *IosDetails) GetIsSupervised() bool {
	if o == nil || IsNil(o.IsSupervised) {
		var ret bool
		return ret
	}
	return *o.IsSupervised
}

// GetIsSupervisedOk returns a tuple with the IsSupervised field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IosDetails) GetIsSupervisedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSupervised) {
		return nil, false
	}
	return o.IsSupervised, true
}

// HasIsSupervised returns a boolean if a field has been set.
func (o *IosDetails) HasIsSupervised() bool {
	if o != nil && !IsNil(o.IsSupervised) {
		return true
	}

	return false
}

// SetIsSupervised gets a reference to the given bool and assigns it to the IsSupervised field.
func (o *IosDetails) SetIsSupervised(v bool) {
	o.IsSupervised = &v
}

// GetBatteryLevel returns the BatteryLevel field value if set, zero value otherwise.
func (o *IosDetails) GetBatteryLevel() int32 {
	if o == nil || IsNil(o.BatteryLevel) {
		var ret int32
		return ret
	}
	return *o.BatteryLevel
}

// GetBatteryLevelOk returns a tuple with the BatteryLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IosDetails) GetBatteryLevelOk() (*int32, bool) {
	if o == nil || IsNil(o.BatteryLevel) {
		return nil, false
	}
	return o.BatteryLevel, true
}

// HasBatteryLevel returns a boolean if a field has been set.
func (o *IosDetails) HasBatteryLevel() bool {
	if o != nil && !IsNil(o.BatteryLevel) {
		return true
	}

	return false
}

// SetBatteryLevel gets a reference to the given int32 and assigns it to the BatteryLevel field.
func (o *IosDetails) SetBatteryLevel(v int32) {
	o.BatteryLevel = &v
}

// GetLastBackupTimestamp returns the LastBackupTimestamp field value if set, zero value otherwise.
func (o *IosDetails) GetLastBackupTimestamp() time.Time {
	if o == nil || IsNil(o.LastBackupTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.LastBackupTimestamp
}

// GetLastBackupTimestampOk returns a tuple with the LastBackupTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IosDetails) GetLastBackupTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastBackupTimestamp) {
		return nil, false
	}
	return o.LastBackupTimestamp, true
}

// HasLastBackupTimestamp returns a boolean if a field has been set.
func (o *IosDetails) HasLastBackupTimestamp() bool {
	if o != nil && !IsNil(o.LastBackupTimestamp) {
		return true
	}

	return false
}

// SetLastBackupTimestamp gets a reference to the given time.Time and assigns it to the LastBackupTimestamp field.
func (o *IosDetails) SetLastBackupTimestamp(v time.Time) {
	o.LastBackupTimestamp = &v
}

// GetCapacityMb returns the CapacityMb field value if set, zero value otherwise.
func (o *IosDetails) GetCapacityMb() int32 {
	if o == nil || IsNil(o.CapacityMb) {
		var ret int32
		return ret
	}
	return *o.CapacityMb
}

// GetCapacityMbOk returns a tuple with the CapacityMb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IosDetails) GetCapacityMbOk() (*int32, bool) {
	if o == nil || IsNil(o.CapacityMb) {
		return nil, false
	}
	return o.CapacityMb, true
}

// HasCapacityMb returns a boolean if a field has been set.
func (o *IosDetails) HasCapacityMb() bool {
	if o != nil && !IsNil(o.CapacityMb) {
		return true
	}

	return false
}

// SetCapacityMb gets a reference to the given int32 and assigns it to the CapacityMb field.
func (o *IosDetails) SetCapacityMb(v int32) {
	o.CapacityMb = &v
}

// GetAvailableMb returns the AvailableMb field value if set, zero value otherwise.
func (o *IosDetails) GetAvailableMb() int32 {
	if o == nil || IsNil(o.AvailableMb) {
		var ret int32
		return ret
	}
	return *o.AvailableMb
}

// GetAvailableMbOk returns a tuple with the AvailableMb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IosDetails) GetAvailableMbOk() (*int32, bool) {
	if o == nil || IsNil(o.AvailableMb) {
		return nil, false
	}
	return o.AvailableMb, true
}

// HasAvailableMb returns a boolean if a field has been set.
func (o *IosDetails) HasAvailableMb() bool {
	if o != nil && !IsNil(o.AvailableMb) {
		return true
	}

	return false
}

// SetAvailableMb gets a reference to the given int32 and assigns it to the AvailableMb field.
func (o *IosDetails) SetAvailableMb(v int32) {
	o.AvailableMb = &v
}

// GetPercentageUsed returns the PercentageUsed field value if set, zero value otherwise.
func (o *IosDetails) GetPercentageUsed() int32 {
	if o == nil || IsNil(o.PercentageUsed) {
		var ret int32
		return ret
	}
	return *o.PercentageUsed
}

// GetPercentageUsedOk returns a tuple with the PercentageUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IosDetails) GetPercentageUsedOk() (*int32, bool) {
	if o == nil || IsNil(o.PercentageUsed) {
		return nil, false
	}
	return o.PercentageUsed, true
}

// HasPercentageUsed returns a boolean if a field has been set.
func (o *IosDetails) HasPercentageUsed() bool {
	if o != nil && !IsNil(o.PercentageUsed) {
		return true
	}

	return false
}

// SetPercentageUsed gets a reference to the given int32 and assigns it to the PercentageUsed field.
func (o *IosDetails) SetPercentageUsed(v int32) {
	o.PercentageUsed = &v
}

// GetIsShared returns the IsShared field value if set, zero value otherwise.
func (o *IosDetails) GetIsShared() bool {
	if o == nil || IsNil(o.IsShared) {
		var ret bool
		return ret
	}
	return *o.IsShared
}

// GetIsSharedOk returns a tuple with the IsShared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IosDetails) GetIsSharedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsShared) {
		return nil, false
	}
	return o.IsShared, true
}

// HasIsShared returns a boolean if a field has been set.
func (o *IosDetails) HasIsShared() bool {
	if o != nil && !IsNil(o.IsShared) {
		return true
	}

	return false
}

// SetIsShared gets a reference to the given bool and assigns it to the IsShared field.
func (o *IosDetails) SetIsShared(v bool) {
	o.IsShared = &v
}

// GetIsDeviceLocatorServiceEnabled returns the IsDeviceLocatorServiceEnabled field value if set, zero value otherwise.
func (o *IosDetails) GetIsDeviceLocatorServiceEnabled() bool {
	if o == nil || IsNil(o.IsDeviceLocatorServiceEnabled) {
		var ret bool
		return ret
	}
	return *o.IsDeviceLocatorServiceEnabled
}

// GetIsDeviceLocatorServiceEnabledOk returns a tuple with the IsDeviceLocatorServiceEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IosDetails) GetIsDeviceLocatorServiceEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDeviceLocatorServiceEnabled) {
		return nil, false
	}
	return o.IsDeviceLocatorServiceEnabled, true
}

// HasIsDeviceLocatorServiceEnabled returns a boolean if a field has been set.
func (o *IosDetails) HasIsDeviceLocatorServiceEnabled() bool {
	if o != nil && !IsNil(o.IsDeviceLocatorServiceEnabled) {
		return true
	}

	return false
}

// SetIsDeviceLocatorServiceEnabled gets a reference to the given bool and assigns it to the IsDeviceLocatorServiceEnabled field.
func (o *IosDetails) SetIsDeviceLocatorServiceEnabled(v bool) {
	o.IsDeviceLocatorServiceEnabled = &v
}

// GetIsDoNotDisturbEnabled returns the IsDoNotDisturbEnabled field value if set, zero value otherwise.
func (o *IosDetails) GetIsDoNotDisturbEnabled() bool {
	if o == nil || IsNil(o.IsDoNotDisturbEnabled) {
		var ret bool
		return ret
	}
	return *o.IsDoNotDisturbEnabled
}

// GetIsDoNotDisturbEnabledOk returns a tuple with the IsDoNotDisturbEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IosDetails) GetIsDoNotDisturbEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDoNotDisturbEnabled) {
		return nil, false
	}
	return o.IsDoNotDisturbEnabled, true
}

// HasIsDoNotDisturbEnabled returns a boolean if a field has been set.
func (o *IosDetails) HasIsDoNotDisturbEnabled() bool {
	if o != nil && !IsNil(o.IsDoNotDisturbEnabled) {
		return true
	}

	return false
}

// SetIsDoNotDisturbEnabled gets a reference to the given bool and assigns it to the IsDoNotDisturbEnabled field.
func (o *IosDetails) SetIsDoNotDisturbEnabled(v bool) {
	o.IsDoNotDisturbEnabled = &v
}

// GetIsCloudBackupEnabled returns the IsCloudBackupEnabled field value if set, zero value otherwise.
func (o *IosDetails) GetIsCloudBackupEnabled() bool {
	if o == nil || IsNil(o.IsCloudBackupEnabled) {
		var ret bool
		return ret
	}
	return *o.IsCloudBackupEnabled
}

// GetIsCloudBackupEnabledOk returns a tuple with the IsCloudBackupEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IosDetails) GetIsCloudBackupEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCloudBackupEnabled) {
		return nil, false
	}
	return o.IsCloudBackupEnabled, true
}

// HasIsCloudBackupEnabled returns a boolean if a field has been set.
func (o *IosDetails) HasIsCloudBackupEnabled() bool {
	if o != nil && !IsNil(o.IsCloudBackupEnabled) {
		return true
	}

	return false
}

// SetIsCloudBackupEnabled gets a reference to the given bool and assigns it to the IsCloudBackupEnabled field.
func (o *IosDetails) SetIsCloudBackupEnabled(v bool) {
	o.IsCloudBackupEnabled = &v
}

// GetLastCloudBackupTimestamp returns the LastCloudBackupTimestamp field value if set, zero value otherwise.
func (o *IosDetails) GetLastCloudBackupTimestamp() time.Time {
	if o == nil || IsNil(o.LastCloudBackupTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.LastCloudBackupTimestamp
}

// GetLastCloudBackupTimestampOk returns a tuple with the LastCloudBackupTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IosDetails) GetLastCloudBackupTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastCloudBackupTimestamp) {
		return nil, false
	}
	return o.LastCloudBackupTimestamp, true
}

// HasLastCloudBackupTimestamp returns a boolean if a field has been set.
func (o *IosDetails) HasLastCloudBackupTimestamp() bool {
	if o != nil && !IsNil(o.LastCloudBackupTimestamp) {
		return true
	}

	return false
}

// SetLastCloudBackupTimestamp gets a reference to the given time.Time and assigns it to the LastCloudBackupTimestamp field.
func (o *IosDetails) SetLastCloudBackupTimestamp(v time.Time) {
	o.LastCloudBackupTimestamp = &v
}

// GetIsLocationServicesEnabled returns the IsLocationServicesEnabled field value if set, zero value otherwise.
func (o *IosDetails) GetIsLocationServicesEnabled() bool {
	if o == nil || IsNil(o.IsLocationServicesEnabled) {
		var ret bool
		return ret
	}
	return *o.IsLocationServicesEnabled
}

// GetIsLocationServicesEnabledOk returns a tuple with the IsLocationServicesEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IosDetails) GetIsLocationServicesEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLocationServicesEnabled) {
		return nil, false
	}
	return o.IsLocationServicesEnabled, true
}

// HasIsLocationServicesEnabled returns a boolean if a field has been set.
func (o *IosDetails) HasIsLocationServicesEnabled() bool {
	if o != nil && !IsNil(o.IsLocationServicesEnabled) {
		return true
	}

	return false
}

// SetIsLocationServicesEnabled gets a reference to the given bool and assigns it to the IsLocationServicesEnabled field.
func (o *IosDetails) SetIsLocationServicesEnabled(v bool) {
	o.IsLocationServicesEnabled = &v
}

// GetIsITunesStoreAccountActive returns the IsITunesStoreAccountActive field value if set, zero value otherwise.
func (o *IosDetails) GetIsITunesStoreAccountActive() bool {
	if o == nil || IsNil(o.IsITunesStoreAccountActive) {
		var ret bool
		return ret
	}
	return *o.IsITunesStoreAccountActive
}

// GetIsITunesStoreAccountActiveOk returns a tuple with the IsITunesStoreAccountActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IosDetails) GetIsITunesStoreAccountActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsITunesStoreAccountActive) {
		return nil, false
	}
	return o.IsITunesStoreAccountActive, true
}

// HasIsITunesStoreAccountActive returns a boolean if a field has been set.
func (o *IosDetails) HasIsITunesStoreAccountActive() bool {
	if o != nil && !IsNil(o.IsITunesStoreAccountActive) {
		return true
	}

	return false
}

// SetIsITunesStoreAccountActive gets a reference to the given bool and assigns it to the IsITunesStoreAccountActive field.
func (o *IosDetails) SetIsITunesStoreAccountActive(v bool) {
	o.IsITunesStoreAccountActive = &v
}

// GetIsBleCapable returns the IsBleCapable field value if set, zero value otherwise.
func (o *IosDetails) GetIsBleCapable() bool {
	if o == nil || IsNil(o.IsBleCapable) {
		var ret bool
		return ret
	}
	return *o.IsBleCapable
}

// GetIsBleCapableOk returns a tuple with the IsBleCapable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IosDetails) GetIsBleCapableOk() (*bool, bool) {
	if o == nil || IsNil(o.IsBleCapable) {
		return nil, false
	}
	return o.IsBleCapable, true
}

// HasIsBleCapable returns a boolean if a field has been set.
func (o *IosDetails) HasIsBleCapable() bool {
	if o != nil && !IsNil(o.IsBleCapable) {
		return true
	}

	return false
}

// SetIsBleCapable gets a reference to the given bool and assigns it to the IsBleCapable field.
func (o *IosDetails) SetIsBleCapable(v bool) {
	o.IsBleCapable = &v
}

// GetComputer returns the Computer field value if set, zero value otherwise.
func (o *IosDetails) GetComputer() IdAndName {
	if o == nil || IsNil(o.Computer) {
		var ret IdAndName
		return ret
	}
	return *o.Computer
}

// GetComputerOk returns a tuple with the Computer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IosDetails) GetComputerOk() (*IdAndName, bool) {
	if o == nil || IsNil(o.Computer) {
		return nil, false
	}
	return o.Computer, true
}

// HasComputer returns a boolean if a field has been set.
func (o *IosDetails) HasComputer() bool {
	if o != nil && !IsNil(o.Computer) {
		return true
	}

	return false
}

// SetComputer gets a reference to the given IdAndName and assigns it to the Computer field.
func (o *IosDetails) SetComputer(v IdAndName) {
	o.Computer = &v
}

// GetPurchasing returns the Purchasing field value if set, zero value otherwise.
func (o *IosDetails) GetPurchasing() Purchasing {
	if o == nil || IsNil(o.Purchasing) {
		var ret Purchasing
		return ret
	}
	return *o.Purchasing
}

// GetPurchasingOk returns a tuple with the Purchasing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IosDetails) GetPurchasingOk() (*Purchasing, bool) {
	if o == nil || IsNil(o.Purchasing) {
		return nil, false
	}
	return o.Purchasing, true
}

// HasPurchasing returns a boolean if a field has been set.
func (o *IosDetails) HasPurchasing() bool {
	if o != nil && !IsNil(o.Purchasing) {
		return true
	}

	return false
}

// SetPurchasing gets a reference to the given Purchasing and assigns it to the Purchasing field.
func (o *IosDetails) SetPurchasing(v Purchasing) {
	o.Purchasing = &v
}

// GetSecurity returns the Security field value if set, zero value otherwise.
func (o *IosDetails) GetSecurity() Security {
	if o == nil || IsNil(o.Security) {
		var ret Security
		return ret
	}
	return *o.Security
}

// GetSecurityOk returns a tuple with the Security field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IosDetails) GetSecurityOk() (*Security, bool) {
	if o == nil || IsNil(o.Security) {
		return nil, false
	}
	return o.Security, true
}

// HasSecurity returns a boolean if a field has been set.
func (o *IosDetails) HasSecurity() bool {
	if o != nil && !IsNil(o.Security) {
		return true
	}

	return false
}

// SetSecurity gets a reference to the given Security and assigns it to the Security field.
func (o *IosDetails) SetSecurity(v Security) {
	o.Security = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *IosDetails) GetNetwork() Network {
	if o == nil || IsNil(o.Network) {
		var ret Network
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IosDetails) GetNetworkOk() (*Network, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *IosDetails) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given Network and assigns it to the Network field.
func (o *IosDetails) SetNetwork(v Network) {
	o.Network = &v
}

// GetApplications returns the Applications field value if set, zero value otherwise.
func (o *IosDetails) GetApplications() []MobileDeviceApplication {
	if o == nil || IsNil(o.Applications) {
		var ret []MobileDeviceApplication
		return ret
	}
	return o.Applications
}

// GetApplicationsOk returns a tuple with the Applications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IosDetails) GetApplicationsOk() ([]MobileDeviceApplication, bool) {
	if o == nil || IsNil(o.Applications) {
		return nil, false
	}
	return o.Applications, true
}

// HasApplications returns a boolean if a field has been set.
func (o *IosDetails) HasApplications() bool {
	if o != nil && !IsNil(o.Applications) {
		return true
	}

	return false
}

// SetApplications gets a reference to the given []MobileDeviceApplication and assigns it to the Applications field.
func (o *IosDetails) SetApplications(v []MobileDeviceApplication) {
	o.Applications = v
}

// GetCertificates returns the Certificates field value if set, zero value otherwise.
func (o *IosDetails) GetCertificates() []MobileDeviceCertificateV1 {
	if o == nil || IsNil(o.Certificates) {
		var ret []MobileDeviceCertificateV1
		return ret
	}
	return o.Certificates
}

// GetCertificatesOk returns a tuple with the Certificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IosDetails) GetCertificatesOk() ([]MobileDeviceCertificateV1, bool) {
	if o == nil || IsNil(o.Certificates) {
		return nil, false
	}
	return o.Certificates, true
}

// HasCertificates returns a boolean if a field has been set.
func (o *IosDetails) HasCertificates() bool {
	if o != nil && !IsNil(o.Certificates) {
		return true
	}

	return false
}

// SetCertificates gets a reference to the given []MobileDeviceCertificateV1 and assigns it to the Certificates field.
func (o *IosDetails) SetCertificates(v []MobileDeviceCertificateV1) {
	o.Certificates = v
}

// GetEbooks returns the Ebooks field value if set, zero value otherwise.
func (o *IosDetails) GetEbooks() []MobileDeviceEbook {
	if o == nil || IsNil(o.Ebooks) {
		var ret []MobileDeviceEbook
		return ret
	}
	return o.Ebooks
}

// GetEbooksOk returns a tuple with the Ebooks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IosDetails) GetEbooksOk() ([]MobileDeviceEbook, bool) {
	if o == nil || IsNil(o.Ebooks) {
		return nil, false
	}
	return o.Ebooks, true
}

// HasEbooks returns a boolean if a field has been set.
func (o *IosDetails) HasEbooks() bool {
	if o != nil && !IsNil(o.Ebooks) {
		return true
	}

	return false
}

// SetEbooks gets a reference to the given []MobileDeviceEbook and assigns it to the Ebooks field.
func (o *IosDetails) SetEbooks(v []MobileDeviceEbook) {
	o.Ebooks = v
}

// GetConfigurationProfiles returns the ConfigurationProfiles field value if set, zero value otherwise.
func (o *IosDetails) GetConfigurationProfiles() []ConfigurationProfile {
	if o == nil || IsNil(o.ConfigurationProfiles) {
		var ret []ConfigurationProfile
		return ret
	}
	return o.ConfigurationProfiles
}

// GetConfigurationProfilesOk returns a tuple with the ConfigurationProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IosDetails) GetConfigurationProfilesOk() ([]ConfigurationProfile, bool) {
	if o == nil || IsNil(o.ConfigurationProfiles) {
		return nil, false
	}
	return o.ConfigurationProfiles, true
}

// HasConfigurationProfiles returns a boolean if a field has been set.
func (o *IosDetails) HasConfigurationProfiles() bool {
	if o != nil && !IsNil(o.ConfigurationProfiles) {
		return true
	}

	return false
}

// SetConfigurationProfiles gets a reference to the given []ConfigurationProfile and assigns it to the ConfigurationProfiles field.
func (o *IosDetails) SetConfigurationProfiles(v []ConfigurationProfile) {
	o.ConfigurationProfiles = v
}

// GetProvisioningProfiles returns the ProvisioningProfiles field value if set, zero value otherwise.
func (o *IosDetails) GetProvisioningProfiles() []MobileDeviceProvisioningProfiles {
	if o == nil || IsNil(o.ProvisioningProfiles) {
		var ret []MobileDeviceProvisioningProfiles
		return ret
	}
	return o.ProvisioningProfiles
}

// GetProvisioningProfilesOk returns a tuple with the ProvisioningProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IosDetails) GetProvisioningProfilesOk() ([]MobileDeviceProvisioningProfiles, bool) {
	if o == nil || IsNil(o.ProvisioningProfiles) {
		return nil, false
	}
	return o.ProvisioningProfiles, true
}

// HasProvisioningProfiles returns a boolean if a field has been set.
func (o *IosDetails) HasProvisioningProfiles() bool {
	if o != nil && !IsNil(o.ProvisioningProfiles) {
		return true
	}

	return false
}

// SetProvisioningProfiles gets a reference to the given []MobileDeviceProvisioningProfiles and assigns it to the ProvisioningProfiles field.
func (o *IosDetails) SetProvisioningProfiles(v []MobileDeviceProvisioningProfiles) {
	o.ProvisioningProfiles = v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *IosDetails) GetAttachments() []MobileDeviceAttachment {
	if o == nil || IsNil(o.Attachments) {
		var ret []MobileDeviceAttachment
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IosDetails) GetAttachmentsOk() ([]MobileDeviceAttachment, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *IosDetails) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []MobileDeviceAttachment and assigns it to the Attachments field.
func (o *IosDetails) SetAttachments(v []MobileDeviceAttachment) {
	o.Attachments = v
}

func (o IosDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IosDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !IsNil(o.ModelIdentifier) {
		toSerialize["modelIdentifier"] = o.ModelIdentifier
	}
	if !IsNil(o.ModelNumber) {
		toSerialize["modelNumber"] = o.ModelNumber
	}
	if !IsNil(o.IsSupervised) {
		toSerialize["isSupervised"] = o.IsSupervised
	}
	if !IsNil(o.BatteryLevel) {
		toSerialize["batteryLevel"] = o.BatteryLevel
	}
	if !IsNil(o.LastBackupTimestamp) {
		toSerialize["lastBackupTimestamp"] = o.LastBackupTimestamp
	}
	if !IsNil(o.CapacityMb) {
		toSerialize["capacityMb"] = o.CapacityMb
	}
	if !IsNil(o.AvailableMb) {
		toSerialize["availableMb"] = o.AvailableMb
	}
	if !IsNil(o.PercentageUsed) {
		toSerialize["percentageUsed"] = o.PercentageUsed
	}
	if !IsNil(o.IsShared) {
		toSerialize["isShared"] = o.IsShared
	}
	if !IsNil(o.IsDeviceLocatorServiceEnabled) {
		toSerialize["isDeviceLocatorServiceEnabled"] = o.IsDeviceLocatorServiceEnabled
	}
	if !IsNil(o.IsDoNotDisturbEnabled) {
		toSerialize["isDoNotDisturbEnabled"] = o.IsDoNotDisturbEnabled
	}
	if !IsNil(o.IsCloudBackupEnabled) {
		toSerialize["isCloudBackupEnabled"] = o.IsCloudBackupEnabled
	}
	if !IsNil(o.LastCloudBackupTimestamp) {
		toSerialize["lastCloudBackupTimestamp"] = o.LastCloudBackupTimestamp
	}
	if !IsNil(o.IsLocationServicesEnabled) {
		toSerialize["isLocationServicesEnabled"] = o.IsLocationServicesEnabled
	}
	if !IsNil(o.IsITunesStoreAccountActive) {
		toSerialize["isITunesStoreAccountActive"] = o.IsITunesStoreAccountActive
	}
	if !IsNil(o.IsBleCapable) {
		toSerialize["isBleCapable"] = o.IsBleCapable
	}
	if !IsNil(o.Computer) {
		toSerialize["computer"] = o.Computer
	}
	if !IsNil(o.Purchasing) {
		toSerialize["purchasing"] = o.Purchasing
	}
	if !IsNil(o.Security) {
		toSerialize["security"] = o.Security
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.Applications) {
		toSerialize["applications"] = o.Applications
	}
	if !IsNil(o.Certificates) {
		toSerialize["certificates"] = o.Certificates
	}
	if !IsNil(o.Ebooks) {
		toSerialize["ebooks"] = o.Ebooks
	}
	if !IsNil(o.ConfigurationProfiles) {
		toSerialize["configurationProfiles"] = o.ConfigurationProfiles
	}
	if !IsNil(o.ProvisioningProfiles) {
		toSerialize["provisioningProfiles"] = o.ProvisioningProfiles
	}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	return toSerialize, nil
}

type NullableIosDetails struct {
	value *IosDetails
	isSet bool
}

func (v NullableIosDetails) Get() *IosDetails {
	return v.value
}

func (v *NullableIosDetails) Set(val *IosDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableIosDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableIosDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIosDetails(val *IosDetails) *NullableIosDetails {
	return &NullableIosDetails{value: val, isSet: true}
}

func (v NullableIosDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIosDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


