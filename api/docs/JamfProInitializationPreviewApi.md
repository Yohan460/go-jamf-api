# \JamfProInitializationPreviewAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SystemInitializeDatabaseConnectionPost**](JamfProInitializationPreviewAPI.md#SystemInitializeDatabaseConnectionPost) | **Post** /system/initialize-database-connection | Provide Database Password during startup 
[**SystemInitializePost**](JamfProInitializationPreviewAPI.md#SystemInitializePost) | **Post** /system/initialize | Set up fresh installed Jamf Pro Server 



## SystemInitializeDatabaseConnectionPost

> SystemInitializeDatabaseConnectionPost(ctx).DatabasePassword(databasePassword).Execute()

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
	r, err := apiClient.JamfProInitializationPreviewAPI.SystemInitializeDatabaseConnectionPost(context.Background()).DatabasePassword(databasePassword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JamfProInitializationPreviewAPI.SystemInitializeDatabaseConnectionPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSystemInitializeDatabaseConnectionPostRequest struct via the builder pattern


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


## SystemInitializePost

> SystemInitializePost(ctx).Initialize(initialize).Execute()

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
	initialize := *openapiclient.NewInitialize("VFAB-YDAB-DFAB-UDAB-DEAB-EFAB-ABAB-DEAB", "Jamf", false, "admin", "12345", "https://jamf.jamfcloud.com") // Initialize | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.JamfProInitializationPreviewAPI.SystemInitializePost(context.Background()).Initialize(initialize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JamfProInitializationPreviewAPI.SystemInitializePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSystemInitializePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **initialize** | [**Initialize**](Initialize.md) |  | 

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

