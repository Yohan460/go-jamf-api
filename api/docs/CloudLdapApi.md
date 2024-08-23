# \CloudLdapAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1LdapKeystoreVerifyPost**](CloudLdapAPI.md#V1LdapKeystoreVerifyPost) | **Post** /v1/ldap-keystore/verify | Validate keystore for Cloud Identity Provider secure connection
[**V2CloudLdapsDefaultsProviderMappingsGet**](CloudLdapAPI.md#V2CloudLdapsDefaultsProviderMappingsGet) | **Get** /v2/cloud-ldaps/defaults/{provider}/mappings | Get default mappings
[**V2CloudLdapsDefaultsProviderServerConfigurationGet**](CloudLdapAPI.md#V2CloudLdapsDefaultsProviderServerConfigurationGet) | **Get** /v2/cloud-ldaps/defaults/{provider}/server-configuration | Get default server configuration
[**V2CloudLdapsIdConnectionBindGet**](CloudLdapAPI.md#V2CloudLdapsIdConnectionBindGet) | **Get** /v2/cloud-ldaps/{id}/connection/bind | Get bind connection pool statistics
[**V2CloudLdapsIdConnectionSearchGet**](CloudLdapAPI.md#V2CloudLdapsIdConnectionSearchGet) | **Get** /v2/cloud-ldaps/{id}/connection/search | Get search connection pool statistics
[**V2CloudLdapsIdConnectionStatusGet**](CloudLdapAPI.md#V2CloudLdapsIdConnectionStatusGet) | **Get** /v2/cloud-ldaps/{id}/connection/status | Tests the communication with the specified cloud connection 
[**V2CloudLdapsIdDelete**](CloudLdapAPI.md#V2CloudLdapsIdDelete) | **Delete** /v2/cloud-ldaps/{id} | Delete Cloud Identity Provider configuration.
[**V2CloudLdapsIdGet**](CloudLdapAPI.md#V2CloudLdapsIdGet) | **Get** /v2/cloud-ldaps/{id} | Get Cloud Identity Provider configuration with given id.
[**V2CloudLdapsIdMappingsGet**](CloudLdapAPI.md#V2CloudLdapsIdMappingsGet) | **Get** /v2/cloud-ldaps/{id}/mappings | Get mappings configurations for Cloud Identity Providers server configuration.
[**V2CloudLdapsIdMappingsPut**](CloudLdapAPI.md#V2CloudLdapsIdMappingsPut) | **Put** /v2/cloud-ldaps/{id}/mappings | Update Cloud Identity Provider mappings configuration.
[**V2CloudLdapsIdPut**](CloudLdapAPI.md#V2CloudLdapsIdPut) | **Put** /v2/cloud-ldaps/{id} | Update Cloud Identity Provider configuration
[**V2CloudLdapsPost**](CloudLdapAPI.md#V2CloudLdapsPost) | **Post** /v2/cloud-ldaps | Create Cloud Identity Provider configuration



## V1LdapKeystoreVerifyPost

> CloudLdapKeystore V1LdapKeystoreVerifyPost(ctx).CloudLdapKeystoreFile(cloudLdapKeystoreFile).Execute()

Validate keystore for Cloud Identity Provider secure connection



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
	cloudLdapKeystoreFile := *openapiclient.NewCloudLdapKeystoreFile("***", string([B@7ff2b8d2), "keystore.p12") // CloudLdapKeystoreFile | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudLdapAPI.V1LdapKeystoreVerifyPost(context.Background()).CloudLdapKeystoreFile(cloudLdapKeystoreFile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudLdapAPI.V1LdapKeystoreVerifyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1LdapKeystoreVerifyPost`: CloudLdapKeystore
	fmt.Fprintf(os.Stdout, "Response from `CloudLdapAPI.V1LdapKeystoreVerifyPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1LdapKeystoreVerifyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudLdapKeystoreFile** | [**CloudLdapKeystoreFile**](CloudLdapKeystoreFile.md) |  | 

### Return type

[**CloudLdapKeystore**](CloudLdapKeystore.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2CloudLdapsDefaultsProviderMappingsGet

> CloudLdapMappingsResponse V2CloudLdapsDefaultsProviderMappingsGet(ctx, provider).Execute()

Get default mappings



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
	provider := "google" // string | Cloud Identity Provider name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudLdapAPI.V2CloudLdapsDefaultsProviderMappingsGet(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudLdapAPI.V2CloudLdapsDefaultsProviderMappingsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2CloudLdapsDefaultsProviderMappingsGet`: CloudLdapMappingsResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudLdapAPI.V2CloudLdapsDefaultsProviderMappingsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Cloud Identity Provider name | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2CloudLdapsDefaultsProviderMappingsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudLdapMappingsResponse**](CloudLdapMappingsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2CloudLdapsDefaultsProviderServerConfigurationGet

> CloudLdapServerResponse V2CloudLdapsDefaultsProviderServerConfigurationGet(ctx, provider).Execute()

Get default server configuration



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
	provider := "google" // string | Cloud Identity Provider name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudLdapAPI.V2CloudLdapsDefaultsProviderServerConfigurationGet(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudLdapAPI.V2CloudLdapsDefaultsProviderServerConfigurationGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2CloudLdapsDefaultsProviderServerConfigurationGet`: CloudLdapServerResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudLdapAPI.V2CloudLdapsDefaultsProviderServerConfigurationGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Cloud Identity Provider name | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2CloudLdapsDefaultsProviderServerConfigurationGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudLdapServerResponse**](CloudLdapServerResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2CloudLdapsIdConnectionBindGet

> CloudLdapConnectionPoolStatistics V2CloudLdapsIdConnectionBindGet(ctx, id).Execute()

Get bind connection pool statistics



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
	id := "id_example" // string | Cloud Identity Provider identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudLdapAPI.V2CloudLdapsIdConnectionBindGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudLdapAPI.V2CloudLdapsIdConnectionBindGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2CloudLdapsIdConnectionBindGet`: CloudLdapConnectionPoolStatistics
	fmt.Fprintf(os.Stdout, "Response from `CloudLdapAPI.V2CloudLdapsIdConnectionBindGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cloud Identity Provider identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2CloudLdapsIdConnectionBindGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudLdapConnectionPoolStatistics**](CloudLdapConnectionPoolStatistics.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2CloudLdapsIdConnectionSearchGet

> CloudLdapConnectionPoolStatistics V2CloudLdapsIdConnectionSearchGet(ctx, id).Execute()

Get search connection pool statistics



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
	id := "id_example" // string | Cloud Identity Provider identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudLdapAPI.V2CloudLdapsIdConnectionSearchGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudLdapAPI.V2CloudLdapsIdConnectionSearchGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2CloudLdapsIdConnectionSearchGet`: CloudLdapConnectionPoolStatistics
	fmt.Fprintf(os.Stdout, "Response from `CloudLdapAPI.V2CloudLdapsIdConnectionSearchGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cloud Identity Provider identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2CloudLdapsIdConnectionSearchGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudLdapConnectionPoolStatistics**](CloudLdapConnectionPoolStatistics.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2CloudLdapsIdConnectionStatusGet

> CloudLdapConnectionStatus V2CloudLdapsIdConnectionStatusGet(ctx, id).Execute()

Tests the communication with the specified cloud connection 



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
	id := "id_example" // string | Cloud Identity Provider identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudLdapAPI.V2CloudLdapsIdConnectionStatusGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudLdapAPI.V2CloudLdapsIdConnectionStatusGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2CloudLdapsIdConnectionStatusGet`: CloudLdapConnectionStatus
	fmt.Fprintf(os.Stdout, "Response from `CloudLdapAPI.V2CloudLdapsIdConnectionStatusGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cloud Identity Provider identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2CloudLdapsIdConnectionStatusGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudLdapConnectionStatus**](CloudLdapConnectionStatus.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2CloudLdapsIdDelete

> V2CloudLdapsIdDelete(ctx, id).Execute()

Delete Cloud Identity Provider configuration.



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
	id := "id_example" // string | Cloud Identity Provider identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CloudLdapAPI.V2CloudLdapsIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudLdapAPI.V2CloudLdapsIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV2CloudLdapsIdDeleteRequest struct via the builder pattern


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


## V2CloudLdapsIdGet

> LdapConfigurationResponse V2CloudLdapsIdGet(ctx, id).Execute()

Get Cloud Identity Provider configuration with given id.



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
	id := "id_example" // string | Cloud Identity Provider identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudLdapAPI.V2CloudLdapsIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudLdapAPI.V2CloudLdapsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2CloudLdapsIdGet`: LdapConfigurationResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudLdapAPI.V2CloudLdapsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cloud Identity Provider identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2CloudLdapsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LdapConfigurationResponse**](LdapConfigurationResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2CloudLdapsIdMappingsGet

> CloudLdapMappingsResponse V2CloudLdapsIdMappingsGet(ctx, id).Execute()

Get mappings configurations for Cloud Identity Providers server configuration.



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
	id := "id_example" // string | Cloud Identity Provider identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudLdapAPI.V2CloudLdapsIdMappingsGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudLdapAPI.V2CloudLdapsIdMappingsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2CloudLdapsIdMappingsGet`: CloudLdapMappingsResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudLdapAPI.V2CloudLdapsIdMappingsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cloud Identity Provider identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2CloudLdapsIdMappingsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudLdapMappingsResponse**](CloudLdapMappingsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2CloudLdapsIdMappingsPut

> CloudLdapMappingsResponse V2CloudLdapsIdMappingsPut(ctx, id).CloudLdapMappingsRequest(cloudLdapMappingsRequest).Execute()

Update Cloud Identity Provider mappings configuration.



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
	id := "id_example" // string | Cloud Identity Provider identifier
	cloudLdapMappingsRequest := *openapiclient.NewCloudLdapMappingsRequest(*openapiclient.NewUserMappings("ANY_OBJECT_CLASSES", "inetOrgPerson", "ou=Users", "ALL_SUBTREES", "mail", "uid", "displayName", "mail", "departmentNumber", "Building_example", "Room_example", "Phone_example", "title", "uid"), *openapiclient.NewGroupMappings("ANY_OBJECT_CLASSES", "groupOfNames", "ou=Groups", "ALL_SUBTREES", "cn", "cn", "gidNumber"), *openapiclient.NewMembershipMappings("memberOf")) // CloudLdapMappingsRequest | Cloud Identity Provider mappings to update.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudLdapAPI.V2CloudLdapsIdMappingsPut(context.Background(), id).CloudLdapMappingsRequest(cloudLdapMappingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudLdapAPI.V2CloudLdapsIdMappingsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2CloudLdapsIdMappingsPut`: CloudLdapMappingsResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudLdapAPI.V2CloudLdapsIdMappingsPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cloud Identity Provider identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2CloudLdapsIdMappingsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudLdapMappingsRequest** | [**CloudLdapMappingsRequest**](CloudLdapMappingsRequest.md) | Cloud Identity Provider mappings to update. | 

### Return type

[**CloudLdapMappingsResponse**](CloudLdapMappingsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2CloudLdapsIdPut

> LdapConfigurationResponse V2CloudLdapsIdPut(ctx, id).LdapConfigurationUpdate(ldapConfigurationUpdate).Execute()

Update Cloud Identity Provider configuration



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
	id := "id_example" // string | Cloud Identity Provider identifier
	ldapConfigurationUpdate := *openapiclient.NewLdapConfigurationUpdate(*openapiclient.NewCloudIdPCommon("1001", "Cloud Identity Provider", "PROVIDER"), *openapiclient.NewCloudLdapServerUpdate("ldap.google.com", true, "jamf.com", int64(636), int64(15), int64(60), true, "LDAPS")) // LdapConfigurationUpdate | Cloud Identity Provider configuration to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudLdapAPI.V2CloudLdapsIdPut(context.Background(), id).LdapConfigurationUpdate(ldapConfigurationUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudLdapAPI.V2CloudLdapsIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2CloudLdapsIdPut`: LdapConfigurationResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudLdapAPI.V2CloudLdapsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cloud Identity Provider identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2CloudLdapsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ldapConfigurationUpdate** | [**LdapConfigurationUpdate**](LdapConfigurationUpdate.md) | Cloud Identity Provider configuration to update | 

### Return type

[**LdapConfigurationResponse**](LdapConfigurationResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2CloudLdapsPost

> HrefResponse V2CloudLdapsPost(ctx).LdapConfigurationRequest(ldapConfigurationRequest).Execute()

Create Cloud Identity Provider configuration



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
	ldapConfigurationRequest := *openapiclient.NewLdapConfigurationRequest(*openapiclient.NewCloudIdPCommonRequest("Cloud Identity Provider", "PROVIDER"), *openapiclient.NewCloudLdapServerRequest("ldap.google.com", true, "jamf.com", int64(636), *openapiclient.NewCloudLdapKeystoreFile("***", string([B@7ff2b8d2), "keystore.p12"), int64(15), int64(60), true, "LDAPS")) // LdapConfigurationRequest | Cloud Identity Provider configuration to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudLdapAPI.V2CloudLdapsPost(context.Background()).LdapConfigurationRequest(ldapConfigurationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudLdapAPI.V2CloudLdapsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2CloudLdapsPost`: HrefResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudLdapAPI.V2CloudLdapsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2CloudLdapsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ldapConfigurationRequest** | [**LdapConfigurationRequest**](LdapConfigurationRequest.md) | Cloud Identity Provider configuration to create | 

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

