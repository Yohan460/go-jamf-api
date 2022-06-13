# CloudLdapConnectionPoolStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumConnectionsClosedDefunct** | Pointer to **int64** |  | [optional] 
**NumConnectionsClosedExpired** | Pointer to **int64** |  | [optional] 
**NumConnectionsClosedUnneeded** | Pointer to **int64** |  | [optional] 
**NumFailedCheckouts** | Pointer to **int64** |  | [optional] 
**NumFailedConnectionAttempts** | Pointer to **int64** |  | [optional] 
**NumReleasedValid** | Pointer to **int64** |  | [optional] 
**NumSuccessfulCheckouts** | Pointer to **int64** |  | [optional] 
**NumSuccessfulCheckoutsNewConnection** | Pointer to **int64** |  | [optional] 
**NumSuccessfulConnectionAttempts** | Pointer to **int64** |  | [optional] 
**MaximumAvailableConnections** | Pointer to **int64** |  | [optional] 
**NumSuccessfulCheckoutsWithoutWaiting** | Pointer to **int64** |  | [optional] 
**NumSuccessfulCheckoutsAfterWaiting** | Pointer to **int64** |  | [optional] 
**NumAvailableConnections** | Pointer to **int64** |  | [optional] 

## Methods

### NewCloudLdapConnectionPoolStatistics

`func NewCloudLdapConnectionPoolStatistics() *CloudLdapConnectionPoolStatistics`

NewCloudLdapConnectionPoolStatistics instantiates a new CloudLdapConnectionPoolStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudLdapConnectionPoolStatisticsWithDefaults

`func NewCloudLdapConnectionPoolStatisticsWithDefaults() *CloudLdapConnectionPoolStatistics`

NewCloudLdapConnectionPoolStatisticsWithDefaults instantiates a new CloudLdapConnectionPoolStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumConnectionsClosedDefunct

`func (o *CloudLdapConnectionPoolStatistics) GetNumConnectionsClosedDefunct() int64`

GetNumConnectionsClosedDefunct returns the NumConnectionsClosedDefunct field if non-nil, zero value otherwise.

### GetNumConnectionsClosedDefunctOk

`func (o *CloudLdapConnectionPoolStatistics) GetNumConnectionsClosedDefunctOk() (*int64, bool)`

GetNumConnectionsClosedDefunctOk returns a tuple with the NumConnectionsClosedDefunct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumConnectionsClosedDefunct

`func (o *CloudLdapConnectionPoolStatistics) SetNumConnectionsClosedDefunct(v int64)`

SetNumConnectionsClosedDefunct sets NumConnectionsClosedDefunct field to given value.

### HasNumConnectionsClosedDefunct

`func (o *CloudLdapConnectionPoolStatistics) HasNumConnectionsClosedDefunct() bool`

HasNumConnectionsClosedDefunct returns a boolean if a field has been set.

### GetNumConnectionsClosedExpired

`func (o *CloudLdapConnectionPoolStatistics) GetNumConnectionsClosedExpired() int64`

GetNumConnectionsClosedExpired returns the NumConnectionsClosedExpired field if non-nil, zero value otherwise.

### GetNumConnectionsClosedExpiredOk

`func (o *CloudLdapConnectionPoolStatistics) GetNumConnectionsClosedExpiredOk() (*int64, bool)`

GetNumConnectionsClosedExpiredOk returns a tuple with the NumConnectionsClosedExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumConnectionsClosedExpired

`func (o *CloudLdapConnectionPoolStatistics) SetNumConnectionsClosedExpired(v int64)`

SetNumConnectionsClosedExpired sets NumConnectionsClosedExpired field to given value.

### HasNumConnectionsClosedExpired

`func (o *CloudLdapConnectionPoolStatistics) HasNumConnectionsClosedExpired() bool`

HasNumConnectionsClosedExpired returns a boolean if a field has been set.

### GetNumConnectionsClosedUnneeded

`func (o *CloudLdapConnectionPoolStatistics) GetNumConnectionsClosedUnneeded() int64`

GetNumConnectionsClosedUnneeded returns the NumConnectionsClosedUnneeded field if non-nil, zero value otherwise.

### GetNumConnectionsClosedUnneededOk

`func (o *CloudLdapConnectionPoolStatistics) GetNumConnectionsClosedUnneededOk() (*int64, bool)`

GetNumConnectionsClosedUnneededOk returns a tuple with the NumConnectionsClosedUnneeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumConnectionsClosedUnneeded

`func (o *CloudLdapConnectionPoolStatistics) SetNumConnectionsClosedUnneeded(v int64)`

SetNumConnectionsClosedUnneeded sets NumConnectionsClosedUnneeded field to given value.

### HasNumConnectionsClosedUnneeded

`func (o *CloudLdapConnectionPoolStatistics) HasNumConnectionsClosedUnneeded() bool`

HasNumConnectionsClosedUnneeded returns a boolean if a field has been set.

### GetNumFailedCheckouts

`func (o *CloudLdapConnectionPoolStatistics) GetNumFailedCheckouts() int64`

GetNumFailedCheckouts returns the NumFailedCheckouts field if non-nil, zero value otherwise.

### GetNumFailedCheckoutsOk

`func (o *CloudLdapConnectionPoolStatistics) GetNumFailedCheckoutsOk() (*int64, bool)`

GetNumFailedCheckoutsOk returns a tuple with the NumFailedCheckouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFailedCheckouts

`func (o *CloudLdapConnectionPoolStatistics) SetNumFailedCheckouts(v int64)`

SetNumFailedCheckouts sets NumFailedCheckouts field to given value.

### HasNumFailedCheckouts

`func (o *CloudLdapConnectionPoolStatistics) HasNumFailedCheckouts() bool`

HasNumFailedCheckouts returns a boolean if a field has been set.

### GetNumFailedConnectionAttempts

`func (o *CloudLdapConnectionPoolStatistics) GetNumFailedConnectionAttempts() int64`

GetNumFailedConnectionAttempts returns the NumFailedConnectionAttempts field if non-nil, zero value otherwise.

### GetNumFailedConnectionAttemptsOk

`func (o *CloudLdapConnectionPoolStatistics) GetNumFailedConnectionAttemptsOk() (*int64, bool)`

GetNumFailedConnectionAttemptsOk returns a tuple with the NumFailedConnectionAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFailedConnectionAttempts

`func (o *CloudLdapConnectionPoolStatistics) SetNumFailedConnectionAttempts(v int64)`

SetNumFailedConnectionAttempts sets NumFailedConnectionAttempts field to given value.

### HasNumFailedConnectionAttempts

`func (o *CloudLdapConnectionPoolStatistics) HasNumFailedConnectionAttempts() bool`

HasNumFailedConnectionAttempts returns a boolean if a field has been set.

### GetNumReleasedValid

`func (o *CloudLdapConnectionPoolStatistics) GetNumReleasedValid() int64`

GetNumReleasedValid returns the NumReleasedValid field if non-nil, zero value otherwise.

### GetNumReleasedValidOk

`func (o *CloudLdapConnectionPoolStatistics) GetNumReleasedValidOk() (*int64, bool)`

GetNumReleasedValidOk returns a tuple with the NumReleasedValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumReleasedValid

`func (o *CloudLdapConnectionPoolStatistics) SetNumReleasedValid(v int64)`

SetNumReleasedValid sets NumReleasedValid field to given value.

### HasNumReleasedValid

`func (o *CloudLdapConnectionPoolStatistics) HasNumReleasedValid() bool`

HasNumReleasedValid returns a boolean if a field has been set.

### GetNumSuccessfulCheckouts

`func (o *CloudLdapConnectionPoolStatistics) GetNumSuccessfulCheckouts() int64`

GetNumSuccessfulCheckouts returns the NumSuccessfulCheckouts field if non-nil, zero value otherwise.

### GetNumSuccessfulCheckoutsOk

`func (o *CloudLdapConnectionPoolStatistics) GetNumSuccessfulCheckoutsOk() (*int64, bool)`

GetNumSuccessfulCheckoutsOk returns a tuple with the NumSuccessfulCheckouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSuccessfulCheckouts

`func (o *CloudLdapConnectionPoolStatistics) SetNumSuccessfulCheckouts(v int64)`

SetNumSuccessfulCheckouts sets NumSuccessfulCheckouts field to given value.

### HasNumSuccessfulCheckouts

`func (o *CloudLdapConnectionPoolStatistics) HasNumSuccessfulCheckouts() bool`

HasNumSuccessfulCheckouts returns a boolean if a field has been set.

### GetNumSuccessfulCheckoutsNewConnection

`func (o *CloudLdapConnectionPoolStatistics) GetNumSuccessfulCheckoutsNewConnection() int64`

GetNumSuccessfulCheckoutsNewConnection returns the NumSuccessfulCheckoutsNewConnection field if non-nil, zero value otherwise.

### GetNumSuccessfulCheckoutsNewConnectionOk

`func (o *CloudLdapConnectionPoolStatistics) GetNumSuccessfulCheckoutsNewConnectionOk() (*int64, bool)`

GetNumSuccessfulCheckoutsNewConnectionOk returns a tuple with the NumSuccessfulCheckoutsNewConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSuccessfulCheckoutsNewConnection

`func (o *CloudLdapConnectionPoolStatistics) SetNumSuccessfulCheckoutsNewConnection(v int64)`

SetNumSuccessfulCheckoutsNewConnection sets NumSuccessfulCheckoutsNewConnection field to given value.

### HasNumSuccessfulCheckoutsNewConnection

`func (o *CloudLdapConnectionPoolStatistics) HasNumSuccessfulCheckoutsNewConnection() bool`

HasNumSuccessfulCheckoutsNewConnection returns a boolean if a field has been set.

### GetNumSuccessfulConnectionAttempts

`func (o *CloudLdapConnectionPoolStatistics) GetNumSuccessfulConnectionAttempts() int64`

GetNumSuccessfulConnectionAttempts returns the NumSuccessfulConnectionAttempts field if non-nil, zero value otherwise.

### GetNumSuccessfulConnectionAttemptsOk

`func (o *CloudLdapConnectionPoolStatistics) GetNumSuccessfulConnectionAttemptsOk() (*int64, bool)`

GetNumSuccessfulConnectionAttemptsOk returns a tuple with the NumSuccessfulConnectionAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSuccessfulConnectionAttempts

`func (o *CloudLdapConnectionPoolStatistics) SetNumSuccessfulConnectionAttempts(v int64)`

SetNumSuccessfulConnectionAttempts sets NumSuccessfulConnectionAttempts field to given value.

### HasNumSuccessfulConnectionAttempts

`func (o *CloudLdapConnectionPoolStatistics) HasNumSuccessfulConnectionAttempts() bool`

HasNumSuccessfulConnectionAttempts returns a boolean if a field has been set.

### GetMaximumAvailableConnections

`func (o *CloudLdapConnectionPoolStatistics) GetMaximumAvailableConnections() int64`

GetMaximumAvailableConnections returns the MaximumAvailableConnections field if non-nil, zero value otherwise.

### GetMaximumAvailableConnectionsOk

`func (o *CloudLdapConnectionPoolStatistics) GetMaximumAvailableConnectionsOk() (*int64, bool)`

GetMaximumAvailableConnectionsOk returns a tuple with the MaximumAvailableConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumAvailableConnections

`func (o *CloudLdapConnectionPoolStatistics) SetMaximumAvailableConnections(v int64)`

SetMaximumAvailableConnections sets MaximumAvailableConnections field to given value.

### HasMaximumAvailableConnections

`func (o *CloudLdapConnectionPoolStatistics) HasMaximumAvailableConnections() bool`

HasMaximumAvailableConnections returns a boolean if a field has been set.

### GetNumSuccessfulCheckoutsWithoutWaiting

`func (o *CloudLdapConnectionPoolStatistics) GetNumSuccessfulCheckoutsWithoutWaiting() int64`

GetNumSuccessfulCheckoutsWithoutWaiting returns the NumSuccessfulCheckoutsWithoutWaiting field if non-nil, zero value otherwise.

### GetNumSuccessfulCheckoutsWithoutWaitingOk

`func (o *CloudLdapConnectionPoolStatistics) GetNumSuccessfulCheckoutsWithoutWaitingOk() (*int64, bool)`

GetNumSuccessfulCheckoutsWithoutWaitingOk returns a tuple with the NumSuccessfulCheckoutsWithoutWaiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSuccessfulCheckoutsWithoutWaiting

`func (o *CloudLdapConnectionPoolStatistics) SetNumSuccessfulCheckoutsWithoutWaiting(v int64)`

SetNumSuccessfulCheckoutsWithoutWaiting sets NumSuccessfulCheckoutsWithoutWaiting field to given value.

### HasNumSuccessfulCheckoutsWithoutWaiting

`func (o *CloudLdapConnectionPoolStatistics) HasNumSuccessfulCheckoutsWithoutWaiting() bool`

HasNumSuccessfulCheckoutsWithoutWaiting returns a boolean if a field has been set.

### GetNumSuccessfulCheckoutsAfterWaiting

`func (o *CloudLdapConnectionPoolStatistics) GetNumSuccessfulCheckoutsAfterWaiting() int64`

GetNumSuccessfulCheckoutsAfterWaiting returns the NumSuccessfulCheckoutsAfterWaiting field if non-nil, zero value otherwise.

### GetNumSuccessfulCheckoutsAfterWaitingOk

`func (o *CloudLdapConnectionPoolStatistics) GetNumSuccessfulCheckoutsAfterWaitingOk() (*int64, bool)`

GetNumSuccessfulCheckoutsAfterWaitingOk returns a tuple with the NumSuccessfulCheckoutsAfterWaiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSuccessfulCheckoutsAfterWaiting

`func (o *CloudLdapConnectionPoolStatistics) SetNumSuccessfulCheckoutsAfterWaiting(v int64)`

SetNumSuccessfulCheckoutsAfterWaiting sets NumSuccessfulCheckoutsAfterWaiting field to given value.

### HasNumSuccessfulCheckoutsAfterWaiting

`func (o *CloudLdapConnectionPoolStatistics) HasNumSuccessfulCheckoutsAfterWaiting() bool`

HasNumSuccessfulCheckoutsAfterWaiting returns a boolean if a field has been set.

### GetNumAvailableConnections

`func (o *CloudLdapConnectionPoolStatistics) GetNumAvailableConnections() int64`

GetNumAvailableConnections returns the NumAvailableConnections field if non-nil, zero value otherwise.

### GetNumAvailableConnectionsOk

`func (o *CloudLdapConnectionPoolStatistics) GetNumAvailableConnectionsOk() (*int64, bool)`

GetNumAvailableConnectionsOk returns a tuple with the NumAvailableConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAvailableConnections

`func (o *CloudLdapConnectionPoolStatistics) SetNumAvailableConnections(v int64)`

SetNumAvailableConnections sets NumAvailableConnections field to given value.

### HasNumAvailableConnections

`func (o *CloudLdapConnectionPoolStatistics) HasNumAvailableConnections() bool`

HasNumAvailableConnections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


