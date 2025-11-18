# DistributionPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Name** | **string** |  | 
**ServerName** | **string** |  | 
**Principal** | Pointer to **bool** |  | [optional] [default to false]
**BackupDistributionPointId** | Pointer to **string** |  | [optional] [default to "-1"]
**SshUsername** | Pointer to **string** |  | [optional] 
**SshPassword** | Pointer to **string** |  | [optional] 
**LocalPathToShare** | Pointer to **string** |  | [optional] 
**FileSharingConnectionType** | **string** | Specify the type of connection , Either of  fileSharingConnectionType (or) https connection type needs to be enabled using httpsEnabled for a distribution point to be created. | [default to "NONE"]
**ShareName** | Pointer to **string** | Required if fileSharingConnectionType is either AFP (or) SMB | [optional] 
**Workgroup** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int64** | Required if fileSharingConnectionType is either AFP (or) SMB | [optional] [default to 139]
**ReadWriteUsername** | Pointer to **string** | Required if fileSharingConnectionType is either AFP (or) SMB | [optional] 
**ReadWritePassword** | Pointer to **string** | Required if fileSharingConnectionType is either AFP (or) SMB | [optional] 
**ReadOnlyUsername** | Pointer to **string** | Required if fileSharingConnectionType is either AFP (or) SMB | [optional] 
**ReadOnlyPassword** | Pointer to **string** | Required if fileSharingConnectionType is either AFP (or) SMB | [optional] 
**HttpsEnabled** | Pointer to **bool** | Allow downloads over HTTPS - requires installation of a valid SSL certificate | [optional] [default to false]
**HttpsPort** | Pointer to **int64** | Port number of the server - required if HTTPS enabled | [optional] [default to 443]
**HttpsContext** | Pointer to **string** | Path to the share (e.g. if the share is accessible at http://192.168.10.10/JamfShare, the context is \&quot;JamfShare\&quot;) - required if HTTPS enabled | [optional] 
**HttpsSecurityType** | Pointer to **string** | Type of authentication required to download files from the distribution point - required if HTTPS enabled | [optional] [default to "NONE"]
**HttpsUsername** | Pointer to **string** | Required if httpsSecurityType is USERNAME_PASSWORD | [optional] 
**HttpsPassword** | Pointer to **string** | Required if httpsSecurityType is USERNAME_PASSWORD | [optional] 
**EnableLoadBalancing** | Pointer to **bool** | This is used to configure load balancing on the backup distribution point. Cannot be enabled when the backup distribution point configured is cloud. | [optional] [default to false]

## Methods

### NewDistributionPoint

`func NewDistributionPoint(name string, serverName string, fileSharingConnectionType string, ) *DistributionPoint`

NewDistributionPoint instantiates a new DistributionPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDistributionPointWithDefaults

`func NewDistributionPointWithDefaults() *DistributionPoint`

NewDistributionPointWithDefaults instantiates a new DistributionPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DistributionPoint) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DistributionPoint) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DistributionPoint) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DistributionPoint) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DistributionPoint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DistributionPoint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DistributionPoint) SetName(v string)`

SetName sets Name field to given value.


### GetServerName

`func (o *DistributionPoint) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *DistributionPoint) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *DistributionPoint) SetServerName(v string)`

SetServerName sets ServerName field to given value.


### GetPrincipal

`func (o *DistributionPoint) GetPrincipal() bool`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *DistributionPoint) GetPrincipalOk() (*bool, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *DistributionPoint) SetPrincipal(v bool)`

SetPrincipal sets Principal field to given value.

### HasPrincipal

`func (o *DistributionPoint) HasPrincipal() bool`

HasPrincipal returns a boolean if a field has been set.

### GetBackupDistributionPointId

`func (o *DistributionPoint) GetBackupDistributionPointId() string`

GetBackupDistributionPointId returns the BackupDistributionPointId field if non-nil, zero value otherwise.

### GetBackupDistributionPointIdOk

`func (o *DistributionPoint) GetBackupDistributionPointIdOk() (*string, bool)`

GetBackupDistributionPointIdOk returns a tuple with the BackupDistributionPointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupDistributionPointId

`func (o *DistributionPoint) SetBackupDistributionPointId(v string)`

SetBackupDistributionPointId sets BackupDistributionPointId field to given value.

### HasBackupDistributionPointId

`func (o *DistributionPoint) HasBackupDistributionPointId() bool`

HasBackupDistributionPointId returns a boolean if a field has been set.

### GetSshUsername

`func (o *DistributionPoint) GetSshUsername() string`

GetSshUsername returns the SshUsername field if non-nil, zero value otherwise.

### GetSshUsernameOk

`func (o *DistributionPoint) GetSshUsernameOk() (*string, bool)`

GetSshUsernameOk returns a tuple with the SshUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUsername

`func (o *DistributionPoint) SetSshUsername(v string)`

SetSshUsername sets SshUsername field to given value.

### HasSshUsername

`func (o *DistributionPoint) HasSshUsername() bool`

HasSshUsername returns a boolean if a field has been set.

### GetSshPassword

`func (o *DistributionPoint) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *DistributionPoint) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *DistributionPoint) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *DistributionPoint) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.

### GetLocalPathToShare

`func (o *DistributionPoint) GetLocalPathToShare() string`

GetLocalPathToShare returns the LocalPathToShare field if non-nil, zero value otherwise.

### GetLocalPathToShareOk

`func (o *DistributionPoint) GetLocalPathToShareOk() (*string, bool)`

GetLocalPathToShareOk returns a tuple with the LocalPathToShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalPathToShare

`func (o *DistributionPoint) SetLocalPathToShare(v string)`

SetLocalPathToShare sets LocalPathToShare field to given value.

### HasLocalPathToShare

`func (o *DistributionPoint) HasLocalPathToShare() bool`

HasLocalPathToShare returns a boolean if a field has been set.

### GetFileSharingConnectionType

`func (o *DistributionPoint) GetFileSharingConnectionType() string`

GetFileSharingConnectionType returns the FileSharingConnectionType field if non-nil, zero value otherwise.

### GetFileSharingConnectionTypeOk

`func (o *DistributionPoint) GetFileSharingConnectionTypeOk() (*string, bool)`

GetFileSharingConnectionTypeOk returns a tuple with the FileSharingConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSharingConnectionType

`func (o *DistributionPoint) SetFileSharingConnectionType(v string)`

SetFileSharingConnectionType sets FileSharingConnectionType field to given value.


### GetShareName

`func (o *DistributionPoint) GetShareName() string`

GetShareName returns the ShareName field if non-nil, zero value otherwise.

### GetShareNameOk

`func (o *DistributionPoint) GetShareNameOk() (*string, bool)`

GetShareNameOk returns a tuple with the ShareName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareName

`func (o *DistributionPoint) SetShareName(v string)`

SetShareName sets ShareName field to given value.

### HasShareName

`func (o *DistributionPoint) HasShareName() bool`

HasShareName returns a boolean if a field has been set.

### GetWorkgroup

`func (o *DistributionPoint) GetWorkgroup() string`

GetWorkgroup returns the Workgroup field if non-nil, zero value otherwise.

### GetWorkgroupOk

`func (o *DistributionPoint) GetWorkgroupOk() (*string, bool)`

GetWorkgroupOk returns a tuple with the Workgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkgroup

`func (o *DistributionPoint) SetWorkgroup(v string)`

SetWorkgroup sets Workgroup field to given value.

### HasWorkgroup

`func (o *DistributionPoint) HasWorkgroup() bool`

HasWorkgroup returns a boolean if a field has been set.

### GetPort

`func (o *DistributionPoint) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DistributionPoint) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DistributionPoint) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *DistributionPoint) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetReadWriteUsername

`func (o *DistributionPoint) GetReadWriteUsername() string`

GetReadWriteUsername returns the ReadWriteUsername field if non-nil, zero value otherwise.

### GetReadWriteUsernameOk

`func (o *DistributionPoint) GetReadWriteUsernameOk() (*string, bool)`

GetReadWriteUsernameOk returns a tuple with the ReadWriteUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadWriteUsername

`func (o *DistributionPoint) SetReadWriteUsername(v string)`

SetReadWriteUsername sets ReadWriteUsername field to given value.

### HasReadWriteUsername

`func (o *DistributionPoint) HasReadWriteUsername() bool`

HasReadWriteUsername returns a boolean if a field has been set.

### GetReadWritePassword

`func (o *DistributionPoint) GetReadWritePassword() string`

GetReadWritePassword returns the ReadWritePassword field if non-nil, zero value otherwise.

### GetReadWritePasswordOk

`func (o *DistributionPoint) GetReadWritePasswordOk() (*string, bool)`

GetReadWritePasswordOk returns a tuple with the ReadWritePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadWritePassword

`func (o *DistributionPoint) SetReadWritePassword(v string)`

SetReadWritePassword sets ReadWritePassword field to given value.

### HasReadWritePassword

`func (o *DistributionPoint) HasReadWritePassword() bool`

HasReadWritePassword returns a boolean if a field has been set.

### GetReadOnlyUsername

`func (o *DistributionPoint) GetReadOnlyUsername() string`

GetReadOnlyUsername returns the ReadOnlyUsername field if non-nil, zero value otherwise.

### GetReadOnlyUsernameOk

`func (o *DistributionPoint) GetReadOnlyUsernameOk() (*string, bool)`

GetReadOnlyUsernameOk returns a tuple with the ReadOnlyUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnlyUsername

`func (o *DistributionPoint) SetReadOnlyUsername(v string)`

SetReadOnlyUsername sets ReadOnlyUsername field to given value.

### HasReadOnlyUsername

`func (o *DistributionPoint) HasReadOnlyUsername() bool`

HasReadOnlyUsername returns a boolean if a field has been set.

### GetReadOnlyPassword

`func (o *DistributionPoint) GetReadOnlyPassword() string`

GetReadOnlyPassword returns the ReadOnlyPassword field if non-nil, zero value otherwise.

### GetReadOnlyPasswordOk

`func (o *DistributionPoint) GetReadOnlyPasswordOk() (*string, bool)`

GetReadOnlyPasswordOk returns a tuple with the ReadOnlyPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnlyPassword

`func (o *DistributionPoint) SetReadOnlyPassword(v string)`

SetReadOnlyPassword sets ReadOnlyPassword field to given value.

### HasReadOnlyPassword

`func (o *DistributionPoint) HasReadOnlyPassword() bool`

HasReadOnlyPassword returns a boolean if a field has been set.

### GetHttpsEnabled

`func (o *DistributionPoint) GetHttpsEnabled() bool`

GetHttpsEnabled returns the HttpsEnabled field if non-nil, zero value otherwise.

### GetHttpsEnabledOk

`func (o *DistributionPoint) GetHttpsEnabledOk() (*bool, bool)`

GetHttpsEnabledOk returns a tuple with the HttpsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsEnabled

`func (o *DistributionPoint) SetHttpsEnabled(v bool)`

SetHttpsEnabled sets HttpsEnabled field to given value.

### HasHttpsEnabled

`func (o *DistributionPoint) HasHttpsEnabled() bool`

HasHttpsEnabled returns a boolean if a field has been set.

### GetHttpsPort

`func (o *DistributionPoint) GetHttpsPort() int64`

GetHttpsPort returns the HttpsPort field if non-nil, zero value otherwise.

### GetHttpsPortOk

`func (o *DistributionPoint) GetHttpsPortOk() (*int64, bool)`

GetHttpsPortOk returns a tuple with the HttpsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsPort

`func (o *DistributionPoint) SetHttpsPort(v int64)`

SetHttpsPort sets HttpsPort field to given value.

### HasHttpsPort

`func (o *DistributionPoint) HasHttpsPort() bool`

HasHttpsPort returns a boolean if a field has been set.

### GetHttpsContext

`func (o *DistributionPoint) GetHttpsContext() string`

GetHttpsContext returns the HttpsContext field if non-nil, zero value otherwise.

### GetHttpsContextOk

`func (o *DistributionPoint) GetHttpsContextOk() (*string, bool)`

GetHttpsContextOk returns a tuple with the HttpsContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsContext

`func (o *DistributionPoint) SetHttpsContext(v string)`

SetHttpsContext sets HttpsContext field to given value.

### HasHttpsContext

`func (o *DistributionPoint) HasHttpsContext() bool`

HasHttpsContext returns a boolean if a field has been set.

### GetHttpsSecurityType

`func (o *DistributionPoint) GetHttpsSecurityType() string`

GetHttpsSecurityType returns the HttpsSecurityType field if non-nil, zero value otherwise.

### GetHttpsSecurityTypeOk

`func (o *DistributionPoint) GetHttpsSecurityTypeOk() (*string, bool)`

GetHttpsSecurityTypeOk returns a tuple with the HttpsSecurityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsSecurityType

`func (o *DistributionPoint) SetHttpsSecurityType(v string)`

SetHttpsSecurityType sets HttpsSecurityType field to given value.

### HasHttpsSecurityType

`func (o *DistributionPoint) HasHttpsSecurityType() bool`

HasHttpsSecurityType returns a boolean if a field has been set.

### GetHttpsUsername

`func (o *DistributionPoint) GetHttpsUsername() string`

GetHttpsUsername returns the HttpsUsername field if non-nil, zero value otherwise.

### GetHttpsUsernameOk

`func (o *DistributionPoint) GetHttpsUsernameOk() (*string, bool)`

GetHttpsUsernameOk returns a tuple with the HttpsUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsUsername

`func (o *DistributionPoint) SetHttpsUsername(v string)`

SetHttpsUsername sets HttpsUsername field to given value.

### HasHttpsUsername

`func (o *DistributionPoint) HasHttpsUsername() bool`

HasHttpsUsername returns a boolean if a field has been set.

### GetHttpsPassword

`func (o *DistributionPoint) GetHttpsPassword() string`

GetHttpsPassword returns the HttpsPassword field if non-nil, zero value otherwise.

### GetHttpsPasswordOk

`func (o *DistributionPoint) GetHttpsPasswordOk() (*string, bool)`

GetHttpsPasswordOk returns a tuple with the HttpsPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsPassword

`func (o *DistributionPoint) SetHttpsPassword(v string)`

SetHttpsPassword sets HttpsPassword field to given value.

### HasHttpsPassword

`func (o *DistributionPoint) HasHttpsPassword() bool`

HasHttpsPassword returns a boolean if a field has been set.

### GetEnableLoadBalancing

`func (o *DistributionPoint) GetEnableLoadBalancing() bool`

GetEnableLoadBalancing returns the EnableLoadBalancing field if non-nil, zero value otherwise.

### GetEnableLoadBalancingOk

`func (o *DistributionPoint) GetEnableLoadBalancingOk() (*bool, bool)`

GetEnableLoadBalancingOk returns a tuple with the EnableLoadBalancing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLoadBalancing

`func (o *DistributionPoint) SetEnableLoadBalancing(v bool)`

SetEnableLoadBalancing sets EnableLoadBalancing field to given value.

### HasEnableLoadBalancing

`func (o *DistributionPoint) HasEnableLoadBalancing() bool`

HasEnableLoadBalancing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


