# PackageManifest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**Hash** | **string** |  | 
**HashType** | **string** |  | 
**DisplayImageUrl** | Pointer to **string** |  | [optional] 
**FullSizeImageUrl** | Pointer to **string** |  | [optional] 
**BundleId** | **string** |  | 
**BundleVersion** | **string** |  | 
**Subtitle** | Pointer to **string** |  | [optional] 
**Title** | **string** |  | 
**SizeInBytes** | **int32** |  | 

## Methods

### NewPackageManifest

`func NewPackageManifest(url string, hash string, hashType string, bundleId string, bundleVersion string, title string, sizeInBytes int32, ) *PackageManifest`

NewPackageManifest instantiates a new PackageManifest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageManifestWithDefaults

`func NewPackageManifestWithDefaults() *PackageManifest`

NewPackageManifestWithDefaults instantiates a new PackageManifest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *PackageManifest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PackageManifest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PackageManifest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHash

`func (o *PackageManifest) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *PackageManifest) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *PackageManifest) SetHash(v string)`

SetHash sets Hash field to given value.


### GetHashType

`func (o *PackageManifest) GetHashType() string`

GetHashType returns the HashType field if non-nil, zero value otherwise.

### GetHashTypeOk

`func (o *PackageManifest) GetHashTypeOk() (*string, bool)`

GetHashTypeOk returns a tuple with the HashType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashType

`func (o *PackageManifest) SetHashType(v string)`

SetHashType sets HashType field to given value.


### GetDisplayImageUrl

`func (o *PackageManifest) GetDisplayImageUrl() string`

GetDisplayImageUrl returns the DisplayImageUrl field if non-nil, zero value otherwise.

### GetDisplayImageUrlOk

`func (o *PackageManifest) GetDisplayImageUrlOk() (*string, bool)`

GetDisplayImageUrlOk returns a tuple with the DisplayImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayImageUrl

`func (o *PackageManifest) SetDisplayImageUrl(v string)`

SetDisplayImageUrl sets DisplayImageUrl field to given value.

### HasDisplayImageUrl

`func (o *PackageManifest) HasDisplayImageUrl() bool`

HasDisplayImageUrl returns a boolean if a field has been set.

### GetFullSizeImageUrl

`func (o *PackageManifest) GetFullSizeImageUrl() string`

GetFullSizeImageUrl returns the FullSizeImageUrl field if non-nil, zero value otherwise.

### GetFullSizeImageUrlOk

`func (o *PackageManifest) GetFullSizeImageUrlOk() (*string, bool)`

GetFullSizeImageUrlOk returns a tuple with the FullSizeImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullSizeImageUrl

`func (o *PackageManifest) SetFullSizeImageUrl(v string)`

SetFullSizeImageUrl sets FullSizeImageUrl field to given value.

### HasFullSizeImageUrl

`func (o *PackageManifest) HasFullSizeImageUrl() bool`

HasFullSizeImageUrl returns a boolean if a field has been set.

### GetBundleId

`func (o *PackageManifest) GetBundleId() string`

GetBundleId returns the BundleId field if non-nil, zero value otherwise.

### GetBundleIdOk

`func (o *PackageManifest) GetBundleIdOk() (*string, bool)`

GetBundleIdOk returns a tuple with the BundleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleId

`func (o *PackageManifest) SetBundleId(v string)`

SetBundleId sets BundleId field to given value.


### GetBundleVersion

`func (o *PackageManifest) GetBundleVersion() string`

GetBundleVersion returns the BundleVersion field if non-nil, zero value otherwise.

### GetBundleVersionOk

`func (o *PackageManifest) GetBundleVersionOk() (*string, bool)`

GetBundleVersionOk returns a tuple with the BundleVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleVersion

`func (o *PackageManifest) SetBundleVersion(v string)`

SetBundleVersion sets BundleVersion field to given value.


### GetSubtitle

`func (o *PackageManifest) GetSubtitle() string`

GetSubtitle returns the Subtitle field if non-nil, zero value otherwise.

### GetSubtitleOk

`func (o *PackageManifest) GetSubtitleOk() (*string, bool)`

GetSubtitleOk returns a tuple with the Subtitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitle

`func (o *PackageManifest) SetSubtitle(v string)`

SetSubtitle sets Subtitle field to given value.

### HasSubtitle

`func (o *PackageManifest) HasSubtitle() bool`

HasSubtitle returns a boolean if a field has been set.

### GetTitle

`func (o *PackageManifest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PackageManifest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PackageManifest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetSizeInBytes

`func (o *PackageManifest) GetSizeInBytes() int32`

GetSizeInBytes returns the SizeInBytes field if non-nil, zero value otherwise.

### GetSizeInBytesOk

`func (o *PackageManifest) GetSizeInBytesOk() (*int32, bool)`

GetSizeInBytesOk returns a tuple with the SizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeInBytes

`func (o *PackageManifest) SetSizeInBytes(v int32)`

SetSizeInBytes sets SizeInBytes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


