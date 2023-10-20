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


type JamfManagementFrameworkAPI interface {

	/*
	V1JamfManagementFrameworkRedeployIdPost Redeploy Jamf Management Framework 

	Redeploys the Jamf Management Framework for enrolled device


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id instance id of computer
	@return JamfManagementFrameworkAPIV1JamfManagementFrameworkRedeployIdPostRequest
	*/
	V1JamfManagementFrameworkRedeployIdPost(ctx context.Context, id string) JamfManagementFrameworkAPIV1JamfManagementFrameworkRedeployIdPostRequest

	// V1JamfManagementFrameworkRedeployIdPostExecute executes the request
	//  @return RedeployJamfManagementFrameworkResponse
	V1JamfManagementFrameworkRedeployIdPostExecute(r JamfManagementFrameworkAPIV1JamfManagementFrameworkRedeployIdPostRequest) (*RedeployJamfManagementFrameworkResponse, *http.Response, error)
}

// JamfManagementFrameworkAPIService JamfManagementFrameworkAPI service
type JamfManagementFrameworkAPIService service

type JamfManagementFrameworkAPIV1JamfManagementFrameworkRedeployIdPostRequest struct {
	ctx context.Context
	ApiService JamfManagementFrameworkAPI
	id string
}

func (r JamfManagementFrameworkAPIV1JamfManagementFrameworkRedeployIdPostRequest) Execute() (*RedeployJamfManagementFrameworkResponse, *http.Response, error) {
	return r.ApiService.V1JamfManagementFrameworkRedeployIdPostExecute(r)
}

/*
V1JamfManagementFrameworkRedeployIdPost Redeploy Jamf Management Framework 

Redeploys the Jamf Management Framework for enrolled device


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id instance id of computer
 @return JamfManagementFrameworkAPIV1JamfManagementFrameworkRedeployIdPostRequest
*/
func (a *JamfManagementFrameworkAPIService) V1JamfManagementFrameworkRedeployIdPost(ctx context.Context, id string) JamfManagementFrameworkAPIV1JamfManagementFrameworkRedeployIdPostRequest {
	return JamfManagementFrameworkAPIV1JamfManagementFrameworkRedeployIdPostRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return RedeployJamfManagementFrameworkResponse
func (a *JamfManagementFrameworkAPIService) V1JamfManagementFrameworkRedeployIdPostExecute(r JamfManagementFrameworkAPIV1JamfManagementFrameworkRedeployIdPostRequest) (*RedeployJamfManagementFrameworkResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RedeployJamfManagementFrameworkResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JamfManagementFrameworkAPIService.V1JamfManagementFrameworkRedeployIdPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/jamf-management-framework/redeploy/{id}"
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
