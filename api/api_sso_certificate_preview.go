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
	"os"
)


type SsoCertificatePreviewAPI interface {

	/*
	V1SsoCertDelete Delete the currently configured certificate used by SSO 

	Deletes the currently configured certificate used by SSO.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return SsoCertificatePreviewAPIV1SsoCertDeleteRequest

	Deprecated
	*/
	V1SsoCertDelete(ctx context.Context) SsoCertificatePreviewAPIV1SsoCertDeleteRequest

	// V1SsoCertDeleteExecute executes the request
	// Deprecated
	V1SsoCertDeleteExecute(r SsoCertificatePreviewAPIV1SsoCertDeleteRequest) (*http.Response, error)

	/*
	V1SsoCertDownloadGet Download the certificate currently configured for use with Jamf Pro's SSO configuration 

	Downloads the certificate currently configured for use with Jamf Pro's SSO configuration

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return SsoCertificatePreviewAPIV1SsoCertDownloadGetRequest

	Deprecated
	*/
	V1SsoCertDownloadGet(ctx context.Context) SsoCertificatePreviewAPIV1SsoCertDownloadGetRequest

	// V1SsoCertDownloadGetExecute executes the request
	//  @return *os.File
	// Deprecated
	V1SsoCertDownloadGetExecute(r SsoCertificatePreviewAPIV1SsoCertDownloadGetRequest) (*os.File, *http.Response, error)

	/*
	V1SsoCertGet Retrieve the certificate currently configured for use with SSO 

	Retrieves the certificate currently configured for use with SSO.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return SsoCertificatePreviewAPIV1SsoCertGetRequest

	Deprecated
	*/
	V1SsoCertGet(ctx context.Context) SsoCertificatePreviewAPIV1SsoCertGetRequest

	// V1SsoCertGetExecute executes the request
	//  @return SsoKeystoreWithDetails
	// Deprecated
	V1SsoCertGetExecute(r SsoCertificatePreviewAPIV1SsoCertGetRequest) (*SsoKeystoreWithDetails, *http.Response, error)

	/*
	V1SsoCertParsePost Parse the certificate to get details about certificate type and keys needed to upload certificate file 

	Parse the certificate to get details about certificate type and keys needed to upload certificate file.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return SsoCertificatePreviewAPIV1SsoCertParsePostRequest

	Deprecated
	*/
	V1SsoCertParsePost(ctx context.Context) SsoCertificatePreviewAPIV1SsoCertParsePostRequest

	// V1SsoCertParsePostExecute executes the request
	//  @return SsoKeystoreCertParseResponse
	// Deprecated
	V1SsoCertParsePostExecute(r SsoCertificatePreviewAPIV1SsoCertParsePostRequest) (*SsoKeystoreCertParseResponse, *http.Response, error)

	/*
	V1SsoCertPost Jamf Pro will generate a new certificate and use it to sign SSO 

	Jamf Pro will generate a new certificate and use it to sign SSO requests to the identity provider.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return SsoCertificatePreviewAPIV1SsoCertPostRequest

	Deprecated
	*/
	V1SsoCertPost(ctx context.Context) SsoCertificatePreviewAPIV1SsoCertPostRequest

	// V1SsoCertPostExecute executes the request
	//  @return SsoKeystoreWithDetails
	// Deprecated
	V1SsoCertPostExecute(r SsoCertificatePreviewAPIV1SsoCertPostRequest) (*SsoKeystoreWithDetails, *http.Response, error)

	/*
	V1SsoCertPut Update the certificate used by Jamf Pro to sign SSO requests to the identify provider 

	Update the certificate used by Jamf Pro to sign SSO requests to the identify provider.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return SsoCertificatePreviewAPIV1SsoCertPutRequest

	Deprecated
	*/
	V1SsoCertPut(ctx context.Context) SsoCertificatePreviewAPIV1SsoCertPutRequest

	// V1SsoCertPutExecute executes the request
	//  @return SsoKeystoreWithDetails
	// Deprecated
	V1SsoCertPutExecute(r SsoCertificatePreviewAPIV1SsoCertPutRequest) (*SsoKeystoreWithDetails, *http.Response, error)
}

// SsoCertificatePreviewAPIService SsoCertificatePreviewAPI service
type SsoCertificatePreviewAPIService service

type SsoCertificatePreviewAPIV1SsoCertDeleteRequest struct {
	ctx context.Context
	ApiService SsoCertificatePreviewAPI
}

func (r SsoCertificatePreviewAPIV1SsoCertDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1SsoCertDeleteExecute(r)
}

/*
V1SsoCertDelete Delete the currently configured certificate used by SSO 

Deletes the currently configured certificate used by SSO.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SsoCertificatePreviewAPIV1SsoCertDeleteRequest

Deprecated
*/
func (a *SsoCertificatePreviewAPIService) V1SsoCertDelete(ctx context.Context) SsoCertificatePreviewAPIV1SsoCertDeleteRequest {
	return SsoCertificatePreviewAPIV1SsoCertDeleteRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
// Deprecated
func (a *SsoCertificatePreviewAPIService) V1SsoCertDeleteExecute(r SsoCertificatePreviewAPIV1SsoCertDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SsoCertificatePreviewAPIService.V1SsoCertDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/sso/cert"

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

type SsoCertificatePreviewAPIV1SsoCertDownloadGetRequest struct {
	ctx context.Context
	ApiService SsoCertificatePreviewAPI
}

func (r SsoCertificatePreviewAPIV1SsoCertDownloadGetRequest) Execute() (*os.File, *http.Response, error) {
	return r.ApiService.V1SsoCertDownloadGetExecute(r)
}

/*
V1SsoCertDownloadGet Download the certificate currently configured for use with Jamf Pro's SSO configuration 

Downloads the certificate currently configured for use with Jamf Pro's SSO configuration

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SsoCertificatePreviewAPIV1SsoCertDownloadGetRequest

Deprecated
*/
func (a *SsoCertificatePreviewAPIService) V1SsoCertDownloadGet(ctx context.Context) SsoCertificatePreviewAPIV1SsoCertDownloadGetRequest {
	return SsoCertificatePreviewAPIV1SsoCertDownloadGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return *os.File
// Deprecated
func (a *SsoCertificatePreviewAPIService) V1SsoCertDownloadGetExecute(r SsoCertificatePreviewAPIV1SsoCertDownloadGetRequest) (*os.File, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SsoCertificatePreviewAPIService.V1SsoCertDownloadGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/sso/cert/download"

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
	localVarHTTPHeaderAccepts := []string{"text/plain"}

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

type SsoCertificatePreviewAPIV1SsoCertGetRequest struct {
	ctx context.Context
	ApiService SsoCertificatePreviewAPI
}

func (r SsoCertificatePreviewAPIV1SsoCertGetRequest) Execute() (*SsoKeystoreWithDetails, *http.Response, error) {
	return r.ApiService.V1SsoCertGetExecute(r)
}

/*
V1SsoCertGet Retrieve the certificate currently configured for use with SSO 

Retrieves the certificate currently configured for use with SSO.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SsoCertificatePreviewAPIV1SsoCertGetRequest

Deprecated
*/
func (a *SsoCertificatePreviewAPIService) V1SsoCertGet(ctx context.Context) SsoCertificatePreviewAPIV1SsoCertGetRequest {
	return SsoCertificatePreviewAPIV1SsoCertGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SsoKeystoreWithDetails
// Deprecated
func (a *SsoCertificatePreviewAPIService) V1SsoCertGetExecute(r SsoCertificatePreviewAPIV1SsoCertGetRequest) (*SsoKeystoreWithDetails, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SsoKeystoreWithDetails
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SsoCertificatePreviewAPIService.V1SsoCertGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/sso/cert"

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

type SsoCertificatePreviewAPIV1SsoCertParsePostRequest struct {
	ctx context.Context
	ApiService SsoCertificatePreviewAPI
	ssoKeystoreParse *SsoKeystoreParse
}

func (r SsoCertificatePreviewAPIV1SsoCertParsePostRequest) SsoKeystoreParse(ssoKeystoreParse SsoKeystoreParse) SsoCertificatePreviewAPIV1SsoCertParsePostRequest {
	r.ssoKeystoreParse = &ssoKeystoreParse
	return r
}

func (r SsoCertificatePreviewAPIV1SsoCertParsePostRequest) Execute() (*SsoKeystoreCertParseResponse, *http.Response, error) {
	return r.ApiService.V1SsoCertParsePostExecute(r)
}

/*
V1SsoCertParsePost Parse the certificate to get details about certificate type and keys needed to upload certificate file 

Parse the certificate to get details about certificate type and keys needed to upload certificate file.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SsoCertificatePreviewAPIV1SsoCertParsePostRequest

Deprecated
*/
func (a *SsoCertificatePreviewAPIService) V1SsoCertParsePost(ctx context.Context) SsoCertificatePreviewAPIV1SsoCertParsePostRequest {
	return SsoCertificatePreviewAPIV1SsoCertParsePostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SsoKeystoreCertParseResponse
// Deprecated
func (a *SsoCertificatePreviewAPIService) V1SsoCertParsePostExecute(r SsoCertificatePreviewAPIV1SsoCertParsePostRequest) (*SsoKeystoreCertParseResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SsoKeystoreCertParseResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SsoCertificatePreviewAPIService.V1SsoCertParsePost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/sso/cert/parse"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.ssoKeystoreParse == nil {
		return localVarReturnValue, nil, reportError("ssoKeystoreParse is required and must be specified")
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
	localVarPostBody = r.ssoKeystoreParse
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

type SsoCertificatePreviewAPIV1SsoCertPostRequest struct {
	ctx context.Context
	ApiService SsoCertificatePreviewAPI
}

func (r SsoCertificatePreviewAPIV1SsoCertPostRequest) Execute() (*SsoKeystoreWithDetails, *http.Response, error) {
	return r.ApiService.V1SsoCertPostExecute(r)
}

/*
V1SsoCertPost Jamf Pro will generate a new certificate and use it to sign SSO 

Jamf Pro will generate a new certificate and use it to sign SSO requests to the identity provider.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SsoCertificatePreviewAPIV1SsoCertPostRequest

Deprecated
*/
func (a *SsoCertificatePreviewAPIService) V1SsoCertPost(ctx context.Context) SsoCertificatePreviewAPIV1SsoCertPostRequest {
	return SsoCertificatePreviewAPIV1SsoCertPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SsoKeystoreWithDetails
// Deprecated
func (a *SsoCertificatePreviewAPIService) V1SsoCertPostExecute(r SsoCertificatePreviewAPIV1SsoCertPostRequest) (*SsoKeystoreWithDetails, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SsoKeystoreWithDetails
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SsoCertificatePreviewAPIService.V1SsoCertPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/sso/cert"

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

type SsoCertificatePreviewAPIV1SsoCertPutRequest struct {
	ctx context.Context
	ApiService SsoCertificatePreviewAPI
	ssoKeystore *SsoKeystore
}

func (r SsoCertificatePreviewAPIV1SsoCertPutRequest) SsoKeystore(ssoKeystore SsoKeystore) SsoCertificatePreviewAPIV1SsoCertPutRequest {
	r.ssoKeystore = &ssoKeystore
	return r
}

func (r SsoCertificatePreviewAPIV1SsoCertPutRequest) Execute() (*SsoKeystoreWithDetails, *http.Response, error) {
	return r.ApiService.V1SsoCertPutExecute(r)
}

/*
V1SsoCertPut Update the certificate used by Jamf Pro to sign SSO requests to the identify provider 

Update the certificate used by Jamf Pro to sign SSO requests to the identify provider.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SsoCertificatePreviewAPIV1SsoCertPutRequest

Deprecated
*/
func (a *SsoCertificatePreviewAPIService) V1SsoCertPut(ctx context.Context) SsoCertificatePreviewAPIV1SsoCertPutRequest {
	return SsoCertificatePreviewAPIV1SsoCertPutRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SsoKeystoreWithDetails
// Deprecated
func (a *SsoCertificatePreviewAPIService) V1SsoCertPutExecute(r SsoCertificatePreviewAPIV1SsoCertPutRequest) (*SsoKeystoreWithDetails, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SsoKeystoreWithDetails
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SsoCertificatePreviewAPIService.V1SsoCertPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/sso/cert"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.ssoKeystore == nil {
		return localVarReturnValue, nil, reportError("ssoKeystore is required and must be specified")
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
	localVarPostBody = r.ssoKeystore
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
