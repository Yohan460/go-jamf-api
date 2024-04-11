# \EnrollmentCustomizationAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnrollmentCustomizationGet**](EnrollmentCustomizationAPI.md#V1EnrollmentCustomizationGet) | **Get** /v1/enrollment-customization | Retrieve sorted and paged Enrollment Customizations 
[**V1EnrollmentCustomizationIdDelete**](EnrollmentCustomizationAPI.md#V1EnrollmentCustomizationIdDelete) | **Delete** /v1/enrollment-customization/{id} | Delete an Enrollment Customization with the supplied id 
[**V1EnrollmentCustomizationIdGet**](EnrollmentCustomizationAPI.md#V1EnrollmentCustomizationIdGet) | **Get** /v1/enrollment-customization/{id} | Retrieve an Enrollment Customization with the supplied id 
[**V1EnrollmentCustomizationIdHistoryGet**](EnrollmentCustomizationAPI.md#V1EnrollmentCustomizationIdHistoryGet) | **Get** /v1/enrollment-customization/{id}/history | Get sorted and paged Enrollment Customization history objects 
[**V1EnrollmentCustomizationIdHistoryPost**](EnrollmentCustomizationAPI.md#V1EnrollmentCustomizationIdHistoryPost) | **Post** /v1/enrollment-customization/{id}/history | Add Enrollment Customization history object notes 
[**V1EnrollmentCustomizationIdPrestagesGet**](EnrollmentCustomizationAPI.md#V1EnrollmentCustomizationIdPrestagesGet) | **Get** /v1/enrollment-customization/{id}/prestages | Retrieve the list of Prestages using this Enrollment Customization 
[**V1EnrollmentCustomizationIdPut**](EnrollmentCustomizationAPI.md#V1EnrollmentCustomizationIdPut) | **Put** /v1/enrollment-customization/{id} | Update an Enrollment Customization 
[**V1EnrollmentCustomizationImagesPost**](EnrollmentCustomizationAPI.md#V1EnrollmentCustomizationImagesPost) | **Post** /v1/enrollment-customization/images | Upload an image
[**V1EnrollmentCustomizationPost**](EnrollmentCustomizationAPI.md#V1EnrollmentCustomizationPost) | **Post** /v1/enrollment-customization | Create an Enrollment Customization 
[**V2EnrollmentCustomizationsGet**](EnrollmentCustomizationAPI.md#V2EnrollmentCustomizationsGet) | **Get** /v2/enrollment-customizations | Retrieve sorted and paged Enrollment Customizations 
[**V2EnrollmentCustomizationsIdDelete**](EnrollmentCustomizationAPI.md#V2EnrollmentCustomizationsIdDelete) | **Delete** /v2/enrollment-customizations/{id} | Delete an Enrollment Customization with the supplied id 
[**V2EnrollmentCustomizationsIdGet**](EnrollmentCustomizationAPI.md#V2EnrollmentCustomizationsIdGet) | **Get** /v2/enrollment-customizations/{id} | Retrieve an Enrollment Customization with the supplied id 
[**V2EnrollmentCustomizationsIdHistoryGet**](EnrollmentCustomizationAPI.md#V2EnrollmentCustomizationsIdHistoryGet) | **Get** /v2/enrollment-customizations/{id}/history | Get sorted and paged Enrollment Customization history objects 
[**V2EnrollmentCustomizationsIdHistoryPost**](EnrollmentCustomizationAPI.md#V2EnrollmentCustomizationsIdHistoryPost) | **Post** /v2/enrollment-customizations/{id}/history | Add Enrollment Customization history object notes 
[**V2EnrollmentCustomizationsIdPrestagesGet**](EnrollmentCustomizationAPI.md#V2EnrollmentCustomizationsIdPrestagesGet) | **Get** /v2/enrollment-customizations/{id}/prestages | Retrieve the list of Prestages using this Enrollment Customization 
[**V2EnrollmentCustomizationsIdPut**](EnrollmentCustomizationAPI.md#V2EnrollmentCustomizationsIdPut) | **Put** /v2/enrollment-customizations/{id} | Update an Enrollment Customization 
[**V2EnrollmentCustomizationsImagesIdGet**](EnrollmentCustomizationAPI.md#V2EnrollmentCustomizationsImagesIdGet) | **Get** /v2/enrollment-customizations/images/{id} | Download an enrollment customization image 
[**V2EnrollmentCustomizationsImagesPost**](EnrollmentCustomizationAPI.md#V2EnrollmentCustomizationsImagesPost) | **Post** /v2/enrollment-customizations/images | Upload an image
[**V2EnrollmentCustomizationsPost**](EnrollmentCustomizationAPI.md#V2EnrollmentCustomizationsPost) | **Post** /v2/enrollment-customizations | Create an Enrollment Customization 



## V1EnrollmentCustomizationGet

> EnrollmentCustomizationSearchResults V1EnrollmentCustomizationGet(ctx).Page(page).Size(size).Pagesize(pagesize).PageSize(pageSize).Sort(sort).Execute()

Retrieve sorted and paged Enrollment Customizations 



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
	size := int64(56) // int64 |  (optional) (default to 100)
	pagesize := int64(56) // int64 |  (optional) (default to 100)
	pageSize := int64(56) // int64 |  (optional) (default to 100)
	sort := "sort_example" // string | Sorting criteria in the format: property:asc/desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc  (optional) (default to "id:asc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnrollmentCustomizationAPI.V1EnrollmentCustomizationGet(context.Background()).Page(page).Size(size).Pagesize(pagesize).PageSize(pageSize).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentCustomizationAPI.V1EnrollmentCustomizationGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1EnrollmentCustomizationGet`: EnrollmentCustomizationSearchResults
	fmt.Fprintf(os.Stdout, "Response from `EnrollmentCustomizationAPI.V1EnrollmentCustomizationGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1EnrollmentCustomizationGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | [default to 0]
 **size** | **int64** |  | [default to 100]
 **pagesize** | **int64** |  | [default to 100]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **string** | Sorting criteria in the format: property:asc/desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;date:desc,name:asc  | [default to &quot;id:asc&quot;]

### Return type

[**EnrollmentCustomizationSearchResults**](EnrollmentCustomizationSearchResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EnrollmentCustomizationIdDelete

> V1EnrollmentCustomizationIdDelete(ctx, id).Execute()

Delete an Enrollment Customization with the supplied id 



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
	id := int64(56) // int64 | Enrollment Customization identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EnrollmentCustomizationAPI.V1EnrollmentCustomizationIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentCustomizationAPI.V1EnrollmentCustomizationIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Enrollment Customization identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnrollmentCustomizationIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EnrollmentCustomizationIdGet

> GetEnrollmentCustomization V1EnrollmentCustomizationIdGet(ctx, id).Execute()

Retrieve an Enrollment Customization with the supplied id 



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
	id := int64(56) // int64 | Enrollment Customization identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnrollmentCustomizationAPI.V1EnrollmentCustomizationIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentCustomizationAPI.V1EnrollmentCustomizationIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1EnrollmentCustomizationIdGet`: GetEnrollmentCustomization
	fmt.Fprintf(os.Stdout, "Response from `EnrollmentCustomizationAPI.V1EnrollmentCustomizationIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Enrollment Customization identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnrollmentCustomizationIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetEnrollmentCustomization**](GetEnrollmentCustomization.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EnrollmentCustomizationIdHistoryGet

> HistorySearchResults V1EnrollmentCustomizationIdHistoryGet(ctx, id).Page(page).Size(size).Pagesize(pagesize).PageSize(pageSize).Sort(sort).Execute()

Get sorted and paged Enrollment Customization history objects 



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
	id := int64(56) // int64 | Enrollment Customization identifier
	page := int64(56) // int64 |  (optional) (default to 0)
	size := int64(56) // int64 |  (optional) (default to 100)
	pagesize := int64(56) // int64 |  (optional) (default to 100)
	pageSize := int64(56) // int64 |  (optional) (default to 100)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property,asc/desc. Default sort order is descending. Multiple sort criteria are supported and must be entered on separate lines in Swagger UI. In the URI the 'sort' query param is duplicated for each sort criterion, e.g., ...&sort=name%2Casc&sort=date%2Cdesc (optional) (default to ["date:desc"])

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnrollmentCustomizationAPI.V1EnrollmentCustomizationIdHistoryGet(context.Background(), id).Page(page).Size(size).Pagesize(pagesize).PageSize(pageSize).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentCustomizationAPI.V1EnrollmentCustomizationIdHistoryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1EnrollmentCustomizationIdHistoryGet`: HistorySearchResults
	fmt.Fprintf(os.Stdout, "Response from `EnrollmentCustomizationAPI.V1EnrollmentCustomizationIdHistoryGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Enrollment Customization identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnrollmentCustomizationIdHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** |  | [default to 0]
 **size** | **int64** |  | [default to 100]
 **pagesize** | **int64** |  | [default to 100]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property,asc/desc. Default sort order is descending. Multiple sort criteria are supported and must be entered on separate lines in Swagger UI. In the URI the &#39;sort&#39; query param is duplicated for each sort criterion, e.g., ...&amp;sort&#x3D;name%2Casc&amp;sort&#x3D;date%2Cdesc | [default to [&quot;date:desc&quot;]]

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


## V1EnrollmentCustomizationIdHistoryPost

> ObjectHistory V1EnrollmentCustomizationIdHistoryPost(ctx, id).ObjectHistoryNote(objectHistoryNote).Execute()

Add Enrollment Customization history object notes 



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
	id := int64(56) // int64 | Enrollment Customization identifier
	objectHistoryNote := *openapiclient.NewObjectHistoryNote("A generic note can sometimes be useful, but generally not.") // ObjectHistoryNote | History notes to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnrollmentCustomizationAPI.V1EnrollmentCustomizationIdHistoryPost(context.Background(), id).ObjectHistoryNote(objectHistoryNote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentCustomizationAPI.V1EnrollmentCustomizationIdHistoryPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1EnrollmentCustomizationIdHistoryPost`: ObjectHistory
	fmt.Fprintf(os.Stdout, "Response from `EnrollmentCustomizationAPI.V1EnrollmentCustomizationIdHistoryPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Enrollment Customization identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnrollmentCustomizationIdHistoryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **objectHistoryNote** | [**ObjectHistoryNote**](ObjectHistoryNote.md) | History notes to create | 

### Return type

[**ObjectHistory**](ObjectHistory.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EnrollmentCustomizationIdPrestagesGet

> PrestageDependencies V1EnrollmentCustomizationIdPrestagesGet(ctx, id).Execute()

Retrieve the list of Prestages using this Enrollment Customization 



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
	id := int64(56) // int64 | Enrollment Customization identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnrollmentCustomizationAPI.V1EnrollmentCustomizationIdPrestagesGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentCustomizationAPI.V1EnrollmentCustomizationIdPrestagesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1EnrollmentCustomizationIdPrestagesGet`: PrestageDependencies
	fmt.Fprintf(os.Stdout, "Response from `EnrollmentCustomizationAPI.V1EnrollmentCustomizationIdPrestagesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Enrollment Customization identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnrollmentCustomizationIdPrestagesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PrestageDependencies**](PrestageDependencies.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EnrollmentCustomizationIdPut

> GetEnrollmentCustomization V1EnrollmentCustomizationIdPut(ctx, id).EnrollmentCustomization(enrollmentCustomization).Execute()

Update an Enrollment Customization 



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
	id := int64(56) // int64 | Enrollment Customization identifier
	enrollmentCustomization := *openapiclient.NewEnrollmentCustomization(int64(2), "Example", "Example description", *openapiclient.NewEnrollmentCustomizationBrandingSettings("0000FF", "0000FF", "0000FF", "0000FF", "https://jamfUrl/api/v2/enrollment-customizations/images/1")) // EnrollmentCustomization | Enrollment Customization to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnrollmentCustomizationAPI.V1EnrollmentCustomizationIdPut(context.Background(), id).EnrollmentCustomization(enrollmentCustomization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentCustomizationAPI.V1EnrollmentCustomizationIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1EnrollmentCustomizationIdPut`: GetEnrollmentCustomization
	fmt.Fprintf(os.Stdout, "Response from `EnrollmentCustomizationAPI.V1EnrollmentCustomizationIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Enrollment Customization identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnrollmentCustomizationIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enrollmentCustomization** | [**EnrollmentCustomization**](EnrollmentCustomization.md) | Enrollment Customization to update | 

### Return type

[**GetEnrollmentCustomization**](GetEnrollmentCustomization.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EnrollmentCustomizationImagesPost

> BrandingImageUrl V1EnrollmentCustomizationImagesPost(ctx).File(file).Execute()

Upload an image



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
	file := os.NewFile(1234, "some_file") // *os.File | The file to upload

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnrollmentCustomizationAPI.V1EnrollmentCustomizationImagesPost(context.Background()).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentCustomizationAPI.V1EnrollmentCustomizationImagesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1EnrollmentCustomizationImagesPost`: BrandingImageUrl
	fmt.Fprintf(os.Stdout, "Response from `EnrollmentCustomizationAPI.V1EnrollmentCustomizationImagesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1EnrollmentCustomizationImagesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | The file to upload | 

### Return type

[**BrandingImageUrl**](BrandingImageUrl.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EnrollmentCustomizationPost

> GetEnrollmentCustomization V1EnrollmentCustomizationPost(ctx).EnrollmentCustomization(enrollmentCustomization).Execute()

Create an Enrollment Customization 



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
	enrollmentCustomization := *openapiclient.NewEnrollmentCustomization(int64(2), "Example", "Example description", *openapiclient.NewEnrollmentCustomizationBrandingSettings("0000FF", "0000FF", "0000FF", "0000FF", "https://jamfUrl/api/v2/enrollment-customizations/images/1")) // EnrollmentCustomization | Enrollment customization to create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnrollmentCustomizationAPI.V1EnrollmentCustomizationPost(context.Background()).EnrollmentCustomization(enrollmentCustomization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentCustomizationAPI.V1EnrollmentCustomizationPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1EnrollmentCustomizationPost`: GetEnrollmentCustomization
	fmt.Fprintf(os.Stdout, "Response from `EnrollmentCustomizationAPI.V1EnrollmentCustomizationPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1EnrollmentCustomizationPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enrollmentCustomization** | [**EnrollmentCustomization**](EnrollmentCustomization.md) | Enrollment customization to create. | 

### Return type

[**GetEnrollmentCustomization**](GetEnrollmentCustomization.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2EnrollmentCustomizationsGet

> EnrollmentCustomizationSearchResultsV2 V2EnrollmentCustomizationsGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Execute()

Retrieve sorted and paged Enrollment Customizations 



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
	resp, r, err := apiClient.EnrollmentCustomizationAPI.V2EnrollmentCustomizationsGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentCustomizationAPI.V2EnrollmentCustomizationsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2EnrollmentCustomizationsGet`: EnrollmentCustomizationSearchResultsV2
	fmt.Fprintf(os.Stdout, "Response from `EnrollmentCustomizationAPI.V2EnrollmentCustomizationsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2EnrollmentCustomizationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;date:desc,name:asc  | [default to [&quot;id:asc&quot;]]

### Return type

[**EnrollmentCustomizationSearchResultsV2**](EnrollmentCustomizationSearchResultsV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2EnrollmentCustomizationsIdDelete

> V2EnrollmentCustomizationsIdDelete(ctx, id).Execute()

Delete an Enrollment Customization with the supplied id 



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
	id := "id_example" // string | Enrollment Customization identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EnrollmentCustomizationAPI.V2EnrollmentCustomizationsIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentCustomizationAPI.V2EnrollmentCustomizationsIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Enrollment Customization identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2EnrollmentCustomizationsIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2EnrollmentCustomizationsIdGet

> EnrollmentCustomizationV2 V2EnrollmentCustomizationsIdGet(ctx, id).Execute()

Retrieve an Enrollment Customization with the supplied id 



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
	id := "id_example" // string | Enrollment Customization identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnrollmentCustomizationAPI.V2EnrollmentCustomizationsIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentCustomizationAPI.V2EnrollmentCustomizationsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2EnrollmentCustomizationsIdGet`: EnrollmentCustomizationV2
	fmt.Fprintf(os.Stdout, "Response from `EnrollmentCustomizationAPI.V2EnrollmentCustomizationsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Enrollment Customization identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2EnrollmentCustomizationsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnrollmentCustomizationV2**](EnrollmentCustomizationV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2EnrollmentCustomizationsIdHistoryGet

> HistorySearchResults V2EnrollmentCustomizationsIdHistoryGet(ctx, id).Page(page).PageSize(pageSize).Sort(sort).Execute()

Get sorted and paged Enrollment Customization history objects 



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
	id := "id_example" // string | Enrollment Customization identifier
	page := int64(56) // int64 |  (optional) (default to 0)
	pageSize := int64(56) // int64 |  (optional) (default to 100)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property,asc/desc. Default sort order is descending. Multiple sort criteria are supported and must be entered on separate lines in Swagger UI. In the URI the 'sort' query param is duplicated for each sort criterion, e.g., ...&sort=name%2Casc&sort=date%2Cdesc (optional) (default to ["date:desc"])

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnrollmentCustomizationAPI.V2EnrollmentCustomizationsIdHistoryGet(context.Background(), id).Page(page).PageSize(pageSize).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentCustomizationAPI.V2EnrollmentCustomizationsIdHistoryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2EnrollmentCustomizationsIdHistoryGet`: HistorySearchResults
	fmt.Fprintf(os.Stdout, "Response from `EnrollmentCustomizationAPI.V2EnrollmentCustomizationsIdHistoryGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Enrollment Customization identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2EnrollmentCustomizationsIdHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property,asc/desc. Default sort order is descending. Multiple sort criteria are supported and must be entered on separate lines in Swagger UI. In the URI the &#39;sort&#39; query param is duplicated for each sort criterion, e.g., ...&amp;sort&#x3D;name%2Casc&amp;sort&#x3D;date%2Cdesc | [default to [&quot;date:desc&quot;]]

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


## V2EnrollmentCustomizationsIdHistoryPost

> ObjectHistory V2EnrollmentCustomizationsIdHistoryPost(ctx, id).ObjectHistoryNote(objectHistoryNote).Execute()

Add Enrollment Customization history object notes 



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
	id := "id_example" // string | Enrollment Customization identifier
	objectHistoryNote := *openapiclient.NewObjectHistoryNote("A generic note can sometimes be useful, but generally not.") // ObjectHistoryNote | History notes to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnrollmentCustomizationAPI.V2EnrollmentCustomizationsIdHistoryPost(context.Background(), id).ObjectHistoryNote(objectHistoryNote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentCustomizationAPI.V2EnrollmentCustomizationsIdHistoryPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2EnrollmentCustomizationsIdHistoryPost`: ObjectHistory
	fmt.Fprintf(os.Stdout, "Response from `EnrollmentCustomizationAPI.V2EnrollmentCustomizationsIdHistoryPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Enrollment Customization identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2EnrollmentCustomizationsIdHistoryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **objectHistoryNote** | [**ObjectHistoryNote**](ObjectHistoryNote.md) | History notes to create | 

### Return type

[**ObjectHistory**](ObjectHistory.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2EnrollmentCustomizationsIdPrestagesGet

> PrestageDependencies V2EnrollmentCustomizationsIdPrestagesGet(ctx, id).Execute()

Retrieve the list of Prestages using this Enrollment Customization 



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
	id := "id_example" // string | Enrollment Customization identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnrollmentCustomizationAPI.V2EnrollmentCustomizationsIdPrestagesGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentCustomizationAPI.V2EnrollmentCustomizationsIdPrestagesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2EnrollmentCustomizationsIdPrestagesGet`: PrestageDependencies
	fmt.Fprintf(os.Stdout, "Response from `EnrollmentCustomizationAPI.V2EnrollmentCustomizationsIdPrestagesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Enrollment Customization identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2EnrollmentCustomizationsIdPrestagesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PrestageDependencies**](PrestageDependencies.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2EnrollmentCustomizationsIdPut

> EnrollmentCustomizationV2 V2EnrollmentCustomizationsIdPut(ctx, id).EnrollmentCustomizationV2(enrollmentCustomizationV2).Execute()

Update an Enrollment Customization 



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
	id := "id_example" // string | Enrollment Customization identifier
	enrollmentCustomizationV2 := *openapiclient.NewEnrollmentCustomizationV2("2", "Example", "Example description", *openapiclient.NewEnrollmentCustomizationBrandingSettings("0000FF", "0000FF", "0000FF", "0000FF", "https://jamfUrl/api/v2/enrollment-customizations/images/1")) // EnrollmentCustomizationV2 | Enrollment Customization to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnrollmentCustomizationAPI.V2EnrollmentCustomizationsIdPut(context.Background(), id).EnrollmentCustomizationV2(enrollmentCustomizationV2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentCustomizationAPI.V2EnrollmentCustomizationsIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2EnrollmentCustomizationsIdPut`: EnrollmentCustomizationV2
	fmt.Fprintf(os.Stdout, "Response from `EnrollmentCustomizationAPI.V2EnrollmentCustomizationsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Enrollment Customization identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2EnrollmentCustomizationsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enrollmentCustomizationV2** | [**EnrollmentCustomizationV2**](EnrollmentCustomizationV2.md) | Enrollment Customization to update | 

### Return type

[**EnrollmentCustomizationV2**](EnrollmentCustomizationV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2EnrollmentCustomizationsImagesIdGet

> *os.File V2EnrollmentCustomizationsImagesIdGet(ctx, id).Execute()

Download an enrollment customization image 



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
	id := "id_example" // string | id of the enrollment customization image

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnrollmentCustomizationAPI.V2EnrollmentCustomizationsImagesIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentCustomizationAPI.V2EnrollmentCustomizationsImagesIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2EnrollmentCustomizationsImagesIdGet`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `EnrollmentCustomizationAPI.V2EnrollmentCustomizationsImagesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of the enrollment customization image | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2EnrollmentCustomizationsImagesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2EnrollmentCustomizationsImagesPost

> BrandingImageUrl V2EnrollmentCustomizationsImagesPost(ctx).File(file).Execute()

Upload an image



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
	file := os.NewFile(1234, "some_file") // *os.File | The file to upload

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnrollmentCustomizationAPI.V2EnrollmentCustomizationsImagesPost(context.Background()).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentCustomizationAPI.V2EnrollmentCustomizationsImagesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2EnrollmentCustomizationsImagesPost`: BrandingImageUrl
	fmt.Fprintf(os.Stdout, "Response from `EnrollmentCustomizationAPI.V2EnrollmentCustomizationsImagesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2EnrollmentCustomizationsImagesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | The file to upload | 

### Return type

[**BrandingImageUrl**](BrandingImageUrl.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2EnrollmentCustomizationsPost

> HrefResponse V2EnrollmentCustomizationsPost(ctx).EnrollmentCustomizationV2(enrollmentCustomizationV2).Execute()

Create an Enrollment Customization 



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
	enrollmentCustomizationV2 := *openapiclient.NewEnrollmentCustomizationV2("2", "Example", "Example description", *openapiclient.NewEnrollmentCustomizationBrandingSettings("0000FF", "0000FF", "0000FF", "0000FF", "https://jamfUrl/api/v2/enrollment-customizations/images/1")) // EnrollmentCustomizationV2 | Enrollment customization to create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnrollmentCustomizationAPI.V2EnrollmentCustomizationsPost(context.Background()).EnrollmentCustomizationV2(enrollmentCustomizationV2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentCustomizationAPI.V2EnrollmentCustomizationsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2EnrollmentCustomizationsPost`: HrefResponse
	fmt.Fprintf(os.Stdout, "Response from `EnrollmentCustomizationAPI.V2EnrollmentCustomizationsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2EnrollmentCustomizationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enrollmentCustomizationV2** | [**EnrollmentCustomizationV2**](EnrollmentCustomizationV2.md) | Enrollment customization to create. | 

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

