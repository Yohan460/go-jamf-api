# EnrollmentCustomizationV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**SiteId** | Pointer to **string** |  | [optional] 
**DisplayName** | **string** |  | 
**Description** | **string** |  | 
**EnrollmentCustomizationBrandingSettings** | [**EnrollmentCustomizationBrandingSettings**](EnrollmentCustomizationBrandingSettings.md) |  | 

## Methods

### NewEnrollmentCustomizationV2

`func NewEnrollmentCustomizationV2(displayName string, description string, enrollmentCustomizationBrandingSettings EnrollmentCustomizationBrandingSettings, ) *EnrollmentCustomizationV2`

NewEnrollmentCustomizationV2 instantiates a new EnrollmentCustomizationV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollmentCustomizationV2WithDefaults

`func NewEnrollmentCustomizationV2WithDefaults() *EnrollmentCustomizationV2`

NewEnrollmentCustomizationV2WithDefaults instantiates a new EnrollmentCustomizationV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnrollmentCustomizationV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnrollmentCustomizationV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnrollmentCustomizationV2) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EnrollmentCustomizationV2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSiteId

`func (o *EnrollmentCustomizationV2) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *EnrollmentCustomizationV2) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *EnrollmentCustomizationV2) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *EnrollmentCustomizationV2) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetDisplayName

`func (o *EnrollmentCustomizationV2) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *EnrollmentCustomizationV2) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *EnrollmentCustomizationV2) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDescription

`func (o *EnrollmentCustomizationV2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnrollmentCustomizationV2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnrollmentCustomizationV2) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEnrollmentCustomizationBrandingSettings

`func (o *EnrollmentCustomizationV2) GetEnrollmentCustomizationBrandingSettings() EnrollmentCustomizationBrandingSettings`

GetEnrollmentCustomizationBrandingSettings returns the EnrollmentCustomizationBrandingSettings field if non-nil, zero value otherwise.

### GetEnrollmentCustomizationBrandingSettingsOk

`func (o *EnrollmentCustomizationV2) GetEnrollmentCustomizationBrandingSettingsOk() (*EnrollmentCustomizationBrandingSettings, bool)`

GetEnrollmentCustomizationBrandingSettingsOk returns a tuple with the EnrollmentCustomizationBrandingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentCustomizationBrandingSettings

`func (o *EnrollmentCustomizationV2) SetEnrollmentCustomizationBrandingSettings(v EnrollmentCustomizationBrandingSettings)`

SetEnrollmentCustomizationBrandingSettings sets EnrollmentCustomizationBrandingSettings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


