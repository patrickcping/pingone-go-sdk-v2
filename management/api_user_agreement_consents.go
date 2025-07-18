/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"bytes"
	"context"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
)

// UserAgreementConsentsApiService UserAgreementConsentsApi service
type UserAgreementConsentsApiService service

type ApiEnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDGetRequest struct {
	ctx                        context.Context
	ApiService                 *UserAgreementConsentsApiService
	environmentID              string
	userID                     string
	agreementID                string
	xPingExternalTransactionID *string
	xPingExternalSessionID     *string
}

// An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiEnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDGetRequest) XPingExternalTransactionID(xPingExternalTransactionID string) ApiEnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDGetRequest {
	r.xPingExternalTransactionID = &xPingExternalTransactionID
	return r
}

// An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiEnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDGetRequest) XPingExternalSessionID(xPingExternalSessionID string) ApiEnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDGetRequest {
	r.xPingExternalSessionID = &xPingExternalSessionID
	return r
}

func (r ApiEnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.EnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDGetExecute(r)
}

/*
EnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDGet READ One User Agreement Consent

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param environmentID
	@param userID
	@param agreementID
	@return ApiEnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDGetRequest
*/
func (a *UserAgreementConsentsApiService) EnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDGet(ctx context.Context, environmentID string, userID string, agreementID string) ApiEnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDGetRequest {
	return ApiEnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDGetRequest{
		ApiService:    a,
		ctx:           ctx,
		environmentID: environmentID,
		userID:        userID,
		agreementID:   agreementID,
	}
}

// Execute executes the request
func (a *UserAgreementConsentsApiService) EnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDGetExecute(r ApiEnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDGetRequest) (*http.Response, error) {
	var (
		err      error
		response *http.Response
	)
	
	response, err = processResponse(
		func() (any, *http.Response, error) {
			resp, err := r.ApiService.internalEnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDGetExecute(r)
			return nil, resp, err
		},
		nil,
	)
	return response, err
}

func (a *UserAgreementConsentsApiService) internalEnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDGetExecute(r ApiEnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserAgreementConsentsApiService.EnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environments/{environmentID}/users/{userID}/agreementConsents/{agreementID}"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentID"+"}", url.PathEscape(parameterValueToString(r.environmentID, "environmentID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userID"+"}", url.PathEscape(parameterValueToString(r.userID, "userID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"agreementID"+"}", url.PathEscape(parameterValueToString(r.agreementID, "agreementID")), -1)

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
	if r.xPingExternalTransactionID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Ping-External-Transaction-ID", r.xPingExternalTransactionID, "")
	}
	if r.xPingExternalSessionID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Ping-External-Session-ID", r.xPingExternalSessionID, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	var bodyBytes []byte
	if req.Body != nil {
		bodyBytes, _ = io.ReadAll(req.Body)
	}

	var localVarHTTPResponse *http.Response
	var localVarBody []byte

	for i := range maxRetries {
		if req.Body != nil {
			req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		if i > 0 {
			slog.Debug("Retrying request", "attempt", i, "method", localVarHTTPMethod, "path", localVarPath)
		}

		localVarHTTPResponse, err = a.client.callAPI(req)
		if err != nil || localVarHTTPResponse == nil {
			return localVarHTTPResponse, err
		}

		localVarBody, err = io.ReadAll(localVarHTTPResponse.Body)
		_ = localVarHTTPResponse.Body.Close()
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
				var v P1Error
				err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
				if err != nil {
					newErr.error = err.Error()
					return localVarHTTPResponse, newErr
				}
				newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
				newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
				newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
				newErr.model = v
				return localVarHTTPResponse, newErr
			}
			if localVarHTTPResponse.StatusCode == 404 {
				var v P1Error
				err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
				if err != nil {
					newErr.error = err.Error()
					return localVarHTTPResponse, newErr
				}
				newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
				newErr.model = v
				return localVarHTTPResponse, newErr
			}
			if localVarHTTPResponse.StatusCode == 409 {
				var v P1Error
				err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
				if err != nil {
					newErr.error = err.Error()
					return localVarHTTPResponse, newErr
				}
				newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
				newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
				newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
				newErr.model = v
				return localVarHTTPResponse, newErr
			}
			return localVarHTTPResponse, newErr
		}
		break
	}

	return localVarHTTPResponse, nil
}

type ApiEnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDPostRequest struct {
	ctx                        context.Context
	ApiService                 *UserAgreementConsentsApiService
	environmentID              string
	userID                     string
	agreementID                string
	contentType                *string
	xPingExternalTransactionID *string
	xPingExternalSessionID     *string
}

func (r ApiEnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDPostRequest) ContentType(contentType string) ApiEnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDPostRequest {
	r.contentType = &contentType
	return r
}

// An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiEnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDPostRequest) XPingExternalTransactionID(xPingExternalTransactionID string) ApiEnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDPostRequest {
	r.xPingExternalTransactionID = &xPingExternalTransactionID
	return r
}

// An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiEnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDPostRequest) XPingExternalSessionID(xPingExternalSessionID string) ApiEnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDPostRequest {
	r.xPingExternalSessionID = &xPingExternalSessionID
	return r
}

func (r ApiEnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.EnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDPostExecute(r)
}

/*
EnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDPost Revoke Agreement

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param environmentID
	@param userID
	@param agreementID
	@return ApiEnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDPostRequest
*/
func (a *UserAgreementConsentsApiService) EnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDPost(ctx context.Context, environmentID string, userID string, agreementID string) ApiEnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDPostRequest {
	return ApiEnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDPostRequest{
		ApiService:    a,
		ctx:           ctx,
		environmentID: environmentID,
		userID:        userID,
		agreementID:   agreementID,
	}
}

// Execute executes the request
func (a *UserAgreementConsentsApiService) EnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDPostExecute(r ApiEnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDPostRequest) (*http.Response, error) {
	var (
		err      error
		response *http.Response
	)
	
	response, err = processResponse(
		func() (any, *http.Response, error) {
			resp, err := r.ApiService.internalEnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDPostExecute(r)
			return nil, resp, err
		},
		nil,
	)
	return response, err
}

func (a *UserAgreementConsentsApiService) internalEnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDPostExecute(r ApiEnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserAgreementConsentsApiService.EnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environments/{environmentID}/users/{userID}/agreementConsents/{agreementID}"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentID"+"}", url.PathEscape(parameterValueToString(r.environmentID, "environmentID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userID"+"}", url.PathEscape(parameterValueToString(r.userID, "userID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"agreementID"+"}", url.PathEscape(parameterValueToString(r.agreementID, "agreementID")), -1)

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
	if r.contentType != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Content-Type", r.contentType, "")
	}
	if r.xPingExternalTransactionID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Ping-External-Transaction-ID", r.xPingExternalTransactionID, "")
	}
	if r.xPingExternalSessionID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Ping-External-Session-ID", r.xPingExternalSessionID, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	var bodyBytes []byte
	if req.Body != nil {
		bodyBytes, _ = io.ReadAll(req.Body)
	}

	var localVarHTTPResponse *http.Response
	var localVarBody []byte

	for i := range maxRetries {
		if req.Body != nil {
			req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		if i > 0 {
			slog.Debug("Retrying request", "attempt", i, "method", localVarHTTPMethod, "path", localVarPath)
		}

		localVarHTTPResponse, err = a.client.callAPI(req)
		if err != nil || localVarHTTPResponse == nil {
			return localVarHTTPResponse, err
		}

		localVarBody, err = io.ReadAll(localVarHTTPResponse.Body)
		_ = localVarHTTPResponse.Body.Close()
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
				var v P1Error
				err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
				if err != nil {
					newErr.error = err.Error()
					return localVarHTTPResponse, newErr
				}
				newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
				newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
				newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
				newErr.model = v
				return localVarHTTPResponse, newErr
			}
			if localVarHTTPResponse.StatusCode == 404 {
				var v P1Error
				err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
				if err != nil {
					newErr.error = err.Error()
					return localVarHTTPResponse, newErr
				}
				newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
				newErr.model = v
				return localVarHTTPResponse, newErr
			}
			if localVarHTTPResponse.StatusCode == 405 {
				var v P1Error
				err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
				if err != nil {
					newErr.error = err.Error()
					return localVarHTTPResponse, newErr
				}
				newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
				newErr.model = v
				return localVarHTTPResponse, newErr
			}
			if localVarHTTPResponse.StatusCode == 409 {
				var v P1Error
				err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
				if err != nil {
					newErr.error = err.Error()
					return localVarHTTPResponse, newErr
				}
				newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
				newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
				newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
				newErr.model = v
				return localVarHTTPResponse, newErr
			}
			return localVarHTTPResponse, newErr
		}
		break
	}

	return localVarHTTPResponse, nil
}

type ApiEnvironmentsEnvironmentIDUsersUserIDAgreementConsentsGetRequest struct {
	ctx                        context.Context
	ApiService                 *UserAgreementConsentsApiService
	environmentID              string
	userID                     string
	xPingExternalTransactionID *string
	xPingExternalSessionID     *string
}

// An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiEnvironmentsEnvironmentIDUsersUserIDAgreementConsentsGetRequest) XPingExternalTransactionID(xPingExternalTransactionID string) ApiEnvironmentsEnvironmentIDUsersUserIDAgreementConsentsGetRequest {
	r.xPingExternalTransactionID = &xPingExternalTransactionID
	return r
}

// An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiEnvironmentsEnvironmentIDUsersUserIDAgreementConsentsGetRequest) XPingExternalSessionID(xPingExternalSessionID string) ApiEnvironmentsEnvironmentIDUsersUserIDAgreementConsentsGetRequest {
	r.xPingExternalSessionID = &xPingExternalSessionID
	return r
}

func (r ApiEnvironmentsEnvironmentIDUsersUserIDAgreementConsentsGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.EnvironmentsEnvironmentIDUsersUserIDAgreementConsentsGetExecute(r)
}

/*
EnvironmentsEnvironmentIDUsersUserIDAgreementConsentsGet READ All User Agreement Consents

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param environmentID
	@param userID
	@return ApiEnvironmentsEnvironmentIDUsersUserIDAgreementConsentsGetRequest
*/
func (a *UserAgreementConsentsApiService) EnvironmentsEnvironmentIDUsersUserIDAgreementConsentsGet(ctx context.Context, environmentID string, userID string) ApiEnvironmentsEnvironmentIDUsersUserIDAgreementConsentsGetRequest {
	return ApiEnvironmentsEnvironmentIDUsersUserIDAgreementConsentsGetRequest{
		ApiService:    a,
		ctx:           ctx,
		environmentID: environmentID,
		userID:        userID,
	}
}

// Execute executes the request
func (a *UserAgreementConsentsApiService) EnvironmentsEnvironmentIDUsersUserIDAgreementConsentsGetExecute(r ApiEnvironmentsEnvironmentIDUsersUserIDAgreementConsentsGetRequest) (*http.Response, error) {
	var (
		err      error
		response *http.Response
	)
	
	response, err = processResponse(
		func() (any, *http.Response, error) {
			resp, err := r.ApiService.internalEnvironmentsEnvironmentIDUsersUserIDAgreementConsentsGetExecute(r)
			return nil, resp, err
		},
		nil,
	)
	return response, err
}

func (a *UserAgreementConsentsApiService) internalEnvironmentsEnvironmentIDUsersUserIDAgreementConsentsGetExecute(r ApiEnvironmentsEnvironmentIDUsersUserIDAgreementConsentsGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserAgreementConsentsApiService.EnvironmentsEnvironmentIDUsersUserIDAgreementConsentsGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environments/{environmentID}/users/{userID}/agreementConsents"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentID"+"}", url.PathEscape(parameterValueToString(r.environmentID, "environmentID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userID"+"}", url.PathEscape(parameterValueToString(r.userID, "userID")), -1)

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
	if r.xPingExternalTransactionID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Ping-External-Transaction-ID", r.xPingExternalTransactionID, "")
	}
	if r.xPingExternalSessionID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Ping-External-Session-ID", r.xPingExternalSessionID, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	var bodyBytes []byte
	if req.Body != nil {
		bodyBytes, _ = io.ReadAll(req.Body)
	}

	var localVarHTTPResponse *http.Response
	var localVarBody []byte

	for i := range maxRetries {
		if req.Body != nil {
			req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		if i > 0 {
			slog.Debug("Retrying request", "attempt", i, "method", localVarHTTPMethod, "path", localVarPath)
		}

		localVarHTTPResponse, err = a.client.callAPI(req)
		if err != nil || localVarHTTPResponse == nil {
			return localVarHTTPResponse, err
		}

		localVarBody, err = io.ReadAll(localVarHTTPResponse.Body)
		_ = localVarHTTPResponse.Body.Close()
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
				var v P1Error
				err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
				if err != nil {
					newErr.error = err.Error()
					return localVarHTTPResponse, newErr
				}
				newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
				newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
				newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
				newErr.model = v
				return localVarHTTPResponse, newErr
			}
			if localVarHTTPResponse.StatusCode == 404 {
				var v P1Error
				err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
				if err != nil {
					newErr.error = err.Error()
					return localVarHTTPResponse, newErr
				}
				newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
				newErr.model = v
				return localVarHTTPResponse, newErr
			}
			if localVarHTTPResponse.StatusCode == 409 {
				var v P1Error
				err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
				if err != nil {
					newErr.error = err.Error()
					return localVarHTTPResponse, newErr
				}
				newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
				newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
				newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
				newErr.model = v
				return localVarHTTPResponse, newErr
			}
			return localVarHTTPResponse, newErr
		}
		break
	}

	return localVarHTTPResponse, nil
}
