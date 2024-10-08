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


type DepartmentsAPI interface {

	/*
	V1DepartmentsDeleteMultiplePost Deletes all departments by ids passed in body 

	Deletes all departments by ids passed in body


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return DepartmentsAPIV1DepartmentsDeleteMultiplePostRequest
	*/
	V1DepartmentsDeleteMultiplePost(ctx context.Context) DepartmentsAPIV1DepartmentsDeleteMultiplePostRequest

	// V1DepartmentsDeleteMultiplePostExecute executes the request
	V1DepartmentsDeleteMultiplePostExecute(r DepartmentsAPIV1DepartmentsDeleteMultiplePostRequest) (*http.Response, error)

	/*
	V1DepartmentsGet Search for Departments 

	Search for Departments


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return DepartmentsAPIV1DepartmentsGetRequest
	*/
	V1DepartmentsGet(ctx context.Context) DepartmentsAPIV1DepartmentsGetRequest

	// V1DepartmentsGetExecute executes the request
	//  @return DepartmentsSearchResults
	V1DepartmentsGetExecute(r DepartmentsAPIV1DepartmentsGetRequest) (*DepartmentsSearchResults, *http.Response, error)

	/*
	V1DepartmentsIdDelete Remove specified department record 

	Removes specified department record


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id instance id of department record
	@return DepartmentsAPIV1DepartmentsIdDeleteRequest
	*/
	V1DepartmentsIdDelete(ctx context.Context, id string) DepartmentsAPIV1DepartmentsIdDeleteRequest

	// V1DepartmentsIdDeleteExecute executes the request
	V1DepartmentsIdDeleteExecute(r DepartmentsAPIV1DepartmentsIdDeleteRequest) (*http.Response, error)

	/*
	V1DepartmentsIdGet Get specified Department object 

	Gets specified Department object


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id instance id of department record
	@return DepartmentsAPIV1DepartmentsIdGetRequest
	*/
	V1DepartmentsIdGet(ctx context.Context, id string) DepartmentsAPIV1DepartmentsIdGetRequest

	// V1DepartmentsIdGetExecute executes the request
	//  @return Department
	V1DepartmentsIdGetExecute(r DepartmentsAPIV1DepartmentsIdGetRequest) (*Department, *http.Response, error)

	/*
	V1DepartmentsIdHistoryGet Get specified Department history object 

	Gets specified Department history object


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id instance id of department history record
	@return DepartmentsAPIV1DepartmentsIdHistoryGetRequest
	*/
	V1DepartmentsIdHistoryGet(ctx context.Context, id string) DepartmentsAPIV1DepartmentsIdHistoryGetRequest

	// V1DepartmentsIdHistoryGetExecute executes the request
	//  @return HistorySearchResults
	V1DepartmentsIdHistoryGetExecute(r DepartmentsAPIV1DepartmentsIdHistoryGetRequest) (*HistorySearchResults, *http.Response, error)

	/*
	V1DepartmentsIdHistoryPost Add specified Department history object notes 

	Adds specified Department history object notes


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id instance id of department history record
	@return DepartmentsAPIV1DepartmentsIdHistoryPostRequest
	*/
	V1DepartmentsIdHistoryPost(ctx context.Context, id string) DepartmentsAPIV1DepartmentsIdHistoryPostRequest

	// V1DepartmentsIdHistoryPostExecute executes the request
	//  @return HrefResponse
	V1DepartmentsIdHistoryPostExecute(r DepartmentsAPIV1DepartmentsIdHistoryPostRequest) (*HrefResponse, *http.Response, error)

	/*
	V1DepartmentsIdPut Update specified department object 

	Update specified department object


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id instance id of department record
	@return DepartmentsAPIV1DepartmentsIdPutRequest
	*/
	V1DepartmentsIdPut(ctx context.Context, id string) DepartmentsAPIV1DepartmentsIdPutRequest

	// V1DepartmentsIdPutExecute executes the request
	//  @return Department
	V1DepartmentsIdPutExecute(r DepartmentsAPIV1DepartmentsIdPutRequest) (*Department, *http.Response, error)

	/*
	V1DepartmentsPost Create department record 

	Create department record


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return DepartmentsAPIV1DepartmentsPostRequest
	*/
	V1DepartmentsPost(ctx context.Context) DepartmentsAPIV1DepartmentsPostRequest

	// V1DepartmentsPostExecute executes the request
	//  @return HrefResponse
	V1DepartmentsPostExecute(r DepartmentsAPIV1DepartmentsPostRequest) (*HrefResponse, *http.Response, error)
}

// DepartmentsAPIService DepartmentsAPI service
type DepartmentsAPIService service

type DepartmentsAPIV1DepartmentsDeleteMultiplePostRequest struct {
	ctx context.Context
	ApiService DepartmentsAPI
	ids *Ids
}

// ids of departments to be deleted. pass in an array of ids
func (r DepartmentsAPIV1DepartmentsDeleteMultiplePostRequest) Ids(ids Ids) DepartmentsAPIV1DepartmentsDeleteMultiplePostRequest {
	r.ids = &ids
	return r
}

func (r DepartmentsAPIV1DepartmentsDeleteMultiplePostRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1DepartmentsDeleteMultiplePostExecute(r)
}

/*
V1DepartmentsDeleteMultiplePost Deletes all departments by ids passed in body 

Deletes all departments by ids passed in body


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return DepartmentsAPIV1DepartmentsDeleteMultiplePostRequest
*/
func (a *DepartmentsAPIService) V1DepartmentsDeleteMultiplePost(ctx context.Context) DepartmentsAPIV1DepartmentsDeleteMultiplePostRequest {
	return DepartmentsAPIV1DepartmentsDeleteMultiplePostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *DepartmentsAPIService) V1DepartmentsDeleteMultiplePostExecute(r DepartmentsAPIV1DepartmentsDeleteMultiplePostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DepartmentsAPIService.V1DepartmentsDeleteMultiplePost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/departments/delete-multiple"

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
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type DepartmentsAPIV1DepartmentsGetRequest struct {
	ctx context.Context
	ApiService DepartmentsAPI
	page *int64
	pageSize *int64
	sort *[]string
	filter *string
}

func (r DepartmentsAPIV1DepartmentsGetRequest) Page(page int64) DepartmentsAPIV1DepartmentsGetRequest {
	r.page = &page
	return r
}

func (r DepartmentsAPIV1DepartmentsGetRequest) PageSize(pageSize int64) DepartmentsAPIV1DepartmentsGetRequest {
	r.pageSize = &pageSize
	return r
}

// Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;date:desc,name:asc 
func (r DepartmentsAPIV1DepartmentsGetRequest) Sort(sort []string) DepartmentsAPIV1DepartmentsGetRequest {
	r.sort = &sort
	return r
}

// Query in the RSQL format, allowing to filter department collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, name. Example: name&#x3D;&#x3D;\&quot;*department*\&quot;
func (r DepartmentsAPIV1DepartmentsGetRequest) Filter(filter string) DepartmentsAPIV1DepartmentsGetRequest {
	r.filter = &filter
	return r
}

func (r DepartmentsAPIV1DepartmentsGetRequest) Execute() (*DepartmentsSearchResults, *http.Response, error) {
	return r.ApiService.V1DepartmentsGetExecute(r)
}

/*
V1DepartmentsGet Search for Departments 

Search for Departments


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return DepartmentsAPIV1DepartmentsGetRequest
*/
func (a *DepartmentsAPIService) V1DepartmentsGet(ctx context.Context) DepartmentsAPIV1DepartmentsGetRequest {
	return DepartmentsAPIV1DepartmentsGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DepartmentsSearchResults
func (a *DepartmentsAPIService) V1DepartmentsGetExecute(r DepartmentsAPIV1DepartmentsGetRequest) (*DepartmentsSearchResults, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DepartmentsSearchResults
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DepartmentsAPIService.V1DepartmentsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/departments"

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
		defaultValue := []string{"id:asc"}
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

type DepartmentsAPIV1DepartmentsIdDeleteRequest struct {
	ctx context.Context
	ApiService DepartmentsAPI
	id string
}

func (r DepartmentsAPIV1DepartmentsIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1DepartmentsIdDeleteExecute(r)
}

/*
V1DepartmentsIdDelete Remove specified department record 

Removes specified department record


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id instance id of department record
 @return DepartmentsAPIV1DepartmentsIdDeleteRequest
*/
func (a *DepartmentsAPIService) V1DepartmentsIdDelete(ctx context.Context, id string) DepartmentsAPIV1DepartmentsIdDeleteRequest {
	return DepartmentsAPIV1DepartmentsIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *DepartmentsAPIService) V1DepartmentsIdDeleteExecute(r DepartmentsAPIV1DepartmentsIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DepartmentsAPIService.V1DepartmentsIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/departments/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

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
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type DepartmentsAPIV1DepartmentsIdGetRequest struct {
	ctx context.Context
	ApiService DepartmentsAPI
	id string
}

func (r DepartmentsAPIV1DepartmentsIdGetRequest) Execute() (*Department, *http.Response, error) {
	return r.ApiService.V1DepartmentsIdGetExecute(r)
}

/*
V1DepartmentsIdGet Get specified Department object 

Gets specified Department object


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id instance id of department record
 @return DepartmentsAPIV1DepartmentsIdGetRequest
*/
func (a *DepartmentsAPIService) V1DepartmentsIdGet(ctx context.Context, id string) DepartmentsAPIV1DepartmentsIdGetRequest {
	return DepartmentsAPIV1DepartmentsIdGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return Department
func (a *DepartmentsAPIService) V1DepartmentsIdGetExecute(r DepartmentsAPIV1DepartmentsIdGetRequest) (*Department, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Department
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DepartmentsAPIService.V1DepartmentsIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/departments/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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
		if localVarHTTPResponse.StatusCode == 404 {
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

type DepartmentsAPIV1DepartmentsIdHistoryGetRequest struct {
	ctx context.Context
	ApiService DepartmentsAPI
	id string
	page *int64
	pageSize *int64
	sort *[]string
	filter *string
}

func (r DepartmentsAPIV1DepartmentsIdHistoryGetRequest) Page(page int64) DepartmentsAPIV1DepartmentsIdHistoryGetRequest {
	r.page = &page
	return r
}

func (r DepartmentsAPIV1DepartmentsIdHistoryGetRequest) PageSize(pageSize int64) DepartmentsAPIV1DepartmentsIdHistoryGetRequest {
	r.pageSize = &pageSize
	return r
}

// Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;date:desc,name:asc 
func (r DepartmentsAPIV1DepartmentsIdHistoryGetRequest) Sort(sort []string) DepartmentsAPIV1DepartmentsIdHistoryGetRequest {
	r.sort = &sort
	return r
}

// Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: filter&#x3D;username!&#x3D;admin and details&#x3D;&#x3D;*disabled* and date&lt;2019-12-15
func (r DepartmentsAPIV1DepartmentsIdHistoryGetRequest) Filter(filter string) DepartmentsAPIV1DepartmentsIdHistoryGetRequest {
	r.filter = &filter
	return r
}

func (r DepartmentsAPIV1DepartmentsIdHistoryGetRequest) Execute() (*HistorySearchResults, *http.Response, error) {
	return r.ApiService.V1DepartmentsIdHistoryGetExecute(r)
}

/*
V1DepartmentsIdHistoryGet Get specified Department history object 

Gets specified Department history object


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id instance id of department history record
 @return DepartmentsAPIV1DepartmentsIdHistoryGetRequest
*/
func (a *DepartmentsAPIService) V1DepartmentsIdHistoryGet(ctx context.Context, id string) DepartmentsAPIV1DepartmentsIdHistoryGetRequest {
	return DepartmentsAPIV1DepartmentsIdHistoryGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return HistorySearchResults
func (a *DepartmentsAPIService) V1DepartmentsIdHistoryGetExecute(r DepartmentsAPIV1DepartmentsIdHistoryGetRequest) (*HistorySearchResults, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *HistorySearchResults
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DepartmentsAPIService.V1DepartmentsIdHistoryGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/departments/{id}/history"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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
		if localVarHTTPResponse.StatusCode == 404 {
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

type DepartmentsAPIV1DepartmentsIdHistoryPostRequest struct {
	ctx context.Context
	ApiService DepartmentsAPI
	id string
	objectHistoryNote *ObjectHistoryNote
}

// history notes to create
func (r DepartmentsAPIV1DepartmentsIdHistoryPostRequest) ObjectHistoryNote(objectHistoryNote ObjectHistoryNote) DepartmentsAPIV1DepartmentsIdHistoryPostRequest {
	r.objectHistoryNote = &objectHistoryNote
	return r
}

func (r DepartmentsAPIV1DepartmentsIdHistoryPostRequest) Execute() (*HrefResponse, *http.Response, error) {
	return r.ApiService.V1DepartmentsIdHistoryPostExecute(r)
}

/*
V1DepartmentsIdHistoryPost Add specified Department history object notes 

Adds specified Department history object notes


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id instance id of department history record
 @return DepartmentsAPIV1DepartmentsIdHistoryPostRequest
*/
func (a *DepartmentsAPIService) V1DepartmentsIdHistoryPost(ctx context.Context, id string) DepartmentsAPIV1DepartmentsIdHistoryPostRequest {
	return DepartmentsAPIV1DepartmentsIdHistoryPostRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return HrefResponse
func (a *DepartmentsAPIService) V1DepartmentsIdHistoryPostExecute(r DepartmentsAPIV1DepartmentsIdHistoryPostRequest) (*HrefResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *HrefResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DepartmentsAPIService.V1DepartmentsIdHistoryPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/departments/{id}/history"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
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

type DepartmentsAPIV1DepartmentsIdPutRequest struct {
	ctx context.Context
	ApiService DepartmentsAPI
	id string
	department *Department
}

// department object to create. ids defined in this body will be ignored
func (r DepartmentsAPIV1DepartmentsIdPutRequest) Department(department Department) DepartmentsAPIV1DepartmentsIdPutRequest {
	r.department = &department
	return r
}

func (r DepartmentsAPIV1DepartmentsIdPutRequest) Execute() (*Department, *http.Response, error) {
	return r.ApiService.V1DepartmentsIdPutExecute(r)
}

/*
V1DepartmentsIdPut Update specified department object 

Update specified department object


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id instance id of department record
 @return DepartmentsAPIV1DepartmentsIdPutRequest
*/
func (a *DepartmentsAPIService) V1DepartmentsIdPut(ctx context.Context, id string) DepartmentsAPIV1DepartmentsIdPutRequest {
	return DepartmentsAPIV1DepartmentsIdPutRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return Department
func (a *DepartmentsAPIService) V1DepartmentsIdPutExecute(r DepartmentsAPIV1DepartmentsIdPutRequest) (*Department, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Department
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DepartmentsAPIService.V1DepartmentsIdPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/departments/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.department == nil {
		return localVarReturnValue, nil, reportError("department is required and must be specified")
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
	localVarPostBody = r.department
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

type DepartmentsAPIV1DepartmentsPostRequest struct {
	ctx context.Context
	ApiService DepartmentsAPI
	department *Department
}

// department object to create. ids defined in this body will be ignored
func (r DepartmentsAPIV1DepartmentsPostRequest) Department(department Department) DepartmentsAPIV1DepartmentsPostRequest {
	r.department = &department
	return r
}

func (r DepartmentsAPIV1DepartmentsPostRequest) Execute() (*HrefResponse, *http.Response, error) {
	return r.ApiService.V1DepartmentsPostExecute(r)
}

/*
V1DepartmentsPost Create department record 

Create department record


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return DepartmentsAPIV1DepartmentsPostRequest
*/
func (a *DepartmentsAPIService) V1DepartmentsPost(ctx context.Context) DepartmentsAPIV1DepartmentsPostRequest {
	return DepartmentsAPIV1DepartmentsPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return HrefResponse
func (a *DepartmentsAPIService) V1DepartmentsPostExecute(r DepartmentsAPIV1DepartmentsPostRequest) (*HrefResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *HrefResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DepartmentsAPIService.V1DepartmentsPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/departments"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.department == nil {
		return localVarReturnValue, nil, reportError("department is required and must be specified")
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
	localVarPostBody = r.department
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
