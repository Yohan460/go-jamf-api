# \CloudAzureAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1CloudAzureDefaultsMappingsGet**](CloudAzureAPI.md#V1CloudAzureDefaultsMappingsGet) | **Get** /v1/cloud-azure/defaults/mappings | Get default mappings
[**V1CloudAzureDefaultsServerConfigurationGet**](CloudAzureAPI.md#V1CloudAzureDefaultsServerConfigurationGet) | **Get** /v1/cloud-azure/defaults/server-configuration | Get default server configuration
[**V1CloudAzureIdDelete**](CloudAzureAPI.md#V1CloudAzureIdDelete) | **Delete** /v1/cloud-azure/{id} | Delete Cloud Identity Provider configuration.
[**V1CloudAzureIdGet**](CloudAzureAPI.md#V1CloudAzureIdGet) | **Get** /v1/cloud-azure/{id} | Get Azure Cloud Identity Provider configuration with given ID.
[**V1CloudAzureIdPut**](CloudAzureAPI.md#V1CloudAzureIdPut) | **Put** /v1/cloud-azure/{id} | Update Azure Cloud Identity Provider configuration
[**V1CloudAzurePost**](CloudAzureAPI.md#V1CloudAzurePost) | **Post** /v1/cloud-azure | Create Azure Cloud Identity Provider configuration



## V1CloudAzureDefaultsMappingsGet

> AzureMappings V1CloudAzureDefaultsMappingsGet(ctx).Execute()

Get default mappings



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
    resp, r, err := apiClient.CloudAzureAPI.V1CloudAzureDefaultsMappingsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudAzureAPI.V1CloudAzureDefaultsMappingsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1CloudAzureDefaultsMappingsGet`: AzureMappings
    fmt.Fprintf(os.Stdout, "Response from `CloudAzureAPI.V1CloudAzureDefaultsMappingsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1CloudAzureDefaultsMappingsGetRequest struct via the builder pattern


### Return type

[**AzureMappings**](AzureMappings.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CloudAzureDefaultsServerConfigurationGet

> AzureServerConfiguration V1CloudAzureDefaultsServerConfigurationGet(ctx).Execute()

Get default server configuration



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
    resp, r, err := apiClient.CloudAzureAPI.V1CloudAzureDefaultsServerConfigurationGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudAzureAPI.V1CloudAzureDefaultsServerConfigurationGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1CloudAzureDefaultsServerConfigurationGet`: AzureServerConfiguration
    fmt.Fprintf(os.Stdout, "Response from `CloudAzureAPI.V1CloudAzureDefaultsServerConfigurationGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1CloudAzureDefaultsServerConfigurationGetRequest struct via the builder pattern


### Return type

[**AzureServerConfiguration**](AzureServerConfiguration.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CloudAzureIdDelete

> V1CloudAzureIdDelete(ctx, id).Execute()

Delete Cloud Identity Provider configuration.



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
    id := "id_example" // string | Cloud Identity Provider identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CloudAzureAPI.V1CloudAzureIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudAzureAPI.V1CloudAzureIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cloud Identity Provider identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CloudAzureIdDeleteRequest struct via the builder pattern


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


## V1CloudAzureIdGet

> AzureConfiguration V1CloudAzureIdGet(ctx, id).Execute()

Get Azure Cloud Identity Provider configuration with given ID.



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
    id := "id_example" // string | Cloud Identity Provider identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudAzureAPI.V1CloudAzureIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudAzureAPI.V1CloudAzureIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1CloudAzureIdGet`: AzureConfiguration
    fmt.Fprintf(os.Stdout, "Response from `CloudAzureAPI.V1CloudAzureIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cloud Identity Provider identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CloudAzureIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AzureConfiguration**](AzureConfiguration.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CloudAzureIdPut

> AzureConfiguration V1CloudAzureIdPut(ctx, id).AzureConfigurationUpdate(azureConfigurationUpdate).Execute()

Update Azure Cloud Identity Provider configuration



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
    id := "id_example" // string | Cloud Identity Provider identifier
    azureConfigurationUpdate := *openapiclient.NewAzureConfigurationUpdate(*openapiclient.NewCloudIdPCommon("1001", "Cloud Identity Provider", "PROVIDER"), *openapiclient.NewAzureServerConfigurationUpdate("1001", true, *openapiclient.NewAzureMappings("id", "userPrincipalName", "displayName", "mail", "department", "companyName", "officeLocation", "mobilePhone", "jobTitle", "id", "displayName"), int32(30), false, "userPrincipalName", false)) // AzureConfigurationUpdate | Azure Cloud Identity Provider configuration to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudAzureAPI.V1CloudAzureIdPut(context.Background(), id).AzureConfigurationUpdate(azureConfigurationUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudAzureAPI.V1CloudAzureIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1CloudAzureIdPut`: AzureConfiguration
    fmt.Fprintf(os.Stdout, "Response from `CloudAzureAPI.V1CloudAzureIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cloud Identity Provider identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CloudAzureIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **azureConfigurationUpdate** | [**AzureConfigurationUpdate**](AzureConfigurationUpdate.md) | Azure Cloud Identity Provider configuration to update | 

### Return type

[**AzureConfiguration**](AzureConfiguration.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CloudAzurePost

> HrefResponse V1CloudAzurePost(ctx).AzureConfigurationRequest(azureConfigurationRequest).Execute()

Create Azure Cloud Identity Provider configuration



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
    azureConfigurationRequest := *openapiclient.NewAzureConfigurationRequest(*openapiclient.NewCloudIdPCommonRequest("Cloud Identity Provider", "PROVIDER"), *openapiclient.NewAzureServerConfigurationRequest("db65d325-0350-4a17-9af9-b302d0fc386b", true, *openapiclient.NewAzureMappings("id", "userPrincipalName", "displayName", "mail", "department", "companyName", "officeLocation", "mobilePhone", "jobTitle", "id", "displayName"), int32(30), false, "userPrincipalName", false, "auth")) // AzureConfigurationRequest | Azure Cloud Identity Provider configuration to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudAzureAPI.V1CloudAzurePost(context.Background()).AzureConfigurationRequest(azureConfigurationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudAzureAPI.V1CloudAzurePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1CloudAzurePost`: HrefResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudAzureAPI.V1CloudAzurePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1CloudAzurePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **azureConfigurationRequest** | [**AzureConfigurationRequest**](AzureConfigurationRequest.md) | Azure Cloud Identity Provider configuration to create | 

### Return type

[**HrefResponse**](HrefResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

