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
	"reflect"
)


type SmtpServerAPI interface {

	/*
	V1SmtpServerGet Finds the Jamf Pro SMTP Server information 

	Finds the Jamf Pro SMTP Server information


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return SmtpServerAPIV1SmtpServerGetRequest

	Deprecated
	*/
	V1SmtpServerGet(ctx context.Context) SmtpServerAPIV1SmtpServerGetRequest

	// V1SmtpServerGetExecute executes the request
	//  @return SmtpServer
	// Deprecated
	V1SmtpServerGetExecute(r SmtpServerAPIV1SmtpServerGetRequest) (*SmtpServer, *http.Response, error)

	/*
	V1SmtpServerHistoryGet Get specified SMTP Server history object 

	Get specified SMTP Server history object


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return SmtpServerAPIV1SmtpServerHistoryGetRequest
	*/
	V1SmtpServerHistoryGet(ctx context.Context) SmtpServerAPIV1SmtpServerHistoryGetRequest

	// V1SmtpServerHistoryGetExecute executes the request
	//  @return HistorySearchResultsV1
	V1SmtpServerHistoryGetExecute(r SmtpServerAPIV1SmtpServerHistoryGetRequest) (*HistorySearchResultsV1, *http.Response, error)

	/*
	V1SmtpServerHistoryPost Add SMTP Server history object notes 

	Adds SMTP Server history object notes


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return SmtpServerAPIV1SmtpServerHistoryPostRequest
	*/
	V1SmtpServerHistoryPost(ctx context.Context) SmtpServerAPIV1SmtpServerHistoryPostRequest

	// V1SmtpServerHistoryPostExecute executes the request
	//  @return HrefResponse
	V1SmtpServerHistoryPostExecute(r SmtpServerAPIV1SmtpServerHistoryPostRequest) (*HrefResponse, *http.Response, error)

	/*
	V1SmtpServerPut Updates Jamf Pro SMTP Server information 

	Updates Jamf Pro SMTP Server information. If requiresAuthentication is set to true, a username and password must be provided


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return SmtpServerAPIV1SmtpServerPutRequest

	Deprecated
	*/
	V1SmtpServerPut(ctx context.Context) SmtpServerAPIV1SmtpServerPutRequest

	// V1SmtpServerPutExecute executes the request
	//  @return SmtpServer
	// Deprecated
	V1SmtpServerPutExecute(r SmtpServerAPIV1SmtpServerPutRequest) (*SmtpServer, *http.Response, error)

	/*
	V1SmtpServerTestPost Test functionality of an SMTP Server

	Test functionality of an SMTP Server

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return SmtpServerAPIV1SmtpServerTestPostRequest
	*/
	V1SmtpServerTestPost(ctx context.Context) SmtpServerAPIV1SmtpServerTestPostRequest

	// V1SmtpServerTestPostExecute executes the request
	V1SmtpServerTestPostExecute(r SmtpServerAPIV1SmtpServerTestPostRequest) (*http.Response, error)

	/*
	V2SmtpServerGet Finds the Jamf Pro SMTP Server information 

	Finds the Jamf Pro SMTP Server information


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return SmtpServerAPIV2SmtpServerGetRequest
	*/
	V2SmtpServerGet(ctx context.Context) SmtpServerAPIV2SmtpServerGetRequest

	// V2SmtpServerGetExecute executes the request
	//  @return SmtpServerV2
	V2SmtpServerGetExecute(r SmtpServerAPIV2SmtpServerGetRequest) (*SmtpServerV2, *http.Response, error)

	/*
	V2SmtpServerPut Updates Jamf Pro SMTP Server information 

	Updates Jamf Pro SMTP Server information. If requiresAuthentication is set to true, a username and password must be provided


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return SmtpServerAPIV2SmtpServerPutRequest
	*/
	V2SmtpServerPut(ctx context.Context) SmtpServerAPIV2SmtpServerPutRequest

	// V2SmtpServerPutExecute executes the request
	//  @return SmtpServerV2
	V2SmtpServerPutExecute(r SmtpServerAPIV2SmtpServerPutRequest) (*SmtpServerV2, *http.Response, error)
}

// SmtpServerAPIService SmtpServerAPI service
type SmtpServerAPIService service

type SmtpServerAPIV1SmtpServerGetRequest struct {
	ctx context.Context
	ApiService SmtpServerAPI
}

func (r SmtpServerAPIV1SmtpServerGetRequest) Execute() (*SmtpServer, *http.Response, error) {
	return r.ApiService.V1SmtpServerGetExecute(r)
}

/*
V1SmtpServerGet Finds the Jamf Pro SMTP Server information 

Finds the Jamf Pro SMTP Server information


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SmtpServerAPIV1SmtpServerGetRequest

Deprecated
*/
func (a *SmtpServerAPIService) V1SmtpServerGet(ctx context.Context) SmtpServerAPIV1SmtpServerGetRequest {
	return SmtpServerAPIV1SmtpServerGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SmtpServer
// Deprecated
func (a *SmtpServerAPIService) V1SmtpServerGetExecute(r SmtpServerAPIV1SmtpServerGetRequest) (*SmtpServer, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SmtpServer
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SmtpServerAPIService.V1SmtpServerGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/smtp-server"

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

type SmtpServerAPIV1SmtpServerHistoryGetRequest struct {
	ctx context.Context
	ApiService SmtpServerAPI
	page *int64
	pageSize *int64
	sort *[]string
	filter *string
}

func (r SmtpServerAPIV1SmtpServerHistoryGetRequest) Page(page int64) SmtpServerAPIV1SmtpServerHistoryGetRequest {
	r.page = &page
	return r
}

func (r SmtpServerAPIV1SmtpServerHistoryGetRequest) PageSize(pageSize int64) SmtpServerAPIV1SmtpServerHistoryGetRequest {
	r.pageSize = &pageSize
	return r
}

// Sorting criteria in the format: property:asc/desc. Default sort is name:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;date:desc,username:asc 
func (r SmtpServerAPIV1SmtpServerHistoryGetRequest) Sort(sort []string) SmtpServerAPIV1SmtpServerHistoryGetRequest {
	r.sort = &sort
	return r
}

// Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: filter&#x3D;username!&#x3D;admin and details&#x3D;&#x3D;*disabled* and date&lt;2019-12-15
func (r SmtpServerAPIV1SmtpServerHistoryGetRequest) Filter(filter string) SmtpServerAPIV1SmtpServerHistoryGetRequest {
	r.filter = &filter
	return r
}

func (r SmtpServerAPIV1SmtpServerHistoryGetRequest) Execute() (*HistorySearchResultsV1, *http.Response, error) {
	return r.ApiService.V1SmtpServerHistoryGetExecute(r)
}

/*
V1SmtpServerHistoryGet Get specified SMTP Server history object 

Get specified SMTP Server history object


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SmtpServerAPIV1SmtpServerHistoryGetRequest
*/
func (a *SmtpServerAPIService) V1SmtpServerHistoryGet(ctx context.Context) SmtpServerAPIV1SmtpServerHistoryGetRequest {
	return SmtpServerAPIV1SmtpServerHistoryGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return HistorySearchResultsV1
func (a *SmtpServerAPIService) V1SmtpServerHistoryGetExecute(r SmtpServerAPIV1SmtpServerHistoryGetRequest) (*HistorySearchResultsV1, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *HistorySearchResultsV1
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SmtpServerAPIService.V1SmtpServerHistoryGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/smtp-server/history"

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
		defaultValue := []string{"date:desc"}
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

type SmtpServerAPIV1SmtpServerHistoryPostRequest struct {
	ctx context.Context
	ApiService SmtpServerAPI
	objectHistoryNote *ObjectHistoryNote
}

// History notes to create
func (r SmtpServerAPIV1SmtpServerHistoryPostRequest) ObjectHistoryNote(objectHistoryNote ObjectHistoryNote) SmtpServerAPIV1SmtpServerHistoryPostRequest {
	r.objectHistoryNote = &objectHistoryNote
	return r
}

func (r SmtpServerAPIV1SmtpServerHistoryPostRequest) Execute() (*HrefResponse, *http.Response, error) {
	return r.ApiService.V1SmtpServerHistoryPostExecute(r)
}

/*
V1SmtpServerHistoryPost Add SMTP Server history object notes 

Adds SMTP Server history object notes


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SmtpServerAPIV1SmtpServerHistoryPostRequest
*/
func (a *SmtpServerAPIService) V1SmtpServerHistoryPost(ctx context.Context) SmtpServerAPIV1SmtpServerHistoryPostRequest {
	return SmtpServerAPIV1SmtpServerHistoryPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return HrefResponse
func (a *SmtpServerAPIService) V1SmtpServerHistoryPostExecute(r SmtpServerAPIV1SmtpServerHistoryPostRequest) (*HrefResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *HrefResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SmtpServerAPIService.V1SmtpServerHistoryPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/smtp-server/history"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.objectHistoryNote == nil {
		return localVarReturnValue, nil, reportError("objectHistoryNote is required and must be specified")
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
	localVarPostBody = r.objectHistoryNote
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
		if localVarHTTPResponse.StatusCode == 503 {
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

type SmtpServerAPIV1SmtpServerPutRequest struct {
	ctx context.Context
	ApiService SmtpServerAPI
	smtpServer *SmtpServer
}

// SMTP Server to update
func (r SmtpServerAPIV1SmtpServerPutRequest) SmtpServer(smtpServer SmtpServer) SmtpServerAPIV1SmtpServerPutRequest {
	r.smtpServer = &smtpServer
	return r
}

func (r SmtpServerAPIV1SmtpServerPutRequest) Execute() (*SmtpServer, *http.Response, error) {
	return r.ApiService.V1SmtpServerPutExecute(r)
}

/*
V1SmtpServerPut Updates Jamf Pro SMTP Server information 

Updates Jamf Pro SMTP Server information. If requiresAuthentication is set to true, a username and password must be provided


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SmtpServerAPIV1SmtpServerPutRequest

Deprecated
*/
func (a *SmtpServerAPIService) V1SmtpServerPut(ctx context.Context) SmtpServerAPIV1SmtpServerPutRequest {
	return SmtpServerAPIV1SmtpServerPutRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SmtpServer
// Deprecated
func (a *SmtpServerAPIService) V1SmtpServerPutExecute(r SmtpServerAPIV1SmtpServerPutRequest) (*SmtpServer, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SmtpServer
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SmtpServerAPIService.V1SmtpServerPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/smtp-server"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.smtpServer == nil {
		return localVarReturnValue, nil, reportError("smtpServer is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/merge-patch+json"}

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
	localVarPostBody = r.smtpServer
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
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
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

type SmtpServerAPIV1SmtpServerTestPostRequest struct {
	ctx context.Context
	ApiService SmtpServerAPI
	smtpServerTest *SmtpServerTest
}

// Recipient email to test SMTP Server
func (r SmtpServerAPIV1SmtpServerTestPostRequest) SmtpServerTest(smtpServerTest SmtpServerTest) SmtpServerAPIV1SmtpServerTestPostRequest {
	r.smtpServerTest = &smtpServerTest
	return r
}

func (r SmtpServerAPIV1SmtpServerTestPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1SmtpServerTestPostExecute(r)
}

/*
V1SmtpServerTestPost Test functionality of an SMTP Server

Test functionality of an SMTP Server

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SmtpServerAPIV1SmtpServerTestPostRequest
*/
func (a *SmtpServerAPIService) V1SmtpServerTestPost(ctx context.Context) SmtpServerAPIV1SmtpServerTestPostRequest {
	return SmtpServerAPIV1SmtpServerTestPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *SmtpServerAPIService) V1SmtpServerTestPostExecute(r SmtpServerAPIV1SmtpServerTestPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SmtpServerAPIService.V1SmtpServerTestPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/smtp-server/test"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.smtpServerTest == nil {
		return nil, reportError("smtpServerTest is required and must be specified")
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
	localVarPostBody = r.smtpServerTest
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

type SmtpServerAPIV2SmtpServerGetRequest struct {
	ctx context.Context
	ApiService SmtpServerAPI
}

func (r SmtpServerAPIV2SmtpServerGetRequest) Execute() (*SmtpServerV2, *http.Response, error) {
	return r.ApiService.V2SmtpServerGetExecute(r)
}

/*
V2SmtpServerGet Finds the Jamf Pro SMTP Server information 

Finds the Jamf Pro SMTP Server information


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SmtpServerAPIV2SmtpServerGetRequest
*/
func (a *SmtpServerAPIService) V2SmtpServerGet(ctx context.Context) SmtpServerAPIV2SmtpServerGetRequest {
	return SmtpServerAPIV2SmtpServerGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SmtpServerV2
func (a *SmtpServerAPIService) V2SmtpServerGetExecute(r SmtpServerAPIV2SmtpServerGetRequest) (*SmtpServerV2, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SmtpServerV2
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SmtpServerAPIService.V2SmtpServerGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/smtp-server"

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

type SmtpServerAPIV2SmtpServerPutRequest struct {
	ctx context.Context
	ApiService SmtpServerAPI
	smtpServerV2 *SmtpServerV2
}

// SMTP Server to update
func (r SmtpServerAPIV2SmtpServerPutRequest) SmtpServerV2(smtpServerV2 SmtpServerV2) SmtpServerAPIV2SmtpServerPutRequest {
	r.smtpServerV2 = &smtpServerV2
	return r
}

func (r SmtpServerAPIV2SmtpServerPutRequest) Execute() (*SmtpServerV2, *http.Response, error) {
	return r.ApiService.V2SmtpServerPutExecute(r)
}

/*
V2SmtpServerPut Updates Jamf Pro SMTP Server information 

Updates Jamf Pro SMTP Server information. If requiresAuthentication is set to true, a username and password must be provided


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SmtpServerAPIV2SmtpServerPutRequest
*/
func (a *SmtpServerAPIService) V2SmtpServerPut(ctx context.Context) SmtpServerAPIV2SmtpServerPutRequest {
	return SmtpServerAPIV2SmtpServerPutRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SmtpServerV2
func (a *SmtpServerAPIService) V2SmtpServerPutExecute(r SmtpServerAPIV2SmtpServerPutRequest) (*SmtpServerV2, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SmtpServerV2
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SmtpServerAPIService.V2SmtpServerPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/smtp-server"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.smtpServerV2 == nil {
		return localVarReturnValue, nil, reportError("smtpServerV2 is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/merge-patch+json"}

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
	localVarPostBody = r.smtpServerV2
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
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
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
