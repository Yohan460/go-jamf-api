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


type CloudDistributionPointAPI interface {

	/*
	V1CloudDistributionPointUploadCapabilityGet Finds specific information for the currently configured Cloud Distribution Point. 

	Finds a variety of values based on the currently configured Cloud Distribution Point.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return CloudDistributionPointAPIV1CloudDistributionPointUploadCapabilityGetRequest
	*/
	V1CloudDistributionPointUploadCapabilityGet(ctx context.Context) CloudDistributionPointAPIV1CloudDistributionPointUploadCapabilityGetRequest

	// V1CloudDistributionPointUploadCapabilityGetExecute executes the request
	//  @return CloudDistributionPointUploadCapability
	V1CloudDistributionPointUploadCapabilityGetExecute(r CloudDistributionPointAPIV1CloudDistributionPointUploadCapabilityGetRequest) (*CloudDistributionPointUploadCapability, *http.Response, error)
}

// CloudDistributionPointAPIService CloudDistributionPointAPI service
type CloudDistributionPointAPIService service

type CloudDistributionPointAPIV1CloudDistributionPointUploadCapabilityGetRequest struct {
	ctx context.Context
	ApiService CloudDistributionPointAPI
}

func (r CloudDistributionPointAPIV1CloudDistributionPointUploadCapabilityGetRequest) Execute() (*CloudDistributionPointUploadCapability, *http.Response, error) {
	return r.ApiService.V1CloudDistributionPointUploadCapabilityGetExecute(r)
}

/*
V1CloudDistributionPointUploadCapabilityGet Finds specific information for the currently configured Cloud Distribution Point. 

Finds a variety of values based on the currently configured Cloud Distribution Point.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return CloudDistributionPointAPIV1CloudDistributionPointUploadCapabilityGetRequest
*/
func (a *CloudDistributionPointAPIService) V1CloudDistributionPointUploadCapabilityGet(ctx context.Context) CloudDistributionPointAPIV1CloudDistributionPointUploadCapabilityGetRequest {
	return CloudDistributionPointAPIV1CloudDistributionPointUploadCapabilityGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CloudDistributionPointUploadCapability
func (a *CloudDistributionPointAPIService) V1CloudDistributionPointUploadCapabilityGetExecute(r CloudDistributionPointAPIV1CloudDistributionPointUploadCapabilityGetRequest) (*CloudDistributionPointUploadCapability, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CloudDistributionPointUploadCapability
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudDistributionPointAPIService.V1CloudDistributionPointUploadCapabilityGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/cloud-distribution-point/upload-capability"

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