# \DeviceEnrollmentsAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1DeviceEnrollmentsGet**](DeviceEnrollmentsAPI.md#V1DeviceEnrollmentsGet) | **Get** /v1/device-enrollments | Read all sorted and paged Device Enrollment instances 
[**V1DeviceEnrollmentsIdDelete**](DeviceEnrollmentsAPI.md#V1DeviceEnrollmentsIdDelete) | **Delete** /v1/device-enrollments/{id} | Delete a Device Enrollment Instance with the supplied id 
[**V1DeviceEnrollmentsIdDisownPost**](DeviceEnrollmentsAPI.md#V1DeviceEnrollmentsIdDisownPost) | **Post** /v1/device-enrollments/{id}/disown | Disown devices from the given Device Enrollment Instance 
[**V1DeviceEnrollmentsIdGet**](DeviceEnrollmentsAPI.md#V1DeviceEnrollmentsIdGet) | **Get** /v1/device-enrollments/{id} | Retrieve a Device Enrollment Instance with the supplied id 
[**V1DeviceEnrollmentsIdHistoryGet**](DeviceEnrollmentsAPI.md#V1DeviceEnrollmentsIdHistoryGet) | **Get** /v1/device-enrollments/{id}/history | Get sorted and paged Device Enrollment history objects 
[**V1DeviceEnrollmentsIdHistoryPost**](DeviceEnrollmentsAPI.md#V1DeviceEnrollmentsIdHistoryPost) | **Post** /v1/device-enrollments/{id}/history | Add Device Enrollment history object notes 
[**V1DeviceEnrollmentsIdPut**](DeviceEnrollmentsAPI.md#V1DeviceEnrollmentsIdPut) | **Put** /v1/device-enrollments/{id} | Update a Device Enrollment Instance with the supplied id 
[**V1DeviceEnrollmentsIdSyncsGet**](DeviceEnrollmentsAPI.md#V1DeviceEnrollmentsIdSyncsGet) | **Get** /v1/device-enrollments/{id}/syncs | Get all instance sync states for a single Device Enrollment Instance 
[**V1DeviceEnrollmentsIdSyncsLatestGet**](DeviceEnrollmentsAPI.md#V1DeviceEnrollmentsIdSyncsLatestGet) | **Get** /v1/device-enrollments/{id}/syncs/latest | Get the latest sync state for a single Device Enrollment Instance 
[**V1DeviceEnrollmentsIdUploadTokenPut**](DeviceEnrollmentsAPI.md#V1DeviceEnrollmentsIdUploadTokenPut) | **Put** /v1/device-enrollments/{id}/upload-token | Update a Device Enrollment Instance with the supplied Token 
[**V1DeviceEnrollmentsPublicKeyGet**](DeviceEnrollmentsAPI.md#V1DeviceEnrollmentsPublicKeyGet) | **Get** /v1/device-enrollments/public-key | Retrieve the Jamf Pro Device Enrollment public key 
[**V1DeviceEnrollmentsSyncsGet**](DeviceEnrollmentsAPI.md#V1DeviceEnrollmentsSyncsGet) | **Get** /v1/device-enrollments/syncs | Get all instance sync states for all Device Enrollment Instances 
[**V1DeviceEnrollmentsUploadTokenPost**](DeviceEnrollmentsAPI.md#V1DeviceEnrollmentsUploadTokenPost) | **Post** /v1/device-enrollments/upload-token | Create a Device Enrollment Instance with the supplied Token 



## V1DeviceEnrollmentsGet

> DeviceEnrollmentInstanceSearchResults V1DeviceEnrollmentsGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Execute()

Read all sorted and paged Device Enrollment instances 



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
	page := int64(56) // int64 |  (optional) (default to 0)
	pageSize := int64(56) // int64 |  (optional) (default to 100)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc  (optional) (default to ["id:asc"])

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceEnrollmentsAPI.V1DeviceEnrollmentsGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceEnrollmentsAPI.V1DeviceEnrollmentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1DeviceEnrollmentsGet`: DeviceEnrollmentInstanceSearchResults
	fmt.Fprintf(os.Stdout, "Response from `DeviceEnrollmentsAPI.V1DeviceEnrollmentsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1DeviceEnrollmentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;date:desc,name:asc  | [default to [&quot;id:asc&quot;]]

### Return type

[**DeviceEnrollmentInstanceSearchResults**](DeviceEnrollmentInstanceSearchResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DeviceEnrollmentsIdDelete

> V1DeviceEnrollmentsIdDelete(ctx, id).Execute()

Delete a Device Enrollment Instance with the supplied id 



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
	id := "id_example" // string | Device Enrollment Instance identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeviceEnrollmentsAPI.V1DeviceEnrollmentsIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceEnrollmentsAPI.V1DeviceEnrollmentsIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Device Enrollment Instance identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DeviceEnrollmentsIdDeleteRequest struct via the builder pattern


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


## V1DeviceEnrollmentsIdDisownPost

> DeviceEnrollmentDisownResponse V1DeviceEnrollmentsIdDisownPost(ctx, id).DeviceEnrollmentDisownBody(deviceEnrollmentDisownBody).Execute()

Disown devices from the given Device Enrollment Instance 



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
	id := "id_example" // string | Device Enrollment Instance identifier
	deviceEnrollmentDisownBody := *openapiclient.NewDeviceEnrollmentDisownBody() // DeviceEnrollmentDisownBody | List of device serial numbers to disown (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceEnrollmentsAPI.V1DeviceEnrollmentsIdDisownPost(context.Background(), id).DeviceEnrollmentDisownBody(deviceEnrollmentDisownBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceEnrollmentsAPI.V1DeviceEnrollmentsIdDisownPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1DeviceEnrollmentsIdDisownPost`: DeviceEnrollmentDisownResponse
	fmt.Fprintf(os.Stdout, "Response from `DeviceEnrollmentsAPI.V1DeviceEnrollmentsIdDisownPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Device Enrollment Instance identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DeviceEnrollmentsIdDisownPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deviceEnrollmentDisownBody** | [**DeviceEnrollmentDisownBody**](DeviceEnrollmentDisownBody.md) | List of device serial numbers to disown | 

### Return type

[**DeviceEnrollmentDisownResponse**](DeviceEnrollmentDisownResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DeviceEnrollmentsIdGet

> DeviceEnrollmentInstance V1DeviceEnrollmentsIdGet(ctx, id).Execute()

Retrieve a Device Enrollment Instance with the supplied id 



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
	id := "id_example" // string | Device Enrollment Instance identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceEnrollmentsAPI.V1DeviceEnrollmentsIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceEnrollmentsAPI.V1DeviceEnrollmentsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1DeviceEnrollmentsIdGet`: DeviceEnrollmentInstance
	fmt.Fprintf(os.Stdout, "Response from `DeviceEnrollmentsAPI.V1DeviceEnrollmentsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Device Enrollment Instance identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DeviceEnrollmentsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeviceEnrollmentInstance**](DeviceEnrollmentInstance.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DeviceEnrollmentsIdHistoryGet

> HistorySearchResults V1DeviceEnrollmentsIdHistoryGet(ctx, id).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Get sorted and paged Device Enrollment history objects 



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
	id := "id_example" // string | Device Enrollment Instance identifier
	page := int64(56) // int64 |  (optional) (default to 0)
	pageSize := int64(56) // int64 |  (optional) (default to 100)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property,asc/desc. Default sort order is descending. Multiple sort criteria are supported and must be entered on separate lines in Swagger UI. In the URI the 'sort' query param is duplicated for each sort criterion, e.g., ...&sort=name%2Casc&sort=date%2Cdesc (optional) (default to ["date:desc"])
	filter := "filter_example" // string | Query in the RSQL format, allowing to filter history notes collection. Default search is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: search=username!=admin and details==*disabled* and date<2019-12-15 (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceEnrollmentsAPI.V1DeviceEnrollmentsIdHistoryGet(context.Background(), id).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceEnrollmentsAPI.V1DeviceEnrollmentsIdHistoryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1DeviceEnrollmentsIdHistoryGet`: HistorySearchResults
	fmt.Fprintf(os.Stdout, "Response from `DeviceEnrollmentsAPI.V1DeviceEnrollmentsIdHistoryGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Device Enrollment Instance identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DeviceEnrollmentsIdHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property,asc/desc. Default sort order is descending. Multiple sort criteria are supported and must be entered on separate lines in Swagger UI. In the URI the &#39;sort&#39; query param is duplicated for each sort criterion, e.g., ...&amp;sort&#x3D;name%2Casc&amp;sort&#x3D;date%2Cdesc | [default to [&quot;date:desc&quot;]]
 **filter** | **string** | Query in the RSQL format, allowing to filter history notes collection. Default search is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: search&#x3D;username!&#x3D;admin and details&#x3D;&#x3D;*disabled* and date&lt;2019-12-15 | [default to &quot;&quot;]

### Return type

[**HistorySearchResults**](HistorySearchResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DeviceEnrollmentsIdHistoryPost

> HrefResponse V1DeviceEnrollmentsIdHistoryPost(ctx, id).ObjectHistoryNote(objectHistoryNote).Execute()

Add Device Enrollment history object notes 



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
	id := "id_example" // string | Device Enrollment Instance identifier
	objectHistoryNote := *openapiclient.NewObjectHistoryNote("A generic note can sometimes be useful, but generally not.") // ObjectHistoryNote | History notes to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceEnrollmentsAPI.V1DeviceEnrollmentsIdHistoryPost(context.Background(), id).ObjectHistoryNote(objectHistoryNote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceEnrollmentsAPI.V1DeviceEnrollmentsIdHistoryPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1DeviceEnrollmentsIdHistoryPost`: HrefResponse
	fmt.Fprintf(os.Stdout, "Response from `DeviceEnrollmentsAPI.V1DeviceEnrollmentsIdHistoryPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Device Enrollment Instance identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DeviceEnrollmentsIdHistoryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **objectHistoryNote** | [**ObjectHistoryNote**](ObjectHistoryNote.md) | History notes to create | 

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


## V1DeviceEnrollmentsIdPut

> DeviceEnrollmentInstance V1DeviceEnrollmentsIdPut(ctx, id).DeviceEnrollmentInstance(deviceEnrollmentInstance).Execute()

Update a Device Enrollment Instance with the supplied id 



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
	id := "id_example" // string | Device Enrollment Instance identifier
	deviceEnrollmentInstance := *openapiclient.NewDeviceEnrollmentInstance("Example Device Enrollment Instance") // DeviceEnrollmentInstance | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceEnrollmentsAPI.V1DeviceEnrollmentsIdPut(context.Background(), id).DeviceEnrollmentInstance(deviceEnrollmentInstance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceEnrollmentsAPI.V1DeviceEnrollmentsIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1DeviceEnrollmentsIdPut`: DeviceEnrollmentInstance
	fmt.Fprintf(os.Stdout, "Response from `DeviceEnrollmentsAPI.V1DeviceEnrollmentsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Device Enrollment Instance identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DeviceEnrollmentsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deviceEnrollmentInstance** | [**DeviceEnrollmentInstance**](DeviceEnrollmentInstance.md) |  | 

### Return type

[**DeviceEnrollmentInstance**](DeviceEnrollmentInstance.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DeviceEnrollmentsIdSyncsGet

> []DeviceEnrollmentInstanceSyncStatus V1DeviceEnrollmentsIdSyncsGet(ctx, id).Execute()

Get all instance sync states for a single Device Enrollment Instance 



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
	id := "id_example" // string | Device Enrollment Instance identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceEnrollmentsAPI.V1DeviceEnrollmentsIdSyncsGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceEnrollmentsAPI.V1DeviceEnrollmentsIdSyncsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1DeviceEnrollmentsIdSyncsGet`: []DeviceEnrollmentInstanceSyncStatus
	fmt.Fprintf(os.Stdout, "Response from `DeviceEnrollmentsAPI.V1DeviceEnrollmentsIdSyncsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Device Enrollment Instance identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DeviceEnrollmentsIdSyncsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]DeviceEnrollmentInstanceSyncStatus**](DeviceEnrollmentInstanceSyncStatus.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DeviceEnrollmentsIdSyncsLatestGet

> DeviceEnrollmentInstanceSyncStatus V1DeviceEnrollmentsIdSyncsLatestGet(ctx, id).Execute()

Get the latest sync state for a single Device Enrollment Instance 



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
	id := "id_example" // string | Device Enrollment Instance identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceEnrollmentsAPI.V1DeviceEnrollmentsIdSyncsLatestGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceEnrollmentsAPI.V1DeviceEnrollmentsIdSyncsLatestGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1DeviceEnrollmentsIdSyncsLatestGet`: DeviceEnrollmentInstanceSyncStatus
	fmt.Fprintf(os.Stdout, "Response from `DeviceEnrollmentsAPI.V1DeviceEnrollmentsIdSyncsLatestGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Device Enrollment Instance identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DeviceEnrollmentsIdSyncsLatestGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeviceEnrollmentInstanceSyncStatus**](DeviceEnrollmentInstanceSyncStatus.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DeviceEnrollmentsIdUploadTokenPut

> DeviceEnrollmentInstance V1DeviceEnrollmentsIdUploadTokenPut(ctx, id).DeviceEnrollmentToken(deviceEnrollmentToken).Execute()

Update a Device Enrollment Instance with the supplied Token 



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
	id := "id_example" // string | Device Enrollment Instance identifier
	deviceEnrollmentToken := *openapiclient.NewDeviceEnrollmentToken() // DeviceEnrollmentToken | The downloaded token base 64 encoded from the MDM server to be used to create a new Device Enrollment Instance.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceEnrollmentsAPI.V1DeviceEnrollmentsIdUploadTokenPut(context.Background(), id).DeviceEnrollmentToken(deviceEnrollmentToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceEnrollmentsAPI.V1DeviceEnrollmentsIdUploadTokenPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1DeviceEnrollmentsIdUploadTokenPut`: DeviceEnrollmentInstance
	fmt.Fprintf(os.Stdout, "Response from `DeviceEnrollmentsAPI.V1DeviceEnrollmentsIdUploadTokenPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Device Enrollment Instance identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DeviceEnrollmentsIdUploadTokenPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deviceEnrollmentToken** | [**DeviceEnrollmentToken**](DeviceEnrollmentToken.md) | The downloaded token base 64 encoded from the MDM server to be used to create a new Device Enrollment Instance. | 

### Return type

[**DeviceEnrollmentInstance**](DeviceEnrollmentInstance.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DeviceEnrollmentsPublicKeyGet

> *os.File V1DeviceEnrollmentsPublicKeyGet(ctx).Execute()

Retrieve the Jamf Pro Device Enrollment public key 



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
	resp, r, err := apiClient.DeviceEnrollmentsAPI.V1DeviceEnrollmentsPublicKeyGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceEnrollmentsAPI.V1DeviceEnrollmentsPublicKeyGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1DeviceEnrollmentsPublicKeyGet`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DeviceEnrollmentsAPI.V1DeviceEnrollmentsPublicKeyGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1DeviceEnrollmentsPublicKeyGetRequest struct via the builder pattern


### Return type

[***os.File**](*os.File.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-pem-file

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DeviceEnrollmentsSyncsGet

> []DeviceEnrollmentInstanceSyncStatus V1DeviceEnrollmentsSyncsGet(ctx).Execute()

Get all instance sync states for all Device Enrollment Instances 



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
	resp, r, err := apiClient.DeviceEnrollmentsAPI.V1DeviceEnrollmentsSyncsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceEnrollmentsAPI.V1DeviceEnrollmentsSyncsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1DeviceEnrollmentsSyncsGet`: []DeviceEnrollmentInstanceSyncStatus
	fmt.Fprintf(os.Stdout, "Response from `DeviceEnrollmentsAPI.V1DeviceEnrollmentsSyncsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1DeviceEnrollmentsSyncsGetRequest struct via the builder pattern


### Return type

[**[]DeviceEnrollmentInstanceSyncStatus**](DeviceEnrollmentInstanceSyncStatus.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DeviceEnrollmentsUploadTokenPost

> HrefResponse V1DeviceEnrollmentsUploadTokenPost(ctx).DeviceEnrollmentToken(deviceEnrollmentToken).Execute()

Create a Device Enrollment Instance with the supplied Token 



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
	deviceEnrollmentToken := *openapiclient.NewDeviceEnrollmentToken() // DeviceEnrollmentToken | The downloaded token base 64 encoded from the MDM server to be used to create a new Device Enrollment Instance.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceEnrollmentsAPI.V1DeviceEnrollmentsUploadTokenPost(context.Background()).DeviceEnrollmentToken(deviceEnrollmentToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceEnrollmentsAPI.V1DeviceEnrollmentsUploadTokenPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1DeviceEnrollmentsUploadTokenPost`: HrefResponse
	fmt.Fprintf(os.Stdout, "Response from `DeviceEnrollmentsAPI.V1DeviceEnrollmentsUploadTokenPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1DeviceEnrollmentsUploadTokenPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceEnrollmentToken** | [**DeviceEnrollmentToken**](DeviceEnrollmentToken.md) | The downloaded token base 64 encoded from the MDM server to be used to create a new Device Enrollment Instance. | 

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

