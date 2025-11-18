# \LogFlushingAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1LogFlushingGet**](LogFlushingAPI.md#V1LogFlushingGet) | **Get** /v1/log-flushing | Get log flushing settings
[**V1LogFlushingTaskGet**](LogFlushingAPI.md#V1LogFlushingTaskGet) | **Get** /v1/log-flushing/task | Get log flushing tasks
[**V1LogFlushingTaskIdDelete**](LogFlushingAPI.md#V1LogFlushingTaskIdDelete) | **Delete** /v1/log-flushing/task/{id} | Cancels a log flushing task
[**V1LogFlushingTaskIdGet**](LogFlushingAPI.md#V1LogFlushingTaskIdGet) | **Get** /v1/log-flushing/task/{id} | Get log flushing task
[**V1LogFlushingTaskPost**](LogFlushingAPI.md#V1LogFlushingTaskPost) | **Post** /v1/log-flushing/task | Queue a log flushing task



## V1LogFlushingGet

> LogFlushingV1 V1LogFlushingGet(ctx).Execute()

Get log flushing settings



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
	resp, r, err := apiClient.LogFlushingAPI.V1LogFlushingGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogFlushingAPI.V1LogFlushingGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1LogFlushingGet`: LogFlushingV1
	fmt.Fprintf(os.Stdout, "Response from `LogFlushingAPI.V1LogFlushingGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1LogFlushingGetRequest struct via the builder pattern


### Return type

[**LogFlushingV1**](LogFlushingV1.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1LogFlushingTaskGet

> []LogFlushingTaskV1 V1LogFlushingTaskGet(ctx).Execute()

Get log flushing tasks



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
	resp, r, err := apiClient.LogFlushingAPI.V1LogFlushingTaskGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogFlushingAPI.V1LogFlushingTaskGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1LogFlushingTaskGet`: []LogFlushingTaskV1
	fmt.Fprintf(os.Stdout, "Response from `LogFlushingAPI.V1LogFlushingTaskGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1LogFlushingTaskGetRequest struct via the builder pattern


### Return type

[**[]LogFlushingTaskV1**](LogFlushingTaskV1.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1LogFlushingTaskIdDelete

> V1LogFlushingTaskIdDelete(ctx, id).Execute()

Cancels a log flushing task



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
	id := "id_example" // string | The identifier of the log flushing task

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LogFlushingAPI.V1LogFlushingTaskIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogFlushingAPI.V1LogFlushingTaskIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The identifier of the log flushing task | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1LogFlushingTaskIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1LogFlushingTaskIdGet

> LogFlushingTaskV1 V1LogFlushingTaskIdGet(ctx, id).Execute()

Get log flushing task



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
	id := "id_example" // string | The identifier of the log flushing task

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogFlushingAPI.V1LogFlushingTaskIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogFlushingAPI.V1LogFlushingTaskIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1LogFlushingTaskIdGet`: LogFlushingTaskV1
	fmt.Fprintf(os.Stdout, "Response from `LogFlushingAPI.V1LogFlushingTaskIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The identifier of the log flushing task | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1LogFlushingTaskIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LogFlushingTaskV1**](LogFlushingTaskV1.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1LogFlushingTaskPost

> HrefResponse V1LogFlushingTaskPost(ctx).LogFlushingTaskV1(logFlushingTaskV1).Execute()

Queue a log flushing task



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
	logFlushingTaskV1 := *openapiclient.NewLogFlushingTaskV1("policy", int64(3), "MONTH") // LogFlushingTaskV1 | The manual log flushing settings (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogFlushingAPI.V1LogFlushingTaskPost(context.Background()).LogFlushingTaskV1(logFlushingTaskV1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogFlushingAPI.V1LogFlushingTaskPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1LogFlushingTaskPost`: HrefResponse
	fmt.Fprintf(os.Stdout, "Response from `LogFlushingAPI.V1LogFlushingTaskPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1LogFlushingTaskPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **logFlushingTaskV1** | [**LogFlushingTaskV1**](LogFlushingTaskV1.md) | The manual log flushing settings | 

### Return type

[**HrefResponse**](HrefResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

