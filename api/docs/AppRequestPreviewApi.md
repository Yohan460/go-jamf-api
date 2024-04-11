# \AppRequestPreviewAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AppRequestFormInputFieldsGet**](AppRequestPreviewAPI.md#V1AppRequestFormInputFieldsGet) | **Get** /v1/app-request/form-input-fields | Search for Form Input Fields 
[**V1AppRequestFormInputFieldsIdDelete**](AppRequestPreviewAPI.md#V1AppRequestFormInputFieldsIdDelete) | **Delete** /v1/app-request/form-input-fields/{id} | Remove specified Form Input Field record 
[**V1AppRequestFormInputFieldsIdGet**](AppRequestPreviewAPI.md#V1AppRequestFormInputFieldsIdGet) | **Get** /v1/app-request/form-input-fields/{id} | Get specified Form Input Field object 
[**V1AppRequestFormInputFieldsIdPut**](AppRequestPreviewAPI.md#V1AppRequestFormInputFieldsIdPut) | **Put** /v1/app-request/form-input-fields/{id} | Update specified Form Input Field object 
[**V1AppRequestFormInputFieldsPost**](AppRequestPreviewAPI.md#V1AppRequestFormInputFieldsPost) | **Post** /v1/app-request/form-input-fields | Create Form Input Field record 
[**V1AppRequestFormInputFieldsPut**](AppRequestPreviewAPI.md#V1AppRequestFormInputFieldsPut) | **Put** /v1/app-request/form-input-fields | Replace all Form Input Fields 
[**V1AppRequestSettingsGet**](AppRequestPreviewAPI.md#V1AppRequestSettingsGet) | **Get** /v1/app-request/settings | Get Applicastion Request Settings 
[**V1AppRequestSettingsPut**](AppRequestPreviewAPI.md#V1AppRequestSettingsPut) | **Put** /v1/app-request/settings | Update Application Request Settings 



## V1AppRequestFormInputFieldsGet

> AppRequestFormInputFieldSearchResults V1AppRequestFormInputFieldsGet(ctx).Execute()

Search for Form Input Fields 



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppRequestPreviewAPI.V1AppRequestFormInputFieldsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppRequestPreviewAPI.V1AppRequestFormInputFieldsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AppRequestFormInputFieldsGet`: AppRequestFormInputFieldSearchResults
	fmt.Fprintf(os.Stdout, "Response from `AppRequestPreviewAPI.V1AppRequestFormInputFieldsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1AppRequestFormInputFieldsGetRequest struct via the builder pattern


### Return type

[**AppRequestFormInputFieldSearchResults**](AppRequestFormInputFieldSearchResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AppRequestFormInputFieldsIdDelete

> V1AppRequestFormInputFieldsIdDelete(ctx, id).Execute()

Remove specified Form Input Field record 



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
	id := int64(56) // int64 | Instance id of form input field record

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AppRequestPreviewAPI.V1AppRequestFormInputFieldsIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppRequestPreviewAPI.V1AppRequestFormInputFieldsIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Instance id of form input field record | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AppRequestFormInputFieldsIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AppRequestFormInputFieldsIdGet

> AppRequestFormInputField V1AppRequestFormInputFieldsIdGet(ctx, id).Execute()

Get specified Form Input Field object 



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
	id := int64(56) // int64 | Instance id of form input field record

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppRequestPreviewAPI.V1AppRequestFormInputFieldsIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppRequestPreviewAPI.V1AppRequestFormInputFieldsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AppRequestFormInputFieldsIdGet`: AppRequestFormInputField
	fmt.Fprintf(os.Stdout, "Response from `AppRequestPreviewAPI.V1AppRequestFormInputFieldsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Instance id of form input field record | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AppRequestFormInputFieldsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppRequestFormInputField**](AppRequestFormInputField.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AppRequestFormInputFieldsIdPut

> AppRequestFormInputField V1AppRequestFormInputFieldsIdPut(ctx, id).AppRequestFormInputField(appRequestFormInputField).Execute()

Update specified Form Input Field object 



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
	id := int64(56) // int64 | Instance id of form input field record
	appRequestFormInputField := *openapiclient.NewAppRequestFormInputField("Quantity", int64(1)) // AppRequestFormInputField | form input field object to create. ids defined in this body will be ignored

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppRequestPreviewAPI.V1AppRequestFormInputFieldsIdPut(context.Background(), id).AppRequestFormInputField(appRequestFormInputField).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppRequestPreviewAPI.V1AppRequestFormInputFieldsIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AppRequestFormInputFieldsIdPut`: AppRequestFormInputField
	fmt.Fprintf(os.Stdout, "Response from `AppRequestPreviewAPI.V1AppRequestFormInputFieldsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Instance id of form input field record | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AppRequestFormInputFieldsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appRequestFormInputField** | [**AppRequestFormInputField**](AppRequestFormInputField.md) | form input field object to create. ids defined in this body will be ignored | 

### Return type

[**AppRequestFormInputField**](AppRequestFormInputField.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AppRequestFormInputFieldsPost

> AppRequestFormInputField V1AppRequestFormInputFieldsPost(ctx).AppRequestFormInputField(appRequestFormInputField).Execute()

Create Form Input Field record 



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
	appRequestFormInputField := *openapiclient.NewAppRequestFormInputField("Quantity", int64(1)) // AppRequestFormInputField | form input field object to create. ids defined in this body will be ignored

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppRequestPreviewAPI.V1AppRequestFormInputFieldsPost(context.Background()).AppRequestFormInputField(appRequestFormInputField).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppRequestPreviewAPI.V1AppRequestFormInputFieldsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AppRequestFormInputFieldsPost`: AppRequestFormInputField
	fmt.Fprintf(os.Stdout, "Response from `AppRequestPreviewAPI.V1AppRequestFormInputFieldsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1AppRequestFormInputFieldsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appRequestFormInputField** | [**AppRequestFormInputField**](AppRequestFormInputField.md) | form input field object to create. ids defined in this body will be ignored | 

### Return type

[**AppRequestFormInputField**](AppRequestFormInputField.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AppRequestFormInputFieldsPut

> []AppRequestFormInputField V1AppRequestFormInputFieldsPut(ctx).AppRequestFormInputField(appRequestFormInputField).Execute()

Replace all Form Input Fields 



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
	appRequestFormInputField := []openapiclient.AppRequestFormInputField{*openapiclient.NewAppRequestFormInputField("Quantity", int64(1))} // []AppRequestFormInputField | list of form input fields to replace all existing fields. Will delete, update, and create all input fields accordingly.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppRequestPreviewAPI.V1AppRequestFormInputFieldsPut(context.Background()).AppRequestFormInputField(appRequestFormInputField).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppRequestPreviewAPI.V1AppRequestFormInputFieldsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AppRequestFormInputFieldsPut`: []AppRequestFormInputField
	fmt.Fprintf(os.Stdout, "Response from `AppRequestPreviewAPI.V1AppRequestFormInputFieldsPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1AppRequestFormInputFieldsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appRequestFormInputField** | [**[]AppRequestFormInputField**](AppRequestFormInputField.md) | list of form input fields to replace all existing fields. Will delete, update, and create all input fields accordingly. | 

### Return type

[**[]AppRequestFormInputField**](AppRequestFormInputField.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AppRequestSettingsGet

> AppRequestSettings V1AppRequestSettingsGet(ctx).Execute()

Get Applicastion Request Settings 



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppRequestPreviewAPI.V1AppRequestSettingsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppRequestPreviewAPI.V1AppRequestSettingsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AppRequestSettingsGet`: AppRequestSettings
	fmt.Fprintf(os.Stdout, "Response from `AppRequestPreviewAPI.V1AppRequestSettingsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1AppRequestSettingsGetRequest struct via the builder pattern


### Return type

[**AppRequestSettings**](AppRequestSettings.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AppRequestSettingsPut

> AppRequestSettings V1AppRequestSettingsPut(ctx).AppRequestSettings(appRequestSettings).Execute()

Update Application Request Settings 



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
	appRequestSettings := *openapiclient.NewAppRequestSettings() // AppRequestSettings | App request settings object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppRequestPreviewAPI.V1AppRequestSettingsPut(context.Background()).AppRequestSettings(appRequestSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppRequestPreviewAPI.V1AppRequestSettingsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AppRequestSettingsPut`: AppRequestSettings
	fmt.Fprintf(os.Stdout, "Response from `AppRequestPreviewAPI.V1AppRequestSettingsPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1AppRequestSettingsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appRequestSettings** | [**AppRequestSettings**](AppRequestSettings.md) | App request settings object | 

### Return type

[**AppRequestSettings**](AppRequestSettings.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

