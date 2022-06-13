# \CsaApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1CsaTokenDelete**](CsaApi.md#V1CsaTokenDelete) | **Delete** /v1/csa/token | Delete the CSA token exchange - This will disable Jamf Pro&#39;s ability to authenticate with cloud-hosted services 
[**V1CsaTokenGet**](CsaApi.md#V1CsaTokenGet) | **Get** /v1/csa/token | Get details regarding the CSA token exchange 
[**V1CsaTokenPost**](CsaApi.md#V1CsaTokenPost) | **Post** /v1/csa/token | Initialize the CSA token exchange 
[**V1CsaTokenPut**](CsaApi.md#V1CsaTokenPut) | **Put** /v1/csa/token | Re-initialize the CSA token exchange with new credentials 



## V1CsaTokenDelete

> V1CsaTokenDelete(ctx).Execute()

Delete the CSA token exchange - This will disable Jamf Pro's ability to authenticate with cloud-hosted services 



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
    resp, r, err := apiClient.CsaApi.V1CsaTokenDelete(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CsaApi.V1CsaTokenDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1CsaTokenDeleteRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CsaTokenGet

> CsaToken V1CsaTokenGet(ctx).Execute()

Get details regarding the CSA token exchange 



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
    resp, r, err := apiClient.CsaApi.V1CsaTokenGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CsaApi.V1CsaTokenGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1CsaTokenGet`: CsaToken
    fmt.Fprintf(os.Stdout, "Response from `CsaApi.V1CsaTokenGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1CsaTokenGetRequest struct via the builder pattern


### Return type

[**CsaToken**](CsaToken.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CsaTokenPost

> CsaToken V1CsaTokenPost(ctx).JamfNationCredentials(jamfNationCredentials).Execute()

Initialize the CSA token exchange 



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
    jamfNationCredentials := *openapiclient.NewJamfNationCredentials() // JamfNationCredentials | Jamf Nation username and password  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CsaApi.V1CsaTokenPost(context.Background()).JamfNationCredentials(jamfNationCredentials).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CsaApi.V1CsaTokenPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1CsaTokenPost`: CsaToken
    fmt.Fprintf(os.Stdout, "Response from `CsaApi.V1CsaTokenPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1CsaTokenPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jamfNationCredentials** | [**JamfNationCredentials**](JamfNationCredentials.md) | Jamf Nation username and password  | 

### Return type

[**CsaToken**](CsaToken.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CsaTokenPut

> CsaToken V1CsaTokenPut(ctx).JamfNationCredentials(jamfNationCredentials).Execute()

Re-initialize the CSA token exchange with new credentials 



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
    jamfNationCredentials := *openapiclient.NewJamfNationCredentials() // JamfNationCredentials | Jamf Nation username and password  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CsaApi.V1CsaTokenPut(context.Background()).JamfNationCredentials(jamfNationCredentials).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CsaApi.V1CsaTokenPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1CsaTokenPut`: CsaToken
    fmt.Fprintf(os.Stdout, "Response from `CsaApi.V1CsaTokenPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1CsaTokenPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jamfNationCredentials** | [**JamfNationCredentials**](JamfNationCredentials.md) | Jamf Nation username and password  | 

### Return type

[**CsaToken**](CsaToken.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

