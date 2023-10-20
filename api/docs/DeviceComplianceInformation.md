# DeviceComplianceInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **string** | ID of the device | [optional] 
**Applicable** | Pointer to **bool** | If device is applicable for compliance calculation | [optional] 
**ComplianceState** | Pointer to **string** | Device compliance state. Possible values are: * &#x60;UNKNOWN&#x60; for unknow compliance state, this usually means that the compliance state is being calculated, * &#x60;NON_COMPLIANT&#x60; for non compliant state, * &#x60;COMPLIANT&#x60; for compliant state  | [optional] 
**ComplianceVendor** | Pointer to **string** | Name of the compliance vendor | [optional] 
**ComplianceVendorDeviceInformation** | Pointer to [**ComplianceVendorDeviceInformation**](ComplianceVendorDeviceInformation.md) |  | [optional] 

## Methods

### NewDeviceComplianceInformation

`func NewDeviceComplianceInformation() *DeviceComplianceInformation`

NewDeviceComplianceInformation instantiates a new DeviceComplianceInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceComplianceInformationWithDefaults

`func NewDeviceComplianceInformationWithDefaults() *DeviceComplianceInformation`

NewDeviceComplianceInformationWithDefaults instantiates a new DeviceComplianceInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *DeviceComplianceInformation) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *DeviceComplianceInformation) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *DeviceComplianceInformation) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *DeviceComplianceInformation) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetApplicable

`func (o *DeviceComplianceInformation) GetApplicable() bool`

GetApplicable returns the Applicable field if non-nil, zero value otherwise.

### GetApplicableOk

`func (o *DeviceComplianceInformation) GetApplicableOk() (*bool, bool)`

GetApplicableOk returns a tuple with the Applicable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicable

`func (o *DeviceComplianceInformation) SetApplicable(v bool)`

SetApplicable sets Applicable field to given value.

### HasApplicable

`func (o *DeviceComplianceInformation) HasApplicable() bool`

HasApplicable returns a boolean if a field has been set.

### GetComplianceState

`func (o *DeviceComplianceInformation) GetComplianceState() string`

GetComplianceState returns the ComplianceState field if non-nil, zero value otherwise.

### GetComplianceStateOk

`func (o *DeviceComplianceInformation) GetComplianceStateOk() (*string, bool)`

GetComplianceStateOk returns a tuple with the ComplianceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceState

`func (o *DeviceComplianceInformation) SetComplianceState(v string)`

SetComplianceState sets ComplianceState field to given value.

### HasComplianceState

`func (o *DeviceComplianceInformation) HasComplianceState() bool`

HasComplianceState returns a boolean if a field has been set.

### GetComplianceVendor

`func (o *DeviceComplianceInformation) GetComplianceVendor() string`

GetComplianceVendor returns the ComplianceVendor field if non-nil, zero value otherwise.

### GetComplianceVendorOk

`func (o *DeviceComplianceInformation) GetComplianceVendorOk() (*string, bool)`

GetComplianceVendorOk returns a tuple with the ComplianceVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceVendor

`func (o *DeviceComplianceInformation) SetComplianceVendor(v string)`

SetComplianceVendor sets ComplianceVendor field to given value.

### HasComplianceVendor

`func (o *DeviceComplianceInformation) HasComplianceVendor() bool`

HasComplianceVendor returns a boolean if a field has been set.

### GetComplianceVendorDeviceInformation

`func (o *DeviceComplianceInformation) GetComplianceVendorDeviceInformation() ComplianceVendorDeviceInformation`

GetComplianceVendorDeviceInformation returns the ComplianceVendorDeviceInformation field if non-nil, zero value otherwise.

### GetComplianceVendorDeviceInformationOk

`func (o *DeviceComplianceInformation) GetComplianceVendorDeviceInformationOk() (*ComplianceVendorDeviceInformation, bool)`

GetComplianceVendorDeviceInformationOk returns a tuple with the ComplianceVendorDeviceInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceVendorDeviceInformation

`func (o *DeviceComplianceInformation) SetComplianceVendorDeviceInformation(v ComplianceVendorDeviceInformation)`

SetComplianceVendorDeviceInformation sets ComplianceVendorDeviceInformation field to given value.

### HasComplianceVendorDeviceInformation

`func (o *DeviceComplianceInformation) HasComplianceVendorDeviceInformation() bool`

HasComplianceVendorDeviceInformation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


