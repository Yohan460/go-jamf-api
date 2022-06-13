# \StartupStatusApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StartupStatusGet**](StartupStatusApi.md#StartupStatusGet) | **Get** /startup-status | Retrieve information about application startup 



## StartupStatusGet

> StartupStatus StartupStatusGet(ctx).Execute()

Retrieve information about application startup 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StartupStatusApi.StartupStatusGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StartupStatusApi.StartupStatusGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartupStatusGet`: StartupStatus
    fmt.Fprintf(os.Stdout, "Response from `StartupStatusApi.StartupStatusGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStartupStatusGetRequest struct via the builder pattern


### Return type

[**StartupStatus**](StartupStatus.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

