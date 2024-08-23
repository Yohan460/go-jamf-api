# \LoginCustomizationAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1LoginCustomizationGet**](LoginCustomizationAPI.md#V1LoginCustomizationGet) | **Get** /v1/login-customization | Get current login disclaimer settings
[**V1LoginCustomizationPut**](LoginCustomizationAPI.md#V1LoginCustomizationPut) | **Put** /v1/login-customization | Update current login disclaimer settings.



## V1LoginCustomizationGet

> LoginContent V1LoginCustomizationGet(ctx).Execute()

Get current login disclaimer settings



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
	resp, r, err := apiClient.LoginCustomizationAPI.V1LoginCustomizationGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoginCustomizationAPI.V1LoginCustomizationGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1LoginCustomizationGet`: LoginContent
	fmt.Fprintf(os.Stdout, "Response from `LoginCustomizationAPI.V1LoginCustomizationGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1LoginCustomizationGetRequest struct via the builder pattern


### Return type

[**LoginContent**](LoginContent.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1LoginCustomizationPut

> LoginContentPut V1LoginCustomizationPut(ctx).LoginContentPut(loginContentPut).Execute()

Update current login disclaimer settings.



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
	loginContentPut := *openapiclient.NewLoginContentPut(true) // LoginContentPut | Login disclaimer settings to save.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoginCustomizationAPI.V1LoginCustomizationPut(context.Background()).LoginContentPut(loginContentPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoginCustomizationAPI.V1LoginCustomizationPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1LoginCustomizationPut`: LoginContentPut
	fmt.Fprintf(os.Stdout, "Response from `LoginCustomizationAPI.V1LoginCustomizationPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1LoginCustomizationPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **loginContentPut** | [**LoginContentPut**](LoginContentPut.md) | Login disclaimer settings to save. | 

### Return type

[**LoginContentPut**](LoginContentPut.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

