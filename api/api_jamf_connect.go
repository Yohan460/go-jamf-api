/*
Jamf Pro API

## Overview The Jamf Pro API is a RESTful API for Jamf Pro built to enable consistent and efficient programmatic access to Jamf Pro.<br/><br/> The swagger schema can be found [here](/api/schema/). 

API version: production
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"reflect"
)


type JamfConnectAPI interface {

	/*
	V1JamfConnectConfigProfilesGet Search for config profiles linked to Jamf Connect 

	Search for config profiles linked to Jamf Connect

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return JamfConnectAPIV1JamfConnectConfigProfilesGetRequest
	*/
	V1JamfConnectConfigProfilesGet(ctx context.Context) JamfConnectAPIV1JamfConnectConfigProfilesGetRequest

	// V1JamfConnectConfigProfilesGetExecute executes the request
	//  @return LinkedConnectProfileSearchResults
	V1JamfConnectConfigProfilesGetExecute(r JamfConnectAPIV1JamfConnectConfigProfilesGetRequest) (*LinkedConnectProfileSearchResults, *http.Response, error)

	/*
	V1JamfConnectConfigProfilesIdPut Update the way the Jamf Connect app gets updated on computers within scope of the associated configuration profile. 

	Update the way the Jamf Connect app gets updated on computers within scope of the associated configuration profile.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id the UUID of the profile to update
	@return JamfConnectAPIV1JamfConnectConfigProfilesIdPutRequest
	*/
	V1JamfConnectConfigProfilesIdPut(ctx context.Context, id string) JamfConnectAPIV1JamfConnectConfigProfilesIdPutRequest

	// V1JamfConnectConfigProfilesIdPutExecute executes the request
	//  @return LinkedConnectProfile
	V1JamfConnectConfigProfilesIdPutExecute(r JamfConnectAPIV1JamfConnectConfigProfilesIdPutRequest) (*LinkedConnectProfile, *http.Response, error)

	/*
	V1JamfConnectDeploymentsIdTasksGet Search for deployment tasks for a config profile linked to Jamf Connect 

	Search for config profiles linked to Jamf Connect

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id the UUID of the Jamf Connect deployment
	@return JamfConnectAPIV1JamfConnectDeploymentsIdTasksGetRequest
	*/
	V1JamfConnectDeploymentsIdTasksGet(ctx context.Context, id string) JamfConnectAPIV1JamfConnectDeploymentsIdTasksGetRequest

	// V1JamfConnectDeploymentsIdTasksGetExecute executes the request
	//  @return DeploymentTaskSearchResults
	V1JamfConnectDeploymentsIdTasksGetExecute(r JamfConnectAPIV1JamfConnectDeploymentsIdTasksGetRequest) (*DeploymentTaskSearchResults, *http.Response, error)

	/*
	V1JamfConnectDeploymentsIdTasksRetryPost Request a retry of Connect install tasks 

	Request a retry of Connect install tasks


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id the UUID of the deployment associated with the retry
	@return JamfConnectAPIV1JamfConnectDeploymentsIdTasksRetryPostRequest
	*/
	V1JamfConnectDeploymentsIdTasksRetryPost(ctx context.Context, id string) JamfConnectAPIV1JamfConnectDeploymentsIdTasksRetryPostRequest

	// V1JamfConnectDeploymentsIdTasksRetryPostExecute executes the request
	V1JamfConnectDeploymentsIdTasksRetryPostExecute(r JamfConnectAPIV1JamfConnectDeploymentsIdTasksRetryPostRequest) (*http.Response, error)

	/*
	V1JamfConnectGet Get the Jamf Connect settings that you have access to see 

	Get the Jamf Connect settings that you have access to see.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return JamfConnectAPIV1JamfConnectGetRequest
	*/
	V1JamfConnectGet(ctx context.Context) JamfConnectAPIV1JamfConnectGetRequest

	// V1JamfConnectGetExecute executes the request
	V1JamfConnectGetExecute(r JamfConnectAPIV1JamfConnectGetRequest) (*http.Response, error)

	/*
	V1JamfConnectHistoryGet Get Jamf Connect history 

	Get Jamf Connect history


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return JamfConnectAPIV1JamfConnectHistoryGetRequest
	*/
	V1JamfConnectHistoryGet(ctx context.Context) JamfConnectAPIV1JamfConnectHistoryGetRequest

	// V1JamfConnectHistoryGetExecute executes the request
	//  @return HistorySearchResults
	V1JamfConnectHistoryGetExecute(r JamfConnectAPIV1JamfConnectHistoryGetRequest) (*HistorySearchResults, *http.Response, error)

	/*
	V1JamfConnectHistoryPost Add Jamf Connect history notes 

	Add Jamf Connect history notes


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return JamfConnectAPIV1JamfConnectHistoryPostRequest
	*/
	V1JamfConnectHistoryPost(ctx context.Context) JamfConnectAPIV1JamfConnectHistoryPostRequest

	// V1JamfConnectHistoryPostExecute executes the request
	//  @return HrefResponse
	V1JamfConnectHistoryPostExecute(r JamfConnectAPIV1JamfConnectHistoryPostRequest) (*HrefResponse, *http.Response, error)
}

// JamfConnectAPIService JamfConnectAPI service
type JamfConnectAPIService service

type JamfConnectAPIV1JamfConnectConfigProfilesGetRequest struct {
	ctx context.Context
	ApiService JamfConnectAPI
	page *int32
	pageSize *int32
	sort *[]string
	filter *string
}

func (r JamfConnectAPIV1JamfConnectConfigProfilesGetRequest) Page(page int32) JamfConnectAPIV1JamfConnectConfigProfilesGetRequest {
	r.page = &page
	return r
}

func (r JamfConnectAPIV1JamfConnectConfigProfilesGetRequest) PageSize(pageSize int32) JamfConnectAPIV1JamfConnectConfigProfilesGetRequest {
	r.pageSize = &pageSize
	return r
}

// Sorting criteria in the format: property:asc/desc. Default sort order is descending. Multiple sort criteria are supported and must be entered on separate lines in Swagger UI. In the URI the &#39;sort&#39; query param is not duplicated for each sort criterion, e.g., ...&amp;sort&#x3D;name:asc,date:desc. Fields that can be sorted: status, updated
func (r JamfConnectAPIV1JamfConnectConfigProfilesGetRequest) Sort(sort []string) JamfConnectAPIV1JamfConnectConfigProfilesGetRequest {
	r.sort = &sort
	return r
}

// Query in the RSQL format, allowing to filter results. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: status, updated, version This param can be combined with paging and sorting. Example: filter&#x3D;username!&#x3D;admin and details&#x3D;&#x3D;*disabled* and date&lt;2019-12-15
func (r JamfConnectAPIV1JamfConnectConfigProfilesGetRequest) Filter(filter string) JamfConnectAPIV1JamfConnectConfigProfilesGetRequest {
	r.filter = &filter
	return r
}

func (r JamfConnectAPIV1JamfConnectConfigProfilesGetRequest) Execute() (*LinkedConnectProfileSearchResults, *http.Response, error) {
	return r.ApiService.V1JamfConnectConfigProfilesGetExecute(r)
}

/*
V1JamfConnectConfigProfilesGet Search for config profiles linked to Jamf Connect 

Search for config profiles linked to Jamf Connect

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return JamfConnectAPIV1JamfConnectConfigProfilesGetRequest
*/
func (a *JamfConnectAPIService) V1JamfConnectConfigProfilesGet(ctx context.Context) JamfConnectAPIV1JamfConnectConfigProfilesGetRequest {
	return JamfConnectAPIV1JamfConnectConfigProfilesGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return LinkedConnectProfileSearchResults
func (a *JamfConnectAPIService) V1JamfConnectConfigProfilesGetExecute(r JamfConnectAPIV1JamfConnectConfigProfilesGetRequest) (*LinkedConnectProfileSearchResults, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *LinkedConnectProfileSearchResults
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JamfConnectAPIService.V1JamfConnectConfigProfilesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/jamf-connect/config-profiles"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	} else {
		var defaultValue int32 = 0
		r.page = &defaultValue
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page-size", r.pageSize, "")
	} else {
		var defaultValue int32 = 100
		r.pageSize = &defaultValue
	}
	if r.sort != nil {
		t := *r.sort
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "sort", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "sort", t, "multi")
		}
	} else {
		var defaultValue []string = []
		r.sort = &defaultValue
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "")
	} else {
		var defaultValue string = ""
		r.filter = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type JamfConnectAPIV1JamfConnectConfigProfilesIdPutRequest struct {
	ctx context.Context
	ApiService JamfConnectAPI
	id string
	linkedConnectProfile *LinkedConnectProfile
}

// Updatable Jamf Connect Settings
func (r JamfConnectAPIV1JamfConnectConfigProfilesIdPutRequest) LinkedConnectProfile(linkedConnectProfile LinkedConnectProfile) JamfConnectAPIV1JamfConnectConfigProfilesIdPutRequest {
	r.linkedConnectProfile = &linkedConnectProfile
	return r
}

func (r JamfConnectAPIV1JamfConnectConfigProfilesIdPutRequest) Execute() (*LinkedConnectProfile, *http.Response, error) {
	return r.ApiService.V1JamfConnectConfigProfilesIdPutExecute(r)
}

/*
V1JamfConnectConfigProfilesIdPut Update the way the Jamf Connect app gets updated on computers within scope of the associated configuration profile. 

Update the way the Jamf Connect app gets updated on computers within scope of the associated configuration profile.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the UUID of the profile to update
 @return JamfConnectAPIV1JamfConnectConfigProfilesIdPutRequest
*/
func (a *JamfConnectAPIService) V1JamfConnectConfigProfilesIdPut(ctx context.Context, id string) JamfConnectAPIV1JamfConnectConfigProfilesIdPutRequest {
	return JamfConnectAPIV1JamfConnectConfigProfilesIdPutRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return LinkedConnectProfile
func (a *JamfConnectAPIService) V1JamfConnectConfigProfilesIdPutExecute(r JamfConnectAPIV1JamfConnectConfigProfilesIdPutRequest) (*LinkedConnectProfile, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *LinkedConnectProfile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JamfConnectAPIService.V1JamfConnectConfigProfilesIdPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/jamf-connect/config-profiles/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.linkedConnectProfile
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type JamfConnectAPIV1JamfConnectDeploymentsIdTasksGetRequest struct {
	ctx context.Context
	ApiService JamfConnectAPI
	id string
	page *int32
	pageSize *int32
	sort *[]string
	filter *string
}

func (r JamfConnectAPIV1JamfConnectDeploymentsIdTasksGetRequest) Page(page int32) JamfConnectAPIV1JamfConnectDeploymentsIdTasksGetRequest {
	r.page = &page
	return r
}

func (r JamfConnectAPIV1JamfConnectDeploymentsIdTasksGetRequest) PageSize(pageSize int32) JamfConnectAPIV1JamfConnectDeploymentsIdTasksGetRequest {
	r.pageSize = &pageSize
	return r
}

// Sorting criteria in the format: property:asc/desc. Default sort order is descending. Multiple sort criteria are supported and must be entered on separate lines in Swagger UI. In the URI the &#39;sort&#39; query param is not duplicated for each sort criterion, e.g., ...&amp;sort&#x3D;name:asc,date:desc. Fields that can be sorted: status, updated
func (r JamfConnectAPIV1JamfConnectDeploymentsIdTasksGetRequest) Sort(sort []string) JamfConnectAPIV1JamfConnectDeploymentsIdTasksGetRequest {
	r.sort = &sort
	return r
}

// Query in the RSQL format, allowing to filter results. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: status, updated, version This param can be combined with paging and sorting. Example: filter&#x3D;username!&#x3D;admin and details&#x3D;&#x3D;*disabled* and date&lt;2019-12-15
func (r JamfConnectAPIV1JamfConnectDeploymentsIdTasksGetRequest) Filter(filter string) JamfConnectAPIV1JamfConnectDeploymentsIdTasksGetRequest {
	r.filter = &filter
	return r
}

func (r JamfConnectAPIV1JamfConnectDeploymentsIdTasksGetRequest) Execute() (*DeploymentTaskSearchResults, *http.Response, error) {
	return r.ApiService.V1JamfConnectDeploymentsIdTasksGetExecute(r)
}

/*
V1JamfConnectDeploymentsIdTasksGet Search for deployment tasks for a config profile linked to Jamf Connect 

Search for config profiles linked to Jamf Connect

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the UUID of the Jamf Connect deployment
 @return JamfConnectAPIV1JamfConnectDeploymentsIdTasksGetRequest
*/
func (a *JamfConnectAPIService) V1JamfConnectDeploymentsIdTasksGet(ctx context.Context, id string) JamfConnectAPIV1JamfConnectDeploymentsIdTasksGetRequest {
	return JamfConnectAPIV1JamfConnectDeploymentsIdTasksGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return DeploymentTaskSearchResults
func (a *JamfConnectAPIService) V1JamfConnectDeploymentsIdTasksGetExecute(r JamfConnectAPIV1JamfConnectDeploymentsIdTasksGetRequest) (*DeploymentTaskSearchResults, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DeploymentTaskSearchResults
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JamfConnectAPIService.V1JamfConnectDeploymentsIdTasksGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/jamf-connect/deployments/{id}/tasks"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	} else {
		var defaultValue int32 = 0
		r.page = &defaultValue
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page-size", r.pageSize, "")
	} else {
		var defaultValue int32 = 100
		r.pageSize = &defaultValue
	}
	if r.sort != nil {
		t := *r.sort
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "sort", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "sort", t, "multi")
		}
	} else {
		var defaultValue []string = []
		r.sort = &defaultValue
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "")
	} else {
		var defaultValue string = ""
		r.filter = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type JamfConnectAPIV1JamfConnectDeploymentsIdTasksRetryPostRequest struct {
	ctx context.Context
	ApiService JamfConnectAPI
	id string
	ids *Ids
}

// task IDs to retry
func (r JamfConnectAPIV1JamfConnectDeploymentsIdTasksRetryPostRequest) Ids(ids Ids) JamfConnectAPIV1JamfConnectDeploymentsIdTasksRetryPostRequest {
	r.ids = &ids
	return r
}

func (r JamfConnectAPIV1JamfConnectDeploymentsIdTasksRetryPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1JamfConnectDeploymentsIdTasksRetryPostExecute(r)
}

/*
V1JamfConnectDeploymentsIdTasksRetryPost Request a retry of Connect install tasks 

Request a retry of Connect install tasks


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the UUID of the deployment associated with the retry
 @return JamfConnectAPIV1JamfConnectDeploymentsIdTasksRetryPostRequest
*/
func (a *JamfConnectAPIService) V1JamfConnectDeploymentsIdTasksRetryPost(ctx context.Context, id string) JamfConnectAPIV1JamfConnectDeploymentsIdTasksRetryPostRequest {
	return JamfConnectAPIV1JamfConnectDeploymentsIdTasksRetryPostRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *JamfConnectAPIService) V1JamfConnectDeploymentsIdTasksRetryPostExecute(r JamfConnectAPIV1JamfConnectDeploymentsIdTasksRetryPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JamfConnectAPIService.V1JamfConnectDeploymentsIdTasksRetryPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/jamf-connect/deployments/{id}/tasks/retry"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.ids == nil {
		return nil, reportError("ids is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.ids
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type JamfConnectAPIV1JamfConnectGetRequest struct {
	ctx context.Context
	ApiService JamfConnectAPI
}

func (r JamfConnectAPIV1JamfConnectGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1JamfConnectGetExecute(r)
}

/*
V1JamfConnectGet Get the Jamf Connect settings that you have access to see 

Get the Jamf Connect settings that you have access to see.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return JamfConnectAPIV1JamfConnectGetRequest
*/
func (a *JamfConnectAPIService) V1JamfConnectGet(ctx context.Context) JamfConnectAPIV1JamfConnectGetRequest {
	return JamfConnectAPIV1JamfConnectGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *JamfConnectAPIService) V1JamfConnectGetExecute(r JamfConnectAPIV1JamfConnectGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JamfConnectAPIService.V1JamfConnectGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/jamf-connect"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type JamfConnectAPIV1JamfConnectHistoryGetRequest struct {
	ctx context.Context
	ApiService JamfConnectAPI
	page *int32
	pageSize *int32
	sort *[]string
	filter *string
}

func (r JamfConnectAPIV1JamfConnectHistoryGetRequest) Page(page int32) JamfConnectAPIV1JamfConnectHistoryGetRequest {
	r.page = &page
	return r
}

func (r JamfConnectAPIV1JamfConnectHistoryGetRequest) PageSize(pageSize int32) JamfConnectAPIV1JamfConnectHistoryGetRequest {
	r.pageSize = &pageSize
	return r
}

// Sorting criteria in the format: property:asc/desc. Default sort order is descending. Multiple sort criteria are supported and must be entered on separate lines in Swagger UI. In the URI the &#39;sort&#39; query param is not duplicated for each sort criterion, e.g., ...&amp;sort&#x3D;name:asc,date:desc. Fields that can be sorted: status, updated
func (r JamfConnectAPIV1JamfConnectHistoryGetRequest) Sort(sort []string) JamfConnectAPIV1JamfConnectHistoryGetRequest {
	r.sort = &sort
	return r
}

// Query in the RSQL format, allowing to filter results. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: status, updated, version This param can be combined with paging and sorting. Example: filter&#x3D;username!&#x3D;admin and details&#x3D;&#x3D;*disabled* and date&lt;2019-12-15
func (r JamfConnectAPIV1JamfConnectHistoryGetRequest) Filter(filter string) JamfConnectAPIV1JamfConnectHistoryGetRequest {
	r.filter = &filter
	return r
}

func (r JamfConnectAPIV1JamfConnectHistoryGetRequest) Execute() (*HistorySearchResults, *http.Response, error) {
	return r.ApiService.V1JamfConnectHistoryGetExecute(r)
}

/*
V1JamfConnectHistoryGet Get Jamf Connect history 

Get Jamf Connect history


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return JamfConnectAPIV1JamfConnectHistoryGetRequest
*/
func (a *JamfConnectAPIService) V1JamfConnectHistoryGet(ctx context.Context) JamfConnectAPIV1JamfConnectHistoryGetRequest {
	return JamfConnectAPIV1JamfConnectHistoryGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return HistorySearchResults
func (a *JamfConnectAPIService) V1JamfConnectHistoryGetExecute(r JamfConnectAPIV1JamfConnectHistoryGetRequest) (*HistorySearchResults, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *HistorySearchResults
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JamfConnectAPIService.V1JamfConnectHistoryGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/jamf-connect/history"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	} else {
		var defaultValue int32 = 0
		r.page = &defaultValue
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page-size", r.pageSize, "")
	} else {
		var defaultValue int32 = 100
		r.pageSize = &defaultValue
	}
	if r.sort != nil {
		t := *r.sort
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "sort", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "sort", t, "multi")
		}
	} else {
		var defaultValue []string = []
		r.sort = &defaultValue
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "")
	} else {
		var defaultValue string = ""
		r.filter = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type JamfConnectAPIV1JamfConnectHistoryPostRequest struct {
	ctx context.Context
	ApiService JamfConnectAPI
	objectHistoryNote *ObjectHistoryNote
}

// history notes to create
func (r JamfConnectAPIV1JamfConnectHistoryPostRequest) ObjectHistoryNote(objectHistoryNote ObjectHistoryNote) JamfConnectAPIV1JamfConnectHistoryPostRequest {
	r.objectHistoryNote = &objectHistoryNote
	return r
}

func (r JamfConnectAPIV1JamfConnectHistoryPostRequest) Execute() (*HrefResponse, *http.Response, error) {
	return r.ApiService.V1JamfConnectHistoryPostExecute(r)
}

/*
V1JamfConnectHistoryPost Add Jamf Connect history notes 

Add Jamf Connect history notes


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return JamfConnectAPIV1JamfConnectHistoryPostRequest
*/
func (a *JamfConnectAPIService) V1JamfConnectHistoryPost(ctx context.Context) JamfConnectAPIV1JamfConnectHistoryPostRequest {
	return JamfConnectAPIV1JamfConnectHistoryPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return HrefResponse
func (a *JamfConnectAPIService) V1JamfConnectHistoryPostExecute(r JamfConnectAPIV1JamfConnectHistoryPostRequest) (*HrefResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *HrefResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JamfConnectAPIService.V1JamfConnectHistoryPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/jamf-connect/history"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.objectHistoryNote == nil {
		return localVarReturnValue, nil, reportError("objectHistoryNote is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.objectHistoryNote
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
