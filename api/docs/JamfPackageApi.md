# \JamfPackageApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1JamfPackageGet**](JamfPackageApi.md#V1JamfPackageGet) | **Get** /v1/jamf-package | Get the packages for a given Jamf application 
[**V2JamfPackageGet**](JamfPackageApi.md#V2JamfPackageGet) | **Get** /v2/jamf-package | Get the packages for a given Jamf application 



## V1JamfPackageGet

> []JamfPackageResponse V1JamfPackageGet(ctx).Application(application).Execute()

Get the packages for a given Jamf application 



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
    application := "protect" // string | The Jamf Application key. The only supported values are protect and connect.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JamfPackageApi.V1JamfPackageGet(context.Background()).Application(application).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JamfPackageApi.V1JamfPackageGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1JamfPackageGet`: []JamfPackageResponse
    fmt.Fprintf(os.Stdout, "Response from `JamfPackageApi.V1JamfPackageGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1JamfPackageGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **application** | **string** | The Jamf Application key. The only supported values are protect and connect. | 

### Return type

[**[]JamfPackageResponse**](JamfPackageResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2JamfPackageGet

> JamfApplicationResponse V2JamfPackageGet(ctx).Application(application).Execute()

Get the packages for a given Jamf application 



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
    application := "protect" // string | The Jamf Application key. The only supported values are protect and connect.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JamfPackageApi.V2JamfPackageGet(context.Background()).Application(application).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JamfPackageApi.V2JamfPackageGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V2JamfPackageGet`: JamfApplicationResponse
    fmt.Fprintf(os.Stdout, "Response from `JamfPackageApi.V2JamfPackageGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2JamfPackageGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **application** | **string** | The Jamf Application key. The only supported values are protect and connect. | 

### Return type

[**JamfApplicationResponse**](JamfApplicationResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

