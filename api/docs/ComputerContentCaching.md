# ComputerContentCaching

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComputerContentCachingInformationId** | Pointer to **string** |  | [optional] [readonly] 
**Parents** | Pointer to [**[]ComputerContentCachingParent**](ComputerContentCachingParent.md) |  | [optional] [readonly] 
**Alerts** | Pointer to [**[]ComputerContentCachingAlert**](ComputerContentCachingAlert.md) |  | [optional] [readonly] 
**Activated** | Pointer to **bool** |  | [optional] [readonly] 
**Active** | Pointer to **bool** |  | [optional] [readonly] 
**ActualCacheBytesUsed** | Pointer to **int64** |  | [optional] [readonly] 
**CacheDetails** | Pointer to [**[]ComputerContentCachingCacheDetail**](ComputerContentCachingCacheDetail.md) |  | [optional] [readonly] 
**CacheBytesFree** | Pointer to **int64** |  | [optional] [readonly] 
**CacheBytesLimit** | Pointer to **int64** |  | [optional] [readonly] 
**CacheStatus** | Pointer to **string** |  | [optional] [readonly] 
**CacheBytesUsed** | Pointer to **int64** |  | [optional] [readonly] 
**DataMigrationCompleted** | Pointer to **bool** |  | [optional] [readonly] 
**DataMigrationProgressPercentage** | Pointer to **int64** |  | [optional] [readonly] 
**DataMigrationError** | Pointer to [**ComputerContentCachingDataMigrationError**](ComputerContentCachingDataMigrationError.md) |  | [optional] 
**MaxCachePressureLast1HourPercentage** | Pointer to **int64** |  | [optional] [readonly] 
**PersonalCacheBytesFree** | Pointer to **int64** |  | [optional] [readonly] 
**PersonalCacheBytesLimit** | Pointer to **int64** |  | [optional] [readonly] 
**PersonalCacheBytesUsed** | Pointer to **int64** |  | [optional] [readonly] 
**Port** | Pointer to **int64** |  | [optional] [readonly] 
**PublicAddress** | Pointer to **string** |  | [optional] [readonly] 
**RegistrationError** | Pointer to **string** |  | [optional] [readonly] 
**RegistrationResponseCode** | Pointer to **int64** |  | [optional] [readonly] 
**RegistrationStarted** | Pointer to **time.Time** |  | [optional] [readonly] 
**RegistrationStatus** | Pointer to **string** |  | [optional] [readonly] 
**RestrictedMedia** | Pointer to **bool** |  | [optional] [readonly] 
**ServerGuid** | Pointer to **string** |  | [optional] [readonly] 
**StartupStatus** | Pointer to **string** |  | [optional] [readonly] 
**TetheratorStatus** | Pointer to **string** |  | [optional] [readonly] 
**TotalBytesAreSince** | Pointer to **time.Time** |  | [optional] [readonly] 
**TotalBytesDropped** | Pointer to **int64** |  | [optional] [readonly] 
**TotalBytesImported** | Pointer to **int64** |  | [optional] [readonly] 
**TotalBytesReturnedToChildren** | Pointer to **int64** |  | [optional] [readonly] 
**TotalBytesReturnedToClients** | Pointer to **int64** |  | [optional] [readonly] 
**TotalBytesReturnedToPeers** | Pointer to **int64** |  | [optional] [readonly] 
**TotalBytesStoredFromOrigin** | Pointer to **int64** |  | [optional] [readonly] 
**TotalBytesStoredFromParents** | Pointer to **int64** |  | [optional] [readonly] 
**TotalBytesStoredFromPeers** | Pointer to **int64** |  | [optional] [readonly] 

## Methods

### NewComputerContentCaching

`func NewComputerContentCaching() *ComputerContentCaching`

NewComputerContentCaching instantiates a new ComputerContentCaching object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputerContentCachingWithDefaults

`func NewComputerContentCachingWithDefaults() *ComputerContentCaching`

NewComputerContentCachingWithDefaults instantiates a new ComputerContentCaching object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComputerContentCachingInformationId

`func (o *ComputerContentCaching) GetComputerContentCachingInformationId() string`

GetComputerContentCachingInformationId returns the ComputerContentCachingInformationId field if non-nil, zero value otherwise.

### GetComputerContentCachingInformationIdOk

`func (o *ComputerContentCaching) GetComputerContentCachingInformationIdOk() (*string, bool)`

GetComputerContentCachingInformationIdOk returns a tuple with the ComputerContentCachingInformationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputerContentCachingInformationId

`func (o *ComputerContentCaching) SetComputerContentCachingInformationId(v string)`

SetComputerContentCachingInformationId sets ComputerContentCachingInformationId field to given value.

### HasComputerContentCachingInformationId

`func (o *ComputerContentCaching) HasComputerContentCachingInformationId() bool`

HasComputerContentCachingInformationId returns a boolean if a field has been set.

### GetParents

`func (o *ComputerContentCaching) GetParents() []ComputerContentCachingParent`

GetParents returns the Parents field if non-nil, zero value otherwise.

### GetParentsOk

`func (o *ComputerContentCaching) GetParentsOk() (*[]ComputerContentCachingParent, bool)`

GetParentsOk returns a tuple with the Parents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParents

`func (o *ComputerContentCaching) SetParents(v []ComputerContentCachingParent)`

SetParents sets Parents field to given value.

### HasParents

`func (o *ComputerContentCaching) HasParents() bool`

HasParents returns a boolean if a field has been set.

### GetAlerts

`func (o *ComputerContentCaching) GetAlerts() []ComputerContentCachingAlert`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *ComputerContentCaching) GetAlertsOk() (*[]ComputerContentCachingAlert, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *ComputerContentCaching) SetAlerts(v []ComputerContentCachingAlert)`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *ComputerContentCaching) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.

### GetActivated

`func (o *ComputerContentCaching) GetActivated() bool`

GetActivated returns the Activated field if non-nil, zero value otherwise.

### GetActivatedOk

`func (o *ComputerContentCaching) GetActivatedOk() (*bool, bool)`

GetActivatedOk returns a tuple with the Activated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivated

`func (o *ComputerContentCaching) SetActivated(v bool)`

SetActivated sets Activated field to given value.

### HasActivated

`func (o *ComputerContentCaching) HasActivated() bool`

HasActivated returns a boolean if a field has been set.

### GetActive

`func (o *ComputerContentCaching) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ComputerContentCaching) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ComputerContentCaching) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ComputerContentCaching) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetActualCacheBytesUsed

`func (o *ComputerContentCaching) GetActualCacheBytesUsed() int64`

GetActualCacheBytesUsed returns the ActualCacheBytesUsed field if non-nil, zero value otherwise.

### GetActualCacheBytesUsedOk

`func (o *ComputerContentCaching) GetActualCacheBytesUsedOk() (*int64, bool)`

GetActualCacheBytesUsedOk returns a tuple with the ActualCacheBytesUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualCacheBytesUsed

`func (o *ComputerContentCaching) SetActualCacheBytesUsed(v int64)`

SetActualCacheBytesUsed sets ActualCacheBytesUsed field to given value.

### HasActualCacheBytesUsed

`func (o *ComputerContentCaching) HasActualCacheBytesUsed() bool`

HasActualCacheBytesUsed returns a boolean if a field has been set.

### GetCacheDetails

`func (o *ComputerContentCaching) GetCacheDetails() []ComputerContentCachingCacheDetail`

GetCacheDetails returns the CacheDetails field if non-nil, zero value otherwise.

### GetCacheDetailsOk

`func (o *ComputerContentCaching) GetCacheDetailsOk() (*[]ComputerContentCachingCacheDetail, bool)`

GetCacheDetailsOk returns a tuple with the CacheDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheDetails

`func (o *ComputerContentCaching) SetCacheDetails(v []ComputerContentCachingCacheDetail)`

SetCacheDetails sets CacheDetails field to given value.

### HasCacheDetails

`func (o *ComputerContentCaching) HasCacheDetails() bool`

HasCacheDetails returns a boolean if a field has been set.

### GetCacheBytesFree

`func (o *ComputerContentCaching) GetCacheBytesFree() int64`

GetCacheBytesFree returns the CacheBytesFree field if non-nil, zero value otherwise.

### GetCacheBytesFreeOk

`func (o *ComputerContentCaching) GetCacheBytesFreeOk() (*int64, bool)`

GetCacheBytesFreeOk returns a tuple with the CacheBytesFree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheBytesFree

`func (o *ComputerContentCaching) SetCacheBytesFree(v int64)`

SetCacheBytesFree sets CacheBytesFree field to given value.

### HasCacheBytesFree

`func (o *ComputerContentCaching) HasCacheBytesFree() bool`

HasCacheBytesFree returns a boolean if a field has been set.

### GetCacheBytesLimit

`func (o *ComputerContentCaching) GetCacheBytesLimit() int64`

GetCacheBytesLimit returns the CacheBytesLimit field if non-nil, zero value otherwise.

### GetCacheBytesLimitOk

`func (o *ComputerContentCaching) GetCacheBytesLimitOk() (*int64, bool)`

GetCacheBytesLimitOk returns a tuple with the CacheBytesLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheBytesLimit

`func (o *ComputerContentCaching) SetCacheBytesLimit(v int64)`

SetCacheBytesLimit sets CacheBytesLimit field to given value.

### HasCacheBytesLimit

`func (o *ComputerContentCaching) HasCacheBytesLimit() bool`

HasCacheBytesLimit returns a boolean if a field has been set.

### GetCacheStatus

`func (o *ComputerContentCaching) GetCacheStatus() string`

GetCacheStatus returns the CacheStatus field if non-nil, zero value otherwise.

### GetCacheStatusOk

`func (o *ComputerContentCaching) GetCacheStatusOk() (*string, bool)`

GetCacheStatusOk returns a tuple with the CacheStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheStatus

`func (o *ComputerContentCaching) SetCacheStatus(v string)`

SetCacheStatus sets CacheStatus field to given value.

### HasCacheStatus

`func (o *ComputerContentCaching) HasCacheStatus() bool`

HasCacheStatus returns a boolean if a field has been set.

### GetCacheBytesUsed

`func (o *ComputerContentCaching) GetCacheBytesUsed() int64`

GetCacheBytesUsed returns the CacheBytesUsed field if non-nil, zero value otherwise.

### GetCacheBytesUsedOk

`func (o *ComputerContentCaching) GetCacheBytesUsedOk() (*int64, bool)`

GetCacheBytesUsedOk returns a tuple with the CacheBytesUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheBytesUsed

`func (o *ComputerContentCaching) SetCacheBytesUsed(v int64)`

SetCacheBytesUsed sets CacheBytesUsed field to given value.

### HasCacheBytesUsed

`func (o *ComputerContentCaching) HasCacheBytesUsed() bool`

HasCacheBytesUsed returns a boolean if a field has been set.

### GetDataMigrationCompleted

`func (o *ComputerContentCaching) GetDataMigrationCompleted() bool`

GetDataMigrationCompleted returns the DataMigrationCompleted field if non-nil, zero value otherwise.

### GetDataMigrationCompletedOk

`func (o *ComputerContentCaching) GetDataMigrationCompletedOk() (*bool, bool)`

GetDataMigrationCompletedOk returns a tuple with the DataMigrationCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataMigrationCompleted

`func (o *ComputerContentCaching) SetDataMigrationCompleted(v bool)`

SetDataMigrationCompleted sets DataMigrationCompleted field to given value.

### HasDataMigrationCompleted

`func (o *ComputerContentCaching) HasDataMigrationCompleted() bool`

HasDataMigrationCompleted returns a boolean if a field has been set.

### GetDataMigrationProgressPercentage

`func (o *ComputerContentCaching) GetDataMigrationProgressPercentage() int64`

GetDataMigrationProgressPercentage returns the DataMigrationProgressPercentage field if non-nil, zero value otherwise.

### GetDataMigrationProgressPercentageOk

`func (o *ComputerContentCaching) GetDataMigrationProgressPercentageOk() (*int64, bool)`

GetDataMigrationProgressPercentageOk returns a tuple with the DataMigrationProgressPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataMigrationProgressPercentage

`func (o *ComputerContentCaching) SetDataMigrationProgressPercentage(v int64)`

SetDataMigrationProgressPercentage sets DataMigrationProgressPercentage field to given value.

### HasDataMigrationProgressPercentage

`func (o *ComputerContentCaching) HasDataMigrationProgressPercentage() bool`

HasDataMigrationProgressPercentage returns a boolean if a field has been set.

### GetDataMigrationError

`func (o *ComputerContentCaching) GetDataMigrationError() ComputerContentCachingDataMigrationError`

GetDataMigrationError returns the DataMigrationError field if non-nil, zero value otherwise.

### GetDataMigrationErrorOk

`func (o *ComputerContentCaching) GetDataMigrationErrorOk() (*ComputerContentCachingDataMigrationError, bool)`

GetDataMigrationErrorOk returns a tuple with the DataMigrationError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataMigrationError

`func (o *ComputerContentCaching) SetDataMigrationError(v ComputerContentCachingDataMigrationError)`

SetDataMigrationError sets DataMigrationError field to given value.

### HasDataMigrationError

`func (o *ComputerContentCaching) HasDataMigrationError() bool`

HasDataMigrationError returns a boolean if a field has been set.

### GetMaxCachePressureLast1HourPercentage

`func (o *ComputerContentCaching) GetMaxCachePressureLast1HourPercentage() int64`

GetMaxCachePressureLast1HourPercentage returns the MaxCachePressureLast1HourPercentage field if non-nil, zero value otherwise.

### GetMaxCachePressureLast1HourPercentageOk

`func (o *ComputerContentCaching) GetMaxCachePressureLast1HourPercentageOk() (*int64, bool)`

GetMaxCachePressureLast1HourPercentageOk returns a tuple with the MaxCachePressureLast1HourPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCachePressureLast1HourPercentage

`func (o *ComputerContentCaching) SetMaxCachePressureLast1HourPercentage(v int64)`

SetMaxCachePressureLast1HourPercentage sets MaxCachePressureLast1HourPercentage field to given value.

### HasMaxCachePressureLast1HourPercentage

`func (o *ComputerContentCaching) HasMaxCachePressureLast1HourPercentage() bool`

HasMaxCachePressureLast1HourPercentage returns a boolean if a field has been set.

### GetPersonalCacheBytesFree

`func (o *ComputerContentCaching) GetPersonalCacheBytesFree() int64`

GetPersonalCacheBytesFree returns the PersonalCacheBytesFree field if non-nil, zero value otherwise.

### GetPersonalCacheBytesFreeOk

`func (o *ComputerContentCaching) GetPersonalCacheBytesFreeOk() (*int64, bool)`

GetPersonalCacheBytesFreeOk returns a tuple with the PersonalCacheBytesFree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalCacheBytesFree

`func (o *ComputerContentCaching) SetPersonalCacheBytesFree(v int64)`

SetPersonalCacheBytesFree sets PersonalCacheBytesFree field to given value.

### HasPersonalCacheBytesFree

`func (o *ComputerContentCaching) HasPersonalCacheBytesFree() bool`

HasPersonalCacheBytesFree returns a boolean if a field has been set.

### GetPersonalCacheBytesLimit

`func (o *ComputerContentCaching) GetPersonalCacheBytesLimit() int64`

GetPersonalCacheBytesLimit returns the PersonalCacheBytesLimit field if non-nil, zero value otherwise.

### GetPersonalCacheBytesLimitOk

`func (o *ComputerContentCaching) GetPersonalCacheBytesLimitOk() (*int64, bool)`

GetPersonalCacheBytesLimitOk returns a tuple with the PersonalCacheBytesLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalCacheBytesLimit

`func (o *ComputerContentCaching) SetPersonalCacheBytesLimit(v int64)`

SetPersonalCacheBytesLimit sets PersonalCacheBytesLimit field to given value.

### HasPersonalCacheBytesLimit

`func (o *ComputerContentCaching) HasPersonalCacheBytesLimit() bool`

HasPersonalCacheBytesLimit returns a boolean if a field has been set.

### GetPersonalCacheBytesUsed

`func (o *ComputerContentCaching) GetPersonalCacheBytesUsed() int64`

GetPersonalCacheBytesUsed returns the PersonalCacheBytesUsed field if non-nil, zero value otherwise.

### GetPersonalCacheBytesUsedOk

`func (o *ComputerContentCaching) GetPersonalCacheBytesUsedOk() (*int64, bool)`

GetPersonalCacheBytesUsedOk returns a tuple with the PersonalCacheBytesUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalCacheBytesUsed

`func (o *ComputerContentCaching) SetPersonalCacheBytesUsed(v int64)`

SetPersonalCacheBytesUsed sets PersonalCacheBytesUsed field to given value.

### HasPersonalCacheBytesUsed

`func (o *ComputerContentCaching) HasPersonalCacheBytesUsed() bool`

HasPersonalCacheBytesUsed returns a boolean if a field has been set.

### GetPort

`func (o *ComputerContentCaching) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ComputerContentCaching) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ComputerContentCaching) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *ComputerContentCaching) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPublicAddress

`func (o *ComputerContentCaching) GetPublicAddress() string`

GetPublicAddress returns the PublicAddress field if non-nil, zero value otherwise.

### GetPublicAddressOk

`func (o *ComputerContentCaching) GetPublicAddressOk() (*string, bool)`

GetPublicAddressOk returns a tuple with the PublicAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicAddress

`func (o *ComputerContentCaching) SetPublicAddress(v string)`

SetPublicAddress sets PublicAddress field to given value.

### HasPublicAddress

`func (o *ComputerContentCaching) HasPublicAddress() bool`

HasPublicAddress returns a boolean if a field has been set.

### GetRegistrationError

`func (o *ComputerContentCaching) GetRegistrationError() string`

GetRegistrationError returns the RegistrationError field if non-nil, zero value otherwise.

### GetRegistrationErrorOk

`func (o *ComputerContentCaching) GetRegistrationErrorOk() (*string, bool)`

GetRegistrationErrorOk returns a tuple with the RegistrationError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationError

`func (o *ComputerContentCaching) SetRegistrationError(v string)`

SetRegistrationError sets RegistrationError field to given value.

### HasRegistrationError

`func (o *ComputerContentCaching) HasRegistrationError() bool`

HasRegistrationError returns a boolean if a field has been set.

### GetRegistrationResponseCode

`func (o *ComputerContentCaching) GetRegistrationResponseCode() int64`

GetRegistrationResponseCode returns the RegistrationResponseCode field if non-nil, zero value otherwise.

### GetRegistrationResponseCodeOk

`func (o *ComputerContentCaching) GetRegistrationResponseCodeOk() (*int64, bool)`

GetRegistrationResponseCodeOk returns a tuple with the RegistrationResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationResponseCode

`func (o *ComputerContentCaching) SetRegistrationResponseCode(v int64)`

SetRegistrationResponseCode sets RegistrationResponseCode field to given value.

### HasRegistrationResponseCode

`func (o *ComputerContentCaching) HasRegistrationResponseCode() bool`

HasRegistrationResponseCode returns a boolean if a field has been set.

### GetRegistrationStarted

`func (o *ComputerContentCaching) GetRegistrationStarted() time.Time`

GetRegistrationStarted returns the RegistrationStarted field if non-nil, zero value otherwise.

### GetRegistrationStartedOk

`func (o *ComputerContentCaching) GetRegistrationStartedOk() (*time.Time, bool)`

GetRegistrationStartedOk returns a tuple with the RegistrationStarted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationStarted

`func (o *ComputerContentCaching) SetRegistrationStarted(v time.Time)`

SetRegistrationStarted sets RegistrationStarted field to given value.

### HasRegistrationStarted

`func (o *ComputerContentCaching) HasRegistrationStarted() bool`

HasRegistrationStarted returns a boolean if a field has been set.

### GetRegistrationStatus

`func (o *ComputerContentCaching) GetRegistrationStatus() string`

GetRegistrationStatus returns the RegistrationStatus field if non-nil, zero value otherwise.

### GetRegistrationStatusOk

`func (o *ComputerContentCaching) GetRegistrationStatusOk() (*string, bool)`

GetRegistrationStatusOk returns a tuple with the RegistrationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationStatus

`func (o *ComputerContentCaching) SetRegistrationStatus(v string)`

SetRegistrationStatus sets RegistrationStatus field to given value.

### HasRegistrationStatus

`func (o *ComputerContentCaching) HasRegistrationStatus() bool`

HasRegistrationStatus returns a boolean if a field has been set.

### GetRestrictedMedia

`func (o *ComputerContentCaching) GetRestrictedMedia() bool`

GetRestrictedMedia returns the RestrictedMedia field if non-nil, zero value otherwise.

### GetRestrictedMediaOk

`func (o *ComputerContentCaching) GetRestrictedMediaOk() (*bool, bool)`

GetRestrictedMediaOk returns a tuple with the RestrictedMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedMedia

`func (o *ComputerContentCaching) SetRestrictedMedia(v bool)`

SetRestrictedMedia sets RestrictedMedia field to given value.

### HasRestrictedMedia

`func (o *ComputerContentCaching) HasRestrictedMedia() bool`

HasRestrictedMedia returns a boolean if a field has been set.

### GetServerGuid

`func (o *ComputerContentCaching) GetServerGuid() string`

GetServerGuid returns the ServerGuid field if non-nil, zero value otherwise.

### GetServerGuidOk

`func (o *ComputerContentCaching) GetServerGuidOk() (*string, bool)`

GetServerGuidOk returns a tuple with the ServerGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGuid

`func (o *ComputerContentCaching) SetServerGuid(v string)`

SetServerGuid sets ServerGuid field to given value.

### HasServerGuid

`func (o *ComputerContentCaching) HasServerGuid() bool`

HasServerGuid returns a boolean if a field has been set.

### GetStartupStatus

`func (o *ComputerContentCaching) GetStartupStatus() string`

GetStartupStatus returns the StartupStatus field if non-nil, zero value otherwise.

### GetStartupStatusOk

`func (o *ComputerContentCaching) GetStartupStatusOk() (*string, bool)`

GetStartupStatusOk returns a tuple with the StartupStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartupStatus

`func (o *ComputerContentCaching) SetStartupStatus(v string)`

SetStartupStatus sets StartupStatus field to given value.

### HasStartupStatus

`func (o *ComputerContentCaching) HasStartupStatus() bool`

HasStartupStatus returns a boolean if a field has been set.

### GetTetheratorStatus

`func (o *ComputerContentCaching) GetTetheratorStatus() string`

GetTetheratorStatus returns the TetheratorStatus field if non-nil, zero value otherwise.

### GetTetheratorStatusOk

`func (o *ComputerContentCaching) GetTetheratorStatusOk() (*string, bool)`

GetTetheratorStatusOk returns a tuple with the TetheratorStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTetheratorStatus

`func (o *ComputerContentCaching) SetTetheratorStatus(v string)`

SetTetheratorStatus sets TetheratorStatus field to given value.

### HasTetheratorStatus

`func (o *ComputerContentCaching) HasTetheratorStatus() bool`

HasTetheratorStatus returns a boolean if a field has been set.

### GetTotalBytesAreSince

`func (o *ComputerContentCaching) GetTotalBytesAreSince() time.Time`

GetTotalBytesAreSince returns the TotalBytesAreSince field if non-nil, zero value otherwise.

### GetTotalBytesAreSinceOk

`func (o *ComputerContentCaching) GetTotalBytesAreSinceOk() (*time.Time, bool)`

GetTotalBytesAreSinceOk returns a tuple with the TotalBytesAreSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBytesAreSince

`func (o *ComputerContentCaching) SetTotalBytesAreSince(v time.Time)`

SetTotalBytesAreSince sets TotalBytesAreSince field to given value.

### HasTotalBytesAreSince

`func (o *ComputerContentCaching) HasTotalBytesAreSince() bool`

HasTotalBytesAreSince returns a boolean if a field has been set.

### GetTotalBytesDropped

`func (o *ComputerContentCaching) GetTotalBytesDropped() int64`

GetTotalBytesDropped returns the TotalBytesDropped field if non-nil, zero value otherwise.

### GetTotalBytesDroppedOk

`func (o *ComputerContentCaching) GetTotalBytesDroppedOk() (*int64, bool)`

GetTotalBytesDroppedOk returns a tuple with the TotalBytesDropped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBytesDropped

`func (o *ComputerContentCaching) SetTotalBytesDropped(v int64)`

SetTotalBytesDropped sets TotalBytesDropped field to given value.

### HasTotalBytesDropped

`func (o *ComputerContentCaching) HasTotalBytesDropped() bool`

HasTotalBytesDropped returns a boolean if a field has been set.

### GetTotalBytesImported

`func (o *ComputerContentCaching) GetTotalBytesImported() int64`

GetTotalBytesImported returns the TotalBytesImported field if non-nil, zero value otherwise.

### GetTotalBytesImportedOk

`func (o *ComputerContentCaching) GetTotalBytesImportedOk() (*int64, bool)`

GetTotalBytesImportedOk returns a tuple with the TotalBytesImported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBytesImported

`func (o *ComputerContentCaching) SetTotalBytesImported(v int64)`

SetTotalBytesImported sets TotalBytesImported field to given value.

### HasTotalBytesImported

`func (o *ComputerContentCaching) HasTotalBytesImported() bool`

HasTotalBytesImported returns a boolean if a field has been set.

### GetTotalBytesReturnedToChildren

`func (o *ComputerContentCaching) GetTotalBytesReturnedToChildren() int64`

GetTotalBytesReturnedToChildren returns the TotalBytesReturnedToChildren field if non-nil, zero value otherwise.

### GetTotalBytesReturnedToChildrenOk

`func (o *ComputerContentCaching) GetTotalBytesReturnedToChildrenOk() (*int64, bool)`

GetTotalBytesReturnedToChildrenOk returns a tuple with the TotalBytesReturnedToChildren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBytesReturnedToChildren

`func (o *ComputerContentCaching) SetTotalBytesReturnedToChildren(v int64)`

SetTotalBytesReturnedToChildren sets TotalBytesReturnedToChildren field to given value.

### HasTotalBytesReturnedToChildren

`func (o *ComputerContentCaching) HasTotalBytesReturnedToChildren() bool`

HasTotalBytesReturnedToChildren returns a boolean if a field has been set.

### GetTotalBytesReturnedToClients

`func (o *ComputerContentCaching) GetTotalBytesReturnedToClients() int64`

GetTotalBytesReturnedToClients returns the TotalBytesReturnedToClients field if non-nil, zero value otherwise.

### GetTotalBytesReturnedToClientsOk

`func (o *ComputerContentCaching) GetTotalBytesReturnedToClientsOk() (*int64, bool)`

GetTotalBytesReturnedToClientsOk returns a tuple with the TotalBytesReturnedToClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBytesReturnedToClients

`func (o *ComputerContentCaching) SetTotalBytesReturnedToClients(v int64)`

SetTotalBytesReturnedToClients sets TotalBytesReturnedToClients field to given value.

### HasTotalBytesReturnedToClients

`func (o *ComputerContentCaching) HasTotalBytesReturnedToClients() bool`

HasTotalBytesReturnedToClients returns a boolean if a field has been set.

### GetTotalBytesReturnedToPeers

`func (o *ComputerContentCaching) GetTotalBytesReturnedToPeers() int64`

GetTotalBytesReturnedToPeers returns the TotalBytesReturnedToPeers field if non-nil, zero value otherwise.

### GetTotalBytesReturnedToPeersOk

`func (o *ComputerContentCaching) GetTotalBytesReturnedToPeersOk() (*int64, bool)`

GetTotalBytesReturnedToPeersOk returns a tuple with the TotalBytesReturnedToPeers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBytesReturnedToPeers

`func (o *ComputerContentCaching) SetTotalBytesReturnedToPeers(v int64)`

SetTotalBytesReturnedToPeers sets TotalBytesReturnedToPeers field to given value.

### HasTotalBytesReturnedToPeers

`func (o *ComputerContentCaching) HasTotalBytesReturnedToPeers() bool`

HasTotalBytesReturnedToPeers returns a boolean if a field has been set.

### GetTotalBytesStoredFromOrigin

`func (o *ComputerContentCaching) GetTotalBytesStoredFromOrigin() int64`

GetTotalBytesStoredFromOrigin returns the TotalBytesStoredFromOrigin field if non-nil, zero value otherwise.

### GetTotalBytesStoredFromOriginOk

`func (o *ComputerContentCaching) GetTotalBytesStoredFromOriginOk() (*int64, bool)`

GetTotalBytesStoredFromOriginOk returns a tuple with the TotalBytesStoredFromOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBytesStoredFromOrigin

`func (o *ComputerContentCaching) SetTotalBytesStoredFromOrigin(v int64)`

SetTotalBytesStoredFromOrigin sets TotalBytesStoredFromOrigin field to given value.

### HasTotalBytesStoredFromOrigin

`func (o *ComputerContentCaching) HasTotalBytesStoredFromOrigin() bool`

HasTotalBytesStoredFromOrigin returns a boolean if a field has been set.

### GetTotalBytesStoredFromParents

`func (o *ComputerContentCaching) GetTotalBytesStoredFromParents() int64`

GetTotalBytesStoredFromParents returns the TotalBytesStoredFromParents field if non-nil, zero value otherwise.

### GetTotalBytesStoredFromParentsOk

`func (o *ComputerContentCaching) GetTotalBytesStoredFromParentsOk() (*int64, bool)`

GetTotalBytesStoredFromParentsOk returns a tuple with the TotalBytesStoredFromParents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBytesStoredFromParents

`func (o *ComputerContentCaching) SetTotalBytesStoredFromParents(v int64)`

SetTotalBytesStoredFromParents sets TotalBytesStoredFromParents field to given value.

### HasTotalBytesStoredFromParents

`func (o *ComputerContentCaching) HasTotalBytesStoredFromParents() bool`

HasTotalBytesStoredFromParents returns a boolean if a field has been set.

### GetTotalBytesStoredFromPeers

`func (o *ComputerContentCaching) GetTotalBytesStoredFromPeers() int64`

GetTotalBytesStoredFromPeers returns the TotalBytesStoredFromPeers field if non-nil, zero value otherwise.

### GetTotalBytesStoredFromPeersOk

`func (o *ComputerContentCaching) GetTotalBytesStoredFromPeersOk() (*int64, bool)`

GetTotalBytesStoredFromPeersOk returns a tuple with the TotalBytesStoredFromPeers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBytesStoredFromPeers

`func (o *ComputerContentCaching) SetTotalBytesStoredFromPeers(v int64)`

SetTotalBytesStoredFromPeers sets TotalBytesStoredFromPeers field to given value.

### HasTotalBytesStoredFromPeers

`func (o *ComputerContentCaching) HasTotalBytesStoredFromPeers() bool`

HasTotalBytesStoredFromPeers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


