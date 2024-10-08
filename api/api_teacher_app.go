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


type TeacherAppAPI interface {

	/*
	V1TeacherAppGet Get the Jamf Teacher settings that you have access to see 

	Get the Jamf Teacher settings that you have access to see.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return TeacherAppAPIV1TeacherAppGetRequest
	*/
	V1TeacherAppGet(ctx context.Context) TeacherAppAPIV1TeacherAppGetRequest

	// V1TeacherAppGetExecute executes the request
	//  @return TeacherSettingsResponse
	V1TeacherAppGetExecute(r TeacherAppAPIV1TeacherAppGetRequest) (*TeacherSettingsResponse, *http.Response, error)

	/*
	V1TeacherAppHistoryGet Get Jamf Teacher app settings history 

	Gets Jamf Teacher app settings history


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return TeacherAppAPIV1TeacherAppHistoryGetRequest
	*/
	V1TeacherAppHistoryGet(ctx context.Context) TeacherAppAPIV1TeacherAppHistoryGetRequest

	// V1TeacherAppHistoryGetExecute executes the request
	//  @return HistorySearchResults
	V1TeacherAppHistoryGetExecute(r TeacherAppAPIV1TeacherAppHistoryGetRequest) (*HistorySearchResults, *http.Response, error)

	/*
	V1TeacherAppHistoryPost Add Jamf Teacher app settings history notes 

	Adds Jamf Teacher app settings history notes


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return TeacherAppAPIV1TeacherAppHistoryPostRequest
	*/
	V1TeacherAppHistoryPost(ctx context.Context) TeacherAppAPIV1TeacherAppHistoryPostRequest

	// V1TeacherAppHistoryPostExecute executes the request
	//  @return HrefResponse
	V1TeacherAppHistoryPostExecute(r TeacherAppAPIV1TeacherAppHistoryPostRequest) (*HrefResponse, *http.Response, error)

	/*
	V1TeacherAppPut Update a Jamf Teacher settings object 

	Update a Jamf Teacher settings object.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return TeacherAppAPIV1TeacherAppPutRequest
	*/
	V1TeacherAppPut(ctx context.Context) TeacherAppAPIV1TeacherAppPutRequest

	// V1TeacherAppPutExecute executes the request
	//  @return TeacherSettingsResponse
	V1TeacherAppPutExecute(r TeacherAppAPIV1TeacherAppPutRequest) (*TeacherSettingsResponse, *http.Response, error)
}

// TeacherAppAPIService TeacherAppAPI service
type TeacherAppAPIService service

type TeacherAppAPIV1TeacherAppGetRequest struct {
	ctx context.Context
	ApiService TeacherAppAPI
}

func (r TeacherAppAPIV1TeacherAppGetRequest) Execute() (*TeacherSettingsResponse, *http.Response, error) {
	return r.ApiService.V1TeacherAppGetExecute(r)
}

/*
V1TeacherAppGet Get the Jamf Teacher settings that you have access to see 

Get the Jamf Teacher settings that you have access to see.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return TeacherAppAPIV1TeacherAppGetRequest
*/
func (a *TeacherAppAPIService) V1TeacherAppGet(ctx context.Context) TeacherAppAPIV1TeacherAppGetRequest {
	return TeacherAppAPIV1TeacherAppGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return TeacherSettingsResponse
func (a *TeacherAppAPIService) V1TeacherAppGetExecute(r TeacherAppAPIV1TeacherAppGetRequest) (*TeacherSettingsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TeacherSettingsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeacherAppAPIService.V1TeacherAppGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/teacher-app"

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

type TeacherAppAPIV1TeacherAppHistoryGetRequest struct {
	ctx context.Context
	ApiService TeacherAppAPI
	page *int64
	pageSize *int64
	sort *[]string
	filter *string
}

func (r TeacherAppAPIV1TeacherAppHistoryGetRequest) Page(page int64) TeacherAppAPIV1TeacherAppHistoryGetRequest {
	r.page = &page
	return r
}

func (r TeacherAppAPIV1TeacherAppHistoryGetRequest) PageSize(pageSize int64) TeacherAppAPIV1TeacherAppHistoryGetRequest {
	r.pageSize = &pageSize
	return r
}

// Sorting criteria in the format: property:asc/desc. Default sort order is descending. Multiple sort criteria are supported and must be entered on separate lines in Swagger UI. In the URI the &#39;sort&#39; query param is not duplicated for each sort criterion, e.g., ...&amp;sort&#x3D;name:asc,date:desc. Fields that can be sorted: status, updated
func (r TeacherAppAPIV1TeacherAppHistoryGetRequest) Sort(sort []string) TeacherAppAPIV1TeacherAppHistoryGetRequest {
	r.sort = &sort
	return r
}

// Query in the RSQL format, allowing to filter results. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: status, updated, version This param can be combined with paging and sorting. Example: filter&#x3D;username!&#x3D;admin and details&#x3D;&#x3D;*disabled* and date&lt;2019-12-15
func (r TeacherAppAPIV1TeacherAppHistoryGetRequest) Filter(filter string) TeacherAppAPIV1TeacherAppHistoryGetRequest {
	r.filter = &filter
	return r
}

func (r TeacherAppAPIV1TeacherAppHistoryGetRequest) Execute() (*HistorySearchResults, *http.Response, error) {
	return r.ApiService.V1TeacherAppHistoryGetExecute(r)
}

/*
V1TeacherAppHistoryGet Get Jamf Teacher app settings history 

Gets Jamf Teacher app settings history


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return TeacherAppAPIV1TeacherAppHistoryGetRequest
*/
func (a *TeacherAppAPIService) V1TeacherAppHistoryGet(ctx context.Context) TeacherAppAPIV1TeacherAppHistoryGetRequest {
	return TeacherAppAPIV1TeacherAppHistoryGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return HistorySearchResults
func (a *TeacherAppAPIService) V1TeacherAppHistoryGetExecute(r TeacherAppAPIV1TeacherAppHistoryGetRequest) (*HistorySearchResults, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *HistorySearchResults
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeacherAppAPIService.V1TeacherAppHistoryGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/teacher-app/history"

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
		defaultValue := []string{}
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

type TeacherAppAPIV1TeacherAppHistoryPostRequest struct {
	ctx context.Context
	ApiService TeacherAppAPI
	objectHistoryNote *ObjectHistoryNote
}

// history notes to create
func (r TeacherAppAPIV1TeacherAppHistoryPostRequest) ObjectHistoryNote(objectHistoryNote ObjectHistoryNote) TeacherAppAPIV1TeacherAppHistoryPostRequest {
	r.objectHistoryNote = &objectHistoryNote
	return r
}

func (r TeacherAppAPIV1TeacherAppHistoryPostRequest) Execute() (*HrefResponse, *http.Response, error) {
	return r.ApiService.V1TeacherAppHistoryPostExecute(r)
}

/*
V1TeacherAppHistoryPost Add Jamf Teacher app settings history notes 

Adds Jamf Teacher app settings history notes


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return TeacherAppAPIV1TeacherAppHistoryPostRequest
*/
func (a *TeacherAppAPIService) V1TeacherAppHistoryPost(ctx context.Context) TeacherAppAPIV1TeacherAppHistoryPostRequest {
	return TeacherAppAPIV1TeacherAppHistoryPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return HrefResponse
func (a *TeacherAppAPIService) V1TeacherAppHistoryPostExecute(r TeacherAppAPIV1TeacherAppHistoryPostRequest) (*HrefResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *HrefResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeacherAppAPIService.V1TeacherAppHistoryPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/teacher-app/history"

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

type TeacherAppAPIV1TeacherAppPutRequest struct {
	ctx context.Context
	ApiService TeacherAppAPI
	teacherSettingsRequest *TeacherSettingsRequest
}

// Teacher settings to create.
func (r TeacherAppAPIV1TeacherAppPutRequest) TeacherSettingsRequest(teacherSettingsRequest TeacherSettingsRequest) TeacherAppAPIV1TeacherAppPutRequest {
	r.teacherSettingsRequest = &teacherSettingsRequest
	return r
}

func (r TeacherAppAPIV1TeacherAppPutRequest) Execute() (*TeacherSettingsResponse, *http.Response, error) {
	return r.ApiService.V1TeacherAppPutExecute(r)
}

/*
V1TeacherAppPut Update a Jamf Teacher settings object 

Update a Jamf Teacher settings object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return TeacherAppAPIV1TeacherAppPutRequest
*/
func (a *TeacherAppAPIService) V1TeacherAppPut(ctx context.Context) TeacherAppAPIV1TeacherAppPutRequest {
	return TeacherAppAPIV1TeacherAppPutRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return TeacherSettingsResponse
func (a *TeacherAppAPIService) V1TeacherAppPutExecute(r TeacherAppAPIV1TeacherAppPutRequest) (*TeacherSettingsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TeacherSettingsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeacherAppAPIService.V1TeacherAppPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/teacher-app"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.teacherSettingsRequest == nil {
		return localVarReturnValue, nil, reportError("teacherSettingsRequest is required and must be specified")
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
	localVarPostBody = r.teacherSettingsRequest
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
		if localVarHTTPResponse.StatusCode == 500 {
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
