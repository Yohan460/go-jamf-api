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


type HealthCheckAPI interface {

	/*
	V1HealthCheckGet Get Jamf Pro API status

	Get Jamf Pro API status. Which response codes might be returned in error states will depend on the specific state encountered.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return HealthCheckAPIV1HealthCheckGetRequest
	*/
	V1HealthCheckGet(ctx context.Context) HealthCheckAPIV1HealthCheckGetRequest

	// V1HealthCheckGetExecute executes the request
	V1HealthCheckGetExecute(r HealthCheckAPIV1HealthCheckGetRequest) (*http.Response, error)
}

// HealthCheckAPIService HealthCheckAPI service
type HealthCheckAPIService service

type HealthCheckAPIV1HealthCheckGetRequest struct {
	ctx context.Context
	ApiService HealthCheckAPI
}

func (r HealthCheckAPIV1HealthCheckGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1HealthCheckGetExecute(r)
}

/*
V1HealthCheckGet Get Jamf Pro API status

Get Jamf Pro API status. Which response codes might be returned in error states will depend on the specific state encountered.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return HealthCheckAPIV1HealthCheckGetRequest
*/
func (a *HealthCheckAPIService) V1HealthCheckGet(ctx context.Context) HealthCheckAPIV1HealthCheckGetRequest {
	return HealthCheckAPIV1HealthCheckGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *HealthCheckAPIService) V1HealthCheckGetExecute(r HealthCheckAPIV1HealthCheckGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HealthCheckAPIService.V1HealthCheckGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/health-check"

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
