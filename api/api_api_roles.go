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


type ApiRolesAPI interface {

	/*
	DeleteApiRole Delete API Integrations Role

	Delete specific Role

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id instance id of API role
	@return ApiRolesAPIDeleteApiRoleRequest
	*/
	DeleteApiRole(ctx context.Context, id string) ApiRolesAPIDeleteApiRoleRequest

	// DeleteApiRoleExecute executes the request
	DeleteApiRoleExecute(r ApiRolesAPIDeleteApiRoleRequest) (*http.Response, error)

	/*
	GetAllApiRoles Get the current Jamf API Roles

	Get roles with Search Criteria

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiRolesAPIGetAllApiRolesRequest
	*/
	GetAllApiRoles(ctx context.Context) ApiRolesAPIGetAllApiRolesRequest

	// GetAllApiRolesExecute executes the request
	//  @return ApiRoleResult
	GetAllApiRolesExecute(r ApiRolesAPIGetAllApiRolesRequest) (*ApiRoleResult, *http.Response, error)

	/*
	GetOneApiRole Get the specific Jamf API Role

	Get specific Role

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id instance id of API role
	@return ApiRolesAPIGetOneApiRoleRequest
	*/
	GetOneApiRole(ctx context.Context, id string) ApiRolesAPIGetOneApiRoleRequest

	// GetOneApiRoleExecute executes the request
	//  @return ApiRole
	GetOneApiRoleExecute(r ApiRolesAPIGetOneApiRoleRequest) (*ApiRole, *http.Response, error)

	/*
	PostCreateApiRole Create a new API role

	Post to create new Role

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiRolesAPIPostCreateApiRoleRequest
	*/
	PostCreateApiRole(ctx context.Context) ApiRolesAPIPostCreateApiRoleRequest

	// PostCreateApiRoleExecute executes the request
	//  @return ApiRole
	PostCreateApiRoleExecute(r ApiRolesAPIPostCreateApiRoleRequest) (*ApiRole, *http.Response, error)

	/*
	PutUpdateApiRole Update API Integrations Role

	Update specific Role

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id instance id of API role
	@return ApiRolesAPIPutUpdateApiRoleRequest
	*/
	PutUpdateApiRole(ctx context.Context, id string) ApiRolesAPIPutUpdateApiRoleRequest

	// PutUpdateApiRoleExecute executes the request
	//  @return ApiRole
	PutUpdateApiRoleExecute(r ApiRolesAPIPutUpdateApiRoleRequest) (*ApiRole, *http.Response, error)
}

// ApiRolesAPIService ApiRolesAPI service
type ApiRolesAPIService service

type ApiRolesAPIDeleteApiRoleRequest struct {
	ctx context.Context
	ApiService ApiRolesAPI
	id string
}

func (r ApiRolesAPIDeleteApiRoleRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteApiRoleExecute(r)
}

/*
DeleteApiRole Delete API Integrations Role

Delete specific Role

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id instance id of API role
 @return ApiRolesAPIDeleteApiRoleRequest
*/
func (a *ApiRolesAPIService) DeleteApiRole(ctx context.Context, id string) ApiRolesAPIDeleteApiRoleRequest {
	return ApiRolesAPIDeleteApiRoleRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *ApiRolesAPIService) DeleteApiRoleExecute(r ApiRolesAPIDeleteApiRoleRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApiRolesAPIService.DeleteApiRole")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/api-roles/{id}"
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
		if localVarHTTPResponse.StatusCode == 409 {
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
		if localVarHTTPResponse.StatusCode == 404 {
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

type ApiRolesAPIGetAllApiRolesRequest struct {
	ctx context.Context
	ApiService ApiRolesAPI
	page *int64
	pageSize *int64
	sort *[]string
	filter *string
}

func (r ApiRolesAPIGetAllApiRolesRequest) Page(page int64) ApiRolesAPIGetAllApiRolesRequest {
	r.page = &page
	return r
}

func (r ApiRolesAPIGetAllApiRolesRequest) PageSize(pageSize int64) ApiRolesAPIGetAllApiRolesRequest {
	r.pageSize = &pageSize
	return r
}

// Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Fields allowed in the query: id, displayName. Example: sort&#x3D;displayName:desc
func (r ApiRolesAPIGetAllApiRolesRequest) Sort(sort []string) ApiRolesAPIGetAllApiRolesRequest {
	r.sort = &sort
	return r
}

// Query in the RSQL format, allowing to filter app titles collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, displayName. Example: displayName&#x3D;&#x3D;\&quot;*myRole*\&quot;
func (r ApiRolesAPIGetAllApiRolesRequest) Filter(filter string) ApiRolesAPIGetAllApiRolesRequest {
	r.filter = &filter
	return r
}

func (r ApiRolesAPIGetAllApiRolesRequest) Execute() (*ApiRoleResult, *http.Response, error) {
	return r.ApiService.GetAllApiRolesExecute(r)
}

/*
GetAllApiRoles Get the current Jamf API Roles

Get roles with Search Criteria

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRolesAPIGetAllApiRolesRequest
*/
func (a *ApiRolesAPIService) GetAllApiRoles(ctx context.Context) ApiRolesAPIGetAllApiRolesRequest {
	return ApiRolesAPIGetAllApiRolesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ApiRoleResult
func (a *ApiRolesAPIService) GetAllApiRolesExecute(r ApiRolesAPIGetAllApiRolesRequest) (*ApiRoleResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiRoleResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApiRolesAPIService.GetAllApiRoles")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/api-roles"

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

type ApiRolesAPIGetOneApiRoleRequest struct {
	ctx context.Context
	ApiService ApiRolesAPI
	id string
}

func (r ApiRolesAPIGetOneApiRoleRequest) Execute() (*ApiRole, *http.Response, error) {
	return r.ApiService.GetOneApiRoleExecute(r)
}

/*
GetOneApiRole Get the specific Jamf API Role

Get specific Role

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id instance id of API role
 @return ApiRolesAPIGetOneApiRoleRequest
*/
func (a *ApiRolesAPIService) GetOneApiRole(ctx context.Context, id string) ApiRolesAPIGetOneApiRoleRequest {
	return ApiRolesAPIGetOneApiRoleRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ApiRole
func (a *ApiRolesAPIService) GetOneApiRoleExecute(r ApiRolesAPIGetOneApiRoleRequest) (*ApiRole, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiRole
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApiRolesAPIService.GetOneApiRole")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/api-roles/{id}"
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

type ApiRolesAPIPostCreateApiRoleRequest struct {
	ctx context.Context
	ApiService ApiRolesAPI
	apiRoleRequest *ApiRoleRequest
}

// API Integrations Role to create
func (r ApiRolesAPIPostCreateApiRoleRequest) ApiRoleRequest(apiRoleRequest ApiRoleRequest) ApiRolesAPIPostCreateApiRoleRequest {
	r.apiRoleRequest = &apiRoleRequest
	return r
}

func (r ApiRolesAPIPostCreateApiRoleRequest) Execute() (*ApiRole, *http.Response, error) {
	return r.ApiService.PostCreateApiRoleExecute(r)
}

/*
PostCreateApiRole Create a new API role

Post to create new Role

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRolesAPIPostCreateApiRoleRequest
*/
func (a *ApiRolesAPIService) PostCreateApiRole(ctx context.Context) ApiRolesAPIPostCreateApiRoleRequest {
	return ApiRolesAPIPostCreateApiRoleRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ApiRole
func (a *ApiRolesAPIService) PostCreateApiRoleExecute(r ApiRolesAPIPostCreateApiRoleRequest) (*ApiRole, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiRole
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApiRolesAPIService.PostCreateApiRole")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/api-roles"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiRoleRequest == nil {
		return localVarReturnValue, nil, reportError("apiRoleRequest is required and must be specified")
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
	localVarPostBody = r.apiRoleRequest
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
		if localVarHTTPResponse.StatusCode == 400 {
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

type ApiRolesAPIPutUpdateApiRoleRequest struct {
	ctx context.Context
	ApiService ApiRolesAPI
	id string
	apiRoleRequest *ApiRoleRequest
}

// API Integrations Role to update
func (r ApiRolesAPIPutUpdateApiRoleRequest) ApiRoleRequest(apiRoleRequest ApiRoleRequest) ApiRolesAPIPutUpdateApiRoleRequest {
	r.apiRoleRequest = &apiRoleRequest
	return r
}

func (r ApiRolesAPIPutUpdateApiRoleRequest) Execute() (*ApiRole, *http.Response, error) {
	return r.ApiService.PutUpdateApiRoleExecute(r)
}

/*
PutUpdateApiRole Update API Integrations Role

Update specific Role

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id instance id of API role
 @return ApiRolesAPIPutUpdateApiRoleRequest
*/
func (a *ApiRolesAPIService) PutUpdateApiRole(ctx context.Context, id string) ApiRolesAPIPutUpdateApiRoleRequest {
	return ApiRolesAPIPutUpdateApiRoleRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ApiRole
func (a *ApiRolesAPIService) PutUpdateApiRoleExecute(r ApiRolesAPIPutUpdateApiRoleRequest) (*ApiRole, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiRole
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApiRolesAPIService.PutUpdateApiRole")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/api-roles/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiRoleRequest == nil {
		return localVarReturnValue, nil, reportError("apiRoleRequest is required and must be specified")
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
	localVarPostBody = r.apiRoleRequest
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
		if localVarHTTPResponse.StatusCode == 400 {
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
