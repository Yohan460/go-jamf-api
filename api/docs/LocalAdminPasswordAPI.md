# \LocalAdminPasswordAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1LocalAdminPasswordClientManagementIdAccountUsernameAuditGet**](LocalAdminPasswordAPI.md#V1LocalAdminPasswordClientManagementIdAccountUsernameAuditGet) | **Get** /v1/local-admin-password/{clientManagementId}/account/{username}/audit | Get LAPS password viewed history.
[**V1LocalAdminPasswordClientManagementIdAccountUsernamePasswordGet**](LocalAdminPasswordAPI.md#V1LocalAdminPasswordClientManagementIdAccountUsernamePasswordGet) | **Get** /v1/local-admin-password/{clientManagementId}/account/{username}/password | Get current LAPS password for specified username on a client.
[**V1LocalAdminPasswordClientManagementIdAccountsGet**](LocalAdminPasswordAPI.md#V1LocalAdminPasswordClientManagementIdAccountsGet) | **Get** /v1/local-admin-password/{clientManagementId}/accounts | Get the LAPS capable admin accounts for a device.
[**V1LocalAdminPasswordClientManagementIdSetPasswordPut**](LocalAdminPasswordAPI.md#V1LocalAdminPasswordClientManagementIdSetPasswordPut) | **Put** /v1/local-admin-password/{clientManagementId}/set-password | Set the LAPS password for a device.
[**V1LocalAdminPasswordSettingsGet**](LocalAdminPasswordAPI.md#V1LocalAdminPasswordSettingsGet) | **Get** /v1/local-admin-password/settings | Get the current LAPS settings.
[**V1LocalAdminPasswordSettingsPut**](LocalAdminPasswordAPI.md#V1LocalAdminPasswordSettingsPut) | **Put** /v1/local-admin-password/settings | Update settings for LAPS.
[**V2LocalAdminPasswordClientManagementIdAccountUsernameAuditGet**](LocalAdminPasswordAPI.md#V2LocalAdminPasswordClientManagementIdAccountUsernameAuditGet) | **Get** /v2/local-admin-password/{clientManagementId}/account/{username}/audit | Get LAPS password viewed history.
[**V2LocalAdminPasswordClientManagementIdAccountUsernameHistoryGet**](LocalAdminPasswordAPI.md#V2LocalAdminPasswordClientManagementIdAccountUsernameHistoryGet) | **Get** /v2/local-admin-password/{clientManagementId}/account/{username}/history | Get LAPS historical records for target device and username.
[**V2LocalAdminPasswordClientManagementIdAccountUsernamePasswordGet**](LocalAdminPasswordAPI.md#V2LocalAdminPasswordClientManagementIdAccountUsernamePasswordGet) | **Get** /v2/local-admin-password/{clientManagementId}/account/{username}/password | Get current LAPS password for specified username on a client.
[**V2LocalAdminPasswordClientManagementIdAccountsGet**](LocalAdminPasswordAPI.md#V2LocalAdminPasswordClientManagementIdAccountsGet) | **Get** /v2/local-admin-password/{clientManagementId}/accounts | Get the LAPS capable admin accounts for a device.
[**V2LocalAdminPasswordClientManagementIdSetPasswordPut**](LocalAdminPasswordAPI.md#V2LocalAdminPasswordClientManagementIdSetPasswordPut) | **Put** /v2/local-admin-password/{clientManagementId}/set-password | Set the LAPS password for a device.
[**V2LocalAdminPasswordPendingRotationsGet**](LocalAdminPasswordAPI.md#V2LocalAdminPasswordPendingRotationsGet) | **Get** /v2/local-admin-password/pending-rotations | Get a list of the current devices and usernames with pending LAPS rotations
[**V2LocalAdminPasswordSettingsGet**](LocalAdminPasswordAPI.md#V2LocalAdminPasswordSettingsGet) | **Get** /v2/local-admin-password/settings | Get the current LAPS settings.
[**V2LocalAdminPasswordSettingsPut**](LocalAdminPasswordAPI.md#V2LocalAdminPasswordSettingsPut) | **Put** /v2/local-admin-password/settings | Update settings for LAPS.



## V1LocalAdminPasswordClientManagementIdAccountUsernameAuditGet

> LapsPasswordAuditsResults V1LocalAdminPasswordClientManagementIdAccountUsernameAuditGet(ctx, clientManagementId, username).Execute()

Get LAPS password viewed history.



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
	clientManagementId := "clientManagementId_example" // string | client management id of target device.
	username := "username_example" // string | user name to view audit information for

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocalAdminPasswordAPI.V1LocalAdminPasswordClientManagementIdAccountUsernameAuditGet(context.Background(), clientManagementId, username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocalAdminPasswordAPI.V1LocalAdminPasswordClientManagementIdAccountUsernameAuditGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1LocalAdminPasswordClientManagementIdAccountUsernameAuditGet`: LapsPasswordAuditsResults
	fmt.Fprintf(os.Stdout, "Response from `LocalAdminPasswordAPI.V1LocalAdminPasswordClientManagementIdAccountUsernameAuditGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientManagementId** | **string** | client management id of target device. | 
**username** | **string** | user name to view audit information for | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1LocalAdminPasswordClientManagementIdAccountUsernameAuditGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**LapsPasswordAuditsResults**](LapsPasswordAuditsResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1LocalAdminPasswordClientManagementIdAccountUsernamePasswordGet

> LapsPasswordResponse V1LocalAdminPasswordClientManagementIdAccountUsernamePasswordGet(ctx, clientManagementId, username).Execute()

Get current LAPS password for specified username on a client.



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
	clientManagementId := "clientManagementId_example" // string | client management id of target device.
	username := "username_example" // string | user name for the account

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocalAdminPasswordAPI.V1LocalAdminPasswordClientManagementIdAccountUsernamePasswordGet(context.Background(), clientManagementId, username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocalAdminPasswordAPI.V1LocalAdminPasswordClientManagementIdAccountUsernamePasswordGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1LocalAdminPasswordClientManagementIdAccountUsernamePasswordGet`: LapsPasswordResponse
	fmt.Fprintf(os.Stdout, "Response from `LocalAdminPasswordAPI.V1LocalAdminPasswordClientManagementIdAccountUsernamePasswordGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientManagementId** | **string** | client management id of target device. | 
**username** | **string** | user name for the account | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1LocalAdminPasswordClientManagementIdAccountUsernamePasswordGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**LapsPasswordResponse**](LapsPasswordResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1LocalAdminPasswordClientManagementIdAccountsGet

> LapsUserResults V1LocalAdminPasswordClientManagementIdAccountsGet(ctx, clientManagementId).Execute()

Get the LAPS capable admin accounts for a device.



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
	clientManagementId := "clientManagementId_example" // string | client management id of target device.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocalAdminPasswordAPI.V1LocalAdminPasswordClientManagementIdAccountsGet(context.Background(), clientManagementId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocalAdminPasswordAPI.V1LocalAdminPasswordClientManagementIdAccountsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1LocalAdminPasswordClientManagementIdAccountsGet`: LapsUserResults
	fmt.Fprintf(os.Stdout, "Response from `LocalAdminPasswordAPI.V1LocalAdminPasswordClientManagementIdAccountsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientManagementId** | **string** | client management id of target device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1LocalAdminPasswordClientManagementIdAccountsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LapsUserResults**](LapsUserResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1LocalAdminPasswordClientManagementIdSetPasswordPut

> LapsUserPasswordResponse V1LocalAdminPasswordClientManagementIdSetPasswordPut(ctx, clientManagementId).LapsUserPasswordRequest(lapsUserPasswordRequest).Execute()

Set the LAPS password for a device.



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
	clientManagementId := "clientManagementId_example" // string | client management id of target device.
	lapsUserPasswordRequest := *openapiclient.NewLapsUserPasswordRequest() // LapsUserPasswordRequest | LAPS password to set

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocalAdminPasswordAPI.V1LocalAdminPasswordClientManagementIdSetPasswordPut(context.Background(), clientManagementId).LapsUserPasswordRequest(lapsUserPasswordRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocalAdminPasswordAPI.V1LocalAdminPasswordClientManagementIdSetPasswordPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1LocalAdminPasswordClientManagementIdSetPasswordPut`: LapsUserPasswordResponse
	fmt.Fprintf(os.Stdout, "Response from `LocalAdminPasswordAPI.V1LocalAdminPasswordClientManagementIdSetPasswordPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientManagementId** | **string** | client management id of target device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1LocalAdminPasswordClientManagementIdSetPasswordPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lapsUserPasswordRequest** | [**LapsUserPasswordRequest**](LapsUserPasswordRequest.md) | LAPS password to set | 

### Return type

[**LapsUserPasswordResponse**](LapsUserPasswordResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1LocalAdminPasswordSettingsGet

> LapsSettingsResponse V1LocalAdminPasswordSettingsGet(ctx).Execute()

Get the current LAPS settings.



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
	resp, r, err := apiClient.LocalAdminPasswordAPI.V1LocalAdminPasswordSettingsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocalAdminPasswordAPI.V1LocalAdminPasswordSettingsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1LocalAdminPasswordSettingsGet`: LapsSettingsResponse
	fmt.Fprintf(os.Stdout, "Response from `LocalAdminPasswordAPI.V1LocalAdminPasswordSettingsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1LocalAdminPasswordSettingsGetRequest struct via the builder pattern


### Return type

[**LapsSettingsResponse**](LapsSettingsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1LocalAdminPasswordSettingsPut

> LapsSettingsResponse V1LocalAdminPasswordSettingsPut(ctx).LapsSettingsRequest(lapsSettingsRequest).Execute()

Update settings for LAPS.



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
	lapsSettingsRequest := *openapiclient.NewLapsSettingsRequest(false, int64(3600), int64(7776000)) // LapsSettingsRequest | LAPS settings to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocalAdminPasswordAPI.V1LocalAdminPasswordSettingsPut(context.Background()).LapsSettingsRequest(lapsSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocalAdminPasswordAPI.V1LocalAdminPasswordSettingsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1LocalAdminPasswordSettingsPut`: LapsSettingsResponse
	fmt.Fprintf(os.Stdout, "Response from `LocalAdminPasswordAPI.V1LocalAdminPasswordSettingsPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1LocalAdminPasswordSettingsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lapsSettingsRequest** | [**LapsSettingsRequest**](LapsSettingsRequest.md) | LAPS settings to update | 

### Return type

[**LapsSettingsResponse**](LapsSettingsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2LocalAdminPasswordClientManagementIdAccountUsernameAuditGet

> LapsPasswordAuditsResultsV2 V2LocalAdminPasswordClientManagementIdAccountUsernameAuditGet(ctx, clientManagementId, username).Execute()

Get LAPS password viewed history.



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
	clientManagementId := "clientManagementId_example" // string | client management id of target device.
	username := "username_example" // string | user name to view audit information for

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocalAdminPasswordAPI.V2LocalAdminPasswordClientManagementIdAccountUsernameAuditGet(context.Background(), clientManagementId, username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocalAdminPasswordAPI.V2LocalAdminPasswordClientManagementIdAccountUsernameAuditGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2LocalAdminPasswordClientManagementIdAccountUsernameAuditGet`: LapsPasswordAuditsResultsV2
	fmt.Fprintf(os.Stdout, "Response from `LocalAdminPasswordAPI.V2LocalAdminPasswordClientManagementIdAccountUsernameAuditGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientManagementId** | **string** | client management id of target device. | 
**username** | **string** | user name to view audit information for | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2LocalAdminPasswordClientManagementIdAccountUsernameAuditGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**LapsPasswordAuditsResultsV2**](LapsPasswordAuditsResultsV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2LocalAdminPasswordClientManagementIdAccountUsernameHistoryGet

> LapsHistoryResponse V2LocalAdminPasswordClientManagementIdAccountUsernameHistoryGet(ctx, clientManagementId, username).Execute()

Get LAPS historical records for target device and username.



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
	clientManagementId := "clientManagementId_example" // string | client management id of target device.
	username := "username_example" // string | user name to view history for

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocalAdminPasswordAPI.V2LocalAdminPasswordClientManagementIdAccountUsernameHistoryGet(context.Background(), clientManagementId, username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocalAdminPasswordAPI.V2LocalAdminPasswordClientManagementIdAccountUsernameHistoryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2LocalAdminPasswordClientManagementIdAccountUsernameHistoryGet`: LapsHistoryResponse
	fmt.Fprintf(os.Stdout, "Response from `LocalAdminPasswordAPI.V2LocalAdminPasswordClientManagementIdAccountUsernameHistoryGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientManagementId** | **string** | client management id of target device. | 
**username** | **string** | user name to view history for | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2LocalAdminPasswordClientManagementIdAccountUsernameHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**LapsHistoryResponse**](LapsHistoryResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2LocalAdminPasswordClientManagementIdAccountUsernamePasswordGet

> LapsPasswordResponseV2 V2LocalAdminPasswordClientManagementIdAccountUsernamePasswordGet(ctx, clientManagementId, username).Execute()

Get current LAPS password for specified username on a client.



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
	clientManagementId := "clientManagementId_example" // string | client management id of target device.
	username := "username_example" // string | user name for the account

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocalAdminPasswordAPI.V2LocalAdminPasswordClientManagementIdAccountUsernamePasswordGet(context.Background(), clientManagementId, username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocalAdminPasswordAPI.V2LocalAdminPasswordClientManagementIdAccountUsernamePasswordGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2LocalAdminPasswordClientManagementIdAccountUsernamePasswordGet`: LapsPasswordResponseV2
	fmt.Fprintf(os.Stdout, "Response from `LocalAdminPasswordAPI.V2LocalAdminPasswordClientManagementIdAccountUsernamePasswordGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientManagementId** | **string** | client management id of target device. | 
**username** | **string** | user name for the account | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2LocalAdminPasswordClientManagementIdAccountUsernamePasswordGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**LapsPasswordResponseV2**](LapsPasswordResponseV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2LocalAdminPasswordClientManagementIdAccountsGet

> LapsUserResultsV2 V2LocalAdminPasswordClientManagementIdAccountsGet(ctx, clientManagementId).Execute()

Get the LAPS capable admin accounts for a device.



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
	clientManagementId := "clientManagementId_example" // string | client management id of target device.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocalAdminPasswordAPI.V2LocalAdminPasswordClientManagementIdAccountsGet(context.Background(), clientManagementId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocalAdminPasswordAPI.V2LocalAdminPasswordClientManagementIdAccountsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2LocalAdminPasswordClientManagementIdAccountsGet`: LapsUserResultsV2
	fmt.Fprintf(os.Stdout, "Response from `LocalAdminPasswordAPI.V2LocalAdminPasswordClientManagementIdAccountsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientManagementId** | **string** | client management id of target device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2LocalAdminPasswordClientManagementIdAccountsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LapsUserResultsV2**](LapsUserResultsV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2LocalAdminPasswordClientManagementIdSetPasswordPut

> LapsUserPasswordResponseV2 V2LocalAdminPasswordClientManagementIdSetPasswordPut(ctx, clientManagementId).LapsUserPasswordRequestV2(lapsUserPasswordRequestV2).Execute()

Set the LAPS password for a device.



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
	clientManagementId := "clientManagementId_example" // string | client management id of target device.
	lapsUserPasswordRequestV2 := *openapiclient.NewLapsUserPasswordRequestV2() // LapsUserPasswordRequestV2 | LAPS password to set

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocalAdminPasswordAPI.V2LocalAdminPasswordClientManagementIdSetPasswordPut(context.Background(), clientManagementId).LapsUserPasswordRequestV2(lapsUserPasswordRequestV2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocalAdminPasswordAPI.V2LocalAdminPasswordClientManagementIdSetPasswordPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2LocalAdminPasswordClientManagementIdSetPasswordPut`: LapsUserPasswordResponseV2
	fmt.Fprintf(os.Stdout, "Response from `LocalAdminPasswordAPI.V2LocalAdminPasswordClientManagementIdSetPasswordPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientManagementId** | **string** | client management id of target device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2LocalAdminPasswordClientManagementIdSetPasswordPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lapsUserPasswordRequestV2** | [**LapsUserPasswordRequestV2**](LapsUserPasswordRequestV2.md) | LAPS password to set | 

### Return type

[**LapsUserPasswordResponseV2**](LapsUserPasswordResponseV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2LocalAdminPasswordPendingRotationsGet

> LapsPendingRotationResponse V2LocalAdminPasswordPendingRotationsGet(ctx).Execute()

Get a list of the current devices and usernames with pending LAPS rotations



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
	resp, r, err := apiClient.LocalAdminPasswordAPI.V2LocalAdminPasswordPendingRotationsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocalAdminPasswordAPI.V2LocalAdminPasswordPendingRotationsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2LocalAdminPasswordPendingRotationsGet`: LapsPendingRotationResponse
	fmt.Fprintf(os.Stdout, "Response from `LocalAdminPasswordAPI.V2LocalAdminPasswordPendingRotationsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV2LocalAdminPasswordPendingRotationsGetRequest struct via the builder pattern


### Return type

[**LapsPendingRotationResponse**](LapsPendingRotationResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2LocalAdminPasswordSettingsGet

> LapsSettingsResponseV2 V2LocalAdminPasswordSettingsGet(ctx).Execute()

Get the current LAPS settings.



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
	resp, r, err := apiClient.LocalAdminPasswordAPI.V2LocalAdminPasswordSettingsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocalAdminPasswordAPI.V2LocalAdminPasswordSettingsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2LocalAdminPasswordSettingsGet`: LapsSettingsResponseV2
	fmt.Fprintf(os.Stdout, "Response from `LocalAdminPasswordAPI.V2LocalAdminPasswordSettingsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV2LocalAdminPasswordSettingsGetRequest struct via the builder pattern


### Return type

[**LapsSettingsResponseV2**](LapsSettingsResponseV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2LocalAdminPasswordSettingsPut

> LapsSettingsResponseV2 V2LocalAdminPasswordSettingsPut(ctx).LapsSettingsRequestV2(lapsSettingsRequestV2).Execute()

Update settings for LAPS.



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
	lapsSettingsRequestV2 := *openapiclient.NewLapsSettingsRequestV2(false, int64(3600), false, int64(7776000)) // LapsSettingsRequestV2 | LAPS settings to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocalAdminPasswordAPI.V2LocalAdminPasswordSettingsPut(context.Background()).LapsSettingsRequestV2(lapsSettingsRequestV2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocalAdminPasswordAPI.V2LocalAdminPasswordSettingsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2LocalAdminPasswordSettingsPut`: LapsSettingsResponseV2
	fmt.Fprintf(os.Stdout, "Response from `LocalAdminPasswordAPI.V2LocalAdminPasswordSettingsPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2LocalAdminPasswordSettingsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lapsSettingsRequestV2** | [**LapsSettingsRequestV2**](LapsSettingsRequestV2.md) | LAPS settings to update | 

### Return type

[**LapsSettingsResponseV2**](LapsSettingsResponseV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

