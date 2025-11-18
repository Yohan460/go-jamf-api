# Attributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VpnUuid** | Pointer to **string** |  | [optional] 
**AssociatedDomains** | Pointer to **[]string** |  | [optional] 
**Removable** | Pointer to **bool** |  | [optional] 
**EnableDirectDownloads** | Pointer to **bool** |  | [optional] 
**ContentFilterUuid** | Pointer to **string** |  | [optional] 
**DnsProxyUuid** | Pointer to **string** |  | [optional] 
**CellularSliceUuid** | Pointer to **string** |  | [optional] 
**RelayUuid** | Pointer to **string** |  | [optional] 
**TapToPayScreenLock** | Pointer to **bool** |  | [optional] 
**Hideable** | Pointer to **bool** |  | [optional] 
**Lockable** | Pointer to **bool** |  | [optional] 

## Methods

### NewAttributes

`func NewAttributes() *Attributes`

NewAttributes instantiates a new Attributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttributesWithDefaults

`func NewAttributesWithDefaults() *Attributes`

NewAttributesWithDefaults instantiates a new Attributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVpnUuid

`func (o *Attributes) GetVpnUuid() string`

GetVpnUuid returns the VpnUuid field if non-nil, zero value otherwise.

### GetVpnUuidOk

`func (o *Attributes) GetVpnUuidOk() (*string, bool)`

GetVpnUuidOk returns a tuple with the VpnUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnUuid

`func (o *Attributes) SetVpnUuid(v string)`

SetVpnUuid sets VpnUuid field to given value.

### HasVpnUuid

`func (o *Attributes) HasVpnUuid() bool`

HasVpnUuid returns a boolean if a field has been set.

### GetAssociatedDomains

`func (o *Attributes) GetAssociatedDomains() []string`

GetAssociatedDomains returns the AssociatedDomains field if non-nil, zero value otherwise.

### GetAssociatedDomainsOk

`func (o *Attributes) GetAssociatedDomainsOk() (*[]string, bool)`

GetAssociatedDomainsOk returns a tuple with the AssociatedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedDomains

`func (o *Attributes) SetAssociatedDomains(v []string)`

SetAssociatedDomains sets AssociatedDomains field to given value.

### HasAssociatedDomains

`func (o *Attributes) HasAssociatedDomains() bool`

HasAssociatedDomains returns a boolean if a field has been set.

### GetRemovable

`func (o *Attributes) GetRemovable() bool`

GetRemovable returns the Removable field if non-nil, zero value otherwise.

### GetRemovableOk

`func (o *Attributes) GetRemovableOk() (*bool, bool)`

GetRemovableOk returns a tuple with the Removable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovable

`func (o *Attributes) SetRemovable(v bool)`

SetRemovable sets Removable field to given value.

### HasRemovable

`func (o *Attributes) HasRemovable() bool`

HasRemovable returns a boolean if a field has been set.

### GetEnableDirectDownloads

`func (o *Attributes) GetEnableDirectDownloads() bool`

GetEnableDirectDownloads returns the EnableDirectDownloads field if non-nil, zero value otherwise.

### GetEnableDirectDownloadsOk

`func (o *Attributes) GetEnableDirectDownloadsOk() (*bool, bool)`

GetEnableDirectDownloadsOk returns a tuple with the EnableDirectDownloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDirectDownloads

`func (o *Attributes) SetEnableDirectDownloads(v bool)`

SetEnableDirectDownloads sets EnableDirectDownloads field to given value.

### HasEnableDirectDownloads

`func (o *Attributes) HasEnableDirectDownloads() bool`

HasEnableDirectDownloads returns a boolean if a field has been set.

### GetContentFilterUuid

`func (o *Attributes) GetContentFilterUuid() string`

GetContentFilterUuid returns the ContentFilterUuid field if non-nil, zero value otherwise.

### GetContentFilterUuidOk

`func (o *Attributes) GetContentFilterUuidOk() (*string, bool)`

GetContentFilterUuidOk returns a tuple with the ContentFilterUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentFilterUuid

`func (o *Attributes) SetContentFilterUuid(v string)`

SetContentFilterUuid sets ContentFilterUuid field to given value.

### HasContentFilterUuid

`func (o *Attributes) HasContentFilterUuid() bool`

HasContentFilterUuid returns a boolean if a field has been set.

### GetDnsProxyUuid

`func (o *Attributes) GetDnsProxyUuid() string`

GetDnsProxyUuid returns the DnsProxyUuid field if non-nil, zero value otherwise.

### GetDnsProxyUuidOk

`func (o *Attributes) GetDnsProxyUuidOk() (*string, bool)`

GetDnsProxyUuidOk returns a tuple with the DnsProxyUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsProxyUuid

`func (o *Attributes) SetDnsProxyUuid(v string)`

SetDnsProxyUuid sets DnsProxyUuid field to given value.

### HasDnsProxyUuid

`func (o *Attributes) HasDnsProxyUuid() bool`

HasDnsProxyUuid returns a boolean if a field has been set.

### GetCellularSliceUuid

`func (o *Attributes) GetCellularSliceUuid() string`

GetCellularSliceUuid returns the CellularSliceUuid field if non-nil, zero value otherwise.

### GetCellularSliceUuidOk

`func (o *Attributes) GetCellularSliceUuidOk() (*string, bool)`

GetCellularSliceUuidOk returns a tuple with the CellularSliceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellularSliceUuid

`func (o *Attributes) SetCellularSliceUuid(v string)`

SetCellularSliceUuid sets CellularSliceUuid field to given value.

### HasCellularSliceUuid

`func (o *Attributes) HasCellularSliceUuid() bool`

HasCellularSliceUuid returns a boolean if a field has been set.

### GetRelayUuid

`func (o *Attributes) GetRelayUuid() string`

GetRelayUuid returns the RelayUuid field if non-nil, zero value otherwise.

### GetRelayUuidOk

`func (o *Attributes) GetRelayUuidOk() (*string, bool)`

GetRelayUuidOk returns a tuple with the RelayUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayUuid

`func (o *Attributes) SetRelayUuid(v string)`

SetRelayUuid sets RelayUuid field to given value.

### HasRelayUuid

`func (o *Attributes) HasRelayUuid() bool`

HasRelayUuid returns a boolean if a field has been set.

### GetTapToPayScreenLock

`func (o *Attributes) GetTapToPayScreenLock() bool`

GetTapToPayScreenLock returns the TapToPayScreenLock field if non-nil, zero value otherwise.

### GetTapToPayScreenLockOk

`func (o *Attributes) GetTapToPayScreenLockOk() (*bool, bool)`

GetTapToPayScreenLockOk returns a tuple with the TapToPayScreenLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTapToPayScreenLock

`func (o *Attributes) SetTapToPayScreenLock(v bool)`

SetTapToPayScreenLock sets TapToPayScreenLock field to given value.

### HasTapToPayScreenLock

`func (o *Attributes) HasTapToPayScreenLock() bool`

HasTapToPayScreenLock returns a boolean if a field has been set.

### GetHideable

`func (o *Attributes) GetHideable() bool`

GetHideable returns the Hideable field if non-nil, zero value otherwise.

### GetHideableOk

`func (o *Attributes) GetHideableOk() (*bool, bool)`

GetHideableOk returns a tuple with the Hideable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideable

`func (o *Attributes) SetHideable(v bool)`

SetHideable sets Hideable field to given value.

### HasHideable

`func (o *Attributes) HasHideable() bool`

HasHideable returns a boolean if a field has been set.

### GetLockable

`func (o *Attributes) GetLockable() bool`

GetLockable returns the Lockable field if non-nil, zero value otherwise.

### GetLockableOk

`func (o *Attributes) GetLockableOk() (*bool, bool)`

GetLockableOk returns a tuple with the Lockable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockable

`func (o *Attributes) SetLockable(v bool)`

SetLockable sets Lockable field to given value.

### HasLockable

`func (o *Attributes) HasLockable() bool`

HasLockable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


