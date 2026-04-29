# WellKnownSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgName** | Pointer to **string** | Organization display name | [optional] 
**ServerUuid** | **string** | Server UUID identifier | 
**EnrollmentType** | [**ServiceDiscoveryVersion**](ServiceDiscoveryVersion.md) |  | 

## Methods

### NewWellKnownSetting

`func NewWellKnownSetting(serverUuid string, enrollmentType ServiceDiscoveryVersion, ) *WellKnownSetting`

NewWellKnownSetting instantiates a new WellKnownSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWellKnownSettingWithDefaults

`func NewWellKnownSettingWithDefaults() *WellKnownSetting`

NewWellKnownSettingWithDefaults instantiates a new WellKnownSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgName

`func (o *WellKnownSetting) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *WellKnownSetting) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *WellKnownSetting) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *WellKnownSetting) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetServerUuid

`func (o *WellKnownSetting) GetServerUuid() string`

GetServerUuid returns the ServerUuid field if non-nil, zero value otherwise.

### GetServerUuidOk

`func (o *WellKnownSetting) GetServerUuidOk() (*string, bool)`

GetServerUuidOk returns a tuple with the ServerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUuid

`func (o *WellKnownSetting) SetServerUuid(v string)`

SetServerUuid sets ServerUuid field to given value.


### GetEnrollmentType

`func (o *WellKnownSetting) GetEnrollmentType() ServiceDiscoveryVersion`

GetEnrollmentType returns the EnrollmentType field if non-nil, zero value otherwise.

### GetEnrollmentTypeOk

`func (o *WellKnownSetting) GetEnrollmentTypeOk() (*ServiceDiscoveryVersion, bool)`

GetEnrollmentTypeOk returns a tuple with the EnrollmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentType

`func (o *WellKnownSetting) SetEnrollmentType(v ServiceDiscoveryVersion)`

SetEnrollmentType sets EnrollmentType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


