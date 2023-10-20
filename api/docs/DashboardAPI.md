# \DashboardAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1DashboardGet**](DashboardAPI.md#V1DashboardGet) | **Get** /v1/dashboard | Get all the dashboard setup information 



## V1DashboardGet

> DashboardSetup V1DashboardGet(ctx).Execute()

Get all the dashboard setup information 



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
    resp, r, err := apiClient.DashboardAPI.V1DashboardGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.V1DashboardGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1DashboardGet`: DashboardSetup
    fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.V1DashboardGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1DashboardGetRequest struct via the builder pattern


### Return type

[**DashboardSetup**](DashboardSetup.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

