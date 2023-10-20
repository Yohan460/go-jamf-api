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


type CategoriesAPI interface {

	/*
	V1CategoriesDeleteMultiplePost Delete multiple Categories by their IDs 

	Delete multiple Categories by their IDs

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return CategoriesAPIV1CategoriesDeleteMultiplePostRequest
	*/
	V1CategoriesDeleteMultiplePost(ctx context.Context) CategoriesAPIV1CategoriesDeleteMultiplePostRequest

	// V1CategoriesDeleteMultiplePostExecute executes the request
	V1CategoriesDeleteMultiplePostExecute(r CategoriesAPIV1CategoriesDeleteMultiplePostRequest) (*http.Response, error)

	/*
	V1CategoriesGet Get Category objects 

	Gets `Category` objects.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return CategoriesAPIV1CategoriesGetRequest
	*/
	V1CategoriesGet(ctx context.Context) CategoriesAPIV1CategoriesGetRequest

	// V1CategoriesGetExecute executes the request
	//  @return CategoriesSearchResults
	V1CategoriesGetExecute(r CategoriesAPIV1CategoriesGetRequest) (*CategoriesSearchResults, *http.Response, error)

	/*
	V1CategoriesIdDelete Remove specified Category record 

	Removes specified category record


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id instance id of category record
	@return CategoriesAPIV1CategoriesIdDeleteRequest
	*/
	V1CategoriesIdDelete(ctx context.Context, id string) CategoriesAPIV1CategoriesIdDeleteRequest

	// V1CategoriesIdDeleteExecute executes the request
	V1CategoriesIdDeleteExecute(r CategoriesAPIV1CategoriesIdDeleteRequest) (*http.Response, error)

	/*
	V1CategoriesIdGet Get specified Category object 

	Gets specified Category object


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id instance id of category record
	@return CategoriesAPIV1CategoriesIdGetRequest
	*/
	V1CategoriesIdGet(ctx context.Context, id string) CategoriesAPIV1CategoriesIdGetRequest

	// V1CategoriesIdGetExecute executes the request
	//  @return Category
	V1CategoriesIdGetExecute(r CategoriesAPIV1CategoriesIdGetRequest) (*Category, *http.Response, error)

	/*
	V1CategoriesIdHistoryGet Get specified Category history object 

	Gets specified Category history object


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id instance id of category history record
	@return CategoriesAPIV1CategoriesIdHistoryGetRequest
	*/
	V1CategoriesIdHistoryGet(ctx context.Context, id string) CategoriesAPIV1CategoriesIdHistoryGetRequest

	// V1CategoriesIdHistoryGetExecute executes the request
	//  @return HistorySearchResults
	V1CategoriesIdHistoryGetExecute(r CategoriesAPIV1CategoriesIdHistoryGetRequest) (*HistorySearchResults, *http.Response, error)

	/*
	V1CategoriesIdHistoryPost Add specified Category history object notes 

	Adds specified Category history object notes


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id instance id of category history record
	@return CategoriesAPIV1CategoriesIdHistoryPostRequest
	*/
	V1CategoriesIdHistoryPost(ctx context.Context, id string) CategoriesAPIV1CategoriesIdHistoryPostRequest

	// V1CategoriesIdHistoryPostExecute executes the request
	//  @return ObjectHistory
	V1CategoriesIdHistoryPostExecute(r CategoriesAPIV1CategoriesIdHistoryPostRequest) (*ObjectHistory, *http.Response, error)

	/*
	V1CategoriesIdPut Update specified Category object 

	Update specified category object


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id instance id of category record
	@return CategoriesAPIV1CategoriesIdPutRequest
	*/
	V1CategoriesIdPut(ctx context.Context, id string) CategoriesAPIV1CategoriesIdPutRequest

	// V1CategoriesIdPutExecute executes the request
	//  @return Category
	V1CategoriesIdPutExecute(r CategoriesAPIV1CategoriesIdPutRequest) (*Category, *http.Response, error)

	/*
	V1CategoriesPost Create Category record 

	Create category record


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return CategoriesAPIV1CategoriesPostRequest
	*/
	V1CategoriesPost(ctx context.Context) CategoriesAPIV1CategoriesPostRequest

	// V1CategoriesPostExecute executes the request
	//  @return HrefResponse
	V1CategoriesPostExecute(r CategoriesAPIV1CategoriesPostRequest) (*HrefResponse, *http.Response, error)
}

// CategoriesAPIService CategoriesAPI service
type CategoriesAPIService service

type CategoriesAPIV1CategoriesDeleteMultiplePostRequest struct {
	ctx context.Context
	ApiService CategoriesAPI
	ids *Ids
}

// IDs of the categories to be deleted
func (r CategoriesAPIV1CategoriesDeleteMultiplePostRequest) Ids(ids Ids) CategoriesAPIV1CategoriesDeleteMultiplePostRequest {
	r.ids = &ids
	return r
}

func (r CategoriesAPIV1CategoriesDeleteMultiplePostRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1CategoriesDeleteMultiplePostExecute(r)
}

/*
V1CategoriesDeleteMultiplePost Delete multiple Categories by their IDs 

Delete multiple Categories by their IDs

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return CategoriesAPIV1CategoriesDeleteMultiplePostRequest
*/
func (a *CategoriesAPIService) V1CategoriesDeleteMultiplePost(ctx context.Context) CategoriesAPIV1CategoriesDeleteMultiplePostRequest {
	return CategoriesAPIV1CategoriesDeleteMultiplePostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *CategoriesAPIService) V1CategoriesDeleteMultiplePostExecute(r CategoriesAPIV1CategoriesDeleteMultiplePostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CategoriesAPIService.V1CategoriesDeleteMultiplePost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/categories/delete-multiple"

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

type CategoriesAPIV1CategoriesGetRequest struct {
	ctx context.Context
	ApiService CategoriesAPI
	page *int32
	pageSize *int32
	sort *[]string
	filter *string
}

func (r CategoriesAPIV1CategoriesGetRequest) Page(page int32) CategoriesAPIV1CategoriesGetRequest {
	r.page = &page
	return r
}

func (r CategoriesAPIV1CategoriesGetRequest) PageSize(pageSize int32) CategoriesAPIV1CategoriesGetRequest {
	r.pageSize = &pageSize
	return r
}

// Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;date:desc,name:asc 
func (r CategoriesAPIV1CategoriesGetRequest) Sort(sort []string) CategoriesAPIV1CategoriesGetRequest {
	r.sort = &sort
	return r
}

// Query in the RSQL format, allowing to filter categories collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: name, priority. This param can be combined with paging and sorting. Example: filter&#x3D;name&#x3D;&#x3D;\&quot;Apps*\&quot; and priority&gt;&#x3D;5
func (r CategoriesAPIV1CategoriesGetRequest) Filter(filter string) CategoriesAPIV1CategoriesGetRequest {
	r.filter = &filter
	return r
}

func (r CategoriesAPIV1CategoriesGetRequest) Execute() (*CategoriesSearchResults, *http.Response, error) {
	return r.ApiService.V1CategoriesGetExecute(r)
}

/*
V1CategoriesGet Get Category objects 

Gets `Category` objects.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return CategoriesAPIV1CategoriesGetRequest
*/
func (a *CategoriesAPIService) V1CategoriesGet(ctx context.Context) CategoriesAPIV1CategoriesGetRequest {
	return CategoriesAPIV1CategoriesGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CategoriesSearchResults
func (a *CategoriesAPIService) V1CategoriesGetExecute(r CategoriesAPIV1CategoriesGetRequest) (*CategoriesSearchResults, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CategoriesSearchResults
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CategoriesAPIService.V1CategoriesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/categories"

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

type CategoriesAPIV1CategoriesIdDeleteRequest struct {
	ctx context.Context
	ApiService CategoriesAPI
	id string
}

func (r CategoriesAPIV1CategoriesIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1CategoriesIdDeleteExecute(r)
}

/*
V1CategoriesIdDelete Remove specified Category record 

Removes specified category record


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id instance id of category record
 @return CategoriesAPIV1CategoriesIdDeleteRequest
*/
func (a *CategoriesAPIService) V1CategoriesIdDelete(ctx context.Context, id string) CategoriesAPIV1CategoriesIdDeleteRequest {
	return CategoriesAPIV1CategoriesIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *CategoriesAPIService) V1CategoriesIdDeleteExecute(r CategoriesAPIV1CategoriesIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CategoriesAPIService.V1CategoriesIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/categories/{id}"
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

type CategoriesAPIV1CategoriesIdGetRequest struct {
	ctx context.Context
	ApiService CategoriesAPI
	id string
}

func (r CategoriesAPIV1CategoriesIdGetRequest) Execute() (*Category, *http.Response, error) {
	return r.ApiService.V1CategoriesIdGetExecute(r)
}

/*
V1CategoriesIdGet Get specified Category object 

Gets specified Category object


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id instance id of category record
 @return CategoriesAPIV1CategoriesIdGetRequest
*/
func (a *CategoriesAPIService) V1CategoriesIdGet(ctx context.Context, id string) CategoriesAPIV1CategoriesIdGetRequest {
	return CategoriesAPIV1CategoriesIdGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return Category
func (a *CategoriesAPIService) V1CategoriesIdGetExecute(r CategoriesAPIV1CategoriesIdGetRequest) (*Category, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Category
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CategoriesAPIService.V1CategoriesIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/categories/{id}"
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

type CategoriesAPIV1CategoriesIdHistoryGetRequest struct {
	ctx context.Context
	ApiService CategoriesAPI
	id string
	page *int32
	pageSize *int32
	sort *[]string
	filter *string
}

func (r CategoriesAPIV1CategoriesIdHistoryGetRequest) Page(page int32) CategoriesAPIV1CategoriesIdHistoryGetRequest {
	r.page = &page
	return r
}

func (r CategoriesAPIV1CategoriesIdHistoryGetRequest) PageSize(pageSize int32) CategoriesAPIV1CategoriesIdHistoryGetRequest {
	r.pageSize = &pageSize
	return r
}

// Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;date:desc,name:asc 
func (r CategoriesAPIV1CategoriesIdHistoryGetRequest) Sort(sort []string) CategoriesAPIV1CategoriesIdHistoryGetRequest {
	r.sort = &sort
	return r
}

// Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: filter&#x3D;username!&#x3D;admin and details&#x3D;&#x3D;*disabled* and date&lt;2019-12-15
func (r CategoriesAPIV1CategoriesIdHistoryGetRequest) Filter(filter string) CategoriesAPIV1CategoriesIdHistoryGetRequest {
	r.filter = &filter
	return r
}

func (r CategoriesAPIV1CategoriesIdHistoryGetRequest) Execute() (*HistorySearchResults, *http.Response, error) {
	return r.ApiService.V1CategoriesIdHistoryGetExecute(r)
}

/*
V1CategoriesIdHistoryGet Get specified Category history object 

Gets specified Category history object


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id instance id of category history record
 @return CategoriesAPIV1CategoriesIdHistoryGetRequest
*/
func (a *CategoriesAPIService) V1CategoriesIdHistoryGet(ctx context.Context, id string) CategoriesAPIV1CategoriesIdHistoryGetRequest {
	return CategoriesAPIV1CategoriesIdHistoryGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return HistorySearchResults
func (a *CategoriesAPIService) V1CategoriesIdHistoryGetExecute(r CategoriesAPIV1CategoriesIdHistoryGetRequest) (*HistorySearchResults, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *HistorySearchResults
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CategoriesAPIService.V1CategoriesIdHistoryGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/categories/{id}/history"
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

type CategoriesAPIV1CategoriesIdHistoryPostRequest struct {
	ctx context.Context
	ApiService CategoriesAPI
	id string
	objectHistoryNote *ObjectHistoryNote
}

// history notes to create
func (r CategoriesAPIV1CategoriesIdHistoryPostRequest) ObjectHistoryNote(objectHistoryNote ObjectHistoryNote) CategoriesAPIV1CategoriesIdHistoryPostRequest {
	r.objectHistoryNote = &objectHistoryNote
	return r
}

func (r CategoriesAPIV1CategoriesIdHistoryPostRequest) Execute() (*ObjectHistory, *http.Response, error) {
	return r.ApiService.V1CategoriesIdHistoryPostExecute(r)
}

/*
V1CategoriesIdHistoryPost Add specified Category history object notes 

Adds specified Category history object notes


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id instance id of category history record
 @return CategoriesAPIV1CategoriesIdHistoryPostRequest
*/
func (a *CategoriesAPIService) V1CategoriesIdHistoryPost(ctx context.Context, id string) CategoriesAPIV1CategoriesIdHistoryPostRequest {
	return CategoriesAPIV1CategoriesIdHistoryPostRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ObjectHistory
func (a *CategoriesAPIService) V1CategoriesIdHistoryPostExecute(r CategoriesAPIV1CategoriesIdHistoryPostRequest) (*ObjectHistory, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ObjectHistory
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CategoriesAPIService.V1CategoriesIdHistoryPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/categories/{id}/history"
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

type CategoriesAPIV1CategoriesIdPutRequest struct {
	ctx context.Context
	ApiService CategoriesAPI
	id string
	category *Category
}

// category object to create. id defined in this body will be ignored
func (r CategoriesAPIV1CategoriesIdPutRequest) Category(category Category) CategoriesAPIV1CategoriesIdPutRequest {
	r.category = &category
	return r
}

func (r CategoriesAPIV1CategoriesIdPutRequest) Execute() (*Category, *http.Response, error) {
	return r.ApiService.V1CategoriesIdPutExecute(r)
}

/*
V1CategoriesIdPut Update specified Category object 

Update specified category object


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id instance id of category record
 @return CategoriesAPIV1CategoriesIdPutRequest
*/
func (a *CategoriesAPIService) V1CategoriesIdPut(ctx context.Context, id string) CategoriesAPIV1CategoriesIdPutRequest {
	return CategoriesAPIV1CategoriesIdPutRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return Category
func (a *CategoriesAPIService) V1CategoriesIdPutExecute(r CategoriesAPIV1CategoriesIdPutRequest) (*Category, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Category
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CategoriesAPIService.V1CategoriesIdPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/categories/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.category == nil {
		return localVarReturnValue, nil, reportError("category is required and must be specified")
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
	localVarPostBody = r.category
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

type CategoriesAPIV1CategoriesPostRequest struct {
	ctx context.Context
	ApiService CategoriesAPI
	category *Category
}

// category object to create. IDs defined in this body will be ignored
func (r CategoriesAPIV1CategoriesPostRequest) Category(category Category) CategoriesAPIV1CategoriesPostRequest {
	r.category = &category
	return r
}

func (r CategoriesAPIV1CategoriesPostRequest) Execute() (*HrefResponse, *http.Response, error) {
	return r.ApiService.V1CategoriesPostExecute(r)
}

/*
V1CategoriesPost Create Category record 

Create category record


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return CategoriesAPIV1CategoriesPostRequest
*/
func (a *CategoriesAPIService) V1CategoriesPost(ctx context.Context) CategoriesAPIV1CategoriesPostRequest {
	return CategoriesAPIV1CategoriesPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return HrefResponse
func (a *CategoriesAPIService) V1CategoriesPostExecute(r CategoriesAPIV1CategoriesPostRequest) (*HrefResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *HrefResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CategoriesAPIService.V1CategoriesPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/categories"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.category == nil {
		return localVarReturnValue, nil, reportError("category is required and must be specified")
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
	localVarPostBody = r.category
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
