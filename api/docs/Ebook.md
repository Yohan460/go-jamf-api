# Ebook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Free** | Pointer to **bool** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Author** | Pointer to **string** |  | [optional] 
**DeployAsManaged** | Pointer to **bool** | If true, it will be automatically installed | [optional] 
**InstallAutomatically** | Pointer to **bool** |  | [optional] 
**CategoryId** | Pointer to **string** |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] 

## Methods

### NewEbook

`func NewEbook() *Ebook`

NewEbook instantiates a new Ebook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEbookWithDefaults

`func NewEbookWithDefaults() *Ebook`

NewEbookWithDefaults instantiates a new Ebook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Ebook) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Ebook) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Ebook) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Ebook) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Ebook) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Ebook) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Ebook) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Ebook) HasName() bool`

HasName returns a boolean if a field has been set.

### GetKind

`func (o *Ebook) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Ebook) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Ebook) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Ebook) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetUrl

`func (o *Ebook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Ebook) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Ebook) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Ebook) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetFree

`func (o *Ebook) GetFree() bool`

GetFree returns the Free field if non-nil, zero value otherwise.

### GetFreeOk

`func (o *Ebook) GetFreeOk() (*bool, bool)`

GetFreeOk returns a tuple with the Free field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree

`func (o *Ebook) SetFree(v bool)`

SetFree sets Free field to given value.

### HasFree

`func (o *Ebook) HasFree() bool`

HasFree returns a boolean if a field has been set.

### GetVersion

`func (o *Ebook) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Ebook) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Ebook) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Ebook) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetAuthor

`func (o *Ebook) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *Ebook) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *Ebook) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *Ebook) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetDeployAsManaged

`func (o *Ebook) GetDeployAsManaged() bool`

GetDeployAsManaged returns the DeployAsManaged field if non-nil, zero value otherwise.

### GetDeployAsManagedOk

`func (o *Ebook) GetDeployAsManagedOk() (*bool, bool)`

GetDeployAsManagedOk returns a tuple with the DeployAsManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployAsManaged

`func (o *Ebook) SetDeployAsManaged(v bool)`

SetDeployAsManaged sets DeployAsManaged field to given value.

### HasDeployAsManaged

`func (o *Ebook) HasDeployAsManaged() bool`

HasDeployAsManaged returns a boolean if a field has been set.

### GetInstallAutomatically

`func (o *Ebook) GetInstallAutomatically() bool`

GetInstallAutomatically returns the InstallAutomatically field if non-nil, zero value otherwise.

### GetInstallAutomaticallyOk

`func (o *Ebook) GetInstallAutomaticallyOk() (*bool, bool)`

GetInstallAutomaticallyOk returns a tuple with the InstallAutomatically field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallAutomatically

`func (o *Ebook) SetInstallAutomatically(v bool)`

SetInstallAutomatically sets InstallAutomatically field to given value.

### HasInstallAutomatically

`func (o *Ebook) HasInstallAutomatically() bool`

HasInstallAutomatically returns a boolean if a field has been set.

### GetCategoryId

`func (o *Ebook) GetCategoryId() string`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *Ebook) GetCategoryIdOk() (*string, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *Ebook) SetCategoryId(v string)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *Ebook) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetSiteId

`func (o *Ebook) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *Ebook) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *Ebook) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *Ebook) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


