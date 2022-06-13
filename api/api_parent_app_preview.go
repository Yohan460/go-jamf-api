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
	"io/ioutil"
	"net/http"
	"net/url"
)


type ParentAppPreviewApi interface {

	/*
	V1ParentAppGet Get the current Jamf Parent app settings 

	Get the current Jamf Parent app settings


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiV1ParentAppGetRequest
	*/
	V1ParentAppGet(ctx context.Context) ApiV1ParentAppGetRequest

	// V1ParentAppGetExecute executes the request
	//  @return ParentApp
	V1ParentAppGetExecute(r ApiV1ParentAppGetRequest) (*ParentApp, *http.Response, error)

	/*
	V1ParentAppHistoryGet Get Jamf Parent app settings history 

	Gets Jamf Parent app settings history


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiV1ParentAppHistoryGetRequest
	*/
	V1ParentAppHistoryGet(ctx context.Context) ApiV1ParentAppHistoryGetRequest

	// V1ParentAppHistoryGetExecute executes the request
	//  @return HistorySearchResults
	V1ParentAppHistoryGetExecute(r ApiV1ParentAppHistoryGetRequest) (*HistorySearchResults, *http.Response, error)

	/*
	V1ParentAppHistoryPost Add Jamf Parent app settings history notes 

	Adds Jamf Parent app settings history notes


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiV1ParentAppHistoryPostRequest
	*/
	V1ParentAppHistoryPost(ctx context.Context) ApiV1ParentAppHistoryPostRequest

	// V1ParentAppHistoryPostExecute executes the request
	//  @return ObjectHistory
	V1ParentAppHistoryPostExecute(r ApiV1ParentAppHistoryPostRequest) (*ObjectHistory, *http.Response, error)

	/*
	V1ParentAppPut Update Jamf Parent app settings 

	Update Jamf Parent app settings


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiV1ParentAppPutRequest
	*/
	V1ParentAppPut(ctx context.Context) ApiV1ParentAppPutRequest

	// V1ParentAppPutExecute executes the request
	//  @return ParentApp
	V1ParentAppPutExecute(r ApiV1ParentAppPutRequest) (*ParentApp, *http.Response, error)
}

// ParentAppPreviewApiService ParentAppPreviewApi service
type ParentAppPreviewApiService service

type ApiV1ParentAppGetRequest struct {
	ctx context.Context
	ApiService ParentAppPreviewApi
}

func (r ApiV1ParentAppGetRequest) Execute() (*ParentApp, *http.Response, error) {
	return r.ApiService.V1ParentAppGetExecute(r)
}

/*
V1ParentAppGet Get the current Jamf Parent app settings 

Get the current Jamf Parent app settings


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV1ParentAppGetRequest
*/
func (a *ParentAppPreviewApiService) V1ParentAppGet(ctx context.Context) ApiV1ParentAppGetRequest {
	return ApiV1ParentAppGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ParentApp
func (a *ParentAppPreviewApiService) V1ParentAppGetExecute(r ApiV1ParentAppGetRequest) (*ParentApp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ParentApp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ParentAppPreviewApiService.V1ParentAppGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/parent-app"

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
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiV1ParentAppHistoryGetRequest struct {
	ctx context.Context
	ApiService ParentAppPreviewApi
	page *int32
	pageSize *int32
	filter *string
	sort *string
}

func (r ApiV1ParentAppHistoryGetRequest) Page(page int32) ApiV1ParentAppHistoryGetRequest {
	r.page = &page
	return r
}

func (r ApiV1ParentAppHistoryGetRequest) PageSize(pageSize int32) ApiV1ParentAppHistoryGetRequest {
	r.pageSize = &pageSize
	return r
}

// Query in the RSQL format, allowing to filter results. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: status, updated, version This param can be combined with paging and sorting. Example: filter&#x3D;username!&#x3D;admin and details&#x3D;&#x3D;*disabled* and date&lt;2019-12-15
func (r ApiV1ParentAppHistoryGetRequest) Filter(filter string) ApiV1ParentAppHistoryGetRequest {
	r.filter = &filter
	return r
}

// Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;date:desc,name:asc 
func (r ApiV1ParentAppHistoryGetRequest) Sort(sort string) ApiV1ParentAppHistoryGetRequest {
	r.sort = &sort
	return r
}

func (r ApiV1ParentAppHistoryGetRequest) Execute() (*HistorySearchResults, *http.Response, error) {
	return r.ApiService.V1ParentAppHistoryGetExecute(r)
}

/*
V1ParentAppHistoryGet Get Jamf Parent app settings history 

Gets Jamf Parent app settings history


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV1ParentAppHistoryGetRequest
*/
func (a *ParentAppPreviewApiService) V1ParentAppHistoryGet(ctx context.Context) ApiV1ParentAppHistoryGetRequest {
	return ApiV1ParentAppHistoryGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return HistorySearchResults
func (a *ParentAppPreviewApiService) V1ParentAppHistoryGetExecute(r ApiV1ParentAppHistoryGetRequest) (*HistorySearchResults, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *HistorySearchResults
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ParentAppPreviewApiService.V1ParentAppHistoryGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/parent-app/history"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("page-size", parameterToString(*r.pageSize, ""))
	}
	if r.filter != nil {
		localVarQueryParams.Add("filter", parameterToString(*r.filter, ""))
	}
	if r.sort != nil {
		localVarQueryParams.Add("sort", parameterToString(*r.sort, ""))
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiV1ParentAppHistoryPostRequest struct {
	ctx context.Context
	ApiService ParentAppPreviewApi
	objectHistoryNote *ObjectHistoryNote
}

// history notes to create
func (r ApiV1ParentAppHistoryPostRequest) ObjectHistoryNote(objectHistoryNote ObjectHistoryNote) ApiV1ParentAppHistoryPostRequest {
	r.objectHistoryNote = &objectHistoryNote
	return r
}

func (r ApiV1ParentAppHistoryPostRequest) Execute() (*ObjectHistory, *http.Response, error) {
	return r.ApiService.V1ParentAppHistoryPostExecute(r)
}

/*
V1ParentAppHistoryPost Add Jamf Parent app settings history notes 

Adds Jamf Parent app settings history notes


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV1ParentAppHistoryPostRequest
*/
func (a *ParentAppPreviewApiService) V1ParentAppHistoryPost(ctx context.Context) ApiV1ParentAppHistoryPostRequest {
	return ApiV1ParentAppHistoryPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ObjectHistory
func (a *ParentAppPreviewApiService) V1ParentAppHistoryPostExecute(r ApiV1ParentAppHistoryPostRequest) (*ObjectHistory, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ObjectHistory
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ParentAppPreviewApiService.V1ParentAppHistoryPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/parent-app/history"

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiV1ParentAppPutRequest struct {
	ctx context.Context
	ApiService ParentAppPreviewApi
	parentApp *ParentApp
}

// Jamf Parent app settings to save.
func (r ApiV1ParentAppPutRequest) ParentApp(parentApp ParentApp) ApiV1ParentAppPutRequest {
	r.parentApp = &parentApp
	return r
}

func (r ApiV1ParentAppPutRequest) Execute() (*ParentApp, *http.Response, error) {
	return r.ApiService.V1ParentAppPutExecute(r)
}

/*
V1ParentAppPut Update Jamf Parent app settings 

Update Jamf Parent app settings


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV1ParentAppPutRequest
*/
func (a *ParentAppPreviewApiService) V1ParentAppPut(ctx context.Context) ApiV1ParentAppPutRequest {
	return ApiV1ParentAppPutRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ParentApp
func (a *ParentAppPreviewApiService) V1ParentAppPutExecute(r ApiV1ParentAppPutRequest) (*ParentApp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ParentApp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ParentAppPreviewApiService.V1ParentAppPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/parent-app"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.parentApp == nil {
		return localVarReturnValue, nil, reportError("parentApp is required and must be specified")
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
	localVarPostBody = r.parentApp
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
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
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
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
