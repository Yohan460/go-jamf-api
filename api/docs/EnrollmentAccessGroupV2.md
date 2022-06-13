# EnrollmentAccessGroupV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Group ID. | [optional] [readonly] 
**LdapServerId** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] 
**EnterpriseEnrollmentEnabled** | Pointer to **bool** |  | [optional] 
**PersonalEnrollmentEnabled** | Pointer to **bool** |  | [optional] 
**AccountDrivenUserEnrollmentEnabled** | Pointer to **bool** |  | [optional] 
**RequireEula** | Pointer to **bool** |  | [optional] 

## Methods

### NewEnrollmentAccessGroupV2

`func NewEnrollmentAccessGroupV2() *EnrollmentAccessGroupV2`

NewEnrollmentAccessGroupV2 instantiates a new EnrollmentAccessGroupV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollmentAccessGroupV2WithDefaults

`func NewEnrollmentAccessGroupV2WithDefaults() *EnrollmentAccessGroupV2`

NewEnrollmentAccessGroupV2WithDefaults instantiates a new EnrollmentAccessGroupV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnrollmentAccessGroupV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnrollmentAccessGroupV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnrollmentAccessGroupV2) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EnrollmentAccessGroupV2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLdapServerId

`func (o *EnrollmentAccessGroupV2) GetLdapServerId() string`

GetLdapServerId returns the LdapServerId field if non-nil, zero value otherwise.

### GetLdapServerIdOk

`func (o *EnrollmentAccessGroupV2) GetLdapServerIdOk() (*string, bool)`

GetLdapServerIdOk returns a tuple with the LdapServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapServerId

`func (o *EnrollmentAccessGroupV2) SetLdapServerId(v string)`

SetLdapServerId sets LdapServerId field to given value.

### HasLdapServerId

`func (o *EnrollmentAccessGroupV2) HasLdapServerId() bool`

HasLdapServerId returns a boolean if a field has been set.

### GetName

`func (o *EnrollmentAccessGroupV2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnrollmentAccessGroupV2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnrollmentAccessGroupV2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EnrollmentAccessGroupV2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSiteId

`func (o *EnrollmentAccessGroupV2) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *EnrollmentAccessGroupV2) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *EnrollmentAccessGroupV2) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *EnrollmentAccessGroupV2) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetEnterpriseEnrollmentEnabled

`func (o *EnrollmentAccessGroupV2) GetEnterpriseEnrollmentEnabled() bool`

GetEnterpriseEnrollmentEnabled returns the EnterpriseEnrollmentEnabled field if non-nil, zero value otherwise.

### GetEnterpriseEnrollmentEnabledOk

`func (o *EnrollmentAccessGroupV2) GetEnterpriseEnrollmentEnabledOk() (*bool, bool)`

GetEnterpriseEnrollmentEnabledOk returns a tuple with the EnterpriseEnrollmentEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseEnrollmentEnabled

`func (o *EnrollmentAccessGroupV2) SetEnterpriseEnrollmentEnabled(v bool)`

SetEnterpriseEnrollmentEnabled sets EnterpriseEnrollmentEnabled field to given value.

### HasEnterpriseEnrollmentEnabled

`func (o *EnrollmentAccessGroupV2) HasEnterpriseEnrollmentEnabled() bool`

HasEnterpriseEnrollmentEnabled returns a boolean if a field has been set.

### GetPersonalEnrollmentEnabled

`func (o *EnrollmentAccessGroupV2) GetPersonalEnrollmentEnabled() bool`

GetPersonalEnrollmentEnabled returns the PersonalEnrollmentEnabled field if non-nil, zero value otherwise.

### GetPersonalEnrollmentEnabledOk

`func (o *EnrollmentAccessGroupV2) GetPersonalEnrollmentEnabledOk() (*bool, bool)`

GetPersonalEnrollmentEnabledOk returns a tuple with the PersonalEnrollmentEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalEnrollmentEnabled

`func (o *EnrollmentAccessGroupV2) SetPersonalEnrollmentEnabled(v bool)`

SetPersonalEnrollmentEnabled sets PersonalEnrollmentEnabled field to given value.

### HasPersonalEnrollmentEnabled

`func (o *EnrollmentAccessGroupV2) HasPersonalEnrollmentEnabled() bool`

HasPersonalEnrollmentEnabled returns a boolean if a field has been set.

### GetAccountDrivenUserEnrollmentEnabled

`func (o *EnrollmentAccessGroupV2) GetAccountDrivenUserEnrollmentEnabled() bool`

GetAccountDrivenUserEnrollmentEnabled returns the AccountDrivenUserEnrollmentEnabled field if non-nil, zero value otherwise.

### GetAccountDrivenUserEnrollmentEnabledOk

`func (o *EnrollmentAccessGroupV2) GetAccountDrivenUserEnrollmentEnabledOk() (*bool, bool)`

GetAccountDrivenUserEnrollmentEnabledOk returns a tuple with the AccountDrivenUserEnrollmentEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountDrivenUserEnrollmentEnabled

`func (o *EnrollmentAccessGroupV2) SetAccountDrivenUserEnrollmentEnabled(v bool)`

SetAccountDrivenUserEnrollmentEnabled sets AccountDrivenUserEnrollmentEnabled field to given value.

### HasAccountDrivenUserEnrollmentEnabled

`func (o *EnrollmentAccessGroupV2) HasAccountDrivenUserEnrollmentEnabled() bool`

HasAccountDrivenUserEnrollmentEnabled returns a boolean if a field has been set.

### GetRequireEula

`func (o *EnrollmentAccessGroupV2) GetRequireEula() bool`

GetRequireEula returns the RequireEula field if non-nil, zero value otherwise.

### GetRequireEulaOk

`func (o *EnrollmentAccessGroupV2) GetRequireEulaOk() (*bool, bool)`

GetRequireEulaOk returns a tuple with the RequireEula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireEula

`func (o *EnrollmentAccessGroupV2) SetRequireEula(v bool)`

SetRequireEula sets RequireEula field to given value.

### HasRequireEula

`func (o *EnrollmentAccessGroupV2) HasRequireEula() bool`

HasRequireEula returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


