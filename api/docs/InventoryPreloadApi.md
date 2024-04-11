# \InventoryPreloadAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InventoryPreloadCsvTemplateGet**](InventoryPreloadAPI.md#InventoryPreloadCsvTemplateGet) | **Get** /inventory-preload/csv-template | Get the Inventory Preload CSV template 
[**InventoryPreloadDelete**](InventoryPreloadAPI.md#InventoryPreloadDelete) | **Delete** /inventory-preload | Delete all Inventory Preload records 
[**InventoryPreloadGet**](InventoryPreloadAPI.md#InventoryPreloadGet) | **Get** /inventory-preload | Return all Inventory Preload records 
[**InventoryPreloadHistoryGet**](InventoryPreloadAPI.md#InventoryPreloadHistoryGet) | **Get** /inventory-preload/history | Get Inventory Preload history entries 
[**InventoryPreloadHistoryNotesPost**](InventoryPreloadAPI.md#InventoryPreloadHistoryNotesPost) | **Post** /inventory-preload/history/notes | Add Inventory Preload history object notes 
[**InventoryPreloadIdDelete**](InventoryPreloadAPI.md#InventoryPreloadIdDelete) | **Delete** /inventory-preload/{id} | Delete an Inventory Preload record 
[**InventoryPreloadIdGet**](InventoryPreloadAPI.md#InventoryPreloadIdGet) | **Get** /inventory-preload/{id} | Get an Inventory Preload record 
[**InventoryPreloadIdPut**](InventoryPreloadAPI.md#InventoryPreloadIdPut) | **Put** /inventory-preload/{id} | Update an Inventory Preload record 
[**InventoryPreloadPost**](InventoryPreloadAPI.md#InventoryPreloadPost) | **Post** /inventory-preload | Create a new Inventory Preload record using JSON or CSV 
[**InventoryPreloadValidateCsvPost**](InventoryPreloadAPI.md#InventoryPreloadValidateCsvPost) | **Post** /inventory-preload/validate-csv | Validate a given CSV file 
[**V1InventoryPreloadCsvTemplateGet**](InventoryPreloadAPI.md#V1InventoryPreloadCsvTemplateGet) | **Get** /v1/inventory-preload/csv-template | Retrieve the Inventory Preload CSV template 
[**V1InventoryPreloadDelete**](InventoryPreloadAPI.md#V1InventoryPreloadDelete) | **Delete** /v1/inventory-preload | Delete all Inventory Preload records 
[**V1InventoryPreloadGet**](InventoryPreloadAPI.md#V1InventoryPreloadGet) | **Get** /v1/inventory-preload | Return all Inventory Preload records 
[**V1InventoryPreloadHistoryGet**](InventoryPreloadAPI.md#V1InventoryPreloadHistoryGet) | **Get** /v1/inventory-preload/history | Get Inventory Preload history entries 
[**V1InventoryPreloadHistoryPost**](InventoryPreloadAPI.md#V1InventoryPreloadHistoryPost) | **Post** /v1/inventory-preload/history | Add Inventory Preload history object notes 
[**V1InventoryPreloadIdDelete**](InventoryPreloadAPI.md#V1InventoryPreloadIdDelete) | **Delete** /v1/inventory-preload/{id} | Delete an Inventory Preload record 
[**V1InventoryPreloadIdGet**](InventoryPreloadAPI.md#V1InventoryPreloadIdGet) | **Get** /v1/inventory-preload/{id} | Get an Inventory Preload record 
[**V1InventoryPreloadIdPut**](InventoryPreloadAPI.md#V1InventoryPreloadIdPut) | **Put** /v1/inventory-preload/{id} | Update an Inventory Preload record 
[**V1InventoryPreloadPost**](InventoryPreloadAPI.md#V1InventoryPreloadPost) | **Post** /v1/inventory-preload | Create a new Inventory Preload record using JSON or CSV 
[**V1InventoryPreloadValidateCsvPost**](InventoryPreloadAPI.md#V1InventoryPreloadValidateCsvPost) | **Post** /v1/inventory-preload/validate-csv | Validate a given CSV file 
[**V2InventoryPreloadCsvGet**](InventoryPreloadAPI.md#V2InventoryPreloadCsvGet) | **Get** /v2/inventory-preload/csv | Download all Inventory Preload records
[**V2InventoryPreloadCsvPost**](InventoryPreloadAPI.md#V2InventoryPreloadCsvPost) | **Post** /v2/inventory-preload/csv | Create one or more new Inventory Preload records using CSV 
[**V2InventoryPreloadCsvTemplateGet**](InventoryPreloadAPI.md#V2InventoryPreloadCsvTemplateGet) | **Get** /v2/inventory-preload/csv-template | Download the Inventory Preload CSV template
[**V2InventoryPreloadCsvValidatePost**](InventoryPreloadAPI.md#V2InventoryPreloadCsvValidatePost) | **Post** /v2/inventory-preload/csv-validate | Validate a given CSV file 
[**V2InventoryPreloadEaColumnsGet**](InventoryPreloadAPI.md#V2InventoryPreloadEaColumnsGet) | **Get** /v2/inventory-preload/ea-columns | Retrieve a list of extension attribute columns 
[**V2InventoryPreloadExportPost**](InventoryPreloadAPI.md#V2InventoryPreloadExportPost) | **Post** /v2/inventory-preload/export | Export a collection of inventory preload records 
[**V2InventoryPreloadHistoryGet**](InventoryPreloadAPI.md#V2InventoryPreloadHistoryGet) | **Get** /v2/inventory-preload/history | Get Inventory Preload history entries 
[**V2InventoryPreloadHistoryPost**](InventoryPreloadAPI.md#V2InventoryPreloadHistoryPost) | **Post** /v2/inventory-preload/history | Add Inventory Preload history object notes
[**V2InventoryPreloadRecordsDeleteAllPost**](InventoryPreloadAPI.md#V2InventoryPreloadRecordsDeleteAllPost) | **Post** /v2/inventory-preload/records/delete-all | Delete all Inventory Preload records 
[**V2InventoryPreloadRecordsGet**](InventoryPreloadAPI.md#V2InventoryPreloadRecordsGet) | **Get** /v2/inventory-preload/records | Return all Inventory Preload records
[**V2InventoryPreloadRecordsIdDelete**](InventoryPreloadAPI.md#V2InventoryPreloadRecordsIdDelete) | **Delete** /v2/inventory-preload/records/{id} | Delete an Inventory Preload record 
[**V2InventoryPreloadRecordsIdGet**](InventoryPreloadAPI.md#V2InventoryPreloadRecordsIdGet) | **Get** /v2/inventory-preload/records/{id} | Get an Inventory Preload record
[**V2InventoryPreloadRecordsIdPut**](InventoryPreloadAPI.md#V2InventoryPreloadRecordsIdPut) | **Put** /v2/inventory-preload/records/{id} | Update an Inventory Preload record
[**V2InventoryPreloadRecordsPost**](InventoryPreloadAPI.md#V2InventoryPreloadRecordsPost) | **Post** /v2/inventory-preload/records | Create a new Inventory Preload record using JSON



## InventoryPreloadCsvTemplateGet

> map[string]interface{} InventoryPreloadCsvTemplateGet(ctx).Execute()

Get the Inventory Preload CSV template 



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
	resp, r, err := apiClient.InventoryPreloadAPI.InventoryPreloadCsvTemplateGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryPreloadAPI.InventoryPreloadCsvTemplateGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryPreloadCsvTemplateGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `InventoryPreloadAPI.InventoryPreloadCsvTemplateGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryPreloadCsvTemplateGetRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryPreloadDelete

> InventoryPreloadDelete(ctx).Execute()

Delete all Inventory Preload records 



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
	r, err := apiClient.InventoryPreloadAPI.InventoryPreloadDelete(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryPreloadAPI.InventoryPreloadDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryPreloadDeleteRequest struct via the builder pattern


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


## InventoryPreloadGet

> []InventoryPreloadRecordSearchResults InventoryPreloadGet(ctx).Page(page).Pagesize(pagesize).Sort(sort).SortBy(sortBy).Execute()

Return all Inventory Preload records 



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
	pagesize := int64(56) // int64 |  (optional) (default to 100)
	sort := "sort_example" // string |  (optional) (default to "ASC")
	sortBy := "sortBy_example" // string |  (optional) (default to "id")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryPreloadAPI.InventoryPreloadGet(context.Background()).Page(page).Pagesize(pagesize).Sort(sort).SortBy(sortBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryPreloadAPI.InventoryPreloadGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryPreloadGet`: []InventoryPreloadRecordSearchResults
	fmt.Fprintf(os.Stdout, "Response from `InventoryPreloadAPI.InventoryPreloadGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInventoryPreloadGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | [default to 0]
 **pagesize** | **int64** |  | [default to 100]
 **sort** | **string** |  | [default to &quot;ASC&quot;]
 **sortBy** | **string** |  | [default to &quot;id&quot;]

### Return type

[**[]InventoryPreloadRecordSearchResults**](InventoryPreloadRecordSearchResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryPreloadHistoryGet

> HistorySearchResults InventoryPreloadHistoryGet(ctx).Page(page).Size(size).Pagesize(pagesize).PageSize(pageSize).Sort(sort).Execute()

Get Inventory Preload history entries 



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
	sort := "sort_example" // string | Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc  (optional) (default to "date:desc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryPreloadAPI.InventoryPreloadHistoryGet(context.Background()).Page(page).Size(size).Pagesize(pagesize).PageSize(pageSize).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryPreloadAPI.InventoryPreloadHistoryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryPreloadHistoryGet`: HistorySearchResults
	fmt.Fprintf(os.Stdout, "Response from `InventoryPreloadAPI.InventoryPreloadHistoryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInventoryPreloadHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | [default to 0]
 **size** | **int64** |  | [default to 100]
 **pagesize** | **int64** |  | [default to 100]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **string** | Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;date:desc,name:asc  | [default to &quot;date:desc&quot;]

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


## InventoryPreloadHistoryNotesPost

> ObjectHistory InventoryPreloadHistoryNotesPost(ctx).ObjectHistoryNote(objectHistoryNote).Execute()

Add Inventory Preload history object notes 



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
	objectHistoryNote := *openapiclient.NewObjectHistoryNote("A generic note can sometimes be useful, but generally not.") // ObjectHistoryNote | History notes to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryPreloadAPI.InventoryPreloadHistoryNotesPost(context.Background()).ObjectHistoryNote(objectHistoryNote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryPreloadAPI.InventoryPreloadHistoryNotesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryPreloadHistoryNotesPost`: ObjectHistory
	fmt.Fprintf(os.Stdout, "Response from `InventoryPreloadAPI.InventoryPreloadHistoryNotesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInventoryPreloadHistoryNotesPostRequest struct via the builder pattern


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


## InventoryPreloadIdDelete

> InventoryPreloadIdDelete(ctx, id).Execute()

Delete an Inventory Preload record 



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
	id := int64(56) // int64 | Inventory Preload identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InventoryPreloadAPI.InventoryPreloadIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryPreloadAPI.InventoryPreloadIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Inventory Preload identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryPreloadIdDeleteRequest struct via the builder pattern


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


## InventoryPreloadIdGet

> InventoryPreloadRecord InventoryPreloadIdGet(ctx, id).Execute()

Get an Inventory Preload record 



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
	id := int64(56) // int64 | Inventory Preload identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryPreloadAPI.InventoryPreloadIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryPreloadAPI.InventoryPreloadIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryPreloadIdGet`: InventoryPreloadRecord
	fmt.Fprintf(os.Stdout, "Response from `InventoryPreloadAPI.InventoryPreloadIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Inventory Preload identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryPreloadIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InventoryPreloadRecord**](InventoryPreloadRecord.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryPreloadIdPut

> InventoryPreloadRecord InventoryPreloadIdPut(ctx, id).InventoryPreloadRecord(inventoryPreloadRecord).Execute()

Update an Inventory Preload record 



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
	id := int64(56) // int64 | Inventory Preload identifier
	inventoryPreloadRecord := *openapiclient.NewInventoryPreloadRecord("C02L29ECF8J1", "Computer") // InventoryPreloadRecord | Inventory Preload record to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryPreloadAPI.InventoryPreloadIdPut(context.Background(), id).InventoryPreloadRecord(inventoryPreloadRecord).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryPreloadAPI.InventoryPreloadIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryPreloadIdPut`: InventoryPreloadRecord
	fmt.Fprintf(os.Stdout, "Response from `InventoryPreloadAPI.InventoryPreloadIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Inventory Preload identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryPreloadIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inventoryPreloadRecord** | [**InventoryPreloadRecord**](InventoryPreloadRecord.md) | Inventory Preload record to update | 

### Return type

[**InventoryPreloadRecord**](InventoryPreloadRecord.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryPreloadPost

> InventoryPreloadRecord InventoryPreloadPost(ctx).InventoryPreloadRecord(inventoryPreloadRecord).Execute()

Create a new Inventory Preload record using JSON or CSV 



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
	inventoryPreloadRecord := *openapiclient.NewInventoryPreloadRecord("C02L29ECF8J1", "Computer") // InventoryPreloadRecord | Inventory Preload record or records to be created

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryPreloadAPI.InventoryPreloadPost(context.Background()).InventoryPreloadRecord(inventoryPreloadRecord).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryPreloadAPI.InventoryPreloadPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryPreloadPost`: InventoryPreloadRecord
	fmt.Fprintf(os.Stdout, "Response from `InventoryPreloadAPI.InventoryPreloadPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInventoryPreloadPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inventoryPreloadRecord** | [**InventoryPreloadRecord**](InventoryPreloadRecord.md) | Inventory Preload record or records to be created | 

### Return type

[**InventoryPreloadRecord**](InventoryPreloadRecord.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json, text/csv
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryPreloadValidateCsvPost

> InventoryPreloadCsvValidationSuccess InventoryPreloadValidateCsvPost(ctx).Body(body).Execute()

Validate a given CSV file 



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
	body := map[string]interface{}{ ... } // map[string]interface{} | Inventory Preload records to be validated. A CSV template can be downloaded from /api/inventory-preload/csv-template

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryPreloadAPI.InventoryPreloadValidateCsvPost(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryPreloadAPI.InventoryPreloadValidateCsvPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InventoryPreloadValidateCsvPost`: InventoryPreloadCsvValidationSuccess
	fmt.Fprintf(os.Stdout, "Response from `InventoryPreloadAPI.InventoryPreloadValidateCsvPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInventoryPreloadValidateCsvPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Inventory Preload records to be validated. A CSV template can be downloaded from /api/inventory-preload/csv-template | 

### Return type

[**InventoryPreloadCsvValidationSuccess**](InventoryPreloadCsvValidationSuccess.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: text/csv
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InventoryPreloadCsvTemplateGet

> map[string]interface{} V1InventoryPreloadCsvTemplateGet(ctx).Execute()

Retrieve the Inventory Preload CSV template 



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
	resp, r, err := apiClient.InventoryPreloadAPI.V1InventoryPreloadCsvTemplateGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryPreloadAPI.V1InventoryPreloadCsvTemplateGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InventoryPreloadCsvTemplateGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `InventoryPreloadAPI.V1InventoryPreloadCsvTemplateGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryPreloadCsvTemplateGetRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InventoryPreloadDelete

> V1InventoryPreloadDelete(ctx).Execute()

Delete all Inventory Preload records 



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
	r, err := apiClient.InventoryPreloadAPI.V1InventoryPreloadDelete(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryPreloadAPI.V1InventoryPreloadDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryPreloadDeleteRequest struct via the builder pattern


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


## V1InventoryPreloadGet

> InventoryPreloadRecordSearchResults V1InventoryPreloadGet(ctx).Page(page).Size(size).Pagesize(pagesize).PageSize(pageSize).Sort(sort).Execute()

Return all Inventory Preload records 



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
	sort := "sort_example" // string | Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc  (optional) (default to "id:asc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryPreloadAPI.V1InventoryPreloadGet(context.Background()).Page(page).Size(size).Pagesize(pagesize).PageSize(pageSize).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryPreloadAPI.V1InventoryPreloadGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InventoryPreloadGet`: InventoryPreloadRecordSearchResults
	fmt.Fprintf(os.Stdout, "Response from `InventoryPreloadAPI.V1InventoryPreloadGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryPreloadGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | [default to 0]
 **size** | **int64** |  | [default to 100]
 **pagesize** | **int64** |  | [default to 100]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **string** | Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;date:desc,name:asc  | [default to &quot;id:asc&quot;]

### Return type

[**InventoryPreloadRecordSearchResults**](InventoryPreloadRecordSearchResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InventoryPreloadHistoryGet

> HistorySearchResults V1InventoryPreloadHistoryGet(ctx).Page(page).Size(size).Pagesize(pagesize).PageSize(pageSize).Sort(sort).Execute()

Get Inventory Preload history entries 



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
	sort := "sort_example" // string | Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc  (optional) (default to "date:desc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryPreloadAPI.V1InventoryPreloadHistoryGet(context.Background()).Page(page).Size(size).Pagesize(pagesize).PageSize(pageSize).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryPreloadAPI.V1InventoryPreloadHistoryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InventoryPreloadHistoryGet`: HistorySearchResults
	fmt.Fprintf(os.Stdout, "Response from `InventoryPreloadAPI.V1InventoryPreloadHistoryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryPreloadHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | [default to 0]
 **size** | **int64** |  | [default to 100]
 **pagesize** | **int64** |  | [default to 100]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **string** | Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;date:desc,name:asc  | [default to &quot;date:desc&quot;]

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


## V1InventoryPreloadHistoryPost

> ObjectHistory V1InventoryPreloadHistoryPost(ctx).ObjectHistoryNote(objectHistoryNote).Execute()

Add Inventory Preload history object notes 



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
	objectHistoryNote := *openapiclient.NewObjectHistoryNote("A generic note can sometimes be useful, but generally not.") // ObjectHistoryNote | History notes to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryPreloadAPI.V1InventoryPreloadHistoryPost(context.Background()).ObjectHistoryNote(objectHistoryNote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryPreloadAPI.V1InventoryPreloadHistoryPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InventoryPreloadHistoryPost`: ObjectHistory
	fmt.Fprintf(os.Stdout, "Response from `InventoryPreloadAPI.V1InventoryPreloadHistoryPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryPreloadHistoryPostRequest struct via the builder pattern


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


## V1InventoryPreloadIdDelete

> V1InventoryPreloadIdDelete(ctx, id).Execute()

Delete an Inventory Preload record 



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
	id := int64(56) // int64 | Inventory Preload identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InventoryPreloadAPI.V1InventoryPreloadIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryPreloadAPI.V1InventoryPreloadIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Inventory Preload identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryPreloadIdDeleteRequest struct via the builder pattern


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


## V1InventoryPreloadIdGet

> InventoryPreloadRecord V1InventoryPreloadIdGet(ctx, id).Execute()

Get an Inventory Preload record 



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
	id := int64(56) // int64 | Inventory Preload identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryPreloadAPI.V1InventoryPreloadIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryPreloadAPI.V1InventoryPreloadIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InventoryPreloadIdGet`: InventoryPreloadRecord
	fmt.Fprintf(os.Stdout, "Response from `InventoryPreloadAPI.V1InventoryPreloadIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Inventory Preload identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryPreloadIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InventoryPreloadRecord**](InventoryPreloadRecord.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InventoryPreloadIdPut

> InventoryPreloadRecord V1InventoryPreloadIdPut(ctx, id).InventoryPreloadRecord(inventoryPreloadRecord).Execute()

Update an Inventory Preload record 



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
	id := int64(56) // int64 | Inventory Preload identifier
	inventoryPreloadRecord := *openapiclient.NewInventoryPreloadRecord("C02L29ECF8J1", "Computer") // InventoryPreloadRecord | Inventory Preload record to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryPreloadAPI.V1InventoryPreloadIdPut(context.Background(), id).InventoryPreloadRecord(inventoryPreloadRecord).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryPreloadAPI.V1InventoryPreloadIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InventoryPreloadIdPut`: InventoryPreloadRecord
	fmt.Fprintf(os.Stdout, "Response from `InventoryPreloadAPI.V1InventoryPreloadIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Inventory Preload identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryPreloadIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inventoryPreloadRecord** | [**InventoryPreloadRecord**](InventoryPreloadRecord.md) | Inventory Preload record to update | 

### Return type

[**InventoryPreloadRecord**](InventoryPreloadRecord.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InventoryPreloadPost

> InventoryPreloadRecord V1InventoryPreloadPost(ctx).InventoryPreloadRecord(inventoryPreloadRecord).Execute()

Create a new Inventory Preload record using JSON or CSV 



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
	inventoryPreloadRecord := *openapiclient.NewInventoryPreloadRecord("C02L29ECF8J1", "Computer") // InventoryPreloadRecord | Inventory Preload record or records to be created

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryPreloadAPI.V1InventoryPreloadPost(context.Background()).InventoryPreloadRecord(inventoryPreloadRecord).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryPreloadAPI.V1InventoryPreloadPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InventoryPreloadPost`: InventoryPreloadRecord
	fmt.Fprintf(os.Stdout, "Response from `InventoryPreloadAPI.V1InventoryPreloadPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryPreloadPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inventoryPreloadRecord** | [**InventoryPreloadRecord**](InventoryPreloadRecord.md) | Inventory Preload record or records to be created | 

### Return type

[**InventoryPreloadRecord**](InventoryPreloadRecord.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json, text/csv
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InventoryPreloadValidateCsvPost

> InventoryPreloadCsvValidationSuccess V1InventoryPreloadValidateCsvPost(ctx).Body(body).Execute()

Validate a given CSV file 



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
	body := map[string]interface{}{ ... } // map[string]interface{} | Inventory Preload records to be validated. A CSV template can be downloaded from /api/inventory-preload/csv-template

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryPreloadAPI.V1InventoryPreloadValidateCsvPost(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryPreloadAPI.V1InventoryPreloadValidateCsvPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InventoryPreloadValidateCsvPost`: InventoryPreloadCsvValidationSuccess
	fmt.Fprintf(os.Stdout, "Response from `InventoryPreloadAPI.V1InventoryPreloadValidateCsvPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryPreloadValidateCsvPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Inventory Preload records to be validated. A CSV template can be downloaded from /api/inventory-preload/csv-template | 

### Return type

[**InventoryPreloadCsvValidationSuccess**](InventoryPreloadCsvValidationSuccess.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: text/csv
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2InventoryPreloadCsvGet

> string V2InventoryPreloadCsvGet(ctx).Execute()

Download all Inventory Preload records



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
	resp, r, err := apiClient.InventoryPreloadAPI.V2InventoryPreloadCsvGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryPreloadAPI.V2InventoryPreloadCsvGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2InventoryPreloadCsvGet`: string
	fmt.Fprintf(os.Stdout, "Response from `InventoryPreloadAPI.V2InventoryPreloadCsvGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV2InventoryPreloadCsvGetRequest struct via the builder pattern


### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2InventoryPreloadCsvPost

> []HrefResponse V2InventoryPreloadCsvPost(ctx).File(file).Execute()

Create one or more new Inventory Preload records using CSV 



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
	file := "file_example" // string | The CSV file to upload

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryPreloadAPI.V2InventoryPreloadCsvPost(context.Background()).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryPreloadAPI.V2InventoryPreloadCsvPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2InventoryPreloadCsvPost`: []HrefResponse
	fmt.Fprintf(os.Stdout, "Response from `InventoryPreloadAPI.V2InventoryPreloadCsvPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2InventoryPreloadCsvPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | **string** | The CSV file to upload | 

### Return type

[**[]HrefResponse**](HrefResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2InventoryPreloadCsvTemplateGet

> string V2InventoryPreloadCsvTemplateGet(ctx).Execute()

Download the Inventory Preload CSV template



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
	resp, r, err := apiClient.InventoryPreloadAPI.V2InventoryPreloadCsvTemplateGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryPreloadAPI.V2InventoryPreloadCsvTemplateGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2InventoryPreloadCsvTemplateGet`: string
	fmt.Fprintf(os.Stdout, "Response from `InventoryPreloadAPI.V2InventoryPreloadCsvTemplateGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV2InventoryPreloadCsvTemplateGetRequest struct via the builder pattern


### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2InventoryPreloadCsvValidatePost

> InventoryPreloadCsvValidationSuccess V2InventoryPreloadCsvValidatePost(ctx).File(file).Execute()

Validate a given CSV file 



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
	file := "file_example" // string | The CSV file to upload

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryPreloadAPI.V2InventoryPreloadCsvValidatePost(context.Background()).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryPreloadAPI.V2InventoryPreloadCsvValidatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2InventoryPreloadCsvValidatePost`: InventoryPreloadCsvValidationSuccess
	fmt.Fprintf(os.Stdout, "Response from `InventoryPreloadAPI.V2InventoryPreloadCsvValidatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2InventoryPreloadCsvValidatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | **string** | The CSV file to upload | 

### Return type

[**InventoryPreloadCsvValidationSuccess**](InventoryPreloadCsvValidationSuccess.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2InventoryPreloadEaColumnsGet

> InventoryPreloadExtensionAttributeColumnResult V2InventoryPreloadEaColumnsGet(ctx).Execute()

Retrieve a list of extension attribute columns 



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
	resp, r, err := apiClient.InventoryPreloadAPI.V2InventoryPreloadEaColumnsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryPreloadAPI.V2InventoryPreloadEaColumnsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2InventoryPreloadEaColumnsGet`: InventoryPreloadExtensionAttributeColumnResult
	fmt.Fprintf(os.Stdout, "Response from `InventoryPreloadAPI.V2InventoryPreloadEaColumnsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV2InventoryPreloadEaColumnsGetRequest struct via the builder pattern


### Return type

[**InventoryPreloadExtensionAttributeColumnResult**](InventoryPreloadExtensionAttributeColumnResult.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2InventoryPreloadExportPost

> interface{} V2InventoryPreloadExportPost(ctx).ExportFields(exportFields).ExportLabels(exportLabels).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).ExportParameters(exportParameters).Execute()

Export a collection of inventory preload records 



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
	exportFields := []string{"Inner_example"} // []string | Export fields parameter, used to change default order or ignore some of the response properties. Default is empty array, which means that all fields of the response entity will be serialized. Example: export-fields=id,username (optional) (default to [])
	exportLabels := []string{"Inner_example"} // []string | Export labels parameter, used to customize fieldnames/columns in the exported file. Default is empty array, which means that response properties names will be used. Number of the provided labels must match the number of export-fields Example: export-labels=identifier,name with matching: export-fields=id,username (optional) (default to [])
	page := int64(56) // int64 |  (optional) (default to 0)
	pageSize := int64(56) // int64 |  (optional) (default to 100)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: `property:asc/desc`. Default sort is `id:asc`. Multiple sort criteria are supported and must be separated with a comma. All inventory preload fields are supported, however fields added by extension attributes are not supported. If sorting by deviceType, use `0` for Computer and `1` for Mobile Device.  Example: `sort=date:desc,name:asc`.  (optional) (default to ["id:asc"])
	filter := "filter_example" // string | Allowing to filter inventory preload records. Default search is empty query - returning all results for the requested page. All inventory preload fields are supported, however fields added by extension attributes are not supported. If filtering by deviceType, use `0` for Computer and `1` for Mobile Device.  Query in the RSQL format, allowing `==`, `!=`, `>`, `<`, and `=in=`.  Example: `filter=categoryName==\"Category\"`  (optional) (default to "")
	exportParameters := *openapiclient.NewExportParameters() // ExportParameters | Optional. Override query parameters since they can make URI exceed 2,000 character limit. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryPreloadAPI.V2InventoryPreloadExportPost(context.Background()).ExportFields(exportFields).ExportLabels(exportLabels).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).ExportParameters(exportParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryPreloadAPI.V2InventoryPreloadExportPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2InventoryPreloadExportPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `InventoryPreloadAPI.V2InventoryPreloadExportPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2InventoryPreloadExportPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **exportFields** | **[]string** | Export fields parameter, used to change default order or ignore some of the response properties. Default is empty array, which means that all fields of the response entity will be serialized. Example: export-fields&#x3D;id,username | [default to []]
 **exportLabels** | **[]string** | Export labels parameter, used to customize fieldnames/columns in the exported file. Default is empty array, which means that response properties names will be used. Number of the provided labels must match the number of export-fields Example: export-labels&#x3D;identifier,name with matching: export-fields&#x3D;id,username | [default to []]
 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: &#x60;property:asc/desc&#x60;. Default sort is &#x60;id:asc&#x60;. Multiple sort criteria are supported and must be separated with a comma. All inventory preload fields are supported, however fields added by extension attributes are not supported. If sorting by deviceType, use &#x60;0&#x60; for Computer and &#x60;1&#x60; for Mobile Device.  Example: &#x60;sort&#x3D;date:desc,name:asc&#x60;.  | [default to [&quot;id:asc&quot;]]
 **filter** | **string** | Allowing to filter inventory preload records. Default search is empty query - returning all results for the requested page. All inventory preload fields are supported, however fields added by extension attributes are not supported. If filtering by deviceType, use &#x60;0&#x60; for Computer and &#x60;1&#x60; for Mobile Device.  Query in the RSQL format, allowing &#x60;&#x3D;&#x3D;&#x60;, &#x60;!&#x3D;&#x60;, &#x60;&gt;&#x60;, &#x60;&lt;&#x60;, and &#x60;&#x3D;in&#x3D;&#x60;.  Example: &#x60;filter&#x3D;categoryName&#x3D;&#x3D;\&quot;Category\&quot;&#x60;  | [default to &quot;&quot;]
 **exportParameters** | [**ExportParameters**](ExportParameters.md) | Optional. Override query parameters since they can make URI exceed 2,000 character limit. | 

### Return type

**interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/csv, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2InventoryPreloadHistoryGet

> HistorySearchResults V2InventoryPreloadHistoryGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Get Inventory Preload history entries 



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
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: `property:asc/desc`. Default sort is `date:desc`. Multiple sort criteria are supported and must be separated with a comma.  Example: `sort=date:desc,name:asc`.  (optional) (default to ["date:desc"])
	filter := "filter_example" // string | Allows filtering inventory preload history records. Default search is empty query - returning all results for the requested page. All inventory preload history fields are supported.  Query in the RSQL format, allowing `==`, `!=`, `>`, `<`, and `=in=`.  Example: `filter=username==\"admin\"`  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryPreloadAPI.V2InventoryPreloadHistoryGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryPreloadAPI.V2InventoryPreloadHistoryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2InventoryPreloadHistoryGet`: HistorySearchResults
	fmt.Fprintf(os.Stdout, "Response from `InventoryPreloadAPI.V2InventoryPreloadHistoryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2InventoryPreloadHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: &#x60;property:asc/desc&#x60;. Default sort is &#x60;date:desc&#x60;. Multiple sort criteria are supported and must be separated with a comma.  Example: &#x60;sort&#x3D;date:desc,name:asc&#x60;.  | [default to [&quot;date:desc&quot;]]
 **filter** | **string** | Allows filtering inventory preload history records. Default search is empty query - returning all results for the requested page. All inventory preload history fields are supported.  Query in the RSQL format, allowing &#x60;&#x3D;&#x3D;&#x60;, &#x60;!&#x3D;&#x60;, &#x60;&gt;&#x60;, &#x60;&lt;&#x60;, and &#x60;&#x3D;in&#x3D;&#x60;.  Example: &#x60;filter&#x3D;username&#x3D;&#x3D;\&quot;admin\&quot;&#x60;  | [default to &quot;&quot;]

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


## V2InventoryPreloadHistoryPost

> HrefResponse V2InventoryPreloadHistoryPost(ctx).ObjectHistoryNote(objectHistoryNote).Execute()

Add Inventory Preload history object notes



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
	objectHistoryNote := *openapiclient.NewObjectHistoryNote("A generic note can sometimes be useful, but generally not.") // ObjectHistoryNote | History notes to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryPreloadAPI.V2InventoryPreloadHistoryPost(context.Background()).ObjectHistoryNote(objectHistoryNote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryPreloadAPI.V2InventoryPreloadHistoryPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2InventoryPreloadHistoryPost`: HrefResponse
	fmt.Fprintf(os.Stdout, "Response from `InventoryPreloadAPI.V2InventoryPreloadHistoryPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2InventoryPreloadHistoryPostRequest struct via the builder pattern


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


## V2InventoryPreloadRecordsDeleteAllPost

> V2InventoryPreloadRecordsDeleteAllPost(ctx).Execute()

Delete all Inventory Preload records 



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
	r, err := apiClient.InventoryPreloadAPI.V2InventoryPreloadRecordsDeleteAllPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryPreloadAPI.V2InventoryPreloadRecordsDeleteAllPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV2InventoryPreloadRecordsDeleteAllPostRequest struct via the builder pattern


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


## V2InventoryPreloadRecordsGet

> InventoryPreloadRecordSearchResultsV2 V2InventoryPreloadRecordsGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Return all Inventory Preload records



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
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: `property:asc/desc`. Default sort is `id:asc`. Multiple sort criteria are supported and must be separated with a comma. All inventory preload fields are supported, however fields added by extension attributes are not supported. If sorting by deviceType, use `0` for Computer and `1` for Mobile Device.  Example: `sort=date:desc,name:asc`.  (optional) (default to ["id:asc"])
	filter := "filter_example" // string | Allowing to filter inventory preload records. Default search is empty query - returning all results for the requested page. All inventory preload fields are supported, however fields added by extension attributes are not supported. If filtering by deviceType, use `0` for Computer and `1` for Mobile Device.  Query in the RSQL format, allowing `==`, `!=`, `>`, `<`, and `=in=`.  Example: `filter=categoryName==\"Category\"`  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryPreloadAPI.V2InventoryPreloadRecordsGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryPreloadAPI.V2InventoryPreloadRecordsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2InventoryPreloadRecordsGet`: InventoryPreloadRecordSearchResultsV2
	fmt.Fprintf(os.Stdout, "Response from `InventoryPreloadAPI.V2InventoryPreloadRecordsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2InventoryPreloadRecordsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: &#x60;property:asc/desc&#x60;. Default sort is &#x60;id:asc&#x60;. Multiple sort criteria are supported and must be separated with a comma. All inventory preload fields are supported, however fields added by extension attributes are not supported. If sorting by deviceType, use &#x60;0&#x60; for Computer and &#x60;1&#x60; for Mobile Device.  Example: &#x60;sort&#x3D;date:desc,name:asc&#x60;.  | [default to [&quot;id:asc&quot;]]
 **filter** | **string** | Allowing to filter inventory preload records. Default search is empty query - returning all results for the requested page. All inventory preload fields are supported, however fields added by extension attributes are not supported. If filtering by deviceType, use &#x60;0&#x60; for Computer and &#x60;1&#x60; for Mobile Device.  Query in the RSQL format, allowing &#x60;&#x3D;&#x3D;&#x60;, &#x60;!&#x3D;&#x60;, &#x60;&gt;&#x60;, &#x60;&lt;&#x60;, and &#x60;&#x3D;in&#x3D;&#x60;.  Example: &#x60;filter&#x3D;categoryName&#x3D;&#x3D;\&quot;Category\&quot;&#x60;  | [default to &quot;&quot;]

### Return type

[**InventoryPreloadRecordSearchResultsV2**](InventoryPreloadRecordSearchResultsV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2InventoryPreloadRecordsIdDelete

> V2InventoryPreloadRecordsIdDelete(ctx, id).Execute()

Delete an Inventory Preload record 



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
	id := "id_example" // string | Inventory Preload identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InventoryPreloadAPI.V2InventoryPreloadRecordsIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryPreloadAPI.V2InventoryPreloadRecordsIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Inventory Preload identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2InventoryPreloadRecordsIdDeleteRequest struct via the builder pattern


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


## V2InventoryPreloadRecordsIdGet

> InventoryPreloadRecordV2 V2InventoryPreloadRecordsIdGet(ctx, id).Execute()

Get an Inventory Preload record



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
	id := "id_example" // string | Inventory Preload identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryPreloadAPI.V2InventoryPreloadRecordsIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryPreloadAPI.V2InventoryPreloadRecordsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2InventoryPreloadRecordsIdGet`: InventoryPreloadRecordV2
	fmt.Fprintf(os.Stdout, "Response from `InventoryPreloadAPI.V2InventoryPreloadRecordsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Inventory Preload identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2InventoryPreloadRecordsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InventoryPreloadRecordV2**](InventoryPreloadRecordV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2InventoryPreloadRecordsIdPut

> InventoryPreloadRecordV2 V2InventoryPreloadRecordsIdPut(ctx, id).InventoryPreloadRecordV2(inventoryPreloadRecordV2).Execute()

Update an Inventory Preload record



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
	id := "id_example" // string | Inventory Preload identifier
	inventoryPreloadRecordV2 := *openapiclient.NewInventoryPreloadRecordV2("C02L29ECF8J1", "Computer") // InventoryPreloadRecordV2 | Inventory Preload record to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryPreloadAPI.V2InventoryPreloadRecordsIdPut(context.Background(), id).InventoryPreloadRecordV2(inventoryPreloadRecordV2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryPreloadAPI.V2InventoryPreloadRecordsIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2InventoryPreloadRecordsIdPut`: InventoryPreloadRecordV2
	fmt.Fprintf(os.Stdout, "Response from `InventoryPreloadAPI.V2InventoryPreloadRecordsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Inventory Preload identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2InventoryPreloadRecordsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inventoryPreloadRecordV2** | [**InventoryPreloadRecordV2**](InventoryPreloadRecordV2.md) | Inventory Preload record to update | 

### Return type

[**InventoryPreloadRecordV2**](InventoryPreloadRecordV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2InventoryPreloadRecordsPost

> HrefResponse V2InventoryPreloadRecordsPost(ctx).InventoryPreloadRecordV2(inventoryPreloadRecordV2).Execute()

Create a new Inventory Preload record using JSON



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
	inventoryPreloadRecordV2 := *openapiclient.NewInventoryPreloadRecordV2("C02L29ECF8J1", "Computer") // InventoryPreloadRecordV2 | Inventory Preload record to be created.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryPreloadAPI.V2InventoryPreloadRecordsPost(context.Background()).InventoryPreloadRecordV2(inventoryPreloadRecordV2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryPreloadAPI.V2InventoryPreloadRecordsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2InventoryPreloadRecordsPost`: HrefResponse
	fmt.Fprintf(os.Stdout, "Response from `InventoryPreloadAPI.V2InventoryPreloadRecordsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2InventoryPreloadRecordsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inventoryPreloadRecordV2** | [**InventoryPreloadRecordV2**](InventoryPreloadRecordV2.md) | Inventory Preload record to be created. | 

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

