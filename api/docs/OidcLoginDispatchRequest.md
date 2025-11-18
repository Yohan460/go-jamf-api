# OidcLoginDispatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginalUrl** | **string** | Original Url | 
**EmailAddress** | **string** | User email address | 

## Methods

### NewOidcLoginDispatchRequest

`func NewOidcLoginDispatchRequest(originalUrl string, emailAddress string, ) *OidcLoginDispatchRequest`

NewOidcLoginDispatchRequest instantiates a new OidcLoginDispatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOidcLoginDispatchRequestWithDefaults

`func NewOidcLoginDispatchRequestWithDefaults() *OidcLoginDispatchRequest`

NewOidcLoginDispatchRequestWithDefaults instantiates a new OidcLoginDispatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginalUrl

`func (o *OidcLoginDispatchRequest) GetOriginalUrl() string`

GetOriginalUrl returns the OriginalUrl field if non-nil, zero value otherwise.

### GetOriginalUrlOk

`func (o *OidcLoginDispatchRequest) GetOriginalUrlOk() (*string, bool)`

GetOriginalUrlOk returns a tuple with the OriginalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalUrl

`func (o *OidcLoginDispatchRequest) SetOriginalUrl(v string)`

SetOriginalUrl sets OriginalUrl field to given value.


### GetEmailAddress

`func (o *OidcLoginDispatchRequest) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *OidcLoginDispatchRequest) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *OidcLoginDispatchRequest) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


