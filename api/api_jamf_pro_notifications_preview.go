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


type JamfProNotificationsPreviewAPI interface {

	/*
	NotificationsAlertsGet Get Notifications for user and site 

	Gets notifications for user and site


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return JamfProNotificationsPreviewAPINotificationsAlertsGetRequest

	Deprecated
	*/
	NotificationsAlertsGet(ctx context.Context) JamfProNotificationsPreviewAPINotificationsAlertsGetRequest

	// NotificationsAlertsGetExecute executes the request
	//  @return []Notification
	// Deprecated
	NotificationsAlertsGetExecute(r JamfProNotificationsPreviewAPINotificationsAlertsGetRequest) ([]Notification, *http.Response, error)

	/*
	NotificationsAlertsIdDelete DEPRECATED - USE \"alerts/{type}/{id}\" INSTEAD. Deletes only Patch Management notifications. 

	DEPRECATED - USE "alerts/{type}/{id}" INSTEAD. Deletes only Patch Management notifications.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id instance ID of the notification
	@return JamfProNotificationsPreviewAPINotificationsAlertsIdDeleteRequest

	Deprecated
	*/
	NotificationsAlertsIdDelete(ctx context.Context, id int64) JamfProNotificationsPreviewAPINotificationsAlertsIdDeleteRequest

	// NotificationsAlertsIdDeleteExecute executes the request
	// Deprecated
	NotificationsAlertsIdDeleteExecute(r JamfProNotificationsPreviewAPINotificationsAlertsIdDeleteRequest) (*http.Response, error)

	/*
	NotificationsAlertsTypeIdDelete Delete Notifications 

	Deletes notifications. Option for 'type' is 'PATCH_UPDATE'.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id instance ID of the notification
	@param type_ type of the notification
	@return JamfProNotificationsPreviewAPINotificationsAlertsTypeIdDeleteRequest

	Deprecated
	*/
	NotificationsAlertsTypeIdDelete(ctx context.Context, id int64, type_ NotificationType) JamfProNotificationsPreviewAPINotificationsAlertsTypeIdDeleteRequest

	// NotificationsAlertsTypeIdDeleteExecute executes the request
	// Deprecated
	NotificationsAlertsTypeIdDeleteExecute(r JamfProNotificationsPreviewAPINotificationsAlertsTypeIdDeleteRequest) (*http.Response, error)
}

// JamfProNotificationsPreviewAPIService JamfProNotificationsPreviewAPI service
type JamfProNotificationsPreviewAPIService service

type JamfProNotificationsPreviewAPINotificationsAlertsGetRequest struct {
	ctx context.Context
	ApiService JamfProNotificationsPreviewAPI
}

func (r JamfProNotificationsPreviewAPINotificationsAlertsGetRequest) Execute() ([]Notification, *http.Response, error) {
	return r.ApiService.NotificationsAlertsGetExecute(r)
}

/*
NotificationsAlertsGet Get Notifications for user and site 

Gets notifications for user and site


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return JamfProNotificationsPreviewAPINotificationsAlertsGetRequest

Deprecated
*/
func (a *JamfProNotificationsPreviewAPIService) NotificationsAlertsGet(ctx context.Context) JamfProNotificationsPreviewAPINotificationsAlertsGetRequest {
	return JamfProNotificationsPreviewAPINotificationsAlertsGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []Notification
// Deprecated
func (a *JamfProNotificationsPreviewAPIService) NotificationsAlertsGetExecute(r JamfProNotificationsPreviewAPINotificationsAlertsGetRequest) ([]Notification, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Notification
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JamfProNotificationsPreviewAPIService.NotificationsAlertsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/notifications/alerts"

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

type JamfProNotificationsPreviewAPINotificationsAlertsIdDeleteRequest struct {
	ctx context.Context
	ApiService JamfProNotificationsPreviewAPI
	id int64
}

func (r JamfProNotificationsPreviewAPINotificationsAlertsIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.NotificationsAlertsIdDeleteExecute(r)
}

/*
NotificationsAlertsIdDelete DEPRECATED - USE \"alerts/{type}/{id}\" INSTEAD. Deletes only Patch Management notifications. 

DEPRECATED - USE "alerts/{type}/{id}" INSTEAD. Deletes only Patch Management notifications.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id instance ID of the notification
 @return JamfProNotificationsPreviewAPINotificationsAlertsIdDeleteRequest

Deprecated
*/
func (a *JamfProNotificationsPreviewAPIService) NotificationsAlertsIdDelete(ctx context.Context, id int64) JamfProNotificationsPreviewAPINotificationsAlertsIdDeleteRequest {
	return JamfProNotificationsPreviewAPINotificationsAlertsIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
// Deprecated
func (a *JamfProNotificationsPreviewAPIService) NotificationsAlertsIdDeleteExecute(r JamfProNotificationsPreviewAPINotificationsAlertsIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JamfProNotificationsPreviewAPIService.NotificationsAlertsIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/notifications/alerts/{id}"
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

type JamfProNotificationsPreviewAPINotificationsAlertsTypeIdDeleteRequest struct {
	ctx context.Context
	ApiService JamfProNotificationsPreviewAPI
	id int64
	type_ NotificationType
}

func (r JamfProNotificationsPreviewAPINotificationsAlertsTypeIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.NotificationsAlertsTypeIdDeleteExecute(r)
}

/*
NotificationsAlertsTypeIdDelete Delete Notifications 

Deletes notifications. Option for 'type' is 'PATCH_UPDATE'.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id instance ID of the notification
 @param type_ type of the notification
 @return JamfProNotificationsPreviewAPINotificationsAlertsTypeIdDeleteRequest

Deprecated
*/
func (a *JamfProNotificationsPreviewAPIService) NotificationsAlertsTypeIdDelete(ctx context.Context, id int64, type_ NotificationType) JamfProNotificationsPreviewAPINotificationsAlertsTypeIdDeleteRequest {
	return JamfProNotificationsPreviewAPINotificationsAlertsTypeIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		type_: type_,
	}
}

// Execute executes the request
// Deprecated
func (a *JamfProNotificationsPreviewAPIService) NotificationsAlertsTypeIdDeleteExecute(r JamfProNotificationsPreviewAPINotificationsAlertsTypeIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JamfProNotificationsPreviewAPIService.NotificationsAlertsTypeIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/notifications/alerts/{type}/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"type"+"}", url.PathEscape(parameterValueToString(r.type_, "type_")), -1)

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
