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
	"os"
)


type BrandingAPI interface {

	/*
	V1BrandingImagesDownloadIdGet Download a self service branding image 

	Download a self service branding image

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id id of the self service branding image
	@return BrandingAPIV1BrandingImagesDownloadIdGetRequest
	*/
	V1BrandingImagesDownloadIdGet(ctx context.Context, id string) BrandingAPIV1BrandingImagesDownloadIdGetRequest

	// V1BrandingImagesDownloadIdGetExecute executes the request
	//  @return *os.File
	V1BrandingImagesDownloadIdGetExecute(r BrandingAPIV1BrandingImagesDownloadIdGetRequest) (*os.File, *http.Response, error)
}

// BrandingAPIService BrandingAPI service
type BrandingAPIService service

type BrandingAPIV1BrandingImagesDownloadIdGetRequest struct {
	ctx context.Context
	ApiService BrandingAPI
	id string
}

func (r BrandingAPIV1BrandingImagesDownloadIdGetRequest) Execute() (*os.File, *http.Response, error) {
	return r.ApiService.V1BrandingImagesDownloadIdGetExecute(r)
}

/*
V1BrandingImagesDownloadIdGet Download a self service branding image 

Download a self service branding image

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id id of the self service branding image
 @return BrandingAPIV1BrandingImagesDownloadIdGetRequest
*/
func (a *BrandingAPIService) V1BrandingImagesDownloadIdGet(ctx context.Context, id string) BrandingAPIV1BrandingImagesDownloadIdGetRequest {
	return BrandingAPIV1BrandingImagesDownloadIdGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return *os.File
func (a *BrandingAPIService) V1BrandingImagesDownloadIdGetExecute(r BrandingAPIV1BrandingImagesDownloadIdGetRequest) (*os.File, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BrandingAPIService.V1BrandingImagesDownloadIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/branding-images/download/{id}"
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
	localVarHTTPHeaderAccepts := []string{"image/*"}

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
