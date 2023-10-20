# EnrollmentSsoConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hosts** | **[]string** |  | [default to []]
**ManagementHint** | Pointer to **string** |  | [optional] [default to ""]

## Methods

### NewEnrollmentSsoConfig

`func NewEnrollmentSsoConfig(hosts []string, ) *EnrollmentSsoConfig`

NewEnrollmentSsoConfig instantiates a new EnrollmentSsoConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollmentSsoConfigWithDefaults

`func NewEnrollmentSsoConfigWithDefaults() *EnrollmentSsoConfig`

NewEnrollmentSsoConfigWithDefaults instantiates a new EnrollmentSsoConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHosts

`func (o *EnrollmentSsoConfig) GetHosts() []string`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *EnrollmentSsoConfig) GetHostsOk() (*[]string, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *EnrollmentSsoConfig) SetHosts(v []string)`

SetHosts sets Hosts field to given value.


### GetManagementHint

`func (o *EnrollmentSsoConfig) GetManagementHint() string`

GetManagementHint returns the ManagementHint field if non-nil, zero value otherwise.

### GetManagementHintOk

`func (o *EnrollmentSsoConfig) GetManagementHintOk() (*string, bool)`

GetManagementHintOk returns a tuple with the ManagementHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementHint

`func (o *EnrollmentSsoConfig) SetManagementHint(v string)`

SetManagementHint sets ManagementHint field to given value.

### HasManagementHint

`func (o *EnrollmentSsoConfig) HasManagementHint() bool`

HasManagementHint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


