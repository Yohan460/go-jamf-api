# MobileDeviceExtensionAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique Id for Mobile Device Extension Attribute. | [optional] [readonly] 
**Name** | **string** | Display name for the extension attribute. | 
**Description** | Pointer to **string** | Description for the extension attribute. | [optional] 
**DataType** | **string** | Type of data being collected. | [default to "STRING"]
**InventoryDisplayType** | **string** | Category in which to display the extension attribute in Jamf Pro. | [default to "GENERAL"]
**InputType** | **string** | Extension attributes collect inventory data by using an input type.The type of the Input used to populate the extension attribute. | [default to "TEXT"]
**PopupMenuChoices** | Pointer to **[]string** | When added with list of choices while creating mobile device extension attributes these Pop-up menu can be displayed in inventory information. User can choose a value from the pop-up menu list when enrolling a mobile device any time using Jamf Pro. Provide popupMenuChoices only when inputType is &#39;POPUP&#39;. | [optional] 
**LdapAttributeMapping** | Pointer to **string** | Directory Service attribute use to populate the extension attribute. &lt;br/&gt; Required when inputType is \&quot;DIRECTORY_SERVICE_ATTRIBUTE_MAPPING\&quot; | [optional] 
**LdapExtensionAttributeAllowed** | Pointer to **bool** | Collect multiple values for this extension attribute. ldapExtensionAttributeAllowed is disabled by default, only for inputType &#39;DIRECTORY_SERVICE_ATTRIBUTE_MAPPING&#39; it can be enabled. It&#39;s value cannot be modified during edit operation.&lt;br/&gt; Possible values are:&lt;br/&gt; false &lt;br/&gt; true &lt;br/&gt; | [optional] [default to false]

## Methods

### NewMobileDeviceExtensionAttributes

`func NewMobileDeviceExtensionAttributes(name string, dataType string, inventoryDisplayType string, inputType string, ) *MobileDeviceExtensionAttributes`

NewMobileDeviceExtensionAttributes instantiates a new MobileDeviceExtensionAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileDeviceExtensionAttributesWithDefaults

`func NewMobileDeviceExtensionAttributesWithDefaults() *MobileDeviceExtensionAttributes`

NewMobileDeviceExtensionAttributesWithDefaults instantiates a new MobileDeviceExtensionAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MobileDeviceExtensionAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MobileDeviceExtensionAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MobileDeviceExtensionAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MobileDeviceExtensionAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *MobileDeviceExtensionAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MobileDeviceExtensionAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MobileDeviceExtensionAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *MobileDeviceExtensionAttributes) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MobileDeviceExtensionAttributes) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MobileDeviceExtensionAttributes) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MobileDeviceExtensionAttributes) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDataType

`func (o *MobileDeviceExtensionAttributes) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *MobileDeviceExtensionAttributes) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *MobileDeviceExtensionAttributes) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetInventoryDisplayType

`func (o *MobileDeviceExtensionAttributes) GetInventoryDisplayType() string`

GetInventoryDisplayType returns the InventoryDisplayType field if non-nil, zero value otherwise.

### GetInventoryDisplayTypeOk

`func (o *MobileDeviceExtensionAttributes) GetInventoryDisplayTypeOk() (*string, bool)`

GetInventoryDisplayTypeOk returns a tuple with the InventoryDisplayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDisplayType

`func (o *MobileDeviceExtensionAttributes) SetInventoryDisplayType(v string)`

SetInventoryDisplayType sets InventoryDisplayType field to given value.


### GetInputType

`func (o *MobileDeviceExtensionAttributes) GetInputType() string`

GetInputType returns the InputType field if non-nil, zero value otherwise.

### GetInputTypeOk

`func (o *MobileDeviceExtensionAttributes) GetInputTypeOk() (*string, bool)`

GetInputTypeOk returns a tuple with the InputType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputType

`func (o *MobileDeviceExtensionAttributes) SetInputType(v string)`

SetInputType sets InputType field to given value.


### GetPopupMenuChoices

`func (o *MobileDeviceExtensionAttributes) GetPopupMenuChoices() []string`

GetPopupMenuChoices returns the PopupMenuChoices field if non-nil, zero value otherwise.

### GetPopupMenuChoicesOk

`func (o *MobileDeviceExtensionAttributes) GetPopupMenuChoicesOk() (*[]string, bool)`

GetPopupMenuChoicesOk returns a tuple with the PopupMenuChoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopupMenuChoices

`func (o *MobileDeviceExtensionAttributes) SetPopupMenuChoices(v []string)`

SetPopupMenuChoices sets PopupMenuChoices field to given value.

### HasPopupMenuChoices

`func (o *MobileDeviceExtensionAttributes) HasPopupMenuChoices() bool`

HasPopupMenuChoices returns a boolean if a field has been set.

### GetLdapAttributeMapping

`func (o *MobileDeviceExtensionAttributes) GetLdapAttributeMapping() string`

GetLdapAttributeMapping returns the LdapAttributeMapping field if non-nil, zero value otherwise.

### GetLdapAttributeMappingOk

`func (o *MobileDeviceExtensionAttributes) GetLdapAttributeMappingOk() (*string, bool)`

GetLdapAttributeMappingOk returns a tuple with the LdapAttributeMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapAttributeMapping

`func (o *MobileDeviceExtensionAttributes) SetLdapAttributeMapping(v string)`

SetLdapAttributeMapping sets LdapAttributeMapping field to given value.

### HasLdapAttributeMapping

`func (o *MobileDeviceExtensionAttributes) HasLdapAttributeMapping() bool`

HasLdapAttributeMapping returns a boolean if a field has been set.

### GetLdapExtensionAttributeAllowed

`func (o *MobileDeviceExtensionAttributes) GetLdapExtensionAttributeAllowed() bool`

GetLdapExtensionAttributeAllowed returns the LdapExtensionAttributeAllowed field if non-nil, zero value otherwise.

### GetLdapExtensionAttributeAllowedOk

`func (o *MobileDeviceExtensionAttributes) GetLdapExtensionAttributeAllowedOk() (*bool, bool)`

GetLdapExtensionAttributeAllowedOk returns a tuple with the LdapExtensionAttributeAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapExtensionAttributeAllowed

`func (o *MobileDeviceExtensionAttributes) SetLdapExtensionAttributeAllowed(v bool)`

SetLdapExtensionAttributeAllowed sets LdapExtensionAttributeAllowed field to given value.

### HasLdapExtensionAttributeAllowed

`func (o *MobileDeviceExtensionAttributes) HasLdapExtensionAttributeAllowed() bool`

HasLdapExtensionAttributeAllowed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


