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
	"reflect"
)


type PatchPoliciesAPI interface {

	/*
	V2PatchPoliciesGet Retrieve Patch Policies

	Retrieves a list of patch policies.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return PatchPoliciesAPIV2PatchPoliciesGetRequest
	*/
	V2PatchPoliciesGet(ctx context.Context) PatchPoliciesAPIV2PatchPoliciesGetRequest

	// V2PatchPoliciesGetExecute executes the request
	//  @return PatchPolicies
	V2PatchPoliciesGetExecute(r PatchPoliciesAPIV2PatchPoliciesGetRequest) (*PatchPolicies, *http.Response, error)

	/*
	V2PatchPoliciesIdDashboardDelete Remove a patch policy from the dashboard 

	Removes a patch policy from the dashboard.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id patch policy id
	@return PatchPoliciesAPIV2PatchPoliciesIdDashboardDeleteRequest
	*/
	V2PatchPoliciesIdDashboardDelete(ctx context.Context, id string) PatchPoliciesAPIV2PatchPoliciesIdDashboardDeleteRequest

	// V2PatchPoliciesIdDashboardDeleteExecute executes the request
	V2PatchPoliciesIdDashboardDeleteExecute(r PatchPoliciesAPIV2PatchPoliciesIdDashboardDeleteRequest) (*http.Response, error)

	/*
	V2PatchPoliciesIdDashboardGet Return whether or not the requested patch policy is on the dashboard 

	Returns whether or not the requested patch policy is on the dashboard

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id patch policy id
	@return PatchPoliciesAPIV2PatchPoliciesIdDashboardGetRequest
	*/
	V2PatchPoliciesIdDashboardGet(ctx context.Context, id string) PatchPoliciesAPIV2PatchPoliciesIdDashboardGetRequest

	// V2PatchPoliciesIdDashboardGetExecute executes the request
	//  @return PatchPolicyV2OnDashboard
	V2PatchPoliciesIdDashboardGetExecute(r PatchPoliciesAPIV2PatchPoliciesIdDashboardGetRequest) (*PatchPolicyV2OnDashboard, *http.Response, error)

	/*
	V2PatchPoliciesIdDashboardPost Add a patch policy to the dashboard 

	Adds a patch policy to the dashboard.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id patch policy id
	@return PatchPoliciesAPIV2PatchPoliciesIdDashboardPostRequest
	*/
	V2PatchPoliciesIdDashboardPost(ctx context.Context, id string) PatchPoliciesAPIV2PatchPoliciesIdDashboardPostRequest

	// V2PatchPoliciesIdDashboardPostExecute executes the request
	V2PatchPoliciesIdDashboardPostExecute(r PatchPoliciesAPIV2PatchPoliciesIdDashboardPostRequest) (*http.Response, error)

	/*
	V2PatchPoliciesPolicyDetailsGet Retrieve Patch Policies

	Retrieves a list of patch policies.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return PatchPoliciesAPIV2PatchPoliciesPolicyDetailsGetRequest
	*/
	V2PatchPoliciesPolicyDetailsGet(ctx context.Context) PatchPoliciesAPIV2PatchPoliciesPolicyDetailsGetRequest

	// V2PatchPoliciesPolicyDetailsGetExecute executes the request
	//  @return PatchPolicyDetails
	V2PatchPoliciesPolicyDetailsGetExecute(r PatchPoliciesAPIV2PatchPoliciesPolicyDetailsGetRequest) (*PatchPolicyDetails, *http.Response, error)
}

// PatchPoliciesAPIService PatchPoliciesAPI service
type PatchPoliciesAPIService service

type PatchPoliciesAPIV2PatchPoliciesGetRequest struct {
	ctx context.Context
	ApiService PatchPoliciesAPI
	page *int64
	pageSize *int64
	sort *[]string
	filter *string
}

func (r PatchPoliciesAPIV2PatchPoliciesGetRequest) Page(page int64) PatchPoliciesAPIV2PatchPoliciesGetRequest {
	r.page = &page
	return r
}

func (r PatchPoliciesAPIV2PatchPoliciesGetRequest) PageSize(pageSize int64) PatchPoliciesAPIV2PatchPoliciesGetRequest {
	r.pageSize = &pageSize
	return r
}

// Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma.
func (r PatchPoliciesAPIV2PatchPoliciesGetRequest) Sort(sort []string) PatchPoliciesAPIV2PatchPoliciesGetRequest {
	r.sort = &sort
	return r
}

// Query in the RSQL format, allowing to filter Patch Policy collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, policyName, policyEnabled, policyTargetVersion, policyDeploymentMethod, softwareTitle, softwareTitleConfigurationId, pending, completed, deferred, and failed. This param can be combined with paging and sorting.
func (r PatchPoliciesAPIV2PatchPoliciesGetRequest) Filter(filter string) PatchPoliciesAPIV2PatchPoliciesGetRequest {
	r.filter = &filter
	return r
}

func (r PatchPoliciesAPIV2PatchPoliciesGetRequest) Execute() (*PatchPolicies, *http.Response, error) {
	return r.ApiService.V2PatchPoliciesGetExecute(r)
}

/*
V2PatchPoliciesGet Retrieve Patch Policies

Retrieves a list of patch policies.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return PatchPoliciesAPIV2PatchPoliciesGetRequest
*/
func (a *PatchPoliciesAPIService) V2PatchPoliciesGet(ctx context.Context) PatchPoliciesAPIV2PatchPoliciesGetRequest {
	return PatchPoliciesAPIV2PatchPoliciesGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PatchPolicies
func (a *PatchPoliciesAPIService) V2PatchPoliciesGetExecute(r PatchPoliciesAPIV2PatchPoliciesGetRequest) (*PatchPolicies, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PatchPolicies
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PatchPoliciesAPIService.V2PatchPoliciesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/patch-policies"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	} else {
		var defaultValue int64 = 0
		r.page = &defaultValue
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page-size", r.pageSize, "")
	} else {
		var defaultValue int64 = 100
		r.pageSize = &defaultValue
	}
	if r.sort != nil {
		t := *r.sort
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "sort", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "sort", t, "multi")
		}
	} else {
		defaultValue := []string{"id:asc"}
		r.sort = &defaultValue
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "")
	} else {
		var defaultValue string = ""
		r.filter = &defaultValue
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

type PatchPoliciesAPIV2PatchPoliciesIdDashboardDeleteRequest struct {
	ctx context.Context
	ApiService PatchPoliciesAPI
	id string
}

func (r PatchPoliciesAPIV2PatchPoliciesIdDashboardDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.V2PatchPoliciesIdDashboardDeleteExecute(r)
}

/*
V2PatchPoliciesIdDashboardDelete Remove a patch policy from the dashboard 

Removes a patch policy from the dashboard.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id patch policy id
 @return PatchPoliciesAPIV2PatchPoliciesIdDashboardDeleteRequest
*/
func (a *PatchPoliciesAPIService) V2PatchPoliciesIdDashboardDelete(ctx context.Context, id string) PatchPoliciesAPIV2PatchPoliciesIdDashboardDeleteRequest {
	return PatchPoliciesAPIV2PatchPoliciesIdDashboardDeleteRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *PatchPoliciesAPIService) V2PatchPoliciesIdDashboardDeleteExecute(r PatchPoliciesAPIV2PatchPoliciesIdDashboardDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PatchPoliciesAPIService.V2PatchPoliciesIdDashboardDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/patch-policies/{id}/dashboard"
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

type PatchPoliciesAPIV2PatchPoliciesIdDashboardGetRequest struct {
	ctx context.Context
	ApiService PatchPoliciesAPI
	id string
}

func (r PatchPoliciesAPIV2PatchPoliciesIdDashboardGetRequest) Execute() (*PatchPolicyV2OnDashboard, *http.Response, error) {
	return r.ApiService.V2PatchPoliciesIdDashboardGetExecute(r)
}

/*
V2PatchPoliciesIdDashboardGet Return whether or not the requested patch policy is on the dashboard 

Returns whether or not the requested patch policy is on the dashboard

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id patch policy id
 @return PatchPoliciesAPIV2PatchPoliciesIdDashboardGetRequest
*/
func (a *PatchPoliciesAPIService) V2PatchPoliciesIdDashboardGet(ctx context.Context, id string) PatchPoliciesAPIV2PatchPoliciesIdDashboardGetRequest {
	return PatchPoliciesAPIV2PatchPoliciesIdDashboardGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return PatchPolicyV2OnDashboard
func (a *PatchPoliciesAPIService) V2PatchPoliciesIdDashboardGetExecute(r PatchPoliciesAPIV2PatchPoliciesIdDashboardGetRequest) (*PatchPolicyV2OnDashboard, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PatchPolicyV2OnDashboard
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PatchPoliciesAPIService.V2PatchPoliciesIdDashboardGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/patch-policies/{id}/dashboard"
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

type PatchPoliciesAPIV2PatchPoliciesIdDashboardPostRequest struct {
	ctx context.Context
	ApiService PatchPoliciesAPI
	id string
}

func (r PatchPoliciesAPIV2PatchPoliciesIdDashboardPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.V2PatchPoliciesIdDashboardPostExecute(r)
}

/*
V2PatchPoliciesIdDashboardPost Add a patch policy to the dashboard 

Adds a patch policy to the dashboard.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id patch policy id
 @return PatchPoliciesAPIV2PatchPoliciesIdDashboardPostRequest
*/
func (a *PatchPoliciesAPIService) V2PatchPoliciesIdDashboardPost(ctx context.Context, id string) PatchPoliciesAPIV2PatchPoliciesIdDashboardPostRequest {
	return PatchPoliciesAPIV2PatchPoliciesIdDashboardPostRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *PatchPoliciesAPIService) V2PatchPoliciesIdDashboardPostExecute(r PatchPoliciesAPIV2PatchPoliciesIdDashboardPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PatchPoliciesAPIService.V2PatchPoliciesIdDashboardPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/patch-policies/{id}/dashboard"
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

type PatchPoliciesAPIV2PatchPoliciesPolicyDetailsGetRequest struct {
	ctx context.Context
	ApiService PatchPoliciesAPI
	page *int64
	pageSize *int64
	sort *[]string
	filter *string
}

func (r PatchPoliciesAPIV2PatchPoliciesPolicyDetailsGetRequest) Page(page int64) PatchPoliciesAPIV2PatchPoliciesPolicyDetailsGetRequest {
	r.page = &page
	return r
}

func (r PatchPoliciesAPIV2PatchPoliciesPolicyDetailsGetRequest) PageSize(pageSize int64) PatchPoliciesAPIV2PatchPoliciesPolicyDetailsGetRequest {
	r.pageSize = &pageSize
	return r
}

// Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma.
func (r PatchPoliciesAPIV2PatchPoliciesPolicyDetailsGetRequest) Sort(sort []string) PatchPoliciesAPIV2PatchPoliciesPolicyDetailsGetRequest {
	r.sort = &sort
	return r
}

// Query in the RSQL format, allowing to filter Patch Policy collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, name, enabled, targetPatchVersion, deploymentMethod, softwareTitleId, softwareTitleConfigurationId, killAppsDelayMinutes, killAppsMessage, isDowngrade, isPatchUnknownVersion, notificationHeader, selfServiceEnforceDeadline, selfServiceDeadline, installButtonText, selfServiceDescription, iconId, reminderFrequency, reminderEnabled. This param can be combined with paging and sorting.
func (r PatchPoliciesAPIV2PatchPoliciesPolicyDetailsGetRequest) Filter(filter string) PatchPoliciesAPIV2PatchPoliciesPolicyDetailsGetRequest {
	r.filter = &filter
	return r
}

func (r PatchPoliciesAPIV2PatchPoliciesPolicyDetailsGetRequest) Execute() (*PatchPolicyDetails, *http.Response, error) {
	return r.ApiService.V2PatchPoliciesPolicyDetailsGetExecute(r)
}

/*
V2PatchPoliciesPolicyDetailsGet Retrieve Patch Policies

Retrieves a list of patch policies.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return PatchPoliciesAPIV2PatchPoliciesPolicyDetailsGetRequest
*/
func (a *PatchPoliciesAPIService) V2PatchPoliciesPolicyDetailsGet(ctx context.Context) PatchPoliciesAPIV2PatchPoliciesPolicyDetailsGetRequest {
	return PatchPoliciesAPIV2PatchPoliciesPolicyDetailsGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PatchPolicyDetails
func (a *PatchPoliciesAPIService) V2PatchPoliciesPolicyDetailsGetExecute(r PatchPoliciesAPIV2PatchPoliciesPolicyDetailsGetRequest) (*PatchPolicyDetails, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PatchPolicyDetails
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PatchPoliciesAPIService.V2PatchPoliciesPolicyDetailsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/patch-policies/policy-details"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	} else {
		var defaultValue int64 = 0
		r.page = &defaultValue
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page-size", r.pageSize, "")
	} else {
		var defaultValue int64 = 100
		r.pageSize = &defaultValue
	}
	if r.sort != nil {
		t := *r.sort
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "sort", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "sort", t, "multi")
		}
	} else {
		defaultValue := []string{"id:asc"}
		r.sort = &defaultValue
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "")
	} else {
		var defaultValue string = ""
		r.filter = &defaultValue
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
