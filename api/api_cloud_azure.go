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


type CloudAzureAPI interface {

	/*
	V1CloudAzureDefaultsMappingsGet Get default mappings

	This is the default set of attributes that allows you to return the data you need from Azure AD. Some fields may be empty and may be edited when creating a new configuration.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return CloudAzureAPIV1CloudAzureDefaultsMappingsGetRequest
	*/
	V1CloudAzureDefaultsMappingsGet(ctx context.Context) CloudAzureAPIV1CloudAzureDefaultsMappingsGetRequest

	// V1CloudAzureDefaultsMappingsGetExecute executes the request
	//  @return AzureMappings
	V1CloudAzureDefaultsMappingsGetExecute(r CloudAzureAPIV1CloudAzureDefaultsMappingsGetRequest) (*AzureMappings, *http.Response, error)

	/*
	V1CloudAzureDefaultsServerConfigurationGet Get default server configuration

	This is the default set of attributes that allows you to return the data you need from Azure AD. Some fields may be empty and may be edited when creating a new configuration.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return CloudAzureAPIV1CloudAzureDefaultsServerConfigurationGetRequest
	*/
	V1CloudAzureDefaultsServerConfigurationGet(ctx context.Context) CloudAzureAPIV1CloudAzureDefaultsServerConfigurationGetRequest

	// V1CloudAzureDefaultsServerConfigurationGetExecute executes the request
	//  @return AzureServerConfiguration
	V1CloudAzureDefaultsServerConfigurationGetExecute(r CloudAzureAPIV1CloudAzureDefaultsServerConfigurationGetRequest) (*AzureServerConfiguration, *http.Response, error)

	/*
	V1CloudAzureIdDelete Delete Cloud Identity Provider configuration.

	Delete Cloud Identity Provider configuration.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Cloud Identity Provider identifier
	@return CloudAzureAPIV1CloudAzureIdDeleteRequest
	*/
	V1CloudAzureIdDelete(ctx context.Context, id string) CloudAzureAPIV1CloudAzureIdDeleteRequest

	// V1CloudAzureIdDeleteExecute executes the request
	V1CloudAzureIdDeleteExecute(r CloudAzureAPIV1CloudAzureIdDeleteRequest) (*http.Response, error)

	/*
	V1CloudAzureIdGet Get Azure Cloud Identity Provider configuration with given ID.

	Get Azure Cloud Identity Provider configuration with given ID.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Cloud Identity Provider identifier
	@return CloudAzureAPIV1CloudAzureIdGetRequest
	*/
	V1CloudAzureIdGet(ctx context.Context, id string) CloudAzureAPIV1CloudAzureIdGetRequest

	// V1CloudAzureIdGetExecute executes the request
	//  @return AzureConfiguration
	V1CloudAzureIdGetExecute(r CloudAzureAPIV1CloudAzureIdGetRequest) (*AzureConfiguration, *http.Response, error)

	/*
	V1CloudAzureIdPut Update Azure Cloud Identity Provider configuration

	Update Azure Cloud Identity Provider configuration. Cannot be used for partial updates, all content body must be sent.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Cloud Identity Provider identifier
	@return CloudAzureAPIV1CloudAzureIdPutRequest
	*/
	V1CloudAzureIdPut(ctx context.Context, id string) CloudAzureAPIV1CloudAzureIdPutRequest

	// V1CloudAzureIdPutExecute executes the request
	//  @return AzureConfiguration
	V1CloudAzureIdPutExecute(r CloudAzureAPIV1CloudAzureIdPutRequest) (*AzureConfiguration, *http.Response, error)

	/*
	V1CloudAzurePost Create Azure Cloud Identity Provider configuration

	Create new Azure Cloud Identity Provider configuration with unique display name.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return CloudAzureAPIV1CloudAzurePostRequest
	*/
	V1CloudAzurePost(ctx context.Context) CloudAzureAPIV1CloudAzurePostRequest

	// V1CloudAzurePostExecute executes the request
	//  @return HrefResponse
	V1CloudAzurePostExecute(r CloudAzureAPIV1CloudAzurePostRequest) (*HrefResponse, *http.Response, error)
}

// CloudAzureAPIService CloudAzureAPI service
type CloudAzureAPIService service

type CloudAzureAPIV1CloudAzureDefaultsMappingsGetRequest struct {
	ctx context.Context
	ApiService CloudAzureAPI
}

func (r CloudAzureAPIV1CloudAzureDefaultsMappingsGetRequest) Execute() (*AzureMappings, *http.Response, error) {
	return r.ApiService.V1CloudAzureDefaultsMappingsGetExecute(r)
}

/*
V1CloudAzureDefaultsMappingsGet Get default mappings

This is the default set of attributes that allows you to return the data you need from Azure AD. Some fields may be empty and may be edited when creating a new configuration.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return CloudAzureAPIV1CloudAzureDefaultsMappingsGetRequest
*/
func (a *CloudAzureAPIService) V1CloudAzureDefaultsMappingsGet(ctx context.Context) CloudAzureAPIV1CloudAzureDefaultsMappingsGetRequest {
	return CloudAzureAPIV1CloudAzureDefaultsMappingsGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AzureMappings
func (a *CloudAzureAPIService) V1CloudAzureDefaultsMappingsGetExecute(r CloudAzureAPIV1CloudAzureDefaultsMappingsGetRequest) (*AzureMappings, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AzureMappings
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudAzureAPIService.V1CloudAzureDefaultsMappingsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/cloud-azure/defaults/mappings"

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
			return localVarReturnValue, localVarHTTPResponse, newErr
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

type CloudAzureAPIV1CloudAzureDefaultsServerConfigurationGetRequest struct {
	ctx context.Context
	ApiService CloudAzureAPI
}

func (r CloudAzureAPIV1CloudAzureDefaultsServerConfigurationGetRequest) Execute() (*AzureServerConfiguration, *http.Response, error) {
	return r.ApiService.V1CloudAzureDefaultsServerConfigurationGetExecute(r)
}

/*
V1CloudAzureDefaultsServerConfigurationGet Get default server configuration

This is the default set of attributes that allows you to return the data you need from Azure AD. Some fields may be empty and may be edited when creating a new configuration.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return CloudAzureAPIV1CloudAzureDefaultsServerConfigurationGetRequest
*/
func (a *CloudAzureAPIService) V1CloudAzureDefaultsServerConfigurationGet(ctx context.Context) CloudAzureAPIV1CloudAzureDefaultsServerConfigurationGetRequest {
	return CloudAzureAPIV1CloudAzureDefaultsServerConfigurationGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AzureServerConfiguration
func (a *CloudAzureAPIService) V1CloudAzureDefaultsServerConfigurationGetExecute(r CloudAzureAPIV1CloudAzureDefaultsServerConfigurationGetRequest) (*AzureServerConfiguration, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AzureServerConfiguration
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudAzureAPIService.V1CloudAzureDefaultsServerConfigurationGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/cloud-azure/defaults/server-configuration"

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
			return localVarReturnValue, localVarHTTPResponse, newErr
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

type CloudAzureAPIV1CloudAzureIdDeleteRequest struct {
	ctx context.Context
	ApiService CloudAzureAPI
	id string
}

func (r CloudAzureAPIV1CloudAzureIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1CloudAzureIdDeleteExecute(r)
}

/*
V1CloudAzureIdDelete Delete Cloud Identity Provider configuration.

Delete Cloud Identity Provider configuration.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Cloud Identity Provider identifier
 @return CloudAzureAPIV1CloudAzureIdDeleteRequest
*/
func (a *CloudAzureAPIService) V1CloudAzureIdDelete(ctx context.Context, id string) CloudAzureAPIV1CloudAzureIdDeleteRequest {
	return CloudAzureAPIV1CloudAzureIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *CloudAzureAPIService) V1CloudAzureIdDeleteExecute(r CloudAzureAPIV1CloudAzureIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudAzureAPIService.V1CloudAzureIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/cloud-azure/{id}"
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
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

type CloudAzureAPIV1CloudAzureIdGetRequest struct {
	ctx context.Context
	ApiService CloudAzureAPI
	id string
}

func (r CloudAzureAPIV1CloudAzureIdGetRequest) Execute() (*AzureConfiguration, *http.Response, error) {
	return r.ApiService.V1CloudAzureIdGetExecute(r)
}

/*
V1CloudAzureIdGet Get Azure Cloud Identity Provider configuration with given ID.

Get Azure Cloud Identity Provider configuration with given ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Cloud Identity Provider identifier
 @return CloudAzureAPIV1CloudAzureIdGetRequest
*/
func (a *CloudAzureAPIService) V1CloudAzureIdGet(ctx context.Context, id string) CloudAzureAPIV1CloudAzureIdGetRequest {
	return CloudAzureAPIV1CloudAzureIdGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return AzureConfiguration
func (a *CloudAzureAPIService) V1CloudAzureIdGetExecute(r CloudAzureAPIV1CloudAzureIdGetRequest) (*AzureConfiguration, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AzureConfiguration
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudAzureAPIService.V1CloudAzureIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/cloud-azure/{id}"
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
			return localVarReturnValue, localVarHTTPResponse, newErr
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

type CloudAzureAPIV1CloudAzureIdPutRequest struct {
	ctx context.Context
	ApiService CloudAzureAPI
	id string
	azureConfigurationUpdate *AzureConfigurationUpdate
}

// Azure Cloud Identity Provider configuration to update
func (r CloudAzureAPIV1CloudAzureIdPutRequest) AzureConfigurationUpdate(azureConfigurationUpdate AzureConfigurationUpdate) CloudAzureAPIV1CloudAzureIdPutRequest {
	r.azureConfigurationUpdate = &azureConfigurationUpdate
	return r
}

func (r CloudAzureAPIV1CloudAzureIdPutRequest) Execute() (*AzureConfiguration, *http.Response, error) {
	return r.ApiService.V1CloudAzureIdPutExecute(r)
}

/*
V1CloudAzureIdPut Update Azure Cloud Identity Provider configuration

Update Azure Cloud Identity Provider configuration. Cannot be used for partial updates, all content body must be sent.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Cloud Identity Provider identifier
 @return CloudAzureAPIV1CloudAzureIdPutRequest
*/
func (a *CloudAzureAPIService) V1CloudAzureIdPut(ctx context.Context, id string) CloudAzureAPIV1CloudAzureIdPutRequest {
	return CloudAzureAPIV1CloudAzureIdPutRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return AzureConfiguration
func (a *CloudAzureAPIService) V1CloudAzureIdPutExecute(r CloudAzureAPIV1CloudAzureIdPutRequest) (*AzureConfiguration, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AzureConfiguration
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudAzureAPIService.V1CloudAzureIdPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/cloud-azure/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.azureConfigurationUpdate == nil {
		return localVarReturnValue, nil, reportError("azureConfigurationUpdate is required and must be specified")
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
	localVarPostBody = r.azureConfigurationUpdate
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
		if localVarHTTPResponse.StatusCode == 404 {
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
		if localVarHTTPResponse.StatusCode == 409 {
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

type CloudAzureAPIV1CloudAzurePostRequest struct {
	ctx context.Context
	ApiService CloudAzureAPI
	azureConfigurationRequest *AzureConfigurationRequest
}

// Azure Cloud Identity Provider configuration to create
func (r CloudAzureAPIV1CloudAzurePostRequest) AzureConfigurationRequest(azureConfigurationRequest AzureConfigurationRequest) CloudAzureAPIV1CloudAzurePostRequest {
	r.azureConfigurationRequest = &azureConfigurationRequest
	return r
}

func (r CloudAzureAPIV1CloudAzurePostRequest) Execute() (*HrefResponse, *http.Response, error) {
	return r.ApiService.V1CloudAzurePostExecute(r)
}

/*
V1CloudAzurePost Create Azure Cloud Identity Provider configuration

Create new Azure Cloud Identity Provider configuration with unique display name.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return CloudAzureAPIV1CloudAzurePostRequest
*/
func (a *CloudAzureAPIService) V1CloudAzurePost(ctx context.Context) CloudAzureAPIV1CloudAzurePostRequest {
	return CloudAzureAPIV1CloudAzurePostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return HrefResponse
func (a *CloudAzureAPIService) V1CloudAzurePostExecute(r CloudAzureAPIV1CloudAzurePostRequest) (*HrefResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *HrefResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudAzureAPIService.V1CloudAzurePost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/cloud-azure"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.azureConfigurationRequest == nil {
		return localVarReturnValue, nil, reportError("azureConfigurationRequest is required and must be specified")
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
	localVarPostBody = r.azureConfigurationRequest
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
