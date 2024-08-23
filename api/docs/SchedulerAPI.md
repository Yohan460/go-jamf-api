# \SchedulerAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1SchedulerJobsGet**](SchedulerAPI.md#V1SchedulerJobsGet) | **Get** /v1/scheduler/jobs | Retrieve all Jamf Pro Scheduler jobs
[**V1SchedulerJobsJobKeyTriggersGet**](SchedulerAPI.md#V1SchedulerJobsJobKeyTriggersGet) | **Get** /v1/scheduler/jobs/{jobKey}/triggers | Retrieve all triggers for a Jamf Pro Scheduler job
[**V1SchedulerSummaryGet**](SchedulerAPI.md#V1SchedulerSummaryGet) | **Get** /v1/scheduler/summary | Retrieve a summary of the Jamf Pro Scheduler



## V1SchedulerJobsGet

> SchedulerJobs V1SchedulerJobsGet(ctx).Execute()

Retrieve all Jamf Pro Scheduler jobs



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
	resp, r, err := apiClient.SchedulerAPI.V1SchedulerJobsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SchedulerAPI.V1SchedulerJobsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SchedulerJobsGet`: SchedulerJobs
	fmt.Fprintf(os.Stdout, "Response from `SchedulerAPI.V1SchedulerJobsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1SchedulerJobsGetRequest struct via the builder pattern


### Return type

[**SchedulerJobs**](SchedulerJobs.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SchedulerJobsJobKeyTriggersGet

> SchedulerJob V1SchedulerJobsJobKeyTriggersGet(ctx, jobKey).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Retrieve all triggers for a Jamf Pro Scheduler job



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
	jobKey := "jobKey_example" // string | Jamf Pro Scheduler Job Key
	page := int64(56) // int64 |  (optional) (default to 0)
	pageSize := int64(56) // int64 |  (optional) (default to 100)
	sort := []string{"Inner_example"} // []string | Sorts results by one or more criteria, following the format property:asc/desc. Default sort is nextFireTime:asc. If using multiple criteria, separate with commas. (optional) (default to ["nextFireTime:asc"])
	filter := "filter_example" // string | Query in the RSQL format, allowing to filter the Jamf Pro Scheduler triggers collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: triggerKey, previousFireTime, nextFireTime. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SchedulerAPI.V1SchedulerJobsJobKeyTriggersGet(context.Background(), jobKey).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SchedulerAPI.V1SchedulerJobsJobKeyTriggersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SchedulerJobsJobKeyTriggersGet`: SchedulerJob
	fmt.Fprintf(os.Stdout, "Response from `SchedulerAPI.V1SchedulerJobsJobKeyTriggersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobKey** | **string** | Jamf Pro Scheduler Job Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1SchedulerJobsJobKeyTriggersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorts results by one or more criteria, following the format property:asc/desc. Default sort is nextFireTime:asc. If using multiple criteria, separate with commas. | [default to [&quot;nextFireTime:asc&quot;]]
 **filter** | **string** | Query in the RSQL format, allowing to filter the Jamf Pro Scheduler triggers collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: triggerKey, previousFireTime, nextFireTime. | [default to &quot;&quot;]

### Return type

[**SchedulerJob**](SchedulerJob.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SchedulerSummaryGet

> SchedulerSummary V1SchedulerSummaryGet(ctx).Execute()

Retrieve a summary of the Jamf Pro Scheduler



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
	resp, r, err := apiClient.SchedulerAPI.V1SchedulerSummaryGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SchedulerAPI.V1SchedulerSummaryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SchedulerSummaryGet`: SchedulerSummary
	fmt.Fprintf(os.Stdout, "Response from `SchedulerAPI.V1SchedulerSummaryGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1SchedulerSummaryGetRequest struct via the builder pattern


### Return type

[**SchedulerSummary**](SchedulerSummary.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

