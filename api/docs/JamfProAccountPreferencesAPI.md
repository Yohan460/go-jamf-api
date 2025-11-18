# \JamfProAccountPreferencesAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V2AccountPreferencesGet**](JamfProAccountPreferencesAPI.md#V2AccountPreferencesGet) | **Get** /v2/account-preferences | Get Jamf Pro account preferences 
[**V2AccountPreferencesPatch**](JamfProAccountPreferencesAPI.md#V2AccountPreferencesPatch) | **Patch** /v2/account-preferences | Update Jamf Pro account preferences 
[**V3AccountPreferencesGet**](JamfProAccountPreferencesAPI.md#V3AccountPreferencesGet) | **Get** /v3/account-preferences | Get Jamf Pro account preferences 
[**V3AccountPreferencesPatch**](JamfProAccountPreferencesAPI.md#V3AccountPreferencesPatch) | **Patch** /v3/account-preferences | Update Jamf Pro account preferences 



## V2AccountPreferencesGet

> AccountPreferencesV5 V2AccountPreferencesGet(ctx).AcceptLanguage(acceptLanguage).Execute()

Get Jamf Pro account preferences 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yohan460/go-jamf-api/api"
)

func main() {
	acceptLanguage := "acceptLanguage_example" // string | Locale to be used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JamfProAccountPreferencesAPI.V2AccountPreferencesGet(context.Background()).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JamfProAccountPreferencesAPI.V2AccountPreferencesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2AccountPreferencesGet`: AccountPreferencesV5
	fmt.Fprintf(os.Stdout, "Response from `JamfProAccountPreferencesAPI.V2AccountPreferencesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2AccountPreferencesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptLanguage** | **string** | Locale to be used. | 

### Return type

[**AccountPreferencesV5**](AccountPreferencesV5.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2AccountPreferencesPatch

> AccountPreferencesV5 V2AccountPreferencesPatch(ctx).AcceptLanguage(acceptLanguage).JSESSIONID(jSESSIONID).AccountPreferencesV5(accountPreferencesV5).Execute()

Update Jamf Pro account preferences 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yohan460/go-jamf-api/api"
)

func main() {
	acceptLanguage := "acceptLanguage_example" // string | Locale to be used, when user has not defined preferred language. (optional)
	jSESSIONID := "jSESSIONID_example" // string | Session cookie, that's used to determine user session where account preferences should be refreshed (optional) (default to "null")
	accountPreferencesV5 := *openapiclient.NewAccountPreferencesV5("Language_example", "MM/dd/yyyy", "America/Chicago", int64(50), openapiclient.AccountPreferencesUserInterfaceDisplayTheme("MATCH_SYSTEM"), false, true, true, true, "ALPHABETICALLY", openapiclient.AccountPreferencesSearchType("EXACT_MATCH"), openapiclient.AccountPreferencesSearchType("EXACT_MATCH"), openapiclient.AccountPreferencesSearchType("EXACT_MATCH"), openapiclient.AccountPreferencesSearchType("EXACT_MATCH"), openapiclient.AccountPreferencesSearchType("EXACT_MATCH"), openapiclient.AccountPreferencesSearchType("EXACT_MATCH"), openapiclient.AccountPreferencesSearchType("EXACT_MATCH"), openapiclient.AccountPreferencesSearchType("EXACT_MATCH"), openapiclient.AccountPreferencesSearchType("EXACT_MATCH"), openapiclient.AccountPreferencesSearchType("EXACT_MATCH"), openapiclient.AccountPreferencesSearchType("EXACT_MATCH"), openapiclient.AccountPreferencesSearchType("EXACT_MATCH"), openapiclient.AccountPreferencesSearchType("EXACT_MATCH"), openapiclient.AccountPreferencesSearchType("EXACT_MATCH")) // AccountPreferencesV5 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JamfProAccountPreferencesAPI.V2AccountPreferencesPatch(context.Background()).AcceptLanguage(acceptLanguage).JSESSIONID(jSESSIONID).AccountPreferencesV5(accountPreferencesV5).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JamfProAccountPreferencesAPI.V2AccountPreferencesPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2AccountPreferencesPatch`: AccountPreferencesV5
	fmt.Fprintf(os.Stdout, "Response from `JamfProAccountPreferencesAPI.V2AccountPreferencesPatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2AccountPreferencesPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptLanguage** | **string** | Locale to be used, when user has not defined preferred language. | 
 **jSESSIONID** | **string** | Session cookie, that&#39;s used to determine user session where account preferences should be refreshed | [default to &quot;null&quot;]
 **accountPreferencesV5** | [**AccountPreferencesV5**](AccountPreferencesV5.md) |  | 

### Return type

[**AccountPreferencesV5**](AccountPreferencesV5.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3AccountPreferencesGet

> AccountPreferencesV6 V3AccountPreferencesGet(ctx).AcceptLanguage(acceptLanguage).Execute()

Get Jamf Pro account preferences 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yohan460/go-jamf-api/api"
)

func main() {
	acceptLanguage := "acceptLanguage_example" // string | Locale to be used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JamfProAccountPreferencesAPI.V3AccountPreferencesGet(context.Background()).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JamfProAccountPreferencesAPI.V3AccountPreferencesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AccountPreferencesGet`: AccountPreferencesV6
	fmt.Fprintf(os.Stdout, "Response from `JamfProAccountPreferencesAPI.V3AccountPreferencesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3AccountPreferencesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptLanguage** | **string** | Locale to be used. | 

### Return type

[**AccountPreferencesV6**](AccountPreferencesV6.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3AccountPreferencesPatch

> V3AccountPreferencesPatch(ctx).AcceptLanguage(acceptLanguage).JSESSIONID(jSESSIONID).AccountPreferencesV6(accountPreferencesV6).Execute()

Update Jamf Pro account preferences 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yohan460/go-jamf-api/api"
)

func main() {
	acceptLanguage := "acceptLanguage_example" // string | Locale to be used, when user has not defined preferred language. (optional)
	jSESSIONID := "jSESSIONID_example" // string | Session cookie, that's used to determine user session where account preferences should be refreshed (optional) (default to "null")
	accountPreferencesV6 := *openapiclient.NewAccountPreferencesV6("Language_example", "MM/dd/yyyy", "America/Chicago", int64(50), openapiclient.AccountPreferencesUserInterfaceDisplayTheme("MATCH_SYSTEM"), false, true, true, true, "ALPHABETICALLY", openapiclient.AccountPreferencesSearchType("EXACT_MATCH"), openapiclient.AccountPreferencesSearchType("EXACT_MATCH"), openapiclient.AccountPreferencesSearchType("EXACT_MATCH"), openapiclient.AccountPreferencesSearchType("EXACT_MATCH"), openapiclient.AccountPreferencesSearchType("EXACT_MATCH"), openapiclient.AccountPreferencesSearchType("EXACT_MATCH"), openapiclient.AccountPreferencesSearchType("EXACT_MATCH"), openapiclient.AccountPreferencesSearchType("EXACT_MATCH"), openapiclient.AccountPreferencesSearchType("EXACT_MATCH"), openapiclient.AccountPreferencesSearchType("EXACT_MATCH"), openapiclient.AccountPreferencesSearchType("EXACT_MATCH"), openapiclient.AccountPreferencesSearchType("EXACT_MATCH"), openapiclient.AccountPreferencesSearchType("EXACT_MATCH"), openapiclient.AccountPreferencesSearchType("EXACT_MATCH")) // AccountPreferencesV6 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.JamfProAccountPreferencesAPI.V3AccountPreferencesPatch(context.Background()).AcceptLanguage(acceptLanguage).JSESSIONID(jSESSIONID).AccountPreferencesV6(accountPreferencesV6).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JamfProAccountPreferencesAPI.V3AccountPreferencesPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3AccountPreferencesPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptLanguage** | **string** | Locale to be used, when user has not defined preferred language. | 
 **jSESSIONID** | **string** | Session cookie, that&#39;s used to determine user session where account preferences should be refreshed | [default to &quot;null&quot;]
 **accountPreferencesV6** | [**AccountPreferencesV6**](AccountPreferencesV6.md) |  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

