# CloudLdapKeystore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**ExpirationDate** | Pointer to **time.Time** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**FileName** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudLdapKeystore

`func NewCloudLdapKeystore() *CloudLdapKeystore`

NewCloudLdapKeystore instantiates a new CloudLdapKeystore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudLdapKeystoreWithDefaults

`func NewCloudLdapKeystoreWithDefaults() *CloudLdapKeystore`

NewCloudLdapKeystoreWithDefaults instantiates a new CloudLdapKeystore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CloudLdapKeystore) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudLdapKeystore) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudLdapKeystore) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CloudLdapKeystore) HasType() bool`

HasType returns a boolean if a field has been set.

### GetExpirationDate

`func (o *CloudLdapKeystore) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *CloudLdapKeystore) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *CloudLdapKeystore) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *CloudLdapKeystore) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetSubject

`func (o *CloudLdapKeystore) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CloudLdapKeystore) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CloudLdapKeystore) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CloudLdapKeystore) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetFileName

`func (o *CloudLdapKeystore) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *CloudLdapKeystore) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *CloudLdapKeystore) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *CloudLdapKeystore) HasFileName() bool`

HasFileName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


