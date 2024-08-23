# DockItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Name** | **string** |  | 
**Type** | **string** |  | 
**Path** | **string** |  | 
**Contents** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewDockItem

`func NewDockItem(name string, type_ string, path string, ) *DockItem`

NewDockItem instantiates a new DockItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDockItemWithDefaults

`func NewDockItemWithDefaults() *DockItem`

NewDockItemWithDefaults instantiates a new DockItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DockItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DockItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DockItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DockItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DockItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DockItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DockItem) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *DockItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DockItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DockItem) SetType(v string)`

SetType sets Type field to given value.


### GetPath

`func (o *DockItem) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DockItem) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DockItem) SetPath(v string)`

SetPath sets Path field to given value.


### GetContents

`func (o *DockItem) GetContents() string`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *DockItem) GetContentsOk() (*string, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *DockItem) SetContents(v string)`

SetContents sets Contents field to given value.

### HasContents

`func (o *DockItem) HasContents() bool`

HasContents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


