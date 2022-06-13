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
)


type SelfServiceApi interface {

	/*
	V1SelfServiceSettingsGet Get an object representation of Self Service settings 

	gets an object representation of Self Service settings


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiV1SelfServiceSettingsGetRequest
	*/
	V1SelfServiceSettingsGet(ctx context.Context) ApiV1SelfServiceSettingsGetRequest

	// V1SelfServiceSettingsGetExecute executes the request
	//  @return SelfServiceSettings
	V1SelfServiceSettingsGetExecute(r ApiV1SelfServiceSettingsGetRequest) (*SelfServiceSettings, *http.Response, error)

	/*
	V1SelfServiceSettingsPut Put an object representation of Self Service settings 

	puts an object representation of Self Service settings


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiV1SelfServiceSettingsPutRequest
	*/
	V1SelfServiceSettingsPut(ctx context.Context) ApiV1SelfServiceSettingsPutRequest

	// V1SelfServiceSettingsPutExecute executes the request
	//  @return SelfServiceSettings
	V1SelfServiceSettingsPutExecute(r ApiV1SelfServiceSettingsPutRequest) (*SelfServiceSettings, *http.Response, error)
}

// SelfServiceApiService SelfServiceApi service
type SelfServiceApiService service

type ApiV1SelfServiceSettingsGetRequest struct {
	ctx context.Context
	ApiService SelfServiceApi
}

func (r ApiV1SelfServiceSettingsGetRequest) Execute() (*SelfServiceSettings, *http.Response, error) {
	return r.ApiService.V1SelfServiceSettingsGetExecute(r)
}

/*
V1SelfServiceSettingsGet Get an object representation of Self Service settings 

gets an object representation of Self Service settings


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV1SelfServiceSettingsGetRequest
*/
func (a *SelfServiceApiService) V1SelfServiceSettingsGet(ctx context.Context) ApiV1SelfServiceSettingsGetRequest {
	return ApiV1SelfServiceSettingsGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SelfServiceSettings
func (a *SelfServiceApiService) V1SelfServiceSettingsGetExecute(r ApiV1SelfServiceSettingsGetRequest) (*SelfServiceSettings, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SelfServiceSettings
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SelfServiceApiService.V1SelfServiceSettingsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/self-service/settings"

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

type ApiV1SelfServiceSettingsPutRequest struct {
	ctx context.Context
	ApiService SelfServiceApi
	selfServiceSettings *SelfServiceSettings
}

// object that contains all editable global fields to alter Self Service settings 
func (r ApiV1SelfServiceSettingsPutRequest) SelfServiceSettings(selfServiceSettings SelfServiceSettings) ApiV1SelfServiceSettingsPutRequest {
	r.selfServiceSettings = &selfServiceSettings
	return r
}

func (r ApiV1SelfServiceSettingsPutRequest) Execute() (*SelfServiceSettings, *http.Response, error) {
	return r.ApiService.V1SelfServiceSettingsPutExecute(r)
}

/*
V1SelfServiceSettingsPut Put an object representation of Self Service settings 

puts an object representation of Self Service settings


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV1SelfServiceSettingsPutRequest
*/
func (a *SelfServiceApiService) V1SelfServiceSettingsPut(ctx context.Context) ApiV1SelfServiceSettingsPutRequest {
	return ApiV1SelfServiceSettingsPutRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SelfServiceSettings
func (a *SelfServiceApiService) V1SelfServiceSettingsPutExecute(r ApiV1SelfServiceSettingsPutRequest) (*SelfServiceSettings, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SelfServiceSettings
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SelfServiceApiService.V1SelfServiceSettingsPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/self-service/settings"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.selfServiceSettings == nil {
		return localVarReturnValue, nil, reportError("selfServiceSettings is required and must be specified")
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
	localVarPostBody = r.selfServiceSettings
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
