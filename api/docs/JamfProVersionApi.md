# \JamfProVersionAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1JamfProVersionGet**](JamfProVersionAPI.md#V1JamfProVersionGet) | **Get** /v1/jamf-pro-version | Return information about the Jamf Pro including the current version 



## V1JamfProVersionGet

> JamfProVersion V1JamfProVersionGet(ctx).Execute()

Return information about the Jamf Pro including the current version 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/yohan460/go-jamf-api"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JamfProVersionAPI.V1JamfProVersionGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JamfProVersionAPI.V1JamfProVersionGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1JamfProVersionGet`: JamfProVersion
    fmt.Fprintf(os.Stdout, "Response from `JamfProVersionAPI.V1JamfProVersionGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1JamfProVersionGetRequest struct via the builder pattern


### Return type

[**JamfProVersion**](JamfProVersion.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

