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
)


type StaticUserGroupsPreviewAPI interface {

	/*
	V1StaticUserGroupsGet Return a list of all Static User Groups 

	Returns a list of all static user groups.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return StaticUserGroupsPreviewAPIV1StaticUserGroupsGetRequest
	*/
	V1StaticUserGroupsGet(ctx context.Context) StaticUserGroupsPreviewAPIV1StaticUserGroupsGetRequest

	// V1StaticUserGroupsGetExecute executes the request
	//  @return []StaticUserGroup
	V1StaticUserGroupsGetExecute(r StaticUserGroupsPreviewAPIV1StaticUserGroupsGetRequest) ([]StaticUserGroup, *http.Response, error)

	/*
	V1StaticUserGroupsIdGet Return a specific Static User Group by id 

	Returns a specific static user group by id.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Instance id of static user group record
	@return StaticUserGroupsPreviewAPIV1StaticUserGroupsIdGetRequest
	*/
	V1StaticUserGroupsIdGet(ctx context.Context, id int64) StaticUserGroupsPreviewAPIV1StaticUserGroupsIdGetRequest

	// V1StaticUserGroupsIdGetExecute executes the request
	//  @return StaticUserGroup
	V1StaticUserGroupsIdGetExecute(r StaticUserGroupsPreviewAPIV1StaticUserGroupsIdGetRequest) (*StaticUserGroup, *http.Response, error)
}

// StaticUserGroupsPreviewAPIService StaticUserGroupsPreviewAPI service
type StaticUserGroupsPreviewAPIService service

type StaticUserGroupsPreviewAPIV1StaticUserGroupsGetRequest struct {
	ctx context.Context
	ApiService StaticUserGroupsPreviewAPI
}

func (r StaticUserGroupsPreviewAPIV1StaticUserGroupsGetRequest) Execute() ([]StaticUserGroup, *http.Response, error) {
	return r.ApiService.V1StaticUserGroupsGetExecute(r)
}

/*
V1StaticUserGroupsGet Return a list of all Static User Groups 

Returns a list of all static user groups.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return StaticUserGroupsPreviewAPIV1StaticUserGroupsGetRequest
*/
func (a *StaticUserGroupsPreviewAPIService) V1StaticUserGroupsGet(ctx context.Context) StaticUserGroupsPreviewAPIV1StaticUserGroupsGetRequest {
	return StaticUserGroupsPreviewAPIV1StaticUserGroupsGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []StaticUserGroup
func (a *StaticUserGroupsPreviewAPIService) V1StaticUserGroupsGetExecute(r StaticUserGroupsPreviewAPIV1StaticUserGroupsGetRequest) ([]StaticUserGroup, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []StaticUserGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StaticUserGroupsPreviewAPIService.V1StaticUserGroupsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/static-user-groups"

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

type StaticUserGroupsPreviewAPIV1StaticUserGroupsIdGetRequest struct {
	ctx context.Context
	ApiService StaticUserGroupsPreviewAPI
	id int64
}

func (r StaticUserGroupsPreviewAPIV1StaticUserGroupsIdGetRequest) Execute() (*StaticUserGroup, *http.Response, error) {
	return r.ApiService.V1StaticUserGroupsIdGetExecute(r)
}

/*
V1StaticUserGroupsIdGet Return a specific Static User Group by id 

Returns a specific static user group by id.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Instance id of static user group record
 @return StaticUserGroupsPreviewAPIV1StaticUserGroupsIdGetRequest
*/
func (a *StaticUserGroupsPreviewAPIService) V1StaticUserGroupsIdGet(ctx context.Context, id int64) StaticUserGroupsPreviewAPIV1StaticUserGroupsIdGetRequest {
	return StaticUserGroupsPreviewAPIV1StaticUserGroupsIdGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return StaticUserGroup
func (a *StaticUserGroupsPreviewAPIService) V1StaticUserGroupsIdGetExecute(r StaticUserGroupsPreviewAPIV1StaticUserGroupsIdGetRequest) (*StaticUserGroup, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *StaticUserGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StaticUserGroupsPreviewAPIService.V1StaticUserGroupsIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/static-user-groups/{id}"
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
