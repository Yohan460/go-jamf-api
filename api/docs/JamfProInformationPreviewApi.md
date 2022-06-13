# \JamfProInformationPreviewApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1JamfProInformationGet**](JamfProInformationPreviewApi.md#V1JamfProInformationGet) | **Get** /v1/jamf-pro-information | Get basic information about the Jamf Pro Server 



## V1JamfProInformationGet

> JamfProInformation V1JamfProInformationGet(ctx).Execute()

Get basic information about the Jamf Pro Server 



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
    resp, r, err := apiClient.JamfProInformationPreviewApi.V1JamfProInformationGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JamfProInformationPreviewApi.V1JamfProInformationGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1JamfProInformationGet`: JamfProInformation
    fmt.Fprintf(os.Stdout, "Response from `JamfProInformationPreviewApi.V1JamfProInformationGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1JamfProInformationGetRequest struct via the builder pattern


### Return type

[**JamfProInformation**](JamfProInformation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

