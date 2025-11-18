# \EnrollmentCustomizationAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
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
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc  (optional) (default to {"id:asc"})

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
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;date:desc,name:asc  | [default to {&quot;id:asc&quot;}]

### Return type

[**EnrollmentCustomizationSearchResultsV2**](EnrollmentCustomizationSearchResultsV2.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

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

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

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

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

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
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property,asc/desc. Default sort order is descending. Multiple sort criteria are supported and must be entered on separate lines in Swagger UI. In the URI the 'sort' query param is duplicated for each sort criterion, e.g., ...&sort=name%2Casc&sort=date%2Cdesc (optional) (default to {"date:desc"})

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
 **sort** | **[]string** | Sorting criteria in the format: property,asc/desc. Default sort order is descending. Multiple sort criteria are supported and must be entered on separate lines in Swagger UI. In the URI the &#39;sort&#39; query param is duplicated for each sort criterion, e.g., ...&amp;sort&#x3D;name%2Casc&amp;sort&#x3D;date%2Cdesc | [default to {&quot;date:desc&quot;}]

### Return type

[**HistorySearchResults**](HistorySearchResults.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

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

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

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

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

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

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

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

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

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

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

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

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

