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
	"reflect"
)


type ClientCheckInAPI interface {

	/*
	V2CheckInGet Get Client Check-In settings 

	Gets `Client Check-In` object.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ClientCheckInAPIV2CheckInGetRequest

	Deprecated
	*/
	V2CheckInGet(ctx context.Context) ClientCheckInAPIV2CheckInGetRequest

	// V2CheckInGetExecute executes the request
	//  @return ClientCheckInV2
	// Deprecated
	V2CheckInGetExecute(r ClientCheckInAPIV2CheckInGetRequest) (*ClientCheckInV2, *http.Response, error)

	/*
	V2CheckInHistoryGet Get Client Check-In history object 

	Gets Client Check-In history object


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ClientCheckInAPIV2CheckInHistoryGetRequest

	Deprecated
	*/
	V2CheckInHistoryGet(ctx context.Context) ClientCheckInAPIV2CheckInHistoryGetRequest

	// V2CheckInHistoryGetExecute executes the request
	//  @return HistorySearchResultsV1
	// Deprecated
	V2CheckInHistoryGetExecute(r ClientCheckInAPIV2CheckInHistoryGetRequest) (*HistorySearchResultsV1, *http.Response, error)

	/*
	V2CheckInHistoryPost Add a Note to Client Check-In History 

	Adds Client Check-In history object notes


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ClientCheckInAPIV2CheckInHistoryPostRequest

	Deprecated
	*/
	V2CheckInHistoryPost(ctx context.Context) ClientCheckInAPIV2CheckInHistoryPostRequest

	// V2CheckInHistoryPostExecute executes the request
	//  @return HrefResponse
	// Deprecated
	V2CheckInHistoryPostExecute(r ClientCheckInAPIV2CheckInHistoryPostRequest) (*HrefResponse, *http.Response, error)

	/*
	V2CheckInPut Update Client Check-In object 

	Update Client Check-In object


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ClientCheckInAPIV2CheckInPutRequest

	Deprecated
	*/
	V2CheckInPut(ctx context.Context) ClientCheckInAPIV2CheckInPutRequest

	// V2CheckInPutExecute executes the request
	//  @return ClientCheckInV2
	// Deprecated
	V2CheckInPutExecute(r ClientCheckInAPIV2CheckInPutRequest) (*ClientCheckInV2, *http.Response, error)

	/*
	V3CheckInGet Get Client Check-In settings 

	Gets `Client Check-In` object.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ClientCheckInAPIV3CheckInGetRequest
	*/
	V3CheckInGet(ctx context.Context) ClientCheckInAPIV3CheckInGetRequest

	// V3CheckInGetExecute executes the request
	//  @return ClientCheckInV3
	V3CheckInGetExecute(r ClientCheckInAPIV3CheckInGetRequest) (*ClientCheckInV3, *http.Response, error)

	/*
	V3CheckInHistoryGet Get Client Check-In history object 

	Gets Client Check-In history object


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ClientCheckInAPIV3CheckInHistoryGetRequest
	*/
	V3CheckInHistoryGet(ctx context.Context) ClientCheckInAPIV3CheckInHistoryGetRequest

	// V3CheckInHistoryGetExecute executes the request
	//  @return HistorySearchResultsV1
	V3CheckInHistoryGetExecute(r ClientCheckInAPIV3CheckInHistoryGetRequest) (*HistorySearchResultsV1, *http.Response, error)

	/*
	V3CheckInHistoryPost Add a Note to Client Check-In History 

	Adds Client Check-In history object notes


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ClientCheckInAPIV3CheckInHistoryPostRequest
	*/
	V3CheckInHistoryPost(ctx context.Context) ClientCheckInAPIV3CheckInHistoryPostRequest

	// V3CheckInHistoryPostExecute executes the request
	//  @return HrefResponse
	V3CheckInHistoryPostExecute(r ClientCheckInAPIV3CheckInHistoryPostRequest) (*HrefResponse, *http.Response, error)

	/*
	V3CheckInPut Update Client Check-In object 

	Update Client Check-In object


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ClientCheckInAPIV3CheckInPutRequest
	*/
	V3CheckInPut(ctx context.Context) ClientCheckInAPIV3CheckInPutRequest

	// V3CheckInPutExecute executes the request
	//  @return ClientCheckInV3
	V3CheckInPutExecute(r ClientCheckInAPIV3CheckInPutRequest) (*ClientCheckInV3, *http.Response, error)
}

// ClientCheckInAPIService ClientCheckInAPI service
type ClientCheckInAPIService service

type ClientCheckInAPIV2CheckInGetRequest struct {
	ctx context.Context
	ApiService ClientCheckInAPI
}

func (r ClientCheckInAPIV2CheckInGetRequest) Execute() (*ClientCheckInV2, *http.Response, error) {
	return r.ApiService.V2CheckInGetExecute(r)
}

/*
V2CheckInGet Get Client Check-In settings 

Gets `Client Check-In` object.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ClientCheckInAPIV2CheckInGetRequest

Deprecated
*/
func (a *ClientCheckInAPIService) V2CheckInGet(ctx context.Context) ClientCheckInAPIV2CheckInGetRequest {
	return ClientCheckInAPIV2CheckInGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ClientCheckInV2
// Deprecated
func (a *ClientCheckInAPIService) V2CheckInGetExecute(r ClientCheckInAPIV2CheckInGetRequest) (*ClientCheckInV2, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ClientCheckInV2
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientCheckInAPIService.V2CheckInGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/check-in"

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

type ClientCheckInAPIV2CheckInHistoryGetRequest struct {
	ctx context.Context
	ApiService ClientCheckInAPI
	page *int64
	pageSize *int64
	sort *[]string
	filter *string
}

func (r ClientCheckInAPIV2CheckInHistoryGetRequest) Page(page int64) ClientCheckInAPIV2CheckInHistoryGetRequest {
	r.page = &page
	return r
}

func (r ClientCheckInAPIV2CheckInHistoryGetRequest) PageSize(pageSize int64) ClientCheckInAPIV2CheckInHistoryGetRequest {
	r.pageSize = &pageSize
	return r
}

// Sorting criteria in the format: property:asc/desc. Default sort is name:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;date:desc,username:asc 
func (r ClientCheckInAPIV2CheckInHistoryGetRequest) Sort(sort []string) ClientCheckInAPIV2CheckInHistoryGetRequest {
	r.sort = &sort
	return r
}

// Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: filter&#x3D;username!&#x3D;admin and details&#x3D;&#x3D;*disabled* and date&lt;2019-12-15
func (r ClientCheckInAPIV2CheckInHistoryGetRequest) Filter(filter string) ClientCheckInAPIV2CheckInHistoryGetRequest {
	r.filter = &filter
	return r
}

func (r ClientCheckInAPIV2CheckInHistoryGetRequest) Execute() (*HistorySearchResultsV1, *http.Response, error) {
	return r.ApiService.V2CheckInHistoryGetExecute(r)
}

/*
V2CheckInHistoryGet Get Client Check-In history object 

Gets Client Check-In history object


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ClientCheckInAPIV2CheckInHistoryGetRequest

Deprecated
*/
func (a *ClientCheckInAPIService) V2CheckInHistoryGet(ctx context.Context) ClientCheckInAPIV2CheckInHistoryGetRequest {
	return ClientCheckInAPIV2CheckInHistoryGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return HistorySearchResultsV1
// Deprecated
func (a *ClientCheckInAPIService) V2CheckInHistoryGetExecute(r ClientCheckInAPIV2CheckInHistoryGetRequest) (*HistorySearchResultsV1, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *HistorySearchResultsV1
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientCheckInAPIService.V2CheckInHistoryGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/check-in/history"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	} else {
		var defaultValue int64 = 0
		r.page = &defaultValue
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page-size", r.pageSize, "")
	} else {
		var defaultValue int64 = 100
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
		defaultValue := []string{"date:desc"}
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

type ClientCheckInAPIV2CheckInHistoryPostRequest struct {
	ctx context.Context
	ApiService ClientCheckInAPI
	objectHistoryNote *ObjectHistoryNote
}

// history notes to create
func (r ClientCheckInAPIV2CheckInHistoryPostRequest) ObjectHistoryNote(objectHistoryNote ObjectHistoryNote) ClientCheckInAPIV2CheckInHistoryPostRequest {
	r.objectHistoryNote = &objectHistoryNote
	return r
}

func (r ClientCheckInAPIV2CheckInHistoryPostRequest) Execute() (*HrefResponse, *http.Response, error) {
	return r.ApiService.V2CheckInHistoryPostExecute(r)
}

/*
V2CheckInHistoryPost Add a Note to Client Check-In History 

Adds Client Check-In history object notes


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ClientCheckInAPIV2CheckInHistoryPostRequest

Deprecated
*/
func (a *ClientCheckInAPIService) V2CheckInHistoryPost(ctx context.Context) ClientCheckInAPIV2CheckInHistoryPostRequest {
	return ClientCheckInAPIV2CheckInHistoryPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return HrefResponse
// Deprecated
func (a *ClientCheckInAPIService) V2CheckInHistoryPostExecute(r ClientCheckInAPIV2CheckInHistoryPostRequest) (*HrefResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *HrefResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientCheckInAPIService.V2CheckInHistoryPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/check-in/history"

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

type ClientCheckInAPIV2CheckInPutRequest struct {
	ctx context.Context
	ApiService ClientCheckInAPI
	clientCheckInV2 *ClientCheckInV2
}

// Client Check-In object to update
func (r ClientCheckInAPIV2CheckInPutRequest) ClientCheckInV2(clientCheckInV2 ClientCheckInV2) ClientCheckInAPIV2CheckInPutRequest {
	r.clientCheckInV2 = &clientCheckInV2
	return r
}

func (r ClientCheckInAPIV2CheckInPutRequest) Execute() (*ClientCheckInV2, *http.Response, error) {
	return r.ApiService.V2CheckInPutExecute(r)
}

/*
V2CheckInPut Update Client Check-In object 

Update Client Check-In object


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ClientCheckInAPIV2CheckInPutRequest

Deprecated
*/
func (a *ClientCheckInAPIService) V2CheckInPut(ctx context.Context) ClientCheckInAPIV2CheckInPutRequest {
	return ClientCheckInAPIV2CheckInPutRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ClientCheckInV2
// Deprecated
func (a *ClientCheckInAPIService) V2CheckInPutExecute(r ClientCheckInAPIV2CheckInPutRequest) (*ClientCheckInV2, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ClientCheckInV2
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientCheckInAPIService.V2CheckInPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/check-in"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.clientCheckInV2 == nil {
		return localVarReturnValue, nil, reportError("clientCheckInV2 is required and must be specified")
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
	localVarPostBody = r.clientCheckInV2
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

type ClientCheckInAPIV3CheckInGetRequest struct {
	ctx context.Context
	ApiService ClientCheckInAPI
}

func (r ClientCheckInAPIV3CheckInGetRequest) Execute() (*ClientCheckInV3, *http.Response, error) {
	return r.ApiService.V3CheckInGetExecute(r)
}

/*
V3CheckInGet Get Client Check-In settings 

Gets `Client Check-In` object.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ClientCheckInAPIV3CheckInGetRequest
*/
func (a *ClientCheckInAPIService) V3CheckInGet(ctx context.Context) ClientCheckInAPIV3CheckInGetRequest {
	return ClientCheckInAPIV3CheckInGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ClientCheckInV3
func (a *ClientCheckInAPIService) V3CheckInGetExecute(r ClientCheckInAPIV3CheckInGetRequest) (*ClientCheckInV3, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ClientCheckInV3
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientCheckInAPIService.V3CheckInGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/check-in"

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

type ClientCheckInAPIV3CheckInHistoryGetRequest struct {
	ctx context.Context
	ApiService ClientCheckInAPI
	page *int64
	pageSize *int64
	sort *[]string
	filter *string
}

func (r ClientCheckInAPIV3CheckInHistoryGetRequest) Page(page int64) ClientCheckInAPIV3CheckInHistoryGetRequest {
	r.page = &page
	return r
}

func (r ClientCheckInAPIV3CheckInHistoryGetRequest) PageSize(pageSize int64) ClientCheckInAPIV3CheckInHistoryGetRequest {
	r.pageSize = &pageSize
	return r
}

// Sorting criteria in the format: property:asc/desc. Default sort is name:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;date:desc,username:asc 
func (r ClientCheckInAPIV3CheckInHistoryGetRequest) Sort(sort []string) ClientCheckInAPIV3CheckInHistoryGetRequest {
	r.sort = &sort
	return r
}

// Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: filter&#x3D;username!&#x3D;admin and details&#x3D;&#x3D;*disabled* and date&lt;2019-12-15
func (r ClientCheckInAPIV3CheckInHistoryGetRequest) Filter(filter string) ClientCheckInAPIV3CheckInHistoryGetRequest {
	r.filter = &filter
	return r
}

func (r ClientCheckInAPIV3CheckInHistoryGetRequest) Execute() (*HistorySearchResultsV1, *http.Response, error) {
	return r.ApiService.V3CheckInHistoryGetExecute(r)
}

/*
V3CheckInHistoryGet Get Client Check-In history object 

Gets Client Check-In history object


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ClientCheckInAPIV3CheckInHistoryGetRequest
*/
func (a *ClientCheckInAPIService) V3CheckInHistoryGet(ctx context.Context) ClientCheckInAPIV3CheckInHistoryGetRequest {
	return ClientCheckInAPIV3CheckInHistoryGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return HistorySearchResultsV1
func (a *ClientCheckInAPIService) V3CheckInHistoryGetExecute(r ClientCheckInAPIV3CheckInHistoryGetRequest) (*HistorySearchResultsV1, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *HistorySearchResultsV1
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientCheckInAPIService.V3CheckInHistoryGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/check-in/history"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	} else {
		var defaultValue int64 = 0
		r.page = &defaultValue
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page-size", r.pageSize, "")
	} else {
		var defaultValue int64 = 100
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
		defaultValue := []string{"date:desc"}
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

type ClientCheckInAPIV3CheckInHistoryPostRequest struct {
	ctx context.Context
	ApiService ClientCheckInAPI
	objectHistoryNote *ObjectHistoryNote
}

// history notes to create
func (r ClientCheckInAPIV3CheckInHistoryPostRequest) ObjectHistoryNote(objectHistoryNote ObjectHistoryNote) ClientCheckInAPIV3CheckInHistoryPostRequest {
	r.objectHistoryNote = &objectHistoryNote
	return r
}

func (r ClientCheckInAPIV3CheckInHistoryPostRequest) Execute() (*HrefResponse, *http.Response, error) {
	return r.ApiService.V3CheckInHistoryPostExecute(r)
}

/*
V3CheckInHistoryPost Add a Note to Client Check-In History 

Adds Client Check-In history object notes


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ClientCheckInAPIV3CheckInHistoryPostRequest
*/
func (a *ClientCheckInAPIService) V3CheckInHistoryPost(ctx context.Context) ClientCheckInAPIV3CheckInHistoryPostRequest {
	return ClientCheckInAPIV3CheckInHistoryPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return HrefResponse
func (a *ClientCheckInAPIService) V3CheckInHistoryPostExecute(r ClientCheckInAPIV3CheckInHistoryPostRequest) (*HrefResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *HrefResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientCheckInAPIService.V3CheckInHistoryPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/check-in/history"

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

type ClientCheckInAPIV3CheckInPutRequest struct {
	ctx context.Context
	ApiService ClientCheckInAPI
	clientCheckInV3 *ClientCheckInV3
}

// Client Check-In object to update
func (r ClientCheckInAPIV3CheckInPutRequest) ClientCheckInV3(clientCheckInV3 ClientCheckInV3) ClientCheckInAPIV3CheckInPutRequest {
	r.clientCheckInV3 = &clientCheckInV3
	return r
}

func (r ClientCheckInAPIV3CheckInPutRequest) Execute() (*ClientCheckInV3, *http.Response, error) {
	return r.ApiService.V3CheckInPutExecute(r)
}

/*
V3CheckInPut Update Client Check-In object 

Update Client Check-In object


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ClientCheckInAPIV3CheckInPutRequest
*/
func (a *ClientCheckInAPIService) V3CheckInPut(ctx context.Context) ClientCheckInAPIV3CheckInPutRequest {
	return ClientCheckInAPIV3CheckInPutRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ClientCheckInV3
func (a *ClientCheckInAPIService) V3CheckInPutExecute(r ClientCheckInAPIV3CheckInPutRequest) (*ClientCheckInV3, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ClientCheckInV3
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientCheckInAPIService.V3CheckInPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/check-in"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.clientCheckInV3 == nil {
		return localVarReturnValue, nil, reportError("clientCheckInV3 is required and must be specified")
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
	localVarPostBody = r.clientCheckInV3
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
