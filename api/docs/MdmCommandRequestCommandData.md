# MdmCommandRequestCommandData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommandType** | [**MdmCommandType**](MdmCommandType.md) |  | 
**Identifier** | Pointer to **string** | The bundle identifier of the app | [optional] 
**RedemptionCode** | Pointer to **string** | The redemption code that applies to the app pending installation | [optional] 
**UnlockToken** | **string** |  | 
**Data** | Pointer to **string** | Base64 encoded data to be sent with the command | [optional] 
**Queries** | Pointer to **[]string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**Pin** | Pointer to **string** | The six-character PIN for Find My. This value is available in macOS 10.8 and later. | [optional] 
**PreserveDataPlan** | Pointer to **bool** | If true, preserve the data plan on an iPhone or iPad with eSIM functionality, if one exists. This value is available in iOS 11 and later. | [optional] [default to false]
**DisallowProximitySetup** | Pointer to **bool** | If true, disable Proximity Setup on the next reboot and skip the pane in Setup Assistant. This value is available in iOS 11 and later. Prior to iOS 14, don’t use this option with any other option. | [optional] [default to false]
**ObliterationBehavior** | Pointer to **string** | This key defines the fallback behavior for erasing a device. | [optional] 
**ReturnToService** | Pointer to [**ReturnToService**](ReturnToService.md) |  | [optional] 
**Identifiers** | Pointer to **[]string** | Array of application bundle identifiers to query. If not provided, all installed applications are returned. | [optional] 
**ManagedAppsOnly** | Pointer to **NullableBool** | If true, only managed applications are returned. If false or not provided, all applications are returned. | [optional] 
**Items** | Pointer to **[]string** | Array of keys to include in the response for each application. If not provided, default keys are returned. | [optional] 
**UserName** | Pointer to **string** | The username of the user account to unlock | [optional] 
**ForceDeletion** | Pointer to **bool** |  | [optional] 
**DeleteAllUsers** | Pointer to **bool** |  | [optional] 
**LostModeMessage** | Pointer to **string** |  | [optional] 
**LostModePhone** | Pointer to **string** |  | [optional] 
**LostModeFootnote** | Pointer to **string** |  | [optional] 
**ManagedOnly** | Pointer to **NullableBool** | If true, only managed provisioning profiles are returned. If false, all provisioning profiles are returned. The default value is false. | [optional] [default to false]
**EsimServerUrl** | **string** | The URL of the eSIM server that the device should contact to refresh cellular plans. | 
**RebuildKernelCache** | Pointer to **bool** |  | [optional] 
**KextPaths** | Pointer to **[]string** | Only used if RebuildKernelCache is true | [optional] 
**NotifyUser** | Pointer to **bool** |  | [optional] 
**DestinationDeviceId** | Pointer to **string** | The hardware address of the destination device to which the screen will be mirrored. This value isn’t case-sensitive. Required if destinationName is not provided. | [optional] 
**DestinationName** | Pointer to **string** | The name of the destination device to which the screen will be mirrored. Required if destinationDeviceId is not provided. | [optional] 
**Password** | Pointer to **string** | The password to verify. | [optional] 
**ScanTime** | Pointer to **NullableInt64** | The scan time which device spends in seconds to find the destination device. | [optional] 
**Guid** | Pointer to **string** | The unique identifier of the local administrator account. Must match the GUID of an administrator account that MDM created during Device Enrollment Program (DEP) enrollment. | [optional] 
**NewPassword** | Pointer to **string** | The new password for Recovery Lock. Set as an empty string to clear the Recovery Lock password. | [optional] 
**BootstrapTokenAllowed** | Pointer to **NullableBool** |  | [optional] 
**Bluetooth** | Pointer to **NullableBool** |  | [optional] 
**AppAnalytics** | Pointer to [**AppAnalyticsSetting**](AppAnalyticsSetting.md) |  | [optional] 
**DiagnosticSubmission** | Pointer to [**DiagnosticSubmissionSetting**](DiagnosticSubmissionSetting.md) |  | [optional] 
**DataRoaming** | Pointer to [**DataRoamingSetting**](DataRoamingSetting.md) |  | [optional] 
**DefaultApplications** | Pointer to [**DefaultApplications**](DefaultApplications.md) |  | [optional] 
**VoiceRoaming** | Pointer to [**VoiceRoamingSetting**](VoiceRoamingSetting.md) |  | [optional] 
**PersonalHotspot** | Pointer to [**PersonalHotspotSetting**](PersonalHotspotSetting.md) |  | [optional] 
**MaximumResidentUsers** | Pointer to **NullableInt64** |  | [optional] 
**DeviceName** | Pointer to **string** |  | [optional] 
**ApplicationAttributes** | Pointer to [**ApplicationAttributes**](ApplicationAttributes.md) |  | [optional] 
**SharedDeviceConfiguration** | Pointer to [**SharedDeviceConfiguration**](SharedDeviceConfiguration.md) |  | [optional] 
**ApplicationConfiguration** | Pointer to [**ApplicationConfiguration**](ApplicationConfiguration.md) |  | [optional] 
**TimeZone** | Pointer to **NullableString** |  | [optional] 
**SoftwareUpdateSettings** | Pointer to [**SoftwareUpdateSettings**](SoftwareUpdateSettings.md) |  | [optional] 
**PasscodeLockGracePeriod** | Pointer to **NullableInt64** | The number of seconds before a locked screen requires the user to enter the device passcode to unlock it. (Shared iPad Only) | [optional] 
**ActivationLockAllowedWhileSupervised** | Pointer to **bool** | If true, a supervised device registers itself with Activation Lock when the user enables Find My. This setting is available for supervised devices in iOS 7 and later, and macOS 10.15 and later. | [optional] 

## Methods

### NewMdmCommandRequestCommandData

`func NewMdmCommandRequestCommandData(commandType MdmCommandType, unlockToken string, esimServerUrl string, ) *MdmCommandRequestCommandData`

NewMdmCommandRequestCommandData instantiates a new MdmCommandRequestCommandData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMdmCommandRequestCommandDataWithDefaults

`func NewMdmCommandRequestCommandDataWithDefaults() *MdmCommandRequestCommandData`

NewMdmCommandRequestCommandDataWithDefaults instantiates a new MdmCommandRequestCommandData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommandType

`func (o *MdmCommandRequestCommandData) GetCommandType() MdmCommandType`

GetCommandType returns the CommandType field if non-nil, zero value otherwise.

### GetCommandTypeOk

`func (o *MdmCommandRequestCommandData) GetCommandTypeOk() (*MdmCommandType, bool)`

GetCommandTypeOk returns a tuple with the CommandType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandType

`func (o *MdmCommandRequestCommandData) SetCommandType(v MdmCommandType)`

SetCommandType sets CommandType field to given value.


### GetIdentifier

`func (o *MdmCommandRequestCommandData) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *MdmCommandRequestCommandData) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *MdmCommandRequestCommandData) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *MdmCommandRequestCommandData) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetRedemptionCode

`func (o *MdmCommandRequestCommandData) GetRedemptionCode() string`

GetRedemptionCode returns the RedemptionCode field if non-nil, zero value otherwise.

### GetRedemptionCodeOk

`func (o *MdmCommandRequestCommandData) GetRedemptionCodeOk() (*string, bool)`

GetRedemptionCodeOk returns a tuple with the RedemptionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedemptionCode

`func (o *MdmCommandRequestCommandData) SetRedemptionCode(v string)`

SetRedemptionCode sets RedemptionCode field to given value.

### HasRedemptionCode

`func (o *MdmCommandRequestCommandData) HasRedemptionCode() bool`

HasRedemptionCode returns a boolean if a field has been set.

### GetUnlockToken

`func (o *MdmCommandRequestCommandData) GetUnlockToken() string`

GetUnlockToken returns the UnlockToken field if non-nil, zero value otherwise.

### GetUnlockTokenOk

`func (o *MdmCommandRequestCommandData) GetUnlockTokenOk() (*string, bool)`

GetUnlockTokenOk returns a tuple with the UnlockToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlockToken

`func (o *MdmCommandRequestCommandData) SetUnlockToken(v string)`

SetUnlockToken sets UnlockToken field to given value.


### GetData

`func (o *MdmCommandRequestCommandData) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MdmCommandRequestCommandData) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MdmCommandRequestCommandData) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *MdmCommandRequestCommandData) HasData() bool`

HasData returns a boolean if a field has been set.

### GetQueries

`func (o *MdmCommandRequestCommandData) GetQueries() []string`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *MdmCommandRequestCommandData) GetQueriesOk() (*[]string, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *MdmCommandRequestCommandData) SetQueries(v []string)`

SetQueries sets Queries field to given value.

### HasQueries

`func (o *MdmCommandRequestCommandData) HasQueries() bool`

HasQueries returns a boolean if a field has been set.

### GetMessage

`func (o *MdmCommandRequestCommandData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MdmCommandRequestCommandData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MdmCommandRequestCommandData) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *MdmCommandRequestCommandData) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *MdmCommandRequestCommandData) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *MdmCommandRequestCommandData) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *MdmCommandRequestCommandData) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *MdmCommandRequestCommandData) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetPin

`func (o *MdmCommandRequestCommandData) GetPin() string`

GetPin returns the Pin field if non-nil, zero value otherwise.

### GetPinOk

`func (o *MdmCommandRequestCommandData) GetPinOk() (*string, bool)`

GetPinOk returns a tuple with the Pin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPin

`func (o *MdmCommandRequestCommandData) SetPin(v string)`

SetPin sets Pin field to given value.

### HasPin

`func (o *MdmCommandRequestCommandData) HasPin() bool`

HasPin returns a boolean if a field has been set.

### GetPreserveDataPlan

`func (o *MdmCommandRequestCommandData) GetPreserveDataPlan() bool`

GetPreserveDataPlan returns the PreserveDataPlan field if non-nil, zero value otherwise.

### GetPreserveDataPlanOk

`func (o *MdmCommandRequestCommandData) GetPreserveDataPlanOk() (*bool, bool)`

GetPreserveDataPlanOk returns a tuple with the PreserveDataPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveDataPlan

`func (o *MdmCommandRequestCommandData) SetPreserveDataPlan(v bool)`

SetPreserveDataPlan sets PreserveDataPlan field to given value.

### HasPreserveDataPlan

`func (o *MdmCommandRequestCommandData) HasPreserveDataPlan() bool`

HasPreserveDataPlan returns a boolean if a field has been set.

### GetDisallowProximitySetup

`func (o *MdmCommandRequestCommandData) GetDisallowProximitySetup() bool`

GetDisallowProximitySetup returns the DisallowProximitySetup field if non-nil, zero value otherwise.

### GetDisallowProximitySetupOk

`func (o *MdmCommandRequestCommandData) GetDisallowProximitySetupOk() (*bool, bool)`

GetDisallowProximitySetupOk returns a tuple with the DisallowProximitySetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisallowProximitySetup

`func (o *MdmCommandRequestCommandData) SetDisallowProximitySetup(v bool)`

SetDisallowProximitySetup sets DisallowProximitySetup field to given value.

### HasDisallowProximitySetup

`func (o *MdmCommandRequestCommandData) HasDisallowProximitySetup() bool`

HasDisallowProximitySetup returns a boolean if a field has been set.

### GetObliterationBehavior

`func (o *MdmCommandRequestCommandData) GetObliterationBehavior() string`

GetObliterationBehavior returns the ObliterationBehavior field if non-nil, zero value otherwise.

### GetObliterationBehaviorOk

`func (o *MdmCommandRequestCommandData) GetObliterationBehaviorOk() (*string, bool)`

GetObliterationBehaviorOk returns a tuple with the ObliterationBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObliterationBehavior

`func (o *MdmCommandRequestCommandData) SetObliterationBehavior(v string)`

SetObliterationBehavior sets ObliterationBehavior field to given value.

### HasObliterationBehavior

`func (o *MdmCommandRequestCommandData) HasObliterationBehavior() bool`

HasObliterationBehavior returns a boolean if a field has been set.

### GetReturnToService

`func (o *MdmCommandRequestCommandData) GetReturnToService() ReturnToService`

GetReturnToService returns the ReturnToService field if non-nil, zero value otherwise.

### GetReturnToServiceOk

`func (o *MdmCommandRequestCommandData) GetReturnToServiceOk() (*ReturnToService, bool)`

GetReturnToServiceOk returns a tuple with the ReturnToService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnToService

`func (o *MdmCommandRequestCommandData) SetReturnToService(v ReturnToService)`

SetReturnToService sets ReturnToService field to given value.

### HasReturnToService

`func (o *MdmCommandRequestCommandData) HasReturnToService() bool`

HasReturnToService returns a boolean if a field has been set.

### GetIdentifiers

`func (o *MdmCommandRequestCommandData) GetIdentifiers() []string`

GetIdentifiers returns the Identifiers field if non-nil, zero value otherwise.

### GetIdentifiersOk

`func (o *MdmCommandRequestCommandData) GetIdentifiersOk() (*[]string, bool)`

GetIdentifiersOk returns a tuple with the Identifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiers

`func (o *MdmCommandRequestCommandData) SetIdentifiers(v []string)`

SetIdentifiers sets Identifiers field to given value.

### HasIdentifiers

`func (o *MdmCommandRequestCommandData) HasIdentifiers() bool`

HasIdentifiers returns a boolean if a field has been set.

### SetIdentifiersNil

`func (o *MdmCommandRequestCommandData) SetIdentifiersNil(b bool)`

 SetIdentifiersNil sets the value for Identifiers to be an explicit nil

### UnsetIdentifiers
`func (o *MdmCommandRequestCommandData) UnsetIdentifiers()`

UnsetIdentifiers ensures that no value is present for Identifiers, not even an explicit nil
### GetManagedAppsOnly

`func (o *MdmCommandRequestCommandData) GetManagedAppsOnly() bool`

GetManagedAppsOnly returns the ManagedAppsOnly field if non-nil, zero value otherwise.

### GetManagedAppsOnlyOk

`func (o *MdmCommandRequestCommandData) GetManagedAppsOnlyOk() (*bool, bool)`

GetManagedAppsOnlyOk returns a tuple with the ManagedAppsOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedAppsOnly

`func (o *MdmCommandRequestCommandData) SetManagedAppsOnly(v bool)`

SetManagedAppsOnly sets ManagedAppsOnly field to given value.

### HasManagedAppsOnly

`func (o *MdmCommandRequestCommandData) HasManagedAppsOnly() bool`

HasManagedAppsOnly returns a boolean if a field has been set.

### SetManagedAppsOnlyNil

`func (o *MdmCommandRequestCommandData) SetManagedAppsOnlyNil(b bool)`

 SetManagedAppsOnlyNil sets the value for ManagedAppsOnly to be an explicit nil

### UnsetManagedAppsOnly
`func (o *MdmCommandRequestCommandData) UnsetManagedAppsOnly()`

UnsetManagedAppsOnly ensures that no value is present for ManagedAppsOnly, not even an explicit nil
### GetItems

`func (o *MdmCommandRequestCommandData) GetItems() []string`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *MdmCommandRequestCommandData) GetItemsOk() (*[]string, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *MdmCommandRequestCommandData) SetItems(v []string)`

SetItems sets Items field to given value.

### HasItems

`func (o *MdmCommandRequestCommandData) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *MdmCommandRequestCommandData) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *MdmCommandRequestCommandData) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetUserName

`func (o *MdmCommandRequestCommandData) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *MdmCommandRequestCommandData) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *MdmCommandRequestCommandData) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *MdmCommandRequestCommandData) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetForceDeletion

`func (o *MdmCommandRequestCommandData) GetForceDeletion() bool`

GetForceDeletion returns the ForceDeletion field if non-nil, zero value otherwise.

### GetForceDeletionOk

`func (o *MdmCommandRequestCommandData) GetForceDeletionOk() (*bool, bool)`

GetForceDeletionOk returns a tuple with the ForceDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceDeletion

`func (o *MdmCommandRequestCommandData) SetForceDeletion(v bool)`

SetForceDeletion sets ForceDeletion field to given value.

### HasForceDeletion

`func (o *MdmCommandRequestCommandData) HasForceDeletion() bool`

HasForceDeletion returns a boolean if a field has been set.

### GetDeleteAllUsers

`func (o *MdmCommandRequestCommandData) GetDeleteAllUsers() bool`

GetDeleteAllUsers returns the DeleteAllUsers field if non-nil, zero value otherwise.

### GetDeleteAllUsersOk

`func (o *MdmCommandRequestCommandData) GetDeleteAllUsersOk() (*bool, bool)`

GetDeleteAllUsersOk returns a tuple with the DeleteAllUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteAllUsers

`func (o *MdmCommandRequestCommandData) SetDeleteAllUsers(v bool)`

SetDeleteAllUsers sets DeleteAllUsers field to given value.

### HasDeleteAllUsers

`func (o *MdmCommandRequestCommandData) HasDeleteAllUsers() bool`

HasDeleteAllUsers returns a boolean if a field has been set.

### GetLostModeMessage

`func (o *MdmCommandRequestCommandData) GetLostModeMessage() string`

GetLostModeMessage returns the LostModeMessage field if non-nil, zero value otherwise.

### GetLostModeMessageOk

`func (o *MdmCommandRequestCommandData) GetLostModeMessageOk() (*string, bool)`

GetLostModeMessageOk returns a tuple with the LostModeMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLostModeMessage

`func (o *MdmCommandRequestCommandData) SetLostModeMessage(v string)`

SetLostModeMessage sets LostModeMessage field to given value.

### HasLostModeMessage

`func (o *MdmCommandRequestCommandData) HasLostModeMessage() bool`

HasLostModeMessage returns a boolean if a field has been set.

### GetLostModePhone

`func (o *MdmCommandRequestCommandData) GetLostModePhone() string`

GetLostModePhone returns the LostModePhone field if non-nil, zero value otherwise.

### GetLostModePhoneOk

`func (o *MdmCommandRequestCommandData) GetLostModePhoneOk() (*string, bool)`

GetLostModePhoneOk returns a tuple with the LostModePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLostModePhone

`func (o *MdmCommandRequestCommandData) SetLostModePhone(v string)`

SetLostModePhone sets LostModePhone field to given value.

### HasLostModePhone

`func (o *MdmCommandRequestCommandData) HasLostModePhone() bool`

HasLostModePhone returns a boolean if a field has been set.

### GetLostModeFootnote

`func (o *MdmCommandRequestCommandData) GetLostModeFootnote() string`

GetLostModeFootnote returns the LostModeFootnote field if non-nil, zero value otherwise.

### GetLostModeFootnoteOk

`func (o *MdmCommandRequestCommandData) GetLostModeFootnoteOk() (*string, bool)`

GetLostModeFootnoteOk returns a tuple with the LostModeFootnote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLostModeFootnote

`func (o *MdmCommandRequestCommandData) SetLostModeFootnote(v string)`

SetLostModeFootnote sets LostModeFootnote field to given value.

### HasLostModeFootnote

`func (o *MdmCommandRequestCommandData) HasLostModeFootnote() bool`

HasLostModeFootnote returns a boolean if a field has been set.

### GetManagedOnly

`func (o *MdmCommandRequestCommandData) GetManagedOnly() bool`

GetManagedOnly returns the ManagedOnly field if non-nil, zero value otherwise.

### GetManagedOnlyOk

`func (o *MdmCommandRequestCommandData) GetManagedOnlyOk() (*bool, bool)`

GetManagedOnlyOk returns a tuple with the ManagedOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedOnly

`func (o *MdmCommandRequestCommandData) SetManagedOnly(v bool)`

SetManagedOnly sets ManagedOnly field to given value.

### HasManagedOnly

`func (o *MdmCommandRequestCommandData) HasManagedOnly() bool`

HasManagedOnly returns a boolean if a field has been set.

### SetManagedOnlyNil

`func (o *MdmCommandRequestCommandData) SetManagedOnlyNil(b bool)`

 SetManagedOnlyNil sets the value for ManagedOnly to be an explicit nil

### UnsetManagedOnly
`func (o *MdmCommandRequestCommandData) UnsetManagedOnly()`

UnsetManagedOnly ensures that no value is present for ManagedOnly, not even an explicit nil
### GetEsimServerUrl

`func (o *MdmCommandRequestCommandData) GetEsimServerUrl() string`

GetEsimServerUrl returns the EsimServerUrl field if non-nil, zero value otherwise.

### GetEsimServerUrlOk

`func (o *MdmCommandRequestCommandData) GetEsimServerUrlOk() (*string, bool)`

GetEsimServerUrlOk returns a tuple with the EsimServerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsimServerUrl

`func (o *MdmCommandRequestCommandData) SetEsimServerUrl(v string)`

SetEsimServerUrl sets EsimServerUrl field to given value.


### GetRebuildKernelCache

`func (o *MdmCommandRequestCommandData) GetRebuildKernelCache() bool`

GetRebuildKernelCache returns the RebuildKernelCache field if non-nil, zero value otherwise.

### GetRebuildKernelCacheOk

`func (o *MdmCommandRequestCommandData) GetRebuildKernelCacheOk() (*bool, bool)`

GetRebuildKernelCacheOk returns a tuple with the RebuildKernelCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebuildKernelCache

`func (o *MdmCommandRequestCommandData) SetRebuildKernelCache(v bool)`

SetRebuildKernelCache sets RebuildKernelCache field to given value.

### HasRebuildKernelCache

`func (o *MdmCommandRequestCommandData) HasRebuildKernelCache() bool`

HasRebuildKernelCache returns a boolean if a field has been set.

### GetKextPaths

`func (o *MdmCommandRequestCommandData) GetKextPaths() []string`

GetKextPaths returns the KextPaths field if non-nil, zero value otherwise.

### GetKextPathsOk

`func (o *MdmCommandRequestCommandData) GetKextPathsOk() (*[]string, bool)`

GetKextPathsOk returns a tuple with the KextPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKextPaths

`func (o *MdmCommandRequestCommandData) SetKextPaths(v []string)`

SetKextPaths sets KextPaths field to given value.

### HasKextPaths

`func (o *MdmCommandRequestCommandData) HasKextPaths() bool`

HasKextPaths returns a boolean if a field has been set.

### GetNotifyUser

`func (o *MdmCommandRequestCommandData) GetNotifyUser() bool`

GetNotifyUser returns the NotifyUser field if non-nil, zero value otherwise.

### GetNotifyUserOk

`func (o *MdmCommandRequestCommandData) GetNotifyUserOk() (*bool, bool)`

GetNotifyUserOk returns a tuple with the NotifyUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyUser

`func (o *MdmCommandRequestCommandData) SetNotifyUser(v bool)`

SetNotifyUser sets NotifyUser field to given value.

### HasNotifyUser

`func (o *MdmCommandRequestCommandData) HasNotifyUser() bool`

HasNotifyUser returns a boolean if a field has been set.

### GetDestinationDeviceId

`func (o *MdmCommandRequestCommandData) GetDestinationDeviceId() string`

GetDestinationDeviceId returns the DestinationDeviceId field if non-nil, zero value otherwise.

### GetDestinationDeviceIdOk

`func (o *MdmCommandRequestCommandData) GetDestinationDeviceIdOk() (*string, bool)`

GetDestinationDeviceIdOk returns a tuple with the DestinationDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationDeviceId

`func (o *MdmCommandRequestCommandData) SetDestinationDeviceId(v string)`

SetDestinationDeviceId sets DestinationDeviceId field to given value.

### HasDestinationDeviceId

`func (o *MdmCommandRequestCommandData) HasDestinationDeviceId() bool`

HasDestinationDeviceId returns a boolean if a field has been set.

### GetDestinationName

`func (o *MdmCommandRequestCommandData) GetDestinationName() string`

GetDestinationName returns the DestinationName field if non-nil, zero value otherwise.

### GetDestinationNameOk

`func (o *MdmCommandRequestCommandData) GetDestinationNameOk() (*string, bool)`

GetDestinationNameOk returns a tuple with the DestinationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationName

`func (o *MdmCommandRequestCommandData) SetDestinationName(v string)`

SetDestinationName sets DestinationName field to given value.

### HasDestinationName

`func (o *MdmCommandRequestCommandData) HasDestinationName() bool`

HasDestinationName returns a boolean if a field has been set.

### GetPassword

`func (o *MdmCommandRequestCommandData) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *MdmCommandRequestCommandData) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *MdmCommandRequestCommandData) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *MdmCommandRequestCommandData) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetScanTime

`func (o *MdmCommandRequestCommandData) GetScanTime() int64`

GetScanTime returns the ScanTime field if non-nil, zero value otherwise.

### GetScanTimeOk

`func (o *MdmCommandRequestCommandData) GetScanTimeOk() (*int64, bool)`

GetScanTimeOk returns a tuple with the ScanTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanTime

`func (o *MdmCommandRequestCommandData) SetScanTime(v int64)`

SetScanTime sets ScanTime field to given value.

### HasScanTime

`func (o *MdmCommandRequestCommandData) HasScanTime() bool`

HasScanTime returns a boolean if a field has been set.

### SetScanTimeNil

`func (o *MdmCommandRequestCommandData) SetScanTimeNil(b bool)`

 SetScanTimeNil sets the value for ScanTime to be an explicit nil

### UnsetScanTime
`func (o *MdmCommandRequestCommandData) UnsetScanTime()`

UnsetScanTime ensures that no value is present for ScanTime, not even an explicit nil
### GetGuid

`func (o *MdmCommandRequestCommandData) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *MdmCommandRequestCommandData) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *MdmCommandRequestCommandData) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *MdmCommandRequestCommandData) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetNewPassword

`func (o *MdmCommandRequestCommandData) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *MdmCommandRequestCommandData) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *MdmCommandRequestCommandData) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.

### HasNewPassword

`func (o *MdmCommandRequestCommandData) HasNewPassword() bool`

HasNewPassword returns a boolean if a field has been set.

### GetBootstrapTokenAllowed

`func (o *MdmCommandRequestCommandData) GetBootstrapTokenAllowed() bool`

GetBootstrapTokenAllowed returns the BootstrapTokenAllowed field if non-nil, zero value otherwise.

### GetBootstrapTokenAllowedOk

`func (o *MdmCommandRequestCommandData) GetBootstrapTokenAllowedOk() (*bool, bool)`

GetBootstrapTokenAllowedOk returns a tuple with the BootstrapTokenAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootstrapTokenAllowed

`func (o *MdmCommandRequestCommandData) SetBootstrapTokenAllowed(v bool)`

SetBootstrapTokenAllowed sets BootstrapTokenAllowed field to given value.

### HasBootstrapTokenAllowed

`func (o *MdmCommandRequestCommandData) HasBootstrapTokenAllowed() bool`

HasBootstrapTokenAllowed returns a boolean if a field has been set.

### SetBootstrapTokenAllowedNil

`func (o *MdmCommandRequestCommandData) SetBootstrapTokenAllowedNil(b bool)`

 SetBootstrapTokenAllowedNil sets the value for BootstrapTokenAllowed to be an explicit nil

### UnsetBootstrapTokenAllowed
`func (o *MdmCommandRequestCommandData) UnsetBootstrapTokenAllowed()`

UnsetBootstrapTokenAllowed ensures that no value is present for BootstrapTokenAllowed, not even an explicit nil
### GetBluetooth

`func (o *MdmCommandRequestCommandData) GetBluetooth() bool`

GetBluetooth returns the Bluetooth field if non-nil, zero value otherwise.

### GetBluetoothOk

`func (o *MdmCommandRequestCommandData) GetBluetoothOk() (*bool, bool)`

GetBluetoothOk returns a tuple with the Bluetooth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBluetooth

`func (o *MdmCommandRequestCommandData) SetBluetooth(v bool)`

SetBluetooth sets Bluetooth field to given value.

### HasBluetooth

`func (o *MdmCommandRequestCommandData) HasBluetooth() bool`

HasBluetooth returns a boolean if a field has been set.

### SetBluetoothNil

`func (o *MdmCommandRequestCommandData) SetBluetoothNil(b bool)`

 SetBluetoothNil sets the value for Bluetooth to be an explicit nil

### UnsetBluetooth
`func (o *MdmCommandRequestCommandData) UnsetBluetooth()`

UnsetBluetooth ensures that no value is present for Bluetooth, not even an explicit nil
### GetAppAnalytics

`func (o *MdmCommandRequestCommandData) GetAppAnalytics() AppAnalyticsSetting`

GetAppAnalytics returns the AppAnalytics field if non-nil, zero value otherwise.

### GetAppAnalyticsOk

`func (o *MdmCommandRequestCommandData) GetAppAnalyticsOk() (*AppAnalyticsSetting, bool)`

GetAppAnalyticsOk returns a tuple with the AppAnalytics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppAnalytics

`func (o *MdmCommandRequestCommandData) SetAppAnalytics(v AppAnalyticsSetting)`

SetAppAnalytics sets AppAnalytics field to given value.

### HasAppAnalytics

`func (o *MdmCommandRequestCommandData) HasAppAnalytics() bool`

HasAppAnalytics returns a boolean if a field has been set.

### GetDiagnosticSubmission

`func (o *MdmCommandRequestCommandData) GetDiagnosticSubmission() DiagnosticSubmissionSetting`

GetDiagnosticSubmission returns the DiagnosticSubmission field if non-nil, zero value otherwise.

### GetDiagnosticSubmissionOk

`func (o *MdmCommandRequestCommandData) GetDiagnosticSubmissionOk() (*DiagnosticSubmissionSetting, bool)`

GetDiagnosticSubmissionOk returns a tuple with the DiagnosticSubmission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiagnosticSubmission

`func (o *MdmCommandRequestCommandData) SetDiagnosticSubmission(v DiagnosticSubmissionSetting)`

SetDiagnosticSubmission sets DiagnosticSubmission field to given value.

### HasDiagnosticSubmission

`func (o *MdmCommandRequestCommandData) HasDiagnosticSubmission() bool`

HasDiagnosticSubmission returns a boolean if a field has been set.

### GetDataRoaming

`func (o *MdmCommandRequestCommandData) GetDataRoaming() DataRoamingSetting`

GetDataRoaming returns the DataRoaming field if non-nil, zero value otherwise.

### GetDataRoamingOk

`func (o *MdmCommandRequestCommandData) GetDataRoamingOk() (*DataRoamingSetting, bool)`

GetDataRoamingOk returns a tuple with the DataRoaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRoaming

`func (o *MdmCommandRequestCommandData) SetDataRoaming(v DataRoamingSetting)`

SetDataRoaming sets DataRoaming field to given value.

### HasDataRoaming

`func (o *MdmCommandRequestCommandData) HasDataRoaming() bool`

HasDataRoaming returns a boolean if a field has been set.

### GetDefaultApplications

`func (o *MdmCommandRequestCommandData) GetDefaultApplications() DefaultApplications`

GetDefaultApplications returns the DefaultApplications field if non-nil, zero value otherwise.

### GetDefaultApplicationsOk

`func (o *MdmCommandRequestCommandData) GetDefaultApplicationsOk() (*DefaultApplications, bool)`

GetDefaultApplicationsOk returns a tuple with the DefaultApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultApplications

`func (o *MdmCommandRequestCommandData) SetDefaultApplications(v DefaultApplications)`

SetDefaultApplications sets DefaultApplications field to given value.

### HasDefaultApplications

`func (o *MdmCommandRequestCommandData) HasDefaultApplications() bool`

HasDefaultApplications returns a boolean if a field has been set.

### GetVoiceRoaming

`func (o *MdmCommandRequestCommandData) GetVoiceRoaming() VoiceRoamingSetting`

GetVoiceRoaming returns the VoiceRoaming field if non-nil, zero value otherwise.

### GetVoiceRoamingOk

`func (o *MdmCommandRequestCommandData) GetVoiceRoamingOk() (*VoiceRoamingSetting, bool)`

GetVoiceRoamingOk returns a tuple with the VoiceRoaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceRoaming

`func (o *MdmCommandRequestCommandData) SetVoiceRoaming(v VoiceRoamingSetting)`

SetVoiceRoaming sets VoiceRoaming field to given value.

### HasVoiceRoaming

`func (o *MdmCommandRequestCommandData) HasVoiceRoaming() bool`

HasVoiceRoaming returns a boolean if a field has been set.

### GetPersonalHotspot

`func (o *MdmCommandRequestCommandData) GetPersonalHotspot() PersonalHotspotSetting`

GetPersonalHotspot returns the PersonalHotspot field if non-nil, zero value otherwise.

### GetPersonalHotspotOk

`func (o *MdmCommandRequestCommandData) GetPersonalHotspotOk() (*PersonalHotspotSetting, bool)`

GetPersonalHotspotOk returns a tuple with the PersonalHotspot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalHotspot

`func (o *MdmCommandRequestCommandData) SetPersonalHotspot(v PersonalHotspotSetting)`

SetPersonalHotspot sets PersonalHotspot field to given value.

### HasPersonalHotspot

`func (o *MdmCommandRequestCommandData) HasPersonalHotspot() bool`

HasPersonalHotspot returns a boolean if a field has been set.

### GetMaximumResidentUsers

`func (o *MdmCommandRequestCommandData) GetMaximumResidentUsers() int64`

GetMaximumResidentUsers returns the MaximumResidentUsers field if non-nil, zero value otherwise.

### GetMaximumResidentUsersOk

`func (o *MdmCommandRequestCommandData) GetMaximumResidentUsersOk() (*int64, bool)`

GetMaximumResidentUsersOk returns a tuple with the MaximumResidentUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumResidentUsers

`func (o *MdmCommandRequestCommandData) SetMaximumResidentUsers(v int64)`

SetMaximumResidentUsers sets MaximumResidentUsers field to given value.

### HasMaximumResidentUsers

`func (o *MdmCommandRequestCommandData) HasMaximumResidentUsers() bool`

HasMaximumResidentUsers returns a boolean if a field has been set.

### SetMaximumResidentUsersNil

`func (o *MdmCommandRequestCommandData) SetMaximumResidentUsersNil(b bool)`

 SetMaximumResidentUsersNil sets the value for MaximumResidentUsers to be an explicit nil

### UnsetMaximumResidentUsers
`func (o *MdmCommandRequestCommandData) UnsetMaximumResidentUsers()`

UnsetMaximumResidentUsers ensures that no value is present for MaximumResidentUsers, not even an explicit nil
### GetDeviceName

`func (o *MdmCommandRequestCommandData) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *MdmCommandRequestCommandData) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *MdmCommandRequestCommandData) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *MdmCommandRequestCommandData) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetApplicationAttributes

`func (o *MdmCommandRequestCommandData) GetApplicationAttributes() ApplicationAttributes`

GetApplicationAttributes returns the ApplicationAttributes field if non-nil, zero value otherwise.

### GetApplicationAttributesOk

`func (o *MdmCommandRequestCommandData) GetApplicationAttributesOk() (*ApplicationAttributes, bool)`

GetApplicationAttributesOk returns a tuple with the ApplicationAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationAttributes

`func (o *MdmCommandRequestCommandData) SetApplicationAttributes(v ApplicationAttributes)`

SetApplicationAttributes sets ApplicationAttributes field to given value.

### HasApplicationAttributes

`func (o *MdmCommandRequestCommandData) HasApplicationAttributes() bool`

HasApplicationAttributes returns a boolean if a field has been set.

### GetSharedDeviceConfiguration

`func (o *MdmCommandRequestCommandData) GetSharedDeviceConfiguration() SharedDeviceConfiguration`

GetSharedDeviceConfiguration returns the SharedDeviceConfiguration field if non-nil, zero value otherwise.

### GetSharedDeviceConfigurationOk

`func (o *MdmCommandRequestCommandData) GetSharedDeviceConfigurationOk() (*SharedDeviceConfiguration, bool)`

GetSharedDeviceConfigurationOk returns a tuple with the SharedDeviceConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedDeviceConfiguration

`func (o *MdmCommandRequestCommandData) SetSharedDeviceConfiguration(v SharedDeviceConfiguration)`

SetSharedDeviceConfiguration sets SharedDeviceConfiguration field to given value.

### HasSharedDeviceConfiguration

`func (o *MdmCommandRequestCommandData) HasSharedDeviceConfiguration() bool`

HasSharedDeviceConfiguration returns a boolean if a field has been set.

### GetApplicationConfiguration

`func (o *MdmCommandRequestCommandData) GetApplicationConfiguration() ApplicationConfiguration`

GetApplicationConfiguration returns the ApplicationConfiguration field if non-nil, zero value otherwise.

### GetApplicationConfigurationOk

`func (o *MdmCommandRequestCommandData) GetApplicationConfigurationOk() (*ApplicationConfiguration, bool)`

GetApplicationConfigurationOk returns a tuple with the ApplicationConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationConfiguration

`func (o *MdmCommandRequestCommandData) SetApplicationConfiguration(v ApplicationConfiguration)`

SetApplicationConfiguration sets ApplicationConfiguration field to given value.

### HasApplicationConfiguration

`func (o *MdmCommandRequestCommandData) HasApplicationConfiguration() bool`

HasApplicationConfiguration returns a boolean if a field has been set.

### GetTimeZone

`func (o *MdmCommandRequestCommandData) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *MdmCommandRequestCommandData) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *MdmCommandRequestCommandData) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *MdmCommandRequestCommandData) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### SetTimeZoneNil

`func (o *MdmCommandRequestCommandData) SetTimeZoneNil(b bool)`

 SetTimeZoneNil sets the value for TimeZone to be an explicit nil

### UnsetTimeZone
`func (o *MdmCommandRequestCommandData) UnsetTimeZone()`

UnsetTimeZone ensures that no value is present for TimeZone, not even an explicit nil
### GetSoftwareUpdateSettings

`func (o *MdmCommandRequestCommandData) GetSoftwareUpdateSettings() SoftwareUpdateSettings`

GetSoftwareUpdateSettings returns the SoftwareUpdateSettings field if non-nil, zero value otherwise.

### GetSoftwareUpdateSettingsOk

`func (o *MdmCommandRequestCommandData) GetSoftwareUpdateSettingsOk() (*SoftwareUpdateSettings, bool)`

GetSoftwareUpdateSettingsOk returns a tuple with the SoftwareUpdateSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareUpdateSettings

`func (o *MdmCommandRequestCommandData) SetSoftwareUpdateSettings(v SoftwareUpdateSettings)`

SetSoftwareUpdateSettings sets SoftwareUpdateSettings field to given value.

### HasSoftwareUpdateSettings

`func (o *MdmCommandRequestCommandData) HasSoftwareUpdateSettings() bool`

HasSoftwareUpdateSettings returns a boolean if a field has been set.

### GetPasscodeLockGracePeriod

`func (o *MdmCommandRequestCommandData) GetPasscodeLockGracePeriod() int64`

GetPasscodeLockGracePeriod returns the PasscodeLockGracePeriod field if non-nil, zero value otherwise.

### GetPasscodeLockGracePeriodOk

`func (o *MdmCommandRequestCommandData) GetPasscodeLockGracePeriodOk() (*int64, bool)`

GetPasscodeLockGracePeriodOk returns a tuple with the PasscodeLockGracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasscodeLockGracePeriod

`func (o *MdmCommandRequestCommandData) SetPasscodeLockGracePeriod(v int64)`

SetPasscodeLockGracePeriod sets PasscodeLockGracePeriod field to given value.

### HasPasscodeLockGracePeriod

`func (o *MdmCommandRequestCommandData) HasPasscodeLockGracePeriod() bool`

HasPasscodeLockGracePeriod returns a boolean if a field has been set.

### SetPasscodeLockGracePeriodNil

`func (o *MdmCommandRequestCommandData) SetPasscodeLockGracePeriodNil(b bool)`

 SetPasscodeLockGracePeriodNil sets the value for PasscodeLockGracePeriod to be an explicit nil

### UnsetPasscodeLockGracePeriod
`func (o *MdmCommandRequestCommandData) UnsetPasscodeLockGracePeriod()`

UnsetPasscodeLockGracePeriod ensures that no value is present for PasscodeLockGracePeriod, not even an explicit nil
### GetActivationLockAllowedWhileSupervised

`func (o *MdmCommandRequestCommandData) GetActivationLockAllowedWhileSupervised() bool`

GetActivationLockAllowedWhileSupervised returns the ActivationLockAllowedWhileSupervised field if non-nil, zero value otherwise.

### GetActivationLockAllowedWhileSupervisedOk

`func (o *MdmCommandRequestCommandData) GetActivationLockAllowedWhileSupervisedOk() (*bool, bool)`

GetActivationLockAllowedWhileSupervisedOk returns a tuple with the ActivationLockAllowedWhileSupervised field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationLockAllowedWhileSupervised

`func (o *MdmCommandRequestCommandData) SetActivationLockAllowedWhileSupervised(v bool)`

SetActivationLockAllowedWhileSupervised sets ActivationLockAllowedWhileSupervised field to given value.

### HasActivationLockAllowedWhileSupervised

`func (o *MdmCommandRequestCommandData) HasActivationLockAllowedWhileSupervised() bool`

HasActivationLockAllowedWhileSupervised returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


