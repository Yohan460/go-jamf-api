# \CloudInformationAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1CloudInformationGet**](CloudInformationAPI.md#V1CloudInformationGet) | **Get** /v1/cloud-information | Retrieve information related to cloud setup. 



## V1CloudInformationGet

> CloudResponse V1CloudInformationGet(ctx).Execute()

Retrieve information related to cloud setup. 



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
    resp, r, err := apiClient.CloudInformationAPI.V1CloudInformationGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudInformationAPI.V1CloudInformationGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1CloudInformationGet`: CloudResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudInformationAPI.V1CloudInformationGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1CloudInformationGetRequest struct via the builder pattern


### Return type

[**CloudResponse**](CloudResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

