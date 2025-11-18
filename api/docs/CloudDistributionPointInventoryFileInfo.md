# CloudDistributionPointInventoryFileInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A unique identifier for the cloud distribution point inventory file table. | [optional] [readonly] 
**InventoryId** | Pointer to **string** | A unique identifier for the cloud distribution point inventory file and cloud distribution point tables.This ID is used to reference a specific inventory resource within the system. | [optional] [readonly] 
**Type** | **string** | The type of the inventory file. This field indicates whether the file is related to a package, mobile device app, or an ebook. | 
**FileName** | **string** | The name of the inventory file. This could be the name of a package, a mobile device app, or an ebook file, depending on the file type. The name should match the actual file or package name as stored in the cloud distribution system. | 
**FileObjectId** | **string** | A unique identifier for each file type (package, ebook, or mobile device app). This ID is used to construct the URL for accessing or navigating to the specific resource related to the file type. | 
**Category** | Pointer to **string** | The category assigned to the inventory file (package, ebook, or mobile device app) during creation. This helps group and organize files based on their type or purpose, such as security software or productivity tools. | [optional] 
**Status** | **string** | The current status of the inventory file, indicating the progress or outcome of the file&#39;s upload process.It reflects whether the file is ready for use, still being processed, or has encountered an error. | 

## Methods

### NewCloudDistributionPointInventoryFileInfo

`func NewCloudDistributionPointInventoryFileInfo(type_ string, fileName string, fileObjectId string, status string, ) *CloudDistributionPointInventoryFileInfo`

NewCloudDistributionPointInventoryFileInfo instantiates a new CloudDistributionPointInventoryFileInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudDistributionPointInventoryFileInfoWithDefaults

`func NewCloudDistributionPointInventoryFileInfoWithDefaults() *CloudDistributionPointInventoryFileInfo`

NewCloudDistributionPointInventoryFileInfoWithDefaults instantiates a new CloudDistributionPointInventoryFileInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CloudDistributionPointInventoryFileInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudDistributionPointInventoryFileInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudDistributionPointInventoryFileInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudDistributionPointInventoryFileInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInventoryId

`func (o *CloudDistributionPointInventoryFileInfo) GetInventoryId() string`

GetInventoryId returns the InventoryId field if non-nil, zero value otherwise.

### GetInventoryIdOk

`func (o *CloudDistributionPointInventoryFileInfo) GetInventoryIdOk() (*string, bool)`

GetInventoryIdOk returns a tuple with the InventoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryId

`func (o *CloudDistributionPointInventoryFileInfo) SetInventoryId(v string)`

SetInventoryId sets InventoryId field to given value.

### HasInventoryId

`func (o *CloudDistributionPointInventoryFileInfo) HasInventoryId() bool`

HasInventoryId returns a boolean if a field has been set.

### GetType

`func (o *CloudDistributionPointInventoryFileInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudDistributionPointInventoryFileInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudDistributionPointInventoryFileInfo) SetType(v string)`

SetType sets Type field to given value.


### GetFileName

`func (o *CloudDistributionPointInventoryFileInfo) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *CloudDistributionPointInventoryFileInfo) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *CloudDistributionPointInventoryFileInfo) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetFileObjectId

`func (o *CloudDistributionPointInventoryFileInfo) GetFileObjectId() string`

GetFileObjectId returns the FileObjectId field if non-nil, zero value otherwise.

### GetFileObjectIdOk

`func (o *CloudDistributionPointInventoryFileInfo) GetFileObjectIdOk() (*string, bool)`

GetFileObjectIdOk returns a tuple with the FileObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileObjectId

`func (o *CloudDistributionPointInventoryFileInfo) SetFileObjectId(v string)`

SetFileObjectId sets FileObjectId field to given value.


### GetCategory

`func (o *CloudDistributionPointInventoryFileInfo) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CloudDistributionPointInventoryFileInfo) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CloudDistributionPointInventoryFileInfo) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CloudDistributionPointInventoryFileInfo) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetStatus

`func (o *CloudDistributionPointInventoryFileInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudDistributionPointInventoryFileInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudDistributionPointInventoryFileInfo) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


