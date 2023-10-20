# \JamfProInformationAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1JamfProInformationGet**](JamfProInformationAPI.md#V1JamfProInformationGet) | **Get** /v1/jamf-pro-information | Get basic information about the Jamf Pro Server 
[**V2JamfProInformationGet**](JamfProInformationAPI.md#V2JamfProInformationGet) | **Get** /v2/jamf-pro-information | Get basic information about the Jamf Pro Server 



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
    openapiclient "github.com/yohan460/go-jamf-api"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JamfProInformationAPI.V1JamfProInformationGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JamfProInformationAPI.V1JamfProInformationGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1JamfProInformationGet`: JamfProInformation
    fmt.Fprintf(os.Stdout, "Response from `JamfProInformationAPI.V1JamfProInformationGet`: %v\n", resp)
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


## V2JamfProInformationGet

> JamfProInformationV2 V2JamfProInformationGet(ctx).Execute()

Get basic information about the Jamf Pro Server 



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
    resp, r, err := apiClient.JamfProInformationAPI.V2JamfProInformationGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JamfProInformationAPI.V2JamfProInformationGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V2JamfProInformationGet`: JamfProInformationV2
    fmt.Fprintf(os.Stdout, "Response from `JamfProInformationAPI.V2JamfProInformationGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV2JamfProInformationGetRequest struct via the builder pattern


### Return type

[**JamfProInformationV2**](JamfProInformationV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

