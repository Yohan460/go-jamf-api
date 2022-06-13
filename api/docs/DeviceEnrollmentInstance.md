# DeviceEnrollmentInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Name** | **string** |  | 
**SupervisionIdentityId** | Pointer to **string** |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] 
**ServerName** | Pointer to **string** |  | [optional] [readonly] 
**ServerUuid** | Pointer to **string** |  | [optional] [readonly] 
**AdminId** | Pointer to **string** |  | [optional] [readonly] 
**OrgName** | Pointer to **string** |  | [optional] [readonly] 
**OrgEmail** | Pointer to **string** |  | [optional] [readonly] 
**OrgPhone** | Pointer to **string** |  | [optional] [readonly] 
**OrgAddress** | Pointer to **string** |  | [optional] [readonly] 
**TokenExpirationDate** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewDeviceEnrollmentInstance

`func NewDeviceEnrollmentInstance(name string, ) *DeviceEnrollmentInstance`

NewDeviceEnrollmentInstance instantiates a new DeviceEnrollmentInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceEnrollmentInstanceWithDefaults

`func NewDeviceEnrollmentInstanceWithDefaults() *DeviceEnrollmentInstance`

NewDeviceEnrollmentInstanceWithDefaults instantiates a new DeviceEnrollmentInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeviceEnrollmentInstance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceEnrollmentInstance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceEnrollmentInstance) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeviceEnrollmentInstance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DeviceEnrollmentInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceEnrollmentInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceEnrollmentInstance) SetName(v string)`

SetName sets Name field to given value.


### GetSupervisionIdentityId

`func (o *DeviceEnrollmentInstance) GetSupervisionIdentityId() string`

GetSupervisionIdentityId returns the SupervisionIdentityId field if non-nil, zero value otherwise.

### GetSupervisionIdentityIdOk

`func (o *DeviceEnrollmentInstance) GetSupervisionIdentityIdOk() (*string, bool)`

GetSupervisionIdentityIdOk returns a tuple with the SupervisionIdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupervisionIdentityId

`func (o *DeviceEnrollmentInstance) SetSupervisionIdentityId(v string)`

SetSupervisionIdentityId sets SupervisionIdentityId field to given value.

### HasSupervisionIdentityId

`func (o *DeviceEnrollmentInstance) HasSupervisionIdentityId() bool`

HasSupervisionIdentityId returns a boolean if a field has been set.

### GetSiteId

`func (o *DeviceEnrollmentInstance) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *DeviceEnrollmentInstance) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *DeviceEnrollmentInstance) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *DeviceEnrollmentInstance) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetServerName

`func (o *DeviceEnrollmentInstance) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *DeviceEnrollmentInstance) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *DeviceEnrollmentInstance) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *DeviceEnrollmentInstance) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetServerUuid

`func (o *DeviceEnrollmentInstance) GetServerUuid() string`

GetServerUuid returns the ServerUuid field if non-nil, zero value otherwise.

### GetServerUuidOk

`func (o *DeviceEnrollmentInstance) GetServerUuidOk() (*string, bool)`

GetServerUuidOk returns a tuple with the ServerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUuid

`func (o *DeviceEnrollmentInstance) SetServerUuid(v string)`

SetServerUuid sets ServerUuid field to given value.

### HasServerUuid

`func (o *DeviceEnrollmentInstance) HasServerUuid() bool`

HasServerUuid returns a boolean if a field has been set.

### GetAdminId

`func (o *DeviceEnrollmentInstance) GetAdminId() string`

GetAdminId returns the AdminId field if non-nil, zero value otherwise.

### GetAdminIdOk

`func (o *DeviceEnrollmentInstance) GetAdminIdOk() (*string, bool)`

GetAdminIdOk returns a tuple with the AdminId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminId

`func (o *DeviceEnrollmentInstance) SetAdminId(v string)`

SetAdminId sets AdminId field to given value.

### HasAdminId

`func (o *DeviceEnrollmentInstance) HasAdminId() bool`

HasAdminId returns a boolean if a field has been set.

### GetOrgName

`func (o *DeviceEnrollmentInstance) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *DeviceEnrollmentInstance) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *DeviceEnrollmentInstance) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *DeviceEnrollmentInstance) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetOrgEmail

`func (o *DeviceEnrollmentInstance) GetOrgEmail() string`

GetOrgEmail returns the OrgEmail field if non-nil, zero value otherwise.

### GetOrgEmailOk

`func (o *DeviceEnrollmentInstance) GetOrgEmailOk() (*string, bool)`

GetOrgEmailOk returns a tuple with the OrgEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgEmail

`func (o *DeviceEnrollmentInstance) SetOrgEmail(v string)`

SetOrgEmail sets OrgEmail field to given value.

### HasOrgEmail

`func (o *DeviceEnrollmentInstance) HasOrgEmail() bool`

HasOrgEmail returns a boolean if a field has been set.

### GetOrgPhone

`func (o *DeviceEnrollmentInstance) GetOrgPhone() string`

GetOrgPhone returns the OrgPhone field if non-nil, zero value otherwise.

### GetOrgPhoneOk

`func (o *DeviceEnrollmentInstance) GetOrgPhoneOk() (*string, bool)`

GetOrgPhoneOk returns a tuple with the OrgPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgPhone

`func (o *DeviceEnrollmentInstance) SetOrgPhone(v string)`

SetOrgPhone sets OrgPhone field to given value.

### HasOrgPhone

`func (o *DeviceEnrollmentInstance) HasOrgPhone() bool`

HasOrgPhone returns a boolean if a field has been set.

### GetOrgAddress

`func (o *DeviceEnrollmentInstance) GetOrgAddress() string`

GetOrgAddress returns the OrgAddress field if non-nil, zero value otherwise.

### GetOrgAddressOk

`func (o *DeviceEnrollmentInstance) GetOrgAddressOk() (*string, bool)`

GetOrgAddressOk returns a tuple with the OrgAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgAddress

`func (o *DeviceEnrollmentInstance) SetOrgAddress(v string)`

SetOrgAddress sets OrgAddress field to given value.

### HasOrgAddress

`func (o *DeviceEnrollmentInstance) HasOrgAddress() bool`

HasOrgAddress returns a boolean if a field has been set.

### GetTokenExpirationDate

`func (o *DeviceEnrollmentInstance) GetTokenExpirationDate() string`

GetTokenExpirationDate returns the TokenExpirationDate field if non-nil, zero value otherwise.

### GetTokenExpirationDateOk

`func (o *DeviceEnrollmentInstance) GetTokenExpirationDateOk() (*string, bool)`

GetTokenExpirationDateOk returns a tuple with the TokenExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExpirationDate

`func (o *DeviceEnrollmentInstance) SetTokenExpirationDate(v string)`

SetTokenExpirationDate sets TokenExpirationDate field to given value.

### HasTokenExpirationDate

`func (o *DeviceEnrollmentInstance) HasTokenExpirationDate() bool`

HasTokenExpirationDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


