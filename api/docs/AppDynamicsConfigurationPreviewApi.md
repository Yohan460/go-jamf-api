# \AppDynamicsConfigurationPreviewApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AppDynamicsScriptConfigurationGet**](AppDynamicsConfigurationPreviewApi.md#V1AppDynamicsScriptConfigurationGet) | **Get** /v1/app-dynamics/script-configuration | Get Application Dynamics Config object 



## V1AppDynamicsScriptConfigurationGet

> AppDynamicsConfig V1AppDynamicsScriptConfigurationGet(ctx).Execute()

Get Application Dynamics Config object 



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
    resp, r, err := apiClient.AppDynamicsConfigurationPreviewApi.V1AppDynamicsScriptConfigurationGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppDynamicsConfigurationPreviewApi.V1AppDynamicsScriptConfigurationGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1AppDynamicsScriptConfigurationGet`: AppDynamicsConfig
    fmt.Fprintf(os.Stdout, "Response from `AppDynamicsConfigurationPreviewApi.V1AppDynamicsScriptConfigurationGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1AppDynamicsScriptConfigurationGetRequest struct via the builder pattern


### Return type

[**AppDynamicsConfig**](AppDynamicsConfig.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

