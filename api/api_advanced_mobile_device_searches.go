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


type AdvancedMobileDeviceSearchesAPI interface {

	/*
	V1AdvancedMobileDeviceSearchesChoicesGet Get Mobile Device Advanced Search criteria choices 

	Gets Mobile Device Advanced Search criteria choices. A list of potentially valid choices can be found by navigating to the Criteria page of the Advanced Mobile Device Search creation process. A few are "App Name", "Building", and "Display Name".


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesChoicesGetRequest
	*/
	V1AdvancedMobileDeviceSearchesChoicesGet(ctx context.Context) AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesChoicesGetRequest

	// V1AdvancedMobileDeviceSearchesChoicesGetExecute executes the request
	//  @return AdvancedSearchCriteriaChoices
	V1AdvancedMobileDeviceSearchesChoicesGetExecute(r AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesChoicesGetRequest) (*AdvancedSearchCriteriaChoices, *http.Response, error)

	/*
	V1AdvancedMobileDeviceSearchesDeleteMultiplePost Remove specified Advanced Search objects 

	Removes specified Advanced Search Objects


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesDeleteMultiplePostRequest
	*/
	V1AdvancedMobileDeviceSearchesDeleteMultiplePost(ctx context.Context) AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesDeleteMultiplePostRequest

	// V1AdvancedMobileDeviceSearchesDeleteMultiplePostExecute executes the request
	V1AdvancedMobileDeviceSearchesDeleteMultiplePostExecute(r AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesDeleteMultiplePostRequest) (*http.Response, error)

	/*
	V1AdvancedMobileDeviceSearchesGet Get Advanced Search objects 

	Gets Advanced Search Objects


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesGetRequest
	*/
	V1AdvancedMobileDeviceSearchesGet(ctx context.Context) AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesGetRequest

	// V1AdvancedMobileDeviceSearchesGetExecute executes the request
	//  @return AdvancedSearchSearchResults
	V1AdvancedMobileDeviceSearchesGetExecute(r AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesGetRequest) (*AdvancedSearchSearchResults, *http.Response, error)

	/*
	V1AdvancedMobileDeviceSearchesIdDelete Remove specified Advanced Search object 

	Removes specified Advanced Search Object


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id instance id of advanced search record
	@return AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesIdDeleteRequest
	*/
	V1AdvancedMobileDeviceSearchesIdDelete(ctx context.Context, id string) AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesIdDeleteRequest

	// V1AdvancedMobileDeviceSearchesIdDeleteExecute executes the request
	V1AdvancedMobileDeviceSearchesIdDeleteExecute(r AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesIdDeleteRequest) (*http.Response, error)

	/*
	V1AdvancedMobileDeviceSearchesIdGet Get specified Advanced Search object 

	Gets Specified Advanced Search Object


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id id of target Advanced Search
	@return AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesIdGetRequest
	*/
	V1AdvancedMobileDeviceSearchesIdGet(ctx context.Context, id string) AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesIdGetRequest

	// V1AdvancedMobileDeviceSearchesIdGetExecute executes the request
	//  @return AdvancedSearch
	V1AdvancedMobileDeviceSearchesIdGetExecute(r AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesIdGetRequest) (*AdvancedSearch, *http.Response, error)

	/*
	V1AdvancedMobileDeviceSearchesIdPut Get specified Advanced Search object 

	Gets Specified Advanced Search Object


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id id of target Advanced Search
	@return AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesIdPutRequest
	*/
	V1AdvancedMobileDeviceSearchesIdPut(ctx context.Context, id string) AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesIdPutRequest

	// V1AdvancedMobileDeviceSearchesIdPutExecute executes the request
	//  @return AdvancedSearch
	V1AdvancedMobileDeviceSearchesIdPutExecute(r AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesIdPutRequest) (*AdvancedSearch, *http.Response, error)

	/*
	V1AdvancedMobileDeviceSearchesPost Create Advanced Search object 

	Creates Advanced Search Object


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesPostRequest
	*/
	V1AdvancedMobileDeviceSearchesPost(ctx context.Context) AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesPostRequest

	// V1AdvancedMobileDeviceSearchesPostExecute executes the request
	//  @return HrefResponse
	V1AdvancedMobileDeviceSearchesPostExecute(r AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesPostRequest) (*HrefResponse, *http.Response, error)
}

// AdvancedMobileDeviceSearchesAPIService AdvancedMobileDeviceSearchesAPI service
type AdvancedMobileDeviceSearchesAPIService service

type AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesChoicesGetRequest struct {
	ctx context.Context
	ApiService AdvancedMobileDeviceSearchesAPI
	criteria *string
	site *string
	contains *string
}

func (r AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesChoicesGetRequest) Criteria(criteria string) AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesChoicesGetRequest {
	r.criteria = &criteria
	return r
}

func (r AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesChoicesGetRequest) Site(site string) AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesChoicesGetRequest {
	r.site = &site
	return r
}

func (r AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesChoicesGetRequest) Contains(contains string) AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesChoicesGetRequest {
	r.contains = &contains
	return r
}

func (r AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesChoicesGetRequest) Execute() (*AdvancedSearchCriteriaChoices, *http.Response, error) {
	return r.ApiService.V1AdvancedMobileDeviceSearchesChoicesGetExecute(r)
}

/*
V1AdvancedMobileDeviceSearchesChoicesGet Get Mobile Device Advanced Search criteria choices 

Gets Mobile Device Advanced Search criteria choices. A list of potentially valid choices can be found by navigating to the Criteria page of the Advanced Mobile Device Search creation process. A few are "App Name", "Building", and "Display Name".


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesChoicesGetRequest
*/
func (a *AdvancedMobileDeviceSearchesAPIService) V1AdvancedMobileDeviceSearchesChoicesGet(ctx context.Context) AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesChoicesGetRequest {
	return AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesChoicesGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AdvancedSearchCriteriaChoices
func (a *AdvancedMobileDeviceSearchesAPIService) V1AdvancedMobileDeviceSearchesChoicesGetExecute(r AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesChoicesGetRequest) (*AdvancedSearchCriteriaChoices, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AdvancedSearchCriteriaChoices
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdvancedMobileDeviceSearchesAPIService.V1AdvancedMobileDeviceSearchesChoicesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/advanced-mobile-device-searches/choices"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.criteria == nil {
		return localVarReturnValue, nil, reportError("criteria is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "criteria", r.criteria, "")
	if r.site != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "site", r.site, "")
	} else {
		var defaultValue string = "-1"
		r.site = &defaultValue
	}
	if r.contains != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "contains", r.contains, "")
	} else {
		var defaultValue string = "null"
		r.contains = &defaultValue
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

type AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesDeleteMultiplePostRequest struct {
	ctx context.Context
	ApiService AdvancedMobileDeviceSearchesAPI
	ids *Ids
}

// ids of the building to be deleted
func (r AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesDeleteMultiplePostRequest) Ids(ids Ids) AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesDeleteMultiplePostRequest {
	r.ids = &ids
	return r
}

func (r AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesDeleteMultiplePostRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1AdvancedMobileDeviceSearchesDeleteMultiplePostExecute(r)
}

/*
V1AdvancedMobileDeviceSearchesDeleteMultiplePost Remove specified Advanced Search objects 

Removes specified Advanced Search Objects


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesDeleteMultiplePostRequest
*/
func (a *AdvancedMobileDeviceSearchesAPIService) V1AdvancedMobileDeviceSearchesDeleteMultiplePost(ctx context.Context) AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesDeleteMultiplePostRequest {
	return AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesDeleteMultiplePostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *AdvancedMobileDeviceSearchesAPIService) V1AdvancedMobileDeviceSearchesDeleteMultiplePostExecute(r AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesDeleteMultiplePostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdvancedMobileDeviceSearchesAPIService.V1AdvancedMobileDeviceSearchesDeleteMultiplePost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/advanced-mobile-device-searches/delete-multiple"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.ids == nil {
		return nil, reportError("ids is required and must be specified")
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
	localVarPostBody = r.ids
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
		if localVarHTTPResponse.StatusCode == 400 {
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

type AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesGetRequest struct {
	ctx context.Context
	ApiService AdvancedMobileDeviceSearchesAPI
}

func (r AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesGetRequest) Execute() (*AdvancedSearchSearchResults, *http.Response, error) {
	return r.ApiService.V1AdvancedMobileDeviceSearchesGetExecute(r)
}

/*
V1AdvancedMobileDeviceSearchesGet Get Advanced Search objects 

Gets Advanced Search Objects


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesGetRequest
*/
func (a *AdvancedMobileDeviceSearchesAPIService) V1AdvancedMobileDeviceSearchesGet(ctx context.Context) AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesGetRequest {
	return AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AdvancedSearchSearchResults
func (a *AdvancedMobileDeviceSearchesAPIService) V1AdvancedMobileDeviceSearchesGetExecute(r AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesGetRequest) (*AdvancedSearchSearchResults, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AdvancedSearchSearchResults
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdvancedMobileDeviceSearchesAPIService.V1AdvancedMobileDeviceSearchesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/advanced-mobile-device-searches"

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

type AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesIdDeleteRequest struct {
	ctx context.Context
	ApiService AdvancedMobileDeviceSearchesAPI
	id string
}

func (r AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1AdvancedMobileDeviceSearchesIdDeleteExecute(r)
}

/*
V1AdvancedMobileDeviceSearchesIdDelete Remove specified Advanced Search object 

Removes specified Advanced Search Object


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id instance id of advanced search record
 @return AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesIdDeleteRequest
*/
func (a *AdvancedMobileDeviceSearchesAPIService) V1AdvancedMobileDeviceSearchesIdDelete(ctx context.Context, id string) AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesIdDeleteRequest {
	return AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *AdvancedMobileDeviceSearchesAPIService) V1AdvancedMobileDeviceSearchesIdDeleteExecute(r AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdvancedMobileDeviceSearchesAPIService.V1AdvancedMobileDeviceSearchesIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/advanced-mobile-device-searches/{id}"
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

type AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesIdGetRequest struct {
	ctx context.Context
	ApiService AdvancedMobileDeviceSearchesAPI
	id string
}

func (r AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesIdGetRequest) Execute() (*AdvancedSearch, *http.Response, error) {
	return r.ApiService.V1AdvancedMobileDeviceSearchesIdGetExecute(r)
}

/*
V1AdvancedMobileDeviceSearchesIdGet Get specified Advanced Search object 

Gets Specified Advanced Search Object


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id id of target Advanced Search
 @return AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesIdGetRequest
*/
func (a *AdvancedMobileDeviceSearchesAPIService) V1AdvancedMobileDeviceSearchesIdGet(ctx context.Context, id string) AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesIdGetRequest {
	return AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesIdGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return AdvancedSearch
func (a *AdvancedMobileDeviceSearchesAPIService) V1AdvancedMobileDeviceSearchesIdGetExecute(r AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesIdGetRequest) (*AdvancedSearch, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AdvancedSearch
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdvancedMobileDeviceSearchesAPIService.V1AdvancedMobileDeviceSearchesIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/advanced-mobile-device-searches/{id}"
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

type AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesIdPutRequest struct {
	ctx context.Context
	ApiService AdvancedMobileDeviceSearchesAPI
	id string
	advancedSearch *AdvancedSearch
}

func (r AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesIdPutRequest) AdvancedSearch(advancedSearch AdvancedSearch) AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesIdPutRequest {
	r.advancedSearch = &advancedSearch
	return r
}

func (r AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesIdPutRequest) Execute() (*AdvancedSearch, *http.Response, error) {
	return r.ApiService.V1AdvancedMobileDeviceSearchesIdPutExecute(r)
}

/*
V1AdvancedMobileDeviceSearchesIdPut Get specified Advanced Search object 

Gets Specified Advanced Search Object


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id id of target Advanced Search
 @return AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesIdPutRequest
*/
func (a *AdvancedMobileDeviceSearchesAPIService) V1AdvancedMobileDeviceSearchesIdPut(ctx context.Context, id string) AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesIdPutRequest {
	return AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesIdPutRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return AdvancedSearch
func (a *AdvancedMobileDeviceSearchesAPIService) V1AdvancedMobileDeviceSearchesIdPutExecute(r AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesIdPutRequest) (*AdvancedSearch, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AdvancedSearch
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdvancedMobileDeviceSearchesAPIService.V1AdvancedMobileDeviceSearchesIdPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/advanced-mobile-device-searches/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advancedSearch == nil {
		return localVarReturnValue, nil, reportError("advancedSearch is required and must be specified")
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
	localVarPostBody = r.advancedSearch
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

type AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesPostRequest struct {
	ctx context.Context
	ApiService AdvancedMobileDeviceSearchesAPI
	advancedSearch *AdvancedSearch
}

func (r AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesPostRequest) AdvancedSearch(advancedSearch AdvancedSearch) AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesPostRequest {
	r.advancedSearch = &advancedSearch
	return r
}

func (r AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesPostRequest) Execute() (*HrefResponse, *http.Response, error) {
	return r.ApiService.V1AdvancedMobileDeviceSearchesPostExecute(r)
}

/*
V1AdvancedMobileDeviceSearchesPost Create Advanced Search object 

Creates Advanced Search Object


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesPostRequest
*/
func (a *AdvancedMobileDeviceSearchesAPIService) V1AdvancedMobileDeviceSearchesPost(ctx context.Context) AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesPostRequest {
	return AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return HrefResponse
func (a *AdvancedMobileDeviceSearchesAPIService) V1AdvancedMobileDeviceSearchesPostExecute(r AdvancedMobileDeviceSearchesAPIV1AdvancedMobileDeviceSearchesPostRequest) (*HrefResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *HrefResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdvancedMobileDeviceSearchesAPIService.V1AdvancedMobileDeviceSearchesPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/advanced-mobile-device-searches"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advancedSearch == nil {
		return localVarReturnValue, nil, reportError("advancedSearch is required and must be specified")
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
	localVarPostBody = r.advancedSearch
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
