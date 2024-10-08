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


type JamfContentDistributionServerAPI interface {

	/*
	V1JcdsFilesFileNameDelete Delete a file from the Jamf Content Distribution Server

	Delete a file by filename from the Jamf Content Distribution Server.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param fileName Name of the file that will be deleted from the Jamf Content Distribution Server.
	@return JamfContentDistributionServerAPIV1JcdsFilesFileNameDeleteRequest
	*/
	V1JcdsFilesFileNameDelete(ctx context.Context, fileName string) JamfContentDistributionServerAPIV1JcdsFilesFileNameDeleteRequest

	// V1JcdsFilesFileNameDeleteExecute executes the request
	V1JcdsFilesFileNameDeleteExecute(r JamfContentDistributionServerAPIV1JcdsFilesFileNameDeleteRequest) (*http.Response, error)

	/*
	V1JcdsFilesFileNameGet Retrieve a download URL for a specific file from the Jamf Content Distribution Server

	Retrieve a download URL for a specific file from the Jamf Content Distribution Server.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param fileName Name of the file stored in the Jamf Content Distribution Server.
	@return JamfContentDistributionServerAPIV1JcdsFilesFileNameGetRequest
	*/
	V1JcdsFilesFileNameGet(ctx context.Context, fileName string) JamfContentDistributionServerAPIV1JcdsFilesFileNameGetRequest

	// V1JcdsFilesFileNameGetExecute executes the request
	//  @return DownloadUrl
	V1JcdsFilesFileNameGetExecute(r JamfContentDistributionServerAPIV1JcdsFilesFileNameGetRequest) (*DownloadUrl, *http.Response, error)

	/*
	V1JcdsFilesGet Retrieve a list of files and file metadata from the Jamf Content Distribution Server

	Retrieve a list of files and file metadata from the Jamf Content Distribution Server.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return JamfContentDistributionServerAPIV1JcdsFilesGetRequest
	*/
	V1JcdsFilesGet(ctx context.Context) JamfContentDistributionServerAPIV1JcdsFilesGetRequest

	// V1JcdsFilesGetExecute executes the request
	//  @return []FileData
	V1JcdsFilesGetExecute(r JamfContentDistributionServerAPIV1JcdsFilesGetRequest) ([]FileData, *http.Response, error)

	/*
	V1JcdsFilesPost Initiate an upload to the Jamf Content Distribution Server

	Creates a temporary record and returns the credentials and information needed for uploading the file to the Jamf Content Distribution Server.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return JamfContentDistributionServerAPIV1JcdsFilesPostRequest
	*/
	V1JcdsFilesPost(ctx context.Context) JamfContentDistributionServerAPIV1JcdsFilesPostRequest

	// V1JcdsFilesPostExecute executes the request
	//  @return Credentials
	V1JcdsFilesPostExecute(r JamfContentDistributionServerAPIV1JcdsFilesPostRequest) (*Credentials, *http.Response, error)

	/*
	V1JcdsPropertiesGet Gets information about JCDS distribution points. 

	Gets string values of jcds2Enabled, fileStreamEndpointEnabled, and maxChunkSize properties


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return JamfContentDistributionServerAPIV1JcdsPropertiesGetRequest
	*/
	V1JcdsPropertiesGet(ctx context.Context) JamfContentDistributionServerAPIV1JcdsPropertiesGetRequest

	// V1JcdsPropertiesGetExecute executes the request
	//  @return JcdsProperties
	V1JcdsPropertiesGetExecute(r JamfContentDistributionServerAPIV1JcdsPropertiesGetRequest) (*JcdsProperties, *http.Response, error)

	/*
	V1JcdsRenewCredentialsPost Renew credentials for an upload to the Jamf Content Distribution Server

	Renews the credentials needed for uploading the file to the Jamf Content Distribution Server.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return JamfContentDistributionServerAPIV1JcdsRenewCredentialsPostRequest
	*/
	V1JcdsRenewCredentialsPost(ctx context.Context) JamfContentDistributionServerAPIV1JcdsRenewCredentialsPostRequest

	// V1JcdsRenewCredentialsPostExecute executes the request
	//  @return Credentials
	V1JcdsRenewCredentialsPostExecute(r JamfContentDistributionServerAPIV1JcdsRenewCredentialsPostRequest) (*Credentials, *http.Response, error)
}

// JamfContentDistributionServerAPIService JamfContentDistributionServerAPI service
type JamfContentDistributionServerAPIService service

type JamfContentDistributionServerAPIV1JcdsFilesFileNameDeleteRequest struct {
	ctx context.Context
	ApiService JamfContentDistributionServerAPI
	fileName string
}

func (r JamfContentDistributionServerAPIV1JcdsFilesFileNameDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1JcdsFilesFileNameDeleteExecute(r)
}

/*
V1JcdsFilesFileNameDelete Delete a file from the Jamf Content Distribution Server

Delete a file by filename from the Jamf Content Distribution Server.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param fileName Name of the file that will be deleted from the Jamf Content Distribution Server.
 @return JamfContentDistributionServerAPIV1JcdsFilesFileNameDeleteRequest
*/
func (a *JamfContentDistributionServerAPIService) V1JcdsFilesFileNameDelete(ctx context.Context, fileName string) JamfContentDistributionServerAPIV1JcdsFilesFileNameDeleteRequest {
	return JamfContentDistributionServerAPIV1JcdsFilesFileNameDeleteRequest{
		ApiService: a,
		ctx: ctx,
		fileName: fileName,
	}
}

// Execute executes the request
func (a *JamfContentDistributionServerAPIService) V1JcdsFilesFileNameDeleteExecute(r JamfContentDistributionServerAPIV1JcdsFilesFileNameDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JamfContentDistributionServerAPIService.V1JcdsFilesFileNameDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/jcds/files/{fileName}"
	localVarPath = strings.Replace(localVarPath, "{"+"fileName"+"}", url.PathEscape(parameterValueToString(r.fileName, "fileName")), -1)

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

type JamfContentDistributionServerAPIV1JcdsFilesFileNameGetRequest struct {
	ctx context.Context
	ApiService JamfContentDistributionServerAPI
	fileName string
}

func (r JamfContentDistributionServerAPIV1JcdsFilesFileNameGetRequest) Execute() (*DownloadUrl, *http.Response, error) {
	return r.ApiService.V1JcdsFilesFileNameGetExecute(r)
}

/*
V1JcdsFilesFileNameGet Retrieve a download URL for a specific file from the Jamf Content Distribution Server

Retrieve a download URL for a specific file from the Jamf Content Distribution Server.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param fileName Name of the file stored in the Jamf Content Distribution Server.
 @return JamfContentDistributionServerAPIV1JcdsFilesFileNameGetRequest
*/
func (a *JamfContentDistributionServerAPIService) V1JcdsFilesFileNameGet(ctx context.Context, fileName string) JamfContentDistributionServerAPIV1JcdsFilesFileNameGetRequest {
	return JamfContentDistributionServerAPIV1JcdsFilesFileNameGetRequest{
		ApiService: a,
		ctx: ctx,
		fileName: fileName,
	}
}

// Execute executes the request
//  @return DownloadUrl
func (a *JamfContentDistributionServerAPIService) V1JcdsFilesFileNameGetExecute(r JamfContentDistributionServerAPIV1JcdsFilesFileNameGetRequest) (*DownloadUrl, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DownloadUrl
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JamfContentDistributionServerAPIService.V1JcdsFilesFileNameGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/jcds/files/{fileName}"
	localVarPath = strings.Replace(localVarPath, "{"+"fileName"+"}", url.PathEscape(parameterValueToString(r.fileName, "fileName")), -1)

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

type JamfContentDistributionServerAPIV1JcdsFilesGetRequest struct {
	ctx context.Context
	ApiService JamfContentDistributionServerAPI
}

func (r JamfContentDistributionServerAPIV1JcdsFilesGetRequest) Execute() ([]FileData, *http.Response, error) {
	return r.ApiService.V1JcdsFilesGetExecute(r)
}

/*
V1JcdsFilesGet Retrieve a list of files and file metadata from the Jamf Content Distribution Server

Retrieve a list of files and file metadata from the Jamf Content Distribution Server.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return JamfContentDistributionServerAPIV1JcdsFilesGetRequest
*/
func (a *JamfContentDistributionServerAPIService) V1JcdsFilesGet(ctx context.Context) JamfContentDistributionServerAPIV1JcdsFilesGetRequest {
	return JamfContentDistributionServerAPIV1JcdsFilesGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []FileData
func (a *JamfContentDistributionServerAPIService) V1JcdsFilesGetExecute(r JamfContentDistributionServerAPIV1JcdsFilesGetRequest) ([]FileData, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []FileData
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JamfContentDistributionServerAPIService.V1JcdsFilesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/jcds/files"

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

type JamfContentDistributionServerAPIV1JcdsFilesPostRequest struct {
	ctx context.Context
	ApiService JamfContentDistributionServerAPI
}

func (r JamfContentDistributionServerAPIV1JcdsFilesPostRequest) Execute() (*Credentials, *http.Response, error) {
	return r.ApiService.V1JcdsFilesPostExecute(r)
}

/*
V1JcdsFilesPost Initiate an upload to the Jamf Content Distribution Server

Creates a temporary record and returns the credentials and information needed for uploading the file to the Jamf Content Distribution Server.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return JamfContentDistributionServerAPIV1JcdsFilesPostRequest
*/
func (a *JamfContentDistributionServerAPIService) V1JcdsFilesPost(ctx context.Context) JamfContentDistributionServerAPIV1JcdsFilesPostRequest {
	return JamfContentDistributionServerAPIV1JcdsFilesPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Credentials
func (a *JamfContentDistributionServerAPIService) V1JcdsFilesPostExecute(r JamfContentDistributionServerAPIV1JcdsFilesPostRequest) (*Credentials, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Credentials
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JamfContentDistributionServerAPIService.V1JcdsFilesPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/jcds/files"

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

type JamfContentDistributionServerAPIV1JcdsPropertiesGetRequest struct {
	ctx context.Context
	ApiService JamfContentDistributionServerAPI
}

func (r JamfContentDistributionServerAPIV1JcdsPropertiesGetRequest) Execute() (*JcdsProperties, *http.Response, error) {
	return r.ApiService.V1JcdsPropertiesGetExecute(r)
}

/*
V1JcdsPropertiesGet Gets information about JCDS distribution points. 

Gets string values of jcds2Enabled, fileStreamEndpointEnabled, and maxChunkSize properties


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return JamfContentDistributionServerAPIV1JcdsPropertiesGetRequest
*/
func (a *JamfContentDistributionServerAPIService) V1JcdsPropertiesGet(ctx context.Context) JamfContentDistributionServerAPIV1JcdsPropertiesGetRequest {
	return JamfContentDistributionServerAPIV1JcdsPropertiesGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return JcdsProperties
func (a *JamfContentDistributionServerAPIService) V1JcdsPropertiesGetExecute(r JamfContentDistributionServerAPIV1JcdsPropertiesGetRequest) (*JcdsProperties, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *JcdsProperties
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JamfContentDistributionServerAPIService.V1JcdsPropertiesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/jcds/properties"

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

type JamfContentDistributionServerAPIV1JcdsRenewCredentialsPostRequest struct {
	ctx context.Context
	ApiService JamfContentDistributionServerAPI
}

func (r JamfContentDistributionServerAPIV1JcdsRenewCredentialsPostRequest) Execute() (*Credentials, *http.Response, error) {
	return r.ApiService.V1JcdsRenewCredentialsPostExecute(r)
}

/*
V1JcdsRenewCredentialsPost Renew credentials for an upload to the Jamf Content Distribution Server

Renews the credentials needed for uploading the file to the Jamf Content Distribution Server.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return JamfContentDistributionServerAPIV1JcdsRenewCredentialsPostRequest
*/
func (a *JamfContentDistributionServerAPIService) V1JcdsRenewCredentialsPost(ctx context.Context) JamfContentDistributionServerAPIV1JcdsRenewCredentialsPostRequest {
	return JamfContentDistributionServerAPIV1JcdsRenewCredentialsPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Credentials
func (a *JamfContentDistributionServerAPIService) V1JcdsRenewCredentialsPostExecute(r JamfContentDistributionServerAPIV1JcdsRenewCredentialsPostRequest) (*Credentials, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Credentials
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JamfContentDistributionServerAPIService.V1JcdsRenewCredentialsPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/jcds/renew-credentials"

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
