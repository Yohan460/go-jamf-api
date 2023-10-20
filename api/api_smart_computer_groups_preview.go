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


type SmartComputerGroupsPreviewAPI interface {

	/*
	V1ComputersIdRecalculateSmartGroupsPost Recalculate a smart group for the given id 

	Recalculates a smart group for the given id and then
returns the count of smart groups the computer falls into


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id id of computer
	@return SmartComputerGroupsPreviewAPIV1ComputersIdRecalculateSmartGroupsPostRequest
	*/
	V1ComputersIdRecalculateSmartGroupsPost(ctx context.Context, id int32) SmartComputerGroupsPreviewAPIV1ComputersIdRecalculateSmartGroupsPostRequest

	// V1ComputersIdRecalculateSmartGroupsPostExecute executes the request
	//  @return RecalculationResults
	V1ComputersIdRecalculateSmartGroupsPostExecute(r SmartComputerGroupsPreviewAPIV1ComputersIdRecalculateSmartGroupsPostRequest) (*RecalculationResults, *http.Response, error)

	/*
	V1SmartComputerGroupsIdRecalculatePost Recalculate the smart group for the given id 

	Recalculates the smart group for the given id and then
returns the ids for the computers in the smart group


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id instance id of smart group
	@return SmartComputerGroupsPreviewAPIV1SmartComputerGroupsIdRecalculatePostRequest
	*/
	V1SmartComputerGroupsIdRecalculatePost(ctx context.Context, id int32) SmartComputerGroupsPreviewAPIV1SmartComputerGroupsIdRecalculatePostRequest

	// V1SmartComputerGroupsIdRecalculatePostExecute executes the request
	//  @return RecalculationResults
	V1SmartComputerGroupsIdRecalculatePostExecute(r SmartComputerGroupsPreviewAPIV1SmartComputerGroupsIdRecalculatePostRequest) (*RecalculationResults, *http.Response, error)
}

// SmartComputerGroupsPreviewAPIService SmartComputerGroupsPreviewAPI service
type SmartComputerGroupsPreviewAPIService service

type SmartComputerGroupsPreviewAPIV1ComputersIdRecalculateSmartGroupsPostRequest struct {
	ctx context.Context
	ApiService SmartComputerGroupsPreviewAPI
	id int32
}

func (r SmartComputerGroupsPreviewAPIV1ComputersIdRecalculateSmartGroupsPostRequest) Execute() (*RecalculationResults, *http.Response, error) {
	return r.ApiService.V1ComputersIdRecalculateSmartGroupsPostExecute(r)
}

/*
V1ComputersIdRecalculateSmartGroupsPost Recalculate a smart group for the given id 

Recalculates a smart group for the given id and then
returns the count of smart groups the computer falls into


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id id of computer
 @return SmartComputerGroupsPreviewAPIV1ComputersIdRecalculateSmartGroupsPostRequest
*/
func (a *SmartComputerGroupsPreviewAPIService) V1ComputersIdRecalculateSmartGroupsPost(ctx context.Context, id int32) SmartComputerGroupsPreviewAPIV1ComputersIdRecalculateSmartGroupsPostRequest {
	return SmartComputerGroupsPreviewAPIV1ComputersIdRecalculateSmartGroupsPostRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return RecalculationResults
func (a *SmartComputerGroupsPreviewAPIService) V1ComputersIdRecalculateSmartGroupsPostExecute(r SmartComputerGroupsPreviewAPIV1ComputersIdRecalculateSmartGroupsPostRequest) (*RecalculationResults, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RecalculationResults
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SmartComputerGroupsPreviewAPIService.V1ComputersIdRecalculateSmartGroupsPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/computers/{id}/recalculate-smart-groups"
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

type SmartComputerGroupsPreviewAPIV1SmartComputerGroupsIdRecalculatePostRequest struct {
	ctx context.Context
	ApiService SmartComputerGroupsPreviewAPI
	id int32
}

func (r SmartComputerGroupsPreviewAPIV1SmartComputerGroupsIdRecalculatePostRequest) Execute() (*RecalculationResults, *http.Response, error) {
	return r.ApiService.V1SmartComputerGroupsIdRecalculatePostExecute(r)
}

/*
V1SmartComputerGroupsIdRecalculatePost Recalculate the smart group for the given id 

Recalculates the smart group for the given id and then
returns the ids for the computers in the smart group


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id instance id of smart group
 @return SmartComputerGroupsPreviewAPIV1SmartComputerGroupsIdRecalculatePostRequest
*/
func (a *SmartComputerGroupsPreviewAPIService) V1SmartComputerGroupsIdRecalculatePost(ctx context.Context, id int32) SmartComputerGroupsPreviewAPIV1SmartComputerGroupsIdRecalculatePostRequest {
	return SmartComputerGroupsPreviewAPIV1SmartComputerGroupsIdRecalculatePostRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return RecalculationResults
func (a *SmartComputerGroupsPreviewAPIService) V1SmartComputerGroupsIdRecalculatePostExecute(r SmartComputerGroupsPreviewAPIV1SmartComputerGroupsIdRecalculatePostRequest) (*RecalculationResults, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RecalculationResults
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SmartComputerGroupsPreviewAPIService.V1SmartComputerGroupsIdRecalculatePost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/smart-computer-groups/{id}/recalculate"
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
