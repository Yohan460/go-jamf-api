# \TimeZonesPreviewAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1TimeZonesGet**](TimeZonesPreviewAPI.md#V1TimeZonesGet) | **Get** /v1/time-zones | Return information about the currently supported Time Zones 



## V1TimeZonesGet

> []TimeZone V1TimeZonesGet(ctx).Execute()

Return information about the currently supported Time Zones 



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
    resp, r, err := apiClient.TimeZonesPreviewAPI.V1TimeZonesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimeZonesPreviewAPI.V1TimeZonesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1TimeZonesGet`: []TimeZone
    fmt.Fprintf(os.Stdout, "Response from `TimeZonesPreviewAPI.V1TimeZonesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1TimeZonesGetRequest struct via the builder pattern


### Return type

[**[]TimeZone**](TimeZone.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

