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
	"strings"
	"reflect"
)


type BuildingsApi interface {

	/*
	V1BuildingsDeleteMultiplePost Delete multiple Buildings by their ids 

	multiple many Buildings by their ids

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiV1BuildingsDeleteMultiplePostRequest
	*/
	V1BuildingsDeleteMultiplePost(ctx context.Context) ApiV1BuildingsDeleteMultiplePostRequest

	// V1BuildingsDeleteMultiplePostExecute executes the request
	V1BuildingsDeleteMultiplePostExecute(r ApiV1BuildingsDeleteMultiplePostRequest) (*http.Response, error)

	/*
	V1BuildingsGet Search for sorted and paged Buildings 

	Search for sorted and paged buildings

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiV1BuildingsGetRequest
	*/
	V1BuildingsGet(ctx context.Context) ApiV1BuildingsGetRequest

	// V1BuildingsGetExecute executes the request
	//  @return BuildingSearchResults
	V1BuildingsGetExecute(r ApiV1BuildingsGetRequest) (*BuildingSearchResults, *http.Response, error)

	/*
	V1BuildingsIdDelete Remove specified Building record 

	Removes specified building record


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id instance id of building record
	@return ApiV1BuildingsIdDeleteRequest
	*/
	V1BuildingsIdDelete(ctx context.Context, id string) ApiV1BuildingsIdDeleteRequest

	// V1BuildingsIdDeleteExecute executes the request
	V1BuildingsIdDeleteExecute(r ApiV1BuildingsIdDeleteRequest) (*http.Response, error)

	/*
	V1BuildingsIdGet Get specified Building object 

	Gets specified Building object


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id instance id of building record
	@return ApiV1BuildingsIdGetRequest
	*/
	V1BuildingsIdGet(ctx context.Context, id string) ApiV1BuildingsIdGetRequest

	// V1BuildingsIdGetExecute executes the request
	//  @return Building
	V1BuildingsIdGetExecute(r ApiV1BuildingsIdGetRequest) (*Building, *http.Response, error)

	/*
	V1BuildingsIdHistoryGet Get specified Building History object 

	Gets specified Building history object


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id instance id of building history record
	@return ApiV1BuildingsIdHistoryGetRequest
	*/
	V1BuildingsIdHistoryGet(ctx context.Context, id string) ApiV1BuildingsIdHistoryGetRequest

	// V1BuildingsIdHistoryGetExecute executes the request
	//  @return HistorySearchResults
	V1BuildingsIdHistoryGetExecute(r ApiV1BuildingsIdHistoryGetRequest) (*HistorySearchResults, *http.Response, error)

	/*
	V1BuildingsIdHistoryPost Add specified Building history object notes 

	Adds specified Building history object notes


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id instance id of building history record
	@return ApiV1BuildingsIdHistoryPostRequest
	*/
	V1BuildingsIdHistoryPost(ctx context.Context, id string) ApiV1BuildingsIdHistoryPostRequest

	// V1BuildingsIdHistoryPostExecute executes the request
	//  @return ObjectHistory
	V1BuildingsIdHistoryPostExecute(r ApiV1BuildingsIdHistoryPostRequest) (*ObjectHistory, *http.Response, error)

	/*
	V1BuildingsIdPut Update specified Building object 

	Update specified building object


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id instance id of building record
	@return ApiV1BuildingsIdPutRequest
	*/
	V1BuildingsIdPut(ctx context.Context, id string) ApiV1BuildingsIdPutRequest

	// V1BuildingsIdPutExecute executes the request
	//  @return Building
	V1BuildingsIdPutExecute(r ApiV1BuildingsIdPutRequest) (*Building, *http.Response, error)

	/*
	V1BuildingsPost Create Building record 

	Create building record


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiV1BuildingsPostRequest
	*/
	V1BuildingsPost(ctx context.Context) ApiV1BuildingsPostRequest

	// V1BuildingsPostExecute executes the request
	//  @return HrefResponse
	V1BuildingsPostExecute(r ApiV1BuildingsPostRequest) (*HrefResponse, *http.Response, error)
}

// BuildingsApiService BuildingsApi service
type BuildingsApiService service

type ApiV1BuildingsDeleteMultiplePostRequest struct {
	ctx context.Context
	ApiService BuildingsApi
	ids *Ids
}

// ids of the building to be deleted
func (r ApiV1BuildingsDeleteMultiplePostRequest) Ids(ids Ids) ApiV1BuildingsDeleteMultiplePostRequest {
	r.ids = &ids
	return r
}

func (r ApiV1BuildingsDeleteMultiplePostRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1BuildingsDeleteMultiplePostExecute(r)
}

/*
V1BuildingsDeleteMultiplePost Delete multiple Buildings by their ids 

multiple many Buildings by their ids

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV1BuildingsDeleteMultiplePostRequest
*/
func (a *BuildingsApiService) V1BuildingsDeleteMultiplePost(ctx context.Context) ApiV1BuildingsDeleteMultiplePostRequest {
	return ApiV1BuildingsDeleteMultiplePostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *BuildingsApiService) V1BuildingsDeleteMultiplePostExecute(r ApiV1BuildingsDeleteMultiplePostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BuildingsApiService.V1BuildingsDeleteMultiplePost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/buildings/delete-multiple"

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiV1BuildingsGetRequest struct {
	ctx context.Context
	ApiService BuildingsApi
	page *int32
	pageSize *int32
	sort *[]string
	filter *string
}

func (r ApiV1BuildingsGetRequest) Page(page int32) ApiV1BuildingsGetRequest {
	r.page = &page
	return r
}

func (r ApiV1BuildingsGetRequest) PageSize(pageSize int32) ApiV1BuildingsGetRequest {
	r.pageSize = &pageSize
	return r
}

// Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;date:desc,name:asc 
func (r ApiV1BuildingsGetRequest) Sort(sort []string) ApiV1BuildingsGetRequest {
	r.sort = &sort
	return r
}

// Query in the RSQL format, allowing to filter buildings collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: name, streetAddress1, streetAddress2, city, stateProvince, zipPostalCode, country. This param can be combined with paging and sorting. Example: filter&#x3D;city&#x3D;&#x3D;\&quot;Chicago\&quot; and name&#x3D;&#x3D;\&quot;*build*\&quot;
func (r ApiV1BuildingsGetRequest) Filter(filter string) ApiV1BuildingsGetRequest {
	r.filter = &filter
	return r
}

func (r ApiV1BuildingsGetRequest) Execute() (*BuildingSearchResults, *http.Response, error) {
	return r.ApiService.V1BuildingsGetExecute(r)
}

/*
V1BuildingsGet Search for sorted and paged Buildings 

Search for sorted and paged buildings

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV1BuildingsGetRequest
*/
func (a *BuildingsApiService) V1BuildingsGet(ctx context.Context) ApiV1BuildingsGetRequest {
	return ApiV1BuildingsGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BuildingSearchResults
func (a *BuildingsApiService) V1BuildingsGetExecute(r ApiV1BuildingsGetRequest) (*BuildingSearchResults, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BuildingSearchResults
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BuildingsApiService.V1BuildingsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/buildings"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("page-size", parameterToString(*r.pageSize, ""))
	}
	if r.sort != nil {
		t := *r.sort
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("sort", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("sort", parameterToString(t, "multi"))
		}
	}
	if r.filter != nil {
		localVarQueryParams.Add("filter", parameterToString(*r.filter, ""))
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

type ApiV1BuildingsIdDeleteRequest struct {
	ctx context.Context
	ApiService BuildingsApi
	id string
}

func (r ApiV1BuildingsIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1BuildingsIdDeleteExecute(r)
}

/*
V1BuildingsIdDelete Remove specified Building record 

Removes specified building record


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id instance id of building record
 @return ApiV1BuildingsIdDeleteRequest
*/
func (a *BuildingsApiService) V1BuildingsIdDelete(ctx context.Context, id string) ApiV1BuildingsIdDeleteRequest {
	return ApiV1BuildingsIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *BuildingsApiService) V1BuildingsIdDeleteExecute(r ApiV1BuildingsIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BuildingsApiService.V1BuildingsIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/buildings/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiV1BuildingsIdGetRequest struct {
	ctx context.Context
	ApiService BuildingsApi
	id string
}

func (r ApiV1BuildingsIdGetRequest) Execute() (*Building, *http.Response, error) {
	return r.ApiService.V1BuildingsIdGetExecute(r)
}

/*
V1BuildingsIdGet Get specified Building object 

Gets specified Building object


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id instance id of building record
 @return ApiV1BuildingsIdGetRequest
*/
func (a *BuildingsApiService) V1BuildingsIdGet(ctx context.Context, id string) ApiV1BuildingsIdGetRequest {
	return ApiV1BuildingsIdGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return Building
func (a *BuildingsApiService) V1BuildingsIdGetExecute(r ApiV1BuildingsIdGetRequest) (*Building, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Building
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BuildingsApiService.V1BuildingsIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/buildings/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

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
		if localVarHTTPResponse.StatusCode == 404 {
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

type ApiV1BuildingsIdHistoryGetRequest struct {
	ctx context.Context
	ApiService BuildingsApi
	id string
	page *int32
	pageSize *int32
	sort *[]string
	filter *string
}

func (r ApiV1BuildingsIdHistoryGetRequest) Page(page int32) ApiV1BuildingsIdHistoryGetRequest {
	r.page = &page
	return r
}

func (r ApiV1BuildingsIdHistoryGetRequest) PageSize(pageSize int32) ApiV1BuildingsIdHistoryGetRequest {
	r.pageSize = &pageSize
	return r
}

// Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;date:desc,name:asc 
func (r ApiV1BuildingsIdHistoryGetRequest) Sort(sort []string) ApiV1BuildingsIdHistoryGetRequest {
	r.sort = &sort
	return r
}

// Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: filter&#x3D;username!&#x3D;admin and details&#x3D;&#x3D;*disabled* and date&lt;2019-12-15
func (r ApiV1BuildingsIdHistoryGetRequest) Filter(filter string) ApiV1BuildingsIdHistoryGetRequest {
	r.filter = &filter
	return r
}

func (r ApiV1BuildingsIdHistoryGetRequest) Execute() (*HistorySearchResults, *http.Response, error) {
	return r.ApiService.V1BuildingsIdHistoryGetExecute(r)
}

/*
V1BuildingsIdHistoryGet Get specified Building History object 

Gets specified Building history object


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id instance id of building history record
 @return ApiV1BuildingsIdHistoryGetRequest
*/
func (a *BuildingsApiService) V1BuildingsIdHistoryGet(ctx context.Context, id string) ApiV1BuildingsIdHistoryGetRequest {
	return ApiV1BuildingsIdHistoryGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return HistorySearchResults
func (a *BuildingsApiService) V1BuildingsIdHistoryGetExecute(r ApiV1BuildingsIdHistoryGetRequest) (*HistorySearchResults, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *HistorySearchResults
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BuildingsApiService.V1BuildingsIdHistoryGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/buildings/{id}/history"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("page-size", parameterToString(*r.pageSize, ""))
	}
	if r.sort != nil {
		t := *r.sort
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("sort", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("sort", parameterToString(t, "multi"))
		}
	}
	if r.filter != nil {
		localVarQueryParams.Add("filter", parameterToString(*r.filter, ""))
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
		if localVarHTTPResponse.StatusCode == 404 {
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

type ApiV1BuildingsIdHistoryPostRequest struct {
	ctx context.Context
	ApiService BuildingsApi
	id string
	objectHistoryNote *ObjectHistoryNote
}

// history notes to create
func (r ApiV1BuildingsIdHistoryPostRequest) ObjectHistoryNote(objectHistoryNote ObjectHistoryNote) ApiV1BuildingsIdHistoryPostRequest {
	r.objectHistoryNote = &objectHistoryNote
	return r
}

func (r ApiV1BuildingsIdHistoryPostRequest) Execute() (*ObjectHistory, *http.Response, error) {
	return r.ApiService.V1BuildingsIdHistoryPostExecute(r)
}

/*
V1BuildingsIdHistoryPost Add specified Building history object notes 

Adds specified Building history object notes


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id instance id of building history record
 @return ApiV1BuildingsIdHistoryPostRequest
*/
func (a *BuildingsApiService) V1BuildingsIdHistoryPost(ctx context.Context, id string) ApiV1BuildingsIdHistoryPostRequest {
	return ApiV1BuildingsIdHistoryPostRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ObjectHistory
func (a *BuildingsApiService) V1BuildingsIdHistoryPostExecute(r ApiV1BuildingsIdHistoryPostRequest) (*ObjectHistory, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ObjectHistory
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BuildingsApiService.V1BuildingsIdHistoryPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/buildings/{id}/history"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

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
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
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

type ApiV1BuildingsIdPutRequest struct {
	ctx context.Context
	ApiService BuildingsApi
	id string
	building *Building
}

// building object to update. ids defined in this body will be ignored
func (r ApiV1BuildingsIdPutRequest) Building(building Building) ApiV1BuildingsIdPutRequest {
	r.building = &building
	return r
}

func (r ApiV1BuildingsIdPutRequest) Execute() (*Building, *http.Response, error) {
	return r.ApiService.V1BuildingsIdPutExecute(r)
}

/*
V1BuildingsIdPut Update specified Building object 

Update specified building object


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id instance id of building record
 @return ApiV1BuildingsIdPutRequest
*/
func (a *BuildingsApiService) V1BuildingsIdPut(ctx context.Context, id string) ApiV1BuildingsIdPutRequest {
	return ApiV1BuildingsIdPutRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return Building
func (a *BuildingsApiService) V1BuildingsIdPutExecute(r ApiV1BuildingsIdPutRequest) (*Building, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Building
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BuildingsApiService.V1BuildingsIdPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/buildings/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.building == nil {
		return localVarReturnValue, nil, reportError("building is required and must be specified")
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
	localVarPostBody = r.building
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

type ApiV1BuildingsPostRequest struct {
	ctx context.Context
	ApiService BuildingsApi
	building *Building
}

// building object to create. ids defined in this body will be ignored
func (r ApiV1BuildingsPostRequest) Building(building Building) ApiV1BuildingsPostRequest {
	r.building = &building
	return r
}

func (r ApiV1BuildingsPostRequest) Execute() (*HrefResponse, *http.Response, error) {
	return r.ApiService.V1BuildingsPostExecute(r)
}

/*
V1BuildingsPost Create Building record 

Create building record


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV1BuildingsPostRequest
*/
func (a *BuildingsApiService) V1BuildingsPost(ctx context.Context) ApiV1BuildingsPostRequest {
	return ApiV1BuildingsPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return HrefResponse
func (a *BuildingsApiService) V1BuildingsPostExecute(r ApiV1BuildingsPostRequest) (*HrefResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *HrefResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BuildingsApiService.V1BuildingsPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/buildings"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.building == nil {
		return localVarReturnValue, nil, reportError("building is required and must be specified")
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
	localVarPostBody = r.building
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
