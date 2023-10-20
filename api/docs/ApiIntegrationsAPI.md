# \ApiIntegrationsAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiIntegration**](ApiIntegrationsAPI.md#DeleteApiIntegration) | **Delete** /v1/api-integrations/{id} | Remove specified API integration
[**GetOneApiIntegration**](ApiIntegrationsAPI.md#GetOneApiIntegration) | **Get** /v1/api-integrations/{id} | Get specified API integration object
[**PostCreateApiIntegration**](ApiIntegrationsAPI.md#PostCreateApiIntegration) | **Post** /v1/api-integrations | Create API integration object
[**PostCreateClientCredentials**](ApiIntegrationsAPI.md#PostCreateClientCredentials) | **Post** /v1/api-integrations/{id}/client-credentials | Create client credentials for specified API integration
[**PutUpdateApiIntegration**](ApiIntegrationsAPI.md#PutUpdateApiIntegration) | **Put** /v1/api-integrations/{id} | Update specified API integration object
[**V1ApiIntegrationsGet**](ApiIntegrationsAPI.md#V1ApiIntegrationsGet) | **Get** /v1/api-integrations | Get the current API Integrations



## DeleteApiIntegration

> DeleteApiIntegration(ctx, id).Execute()

Remove specified API integration



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
    id := "id_example" // string | instance id of api integration object

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ApiIntegrationsAPI.DeleteApiIntegration(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiIntegrationsAPI.DeleteApiIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of api integration object | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetOneApiIntegration

> ApiIntegrationResponse GetOneApiIntegration(ctx, id).Execute()

Get specified API integration object



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
    id := "id_example" // string | instance id of api integration object

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApiIntegrationsAPI.GetOneApiIntegration(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiIntegrationsAPI.GetOneApiIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOneApiIntegration`: ApiIntegrationResponse
    fmt.Fprintf(os.Stdout, "Response from `ApiIntegrationsAPI.GetOneApiIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of api integration object | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOneApiIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiIntegrationResponse**](ApiIntegrationResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCreateApiIntegration

> ApiIntegrationResponse PostCreateApiIntegration(ctx).ApiIntegrationRequest(apiIntegrationRequest).Execute()

Create API integration object



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
    apiIntegrationRequest := *openapiclient.NewApiIntegrationRequest([]string{"AuthorizationScopes_example"}, "My API Integration") // ApiIntegrationRequest | api integration object to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApiIntegrationsAPI.PostCreateApiIntegration(context.Background()).ApiIntegrationRequest(apiIntegrationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiIntegrationsAPI.PostCreateApiIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCreateApiIntegration`: ApiIntegrationResponse
    fmt.Fprintf(os.Stdout, "Response from `ApiIntegrationsAPI.PostCreateApiIntegration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCreateApiIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiIntegrationRequest** | [**ApiIntegrationRequest**](ApiIntegrationRequest.md) | api integration object to create | 

### Return type

[**ApiIntegrationResponse**](ApiIntegrationResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCreateClientCredentials

> OAuthClientCredentials PostCreateClientCredentials(ctx, id).Execute()

Create client credentials for specified API integration



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
    id := "id_example" // string | instance id of api integration object

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApiIntegrationsAPI.PostCreateClientCredentials(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiIntegrationsAPI.PostCreateClientCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCreateClientCredentials`: OAuthClientCredentials
    fmt.Fprintf(os.Stdout, "Response from `ApiIntegrationsAPI.PostCreateClientCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of api integration object | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCreateClientCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OAuthClientCredentials**](OAuthClientCredentials.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutUpdateApiIntegration

> ApiIntegrationResponse PutUpdateApiIntegration(ctx, id).ApiIntegrationRequest(apiIntegrationRequest).Execute()

Update specified API integration object



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
    id := "id_example" // string | instance id of api integration object
    apiIntegrationRequest := *openapiclient.NewApiIntegrationRequest([]string{"AuthorizationScopes_example"}, "My API Integration") // ApiIntegrationRequest | api object to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApiIntegrationsAPI.PutUpdateApiIntegration(context.Background(), id).ApiIntegrationRequest(apiIntegrationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiIntegrationsAPI.PutUpdateApiIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutUpdateApiIntegration`: ApiIntegrationResponse
    fmt.Fprintf(os.Stdout, "Response from `ApiIntegrationsAPI.PutUpdateApiIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of api integration object | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutUpdateApiIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiIntegrationRequest** | [**ApiIntegrationRequest**](ApiIntegrationRequest.md) | api object to update | 

### Return type

[**ApiIntegrationResponse**](ApiIntegrationResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ApiIntegrationsGet

> ApiIntegrationSearchResult V1ApiIntegrationsGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Get the current API Integrations



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
    page := int32(56) // int32 |  (optional) (default to 0)
    pageSize := int32(56) // int32 |  (optional) (default to 100)
    sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Fields allowed in the query: id, displayName. Example: sort=displayName:desc (optional) (default to ["id:asc"])
    filter := "filter_example" // string | Query in the RSQL format, allowing to filter app titles collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, displayName. Example: displayName==\"*IntegrationName*\" (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApiIntegrationsAPI.V1ApiIntegrationsGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiIntegrationsAPI.V1ApiIntegrationsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ApiIntegrationsGet`: ApiIntegrationSearchResult
    fmt.Fprintf(os.Stdout, "Response from `ApiIntegrationsAPI.V1ApiIntegrationsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ApiIntegrationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 0]
 **pageSize** | **int32** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Fields allowed in the query: id, displayName. Example: sort&#x3D;displayName:desc | [default to [&quot;id:asc&quot;]]
 **filter** | **string** | Query in the RSQL format, allowing to filter app titles collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, displayName. Example: displayName&#x3D;&#x3D;\&quot;*IntegrationName*\&quot; | [default to &quot;&quot;]

### Return type

[**ApiIntegrationSearchResult**](ApiIntegrationSearchResult.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

