# StartupStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Step** | Pointer to **string** |  | [optional] 
**StepCode** | Pointer to **string** |  | [optional] 
**StepParam** | Pointer to **NullableString** |  | [optional] 
**Percentage** | Pointer to **int64** |  | [optional] 
**Warning** | Pointer to **NullableString** |  | [optional] 
**WarningCode** | Pointer to **NullableString** |  | [optional] 
**WarningParam** | Pointer to **NullableString** |  | [optional] 
**Error** | Pointer to **NullableString** |  | [optional] 
**ErrorCode** | Pointer to **NullableString** |  | [optional] 
**SetupAssistantNecessary** | Pointer to **bool** |  | [optional] 

## Methods

### NewStartupStatus

`func NewStartupStatus() *StartupStatus`

NewStartupStatus instantiates a new StartupStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartupStatusWithDefaults

`func NewStartupStatusWithDefaults() *StartupStatus`

NewStartupStatusWithDefaults instantiates a new StartupStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStep

`func (o *StartupStatus) GetStep() string`

GetStep returns the Step field if non-nil, zero value otherwise.

### GetStepOk

`func (o *StartupStatus) GetStepOk() (*string, bool)`

GetStepOk returns a tuple with the Step field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStep

`func (o *StartupStatus) SetStep(v string)`

SetStep sets Step field to given value.

### HasStep

`func (o *StartupStatus) HasStep() bool`

HasStep returns a boolean if a field has been set.

### GetStepCode

`func (o *StartupStatus) GetStepCode() string`

GetStepCode returns the StepCode field if non-nil, zero value otherwise.

### GetStepCodeOk

`func (o *StartupStatus) GetStepCodeOk() (*string, bool)`

GetStepCodeOk returns a tuple with the StepCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepCode

`func (o *StartupStatus) SetStepCode(v string)`

SetStepCode sets StepCode field to given value.

### HasStepCode

`func (o *StartupStatus) HasStepCode() bool`

HasStepCode returns a boolean if a field has been set.

### GetStepParam

`func (o *StartupStatus) GetStepParam() string`

GetStepParam returns the StepParam field if non-nil, zero value otherwise.

### GetStepParamOk

`func (o *StartupStatus) GetStepParamOk() (*string, bool)`

GetStepParamOk returns a tuple with the StepParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepParam

`func (o *StartupStatus) SetStepParam(v string)`

SetStepParam sets StepParam field to given value.

### HasStepParam

`func (o *StartupStatus) HasStepParam() bool`

HasStepParam returns a boolean if a field has been set.

### SetStepParamNil

`func (o *StartupStatus) SetStepParamNil(b bool)`

 SetStepParamNil sets the value for StepParam to be an explicit nil

### UnsetStepParam
`func (o *StartupStatus) UnsetStepParam()`

UnsetStepParam ensures that no value is present for StepParam, not even an explicit nil
### GetPercentage

`func (o *StartupStatus) GetPercentage() int64`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *StartupStatus) GetPercentageOk() (*int64, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *StartupStatus) SetPercentage(v int64)`

SetPercentage sets Percentage field to given value.

### HasPercentage

`func (o *StartupStatus) HasPercentage() bool`

HasPercentage returns a boolean if a field has been set.

### GetWarning

`func (o *StartupStatus) GetWarning() string`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *StartupStatus) GetWarningOk() (*string, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *StartupStatus) SetWarning(v string)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *StartupStatus) HasWarning() bool`

HasWarning returns a boolean if a field has been set.

### SetWarningNil

`func (o *StartupStatus) SetWarningNil(b bool)`

 SetWarningNil sets the value for Warning to be an explicit nil

### UnsetWarning
`func (o *StartupStatus) UnsetWarning()`

UnsetWarning ensures that no value is present for Warning, not even an explicit nil
### GetWarningCode

`func (o *StartupStatus) GetWarningCode() string`

GetWarningCode returns the WarningCode field if non-nil, zero value otherwise.

### GetWarningCodeOk

`func (o *StartupStatus) GetWarningCodeOk() (*string, bool)`

GetWarningCodeOk returns a tuple with the WarningCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningCode

`func (o *StartupStatus) SetWarningCode(v string)`

SetWarningCode sets WarningCode field to given value.

### HasWarningCode

`func (o *StartupStatus) HasWarningCode() bool`

HasWarningCode returns a boolean if a field has been set.

### SetWarningCodeNil

`func (o *StartupStatus) SetWarningCodeNil(b bool)`

 SetWarningCodeNil sets the value for WarningCode to be an explicit nil

### UnsetWarningCode
`func (o *StartupStatus) UnsetWarningCode()`

UnsetWarningCode ensures that no value is present for WarningCode, not even an explicit nil
### GetWarningParam

`func (o *StartupStatus) GetWarningParam() string`

GetWarningParam returns the WarningParam field if non-nil, zero value otherwise.

### GetWarningParamOk

`func (o *StartupStatus) GetWarningParamOk() (*string, bool)`

GetWarningParamOk returns a tuple with the WarningParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningParam

`func (o *StartupStatus) SetWarningParam(v string)`

SetWarningParam sets WarningParam field to given value.

### HasWarningParam

`func (o *StartupStatus) HasWarningParam() bool`

HasWarningParam returns a boolean if a field has been set.

### SetWarningParamNil

`func (o *StartupStatus) SetWarningParamNil(b bool)`

 SetWarningParamNil sets the value for WarningParam to be an explicit nil

### UnsetWarningParam
`func (o *StartupStatus) UnsetWarningParam()`

UnsetWarningParam ensures that no value is present for WarningParam, not even an explicit nil
### GetError

`func (o *StartupStatus) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *StartupStatus) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *StartupStatus) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *StartupStatus) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *StartupStatus) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *StartupStatus) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetErrorCode

`func (o *StartupStatus) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *StartupStatus) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *StartupStatus) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *StartupStatus) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *StartupStatus) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *StartupStatus) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetSetupAssistantNecessary

`func (o *StartupStatus) GetSetupAssistantNecessary() bool`

GetSetupAssistantNecessary returns the SetupAssistantNecessary field if non-nil, zero value otherwise.

### GetSetupAssistantNecessaryOk

`func (o *StartupStatus) GetSetupAssistantNecessaryOk() (*bool, bool)`

GetSetupAssistantNecessaryOk returns a tuple with the SetupAssistantNecessary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetupAssistantNecessary

`func (o *StartupStatus) SetSetupAssistantNecessary(v bool)`

SetSetupAssistantNecessary sets SetupAssistantNecessary field to given value.

### HasSetupAssistantNecessary

`func (o *StartupStatus) HasSetupAssistantNecessary() bool`

HasSetupAssistantNecessary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


