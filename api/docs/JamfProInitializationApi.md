# \JamfProInitializationAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1SystemInitializeDatabaseConnectionPost**](JamfProInitializationAPI.md#V1SystemInitializeDatabaseConnectionPost) | **Post** /v1/system/initialize-database-connection | Provide Database Password during startup 
[**V1SystemInitializePost**](JamfProInitializationAPI.md#V1SystemInitializePost) | **Post** /v1/system/initialize | Set up fresh installed Jamf Pro Server 



## V1SystemInitializeDatabaseConnectionPost

> V1SystemInitializeDatabaseConnectionPost(ctx).DatabasePassword(databasePassword).Execute()

Provide Database Password during startup 



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
	databasePassword := *openapiclient.NewDatabasePassword("12345") // DatabasePassword | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.JamfProInitializationAPI.V1SystemInitializeDatabaseConnectionPost(context.Background()).DatabasePassword(databasePassword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JamfProInitializationAPI.V1SystemInitializeDatabaseConnectionPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1SystemInitializeDatabaseConnectionPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **databasePassword** | [**DatabasePassword**](DatabasePassword.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SystemInitializePost

> V1SystemInitializePost(ctx).InitializeV1(initializeV1).Execute()

Set up fresh installed Jamf Pro Server 



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
	initializeV1 := *openapiclient.NewInitializeV1("VFAB-YDAB-DFAB-UDAB-DEAB-EFAB-ABAB-DEAB", "Jamf", false, "admin", "12345", "https://jamf.jamfcloud.com") // InitializeV1 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.JamfProInitializationAPI.V1SystemInitializePost(context.Background()).InitializeV1(initializeV1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JamfProInitializationAPI.V1SystemInitializePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1SystemInitializePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **initializeV1** | [**InitializeV1**](InitializeV1.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

