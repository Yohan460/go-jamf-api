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


type PatchManagementAPI interface {

	/*
	V2PatchManagementAcceptDisclaimerPost Accept Patch Management disclaimer 

	Accept Patch Management disclaimer

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return PatchManagementAPIV2PatchManagementAcceptDisclaimerPostRequest
	*/
	V2PatchManagementAcceptDisclaimerPost(ctx context.Context) PatchManagementAPIV2PatchManagementAcceptDisclaimerPostRequest

	// V2PatchManagementAcceptDisclaimerPostExecute executes the request
	V2PatchManagementAcceptDisclaimerPostExecute(r PatchManagementAPIV2PatchManagementAcceptDisclaimerPostRequest) (*http.Response, error)
}

// PatchManagementAPIService PatchManagementAPI service
type PatchManagementAPIService service

type PatchManagementAPIV2PatchManagementAcceptDisclaimerPostRequest struct {
	ctx context.Context
	ApiService PatchManagementAPI
}

func (r PatchManagementAPIV2PatchManagementAcceptDisclaimerPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.V2PatchManagementAcceptDisclaimerPostExecute(r)
}

/*
V2PatchManagementAcceptDisclaimerPost Accept Patch Management disclaimer 

Accept Patch Management disclaimer

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return PatchManagementAPIV2PatchManagementAcceptDisclaimerPostRequest
*/
func (a *PatchManagementAPIService) V2PatchManagementAcceptDisclaimerPost(ctx context.Context) PatchManagementAPIV2PatchManagementAcceptDisclaimerPostRequest {
	return PatchManagementAPIV2PatchManagementAcceptDisclaimerPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *PatchManagementAPIService) V2PatchManagementAcceptDisclaimerPostExecute(r PatchManagementAPIV2PatchManagementAcceptDisclaimerPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PatchManagementAPIService.V2PatchManagementAcceptDisclaimerPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/patch-management-accept-disclaimer"

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
