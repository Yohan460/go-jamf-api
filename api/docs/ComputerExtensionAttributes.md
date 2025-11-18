# ComputerExtensionAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique Id for Mobile Device Extension Attribute. | [optional] [readonly] 
**Name** | **string** | Display name for the extension attribute. | 
**Description** | Pointer to **string** | Description for the extension attribute. | [optional] 
**DataType** | **string** | Type of data being collected. | [default to "STRING"]
**Enabled** | Pointer to **bool** | Enabled by default, but for inputType Script we can disable it as well. &lt;br/&gt; Possible values are:&lt;br/&gt; false &lt;br/&gt; true | [optional] [default to true]
**InventoryDisplayType** | **string** | Category in which to display the extension attribute in Jamf Pro. | [default to "GENERAL"]
**InputType** | **string** | Extension attributes collect inventory data by using an input type.The type of the Input used to populate the extension attribute. | [default to "TEXT"]
**ScriptContents** | Pointer to **NullableString** | When we run this script it returns a data value each time a computer submits inventory to Jamf Pro. Provide scriptContents only when inputType is &#39;SCRIPT&#39;. | [optional] 
**PopupMenuChoices** | Pointer to **[]string** | When added with list of choices while creating computer extension attributes these Pop-up menu can be displayed in inventory information. User can choose a value from the pop-up menu list when enrolling a computer any time using Jamf Pro. Provide popupMenuChoices only when inputType is &#39;POPUP&#39;. | [optional] 
**LdapAttributeMapping** | Pointer to **string** | Directory Service attribute use to populate the extension attribute. &lt;br/&gt; Required when inputType is \&quot;DIRECTORY_SERVICE_ATTRIBUTE_MAPPING\&quot; | [optional] 
**LdapExtensionAttributeAllowed** | Pointer to **bool** | Collect multiple values for this extension attribute. ldapExtensionAttributeAllowed is disabled by default, only for inputType &#39;DIRECTORY_SERVICE_ATTRIBUTE_MAPPING&#39; it can be enabled. It&#39;s value cannot be modified during edit operation.&lt;br/&gt; Possible values are:&lt;br/&gt; false &lt;br/&gt; true &lt;br/&gt; | [optional] [default to false]
**ManageExistingData** | Pointer to **string** | It is used to specify to either delete or retain the extension attributes values when inputType is Script and enabled is false. | [optional] 

## Methods

### NewComputerExtensionAttributes

`func NewComputerExtensionAttributes(name string, dataType string, inventoryDisplayType string, inputType string, ) *ComputerExtensionAttributes`

NewComputerExtensionAttributes instantiates a new ComputerExtensionAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputerExtensionAttributesWithDefaults

`func NewComputerExtensionAttributesWithDefaults() *ComputerExtensionAttributes`

NewComputerExtensionAttributesWithDefaults instantiates a new ComputerExtensionAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ComputerExtensionAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ComputerExtensionAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ComputerExtensionAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ComputerExtensionAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ComputerExtensionAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComputerExtensionAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComputerExtensionAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ComputerExtensionAttributes) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ComputerExtensionAttributes) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ComputerExtensionAttributes) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ComputerExtensionAttributes) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDataType

`func (o *ComputerExtensionAttributes) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *ComputerExtensionAttributes) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *ComputerExtensionAttributes) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetEnabled

`func (o *ComputerExtensionAttributes) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ComputerExtensionAttributes) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ComputerExtensionAttributes) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ComputerExtensionAttributes) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetInventoryDisplayType

`func (o *ComputerExtensionAttributes) GetInventoryDisplayType() string`

GetInventoryDisplayType returns the InventoryDisplayType field if non-nil, zero value otherwise.

### GetInventoryDisplayTypeOk

`func (o *ComputerExtensionAttributes) GetInventoryDisplayTypeOk() (*string, bool)`

GetInventoryDisplayTypeOk returns a tuple with the InventoryDisplayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDisplayType

`func (o *ComputerExtensionAttributes) SetInventoryDisplayType(v string)`

SetInventoryDisplayType sets InventoryDisplayType field to given value.


### GetInputType

`func (o *ComputerExtensionAttributes) GetInputType() string`

GetInputType returns the InputType field if non-nil, zero value otherwise.

### GetInputTypeOk

`func (o *ComputerExtensionAttributes) GetInputTypeOk() (*string, bool)`

GetInputTypeOk returns a tuple with the InputType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputType

`func (o *ComputerExtensionAttributes) SetInputType(v string)`

SetInputType sets InputType field to given value.


### GetScriptContents

`func (o *ComputerExtensionAttributes) GetScriptContents() string`

GetScriptContents returns the ScriptContents field if non-nil, zero value otherwise.

### GetScriptContentsOk

`func (o *ComputerExtensionAttributes) GetScriptContentsOk() (*string, bool)`

GetScriptContentsOk returns a tuple with the ScriptContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptContents

`func (o *ComputerExtensionAttributes) SetScriptContents(v string)`

SetScriptContents sets ScriptContents field to given value.

### HasScriptContents

`func (o *ComputerExtensionAttributes) HasScriptContents() bool`

HasScriptContents returns a boolean if a field has been set.

### SetScriptContentsNil

`func (o *ComputerExtensionAttributes) SetScriptContentsNil(b bool)`

 SetScriptContentsNil sets the value for ScriptContents to be an explicit nil

### UnsetScriptContents
`func (o *ComputerExtensionAttributes) UnsetScriptContents()`

UnsetScriptContents ensures that no value is present for ScriptContents, not even an explicit nil
### GetPopupMenuChoices

`func (o *ComputerExtensionAttributes) GetPopupMenuChoices() []string`

GetPopupMenuChoices returns the PopupMenuChoices field if non-nil, zero value otherwise.

### GetPopupMenuChoicesOk

`func (o *ComputerExtensionAttributes) GetPopupMenuChoicesOk() (*[]string, bool)`

GetPopupMenuChoicesOk returns a tuple with the PopupMenuChoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopupMenuChoices

`func (o *ComputerExtensionAttributes) SetPopupMenuChoices(v []string)`

SetPopupMenuChoices sets PopupMenuChoices field to given value.

### HasPopupMenuChoices

`func (o *ComputerExtensionAttributes) HasPopupMenuChoices() bool`

HasPopupMenuChoices returns a boolean if a field has been set.

### GetLdapAttributeMapping

`func (o *ComputerExtensionAttributes) GetLdapAttributeMapping() string`

GetLdapAttributeMapping returns the LdapAttributeMapping field if non-nil, zero value otherwise.

### GetLdapAttributeMappingOk

`func (o *ComputerExtensionAttributes) GetLdapAttributeMappingOk() (*string, bool)`

GetLdapAttributeMappingOk returns a tuple with the LdapAttributeMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapAttributeMapping

`func (o *ComputerExtensionAttributes) SetLdapAttributeMapping(v string)`

SetLdapAttributeMapping sets LdapAttributeMapping field to given value.

### HasLdapAttributeMapping

`func (o *ComputerExtensionAttributes) HasLdapAttributeMapping() bool`

HasLdapAttributeMapping returns a boolean if a field has been set.

### GetLdapExtensionAttributeAllowed

`func (o *ComputerExtensionAttributes) GetLdapExtensionAttributeAllowed() bool`

GetLdapExtensionAttributeAllowed returns the LdapExtensionAttributeAllowed field if non-nil, zero value otherwise.

### GetLdapExtensionAttributeAllowedOk

`func (o *ComputerExtensionAttributes) GetLdapExtensionAttributeAllowedOk() (*bool, bool)`

GetLdapExtensionAttributeAllowedOk returns a tuple with the LdapExtensionAttributeAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapExtensionAttributeAllowed

`func (o *ComputerExtensionAttributes) SetLdapExtensionAttributeAllowed(v bool)`

SetLdapExtensionAttributeAllowed sets LdapExtensionAttributeAllowed field to given value.

### HasLdapExtensionAttributeAllowed

`func (o *ComputerExtensionAttributes) HasLdapExtensionAttributeAllowed() bool`

HasLdapExtensionAttributeAllowed returns a boolean if a field has been set.

### GetManageExistingData

`func (o *ComputerExtensionAttributes) GetManageExistingData() string`

GetManageExistingData returns the ManageExistingData field if non-nil, zero value otherwise.

### GetManageExistingDataOk

`func (o *ComputerExtensionAttributes) GetManageExistingDataOk() (*string, bool)`

GetManageExistingDataOk returns a tuple with the ManageExistingData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageExistingData

`func (o *ComputerExtensionAttributes) SetManageExistingData(v string)`

SetManageExistingData sets ManageExistingData field to given value.

### HasManageExistingData

`func (o *ComputerExtensionAttributes) HasManageExistingData() bool`

HasManageExistingData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


