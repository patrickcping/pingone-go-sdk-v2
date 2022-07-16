/*
PingOne Platform API - PingOne MFA

The PingOne Platform API covering the PingOne MFA service

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mfa

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// UsersEnableUsersMFAApiService UsersEnableUsersMFAApi service
type UsersEnableUsersMFAApiService service

type ApiV1EnvironmentsEnvIDUsersUserIDMfaEnabledGetRequest struct {
	ctx context.Context
	ApiService *UsersEnableUsersMFAApiService
	envID string
	userID string
}

func (r ApiV1EnvironmentsEnvIDUsersUserIDMfaEnabledGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDUsersUserIDMfaEnabledGetExecute(r)
}

/*
V1EnvironmentsEnvIDUsersUserIDMfaEnabledGet READ User MFA Enabled

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param userID
 @return ApiV1EnvironmentsEnvIDUsersUserIDMfaEnabledGetRequest
*/
func (a *UsersEnableUsersMFAApiService) V1EnvironmentsEnvIDUsersUserIDMfaEnabledGet(ctx context.Context, envID string, userID string) ApiV1EnvironmentsEnvIDUsersUserIDMfaEnabledGetRequest {
	return ApiV1EnvironmentsEnvIDUsersUserIDMfaEnabledGetRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		userID: userID,
	}
}

// Execute executes the request
func (a *UsersEnableUsersMFAApiService) V1EnvironmentsEnvIDUsersUserIDMfaEnabledGetExecute(r ApiV1EnvironmentsEnvIDUsersUserIDMfaEnabledGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersEnableUsersMFAApiService.V1EnvironmentsEnvIDUsersUserIDMfaEnabledGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/users/{userID}/mfaEnabled"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", url.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userID"+"}", url.PathEscape(parameterToString(r.userID, "")), -1)

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiV1EnvironmentsEnvIDUsersUserIDMfaEnabledPutRequest struct {
	ctx context.Context
	ApiService *UsersEnableUsersMFAApiService
	envID string
	userID string
	body *map[string]interface{}
}

func (r ApiV1EnvironmentsEnvIDUsersUserIDMfaEnabledPutRequest) Body(body map[string]interface{}) ApiV1EnvironmentsEnvIDUsersUserIDMfaEnabledPutRequest {
	r.body = &body
	return r
}

func (r ApiV1EnvironmentsEnvIDUsersUserIDMfaEnabledPutRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1EnvironmentsEnvIDUsersUserIDMfaEnabledPutExecute(r)
}

/*
V1EnvironmentsEnvIDUsersUserIDMfaEnabledPut UPDATE User MFA Enabled

By design, PingOne requests solely comprise this collection. For complete documentation, direct a browser to <a href='https://apidocs.pingidentity.com/pingone/platform/v1/api/'>apidocs.pingidentity.com</a>.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param userID
 @return ApiV1EnvironmentsEnvIDUsersUserIDMfaEnabledPutRequest
*/
func (a *UsersEnableUsersMFAApiService) V1EnvironmentsEnvIDUsersUserIDMfaEnabledPut(ctx context.Context, envID string, userID string) ApiV1EnvironmentsEnvIDUsersUserIDMfaEnabledPutRequest {
	return ApiV1EnvironmentsEnvIDUsersUserIDMfaEnabledPutRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		userID: userID,
	}
}

// Execute executes the request
func (a *UsersEnableUsersMFAApiService) V1EnvironmentsEnvIDUsersUserIDMfaEnabledPutExecute(r ApiV1EnvironmentsEnvIDUsersUserIDMfaEnabledPutRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersEnableUsersMFAApiService.V1EnvironmentsEnvIDUsersUserIDMfaEnabledPut")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{envID}/users/{userID}/mfaEnabled"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", url.PathEscape(parameterToString(r.envID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userID"+"}", url.PathEscape(parameterToString(r.userID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
