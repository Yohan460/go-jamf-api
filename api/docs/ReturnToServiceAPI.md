# \ReturnToServiceAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ReturnToServiceGet**](ReturnToServiceAPI.md#V1ReturnToServiceGet) | **Get** /v1/return-to-service | Get all Return to Service Configurations 
[**V1ReturnToServiceIdDelete**](ReturnToServiceAPI.md#V1ReturnToServiceIdDelete) | **Delete** /v1/return-to-service/{id} | Delete a Return To Service Configuration with the supplied id 
[**V1ReturnToServiceIdGet**](ReturnToServiceAPI.md#V1ReturnToServiceIdGet) | **Get** /v1/return-to-service/{id} | Retrieve a Return to Service Configuration with the supplied id 
[**V1ReturnToServiceIdPut**](ReturnToServiceAPI.md#V1ReturnToServiceIdPut) | **Put** /v1/return-to-service/{id} | Update a Return to Service Configuration 
[**V1ReturnToServicePost**](ReturnToServiceAPI.md#V1ReturnToServicePost) | **Post** /v1/return-to-service | Create a Return to Service Configuration 



## V1ReturnToServiceGet

> ReturnToServiceConfigurationSearchResults V1ReturnToServiceGet(ctx).Execute()

Get all Return to Service Configurations 



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
	resp, r, err := apiClient.ReturnToServiceAPI.V1ReturnToServiceGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReturnToServiceAPI.V1ReturnToServiceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ReturnToServiceGet`: ReturnToServiceConfigurationSearchResults
	fmt.Fprintf(os.Stdout, "Response from `ReturnToServiceAPI.V1ReturnToServiceGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1ReturnToServiceGetRequest struct via the builder pattern


### Return type

[**ReturnToServiceConfigurationSearchResults**](ReturnToServiceConfigurationSearchResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ReturnToServiceIdDelete

> V1ReturnToServiceIdDelete(ctx, id).Execute()

Delete a Return To Service Configuration with the supplied id 



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
	id := "id_example" // string | Return To Service Configurations identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ReturnToServiceAPI.V1ReturnToServiceIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReturnToServiceAPI.V1ReturnToServiceIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Return To Service Configurations identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ReturnToServiceIdDeleteRequest struct via the builder pattern


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


## V1ReturnToServiceIdGet

> ReturnToServiceConfiguration V1ReturnToServiceIdGet(ctx, id).Execute()

Retrieve a Return to Service Configuration with the supplied id 



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
	id := "id_example" // string | Return to Service Configuration identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReturnToServiceAPI.V1ReturnToServiceIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReturnToServiceAPI.V1ReturnToServiceIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ReturnToServiceIdGet`: ReturnToServiceConfiguration
	fmt.Fprintf(os.Stdout, "Response from `ReturnToServiceAPI.V1ReturnToServiceIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Return to Service Configuration identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ReturnToServiceIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReturnToServiceConfiguration**](ReturnToServiceConfiguration.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ReturnToServiceIdPut

> ReturnToServiceConfiguration V1ReturnToServiceIdPut(ctx, id).ReturnToServiceConfigurationRequest(returnToServiceConfigurationRequest).Execute()

Update a Return to Service Configuration 



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
	id := "id_example" // string | Return to Service Configuration identifier
	returnToServiceConfigurationRequest := *openapiclient.NewReturnToServiceConfigurationRequest() // ReturnToServiceConfigurationRequest | Return to Service Configuration to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReturnToServiceAPI.V1ReturnToServiceIdPut(context.Background(), id).ReturnToServiceConfigurationRequest(returnToServiceConfigurationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReturnToServiceAPI.V1ReturnToServiceIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ReturnToServiceIdPut`: ReturnToServiceConfiguration
	fmt.Fprintf(os.Stdout, "Response from `ReturnToServiceAPI.V1ReturnToServiceIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Return to Service Configuration identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ReturnToServiceIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **returnToServiceConfigurationRequest** | [**ReturnToServiceConfigurationRequest**](ReturnToServiceConfigurationRequest.md) | Return to Service Configuration to update | 

### Return type

[**ReturnToServiceConfiguration**](ReturnToServiceConfiguration.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ReturnToServicePost

> HrefResponse V1ReturnToServicePost(ctx).ReturnToServiceConfigurationRequest(returnToServiceConfigurationRequest).Execute()

Create a Return to Service Configuration 



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
	returnToServiceConfigurationRequest := *openapiclient.NewReturnToServiceConfigurationRequest() // ReturnToServiceConfigurationRequest | Return to Service Configuration to create. ids defined in this body will be ignored

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReturnToServiceAPI.V1ReturnToServicePost(context.Background()).ReturnToServiceConfigurationRequest(returnToServiceConfigurationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReturnToServiceAPI.V1ReturnToServicePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ReturnToServicePost`: HrefResponse
	fmt.Fprintf(os.Stdout, "Response from `ReturnToServiceAPI.V1ReturnToServicePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ReturnToServicePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **returnToServiceConfigurationRequest** | [**ReturnToServiceConfigurationRequest**](ReturnToServiceConfigurationRequest.md) | Return to Service Configuration to create. ids defined in this body will be ignored | 

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

