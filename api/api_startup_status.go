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
)


type StartupStatusAPI interface {

	/*
	StartupStatusGet Retrieve information about application startup 

	Retrieves information about application startup. Current startup operation taking place (if any) and overall startup completion percentage.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return StartupStatusAPIStartupStatusGetRequest
	*/
	StartupStatusGet(ctx context.Context) StartupStatusAPIStartupStatusGetRequest

	// StartupStatusGetExecute executes the request
	//  @return StartupStatus
	StartupStatusGetExecute(r StartupStatusAPIStartupStatusGetRequest) (*StartupStatus, *http.Response, error)
}

// StartupStatusAPIService StartupStatusAPI service
type StartupStatusAPIService service

type StartupStatusAPIStartupStatusGetRequest struct {
	ctx context.Context
	ApiService StartupStatusAPI
}

func (r StartupStatusAPIStartupStatusGetRequest) Execute() (*StartupStatus, *http.Response, error) {
	return r.ApiService.StartupStatusGetExecute(r)
}

/*
StartupStatusGet Retrieve information about application startup 

Retrieves information about application startup. Current startup operation taking place (if any) and overall startup completion percentage.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return StartupStatusAPIStartupStatusGetRequest
*/
func (a *StartupStatusAPIService) StartupStatusGet(ctx context.Context) StartupStatusAPIStartupStatusGetRequest {
	return StartupStatusAPIStartupStatusGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return StartupStatus
func (a *StartupStatusAPIService) StartupStatusGetExecute(r StartupStatusAPIStartupStatusGetRequest) (*StartupStatus, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *StartupStatus
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StartupStatusAPIService.StartupStatusGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/startup-status"

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
